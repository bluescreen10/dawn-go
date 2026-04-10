//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"

type BindGroup struct {
	ref C.WGPUBindGroup
}

func (b *BindGroup) SetLabel(label string) {
	cLabel := toCStr(label)
	C.wgpuBindGroupSetLabel(b.ref, cLabel)
}

func (b *BindGroup) Release() {
	C.wgpuBindGroupRelease(b.ref)
}
