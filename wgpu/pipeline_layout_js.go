//go:build js

package wgpu

import "syscall/js"

// PipelineLayout represents a pipeline layout that defines the bind group layouts used by a pipeline.
// Pipeline layouts are created from a device and are used when creating pipelines.
type PipelineLayout struct {
	ref js.Value
}

// SetLabel sets the debug label for the pipeline layout.
// This label appears in debuggers and validation layers.
func (p *PipelineLayout) SetLabel(label string) {
	p.ref.Set("label", label)
}

// Release is a no-op in JS (GC handles cleanup).
func (p *PipelineLayout) Release() {}
