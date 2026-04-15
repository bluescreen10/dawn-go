//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"

// Texture represents a GPU texture, which is a structured collection of pixels used for rendering and data storage.
// Textures are created from a device and can be used as render targets or sampled in shaders.
type Texture struct {
	ref C.WGPUTexture
}

// AsImageCopy returns a TexelCopyTextureInfo for use in image copy operations.
// The texture is treated as having all aspects, full mip levels, and the full extent.
func (t *Texture) AsImageCopy() TexelCopyTextureInfo {
	return TexelCopyTextureInfo{
		Texture:  t,
		MipLevel: 0,
		Origin:   Origin3D{},
		Aspect:   TextureAspectAll,
	}
}

// CreateView creates a texture view from this texture.
// Texture views can have different dimensions and mip level ranges than the underlying texture.
func (t *Texture) CreateView(descriptor *TextureViewDescriptor) *TextureView {

	var cDescriptor *C.WGPUTextureViewDescriptor
	if descriptor != nil {
		cDescriptor = &C.WGPUTextureViewDescriptor{
			label:           toCStr(descriptor.Label),
			format:          C.WGPUTextureFormat(descriptor.Format),
			dimension:       C.WGPUTextureViewDimension(descriptor.Dimension),
			baseMipLevel:    C.uint32_t(descriptor.BaseMipLevel),
			mipLevelCount:   C.uint32_t(descriptor.MipLevelCount),
			baseArrayLayer:  C.uint32_t(descriptor.BaseArrayLayer),
			arrayLayerCount: C.uint32_t(descriptor.ArrayLayerCount),
			aspect:          C.WGPUTextureAspect(descriptor.Aspect),
			usage:           C.WGPUTextureUsage(descriptor.Usage),
		}

		if cDescriptor.mipLevelCount == 0 {
			cDescriptor.mipLevelCount = MipLevelCountUndefined
		}

		if cDescriptor.arrayLayerCount == 0 {
			cDescriptor.arrayLayerCount = ArrayLayerCountUndefined
		}
	}

	return &TextureView{ref: C.wgpuTextureCreateView(t.ref, cDescriptor)}
}

// SetLabel sets the debug label for the texture.
// This label appears in debuggers and validation layers.
func (t *Texture) SetLabel(label string) {
	C.wgpuTextureSetLabel(t.ref, toCStr(label))
}

// GetWidth returns the width of the texture in pixels.
func (t *Texture) GetWidth() uint32 {
	return uint32(C.wgpuTextureGetWidth(t.ref))
}

// GetHeight returns the height of the texture in pixels.
func (t *Texture) GetHeight() uint32 {
	return uint32(C.wgpuTextureGetHeight(t.ref))
}

// GetDepthOrArrayLayers returns the depth or array layer count of the texture.
// For 3D textures, this is the depth; for 2D array textures, this is the number of layers.
func (t *Texture) GetDepthOrArrayLayers() uint32 {
	return uint32(C.wgpuTextureGetDepthOrArrayLayers(t.ref))
}

// GetMipLevelCount returns the number of mip levels in the texture.
func (t *Texture) GetMipLevelCount() uint32 {
	return uint32(C.wgpuTextureGetMipLevelCount(t.ref))
}

// GetSampleCount returns the sample count of the texture.
// A value of 1 means single-sampled; values greater than 1 indicate multisampling.
func (t *Texture) GetSampleCount() uint32 {
	return uint32(C.wgpuTextureGetSampleCount(t.ref))
}

// GetDimension returns the dimension of the texture (1D, 2D, or 3D).
func (t *Texture) GetDimension() TextureDimension {
	return TextureDimension(C.wgpuTextureGetDimension(t.ref))
}

// GetFormat returns the pixel format of the texture.
func (t *Texture) GetFormat() TextureFormat {
	return TextureFormat(C.wgpuTextureGetFormat(t.ref))
}

// GetUsage returns the usage flags for the texture.
func (t *Texture) GetUsage() TextureUsage {
	return TextureUsage(C.wgpuTextureGetUsage(t.ref))
}

// GetTextureBindingViewDimension returns the view dimension for texture binding.
// This is useful for determining how the texture can be sampled in shaders.
func (t *Texture) GetTextureBindingViewDimension() TextureViewDimension {
	return TextureViewDimension(C.wgpuTextureGetTextureBindingViewDimension(t.ref))
}

// Destroy destroys the texture and frees all associated GPU resources.
func (t *Texture) Destroy() {
	C.wgpuTextureDestroy(t.ref)
}

// TextureView represents a view of a texture with specific dimensions, mip levels, and array layers.
// Texture views are created from textures and can be used as render targets or sampled in shaders.
type TextureView struct {
	ref C.WGPUTextureView
}

// SetLabel sets the debug label for the texture view.
// This label appears in debuggers and validation layers.
func (t *TextureView) SetLabel(label string) {
	C.wgpuTextureViewSetLabel(t.ref, toCStr(label))
}

// Release releases the texture view and all associated resources.
// After calling this method, the view should no longer be used.
func (t *TextureView) Release() {
	C.wgpuTextureViewRelease(t.ref)
}
