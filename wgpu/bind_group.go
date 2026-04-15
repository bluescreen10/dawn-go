//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"

// BindGroup represents a group of resources that are bound together and used in GPU commands.
// Bind groups are created from a device using a bind group layout and a set of resources.
type BindGroup struct {
	ref C.WGPUBindGroup
}

// SetLabel sets the debug label for the bind group.
// This label appears in debuggers and validation layers.
func (b *BindGroup) SetLabel(label string) {
	cLabel := toCStr(label)
	C.wgpuBindGroupSetLabel(b.ref, cLabel)
}

// Release releases the bind group and all associated resources.
// After calling this method, the group should no longer be used.
func (b *BindGroup) Release() {
	C.wgpuBindGroupRelease(b.ref)
}
