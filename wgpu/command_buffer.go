//go:build !js

package wgpu

/*
#include "webgpu/webgpu.h"
*/
import "C"
import "runtime"

// CommandBuffer represents a sequence of GPU commands that can be submitted to a queue.
// Command buffers are created from a command encoder and contain recorded commands.
type CommandBuffer struct {
	ref C.WGPUCommandBuffer
}

// SetLabel sets the debug label for the command buffer.
// This label appears in debuggers and validation layers.
func (c *CommandBuffer) SetLabel(label string) {
	var pinner runtime.Pinner
	defer pinner.Unpin()

	cLabel := toCStr(label)
	pinner.Pin(cLabel.data)

	C.wgpuCommandBufferSetLabel(c.ref, cLabel)
}

// Release releases the command buffer and all associated resources.
// After calling this method, the buffer should no longer be used.
func (c *CommandBuffer) Release() {
	C.wgpuCommandBufferRelease(c.ref)
}
