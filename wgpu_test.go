package wgpu_test

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/bluescreen10/dawn-go/wgpu"
)

func init() {
	runtime.LockOSThread()
}

func TestTriangleRendering(t *testing.T) {
	ctx, err := createTestContext()
	if err != nil {
		t.Fatalf("error creating test context: %v", err)
	}
	defer ctx.CleanUp()

	const width, height = 256, 256
	format := wgpu.TextureFormatRGBA8Unorm

	texture := ctx.device.CreateTexture(&wgpu.TextureDescriptor{
		Size:          wgpu.Extent3D{Width: width, Height: height, DepthOrArrayLayers: 1},
		Format:        format,
		Usage:         wgpu.TextureUsageRenderAttachment | wgpu.TextureUsageCopySrc,
		SampleCount:   1,
		MipLevelCount: 1,
	})
	defer texture.Destroy()

	shader := ctx.device.CreateShaderModule(wgpu.ShaderModuleDescriptor{
		WGSLSource: &wgpu.ShaderSourceWGSL{
			Code: `
			@vertex
			fn vs_main(@builtin(vertex_index) in_vertex_index: u32) -> @builtin(position) vec4<f32> {
				var pos = array<vec2<f32>, 3>(
					vec2<f32>(0.0, 0.5),
					vec2<f32>(-0.5, -0.5),
					vec2<f32>(0.5, -0.5)
				);
				return vec4<f32>(pos[in_vertex_index], 0.0, 1.0);
			}

			@fragment
			fn fs_main() -> @location(0) vec4<f32> {
				return vec4<f32>(1.0, 0.0, 0.0, 1.0);
			}`,
		},
	})
	defer shader.Release()

	pipeline := ctx.device.CreateRenderPipeline(wgpu.RenderPipelineDescriptor{
		Vertex: wgpu.VertexState{Module: shader, EntryPoint: "vs_main"},
		Fragment: &wgpu.FragmentState{
			Module: shader, EntryPoint: "fs_main",
			Targets: []wgpu.ColorTargetState{
				{
					Format:    format,
					WriteMask: wgpu.ColorWriteMaskAll,
				}},
		},
		Primitive: wgpu.PrimitiveState{Topology: wgpu.PrimitiveTopologyTriangleList},
		Multisample: wgpu.MultisampleState{
			Count:                  1,
			Mask:                   0xFFFFFFFF,
			AlphaToCoverageEnabled: false,
		},
	})
	defer pipeline.Release()

	encoder := ctx.device.CreateCommandEncoder(nil)
	view := texture.CreateView(nil)

	pass := encoder.BeginRenderPass(wgpu.RenderPassDescriptor{
		ColorAttachments: []wgpu.RenderPassColorAttachment{{
			View: view, LoadOp: wgpu.LoadOpClear, StoreOp: wgpu.StoreOpStore,
			ClearValue: wgpu.Color{R: 0, G: 0, B: 0, A: 1},
		}},
	})

	pass.SetPipeline(pipeline)
	pass.Draw(3, 1, 0, 0)
	pass.End()

	bytesPerPixel := uint32(4)
	paddedBytesPerRow := (width*bytesPerPixel + 255) &^ 255
	bufferSize := paddedBytesPerRow * height

	readbackBuffer := ctx.device.CreateBuffer(wgpu.BufferDescriptor{
		Size:  uint64(bufferSize),
		Usage: wgpu.BufferUsageCopyDst | wgpu.BufferUsageMapRead,
	})

	encoder.CopyTextureToBuffer(
		texture.AsImageCopy(),
		readbackBuffer.AsImageCopyBuffer(paddedBytesPerRow, height),
		wgpu.Extent3D{Width: width, Height: height, DepthOrArrayLayers: 1},
	)

	ctx.device.GetQueue().Submit(encoder.Finish(nil))

	future := readbackBuffer.MapAsync(wgpu.MapModeRead, 0, int(bufferSize), func(status wgpu.MapAsyncStatus, message string) {
		if status != wgpu.MapAsyncStatusSuccess {
			t.Fatalf("MapAsync failed: %s", message)
		}
	})

	err = ctx.instance.WaitAny([]wgpu.Future{future}, uint64(10*time.Second))
	if err != nil {
		t.Fatalf("error waiting futures: %v", err)
	}

	pixels := readbackBuffer.GetConstMappedRange(0, int(bufferSize))
	defer readbackBuffer.Unmap()

	expected, err := loadImage("testdata/triangle.png")
	if err != nil {
		t.Fatalf("error loading image: %v", err)
	}
	got := createImageFromPixels(pixels, width, height)

	// Uncomment the following lines to update the golden copy
	// err = saveImage("testdata/triangle.png", got)
	// if err != nil {
	// 	t.Fatalf("error saving file: %v", err)
	// }

	imagesAreEqual(t, expected, got)
}

type wgpuContext struct {
	instance *wgpu.Instance
	adapter  *wgpu.Adapter
	device   *wgpu.Device
}

func (c *wgpuContext) CleanUp() {
	c.device.Destroy()
	c.adapter.Release()
	c.instance.Release()
}

func createTestContext() (*wgpuContext, error) {
	var ctx wgpuContext
	var err error

	ctx.instance = wgpu.CreateInstance(&wgpu.InstanceDescriptor{
		RequiredFeatures: []wgpu.InstanceFeatureName{
			wgpu.InstanceFeatureNameTimedWaitAny,
		},
	})

	ctx.adapter, err = ctx.instance.RequestAdapter(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get adapter: %v", err)
	}

	ctx.device = ctx.adapter.RequestDevice(&wgpu.DeviceDescriptor{
		UncapturedErrorCallback: func(device *wgpu.Device, typ wgpu.ErrorType, message string) {
			panic(fmt.Sprintf("%s error: %s", typ, message))
		},
	})

	return &ctx, nil
}

func loadImage(path string) (image.Image, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(r)
	return img, err
}

func saveImage(path string, img image.Image) error {

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}

func createImageFromPixels(pixels []byte, width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	bytesPerPixel := 4
	paddedBytesPerRow := (width*bytesPerPixel + 255) &^ 255

	for y := 0; y < height; y++ {
		rowStart := y * int(paddedBytesPerRow)
		copy(img.Pix[y*width*4:(y+1)*width*4], pixels[rowStart:rowStart+width*4])
	}

	return img
}

func imagesAreEqual(t *testing.T, expected, got image.Image) {
	boundsA := expected.Bounds()
	boundsB := got.Bounds()
	if boundsA != boundsB {
		t.Errorf("bounds error, expected %v got %v", expected.Bounds(), got.Bounds())
		return
	}

	for y := boundsA.Min.Y; y < boundsA.Max.Y; y++ {
		for x := boundsA.Min.X; x < boundsA.Max.X; x++ {
			if expected.At(x, y) != got.At(x, y) {
				eR, eG, eB, eA := expected.At(x, y).RGBA()
				gR, gG, gB, gA := got.At(x, y).RGBA()
				t.Errorf("pixel error, expected (%d, %d, %d, %d) got (%d, %d, %d, %d)", eR, eG, eB, eA, gR, gG, gB, gA)
				return
			}
		}
	}
}
