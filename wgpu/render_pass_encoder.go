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

type RenderPassEncoder struct {
	ref C.WGPURenderPassEncoder
}

func (r *RenderPassEncoder) SetPipeline(pipeline *RenderPipeline) {
	C.wgpuRenderPassEncoderSetPipeline(r.ref, pipeline.ref)
}

func (r *RenderPassEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsets []uint32) {
	cDynamicOffsetsCount := C.size_t(len(dynamicOffsets))
	var cDynamicOffsets *C.uint32_t

	if cDynamicOffsetsCount > 0 {
		cDynamicOffsets = (*C.uint32_t)(&dynamicOffsets[0])
	}

	C.wgpuRenderPassEncoderSetBindGroup(r.ref, C.uint32_t(groupIndex), group.ref, cDynamicOffsetsCount, cDynamicOffsets)
}

func (r *RenderPassEncoder) Draw(vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	C.wgpuRenderPassEncoderDraw(r.ref, C.uint32_t(vertexCount), C.uint32_t(instanceCount), C.uint32_t(firstVertex), C.uint32_t(firstInstance))
}

func (r *RenderPassEncoder) DrawIndexed(indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	C.wgpuRenderPassEncoderDrawIndexed(r.ref, C.uint32_t(indexCount), C.uint32_t(instanceCount), C.uint32_t(firstIndex), C.int32_t(baseVertex), C.uint32_t(firstInstance))
}

func (r *RenderPassEncoder) DrawIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	C.wgpuRenderPassEncoderDrawIndirect(r.ref, indirectBuffer.ref, C.uint64_t(indirectOffset))
}

func (r *RenderPassEncoder) DrawIndexedIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	C.wgpuRenderPassEncoderDrawIndexedIndirect(r.ref, indirectBuffer.ref, C.uint64_t(indirectOffset))
}

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

func (r *RenderPassEncoder) InsertDebugMarker(markerLabel string) {
	C.wgpuRenderPassEncoderInsertDebugMarker(r.ref, toCStr(markerLabel))
}

func (r *RenderPassEncoder) PopDebugGroup() {
	C.wgpuRenderPassEncoderPopDebugGroup(r.ref)
}

func (r *RenderPassEncoder) PushDebugGroup(groupLabel string) {
	C.wgpuRenderPassEncoderPushDebugGroup(r.ref, toCStr(groupLabel))
}

func (r *RenderPassEncoder) SetStencilReference(reference uint32) {
	C.wgpuRenderPassEncoderSetStencilReference(r.ref, C.uint32_t(reference))
}

func (r *RenderPassEncoder) SetBlendConstant(color Color) {
	cColor := C.WGPUColor{
		r: C.double(color.R),
		g: C.double(color.G),
		b: C.double(color.B),
		a: C.double(color.A),
	}

	C.wgpuRenderPassEncoderSetBlendConstant(r.ref, &cColor)
}

func (r *RenderPassEncoder) SetViewport(x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32) {
	C.wgpuRenderPassEncoderSetViewport(r.ref, C.float(x), C.float(y), C.float(width), C.float(height), C.float(minDepth), C.float(maxDepth))
}

func (r *RenderPassEncoder) SetScissorRect(x uint32, y uint32, width uint32, height uint32) {
	C.wgpuRenderPassEncoderSetScissorRect(r.ref, C.uint32_t(x), C.uint32_t(y), C.uint32_t(width), C.uint32_t(height))
}

func (r *RenderPassEncoder) SetVertexBuffer(slot uint32, buffer *Buffer, offset uint64, size uint64) {
	C.wgpuRenderPassEncoderSetVertexBuffer(r.ref, C.uint32_t(slot), buffer.ref, C.uint64_t(offset), C.uint64_t(size))
}

func (r *RenderPassEncoder) SetIndexBuffer(buffer *Buffer, format IndexFormat, offset uint64, size uint64) {
	C.wgpuRenderPassEncoderSetIndexBuffer(r.ref, buffer.ref, C.WGPUIndexFormat(format), C.uint64_t(offset), C.uint64_t(size))
}

func (r *RenderPassEncoder) BeginOcclusionQuery(queryIndex uint32) {
	C.wgpuRenderPassEncoderBeginOcclusionQuery(r.ref, C.uint32_t(queryIndex))
}

func (r *RenderPassEncoder) EndOcclusionQuery() {
	C.wgpuRenderPassEncoderEndOcclusionQuery(r.ref)
}

func (r *RenderPassEncoder) End() {
	C.wgpuRenderPassEncoderEnd(r.ref)
}

func (r *RenderPassEncoder) SetLabel(label string) {
	C.wgpuRenderPassEncoderSetLabel(r.ref, toCStr(label))
}
