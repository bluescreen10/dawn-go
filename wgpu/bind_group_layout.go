//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"

type BindGroupLayout struct {
	ref C.WGPUBindGroupLayout
}

func (b *BindGroupLayout) SetLabel(label string) {
	cLabel := toCStr(label)
	C.wgpuBindGroupLayoutSetLabel(b.ref, cLabel)
}

func (b *BindGroupLayout) Release() {
	C.wgpuBindGroupLayoutRelease(b.ref)
}
