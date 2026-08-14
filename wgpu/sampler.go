//go:build !js

package wgpu

/*
#include "webgpu/webgpu.h"
*/
import "C"
import "runtime"

// Sampler represents a sampler that defines how textures are sampled in shaders.
// Samplers are created from a device and define filtering modes, addressing modes, and other sampling parameters.
type Sampler struct {
	ref C.WGPUSampler
}

// SetLabel sets the debug label for the sampler.
// This label appears in debuggers and validation layers.
func (s *Sampler) SetLabel(label string) {
	var pinner runtime.Pinner
	defer pinner.Unpin()

	cLabel := toCStr(label)
	pinner.Pin(cLabel.data)

	C.wgpuSamplerSetLabel(s.ref, cLabel)
}

// Release releases the sampler and all associated resources.
// After calling this method, the sampler should no longer be used.
func (s *Sampler) Release() {
	C.wgpuSamplerRelease(s.ref)
}
