//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"

// PipelineLayout represents a pipeline layout that defines the bind group layouts used by a pipeline.
// Pipeline layouts are created from a device and are used when creating pipelines.
type PipelineLayout struct {
	ref C.WGPUPipelineLayout
}

// SetLabel sets the debug label for the pipeline layout.
// This label appears in debuggers and validation layers.
func (p *PipelineLayout) SetLabel(label string) {
	C.wgpuPipelineLayoutSetLabel(p.ref, toCStr(label))
}

// Release releases the device and all associated resources.
// After calling this method, the device should no longer be used.
func (p *PipelineLayout) Release() {
	C.wgpuPipelineLayoutRelease(p.ref)
}
