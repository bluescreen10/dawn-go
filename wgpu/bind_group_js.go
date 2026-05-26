//go:build js

package wgpu

import "syscall/js"

// BindGroup represents a group of resources that are bound together and used in GPU commands.
// Bind groups are created from a device using a bind group layout and a set of resources.
type BindGroup struct {
	ref js.Value
}

// SetLabel sets the debug label for the bind group.
// This label appears in debuggers and validation layers.
func (b *BindGroup) SetLabel(label string) {
	b.ref.Set("label", label)
}

// Release is a no-op in JS (GC handles cleanup).
func (b *BindGroup) Release() {}
