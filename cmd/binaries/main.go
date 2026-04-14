package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Configuration from Env and Flags
type Config struct {
	R2AccountID  string
	R2AccessKey  string
	R2SecretKey  string
	R2Bucket     string
	ModuleDomain string
	NewVersion   string
}

type ReleaseAsset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

type GithubRelease struct {
	TagName string         `json:"tag_name"`
	Assets  []ReleaseAsset `json:"assets"`
}

type Target struct {
	Pattern     string
	FileMapping map[string]string
}

func main() {
	versionFlag := flag.String("ver", os.Getenv("NEW_VERSION"), "version to release (e.g. 0.0.1)")
	flag.Parse()

	cfg := Config{
		R2AccountID:  os.Getenv("R2_ACCOUNT_ID"),
		R2AccessKey:  os.Getenv("R2_ACCESS_KEY"),
		R2SecretKey:  os.Getenv("R2_SECRET_KEY"),
		R2Bucket:     os.Getenv("R2_BUCKET"),
		ModuleDomain: os.Getenv("MODULE_DOMAIN"),
		NewVersion:   *versionFlag,
	}

	if cfg.NewVersion == "" {
		log.Fatal("Version must be specified via -ver flag or NEW_VERSION env var")
	}

	// Ensure version has 'v' prefix for Go modules
	if !strings.HasPrefix(cfg.NewVersion, "v") {
		cfg.NewVersion = "v" + cfg.NewVersion
	}

	fmt.Printf("Fetching latest Dawn release for version %s...\n", cfg.NewVersion)
	release, err := getLatestDawnRelease()
	if err != nil {
		log.Fatalf("Failed to fetch release: %v", err)
	}

	targets := map[string][]Target{
		"linux": {
			{
				Pattern: "ubuntu-latest-Release",
				FileMapping: map[string]string{
					"lib64/libwebgpu_dawn.a": "amd64/libwebgpu_dawn.a",
				},
			},
		},
		"windows": {
			{
				Pattern: "windows-latest-Release",
				FileMapping: map[string]string{
					"lib/webgpu_dawn.lib": "amd64/webgpu_dawn.lib",
				},
			},
		},
		"darwin": {
			{
				Pattern: "macos-latest-Release",
				FileMapping: map[string]string{
					"lib/libwebgpu_dawn.a": "arm64/libwebgpu_dawn.a",
				},
			},
			{
				Pattern: "macos-15-intel-Release",
				FileMapping: map[string]string{
					"lib/libwebgpu_dawn.a": "amd64/libwebgpu_dawn.a",
				},
			},
		},
		"android": {
			{
				Pattern: "dawn-android",
				FileMapping: map[string]string{
					"x86_64/libwebgpu_dawn.a":      "amd64/libwebgpu_dawn.a",
					"x86/libwebgpu_dawn.a":         "386/libwebgpu_dawn.a",
					"arm64-v8a/libwebgpu_dawn.a":   "arm64/libwebgpu_dawn.a",
					"armeabi-v7a/libwebgpu_dawn.a": "arm/libwebgpu_dawn.a",
				},
			},
		},
	}

	s3Client := getR2Client(cfg)

	for os, targets := range targets {
		moduleSuffix := "wgpu-" + os
		fmt.Printf("Processing %s...\n", moduleSuffix)

		var buf bytes.Buffer
		fullModuleName := fmt.Sprintf("%s/%s", cfg.ModuleDomain, moduleSuffix)
		prefix := fmt.Sprintf("%s@%s/", fullModuleName, cfg.NewVersion)

		zw := zip.NewWriter(&buf)

		for _, target := range targets {
			for _, asset := range release.Assets {
				if strings.Contains(asset.Name, target.Pattern) {
					fmt.Printf("  Downloading %s...\n", asset.Name)
					if err := downloadAndExtractToZip(asset.BrowserDownloadURL, zw, prefix, target.FileMapping); err != nil {
						log.Printf("    Error processing asset %s: %v", asset.Name, err)
						continue
					}
				}
			}
		}

		addGoModToZip(zw, prefix, moduleSuffix)
		addLibGoToZip(zw, prefix, os)
		zw.Close()

		uploadModuleToR2(s3Client, cfg, moduleSuffix, buf.Bytes())
		fmt.Printf("  %s uploaded successfully.\n", moduleSuffix)
	}
}

func downloadAndExtractToZip(url string, zw *zip.Writer, prefix string, fileToCopy map[string]string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download asset: %s", resp.Status)
	}

	gr, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		dir := filepath.Base(filepath.Dir(header.Name))
		base := filepath.Base(header.Name)
		file := filepath.Join(dir, base)

		dst, ok := fileToCopy[file]
		if !ok {
			continue
		}

		zipPath := filepath.Join(prefix, dst)

		f, err := zw.Create(zipPath)
		if err != nil {
			return err
		}
		if _, err := io.Copy(f, tr); err != nil {
			return err
		}
	}
	return nil
}

func addGoModToZip(zw *zip.Writer, prefix, module string) {
	f, _ := zw.Create(prefix + "go.mod")
	content := fmt.Sprintf("module %s\n\ngo 1.23\n", module)
	f.Write([]byte(content))
}

func addLibGoToZip(zw *zip.Writer, prefix string, osName string) {
	f, _ := zw.Create(prefix + "lib.go")

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("package wgpu_%s\n\n", osName))
	sb.WriteString("/*\n")

	switch osName {
	case "darwin":
		sb.WriteString(fmt.Sprintf("#cgo %s,amd64 LDFLAGS: -L${SRCDIR}/amd64\n", osName))
		sb.WriteString(fmt.Sprintf("#cgo %s,arm64 LDFLAGS: -L${SRCDIR}/arm64\n", osName))
		sb.WriteString("#cgo darwin LDFLAGS: -framework Metal -framework IOKit -framework QuartzCore -framework Foundation -framework IOSurface -lc++ -lwebgpu_dawn\n")
	case "linux":
		sb.WriteString(fmt.Sprintf("#cgo %s,amd64 LDFLAGS: -L${SRCDIR}/amd64\n", osName))
		sb.WriteString("#cgo linux LDFLAGS: -lwebgpu_dawn -lstdc++ -lm -lpthread -ldl\n")
	case "windows":
		sb.WriteString(fmt.Sprintf("#cgo %s,amd64 LDFLAGS: -L${SRCDIR}/amd64\n", osName))
		sb.WriteString("#cgo windows LDFLAGS: -lwebgpu_dawn -ld3d12 -ldxgi -ld3dcompiler\n")
	case "android":
		sb.WriteString(fmt.Sprintf("#cgo %s,amd64 LDFLAGS: -L${SRCDIR}/amd64\n", osName))
		sb.WriteString(fmt.Sprintf("#cgo %s,386 LDFLAGS: -L${SRCDIR}/386\n", osName))
		sb.WriteString(fmt.Sprintf("#cgo %s,arm64 LDFLAGS: -L${SRCDIR}/arm64\n", osName))
		sb.WriteString(fmt.Sprintf("#cgo %s,arm LDFLAGS: -L${SRCDIR}/arm\n", osName))
		sb.WriteString("#cgo android LDFLAGS: -lwebgpu_dawn -landroid -lvulkan -lm -llog\n")
	}

	sb.WriteString("*/\nimport \"C\"\n")
	f.Write([]byte(sb.String()))
}

func uploadModuleToR2(client *s3.Client, cfg Config, moduleSuffix string, zipData []byte) {
	ctx := context.TODO()
	prefix := fmt.Sprintf("%s/@v", moduleSuffix)

	// 1. Upload ZIP
	_, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(cfg.R2Bucket),
		Key:         aws.String(fmt.Sprintf("%s/%s.zip", prefix, cfg.NewVersion)),
		Body:        bytes.NewReader(zipData),
		ContentType: aws.String("application/zip"),
	})
	if err != nil {
		log.Printf("Failed to upload ZIP for %s: %v", moduleSuffix, err)
		return
	}

	// 2. Upload .mod
	modContent := fmt.Sprintf("module %s/%s\n\ngo 1.23\n", cfg.ModuleDomain, moduleSuffix)
	client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(cfg.R2Bucket),
		Key:         aws.String(fmt.Sprintf("%s/%s.mod", prefix, cfg.NewVersion)),
		Body:        strings.NewReader(modContent),
		ContentType: aws.String("text/plain"),
	})

	// 3. Upload .info
	info := map[string]string{
		"Version": cfg.NewVersion,
		"Time":    time.Now().Format(time.RFC3339),
	}
	infoJSON, _ := json.Marshal(info)
	client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(cfg.R2Bucket),
		Key:         aws.String(fmt.Sprintf("%s/%s.info", prefix, cfg.NewVersion)),
		Body:        bytes.NewReader(infoJSON),
		ContentType: aws.String("application/json"),
	})

	// 4. Update List
	updateListFile(client, cfg.R2Bucket, prefix, cfg.NewVersion)
}

func updateListFile(client *s3.Client, bucket, prefix, version string) {
	key := prefix + "/list"
	existing, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	var listContent string
	if err == nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(existing.Body)
		listContent = buf.String()
	}

	if !strings.Contains(listContent, version) {
		listContent += version + "\n"
		client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket:      aws.String(bucket),
			Key:         aws.String(key),
			Body:        strings.NewReader(listContent),
			ContentType: aws.String("text/plain"),
		})
	}
}

func getLatestDawnRelease() (*GithubRelease, error) {
	req, _ := http.NewRequest("GET", "https://api.github.com/repos/google/dawn/releases/latest", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github API returned status: %s", resp.Status)
	}

	var release GithubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}
	return &release, nil
}

func getR2Client(cfg Config) *s3.Client {
	awsCfg, _ := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.R2AccessKey, cfg.R2SecretKey, "")),
		config.WithRegion("auto"),
	)

	return s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cfg.R2AccountID))
	})
}
