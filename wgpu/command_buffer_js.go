//go:build js

package wgpu

import "syscall/js"

// CommandBuffer represents a sequence of GPU commands that can be submitted to a queue.
// Command buffers are created from a command encoder and contain recorded commands.
type CommandBuffer struct {
	ref js.Value
}

// SetLabel sets the debug label for the command buffer.
// This label appears in debuggers and validation layers.
func (c *CommandBuffer) SetLabel(label string) {
	c.ref.Set("label", label)
}

// Release is a no-op in JS (GC handles cleanup).
func (c *CommandBuffer) Release() {}
