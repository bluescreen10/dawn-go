//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"
import "unsafe"

// RenderBundle represents a pre-recorded sequence of render commands that can be executed efficiently multiple times.
// Render bundles are created from a render bundle encoder and can be executed in render passes.
type RenderBundle struct {
	ref C.WGPURenderBundle
}

// SetLabel sets the debug label for the render bundle.
// This label appears in debuggers and validation layers.
func (r *RenderBundle) SetLabel(label string) {
	C.wgpuRenderBundleSetLabel(r.ref, toCStr(label))
}

// RenderBundleEncoder encodes a sequence of render commands that can be recorded into a render bundle.
// Render bundle encoders are created from a device and are used to pre-record render commands.
type RenderBundleEncoder struct {
	ref C.WGPURenderBundleEncoder
}

// SetPipeline sets the render pipeline to be used for subsequent render commands in the bundle.
func (r *RenderBundleEncoder) SetPipeline(pipeline *RenderPipeline) {
	C.wgpuRenderBundleEncoderSetPipeline(r.ref, pipeline.ref)
}

// SetBindGroup sets a bind group to be used for subsequent render commands in the bundle.
// The groupIndex specifies which bind group slot to use.
// The dynamicOffsets provide values for any dynamic buffer offsets in the bind group.
func (r *RenderBundleEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsets []uint32) {
	cDynamicOffsetCount := C.size_t(len(dynamicOffsets))
	var cDynamicOffsets *C.uint32_t

	if cDynamicOffsetCount > 0 {
		cDynamicOffsets = (*C.uint32_t)(unsafe.Pointer(&dynamicOffsets[0]))
	}

	C.wgpuRenderBundleEncoderSetBindGroup(r.ref, C.uint32_t(groupIndex), group.ref, cDynamicOffsetCount, cDynamicOffsets)
}

// Draw draws non-indexed primitives using the currently set pipeline and vertex buffers in the bundle.
// The vertexCount specifies the number of vertices to draw.
func (r *RenderBundleEncoder) Draw(vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	C.wgpuRenderBundleEncoderDraw(r.ref, C.uint32_t(vertexCount), C.uint32_t(instanceCount), C.uint32_t(firstVertex), C.uint32_t(firstInstance))
}

// DrawIndexed draws indexed primitives using the currently set pipeline, index buffer, and vertex buffers in the bundle.
// The indexCount specifies the number of indices to draw.
func (r *RenderBundleEncoder) DrawIndexed(indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	C.wgpuRenderBundleEncoderDrawIndexed(r.ref, C.uint32_t(indexCount), C.uint32_t(instanceCount), C.uint32_t(firstIndex), C.int32_t(baseVertex), C.uint32_t(firstInstance))
}

// DrawIndirect draws primitives using parameters from a buffer in the bundle.
// The indirectBuffer contains the draw parameters and indirectOffset specifies the offset into that buffer.
func (r *RenderBundleEncoder) DrawIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	C.wgpuRenderBundleEncoderDrawIndirect(r.ref, indirectBuffer.ref, C.uint64_t(indirectOffset))
}

// DrawIndexedIndirect draws indexed primitives using parameters from a buffer in the bundle.
// The indirectBuffer contains the draw parameters and indirectOffset specifies the offset into that buffer.
func (r *RenderBundleEncoder) DrawIndexedIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	C.wgpuRenderBundleEncoderDrawIndexedIndirect(r.ref, indirectBuffer.ref, C.uint64_t(indirectOffset))
}

// InsertDebugMarker inserts a debug marker into the render bundle encoder.
// The marker label is used to identify the marker in debuggers and profilers.
func (r *RenderBundleEncoder) InsertDebugMarker(markerLabel string) {
	C.wgpuRenderBundleEncoderInsertDebugMarker(r.ref, toCStr(markerLabel))
}

// PopDebugGroup pops the most recently pushed debug group from the render bundle encoder.
func (r *RenderBundleEncoder) PopDebugGroup() {
	C.wgpuRenderBundleEncoderPopDebugGroup(r.ref)
}

// PushDebugGroup pushes a debug group into the render bundle encoder with the given label.
// Debug groups can be nested and are used to group commands in debuggers and profilers.
func (r *RenderBundleEncoder) PushDebugGroup(groupLabel string) {
	C.wgpuRenderBundleEncoderPushDebugGroup(r.ref, toCStr(groupLabel))
}

// SetVertexBuffer sets a vertex buffer at the specified slot for subsequent draw commands in the bundle.
// The offset and size specify the region of the buffer to use.
func (r *RenderBundleEncoder) SetVertexBuffer(slot uint32, buffer *Buffer, offset uint64, size uint64) {
	C.wgpuRenderBundleEncoderSetVertexBuffer(r.ref, C.uint32_t(slot), buffer.ref, C.uint64_t(offset), C.uint64_t(size))
}

// SetIndexBuffer sets an index buffer for subsequent indexed draw commands in the bundle.
// The format specifies the type of indices and the offset and size specify the region of the buffer to use.
func (r *RenderBundleEncoder) SetIndexBuffer(buffer *Buffer, format IndexFormat, offset uint64, size uint64) {
	C.wgpuRenderBundleEncoderSetIndexBuffer(r.ref, buffer.ref, C.WGPUIndexFormat(format), C.uint64_t(offset), C.uint64_t(size))
}

// Finish finishes recording and returns a render bundle.
// The descriptor can be used to set the label of the render bundle.
func (r *RenderBundleEncoder) Finish(descriptor *RenderBundleDescriptor) RenderBundle {
	var cDescriptor *C.WGPURenderBundleDescriptor

	if descriptor != nil {
		cDescriptor = &C.WGPURenderBundleDescriptor{
			label: toCStr(descriptor.Label),
		}
	}

	return RenderBundle{ref: C.wgpuRenderBundleEncoderFinish(r.ref, cDescriptor)}
}

// SetLabel sets the debug label for the render bundle encoder.
// This label appears in debuggers and validation layers.
func (r *RenderBundleEncoder) SetLabel(label string) {
	C.wgpuRenderBundleEncoderSetLabel(r.ref, toCStr(label))
}
