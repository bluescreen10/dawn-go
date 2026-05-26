//go:build js

package wgpu

import "syscall/js"

// BindGroupLayout represents the interface for a bind group, defining the types and access patterns of its resources.
// Bind group layouts are created from a device and are used to create bind groups.
type BindGroupLayout struct {
	ref js.Value
}

// SetLabel sets the debug label for the bind group layout.
// This label appears in debuggers and validation layers.
func (b *BindGroupLayout) SetLabel(label string) {
	b.ref.Set("label", label)
}

// Release is a no-op in JS (GC handles cleanup).
func (b *BindGroupLayout) Release() {}
