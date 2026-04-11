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
	C.wgpuBindGroupLayoutSetLabel(b.ref, toCStr(label))
}

func (b *BindGroupLayout) Release() {
	C.wgpuBindGroupLayoutRelease(b.ref)
}
