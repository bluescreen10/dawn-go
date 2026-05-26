//go:build js

package wgpu

import "syscall/js"

// Texture represents a GPU texture, which is a structured collection of pixels used for rendering and data storage.
// Textures are created from a device and can be used as render targets or sampled in shaders.
type Texture struct {
	ref js.Value
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
	var desc map[string]any
	if descriptor != nil {
		desc = map[string]any{
			"label":           descriptor.Label,
			"format":          descriptor.Format.toJS(),
			"dimension":       descriptor.Dimension.toJS(),
			"aspect":          descriptor.Aspect.toJS(),
			"baseMipLevel":    descriptor.BaseMipLevel,
			"mipLevelCount":   descriptor.MipLevelCount,
			"baseArrayLayer":  descriptor.BaseArrayLayer,
			"arrayLayerCount": descriptor.ArrayLayerCount,
			"usage":           int(descriptor.Usage),
		}
		if descriptor.MipLevelCount == 0 {
			delete(desc, "mipLevelCount")
		}
		if descriptor.ArrayLayerCount == 0 {
			delete(desc, "arrayLayerCount")
		}
	}
	return &TextureView{ref: t.ref.Call("createView", desc)}
}

// SetLabel sets the debug label for the texture.
// This label appears in debuggers and validation layers.
func (t *Texture) SetLabel(label string) {
	t.ref.Set("label", label)
}

// GetWidth returns the width of the texture in pixels.
func (t *Texture) GetWidth() uint32 {
	return uint32(t.ref.Get("width").Int())
}

// GetHeight returns the height of the texture in pixels.
func (t *Texture) GetHeight() uint32 {
	return uint32(t.ref.Get("height").Int())
}

// GetDepthOrArrayLayers returns the depth or array layer count of the texture.
// For 3D textures, this is the depth; for 2D array textures, this is the number of layers.
func (t *Texture) GetDepthOrArrayLayers() uint32 {
	return uint32(t.ref.Get("depthOrArrayLayers").Int())
}

// GetMipLevelCount returns the number of mip levels in the texture.
func (t *Texture) GetMipLevelCount() uint32 {
	return uint32(t.ref.Get("mipLevelCount").Int())
}

// GetSampleCount returns the sample count of the texture.
// A value of 1 means single-sampled; values greater than 1 indicate multisampling.
func (t *Texture) GetSampleCount() uint32 {
	return uint32(t.ref.Get("sampleCount").Int())
}

// GetDimension returns the dimension of the texture (1D, 2D, or 3D).
func (t *Texture) GetDimension() TextureDimension {
	switch t.ref.Get("dimension").String() {
	case "1d":
		return TextureDimension1D
	case "3d":
		return TextureDimension3D
	default:
		return TextureDimension2D
	}
}

// GetFormat returns the pixel format of the texture.
func (t *Texture) GetFormat() TextureFormat {
	return textureFormatFromJS(t.ref.Get("format").String())
}

// GetUsage returns the usage flags for the texture.
func (t *Texture) GetUsage() TextureUsage {
	return TextureUsage(t.ref.Get("usage").Int())
}

// GetTextureBindingViewDimension returns the view dimension for texture binding.
// This is useful for determining how the texture can be sampled in shaders.
func (t *Texture) GetTextureBindingViewDimension() TextureViewDimension {
	return TextureViewDimension2D
}

// Destroy destroys the texture and frees all associated GPU resources.
func (t *Texture) Destroy() {
	t.ref.Call("destroy")
}

// TextureView represents a view of a texture with specific dimensions, mip levels, and array layers.
// Texture views are created from textures and can be used as render targets or sampled in shaders.
type TextureView struct {
	ref js.Value
}

// SetLabel sets the debug label for the texture view.
// This label appears in debuggers and validation layers.
func (t *TextureView) SetLabel(label string) {
	t.ref.Set("label", label)
}

// Release releases the texture view and all associated resources.
// In JS, this is a no-op as GC handles cleanup.
func (t *TextureView) Release() {}
