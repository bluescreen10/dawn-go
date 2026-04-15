//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

// Surface represents a surface that can be used to present rendered graphics to a window.
// Surfaces are created from an instance and a platform-specific surface descriptor.
type Surface struct {
	ref C.WGPUSurface
}

// Configure configures the surface for rendering with the specified device and settings.
// The configuration defines the format, size, and present mode for the surface.
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

// GetCapabilities returns the capabilities of the surface when used with the given adapter.
// The capabilities include supported usages, formats, present modes, and alpha modes.
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

// GetCurrentTexture obtains the current texture to render to from the surface.
// Panics if the texture cannot be obtained.
func (s *Surface) GetCurrentTexture() *Texture {
	surfaceTexture, err := s.TryGetCurrentTexture()
	if err != nil {
		panic(err)
	}
	return surfaceTexture
}

// TryGetCurrentTexture obtains the current texture to render to from the surface, or returns an error if it cannot be obtained.
func (s *Surface) TryGetCurrentTexture() (*Texture, error) {

	var cSurfaceTexture C.WGPUSurfaceTexture
	C.wgpuSurfaceGetCurrentTexture(s.ref, &cSurfaceTexture)

	status := surfaceGetCurrentTextureStatus(cSurfaceTexture.status)
	if status != surfaceGetCurrentTextureStatusSuccessOptimal && status != surfaceGetCurrentTextureStatusSuccessSuboptimal {
		return nil, fmt.Errorf("error getting current texture: %s", status)
	}

	return &Texture{ref: cSurfaceTexture.texture}, nil
}

// Release releases the surface and all associated resources.
// After calling this method, the surface should no longer be used.
func (s *Surface) Release() {
	C.wgpuSurfaceRelease(s.ref)
}

// Present presents the current texture to the screen.
// Panics if presentation fails.
func (s *Surface) Present() {
	err := s.TryPresent()
	if err != nil {
		panic(err)
	}
}

// TryPresent presents the current texture to the screen, returning an error if presentation fails.
func (s *Surface) TryPresent() error {
	status := C.wgpuSurfacePresent(s.ref)

	if statusCode(status) != statusCodeSuccess {
		return fmt.Errorf("error presenting")
	}

	return nil
}

// Unconfigure unconfigures the surface, releasing any resources associated with it.
// After calling this method, the surface must be reconfigured before it can be used again.
func (s *Surface) Unconfigure() {
	C.wgpuSurfaceUnconfigure(s.ref)
}

// SetLabel sets the debug label for the surface.
// This label appears in debuggers and validation layers.
func (s *Surface) SetLabel(label string) {
	C.wgpuSurfaceSetLabel(s.ref, toCStr(label))
}
