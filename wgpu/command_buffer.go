//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"

// CommandBuffer represents a sequence of GPU commands that can be submitted to a queue.
// Command buffers are created from a command encoder and contain recorded commands.
type CommandBuffer struct {
	ref C.WGPUCommandBuffer
}

// SetLabel sets the debug label for the command buffer.
// This label appears in debuggers and validation layers.
func (c *CommandBuffer) SetLabel(label string) {
	C.wgpuCommandBufferSetLabel(c.ref, toCStr(label))
}

// Release releases the command buffer and all associated resources.
// After calling this method, the buffer should no longer be used.
func (c *CommandBuffer) Release() {
	C.wgpuCommandBufferRelease(c.ref)
}
