package wgpu

/*
#include "webgpu.h"
*/
import "C"

type RenderPipeline struct {
	ref C.WGPURenderPipeline
}

func (r *RenderPipeline) GetBindGroupLayout(groupIndex uint32) *BindGroupLayout {
	cGroupIndex := C.uint32_t(groupIndex)
	return &BindGroupLayout{ref: C.wgpuRenderPipelineGetBindGroupLayout(r.ref, cGroupIndex)}
}

func (r *RenderPipeline) SetLabel(label string) {
	cLabel := toCStr(label)
	C.wgpuRenderPipelineSetLabel(r.ref, cLabel)
}

func (r *RenderPipeline) Release() {
	C.wgpuRenderPipelineRelease(r.ref)
}
