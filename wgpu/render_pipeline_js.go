//go:build js

package wgpu

import "syscall/js"

// RenderPipeline represents a render pipeline that defines how graphics are rendered.
// Render pipelines are created from a device and a render pipeline descriptor.
// They define the shader modules, vertex and fragment stages, and render state.
type RenderPipeline struct {
	ref js.Value
}

// GetBindGroupLayout returns the bind group layout at the specified group index for this render pipeline.
// This layout defines the interface for bind groups used with this pipeline.
func (r *RenderPipeline) GetBindGroupLayout(groupIndex uint32) *BindGroupLayout {
	return &BindGroupLayout{ref: r.ref.Call("getBindGroupLayout", groupIndex)}
}

// SetLabel sets the debug label for the render pipeline.
// This label appears in debuggers and validation layers.
func (r *RenderPipeline) SetLabel(label string) {
	r.ref.Set("label", label)
}

// Release is a no-op in JS (GC handles cleanup).
func (r *RenderPipeline) Release() {}
