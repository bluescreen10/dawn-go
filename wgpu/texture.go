//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"

type Texture struct {
	ref C.WGPUTexture
}

func (t *Texture) AsImageCopy() TexelCopyTextureInfo {
	return TexelCopyTextureInfo{
		Texture:  t,
		MipLevel: 0,
		Origin:   Origin3D{},
		Aspect:   TextureAspectAll,
	}
}

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

func (t *Texture) SetLabel(label string) {
	C.wgpuTextureSetLabel(t.ref, toCStr(label))
}

func (t *Texture) GetWidth() uint32 {
	return uint32(C.wgpuTextureGetWidth(t.ref))
}

func (t *Texture) GetHeight() uint32 {
	return uint32(C.wgpuTextureGetHeight(t.ref))
}

func (t *Texture) GetDepthOrArrayLayers() uint32 {
	return uint32(C.wgpuTextureGetDepthOrArrayLayers(t.ref))
}

func (t *Texture) GetMipLevelCount() uint32 {
	return uint32(C.wgpuTextureGetMipLevelCount(t.ref))
}

func (t *Texture) GetSampleCount() uint32 {
	return uint32(C.wgpuTextureGetSampleCount(t.ref))
}

func (t *Texture) GetDimension() TextureDimension {
	return TextureDimension(C.wgpuTextureGetDimension(t.ref))
}

func (t *Texture) GetFormat() TextureFormat {
	return TextureFormat(C.wgpuTextureGetFormat(t.ref))
}

func (t *Texture) GetUsage() TextureUsage {
	return TextureUsage(C.wgpuTextureGetUsage(t.ref))
}

func (t *Texture) GetTextureBindingViewDimension() TextureViewDimension {
	return TextureViewDimension(C.wgpuTextureGetTextureBindingViewDimension(t.ref))
}

func (t *Texture) Destroy() {
	C.wgpuTextureDestroy(t.ref)
}

type TextureView struct {
	ref C.WGPUTextureView
}

func (t *TextureView) SetLabel(label string) {
	C.wgpuTextureViewSetLabel(t.ref, toCStr(label))
}

func (t *TextureView) Release() {
	C.wgpuTextureViewRelease(t.ref)
}
