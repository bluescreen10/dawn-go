//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// RenderPassEncoder encodes render commands that will be drawn to a set of render targets.
// Render pass encoders are created from a command encoder and are used to issue rendering commands.
type RenderPassEncoder struct {
	ref C.WGPURenderPassEncoder
}

// SetPipeline sets the render pipeline to be used for subsequent render commands.
func (r *RenderPassEncoder) SetPipeline(pipeline *RenderPipeline) {
	C.wgpuRenderPassEncoderSetPipeline(r.ref, pipeline.ref)
}

// SetBindGroup sets a bind group to be used for subsequent render commands.
// The groupIndex specifies which bind group slot to use.
// The dynamicOffsets provide values for any dynamic buffer offsets in the bind group.
func (r *RenderPassEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsets []uint32) {
	cDynamicOffsetsCount := C.size_t(len(dynamicOffsets))
	var cDynamicOffsets *C.uint32_t

	if cDynamicOffsetsCount > 0 {
		cDynamicOffsets = (*C.uint32_t)(&dynamicOffsets[0])
	}

	C.wgpuRenderPassEncoderSetBindGroup(r.ref, C.uint32_t(groupIndex), group.ref, cDynamicOffsetsCount, cDynamicOffsets)
}

// Draw draws non-indexed primitives using the currently set pipeline and vertex buffers.
// The vertexCount specifies the number of vertices to draw.
func (r *RenderPassEncoder) Draw(vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	C.wgpuRenderPassEncoderDraw(r.ref, C.uint32_t(vertexCount), C.uint32_t(instanceCount), C.uint32_t(firstVertex), C.uint32_t(firstInstance))
}

// DrawIndexed draws indexed primitives using the currently set pipeline, index buffer, and vertex buffers.
// The indexCount specifies the number of indices to draw.
func (r *RenderPassEncoder) DrawIndexed(indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	C.wgpuRenderPassEncoderDrawIndexed(r.ref, C.uint32_t(indexCount), C.uint32_t(instanceCount), C.uint32_t(firstIndex), C.int32_t(baseVertex), C.uint32_t(firstInstance))
}

// DrawIndirect draws primitives using parameters from a buffer.
// The indirectBuffer contains the draw parameters and indirectOffset specifies the offset into that buffer.
func (r *RenderPassEncoder) DrawIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	C.wgpuRenderPassEncoderDrawIndirect(r.ref, indirectBuffer.ref, C.uint64_t(indirectOffset))
}

// DrawIndexedIndirect draws indexed primitives using parameters from a buffer.
// The indirectBuffer contains the draw parameters and indirectOffset specifies the offset into that buffer.
func (r *RenderPassEncoder) DrawIndexedIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	C.wgpuRenderPassEncoderDrawIndexedIndirect(r.ref, indirectBuffer.ref, C.uint64_t(indirectOffset))
}

// ExecuteBundles executes a sequence of render bundles in the render pass.
// Render bundles are pre-recorded render commands that can be executed efficiently.
func (r *RenderPassEncoder) ExecuteBundles(bundles ...*RenderBundle) {
	var pinner runtime.Pinner
	defer pinner.Unpin()

	cBundleCount := C.size_t(len(bundles))
	var cBundles *C.WGPURenderBundle

	if cBundleCount > 0 {
		bds := make([]C.WGPURenderBundle, cBundleCount)
		pinner.Pin(&bundles[0])

		cBundles = (*C.WGPURenderBundle)(unsafe.Pointer(&bds[0]))

		for i, b := range bundles {
			bds[i] = C.WGPURenderBundle(b.ref)
		}
	}

	C.wgpuRenderPassEncoderExecuteBundles(r.ref, cBundleCount, cBundles)
}

// InsertDebugMarker inserts a debug marker into the render pass.
// The marker label is used to identify the marker in debuggers and profilers.
func (r *RenderPassEncoder) InsertDebugMarker(markerLabel string) {
	C.wgpuRenderPassEncoderInsertDebugMarker(r.ref, toCStr(markerLabel))
}

// PopDebugGroup pops the most recently pushed debug group from the render pass.
func (r *RenderPassEncoder) PopDebugGroup() {
	C.wgpuRenderPassEncoderPopDebugGroup(r.ref)
}

// PushDebugGroup pushes a debug group into the render pass with the given label.
// Debug groups can be nested and are used to group commands in debuggers and profilers.
func (r *RenderPassEncoder) PushDebugGroup(groupLabel string) {
	C.wgpuRenderPassEncoderPushDebugGroup(r.ref, toCStr(groupLabel))
}

// SetStencilReference sets the stencil reference value for subsequent stencil operations.
func (r *RenderPassEncoder) SetStencilReference(reference uint32) {
	C.wgpuRenderPassEncoderSetStencilReference(r.ref, C.uint32_t(reference))
}

// SetBlendConstant sets the blend constant color used for blend operations that use a constant color.
func (r *RenderPassEncoder) SetBlendConstant(color Color) {
	cColor := C.WGPUColor{
		r: C.double(color.R),
		g: C.double(color.G),
		b: C.double(color.B),
		a: C.double(color.A),
	}

	C.wgpuRenderPassEncoderSetBlendConstant(r.ref, &cColor)
}

// SetViewport sets the viewport for subsequent render commands.
// The viewport defines the region of the render target that output is drawn to.
func (r *RenderPassEncoder) SetViewport(x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32) {
	C.wgpuRenderPassEncoderSetViewport(r.ref, C.float(x), C.float(y), C.float(width), C.float(height), C.float(minDepth), C.float(maxDepth))
}

// SetScissorRect sets the scissor rectangle for subsequent render commands.
// Pixels outside this rectangle will be discarded.
func (r *RenderPassEncoder) SetScissorRect(x uint32, y uint32, width uint32, height uint32) {
	C.wgpuRenderPassEncoderSetScissorRect(r.ref, C.uint32_t(x), C.uint32_t(y), C.uint32_t(width), C.uint32_t(height))
}

// SetVertexBuffer sets a vertex buffer at the specified slot for subsequent draw commands.
// The offset and size specify the region of the buffer to use.
func (r *RenderPassEncoder) SetVertexBuffer(slot uint32, buffer *Buffer, offset uint64, size uint64) {
	C.wgpuRenderPassEncoderSetVertexBuffer(r.ref, C.uint32_t(slot), buffer.ref, C.uint64_t(offset), C.uint64_t(size))
}

// SetIndexBuffer sets an index buffer for subsequent indexed draw commands.
// The format specifies the type of indices and the offset and size specify the region of the buffer to use.
func (r *RenderPassEncoder) SetIndexBuffer(buffer *Buffer, format IndexFormat, offset uint64, size uint64) {
	C.wgpuRenderPassEncoderSetIndexBuffer(r.ref, buffer.ref, C.WGPUIndexFormat(format), C.uint64_t(offset), C.uint64_t(size))
}

// BeginOcclusionQuery begins an occlusion query at the specified query index.
// Occlusion queries can be used to determine how many samples pass the depth test.
func (r *RenderPassEncoder) BeginOcclusionQuery(queryIndex uint32) {
	C.wgpuRenderPassEncoderBeginOcclusionQuery(r.ref, C.uint32_t(queryIndex))
}

// EndOcclusionQuery ends the current occlusion query.
func (r *RenderPassEncoder) EndOcclusionQuery() {
	C.wgpuRenderPassEncoderEndOcclusionQuery(r.ref)
}

// End ends the render pass.
// After calling this method, no more commands can be recorded in this pass.
func (r *RenderPassEncoder) End() {
	C.wgpuRenderPassEncoderEnd(r.ref)
}

// SetLabel sets the debug label for the render pass encoder.
// This label appears in debuggers and validation layers.
func (r *RenderPassEncoder) SetLabel(label string) {
	C.wgpuRenderPassEncoderSetLabel(r.ref, toCStr(label))
}
