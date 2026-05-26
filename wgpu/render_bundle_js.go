//go:build js

package wgpu

import "syscall/js"

// RenderBundle represents a pre-recorded sequence of render commands that can be executed efficiently multiple times.
// Render bundles are created from a render bundle encoder and can be executed in render passes.
type RenderBundle struct {
	ref js.Value
}

// SetLabel sets the debug label for the render bundle.
// This label appears in debuggers and validation layers.
func (r *RenderBundle) SetLabel(label string) {
	r.ref.Set("label", label)
}

// RenderBundleEncoder encodes a sequence of render commands that can be recorded into a render bundle.
// Render bundle encoders are created from a device and are used to pre-record render commands.
type RenderBundleEncoder struct {
	ref js.Value
}

// SetPipeline sets the render pipeline to be used for subsequent render commands in the bundle.
func (r *RenderBundleEncoder) SetPipeline(pipeline *RenderPipeline) {
	r.ref.Call("setPipeline", pipeline.ref)
}

// SetBindGroup sets a bind group to be used for subsequent render commands in the bundle.
// The groupIndex specifies which bind group slot to use.
// The dynamicOffsets provide values for any dynamic buffer offsets in the bind group.
func (r *RenderBundleEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsets []uint32) {
	arr := js.Global().Get("Uint32Array").New(len(dynamicOffsets))
	for i, o := range dynamicOffsets {
		arr.SetIndex(i, o)
	}
	r.ref.Call("setBindGroup", groupIndex, group.ref, arr)
}

// Draw draws non-indexed primitives using the currently set pipeline and vertex buffers in the bundle.
// The vertexCount specifies the number of vertices to draw.
func (r *RenderBundleEncoder) Draw(vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	r.ref.Call("draw", vertexCount, instanceCount, firstVertex, firstInstance)
}

// DrawIndexed draws indexed primitives using the currently set pipeline, index buffer, and vertex buffers in the bundle.
// The indexCount specifies the number of indices to draw.
func (r *RenderBundleEncoder) DrawIndexed(indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	r.ref.Call("drawIndexed", indexCount, instanceCount, firstIndex, baseVertex, firstInstance)
}

// DrawIndirect draws primitives using parameters from a buffer in the bundle.
// The indirectBuffer contains the draw parameters and indirectOffset specifies the offset into that buffer.
func (r *RenderBundleEncoder) DrawIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	r.ref.Call("drawIndirect", indirectBuffer.ref, indirectOffset)
}

// DrawIndexedIndirect draws indexed primitives using parameters from a buffer in the bundle.
// The indirectBuffer contains the draw parameters and indirectOffset specifies the offset into that buffer.
func (r *RenderBundleEncoder) DrawIndexedIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	r.ref.Call("drawIndexedIndirect", indirectBuffer.ref, indirectOffset)
}

// InsertDebugMarker inserts a debug marker into the render bundle encoder.
// The marker label is used to identify the marker in debuggers and profilers.
func (r *RenderBundleEncoder) InsertDebugMarker(markerLabel string) {
	r.ref.Call("insertDebugMarker", markerLabel)
}

// PopDebugGroup pops the most recently pushed debug group from the render bundle encoder.
func (r *RenderBundleEncoder) PopDebugGroup() {
	r.ref.Call("popDebugGroup")
}

// PushDebugGroup pushes a debug group into the render bundle encoder with the given label.
// Debug groups can be nested and are used to group commands in debuggers and profilers.
func (r *RenderBundleEncoder) PushDebugGroup(groupLabel string) {
	r.ref.Call("pushDebugGroup", groupLabel)
}

// SetVertexBuffer sets a vertex buffer at the specified slot for subsequent draw commands in the bundle.
// The offset and size specify the region of the buffer to use.
func (r *RenderBundleEncoder) SetVertexBuffer(slot uint32, buffer *Buffer, offset uint64, size uint64) {
	r.ref.Call("setVertexBuffer", slot, buffer.ref, offset, size)
}

// SetIndexBuffer sets an index buffer for subsequent indexed draw commands in the bundle.
// The format specifies the type of indices and the offset and size specify the region of the buffer to use.
func (r *RenderBundleEncoder) SetIndexBuffer(buffer *Buffer, format IndexFormat, offset uint64, size uint64) {
	r.ref.Call("setIndexBuffer", buffer.ref, format.toJS(), offset, size)
}

// Finish finishes recording and returns a render bundle.
// The descriptor can be used to set the label of the render bundle.
func (r *RenderBundleEncoder) Finish(descriptor *RenderBundleDescriptor) RenderBundle {
	var desc map[string]any
	if descriptor != nil {
		desc = map[string]any{"label": descriptor.Label}
	}
	return RenderBundle{ref: r.ref.Call("finish", desc)}
}

// SetLabel sets the debug label for the render bundle encoder.
// This label appears in debuggers and validation layers.
func (r *RenderBundleEncoder) SetLabel(label string) {
	r.ref.Set("label", label)
}
