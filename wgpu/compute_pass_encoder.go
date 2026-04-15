//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"
import "unsafe"

// ComputePassEncoder encodes compute commands that will be dispatched to the GPU.
// Compute pass encoders are created from a command encoder and are used to issue compute work.
type ComputePassEncoder struct {
	ref C.WGPUComputePassEncoder
}

// InsertDebugMarker inserts a debug marker into the compute pass.
// The marker label is used to identify the marker in debuggers and profilers.
func (c *ComputePassEncoder) InsertDebugMarker(markerLabel string) {
	C.wgpuComputePassEncoderInsertDebugMarker(c.ref, toCStr(markerLabel))
}

// PopDebugGroup pops the most recently pushed debug group from the compute pass.
func (c *ComputePassEncoder) PopDebugGroup() {
	C.wgpuComputePassEncoderPopDebugGroup(c.ref)
}

// PushDebugGroup pushes a debug group into the compute pass with the given label.
// Debug groups can be nested and are used to group commands in debuggers and profilers.
func (c *ComputePassEncoder) PushDebugGroup(groupLabel string) {
	C.wgpuComputePassEncoderPushDebugGroup(c.ref, toCStr(groupLabel))
}

// SetPipeline sets the compute pipeline to be used for subsequent compute commands.
func (c *ComputePassEncoder) SetPipeline(pipeline *ComputePipeline) {
	C.wgpuComputePassEncoderSetPipeline(c.ref, pipeline.ref)
}

// SetBindGroup sets a bind group to be used for subsequent compute commands.
// The groupIndex specifies which bind group slot to use.
// The dynamicOffsets provide values for any dynamic buffer offsets in the bind group.
func (c *ComputePassEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsets []uint32) {
	cDynamicOffsetCount := C.size_t(len(dynamicOffsets))
	var cDynamicOffsets *C.uint32_t

	if cDynamicOffsetCount > 0 {
		cDynamicOffsets = (*C.uint32_t)(unsafe.Pointer(&dynamicOffsets[0]))
	}
	C.wgpuComputePassEncoderSetBindGroup(c.ref, C.uint32_t(groupIndex), group.ref, cDynamicOffsetCount, cDynamicOffsets)
}

// DispatchWorkgroups dispatches compute workgroups with the specified dimensions.
// Each workgroup runs the compute shader with the given number of threads.
func (c *ComputePassEncoder) DispatchWorkgroups(workgroupCountX uint32, workgroupCountY uint32, workgroupCountZ uint32) {
	C.wgpuComputePassEncoderDispatchWorkgroups(c.ref, C.uint32_t(workgroupCountX), C.uint32_t(workgroupCountY), C.uint32_t(workgroupCountZ))
}

// DispatchWorkgroupsIndirect dispatches compute workgroups indirectly using parameters from a buffer.
// The indirectBuffer contains the workgroup count parameters and indirectOffset specifies the offset into that buffer.
func (c *ComputePassEncoder) DispatchWorkgroupsIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	C.wgpuComputePassEncoderDispatchWorkgroupsIndirect(c.ref, indirectBuffer.ref, C.uint64_t(indirectOffset))
}

// End ends the compute pass.
// After calling this method, no more commands can be recorded in this pass.
func (c *ComputePassEncoder) End() {
	C.wgpuComputePassEncoderEnd(c.ref)
}

// SetLabel sets the debug label for the compute pass encoder.
// This label appears in debuggers and validation layers.
func (c *ComputePassEncoder) SetLabel(label string) {
	C.wgpuComputePassEncoderSetLabel(c.ref, toCStr(label))
}

// Release releases the compute pass encoder and all associated resources.
// After calling this method, the encoder should no longer be used.
func (c *ComputePassEncoder) Release() {
	C.wgpuComputePassEncoderRelease(c.ref)
}
