//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"

// ComputePipeline represents a compute pipeline that can execute compute shaders.
// Compute pipelines are created from a device and a compute pipeline descriptor.
type ComputePipeline struct {
	ref C.WGPUComputePipeline
}

// GetBindGroupLayout returns the bind group layout at the specified group index for this compute pipeline.
// This layout defines the interface for bind groups used with this pipeline.
func (c *ComputePipeline) GetBindGroupLayout(groupIndex uint32) *BindGroupLayout {
	return &BindGroupLayout{ref: C.wgpuComputePipelineGetBindGroupLayout(c.ref, C.uint32_t(groupIndex))}
}

// SetLabel sets the debug label for the compute pipeline.
// This label appears in debuggers and validation layers.
func (c *ComputePipeline) SetLabel(label string) {
	C.wgpuComputePipelineSetLabel(c.ref, toCStr(label))
}

// Release releases the compute pipeline and all associated resources.
// After calling this method, the pipeline should no longer be used.
func (c *ComputePipeline) Release() {
	C.wgpuComputePipelineRelease(c.ref)
}
