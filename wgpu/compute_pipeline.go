//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"

type ComputePipeline struct {
	ref C.WGPUComputePipeline
}

func (c *ComputePipeline) GetBindGroupLayout(groupIndex uint32) BindGroupLayout {
	return BindGroupLayout{ref: C.wgpuComputePipelineGetBindGroupLayout(c.ref, C.uint32_t(groupIndex))}
}

func (c *ComputePipeline) SetLabel(label string) {
	C.wgpuComputePipelineSetLabel(c.ref, toCStr(label))
}
