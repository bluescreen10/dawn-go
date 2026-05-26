//go:build js

package wgpu

import "syscall/js"

// ComputePipeline represents a compute pipeline that can execute compute shaders.
// Compute pipelines are created from a device and a compute pipeline descriptor.
type ComputePipeline struct {
	ref js.Value
}

// GetBindGroupLayout returns the bind group layout at the specified group index for this compute pipeline.
// This layout defines the interface for bind groups used with this pipeline.
func (c *ComputePipeline) GetBindGroupLayout(groupIndex uint32) *BindGroupLayout {
	return &BindGroupLayout{ref: c.ref.Call("getBindGroupLayout", groupIndex)}
}

// SetLabel sets the debug label for the compute pipeline.
// This label appears in debuggers and validation layers.
func (c *ComputePipeline) SetLabel(label string) {
	c.ref.Set("label", label)
}

// Release is a no-op in JS (GC handles cleanup).
func (c *ComputePipeline) Release() {}
