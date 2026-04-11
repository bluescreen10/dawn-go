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
	C.wgpuCommandBufferSetLabel(c.ref, toCStr(label))
}

func (c *CommandBuffer) Release() {
	C.wgpuCommandBufferRelease(c.ref)
}
