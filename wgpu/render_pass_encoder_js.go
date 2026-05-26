//go:build js

package wgpu

import "syscall/js"

// RenderPassEncoder encodes render commands that will be drawn to a set of render targets.
// Render pass encoders are created from a command encoder and are used to issue rendering commands.
type RenderPassEncoder struct {
	ref js.Value
}

// SetPipeline sets the render pipeline to be used for subsequent render commands.
func (r *RenderPassEncoder) SetPipeline(pipeline *RenderPipeline) {
	r.ref.Call("setPipeline", pipeline.ref)
}

// SetBindGroup sets a bind group to be used for subsequent render commands.
// The groupIndex specifies which bind group slot to use.
// The dynamicOffsets provide values for any dynamic buffer offsets in the bind group.
func (r *RenderPassEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsets []uint32) {
	arr := js.Global().Get("Uint32Array").New(len(dynamicOffsets))
	for i, o := range dynamicOffsets {
		arr.SetIndex(i, o)
	}
	r.ref.Call("setBindGroup", groupIndex, group.ref, arr)
}

// Draw draws non-indexed primitives using the currently set pipeline and vertex buffers.
// The vertexCount specifies the number of vertices to draw.
func (r *RenderPassEncoder) Draw(vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	r.ref.Call("draw", vertexCount, instanceCount, firstVertex, firstInstance)
}

// DrawIndexed draws indexed primitives using the currently set pipeline, index buffer, and vertex buffers.
// The indexCount specifies the number of indices to draw.
func (r *RenderPassEncoder) DrawIndexed(indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	r.ref.Call("drawIndexed", indexCount, instanceCount, firstIndex, baseVertex, firstInstance)
}

// DrawIndirect draws primitives using parameters from a buffer.
// The indirectBuffer contains the draw parameters and indirectOffset specifies the offset into that buffer.
func (r *RenderPassEncoder) DrawIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	r.ref.Call("drawIndirect", indirectBuffer.ref, indirectOffset)
}

// DrawIndexedIndirect draws indexed primitives using parameters from a buffer.
// The indirectBuffer contains the draw parameters and indirectOffset specifies the offset into that buffer.
func (r *RenderPassEncoder) DrawIndexedIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	r.ref.Call("drawIndexedIndirect", indirectBuffer.ref, indirectOffset)
}

// ExecuteBundles executes a sequence of render bundles in the render pass.
// Render bundles are pre-recorded render commands that can be executed efficiently.
func (r *RenderPassEncoder) ExecuteBundles(bundles ...*RenderBundle) {
	bds := make([]any, len(bundles))
	for i, b := range bundles {
		bds[i] = b.ref
	}
	r.ref.Call("executeBundles", bds)
}

// InsertDebugMarker inserts a debug marker into the render pass.
// The marker label is used to identify the marker in debuggers and profilers.
func (r *RenderPassEncoder) InsertDebugMarker(markerLabel string) {
	r.ref.Call("insertDebugMarker", markerLabel)
}

// PopDebugGroup pops the most recently pushed debug group from the render pass.
func (r *RenderPassEncoder) PopDebugGroup() {
	r.ref.Call("popDebugGroup")
}

// PushDebugGroup pushes a debug group into the render pass with the given label.
// Debug groups can be nested and are used to group commands in debuggers and profilers.
func (r *RenderPassEncoder) PushDebugGroup(groupLabel string) {
	r.ref.Call("pushDebugGroup", groupLabel)
}

// SetStencilReference sets the stencil reference value for subsequent stencil operations.
func (r *RenderPassEncoder) SetStencilReference(reference uint32) {
	r.ref.Call("setStencilReference", reference)
}

// SetBlendConstant sets the blend constant color used for blend operations that use a constant color.
func (r *RenderPassEncoder) SetBlendConstant(color Color) {
	r.ref.Call("setBlendConstant", map[string]any{
		"r": color.R,
		"g": color.G,
		"b": color.B,
		"a": color.A,
	})
}

// SetViewport sets the viewport for subsequent render commands.
// The viewport defines the region of the render target that output is drawn to.
func (r *RenderPassEncoder) SetViewport(x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32) {
	r.ref.Call("setViewport", x, y, width, height, minDepth, maxDepth)
}

// SetScissorRect sets the scissor rectangle for subsequent render commands.
// Pixels outside this rectangle will be discarded.
func (r *RenderPassEncoder) SetScissorRect(x uint32, y uint32, width uint32, height uint32) {
	r.ref.Call("setScissorRect", x, y, width, height)
}

// SetVertexBuffer sets a vertex buffer at the specified slot for subsequent draw commands.
// The offset and size specify the region of the buffer to use.
func (r *RenderPassEncoder) SetVertexBuffer(slot uint32, buffer *Buffer, offset uint64, size uint64) {
	r.ref.Call("setVertexBuffer", slot, buffer.ref, offset, size)
}

// SetIndexBuffer sets an index buffer for subsequent indexed draw commands.
// The format specifies the type of indices and the offset and size specify the region of the buffer to use.
func (r *RenderPassEncoder) SetIndexBuffer(buffer *Buffer, format IndexFormat, offset uint64, size uint64) {
	r.ref.Call("setIndexBuffer", buffer.ref, format.toJS(), offset, size)
}

// BeginOcclusionQuery begins an occlusion query at the specified query index.
// Occlusion queries can be used to determine how many samples pass the depth test.
func (r *RenderPassEncoder) BeginOcclusionQuery(queryIndex uint32) {
	r.ref.Call("beginOcclusionQuery", queryIndex)
}

// EndOcclusionQuery ends the current occlusion query.
func (r *RenderPassEncoder) EndOcclusionQuery() {
	r.ref.Call("endOcclusionQuery")
}

// End ends the render pass.
// After calling this method, no more commands can be recorded in this pass.
func (r *RenderPassEncoder) End() {
	r.ref.Call("end")
}

// SetLabel sets the debug label for the render pass encoder.
// This label appears in debuggers and validation layers.
func (r *RenderPassEncoder) SetLabel(label string) {
	r.ref.Set("label", label)
}

// Release is a no-op in JS (GC handles cleanup).
func (r *RenderPassEncoder) Release() {}
