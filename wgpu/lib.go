//go:build !js

package wgpu

/*
#cgo CFLAGS: -I./lib
#cgo LDFLAGS: -lwebgpu_dawn

// Android
#cgo android,amd64 LDFLAGS: -L${SRCDIR}/lib/android/amd64
#cgo android,386 LDFLAGS: -L${SRCDIR}/lib/android/386
#cgo android,arm64 LDFLAGS: -L${SRCDIR}/lib/android/arm64
#cgo android,arm LDFLAGS: -L${SRCDIR}/lib/android/arm
#cgo android LDFLAGS: -landroid -lm -llog

// Linux
#cgo linux,!android,amd64 LDFLAGS: -L${SRCDIR}/lib/linux/amd64
#cgo linux,!android LDFLAGS: -lm -ldl

// Darwin
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/lib/darwin/amd64
#cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/lib/darwin/arm64
#cgo darwin LDFLAGS: -framework Metal -framework IOKit -framework QuartzCore -framework Foundation -framework IOSurface -lc++

// Windows
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/lib/windows/amd64
#cgo windows LDFLAGS: -ld3dcompiler_47 -lws2_32 -luserenv -lbcrypt -lntdll

#include <stdio.h>
#include <stdlib.h>
#include "webgpu.h"

*/
import "C"

import (
	"unsafe"
)

type ComputePipeline struct {
	ref C.WGPUComputePipeline
}

func (c *ComputePipeline) GetBindGroupLayout(groupIndex uint32) BindGroupLayout {
	return BindGroupLayout{ref: C.wgpuComputePipelineGetBindGroupLayout(c.ref, C.uint32_t(groupIndex))}
}

func (c *ComputePipeline) SetLabel(label string) {
	C.wgpuComputePipelineSetLabel(c.ref, toCStr(label))
}

type PipelineLayout struct {
	ref C.WGPUPipelineLayout
}

func (p *PipelineLayout) SetLabel(label string) {
	C.wgpuPipelineLayoutSetLabel(p.ref, toCStr(label))
}

type QuerySet struct {
	ref C.WGPUQuerySet
}

func (q *QuerySet) SetLabel(label string) {
	C.wgpuQuerySetSetLabel(q.ref, toCStr(label))
}

func (q *QuerySet) GetType() QueryType {
	return QueryType(C.wgpuQuerySetGetType(q.ref))
}

func (q *QuerySet) GetCount() uint32 {
	return uint32(C.wgpuQuerySetGetCount(q.ref))
}

func (q *QuerySet) Destroy() {
	C.wgpuQuerySetDestroy(q.ref)
}

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
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))

	cSlot := C.uint32_t(slot)
	pBuffer := C.WGPUBuffer(unsafe.Pointer(buffer.ref))
	cOffset := C.uint64_t(offset)
	cSize := C.uint64_t(size)
	// Call and return
	C.wgpuRenderBundleEncoderSetVertexBuffer(cRenderBundleEncoder, cSlot, pBuffer, cOffset, cSize)
}

func (r *RenderBundleEncoder) SetIndexBuffer(buffer *Buffer, format IndexFormat, offset uint64, size uint64) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))

	pBuffer := C.WGPUBuffer(unsafe.Pointer(buffer.ref))
	cFormat := C.WGPUIndexFormat(format)
	cOffset := C.uint64_t(offset)
	cSize := C.uint64_t(size)
	// Call and return
	C.wgpuRenderBundleEncoderSetIndexBuffer(cRenderBundleEncoder, pBuffer, cFormat, cOffset, cSize)
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

type RenderPassEncoder struct {
	ref C.WGPURenderPassEncoder
}

func (r *RenderPassEncoder) SetPipeline(pipeline *RenderPipeline) {
	C.wgpuRenderPassEncoderSetPipeline(r.ref, pipeline.ref)
}

func (r *RenderPassEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsets []uint32) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	cGroupIndex := C.uint32_t(groupIndex)

	pGroup := C.WGPUBindGroup(unsafe.Pointer(group.ref))

	var cDynamicOffsetCount C.size_t
	var cDynamicOffsets *C.uint32_t

	dynamicOffsetsCount := len(dynamicOffsets)
	if dynamicOffsetsCount > 0 {
		offsets := C.malloc(C.size_t(dynamicOffsetsCount) * C.size_t(unsafe.Sizeof(C.uint32_t(0))))
		slice := unsafe.Slice((*uint32)(offsets), dynamicOffsetsCount)
		defer C.free(offsets)
		copy(slice, dynamicOffsets)

		cDynamicOffsets = (*C.uint32_t)(offsets)
		cDynamicOffsetCount = C.size_t(dynamicOffsetsCount)
	}

	C.wgpuRenderPassEncoderSetBindGroup(cRenderPassEncoder, cGroupIndex, pGroup, cDynamicOffsetCount, cDynamicOffsets)
}

func (r *RenderPassEncoder) Draw(vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	cVertexCount := C.uint32_t(vertexCount)
	cInstanceCount := C.uint32_t(instanceCount)
	cFirstVertex := C.uint32_t(firstVertex)
	cFirstInstance := C.uint32_t(firstInstance)
	// Call and return
	C.wgpuRenderPassEncoderDraw(cRenderPassEncoder, cVertexCount, cInstanceCount, cFirstVertex, cFirstInstance)
}

func (r *RenderPassEncoder) DrawIndexed(indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	cIndexCount := C.uint32_t(indexCount)
	cInstanceCount := C.uint32_t(instanceCount)
	cFirstIndex := C.uint32_t(firstIndex)
	cBaseVertex := C.int32_t(baseVertex)
	cFirstInstance := C.uint32_t(firstInstance)
	// Call and return
	C.wgpuRenderPassEncoderDrawIndexed(cRenderPassEncoder, cIndexCount, cInstanceCount, cFirstIndex, cBaseVertex, cFirstInstance)
}

func (r *RenderPassEncoder) DrawIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	pIndirectBuffer := C.WGPUBuffer(unsafe.Pointer(indirectBuffer.ref))
	cIndirectOffset := C.uint64_t(indirectOffset)
	// Call and return
	C.wgpuRenderPassEncoderDrawIndirect(cRenderPassEncoder, pIndirectBuffer, cIndirectOffset)
}

func (r *RenderPassEncoder) DrawIndexedIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	C.wgpuRenderPassEncoderDrawIndexedIndirect(r.ref, indirectBuffer.ref, C.uint64_t(indirectOffset))
}

func (r *RenderPassEncoder) ExecuteBundles(bundleCount int, bundles *RenderBundle) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	cBundleCount := C.size_t(bundleCount)
	pBundles := C.WGPURenderBundle(unsafe.Pointer(bundles.ref))
	// Call and return
	C.wgpuRenderPassEncoderExecuteBundles(cRenderPassEncoder, cBundleCount, &pBundles)
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
