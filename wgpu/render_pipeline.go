package wgpu

/*
#include "webgpu.h"
*/
import "C"

// RenderPipeline represents a render pipeline that defines how graphics are rendered.
// Render pipelines are created from a device and a render pipeline descriptor.
// They define the shader modules, vertex and fragment stages, and render state.
type RenderPipeline struct {
	ref C.WGPURenderPipeline
}

// GetBindGroupLayout returns the bind group layout at the specified group index for this render pipeline.
// This layout defines the interface for bind groups used with this pipeline.
func (r *RenderPipeline) GetBindGroupLayout(groupIndex uint32) *BindGroupLayout {
	cGroupIndex := C.uint32_t(groupIndex)
	return &BindGroupLayout{ref: C.wgpuRenderPipelineGetBindGroupLayout(r.ref, cGroupIndex)}
}

// SetLabel sets the debug label for the render pipeline.
// This label appears in debuggers and validation layers.
func (r *RenderPipeline) SetLabel(label string) {
	cLabel := toCStr(label)
	C.wgpuRenderPipelineSetLabel(r.ref, cLabel)
}

// Release releases the render pipeline and all associated resources.
// After calling this method, the pipeline should no longer be used.
func (r *RenderPipeline) Release() {
	C.wgpuRenderPipelineRelease(r.ref)
}
