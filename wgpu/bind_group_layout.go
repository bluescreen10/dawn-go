//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"
import "runtime"

// BindGroupLayout represents the interface for a bind group, defining the types and access patterns of its resources.
// Bind group layouts are created from a device and are used to create bind groups.
type BindGroupLayout struct {
	ref C.WGPUBindGroupLayout
}

// SetLabel sets the debug label for the bind group layout.
// This label appears in debuggers and validation layers.
func (b *BindGroupLayout) SetLabel(label string) {
	var pinner runtime.Pinner
	defer pinner.Unpin()

	cLabel := toCStr(label)
	pinner.Pin(cLabel.data)

	C.wgpuBindGroupLayoutSetLabel(b.ref, cLabel)
}

// Release releases the bind group layout and all associated resources.
// After calling this method, the layout should no longer be used.
func (b *BindGroupLayout) Release() {
	C.wgpuBindGroupLayoutRelease(b.ref)
}
