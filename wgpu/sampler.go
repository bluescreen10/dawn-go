//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"

type Sampler struct {
	ref C.WGPUSampler
}

func (s *Sampler) SetLabel(label string) {
	C.wgpuSamplerSetLabel(s.ref, toCStr(label))
}
