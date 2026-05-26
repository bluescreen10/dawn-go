//go:build js

package wgpu

import "syscall/js"

// ComputePassEncoder encodes compute commands that will be dispatched to the GPU.
// Compute pass encoders are created from a command encoder and are used to issue compute work.
type ComputePassEncoder struct {
	ref js.Value
}

// InsertDebugMarker inserts a debug marker into the compute pass.
// The marker label is used to identify the marker in debuggers and profilers.
func (c *ComputePassEncoder) InsertDebugMarker(markerLabel string) {
	c.ref.Call("insertDebugMarker", markerLabel)
}

// PopDebugGroup pops the most recently pushed debug group from the compute pass.
func (c *ComputePassEncoder) PopDebugGroup() {
	c.ref.Call("popDebugGroup")
}

// PushDebugGroup pushes a debug group into the compute pass with the given label.
// Debug groups can be nested and are used to group commands in debuggers and profilers.
func (c *ComputePassEncoder) PushDebugGroup(groupLabel string) {
	c.ref.Call("pushDebugGroup", groupLabel)
}

// SetPipeline sets the compute pipeline to be used for subsequent compute commands.
func (c *ComputePassEncoder) SetPipeline(pipeline *ComputePipeline) {
	c.ref.Call("setPipeline", pipeline.ref)
}

// SetBindGroup sets a bind group to be used for subsequent compute commands.
// The groupIndex specifies which bind group slot to use.
// The dynamicOffsets provide values for any dynamic buffer offsets in the bind group.
func (c *ComputePassEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsets []uint32) {
	arr := js.Global().Get("Uint32Array").New(len(dynamicOffsets))
	for i, o := range dynamicOffsets {
		arr.SetIndex(i, o)
	}
	c.ref.Call("setBindGroup", groupIndex, group.ref, arr)
}

// DispatchWorkgroups dispatches compute workgroups with the specified dimensions.
// Each workgroup runs the compute shader with the given number of threads.
func (c *ComputePassEncoder) DispatchWorkgroups(workgroupCountX uint32, workgroupCountY uint32, workgroupCountZ uint32) {
	c.ref.Call("dispatchWorkgroups", workgroupCountX, workgroupCountY, workgroupCountZ)
}

// DispatchWorkgroupsIndirect dispatches compute workgroups indirectly using parameters from a buffer.
// The indirectBuffer contains the workgroup count parameters and indirectOffset specifies the offset into that buffer.
func (c *ComputePassEncoder) DispatchWorkgroupsIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	c.ref.Call("dispatchWorkgroupsIndirect", indirectBuffer.ref, indirectOffset)
}

// End ends the compute pass.
// After calling this method, no more commands can be recorded in this pass.
func (c *ComputePassEncoder) End() {
	c.ref.Call("end")
}

// SetLabel sets the debug label for the compute pass encoder.
// This label appears in debuggers and validation layers.
func (c *ComputePassEncoder) SetLabel(label string) {
	c.ref.Set("label", label)
}

// Release is a no-op in JS (GC handles cleanup).
func (c *ComputePassEncoder) Release() {}
