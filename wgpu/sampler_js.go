//go:build js

package wgpu

import "syscall/js"

// Sampler represents a sampler that defines how textures are sampled in shaders.
// Samplers are created from a device and define filtering modes, addressing modes, and other sampling parameters.
type Sampler struct {
	ref js.Value
}

// SetLabel sets the debug label for the sampler.
// This label appears in debuggers and validation layers.
func (s *Sampler) SetLabel(label string) {
	s.ref.Set("label", label)
}

// Release is a no-op in JS (GC handles cleanup).
func (s *Sampler) Release() {}
