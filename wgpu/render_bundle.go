//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"
import "unsafe"

type RenderBundle struct {
	ref C.WGPURenderBundle
}

func (r *RenderBundle) SetLabel(label string) {
	C.wgpuRenderBundleSetLabel(r.ref, toCStr(label))
}

type RenderBundleEncoder struct {
	ref C.WGPURenderBundleEncoder
}

func (r *RenderBundleEncoder) SetPipeline(pipeline *RenderPipeline) {
	C.wgpuRenderBundleEncoderSetPipeline(r.ref, pipeline.ref)
}

func (r *RenderBundleEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsets []uint32) {
	cDynamicOffsetCount := C.size_t(len(dynamicOffsets))
	var cDynamicOffsets *C.uint32_t

	if cDynamicOffsetCount > 0 {
		cDynamicOffsets = (*C.uint32_t)(unsafe.Pointer(&dynamicOffsets[0]))
	}

	C.wgpuRenderBundleEncoderSetBindGroup(r.ref, C.uint32_t(groupIndex), group.ref, cDynamicOffsetCount, cDynamicOffsets)
}

func (r *RenderBundleEncoder) Draw(vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	C.wgpuRenderBundleEncoderDraw(r.ref, C.uint32_t(vertexCount), C.uint32_t(instanceCount), C.uint32_t(firstVertex), C.uint32_t(firstInstance))
}

func (r *RenderBundleEncoder) DrawIndexed(indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	C.wgpuRenderBundleEncoderDrawIndexed(r.ref, C.uint32_t(indexCount), C.uint32_t(instanceCount), C.uint32_t(firstIndex), C.int32_t(baseVertex), C.uint32_t(firstInstance))
}

func (r *RenderBundleEncoder) DrawIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	C.wgpuRenderBundleEncoderDrawIndirect(r.ref, indirectBuffer.ref, C.uint64_t(indirectOffset))
}

func (r *RenderBundleEncoder) DrawIndexedIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	C.wgpuRenderBundleEncoderDrawIndexedIndirect(r.ref, indirectBuffer.ref, C.uint64_t(indirectOffset))
}

func (r *RenderBundleEncoder) InsertDebugMarker(markerLabel string) {
	C.wgpuRenderBundleEncoderInsertDebugMarker(r.ref, toCStr(markerLabel))
}

func (r *RenderBundleEncoder) PopDebugGroup() {
	C.wgpuRenderBundleEncoderPopDebugGroup(r.ref)
}

func (r *RenderBundleEncoder) PushDebugGroup(groupLabel string) {
	C.wgpuRenderBundleEncoderPushDebugGroup(r.ref, toCStr(groupLabel))
}

func (r *RenderBundleEncoder) SetVertexBuffer(slot uint32, buffer *Buffer, offset uint64, size uint64) {
	C.wgpuRenderBundleEncoderSetVertexBuffer(r.ref, C.uint32_t(slot), buffer.ref, C.uint64_t(offset), C.uint64_t(size))
}

func (r *RenderBundleEncoder) SetIndexBuffer(buffer *Buffer, format IndexFormat, offset uint64, size uint64) {
	C.wgpuRenderBundleEncoderSetIndexBuffer(r.ref, buffer.ref, C.WGPUIndexFormat(format), C.uint64_t(offset), C.uint64_t(size))
}

func (r *RenderBundleEncoder) Finish(descriptor *RenderBundleDescriptor) RenderBundle {
	var cDescriptor *C.WGPURenderBundleDescriptor

	if descriptor != nil {
		cDescriptor = &C.WGPURenderBundleDescriptor{
			label: toCStr(descriptor.Label),
		}
	}

	return RenderBundle{ref: C.wgpuRenderBundleEncoderFinish(r.ref, cDescriptor)}
}

func (r *RenderBundleEncoder) SetLabel(label string) {
	C.wgpuRenderBundleEncoderSetLabel(r.ref, toCStr(label))
}
