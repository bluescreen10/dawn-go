//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"

type PipelineLayout struct {
	ref C.WGPUPipelineLayout
}

func (p *PipelineLayout) SetLabel(label string) {
	C.wgpuPipelineLayoutSetLabel(p.ref, toCStr(label))
}
