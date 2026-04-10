//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"

type CommandBuffer struct {
	ref C.WGPUCommandBuffer
}

func (c *CommandBuffer) SetLabel(label string) {
	cLabel := toCStr(label)
	C.wgpuCommandBufferSetLabel(c.ref, cLabel)
}

func (c *CommandBuffer) Release() {
	C.wgpuCommandBufferRelease(c.ref)
}
