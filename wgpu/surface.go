//go:build !js

package wgpu

/*
#include <stdlib.h>
#include "webgpu.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type Surface struct {
	ref C.WGPUSurface
}

type SurfaceTexture struct {
	Texture *Texture
	Status  SurfaceGetCurrentTextureStatus
}

func (s *Surface) Configure(config SurfaceConfiguration) {

	cConfig := C.WGPUSurfaceConfiguration{
		device:      C.WGPUDevice(config.Device.ref),
		format:      C.WGPUTextureFormat(config.Format),
		usage:       C.WGPUTextureUsage(config.Usage),
		width:       C.uint32_t(config.Width),
		height:      C.uint32_t(config.Height),
		alphaMode:   C.WGPUCompositeAlphaMode(config.AlphaMode),
		presentMode: C.WGPUPresentMode(config.PresentMode),
	}

	if len(config.ViewFormats) > 0 {
		cConfig.viewFormats = (*C.WGPUTextureFormat)(unsafe.Pointer(&config.ViewFormats[0]))
		cConfig.viewFormatCount = C.size_t(len(config.ViewFormats))

		var pinner runtime.Pinner
		pinner.Pin(&config.ViewFormats[0])
		defer pinner.Unpin()
	}

	C.wgpuSurfaceConfigure(s.ref, &cConfig)
}

func (s *Surface) GetCapabilities(adapter *Adapter) (SurfaceCapabilities, error) {
	pAdapter := C.WGPUAdapter(unsafe.Pointer(adapter.ref))

	var cCapabilities C.WGPUSurfaceCapabilities

	status := C.wgpuSurfaceGetCapabilities(s.ref, pAdapter, &cCapabilities)
	defer C.wgpuSurfaceCapabilitiesFreeMembers(cCapabilities)

	if statusCode(status) != statusCodeSuccess {
		return SurfaceCapabilities{}, fmt.Errorf("error getting surface capabilities: %v", status)
	}

	capabilities := SurfaceCapabilities{
		Usages: TextureUsage(cCapabilities.usages),
	}

	if count := cCapabilities.formatCount; count > 0 {
		cFormats := unsafe.Slice((*C.WGPUTextureFormat)(cCapabilities.formats), count)
		capabilities.Formats = make([]TextureFormat, count)
		for i := range cFormats {
			capabilities.Formats[i] = TextureFormat(cFormats[i])
		}
	}

	if count := cCapabilities.presentModeCount; count > 0 {
		cPresentModes := unsafe.Slice((*C.WGPUPresentMode)(cCapabilities.presentModes), count)
		capabilities.PresentModes = make([]PresentMode, count)
		for i := range cPresentModes {
			capabilities.PresentModes[i] = PresentMode(cPresentModes[i])
		}
	}

	if count := cCapabilities.alphaModeCount; count > 0 {
		cAlphaModes := unsafe.Slice((*C.WGPUCompositeAlphaMode)(cCapabilities.alphaModes), count)
		capabilities.AlphaModes = make([]CompositeAlphaMode, count)
		for i := range cAlphaModes {
			capabilities.AlphaModes[i] = CompositeAlphaMode(cAlphaModes[i])
		}
	}

	return capabilities, nil
}

func (s *Surface) GetCurrentTexture() *SurfaceTexture {
	surfaceTexture, err := s.TryGetCurrentTexture()
	if err != nil {
		panic(err)
	}
	return surfaceTexture
}

func (s *Surface) TryGetCurrentTexture() (*SurfaceTexture, error) {

	var cSurfaceTexture C.WGPUSurfaceTexture
	C.wgpuSurfaceGetCurrentTexture(s.ref, &cSurfaceTexture)

	status := SurfaceGetCurrentTextureStatus(cSurfaceTexture.status)
	if status != SurfaceGetCurrentTextureStatusSuccessOptimal && status != SurfaceGetCurrentTextureStatusSuccessSuboptimal {
		return nil, fmt.Errorf("error getting current texture")
	}

	surfaceTexture := &SurfaceTexture{
		Texture: &Texture{ref: cSurfaceTexture.texture},
		Status:  status,
	}

	return surfaceTexture, nil
}

func (s *Surface) Release() {
	C.wgpuSurfaceRelease(s.ref)
}

func (s *SurfaceTexture) CreateView(descriptor *TextureViewDescriptor) *TextureView {
	return s.Texture.CreateView(descriptor)
}

func (s *Surface) Present() {
	err := s.TryPresent()
	if err != nil {
		panic(err)
	}
}

func (s *Surface) TryPresent() error {
	status := C.wgpuSurfacePresent(s.ref)

	if statusCode(status) != statusCodeSuccess {
		return fmt.Errorf("error presenting")
	}

	return nil
}

func (s *Surface) Unconfigure() {
	C.wgpuSurfaceUnconfigure(s.ref)
}

func (s *Surface) SetLabel(label string) {
	C.wgpuSurfaceSetLabel(s.ref, toCStr(label))
}
