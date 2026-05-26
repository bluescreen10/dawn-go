//go:build js

package wgpu

import (
	"fmt"
	"syscall/js"
)

// Surface represents a surface that can be used to present rendered graphics to a window.
// In JS/WASM this wraps a GPUCanvasContext from a canvas element.
type Surface struct {
	ref    js.Value // GPUCanvasContext
	canvas js.Value // HTMLCanvasElement
}

// Configure configures the surface for rendering with the specified device and settings.
// The configuration defines the format, size, and present mode for the surface.
func (s *Surface) Configure(config SurfaceConfiguration) {
	desc := map[string]any{
		"device": config.Device.ref,
		"format": config.Format.toJS(),
		"usage":  int(config.Usage),
	}
	if alphaMode := config.AlphaMode.toJS(); alphaMode != "" {
		desc["alphaMode"] = alphaMode
	}
	if len(config.ViewFormats) > 0 {
		vf := make([]any, len(config.ViewFormats))
		for i, f := range config.ViewFormats {
			vf[i] = f.toJS()
		}
		desc["viewFormats"] = vf
	}
	s.ref.Call("configure", desc)
}

// GetCapabilities returns the capabilities of the surface when used with the given adapter.
// The capabilities include supported usages, formats, present modes, and alpha modes.
func (s *Surface) GetCapabilities(adapter *Adapter) (SurfaceCapabilities, error) {
	// getPreferredCanvasFormat lives on navigator.gpu, not on the canvas context, and takes no arguments.
	gpu := js.Global().Get("navigator").Get("gpu")
	var format TextureFormat
	preferredFormat := gpu.Call("getPreferredCanvasFormat")
	if !preferredFormat.IsUndefined() && !preferredFormat.IsNull() {
		format = textureFormatFromJS(preferredFormat.String())
	} else {
		format = TextureFormatBGRA8Unorm
	}

	return SurfaceCapabilities{
		Usages:       TextureUsageRenderAttachment,
		Formats:      []TextureFormat{format, TextureFormatRGBA8Unorm},
		PresentModes: []PresentMode{PresentModeFifo},
		AlphaModes:   []CompositeAlphaMode{CompositeAlphaModeAuto, CompositeAlphaModeOpaque, CompositeAlphaModePremultiplied},
	}, nil
}

// GetCurrentTexture obtains the current texture to render to from the surface.
// Panics if the texture cannot be obtained.
func (s *Surface) GetCurrentTexture() *Texture {
	tex, err := s.TryGetCurrentTexture()
	if err != nil {
		panic(err)
	}
	return tex
}

// TryGetCurrentTexture obtains the current texture to render to from the surface, or returns an error if it cannot be obtained.
func (s *Surface) TryGetCurrentTexture() (*Texture, error) {
	jsTexture := s.ref.Call("getCurrentTexture")
	if jsTexture.IsNull() || jsTexture.IsUndefined() {
		return nil, fmt.Errorf("getCurrentTexture: no texture available")
	}
	return &Texture{ref: jsTexture}, nil
}

// Present presents the current texture to the screen.
// In JS, presentation happens automatically so this is a no-op.
func (s *Surface) Present() {}

// TryPresent presents the current texture to the screen, returning an error if presentation fails.
// In JS, this is a no-op as presentation happens automatically.
func (s *Surface) TryPresent() error { return nil }

// Unconfigure unconfigures the surface, releasing any resources associated with it.
// After calling this method, the surface must be reconfigured before it can be used again.
func (s *Surface) Unconfigure() {
	s.ref.Call("unconfigure")
}

// SetLabel sets the debug label for the surface.
// This label appears in debuggers and validation layers.
func (s *Surface) SetLabel(label string) {
	if !s.canvas.IsUndefined() {
		s.canvas.Set("id", label)
	}
}

// Release releases the surface and all associated resources.
// In JS, this is a no-op as GC handles cleanup.
func (s *Surface) Release() {}
