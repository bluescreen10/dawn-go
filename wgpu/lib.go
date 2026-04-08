// CODE GENERATED. DO NOT EDIT
//
//go:generate go run ./cmd/wrapper/.
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

extern void cgo_callback_BufferMapCallback(WGPUMapAsyncStatus status, WGPUStringView message, void *userData1, void *userData2);
extern void cgo_callback_CompilationInfoCallback(WGPUCompilationInfoRequestStatus status, WGPUCompilationInfo compilationInfo, void *userData1, void *userData2);
extern void cgo_callback_CreateComputePipelineAsyncCallback(WGPUCreatePipelineAsyncStatus status, WGPUComputePipeline pipeline, WGPUStringView message, void *userData1, void *userData2);
extern void cgo_callback_CreateRenderPipelineAsyncCallback(WGPUCreatePipelineAsyncStatus status, WGPURenderPipeline pipeline, WGPUStringView message, void *userData1, void *userData2);
extern void cgo_callback_PopErrorScopeCallback(WGPUPopErrorScopeStatus status, WGPUErrorType typ, WGPUStringView message, void *userData1, void *userData2);
extern void cgo_callback_QueueWorkDoneCallback(WGPUQueueWorkDoneStatus status, WGPUStringView message, void *userData1, void *userData2);
extern void cgo_callback_RequestAdapterCallback(WGPURequestAdapterStatus status, WGPUAdapter adapter, WGPUStringView message, void *userData1, void *userData2);
*/
import "C"

import (
	"fmt"
	"runtime/cgo"
	"unsafe"
)

type BindGroup struct {
	ref uintptr
}

func (b *BindGroup) SetLabel(label string) {
	cBindGroup := C.WGPUBindGroup(unsafe.Pointer(b.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuBindGroupSetLabel(cBindGroup, cLabel)
}

type BindGroupLayout struct {
	ref uintptr
}

func (b *BindGroupLayout) SetLabel(label string) {
	cBindGroupLayout := C.WGPUBindGroupLayout(unsafe.Pointer(b.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuBindGroupLayoutSetLabel(cBindGroupLayout, cLabel)
}

type Buffer struct {
	ref uintptr
}

//export goBufferMapCallbackHandler
func goBufferMapCallbackHandler(status C.WGPUMapAsyncStatus, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	defer handle.Delete()
	fn := handle.Value().(BufferMapCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}

	fn(
		MapAsyncStatus(status),
		msg,
	)
}
func (b *Buffer) MapAsync(mode MapMode, offset int, size int, callbackInfo BufferMapCallbackInfo) Future {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))

	cMode := C.WGPUMapMode(mode)
	cOffset := C.size_t(offset)
	cSize := C.size_t(size)
	var cCallbackInfo C.WGPUBufferMapCallbackInfo
	cCallbackInfo.mode = C.WGPUCallbackMode(callbackInfo.Mode)
	//TODO: replace with callbacks
	//cCallbackInfo.callback = C.WGPUBufferMapCallback(callbackInfo.Callback)
	// Call and return
	return Future{Id: uint64(C.wgpuBufferMapAsync(cBuffer, cMode, cOffset, cSize, cCallbackInfo).id)}
}

func (b *Buffer) GetMappedRange(offset int, size int) uintptr {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))

	cOffset := C.size_t(offset)
	cSize := C.size_t(size)
	// Call and return
	return uintptr(C.wgpuBufferGetMappedRange(cBuffer, cOffset, cSize))
}

func (b *Buffer) GetConstMappedRange(offset int, size int) []byte {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))

	cOffset := C.size_t(offset)
	cSize := C.size_t(size)
	buf := C.wgpuBufferGetConstMappedRange(cBuffer, cOffset, cSize)

	return unsafe.Slice((*byte)(buf), size)
}

func (b *Buffer) WriteMappedRange(offset int, data []byte, size int) statusCode {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))

	cOffset := C.size_t(offset)
	var cData unsafe.Pointer
	if len(data) > 0 {
		cData = unsafe.Pointer(&data[0])
	}
	cSize := C.size_t(size)

	return statusCode(C.wgpuBufferWriteMappedRange(cBuffer, cOffset, cData, cSize))
}

func (b *Buffer) ReadMappedRange(offset int, data []byte, size int) statusCode {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))

	cOffset := C.size_t(offset)
	var cData unsafe.Pointer
	if len(data) > 0 {
		cData = unsafe.Pointer(&data[0])
	}
	cSize := C.size_t(size)
	// Call and return
	return statusCode(C.wgpuBufferReadMappedRange(cBuffer, cOffset, cData, cSize))
}

func (b *Buffer) SetLabel(label string) {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuBufferSetLabel(cBuffer, cLabel)
}

func (b *Buffer) GetUsage() BufferUsage {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))

	// Call and return
	return BufferUsage(C.wgpuBufferGetUsage(cBuffer))
}

func (b *Buffer) GetSize() uint64 {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))

	// Call and return
	return uint64(C.wgpuBufferGetSize(cBuffer))
}

func (b *Buffer) GetMapState() BufferMapState {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))

	// Call and return
	return BufferMapState(C.wgpuBufferGetMapState(cBuffer))
}

func (b *Buffer) Unmap() {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))

	// Call and return
	C.wgpuBufferUnmap(cBuffer)
}

func (b *Buffer) Destroy() {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))

	// Call and return
	C.wgpuBufferDestroy(cBuffer)
}

type CommandBuffer struct {
	ref uintptr
}

func (c *CommandBuffer) SetLabel(label string) {
	cCommandBuffer := C.WGPUCommandBuffer(unsafe.Pointer(c.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuCommandBufferSetLabel(cCommandBuffer, cLabel)
}

type CommandEncoder struct {
	ref uintptr
}

func (c *CommandEncoder) Finish(descriptor *CommandBufferDescriptor) CommandBuffer {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	var pDescriptor C.WGPUCommandBufferDescriptor
	if descriptor != nil {
		pDescriptorlabelStr := C.CString(descriptor.Label)
		defer C.free(unsafe.Pointer(pDescriptorlabelStr))
		pDescriptor.label.data = pDescriptorlabelStr
		pDescriptor.label.length = C.size_t(len(descriptor.Label))
	}
	// Call and return
	return CommandBuffer{ref: uintptr(unsafe.Pointer(C.wgpuCommandEncoderFinish(cCommandEncoder, &pDescriptor)))}
}

func (c *CommandEncoder) BeginComputePass(descriptor *ComputePassDescriptor) ComputePassEncoder {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	var pDescriptor C.WGPUComputePassDescriptor
	if descriptor != nil {
		pDescriptorlabelStr := C.CString(descriptor.Label)
		defer C.free(unsafe.Pointer(pDescriptorlabelStr))
		pDescriptor.label.data = pDescriptorlabelStr
		pDescriptor.label.length = C.size_t(len(descriptor.Label))

		if descriptor.TimestampWrites != nil {
			pDescriptor.timestampWrites.querySet = C.WGPUQuerySet(unsafe.Pointer(descriptor.TimestampWrites.QuerySet.ref))
			pDescriptor.timestampWrites.beginningOfPassWriteIndex = C.uint32_t(descriptor.TimestampWrites.BeginningOfPassWriteIndex)
			pDescriptor.timestampWrites.endOfPassWriteIndex = C.uint32_t(descriptor.TimestampWrites.EndOfPassWriteIndex)
		}
	}
	// Call and return
	return ComputePassEncoder{ref: uintptr(unsafe.Pointer(C.wgpuCommandEncoderBeginComputePass(cCommandEncoder, &pDescriptor)))}
}

func (c *CommandEncoder) BeginRenderPass(descriptor RenderPassDescriptor) RenderPassEncoder {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	var cDescriptor C.WGPURenderPassDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor.colorAttachmentCount = C.size_t(descriptor.ColorAttachmentCount)
	cDescriptor.colorAttachments.view = C.WGPUTextureView(unsafe.Pointer(descriptor.ColorAttachments.View.ref))
	cDescriptor.colorAttachments.depthSlice = C.uint32_t(descriptor.ColorAttachments.DepthSlice)
	cDescriptor.colorAttachments.resolveTarget = C.WGPUTextureView(unsafe.Pointer(descriptor.ColorAttachments.ResolveTarget.ref))
	cDescriptor.colorAttachments.loadOp = C.WGPULoadOp(descriptor.ColorAttachments.LoadOp)
	cDescriptor.colorAttachments.storeOp = C.WGPUStoreOp(descriptor.ColorAttachments.StoreOp)
	cDescriptor.colorAttachments.clearValue.r = C.double(descriptor.ColorAttachments.ClearValue.R)
	cDescriptor.colorAttachments.clearValue.g = C.double(descriptor.ColorAttachments.ClearValue.G)
	cDescriptor.colorAttachments.clearValue.b = C.double(descriptor.ColorAttachments.ClearValue.B)
	cDescriptor.colorAttachments.clearValue.a = C.double(descriptor.ColorAttachments.ClearValue.A)

	if descriptor.DepthStencilAttachment != nil {
		cDescriptor.depthStencilAttachment.view = C.WGPUTextureView(unsafe.Pointer(descriptor.DepthStencilAttachment.View.ref))
		cDescriptor.depthStencilAttachment.depthLoadOp = C.WGPULoadOp(descriptor.DepthStencilAttachment.DepthLoadOp)
		cDescriptor.depthStencilAttachment.depthStoreOp = C.WGPUStoreOp(descriptor.DepthStencilAttachment.DepthStoreOp)
		cDescriptor.depthStencilAttachment.depthClearValue = C.float(descriptor.DepthStencilAttachment.DepthClearValue)
		cDescriptor.depthStencilAttachment.depthReadOnly = boolToWGPUBool(descriptor.DepthStencilAttachment.DepthReadOnly)
		cDescriptor.depthStencilAttachment.stencilLoadOp = C.WGPULoadOp(descriptor.DepthStencilAttachment.StencilLoadOp)
		cDescriptor.depthStencilAttachment.stencilStoreOp = C.WGPUStoreOp(descriptor.DepthStencilAttachment.StencilStoreOp)
		cDescriptor.depthStencilAttachment.stencilClearValue = C.uint32_t(descriptor.DepthStencilAttachment.StencilClearValue)
		cDescriptor.depthStencilAttachment.stencilReadOnly = boolToWGPUBool(descriptor.DepthStencilAttachment.StencilReadOnly)
	}

	if descriptor.OcclusionQuerySet != nil {
		cDescriptor.occlusionQuerySet = C.WGPUQuerySet(unsafe.Pointer(descriptor.OcclusionQuerySet.ref))
	}

	if descriptor.TimestampWrites != nil {
		cDescriptor.timestampWrites.querySet = C.WGPUQuerySet(unsafe.Pointer(descriptor.TimestampWrites.QuerySet.ref))
		cDescriptor.timestampWrites.beginningOfPassWriteIndex = C.uint32_t(descriptor.TimestampWrites.BeginningOfPassWriteIndex)
		cDescriptor.timestampWrites.endOfPassWriteIndex = C.uint32_t(descriptor.TimestampWrites.EndOfPassWriteIndex)
	}
	// Call and return
	return RenderPassEncoder{ref: uintptr(unsafe.Pointer(C.wgpuCommandEncoderBeginRenderPass(cCommandEncoder, &cDescriptor)))}
}

func (c *CommandEncoder) CopyBufferToBuffer(source *Buffer, sourceOffset uint64, destination *Buffer, destinationOffset uint64, size uint64) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	pSource := C.WGPUBuffer(unsafe.Pointer(source.ref))
	cSourceOffset := C.uint64_t(sourceOffset)
	pDestination := C.WGPUBuffer(unsafe.Pointer(destination.ref))
	cDestinationOffset := C.uint64_t(destinationOffset)
	cSize := C.uint64_t(size)
	// Call and return
	C.wgpuCommandEncoderCopyBufferToBuffer(cCommandEncoder, pSource, cSourceOffset, pDestination, cDestinationOffset, cSize)
}

func (c *CommandEncoder) CopyBufferToTexture(source TexelCopyBufferInfo, destination TexelCopyTextureInfo, copySize Extent3D) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	var cSource C.WGPUTexelCopyBufferInfo
	cSource.layout.offset = C.uint64_t(source.Layout.Offset)
	cSource.layout.bytesPerRow = C.uint32_t(source.Layout.BytesPerRow)
	cSource.layout.rowsPerImage = C.uint32_t(source.Layout.RowsPerImage)
	cSource.buffer = C.WGPUBuffer(unsafe.Pointer(source.Buffer.ref))
	var cDestination C.WGPUTexelCopyTextureInfo
	cDestination.texture = C.WGPUTexture(unsafe.Pointer(destination.Texture.ref))
	cDestination.mipLevel = C.uint32_t(destination.MipLevel)
	cDestination.origin.x = C.uint32_t(destination.Origin.X)
	cDestination.origin.y = C.uint32_t(destination.Origin.Y)
	cDestination.origin.z = C.uint32_t(destination.Origin.Z)
	cDestination.aspect = C.WGPUTextureAspect(destination.Aspect)
	var cCopySize C.WGPUExtent3D
	cCopySize.width = C.uint32_t(copySize.Width)
	cCopySize.height = C.uint32_t(copySize.Height)
	cCopySize.depthOrArrayLayers = C.uint32_t(copySize.DepthOrArrayLayers)
	// Call and return
	C.wgpuCommandEncoderCopyBufferToTexture(cCommandEncoder, &cSource, &cDestination, &cCopySize)
}

func (c *CommandEncoder) CopyTextureToBuffer(source TexelCopyTextureInfo, destination TexelCopyBufferInfo, copySize Extent3D) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	var cSource C.WGPUTexelCopyTextureInfo
	cSource.texture = C.WGPUTexture(unsafe.Pointer(source.Texture.ref))
	cSource.mipLevel = C.uint32_t(source.MipLevel)
	cSource.origin.x = C.uint32_t(source.Origin.X)
	cSource.origin.y = C.uint32_t(source.Origin.Y)
	cSource.origin.z = C.uint32_t(source.Origin.Z)
	cSource.aspect = C.WGPUTextureAspect(source.Aspect)
	var cDestination C.WGPUTexelCopyBufferInfo
	cDestination.layout.offset = C.uint64_t(destination.Layout.Offset)
	cDestination.layout.bytesPerRow = C.uint32_t(destination.Layout.BytesPerRow)
	cDestination.layout.rowsPerImage = C.uint32_t(destination.Layout.RowsPerImage)
	cDestination.buffer = C.WGPUBuffer(unsafe.Pointer(destination.Buffer.ref))
	var cCopySize C.WGPUExtent3D
	cCopySize.width = C.uint32_t(copySize.Width)
	cCopySize.height = C.uint32_t(copySize.Height)
	cCopySize.depthOrArrayLayers = C.uint32_t(copySize.DepthOrArrayLayers)
	// Call and return
	C.wgpuCommandEncoderCopyTextureToBuffer(cCommandEncoder, &cSource, &cDestination, &cCopySize)
}

func (c *CommandEncoder) CopyTextureToTexture(source TexelCopyTextureInfo, destination TexelCopyTextureInfo, copySize Extent3D) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	var cSource C.WGPUTexelCopyTextureInfo
	cSource.texture = C.WGPUTexture(unsafe.Pointer(source.Texture.ref))
	cSource.mipLevel = C.uint32_t(source.MipLevel)
	cSource.origin.x = C.uint32_t(source.Origin.X)
	cSource.origin.y = C.uint32_t(source.Origin.Y)
	cSource.origin.z = C.uint32_t(source.Origin.Z)
	cSource.aspect = C.WGPUTextureAspect(source.Aspect)
	var cDestination C.WGPUTexelCopyTextureInfo
	cDestination.texture = C.WGPUTexture(unsafe.Pointer(destination.Texture.ref))
	cDestination.mipLevel = C.uint32_t(destination.MipLevel)
	cDestination.origin.x = C.uint32_t(destination.Origin.X)
	cDestination.origin.y = C.uint32_t(destination.Origin.Y)
	cDestination.origin.z = C.uint32_t(destination.Origin.Z)
	cDestination.aspect = C.WGPUTextureAspect(destination.Aspect)
	var cCopySize C.WGPUExtent3D
	cCopySize.width = C.uint32_t(copySize.Width)
	cCopySize.height = C.uint32_t(copySize.Height)
	cCopySize.depthOrArrayLayers = C.uint32_t(copySize.DepthOrArrayLayers)
	// Call and return
	C.wgpuCommandEncoderCopyTextureToTexture(cCommandEncoder, &cSource, &cDestination, &cCopySize)
}

func (c *CommandEncoder) ClearBuffer(buffer *Buffer, offset uint64, size uint64) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	pBuffer := C.WGPUBuffer(unsafe.Pointer(buffer.ref))
	cOffset := C.uint64_t(offset)
	cSize := C.uint64_t(size)
	// Call and return
	C.wgpuCommandEncoderClearBuffer(cCommandEncoder, pBuffer, cOffset, cSize)
}

func (c *CommandEncoder) InsertDebugMarker(markerLabel string) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	cMarkerLabelStr := C.CString(markerLabel)
	defer C.free(unsafe.Pointer(cMarkerLabelStr))
	var cMarkerLabel C.WGPUStringView
	cMarkerLabel.data = cMarkerLabelStr
	cMarkerLabel.length = C.size_t(len(markerLabel))
	// Call and return
	C.wgpuCommandEncoderInsertDebugMarker(cCommandEncoder, cMarkerLabel)
}

func (c *CommandEncoder) PopDebugGroup() {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	// Call and return
	C.wgpuCommandEncoderPopDebugGroup(cCommandEncoder)
}

func (c *CommandEncoder) PushDebugGroup(groupLabel string) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	cGroupLabelStr := C.CString(groupLabel)
	defer C.free(unsafe.Pointer(cGroupLabelStr))
	var cGroupLabel C.WGPUStringView
	cGroupLabel.data = cGroupLabelStr
	cGroupLabel.length = C.size_t(len(groupLabel))
	// Call and return
	C.wgpuCommandEncoderPushDebugGroup(cCommandEncoder, cGroupLabel)
}

func (c *CommandEncoder) ResolveQuerySet(querySet *QuerySet, firstQuery uint32, queryCount uint32, destination *Buffer, destinationOffset uint64) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	pQuerySet := C.WGPUQuerySet(unsafe.Pointer(querySet.ref))
	cFirstQuery := C.uint32_t(firstQuery)
	cQueryCount := C.uint32_t(queryCount)
	pDestination := C.WGPUBuffer(unsafe.Pointer(destination.ref))
	cDestinationOffset := C.uint64_t(destinationOffset)
	// Call and return
	C.wgpuCommandEncoderResolveQuerySet(cCommandEncoder, pQuerySet, cFirstQuery, cQueryCount, pDestination, cDestinationOffset)
}

func (c *CommandEncoder) WriteTimestamp(querySet *QuerySet, queryIndex uint32) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	pQuerySet := C.WGPUQuerySet(unsafe.Pointer(querySet.ref))
	cQueryIndex := C.uint32_t(queryIndex)
	// Call and return
	C.wgpuCommandEncoderWriteTimestamp(cCommandEncoder, pQuerySet, cQueryIndex)
}

func (c *CommandEncoder) SetLabel(label string) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuCommandEncoderSetLabel(cCommandEncoder, cLabel)
}

type ComputePassEncoder struct {
	ref uintptr
}

func (c *ComputePassEncoder) InsertDebugMarker(markerLabel string) {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))

	cMarkerLabelStr := C.CString(markerLabel)
	defer C.free(unsafe.Pointer(cMarkerLabelStr))
	var cMarkerLabel C.WGPUStringView
	cMarkerLabel.data = cMarkerLabelStr
	cMarkerLabel.length = C.size_t(len(markerLabel))
	// Call and return
	C.wgpuComputePassEncoderInsertDebugMarker(cComputePassEncoder, cMarkerLabel)
}

func (c *ComputePassEncoder) PopDebugGroup() {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))

	// Call and return
	C.wgpuComputePassEncoderPopDebugGroup(cComputePassEncoder)
}

func (c *ComputePassEncoder) PushDebugGroup(groupLabel string) {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))

	cGroupLabelStr := C.CString(groupLabel)
	defer C.free(unsafe.Pointer(cGroupLabelStr))
	var cGroupLabel C.WGPUStringView
	cGroupLabel.data = cGroupLabelStr
	cGroupLabel.length = C.size_t(len(groupLabel))
	// Call and return
	C.wgpuComputePassEncoderPushDebugGroup(cComputePassEncoder, cGroupLabel)
}

func (c *ComputePassEncoder) SetPipeline(pipeline *ComputePipeline) {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))

	pPipeline := C.WGPUComputePipeline(unsafe.Pointer(pipeline.ref))
	// Call and return
	C.wgpuComputePassEncoderSetPipeline(cComputePassEncoder, pPipeline)
}

func (c *ComputePassEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsetCount int, dynamicOffsets uint32) {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))

	cGroupIndex := C.uint32_t(groupIndex)
	pGroup := C.WGPUBindGroup(unsafe.Pointer(group.ref))
	cDynamicOffsetCount := C.size_t(dynamicOffsetCount)
	cDynamicOffsets := C.uint32_t(dynamicOffsets)
	// Call and return
	C.wgpuComputePassEncoderSetBindGroup(cComputePassEncoder, cGroupIndex, pGroup, cDynamicOffsetCount, &cDynamicOffsets)
}

func (c *ComputePassEncoder) DispatchWorkgroups(workgroupCountX uint32, workgroupCountY uint32, workgroupCountZ uint32) {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))

	cWorkgroupCountX := C.uint32_t(workgroupCountX)
	cWorkgroupCountY := C.uint32_t(workgroupCountY)
	cWorkgroupCountZ := C.uint32_t(workgroupCountZ)
	// Call and return
	C.wgpuComputePassEncoderDispatchWorkgroups(cComputePassEncoder, cWorkgroupCountX, cWorkgroupCountY, cWorkgroupCountZ)
}

func (c *ComputePassEncoder) DispatchWorkgroupsIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))

	pIndirectBuffer := C.WGPUBuffer(unsafe.Pointer(indirectBuffer.ref))
	cIndirectOffset := C.uint64_t(indirectOffset)
	// Call and return
	C.wgpuComputePassEncoderDispatchWorkgroupsIndirect(cComputePassEncoder, pIndirectBuffer, cIndirectOffset)
}

func (c *ComputePassEncoder) End() {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))

	// Call and return
	C.wgpuComputePassEncoderEnd(cComputePassEncoder)
}

func (c *ComputePassEncoder) SetLabel(label string) {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuComputePassEncoderSetLabel(cComputePassEncoder, cLabel)
}

type ComputePipeline struct {
	ref uintptr
}

func (c *ComputePipeline) GetBindGroupLayout(groupIndex uint32) BindGroupLayout {
	cComputePipeline := C.WGPUComputePipeline(unsafe.Pointer(c.ref))

	cGroupIndex := C.uint32_t(groupIndex)
	// Call and return
	return BindGroupLayout{ref: uintptr(unsafe.Pointer(C.wgpuComputePipelineGetBindGroupLayout(cComputePipeline, cGroupIndex)))}
}

func (c *ComputePipeline) SetLabel(label string) {
	cComputePipeline := C.WGPUComputePipeline(unsafe.Pointer(c.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuComputePipelineSetLabel(cComputePipeline, cLabel)
}

type Device struct {
	ref uintptr
}

func (d *Device) CreateBindGroup(descriptor BindGroupDescriptor) BindGroup {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUBindGroupDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor.layout = C.WGPUBindGroupLayout(unsafe.Pointer(descriptor.Layout.ref))
	cDescriptor.entryCount = C.size_t(descriptor.EntryCount)
	cDescriptor.entries.binding = C.uint32_t(descriptor.Entries.Binding)
	cDescriptor.entries.buffer = C.WGPUBuffer(unsafe.Pointer(descriptor.Entries.Buffer.ref))
	cDescriptor.entries.offset = C.uint64_t(descriptor.Entries.Offset)
	cDescriptor.entries.size = C.uint64_t(descriptor.Entries.Size)
	cDescriptor.entries.sampler = C.WGPUSampler(unsafe.Pointer(descriptor.Entries.Sampler.ref))
	cDescriptor.entries.textureView = C.WGPUTextureView(unsafe.Pointer(descriptor.Entries.TextureView.ref))
	// Call and return
	return BindGroup{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateBindGroup(cDevice, &cDescriptor)))}
}

func (d *Device) CreateBindGroupLayout(descriptor BindGroupLayoutDescriptor) BindGroupLayout {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUBindGroupLayoutDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor.entryCount = C.size_t(descriptor.EntryCount)
	cDescriptor.entries.binding = C.uint32_t(descriptor.Entries.Binding)
	cDescriptor.entries.visibility = C.WGPUShaderStage(descriptor.Entries.Visibility)
	cDescriptor.entries.bindingArraySize = C.uint32_t(descriptor.Entries.BindingArraySize)
	cDescriptor.entries.buffer._type = C.WGPUBufferBindingType(descriptor.Entries.Buffer.Type)
	cDescriptor.entries.buffer.hasDynamicOffset = boolToWGPUBool(descriptor.Entries.Buffer.HasDynamicOffset)
	cDescriptor.entries.buffer.minBindingSize = C.uint64_t(descriptor.Entries.Buffer.MinBindingSize)
	cDescriptor.entries.sampler._type = C.WGPUSamplerBindingType(descriptor.Entries.Sampler.Type)
	cDescriptor.entries.texture.sampleType = C.WGPUTextureSampleType(descriptor.Entries.Texture.SampleType)
	cDescriptor.entries.texture.viewDimension = C.WGPUTextureViewDimension(descriptor.Entries.Texture.ViewDimension)
	cDescriptor.entries.texture.multisampled = boolToWGPUBool(descriptor.Entries.Texture.Multisampled)
	cDescriptor.entries.storageTexture.access = C.WGPUStorageTextureAccess(descriptor.Entries.StorageTexture.Access)
	cDescriptor.entries.storageTexture.format = C.WGPUTextureFormat(descriptor.Entries.StorageTexture.Format)
	cDescriptor.entries.storageTexture.viewDimension = C.WGPUTextureViewDimension(descriptor.Entries.StorageTexture.ViewDimension)
	// Call and return
	return BindGroupLayout{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateBindGroupLayout(cDevice, &cDescriptor)))}
}

func (d *Device) CreateBuffer(descriptor BufferDescriptor) *Buffer {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUBufferDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor.usage = C.WGPUBufferUsage(descriptor.Usage)
	cDescriptor.size = C.uint64_t(descriptor.Size)
	cDescriptor.mappedAtCreation = boolToWGPUBool(descriptor.MappedAtCreation)
	// Call and return
	return &Buffer{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateBuffer(cDevice, &cDescriptor)))}
}

func (d *Device) CreateCommandEncoder(descriptor *CommandEncoderDescriptor) CommandEncoder {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var pDescriptor C.WGPUCommandEncoderDescriptor
	if descriptor != nil {
		pDescriptorlabelStr := C.CString(descriptor.Label)
		defer C.free(unsafe.Pointer(pDescriptorlabelStr))
		pDescriptor.label.data = pDescriptorlabelStr
		pDescriptor.label.length = C.size_t(len(descriptor.Label))
	}
	// Call and return
	return CommandEncoder{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateCommandEncoder(cDevice, &pDescriptor)))}
}

func (d *Device) CreateComputePipeline(descriptor ComputePipelineDescriptor) ComputePipeline {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUComputePipelineDescriptor

	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))

	cDescriptor.layout = C.WGPUPipelineLayout(unsafe.Pointer(descriptor.Layout.ref))
	cDescriptor.compute.module = C.WGPUShaderModule(unsafe.Pointer(descriptor.Compute.Module.ref))

	cDescriptorcomputeentryPointStr := C.CString(descriptor.Compute.EntryPoint)
	defer C.free(unsafe.Pointer(cDescriptorcomputeentryPointStr))
	cDescriptor.compute.entryPoint.data = cDescriptorcomputeentryPointStr
	cDescriptor.compute.entryPoint.length = C.size_t(len(descriptor.Compute.EntryPoint))
	cDescriptor.compute.constantCount = C.size_t(descriptor.Compute.ConstantCount)
	cDescriptorcomputeconstantskeyStr := C.CString(descriptor.Compute.Constants.Key)
	defer C.free(unsafe.Pointer(cDescriptorcomputeconstantskeyStr))
	cDescriptor.compute.constants.key.data = cDescriptorcomputeconstantskeyStr
	cDescriptor.compute.constants.key.length = C.size_t(len(descriptor.Compute.Constants.Key))
	cDescriptor.compute.constants.value = C.double(descriptor.Compute.Constants.Value)

	return ComputePipeline{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateComputePipeline(cDevice, &cDescriptor)))}
}

//export goCreateComputePipelineAsyncCallbackHandler
func goCreateComputePipelineAsyncCallbackHandler(status C.WGPUCreatePipelineAsyncStatus, pipeline C.WGPUComputePipeline, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	defer handle.Delete()
	fn := handle.Value().(CreateComputePipelineAsyncCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}

	fn(
		CreatePipelineAsyncStatus(status),
		&ComputePipeline{ref: uintptr(unsafe.Pointer(pipeline))},
		msg,
	)
}
func (d *Device) CreateComputePipelineAsync(descriptor ComputePipelineDescriptor, callback CreateComputePipelineAsyncCallback) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUComputePipelineDescriptor
	cDescriptor.layout = C.WGPUPipelineLayout(unsafe.Pointer(descriptor.Layout.ref))
	cDescriptor.compute.module = C.WGPUShaderModule(unsafe.Pointer(descriptor.Compute.Module.ref))
	cDescriptor.compute.constantCount = C.size_t(descriptor.Compute.ConstantCount)
	cDescriptorcomputeconstantskeyStr := C.CString(descriptor.Compute.Constants.Key)
	cDescriptor.compute.constants.value = C.double(descriptor.Compute.Constants.Value)

	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))

	cDescriptorcomputeentryPointStr := C.CString(descriptor.Compute.EntryPoint)
	defer C.free(unsafe.Pointer(cDescriptorcomputeentryPointStr))
	cDescriptor.compute.entryPoint.data = cDescriptorcomputeentryPointStr
	cDescriptor.compute.entryPoint.length = C.size_t(len(descriptor.Compute.EntryPoint))

	defer C.free(unsafe.Pointer(cDescriptorcomputeconstantskeyStr))
	cDescriptor.compute.constants.key.data = cDescriptorcomputeconstantskeyStr
	cDescriptor.compute.constants.key.length = C.size_t(len(descriptor.Compute.Constants.Key))

	// Setup callback
	handle := cgo.NewHandle(callback)
	var cCallbackInfo C.WGPUCreateComputePipelineAsyncCallbackInfo
	cCallbackInfo.mode = C.WGPUCallbackMode(callbackModeAllowProcessEvents)
	cCallbackInfo.callback = C.WGPUCreateComputePipelineAsyncCallback(C.cgo_callback_CreateComputePipelineAsyncCallback)
	cCallbackInfo.userdata1 = unsafe.Pointer(handle)
	cCallbackInfo.userdata2 = nil

	C.wgpuDeviceCreateComputePipelineAsync(cDevice, &cDescriptor, cCallbackInfo)
}

func (d *Device) CreatePipelineLayout(descriptor PipelineLayoutDescriptor) PipelineLayout {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUPipelineLayoutDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor.bindGroupLayoutCount = C.size_t(descriptor.BindGroupLayoutCount)
	cDescriptor.bindGroupLayouts = (*C.WGPUBindGroupLayout)(unsafe.Pointer(descriptor.BindGroupLayouts.ref))
	cDescriptor.immediateSize = C.uint32_t(descriptor.ImmediateSize)
	// Call and return
	return PipelineLayout{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreatePipelineLayout(cDevice, &cDescriptor)))}
}

func (d *Device) CreateQuerySet(descriptor QuerySetDescriptor) QuerySet {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUQuerySetDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor._type = C.WGPUQueryType(descriptor.Type)
	cDescriptor.count = C.uint32_t(descriptor.Count)
	// Call and return
	return QuerySet{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateQuerySet(cDevice, &cDescriptor)))}
}

//export goCreateRenderPipelineAsyncCallbackHandler
func goCreateRenderPipelineAsyncCallbackHandler(status C.WGPUCreatePipelineAsyncStatus, pipeline C.WGPURenderPipeline, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	defer handle.Delete()
	fn := handle.Value().(CreateRenderPipelineAsyncCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}
	fn(
		CreatePipelineAsyncStatus(status),
		&RenderPipeline{ref: uintptr(unsafe.Pointer(pipeline))},
		msg,
	)
}
func (d *Device) CreateRenderPipelineAsync(descriptor RenderPipelineDescriptor, callback CreateRenderPipelineAsyncCallback) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPURenderPipelineDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor.layout = C.WGPUPipelineLayout(unsafe.Pointer(descriptor.Layout.ref))
	cDescriptor.vertex.module = C.WGPUShaderModule(unsafe.Pointer(descriptor.Vertex.Module.ref))
	cDescriptorvertexentryPointStr := C.CString(descriptor.Vertex.EntryPoint)
	defer C.free(unsafe.Pointer(cDescriptorvertexentryPointStr))
	cDescriptor.vertex.entryPoint.data = cDescriptorvertexentryPointStr
	cDescriptor.vertex.entryPoint.length = C.size_t(len(descriptor.Vertex.EntryPoint))
	cDescriptor.vertex.constantCount = C.size_t(descriptor.Vertex.ConstantCount)
	cDescriptorvertexconstantskeyStr := C.CString(descriptor.Vertex.Constants.Key)
	defer C.free(unsafe.Pointer(cDescriptorvertexconstantskeyStr))
	cDescriptor.vertex.constants.key.data = cDescriptorvertexconstantskeyStr
	cDescriptor.vertex.constants.key.length = C.size_t(len(descriptor.Vertex.Constants.Key))
	cDescriptor.vertex.constants.value = C.double(descriptor.Vertex.Constants.Value)
	cDescriptor.vertex.bufferCount = C.size_t(descriptor.Vertex.BufferCount)
	cDescriptor.vertex.buffers.stepMode = C.WGPUVertexStepMode(descriptor.Vertex.Buffers.StepMode)
	cDescriptor.vertex.buffers.arrayStride = C.uint64_t(descriptor.Vertex.Buffers.ArrayStride)
	cDescriptor.vertex.buffers.attributeCount = C.size_t(descriptor.Vertex.Buffers.AttributeCount)
	cDescriptor.vertex.buffers.attributes.format = C.WGPUVertexFormat(descriptor.Vertex.Buffers.Attributes.Format)
	cDescriptor.vertex.buffers.attributes.offset = C.uint64_t(descriptor.Vertex.Buffers.Attributes.Offset)
	cDescriptor.vertex.buffers.attributes.shaderLocation = C.uint32_t(descriptor.Vertex.Buffers.Attributes.ShaderLocation)
	cDescriptor.primitive.topology = C.WGPUPrimitiveTopology(descriptor.Primitive.Topology)
	cDescriptor.primitive.stripIndexFormat = C.WGPUIndexFormat(descriptor.Primitive.StripIndexFormat)
	cDescriptor.primitive.frontFace = C.WGPUFrontFace(descriptor.Primitive.FrontFace)
	cDescriptor.primitive.cullMode = C.WGPUCullMode(descriptor.Primitive.CullMode)
	cDescriptor.primitive.unclippedDepth = boolToWGPUBool(descriptor.Primitive.UnclippedDepth)

	if descriptor.DepthStencil != nil {
		cDescriptor.depthStencil.format = C.WGPUTextureFormat(descriptor.DepthStencil.Format)
		cDescriptor.depthStencil.depthWriteEnabled = C.WGPUOptionalBool(descriptor.DepthStencil.DepthWriteEnabled)
		cDescriptor.depthStencil.depthCompare = C.WGPUCompareFunction(descriptor.DepthStencil.DepthCompare)
		cDescriptor.depthStencil.stencilFront.compare = C.WGPUCompareFunction(descriptor.DepthStencil.StencilFront.Compare)
		cDescriptor.depthStencil.stencilFront.failOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.FailOp)
		cDescriptor.depthStencil.stencilFront.depthFailOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.DepthFailOp)
		cDescriptor.depthStencil.stencilFront.passOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.PassOp)
		cDescriptor.depthStencil.stencilBack.compare = C.WGPUCompareFunction(descriptor.DepthStencil.StencilBack.Compare)
		cDescriptor.depthStencil.stencilBack.failOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.FailOp)
		cDescriptor.depthStencil.stencilBack.depthFailOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.DepthFailOp)
		cDescriptor.depthStencil.stencilBack.passOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.PassOp)
		cDescriptor.depthStencil.stencilReadMask = C.uint32_t(descriptor.DepthStencil.StencilReadMask)
		cDescriptor.depthStencil.stencilWriteMask = C.uint32_t(descriptor.DepthStencil.StencilWriteMask)
		cDescriptor.depthStencil.depthBias = C.int32_t(descriptor.DepthStencil.DepthBias)
		cDescriptor.depthStencil.depthBiasSlopeScale = C.float(descriptor.DepthStencil.DepthBiasSlopeScale)
		cDescriptor.depthStencil.depthBiasClamp = C.float(descriptor.DepthStencil.DepthBiasClamp)
	}

	cDescriptor.multisample.count = C.uint32_t(descriptor.Multisample.Count)
	cDescriptor.multisample.mask = C.uint32_t(descriptor.Multisample.Mask)
	cDescriptor.multisample.alphaToCoverageEnabled = boolToWGPUBool(descriptor.Multisample.AlphaToCoverageEnabled)

	if descriptor.Fragment != nil {
		cDescriptor.fragment.module = C.WGPUShaderModule(unsafe.Pointer(descriptor.Fragment.Module.ref))
		cDescriptorfragmententryPointStr := C.CString(descriptor.Fragment.EntryPoint)
		defer C.free(unsafe.Pointer(cDescriptorfragmententryPointStr))
		cDescriptor.fragment.entryPoint.data = cDescriptorfragmententryPointStr
		cDescriptor.fragment.entryPoint.length = C.size_t(len(descriptor.Fragment.EntryPoint))
		cDescriptor.fragment.constantCount = C.size_t(descriptor.Fragment.ConstantCount)
		cDescriptorfragmentconstantskeyStr := C.CString(descriptor.Fragment.Constants.Key)
		defer C.free(unsafe.Pointer(cDescriptorfragmentconstantskeyStr))
		cDescriptor.fragment.constants.key.data = cDescriptorfragmentconstantskeyStr
		cDescriptor.fragment.constants.key.length = C.size_t(len(descriptor.Fragment.Constants.Key))
		cDescriptor.fragment.constants.value = C.double(descriptor.Fragment.Constants.Value)
		cDescriptor.fragment.targetCount = C.size_t(descriptor.Fragment.TargetCount)
		cDescriptor.fragment.targets.format = C.WGPUTextureFormat(descriptor.Fragment.Targets.Format)
		if descriptor.Fragment.Targets.Blend != nil {
			cDescriptor.fragment.targets.blend.color.operation = C.WGPUBlendOperation(descriptor.Fragment.Targets.Blend.Color.Operation)
			cDescriptor.fragment.targets.blend.color.srcFactor = C.WGPUBlendFactor(descriptor.Fragment.Targets.Blend.Color.SrcFactor)
			cDescriptor.fragment.targets.blend.color.dstFactor = C.WGPUBlendFactor(descriptor.Fragment.Targets.Blend.Color.DstFactor)
			cDescriptor.fragment.targets.blend.alpha.operation = C.WGPUBlendOperation(descriptor.Fragment.Targets.Blend.Alpha.Operation)
			cDescriptor.fragment.targets.blend.alpha.srcFactor = C.WGPUBlendFactor(descriptor.Fragment.Targets.Blend.Alpha.SrcFactor)
			cDescriptor.fragment.targets.blend.alpha.dstFactor = C.WGPUBlendFactor(descriptor.Fragment.Targets.Blend.Alpha.DstFactor)
		}
		cDescriptor.fragment.targets.writeMask = C.WGPUColorWriteMask(descriptor.Fragment.Targets.WriteMask)
	}

	handle := cgo.NewHandle(callback)

	var cCallbackInfo C.WGPUCreateRenderPipelineAsyncCallbackInfo
	cCallbackInfo.mode = C.WGPUCallbackMode(callbackModeAllowProcessEvents)
	cCallbackInfo.callback = C.WGPUCreateRenderPipelineAsyncCallback(C.cgo_callback_CreateRenderPipelineAsyncCallback)
	cCallbackInfo.userdata1 = unsafe.Pointer(handle)
	cCallbackInfo.userdata2 = nil

	C.wgpuDeviceCreateRenderPipelineAsync(cDevice, &cDescriptor, cCallbackInfo)
}

func (d *Device) CreateRenderBundleEncoder(descriptor RenderBundleEncoderDescriptor) RenderBundleEncoder {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPURenderBundleEncoderDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor.depthStencilFormat = C.WGPUTextureFormat(descriptor.DepthStencilFormat)
	cDescriptor.sampleCount = C.uint32_t(descriptor.SampleCount)
	cDescriptor.depthReadOnly = boolToWGPUBool(descriptor.DepthReadOnly)
	cDescriptor.stencilReadOnly = boolToWGPUBool(descriptor.StencilReadOnly)

	colorFormatCount := len(descriptor.ColorFormats)
	if colorFormatCount > 0 {
		colorFormats := C.malloc(C.size_t(colorFormatCount) * C.size_t(unsafe.Sizeof(C.WGPUTextureFormat(0))))
		defer C.free(colorFormats)

		slice := unsafe.Slice((*TextureFormat)(colorFormats), colorFormatCount)
		copy(slice, descriptor.ColorFormats)

		cDescriptor.colorFormats = (*C.WGPUTextureFormat)(colorFormats)
		cDescriptor.colorFormatCount = C.size_t(colorFormatCount)
	}

	return RenderBundleEncoder{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateRenderBundleEncoder(cDevice, &cDescriptor)))}
}

func (d *Device) CreateRenderPipeline(descriptor RenderPipelineDescriptor) RenderPipeline {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPURenderPipelineDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor.layout = C.WGPUPipelineLayout(unsafe.Pointer(descriptor.Layout.ref))
	cDescriptor.vertex.module = C.WGPUShaderModule(unsafe.Pointer(descriptor.Vertex.Module.ref))
	cDescriptorvertexentryPointStr := C.CString(descriptor.Vertex.EntryPoint)
	defer C.free(unsafe.Pointer(cDescriptorvertexentryPointStr))
	cDescriptor.vertex.entryPoint.data = cDescriptorvertexentryPointStr
	cDescriptor.vertex.entryPoint.length = C.size_t(len(descriptor.Vertex.EntryPoint))
	cDescriptor.vertex.constantCount = C.size_t(descriptor.Vertex.ConstantCount)
	cDescriptorvertexconstantskeyStr := C.CString(descriptor.Vertex.Constants.Key)
	defer C.free(unsafe.Pointer(cDescriptorvertexconstantskeyStr))
	cDescriptor.vertex.constants.key.data = cDescriptorvertexconstantskeyStr
	cDescriptor.vertex.constants.key.length = C.size_t(len(descriptor.Vertex.Constants.Key))
	cDescriptor.vertex.constants.value = C.double(descriptor.Vertex.Constants.Value)
	cDescriptor.vertex.bufferCount = C.size_t(descriptor.Vertex.BufferCount)
	cDescriptor.vertex.buffers.stepMode = C.WGPUVertexStepMode(descriptor.Vertex.Buffers.StepMode)
	cDescriptor.vertex.buffers.arrayStride = C.uint64_t(descriptor.Vertex.Buffers.ArrayStride)
	cDescriptor.vertex.buffers.attributeCount = C.size_t(descriptor.Vertex.Buffers.AttributeCount)
	cDescriptor.vertex.buffers.attributes.format = C.WGPUVertexFormat(descriptor.Vertex.Buffers.Attributes.Format)
	cDescriptor.vertex.buffers.attributes.offset = C.uint64_t(descriptor.Vertex.Buffers.Attributes.Offset)
	cDescriptor.vertex.buffers.attributes.shaderLocation = C.uint32_t(descriptor.Vertex.Buffers.Attributes.ShaderLocation)
	cDescriptor.primitive.topology = C.WGPUPrimitiveTopology(descriptor.Primitive.Topology)
	cDescriptor.primitive.stripIndexFormat = C.WGPUIndexFormat(descriptor.Primitive.StripIndexFormat)
	cDescriptor.primitive.frontFace = C.WGPUFrontFace(descriptor.Primitive.FrontFace)
	cDescriptor.primitive.cullMode = C.WGPUCullMode(descriptor.Primitive.CullMode)
	cDescriptor.primitive.unclippedDepth = boolToWGPUBool(descriptor.Primitive.UnclippedDepth)
	if descriptor.DepthStencil != nil {
		cDescriptor.depthStencil.format = C.WGPUTextureFormat(descriptor.DepthStencil.Format)
		cDescriptor.depthStencil.depthWriteEnabled = C.WGPUOptionalBool(descriptor.DepthStencil.DepthWriteEnabled)
		cDescriptor.depthStencil.depthCompare = C.WGPUCompareFunction(descriptor.DepthStencil.DepthCompare)
		cDescriptor.depthStencil.stencilFront.compare = C.WGPUCompareFunction(descriptor.DepthStencil.StencilFront.Compare)
		cDescriptor.depthStencil.stencilFront.failOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.FailOp)
		cDescriptor.depthStencil.stencilFront.depthFailOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.DepthFailOp)
		cDescriptor.depthStencil.stencilFront.passOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.PassOp)
		cDescriptor.depthStencil.stencilBack.compare = C.WGPUCompareFunction(descriptor.DepthStencil.StencilBack.Compare)
		cDescriptor.depthStencil.stencilBack.failOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.FailOp)
		cDescriptor.depthStencil.stencilBack.depthFailOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.DepthFailOp)
		cDescriptor.depthStencil.stencilBack.passOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.PassOp)
		cDescriptor.depthStencil.stencilReadMask = C.uint32_t(descriptor.DepthStencil.StencilReadMask)
		cDescriptor.depthStencil.stencilWriteMask = C.uint32_t(descriptor.DepthStencil.StencilWriteMask)
		cDescriptor.depthStencil.depthBias = C.int32_t(descriptor.DepthStencil.DepthBias)
		cDescriptor.depthStencil.depthBiasSlopeScale = C.float(descriptor.DepthStencil.DepthBiasSlopeScale)
		cDescriptor.depthStencil.depthBiasClamp = C.float(descriptor.DepthStencil.DepthBiasClamp)
	}
	cDescriptor.multisample.count = C.uint32_t(descriptor.Multisample.Count)
	cDescriptor.multisample.mask = C.uint32_t(descriptor.Multisample.Mask)
	cDescriptor.multisample.alphaToCoverageEnabled = boolToWGPUBool(descriptor.Multisample.AlphaToCoverageEnabled)
	if descriptor.Fragment != nil {
		cDescriptor.fragment.module = C.WGPUShaderModule(unsafe.Pointer(descriptor.Fragment.Module.ref))
		cDescriptorfragmententryPointStr := C.CString(descriptor.Fragment.EntryPoint)
		defer C.free(unsafe.Pointer(cDescriptorfragmententryPointStr))
		cDescriptor.fragment.entryPoint.data = cDescriptorfragmententryPointStr
		cDescriptor.fragment.entryPoint.length = C.size_t(len(descriptor.Fragment.EntryPoint))
		cDescriptor.fragment.constantCount = C.size_t(descriptor.Fragment.ConstantCount)
		cDescriptorfragmentconstantskeyStr := C.CString(descriptor.Fragment.Constants.Key)
		defer C.free(unsafe.Pointer(cDescriptorfragmentconstantskeyStr))
		cDescriptor.fragment.constants.key.data = cDescriptorfragmentconstantskeyStr
		cDescriptor.fragment.constants.key.length = C.size_t(len(descriptor.Fragment.Constants.Key))
		cDescriptor.fragment.constants.value = C.double(descriptor.Fragment.Constants.Value)
		cDescriptor.fragment.targetCount = C.size_t(descriptor.Fragment.TargetCount)
		cDescriptor.fragment.targets.format = C.WGPUTextureFormat(descriptor.Fragment.Targets.Format)
		if descriptor.Fragment.Targets.Blend != nil {
			cDescriptor.fragment.targets.blend.color.operation = C.WGPUBlendOperation(descriptor.Fragment.Targets.Blend.Color.Operation)
			cDescriptor.fragment.targets.blend.color.srcFactor = C.WGPUBlendFactor(descriptor.Fragment.Targets.Blend.Color.SrcFactor)
			cDescriptor.fragment.targets.blend.color.dstFactor = C.WGPUBlendFactor(descriptor.Fragment.Targets.Blend.Color.DstFactor)
			cDescriptor.fragment.targets.blend.alpha.operation = C.WGPUBlendOperation(descriptor.Fragment.Targets.Blend.Alpha.Operation)
			cDescriptor.fragment.targets.blend.alpha.srcFactor = C.WGPUBlendFactor(descriptor.Fragment.Targets.Blend.Alpha.SrcFactor)
			cDescriptor.fragment.targets.blend.alpha.dstFactor = C.WGPUBlendFactor(descriptor.Fragment.Targets.Blend.Alpha.DstFactor)
		}
		cDescriptor.fragment.targets.writeMask = C.WGPUColorWriteMask(descriptor.Fragment.Targets.WriteMask)
	}
	// Call and return
	return RenderPipeline{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateRenderPipeline(cDevice, &cDescriptor)))}
}

func (d *Device) CreateSampler(descriptor *SamplerDescriptor) Sampler {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var pDescriptor C.WGPUSamplerDescriptor
	if descriptor != nil {
		pDescriptorlabelStr := C.CString(descriptor.Label)
		defer C.free(unsafe.Pointer(pDescriptorlabelStr))
		pDescriptor.label.data = pDescriptorlabelStr
		pDescriptor.label.length = C.size_t(len(descriptor.Label))
		pDescriptor.addressModeU = C.WGPUAddressMode(descriptor.AddressModeU)
		pDescriptor.addressModeV = C.WGPUAddressMode(descriptor.AddressModeV)
		pDescriptor.addressModeW = C.WGPUAddressMode(descriptor.AddressModeW)
		pDescriptor.magFilter = C.WGPUFilterMode(descriptor.MagFilter)
		pDescriptor.minFilter = C.WGPUFilterMode(descriptor.MinFilter)
		pDescriptor.mipmapFilter = C.WGPUMipmapFilterMode(descriptor.MipmapFilter)
		pDescriptor.lodMinClamp = C.float(descriptor.LodMinClamp)
		pDescriptor.lodMaxClamp = C.float(descriptor.LodMaxClamp)
		pDescriptor.compare = C.WGPUCompareFunction(descriptor.Compare)
		pDescriptor.maxAnisotropy = C.uint16_t(descriptor.MaxAnisotropy)
	}
	// Call and return
	return Sampler{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateSampler(cDevice, &pDescriptor)))}
}

func (d *Device) CreateShaderModule(descriptor ShaderModuleDescriptor) ShaderModule {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUShaderModuleDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	// Call and return
	return ShaderModule{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateShaderModule(cDevice, &cDescriptor)))}
}

func (d *Device) CreateTexture(descriptor TextureDescriptor) Texture {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUTextureDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor.usage = C.WGPUTextureUsage(descriptor.Usage)
	cDescriptor.dimension = C.WGPUTextureDimension(descriptor.Dimension)
	cDescriptor.size.width = C.uint32_t(descriptor.Size.Width)
	cDescriptor.size.height = C.uint32_t(descriptor.Size.Height)
	cDescriptor.size.depthOrArrayLayers = C.uint32_t(descriptor.Size.DepthOrArrayLayers)
	cDescriptor.format = C.WGPUTextureFormat(descriptor.Format)
	cDescriptor.mipLevelCount = C.uint32_t(descriptor.MipLevelCount)
	cDescriptor.sampleCount = C.uint32_t(descriptor.SampleCount)

	viewFormatsCount := len(descriptor.ViewFormats)
	if viewFormatsCount > 0 {
		viewFormats := C.malloc(C.size_t(viewFormatsCount) * C.size_t(unsafe.Sizeof(C.WGPUTextureFormat(0))))
		defer C.free(viewFormats)

		slice := unsafe.Slice((*TextureFormat)(viewFormats), viewFormatsCount)
		copy(slice, descriptor.ViewFormats)

		cDescriptor.viewFormats = (*C.WGPUTextureFormat)(viewFormats)
		cDescriptor.viewFormatCount = C.size_t(viewFormatsCount)
	}

	return Texture{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateTexture(cDevice, &cDescriptor)))}
}

func (d *Device) Destroy() {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	// Call and return
	C.wgpuDeviceDestroy(cDevice)
}

func (d *Device) GetLimits(limits Limits) statusCode {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cLimits C.WGPULimits
	cLimits.maxTextureDimension1D = C.uint32_t(limits.MaxTextureDimension1D)
	cLimits.maxTextureDimension2D = C.uint32_t(limits.MaxTextureDimension2D)
	cLimits.maxTextureDimension3D = C.uint32_t(limits.MaxTextureDimension3D)
	cLimits.maxTextureArrayLayers = C.uint32_t(limits.MaxTextureArrayLayers)
	cLimits.maxBindGroups = C.uint32_t(limits.MaxBindGroups)
	cLimits.maxBindGroupsPlusVertexBuffers = C.uint32_t(limits.MaxBindGroupsPlusVertexBuffers)
	cLimits.maxBindingsPerBindGroup = C.uint32_t(limits.MaxBindingsPerBindGroup)
	cLimits.maxDynamicUniformBuffersPerPipelineLayout = C.uint32_t(limits.MaxDynamicUniformBuffersPerPipelineLayout)
	cLimits.maxDynamicStorageBuffersPerPipelineLayout = C.uint32_t(limits.MaxDynamicStorageBuffersPerPipelineLayout)
	cLimits.maxSampledTexturesPerShaderStage = C.uint32_t(limits.MaxSampledTexturesPerShaderStage)
	cLimits.maxSamplersPerShaderStage = C.uint32_t(limits.MaxSamplersPerShaderStage)
	cLimits.maxStorageBuffersPerShaderStage = C.uint32_t(limits.MaxStorageBuffersPerShaderStage)
	cLimits.maxStorageTexturesPerShaderStage = C.uint32_t(limits.MaxStorageTexturesPerShaderStage)
	cLimits.maxUniformBuffersPerShaderStage = C.uint32_t(limits.MaxUniformBuffersPerShaderStage)
	cLimits.maxUniformBufferBindingSize = C.uint64_t(limits.MaxUniformBufferBindingSize)
	cLimits.maxStorageBufferBindingSize = C.uint64_t(limits.MaxStorageBufferBindingSize)
	cLimits.minUniformBufferOffsetAlignment = C.uint32_t(limits.MinUniformBufferOffsetAlignment)
	cLimits.minStorageBufferOffsetAlignment = C.uint32_t(limits.MinStorageBufferOffsetAlignment)
	cLimits.maxVertexBuffers = C.uint32_t(limits.MaxVertexBuffers)
	cLimits.maxBufferSize = C.uint64_t(limits.MaxBufferSize)
	cLimits.maxVertexAttributes = C.uint32_t(limits.MaxVertexAttributes)
	cLimits.maxVertexBufferArrayStride = C.uint32_t(limits.MaxVertexBufferArrayStride)
	cLimits.maxInterStageShaderVariables = C.uint32_t(limits.MaxInterStageShaderVariables)
	cLimits.maxColorAttachments = C.uint32_t(limits.MaxColorAttachments)
	cLimits.maxColorAttachmentBytesPerSample = C.uint32_t(limits.MaxColorAttachmentBytesPerSample)
	cLimits.maxComputeWorkgroupStorageSize = C.uint32_t(limits.MaxComputeWorkgroupStorageSize)
	cLimits.maxComputeInvocationsPerWorkgroup = C.uint32_t(limits.MaxComputeInvocationsPerWorkgroup)
	cLimits.maxComputeWorkgroupSizeX = C.uint32_t(limits.MaxComputeWorkgroupSizeX)
	cLimits.maxComputeWorkgroupSizeY = C.uint32_t(limits.MaxComputeWorkgroupSizeY)
	cLimits.maxComputeWorkgroupSizeZ = C.uint32_t(limits.MaxComputeWorkgroupSizeZ)
	cLimits.maxComputeWorkgroupsPerDimension = C.uint32_t(limits.MaxComputeWorkgroupsPerDimension)
	cLimits.maxImmediateSize = C.uint32_t(limits.MaxImmediateSize)
	// Call and return
	return statusCode(C.wgpuDeviceGetLimits(cDevice, &cLimits))
}

func (d *Device) HasFeature(feature FeatureName) bool {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	cFeature := C.WGPUFeatureName(feature)
	// Call and return
	return bool(C.wgpuDeviceHasFeature(cDevice, cFeature) != 0)
}

func (d *Device) GetFeatures() []FeatureName {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cFeatures C.WGPUSupportedFeatures
	//TODO:here
	// cFeatures.featureCount = C.size_t(features.FeatureCount)
	// cFeatures.features = C.WGPUFeatureName(features.Features)
	// Call and return
	C.wgpuDeviceGetFeatures(cDevice, &cFeatures)

	return nil
}

func (d *Device) GetAdapterInfo(adapterInfo AdapterInfo) statusCode {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cAdapterInfo C.WGPUAdapterInfo
	cAdapterInfovendorStr := C.CString(adapterInfo.Vendor)
	defer C.free(unsafe.Pointer(cAdapterInfovendorStr))
	cAdapterInfo.vendor.data = cAdapterInfovendorStr
	cAdapterInfo.vendor.length = C.size_t(len(adapterInfo.Vendor))
	cAdapterInfoarchitectureStr := C.CString(adapterInfo.Architecture)
	defer C.free(unsafe.Pointer(cAdapterInfoarchitectureStr))
	cAdapterInfo.architecture.data = cAdapterInfoarchitectureStr
	cAdapterInfo.architecture.length = C.size_t(len(adapterInfo.Architecture))
	cAdapterInfodeviceStr := C.CString(adapterInfo.Device)
	defer C.free(unsafe.Pointer(cAdapterInfodeviceStr))
	cAdapterInfo.device.data = cAdapterInfodeviceStr
	cAdapterInfo.device.length = C.size_t(len(adapterInfo.Device))
	cAdapterInfodescriptionStr := C.CString(adapterInfo.Description)
	defer C.free(unsafe.Pointer(cAdapterInfodescriptionStr))
	cAdapterInfo.description.data = cAdapterInfodescriptionStr
	cAdapterInfo.description.length = C.size_t(len(adapterInfo.Description))
	cAdapterInfo.backendType = C.WGPUBackendType(adapterInfo.BackendType)
	cAdapterInfo.adapterType = C.WGPUAdapterType(adapterInfo.AdapterType)
	cAdapterInfo.vendorID = C.uint32_t(adapterInfo.VendorID)
	cAdapterInfo.deviceID = C.uint32_t(adapterInfo.DeviceID)
	cAdapterInfo.subgroupMinSize = C.uint32_t(adapterInfo.SubgroupMinSize)
	cAdapterInfo.subgroupMaxSize = C.uint32_t(adapterInfo.SubgroupMaxSize)
	// Call and return
	return statusCode(C.wgpuDeviceGetAdapterInfo(cDevice, &cAdapterInfo))
}

func (d *Device) GetQueue() Queue {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	// Call and return
	return Queue{ref: uintptr(unsafe.Pointer(C.wgpuDeviceGetQueue(cDevice)))}
}

func (d *Device) PushErrorScope(filter ErrorFilter) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	cFilter := C.WGPUErrorFilter(filter)
	// Call and return
	C.wgpuDevicePushErrorScope(cDevice, cFilter)
}

//export goPopErrorScopeCallbackHandler
func goPopErrorScopeCallbackHandler(status C.WGPUPopErrorScopeStatus, typ C.WGPUErrorType, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	defer handle.Delete()
	fn := handle.Value().(PopErrorScopeCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}
	fn(
		PopErrorScopeStatus(status),
		ErrorType(typ),
		msg,
	)
}
func (d *Device) PopErrorScope(callback PopErrorScopeCallback) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	handle := cgo.NewHandle(callback)
	var cCallbackInfo C.WGPUPopErrorScopeCallbackInfo
	cCallbackInfo.mode = C.WGPUCallbackMode(callbackModeAllowProcessEvents)
	cCallbackInfo.callback = C.WGPUPopErrorScopeCallback(C.cgo_callback_PopErrorScopeCallback)
	cCallbackInfo.userdata1 = unsafe.Pointer(handle)
	cCallbackInfo.userdata2 = nil

	C.wgpuDevicePopErrorScope(cDevice, cCallbackInfo)
}

func (d *Device) SetLabel(label string) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuDeviceSetLabel(cDevice, cLabel)
}

type ExternalTexture struct {
	ref uintptr
}

func (e *ExternalTexture) SetLabel(label string) {
	cExternalTexture := C.WGPUExternalTexture(unsafe.Pointer(e.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuExternalTextureSetLabel(cExternalTexture, cLabel)
}

type Instance struct {
	ref uintptr
}

func (i *Instance) CreateSurface(descriptor SurfaceDescriptor) Surface {
	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))

	var cDescriptor C.WGPUSurfaceDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	// Call and return
	return Surface{ref: uintptr(unsafe.Pointer(C.wgpuInstanceCreateSurface(cInstance, &cDescriptor)))}
}

func (i *Instance) ProcessEvents() {
	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))

	// Call and return
	C.wgpuInstanceProcessEvents(cInstance)
}

// func (i *Instance) WaitAny(futureCount int, futures *FutureWaitInfo, timeoutNS uint64) WaitStatus {
// 	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))

// 	cFutureCount := C.size_t(futureCount)
// 	if futures != nil {
// 		var pFutures C.WGPUFutureWaitInfo
// 		pFutures.future.id = C.uint64_t(futures.Future.Id)
// 		pFutures.completed = boolToWGPUBool(futures.Completed)
// 	}
// 	cTimeoutNS := C.uint64_t(timeoutNS)
// 	// Call and return
// 	return WaitStatus(C.wgpuInstanceWaitAny(cInstance, cFutureCount, pFutures, cTimeoutNS))
// }

//export goRequestAdapterCallbackHandler
func goRequestAdapterCallbackHandler(status C.WGPURequestAdapterStatus, adapter C.WGPUAdapter, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	defer handle.Delete()
	fn := handle.Value().(requestAdapterCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}
	fn(
		requestAdapterStatus(status),
		&Adapter{ref: uintptr(unsafe.Pointer(adapter))},
		msg,
	)
}

func (i *Instance) RequestAdapter(options *RequestAdapterOptions) *Adapter {
	adapter, err := i.TryRequestAdapter(options)
	if err != nil {
		panic(err)
	}
	return adapter
}

func (i *Instance) TryRequestAdapter(options *RequestAdapterOptions) (*Adapter, error) {
	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))

	var pOptions C.WGPURequestAdapterOptions
	if options != nil {
		pOptions.featureLevel = C.WGPUFeatureLevel(options.FeatureLevel)
		pOptions.powerPreference = C.WGPUPowerPreference(options.PowerPreference)
		pOptions.forceFallbackAdapter = boolToWGPUBool(options.ForceFallbackAdapter)
		pOptions.backendType = C.WGPUBackendType(options.BackendType)
		pOptions.compatibleSurface = C.WGPUSurface(unsafe.Pointer(options.CompatibleSurface.ref))
	}

	var status requestAdapterStatus
	var adapter *Adapter
	var message string

	callback := requestAdapterCallback(func(s requestAdapterStatus, a *Adapter, m string) {
		status = s
		adapter = a
		message = m
	})

	handle := cgo.NewHandle(callback)

	var cCallbackInfo C.WGPURequestAdapterCallbackInfo
	cCallbackInfo.mode = C.WGPUCallbackMode(callbackModeAllowSpontaneous)
	cCallbackInfo.callback = C.WGPURequestAdapterCallback(C.cgo_callback_RequestAdapterCallback)
	cCallbackInfo.userdata1 = unsafe.Pointer(handle)
	cCallbackInfo.userdata2 = nil

	C.wgpuInstanceRequestAdapter(cInstance, &pOptions, cCallbackInfo)

	if status != requestAdapterStatusSuccess {
		return nil, fmt.Errorf("error request adapter: %s", message)
	}

	return adapter, nil
}

func (i *Instance) HasWGSLLanguageFeature(feature WGSLLanguageFeatureName) bool {
	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))

	cFeature := C.WGPUWGSLLanguageFeatureName(feature)
	// Call and return
	return bool(C.wgpuInstanceHasWGSLLanguageFeature(cInstance, cFeature) != 0)
}

func (i *Instance) GetWGSLLanguageFeatures(features SupportedWGSLLanguageFeatures) {
	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))

	var cFeatures C.WGPUSupportedWGSLLanguageFeatures
	// cFeatures.featureCount = C.size_t(features.FeatureCount)
	// cFeatures.features = C.WGPUWGSLLanguageFeatureName(features.Features)
	// Call and return
	C.wgpuInstanceGetWGSLLanguageFeatures(cInstance, &cFeatures)
}

type PipelineLayout struct {
	ref uintptr
}

func (p *PipelineLayout) SetLabel(label string) {
	cPipelineLayout := C.WGPUPipelineLayout(unsafe.Pointer(p.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuPipelineLayoutSetLabel(cPipelineLayout, cLabel)
}

type QuerySet struct {
	ref uintptr
}

func (q *QuerySet) SetLabel(label string) {
	cQuerySet := C.WGPUQuerySet(unsafe.Pointer(q.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuQuerySetSetLabel(cQuerySet, cLabel)
}

func (q *QuerySet) GetType() QueryType {
	cQuerySet := C.WGPUQuerySet(unsafe.Pointer(q.ref))

	// Call and return
	return QueryType(C.wgpuQuerySetGetType(cQuerySet))
}

func (q *QuerySet) GetCount() uint32 {
	cQuerySet := C.WGPUQuerySet(unsafe.Pointer(q.ref))

	// Call and return
	return uint32(C.wgpuQuerySetGetCount(cQuerySet))
}

func (q *QuerySet) Destroy() {
	cQuerySet := C.WGPUQuerySet(unsafe.Pointer(q.ref))

	// Call and return
	C.wgpuQuerySetDestroy(cQuerySet)
}

type Queue struct {
	ref uintptr
}

func (q *Queue) Submit(commandCount int, commands *CommandBuffer) {
	cQueue := C.WGPUQueue(unsafe.Pointer(q.ref))

	cCommandCount := C.size_t(commandCount)
	pCommands := C.WGPUCommandBuffer(unsafe.Pointer(commands.ref))
	// Call and return
	C.wgpuQueueSubmit(cQueue, cCommandCount, &pCommands)
}

//export goQueueWorkDoneCallbackHandler
func goQueueWorkDoneCallbackHandler(status C.WGPUQueueWorkDoneStatus, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	defer handle.Delete()
	fn := handle.Value().(QueueWorkDoneCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}

	fn(
		QueueWorkDoneStatus(status),
		msg,
	)
}
func (q *Queue) OnSubmittedWorkDone(callback QueueWorkDoneCallback) {
	cQueue := C.WGPUQueue(unsafe.Pointer(q.ref))

	handle := cgo.NewHandle(callback)

	var cCallbackInfo C.WGPUQueueWorkDoneCallbackInfo
	cCallbackInfo.mode = C.WGPUCallbackMode(callbackModeAllowProcessEvents)
	cCallbackInfo.callback = C.WGPUQueueWorkDoneCallback(C.cgo_callback_QueueWorkDoneCallback)
	cCallbackInfo.userdata1 = unsafe.Pointer(handle)
	cCallbackInfo.userdata2 = nil

	C.wgpuQueueOnSubmittedWorkDone(cQueue, cCallbackInfo)
}

func (q *Queue) WriteBuffer(buffer *Buffer, bufferOffset uint64, data []byte, size int) {
	cQueue := C.WGPUQueue(unsafe.Pointer(q.ref))

	pBuffer := C.WGPUBuffer(unsafe.Pointer(buffer.ref))
	cBufferOffset := C.uint64_t(bufferOffset)
	var cData unsafe.Pointer
	if len(data) > 0 {
		cData = unsafe.Pointer(&data[0])
	}
	cSize := C.size_t(size)

	C.wgpuQueueWriteBuffer(cQueue, pBuffer, cBufferOffset, cData, cSize)
}

func (q *Queue) WriteTexture(destination TexelCopyTextureInfo, data []byte, dataSize int, dataLayout TexelCopyBufferLayout, writeSize Extent3D) {
	cQueue := C.WGPUQueue(unsafe.Pointer(q.ref))

	var cDestination C.WGPUTexelCopyTextureInfo
	cDestination.texture = C.WGPUTexture(unsafe.Pointer(destination.Texture.ref))
	cDestination.mipLevel = C.uint32_t(destination.MipLevel)
	cDestination.origin.x = C.uint32_t(destination.Origin.X)
	cDestination.origin.y = C.uint32_t(destination.Origin.Y)
	cDestination.origin.z = C.uint32_t(destination.Origin.Z)
	cDestination.aspect = C.WGPUTextureAspect(destination.Aspect)
	var cData unsafe.Pointer
	if len(data) > 0 {
		cData = unsafe.Pointer(&data[0])
	}
	cDataSize := C.size_t(dataSize)
	var cDataLayout C.WGPUTexelCopyBufferLayout
	cDataLayout.offset = C.uint64_t(dataLayout.Offset)
	cDataLayout.bytesPerRow = C.uint32_t(dataLayout.BytesPerRow)
	cDataLayout.rowsPerImage = C.uint32_t(dataLayout.RowsPerImage)
	var cWriteSize C.WGPUExtent3D
	cWriteSize.width = C.uint32_t(writeSize.Width)
	cWriteSize.height = C.uint32_t(writeSize.Height)
	cWriteSize.depthOrArrayLayers = C.uint32_t(writeSize.DepthOrArrayLayers)
	// Call and return
	C.wgpuQueueWriteTexture(cQueue, &cDestination, cData, cDataSize, &cDataLayout, &cWriteSize)
}

func (q *Queue) SetLabel(label string) {
	cQueue := C.WGPUQueue(unsafe.Pointer(q.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuQueueSetLabel(cQueue, cLabel)
}

type RenderBundle struct {
	ref uintptr
}

func (r *RenderBundle) SetLabel(label string) {
	cRenderBundle := C.WGPURenderBundle(unsafe.Pointer(r.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuRenderBundleSetLabel(cRenderBundle, cLabel)
}

type RenderBundleEncoder struct {
	ref uintptr
}

func (r *RenderBundleEncoder) SetPipeline(pipeline *RenderPipeline) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))

	pPipeline := C.WGPURenderPipeline(unsafe.Pointer(pipeline.ref))
	// Call and return
	C.wgpuRenderBundleEncoderSetPipeline(cRenderBundleEncoder, pPipeline)
}

func (r *RenderBundleEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsetCount int, dynamicOffsets uint32) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))

	cGroupIndex := C.uint32_t(groupIndex)
	pGroup := C.WGPUBindGroup(unsafe.Pointer(group.ref))
	cDynamicOffsetCount := C.size_t(dynamicOffsetCount)
	cDynamicOffsets := C.uint32_t(dynamicOffsets)
	// Call and return
	C.wgpuRenderBundleEncoderSetBindGroup(cRenderBundleEncoder, cGroupIndex, pGroup, cDynamicOffsetCount, &cDynamicOffsets)
}

func (r *RenderBundleEncoder) Draw(vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))

	cVertexCount := C.uint32_t(vertexCount)
	cInstanceCount := C.uint32_t(instanceCount)
	cFirstVertex := C.uint32_t(firstVertex)
	cFirstInstance := C.uint32_t(firstInstance)
	// Call and return
	C.wgpuRenderBundleEncoderDraw(cRenderBundleEncoder, cVertexCount, cInstanceCount, cFirstVertex, cFirstInstance)
}

func (r *RenderBundleEncoder) DrawIndexed(indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))

	cIndexCount := C.uint32_t(indexCount)
	cInstanceCount := C.uint32_t(instanceCount)
	cFirstIndex := C.uint32_t(firstIndex)
	cBaseVertex := C.int32_t(baseVertex)
	cFirstInstance := C.uint32_t(firstInstance)
	// Call and return
	C.wgpuRenderBundleEncoderDrawIndexed(cRenderBundleEncoder, cIndexCount, cInstanceCount, cFirstIndex, cBaseVertex, cFirstInstance)
}

func (r *RenderBundleEncoder) DrawIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))

	pIndirectBuffer := C.WGPUBuffer(unsafe.Pointer(indirectBuffer.ref))
	cIndirectOffset := C.uint64_t(indirectOffset)
	// Call and return
	C.wgpuRenderBundleEncoderDrawIndirect(cRenderBundleEncoder, pIndirectBuffer, cIndirectOffset)
}

func (r *RenderBundleEncoder) DrawIndexedIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))

	pIndirectBuffer := C.WGPUBuffer(unsafe.Pointer(indirectBuffer.ref))
	cIndirectOffset := C.uint64_t(indirectOffset)
	// Call and return
	C.wgpuRenderBundleEncoderDrawIndexedIndirect(cRenderBundleEncoder, pIndirectBuffer, cIndirectOffset)
}

func (r *RenderBundleEncoder) InsertDebugMarker(markerLabel string) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))

	cMarkerLabelStr := C.CString(markerLabel)
	defer C.free(unsafe.Pointer(cMarkerLabelStr))
	var cMarkerLabel C.WGPUStringView
	cMarkerLabel.data = cMarkerLabelStr
	cMarkerLabel.length = C.size_t(len(markerLabel))
	// Call and return
	C.wgpuRenderBundleEncoderInsertDebugMarker(cRenderBundleEncoder, cMarkerLabel)
}

func (r *RenderBundleEncoder) PopDebugGroup() {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))

	// Call and return
	C.wgpuRenderBundleEncoderPopDebugGroup(cRenderBundleEncoder)
}

func (r *RenderBundleEncoder) PushDebugGroup(groupLabel string) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))

	cGroupLabelStr := C.CString(groupLabel)
	defer C.free(unsafe.Pointer(cGroupLabelStr))
	var cGroupLabel C.WGPUStringView
	cGroupLabel.data = cGroupLabelStr
	cGroupLabel.length = C.size_t(len(groupLabel))
	// Call and return
	C.wgpuRenderBundleEncoderPushDebugGroup(cRenderBundleEncoder, cGroupLabel)
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
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))

	var pDescriptor C.WGPURenderBundleDescriptor
	if descriptor != nil {
		pDescriptorlabelStr := C.CString(descriptor.Label)
		defer C.free(unsafe.Pointer(pDescriptorlabelStr))
		pDescriptor.label.data = pDescriptorlabelStr
		pDescriptor.label.length = C.size_t(len(descriptor.Label))
	}
	// Call and return
	return RenderBundle{ref: uintptr(unsafe.Pointer(C.wgpuRenderBundleEncoderFinish(cRenderBundleEncoder, &pDescriptor)))}
}

func (r *RenderBundleEncoder) SetLabel(label string) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuRenderBundleEncoderSetLabel(cRenderBundleEncoder, cLabel)
}

type RenderPassEncoder struct {
	ref uintptr
}

func (r *RenderPassEncoder) SetPipeline(pipeline *RenderPipeline) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	pPipeline := C.WGPURenderPipeline(unsafe.Pointer(pipeline.ref))
	// Call and return
	C.wgpuRenderPassEncoderSetPipeline(cRenderPassEncoder, pPipeline)
}

func (r *RenderPassEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsetCount int, dynamicOffsets uint32) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	cGroupIndex := C.uint32_t(groupIndex)
	pGroup := C.WGPUBindGroup(unsafe.Pointer(group.ref))
	cDynamicOffsetCount := C.size_t(dynamicOffsetCount)
	cDynamicOffsets := C.uint32_t(dynamicOffsets)
	// Call and return
	C.wgpuRenderPassEncoderSetBindGroup(cRenderPassEncoder, cGroupIndex, pGroup, cDynamicOffsetCount, &cDynamicOffsets)
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
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	pIndirectBuffer := C.WGPUBuffer(unsafe.Pointer(indirectBuffer.ref))
	cIndirectOffset := C.uint64_t(indirectOffset)
	// Call and return
	C.wgpuRenderPassEncoderDrawIndexedIndirect(cRenderPassEncoder, pIndirectBuffer, cIndirectOffset)
}

func (r *RenderPassEncoder) ExecuteBundles(bundleCount int, bundles *RenderBundle) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	cBundleCount := C.size_t(bundleCount)
	pBundles := C.WGPURenderBundle(unsafe.Pointer(bundles.ref))
	// Call and return
	C.wgpuRenderPassEncoderExecuteBundles(cRenderPassEncoder, cBundleCount, &pBundles)
}

func (r *RenderPassEncoder) InsertDebugMarker(markerLabel string) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	cMarkerLabelStr := C.CString(markerLabel)
	defer C.free(unsafe.Pointer(cMarkerLabelStr))
	var cMarkerLabel C.WGPUStringView
	cMarkerLabel.data = cMarkerLabelStr
	cMarkerLabel.length = C.size_t(len(markerLabel))
	// Call and return
	C.wgpuRenderPassEncoderInsertDebugMarker(cRenderPassEncoder, cMarkerLabel)
}

func (r *RenderPassEncoder) PopDebugGroup() {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	// Call and return
	C.wgpuRenderPassEncoderPopDebugGroup(cRenderPassEncoder)
}

func (r *RenderPassEncoder) PushDebugGroup(groupLabel string) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	cGroupLabelStr := C.CString(groupLabel)
	defer C.free(unsafe.Pointer(cGroupLabelStr))
	var cGroupLabel C.WGPUStringView
	cGroupLabel.data = cGroupLabelStr
	cGroupLabel.length = C.size_t(len(groupLabel))
	// Call and return
	C.wgpuRenderPassEncoderPushDebugGroup(cRenderPassEncoder, cGroupLabel)
}

func (r *RenderPassEncoder) SetStencilReference(reference uint32) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	cReference := C.uint32_t(reference)
	// Call and return
	C.wgpuRenderPassEncoderSetStencilReference(cRenderPassEncoder, cReference)
}

func (r *RenderPassEncoder) SetBlendConstant(color Color) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	var cColor C.WGPUColor
	cColor.r = C.double(color.R)
	cColor.g = C.double(color.G)
	cColor.b = C.double(color.B)
	cColor.a = C.double(color.A)
	// Call and return
	C.wgpuRenderPassEncoderSetBlendConstant(cRenderPassEncoder, &cColor)
}

func (r *RenderPassEncoder) SetViewport(x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	cX := C.float(x)
	cY := C.float(y)
	cWidth := C.float(width)
	cHeight := C.float(height)
	cMinDepth := C.float(minDepth)
	cMaxDepth := C.float(maxDepth)
	// Call and return
	C.wgpuRenderPassEncoderSetViewport(cRenderPassEncoder, cX, cY, cWidth, cHeight, cMinDepth, cMaxDepth)
}

func (r *RenderPassEncoder) SetScissorRect(x uint32, y uint32, width uint32, height uint32) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	cX := C.uint32_t(x)
	cY := C.uint32_t(y)
	cWidth := C.uint32_t(width)
	cHeight := C.uint32_t(height)
	// Call and return
	C.wgpuRenderPassEncoderSetScissorRect(cRenderPassEncoder, cX, cY, cWidth, cHeight)
}

func (r *RenderPassEncoder) SetVertexBuffer(slot uint32, buffer *Buffer, offset uint64, size uint64) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	cSlot := C.uint32_t(slot)
	pBuffer := C.WGPUBuffer(unsafe.Pointer(buffer.ref))
	cOffset := C.uint64_t(offset)
	cSize := C.uint64_t(size)
	// Call and return
	C.wgpuRenderPassEncoderSetVertexBuffer(cRenderPassEncoder, cSlot, pBuffer, cOffset, cSize)
}

func (r *RenderPassEncoder) SetIndexBuffer(buffer *Buffer, format IndexFormat, offset uint64, size uint64) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	pBuffer := C.WGPUBuffer(unsafe.Pointer(buffer.ref))
	cFormat := C.WGPUIndexFormat(format)
	cOffset := C.uint64_t(offset)
	cSize := C.uint64_t(size)
	// Call and return
	C.wgpuRenderPassEncoderSetIndexBuffer(cRenderPassEncoder, pBuffer, cFormat, cOffset, cSize)
}

func (r *RenderPassEncoder) BeginOcclusionQuery(queryIndex uint32) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	cQueryIndex := C.uint32_t(queryIndex)
	// Call and return
	C.wgpuRenderPassEncoderBeginOcclusionQuery(cRenderPassEncoder, cQueryIndex)
}

func (r *RenderPassEncoder) EndOcclusionQuery() {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	// Call and return
	C.wgpuRenderPassEncoderEndOcclusionQuery(cRenderPassEncoder)
}

func (r *RenderPassEncoder) End() {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	// Call and return
	C.wgpuRenderPassEncoderEnd(cRenderPassEncoder)
}

func (r *RenderPassEncoder) SetLabel(label string) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuRenderPassEncoderSetLabel(cRenderPassEncoder, cLabel)
}

type RenderPipeline struct {
	ref uintptr
}

func (r *RenderPipeline) GetBindGroupLayout(groupIndex uint32) BindGroupLayout {
	cRenderPipeline := C.WGPURenderPipeline(unsafe.Pointer(r.ref))

	cGroupIndex := C.uint32_t(groupIndex)
	// Call and return
	return BindGroupLayout{ref: uintptr(unsafe.Pointer(C.wgpuRenderPipelineGetBindGroupLayout(cRenderPipeline, cGroupIndex)))}
}

func (r *RenderPipeline) SetLabel(label string) {
	cRenderPipeline := C.WGPURenderPipeline(unsafe.Pointer(r.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuRenderPipelineSetLabel(cRenderPipeline, cLabel)
}

type Sampler struct {
	ref uintptr
}

func (s *Sampler) SetLabel(label string) {
	cSampler := C.WGPUSampler(unsafe.Pointer(s.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuSamplerSetLabel(cSampler, cLabel)
}

type ShaderModule struct {
	ref uintptr
}

//export goCompilationInfoCallbackHandler
func goCompilationInfoCallbackHandler(status C.WGPUCompilationInfoRequestStatus, compilationInfo C.WGPUCompilationInfo, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	defer handle.Delete()
	fn := handle.Value().(compilationInfoCallback)

	var compInfo CompilationInfo
	compInfo.Messages = unsafe.Slice((*CompilationMessage)(unsafe.Pointer(compilationInfo.messages)), compilationInfo.messageCount)

	fn(
		compilationInfoRequestStatus(status),
		compInfo,
	)
}

func (s *ShaderModule) GetCompilationInfo() CompilationInfo {
	info, err := s.TryGetCompilationInfo()
	if err != nil {
		panic(err)
	}
	return info
}

func (s *ShaderModule) TryGetCompilationInfo() (CompilationInfo, error) {
	cShaderModule := C.WGPUShaderModule(unsafe.Pointer(s.ref))

	var status compilationInfoRequestStatus
	var info CompilationInfo

	callback := func(s compilationInfoRequestStatus, i CompilationInfo) {
		status = s
		info = i
	}

	handle := cgo.NewHandle(callback)

	var cCallbackInfo C.WGPUCompilationInfoCallbackInfo
	cCallbackInfo.mode = C.WGPUCallbackMode(callbackModeAllowSpontaneous)
	cCallbackInfo.callback = C.WGPUCompilationInfoCallback(C.cgo_callback_CompilationInfoCallback)
	cCallbackInfo.userdata1 = unsafe.Pointer(handle)
	cCallbackInfo.userdata2 = nil

	C.wgpuShaderModuleGetCompilationInfo(cShaderModule, cCallbackInfo)

	if status != compilationInfoRequestStatusSuccess {
		return CompilationInfo{}, fmt.Errorf("error getting compilation info")
	}

	return info, nil
}

func (s *ShaderModule) SetLabel(label string) {
	cShaderModule := C.WGPUShaderModule(unsafe.Pointer(s.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuShaderModuleSetLabel(cShaderModule, cLabel)
}

type Surface struct {
	ref uintptr
}

func (s *Surface) Configure(config SurfaceConfiguration) {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))

	var cConfig C.WGPUSurfaceConfiguration
	cConfig.device = C.WGPUDevice(unsafe.Pointer(config.Device.ref))
	cConfig.format = C.WGPUTextureFormat(config.Format)
	cConfig.usage = C.WGPUTextureUsage(config.Usage)
	cConfig.width = C.uint32_t(config.Width)
	cConfig.height = C.uint32_t(config.Height)
	cConfig.alphaMode = C.WGPUCompositeAlphaMode(config.AlphaMode)
	cConfig.presentMode = C.WGPUPresentMode(config.PresentMode)

	viewFormatsCount := len(config.ViewFormats)
	if viewFormatsCount > 0 {
		viewFormats := C.malloc(C.size_t(viewFormatsCount) * C.size_t(unsafe.Sizeof(C.WGPUTextureFormat(0))))
		defer C.free(viewFormats)

		slice := unsafe.Slice((*TextureFormat)(viewFormats), viewFormatsCount)
		copy(slice, config.ViewFormats)

		cConfig.viewFormats = (*C.WGPUTextureFormat)(viewFormats)
		cConfig.viewFormatCount = C.size_t(viewFormatsCount)
	}

	C.wgpuSurfaceConfigure(cSurface, &cConfig)
}

func (s *Surface) GetCapabilities(adapter *Adapter) SurfaceCapabilities {
	capabilities, err := s.TryGetCapabilities(adapter)
	if err != nil {
		panic(err)
	}
	return capabilities
}

func (s *Surface) TryGetCapabilities(adapter *Adapter) (SurfaceCapabilities, error) {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))
	pAdapter := C.WGPUAdapter(unsafe.Pointer(adapter.ref))

	var cCapabilities C.WGPUSurfaceCapabilities
	status := statusCode(C.wgpuSurfaceGetCapabilities(cSurface, pAdapter, &cCapabilities))

	if status != statusCodeSuccess {
		return SurfaceCapabilities{}, fmt.Errorf("error getting surface capabilities")
	}

	var capabilities SurfaceCapabilities
	capabilities.Usages = TextureUsage(cCapabilities.usages)
	capabilities.Formats = unsafe.Slice((*TextureFormat)(unsafe.Pointer(cCapabilities.formats)), cCapabilities.formatCount)
	capabilities.PresentModes = unsafe.Slice((*PresentMode)(unsafe.Pointer(cCapabilities.presentModes)), cCapabilities.presentModeCount)
	capabilities.AlphaModes = unsafe.Slice((*CompositeAlphaMode)(unsafe.Pointer(cCapabilities.alphaModes)), cCapabilities.alphaModeCount)

	return capabilities, nil
}

func (s *Surface) GetCurrentTexture(surfaceTexture SurfaceTexture) {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))

	var cSurfaceTexture C.WGPUSurfaceTexture
	cSurfaceTexture.texture = C.WGPUTexture(unsafe.Pointer(surfaceTexture.Texture.ref))
	cSurfaceTexture.status = C.WGPUSurfaceGetCurrentTextureStatus(surfaceTexture.Status)
	// Call and return
	C.wgpuSurfaceGetCurrentTexture(cSurface, &cSurfaceTexture)
}

func (s *Surface) Present() {
	err := s.TryPresent()
	if err != nil {
		panic(err)
	}
}

func (s *Surface) TryPresent() error {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))

	status := statusCode(C.wgpuSurfacePresent(cSurface))
	if status != statusCodeSuccess {
		return fmt.Errorf("error presenting")
	}

	return nil
}

func (s *Surface) Unconfigure() {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))

	// Call and return
	C.wgpuSurfaceUnconfigure(cSurface)
}

func (s *Surface) SetLabel(label string) {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuSurfaceSetLabel(cSurface, cLabel)
}

type Texture struct {
	ref uintptr
}

func (t *Texture) CreateView(descriptor *TextureViewDescriptor) TextureView {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))

	var pDescriptor C.WGPUTextureViewDescriptor
	if descriptor != nil {
		pDescriptorlabelStr := C.CString(descriptor.Label)
		defer C.free(unsafe.Pointer(pDescriptorlabelStr))
		pDescriptor.label.data = pDescriptorlabelStr
		pDescriptor.label.length = C.size_t(len(descriptor.Label))
		pDescriptor.format = C.WGPUTextureFormat(descriptor.Format)
		pDescriptor.dimension = C.WGPUTextureViewDimension(descriptor.Dimension)
		pDescriptor.baseMipLevel = C.uint32_t(descriptor.BaseMipLevel)
		pDescriptor.mipLevelCount = C.uint32_t(descriptor.MipLevelCount)
		pDescriptor.baseArrayLayer = C.uint32_t(descriptor.BaseArrayLayer)
		pDescriptor.arrayLayerCount = C.uint32_t(descriptor.ArrayLayerCount)
		pDescriptor.aspect = C.WGPUTextureAspect(descriptor.Aspect)
		pDescriptor.usage = C.WGPUTextureUsage(descriptor.Usage)
	}
	// Call and return
	return TextureView{ref: uintptr(unsafe.Pointer(C.wgpuTextureCreateView(cTexture, &pDescriptor)))}
}

func (t *Texture) SetLabel(label string) {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuTextureSetLabel(cTexture, cLabel)
}

func (t *Texture) GetWidth() uint32 {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))

	// Call and return
	return uint32(C.wgpuTextureGetWidth(cTexture))
}

func (t *Texture) GetHeight() uint32 {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))

	// Call and return
	return uint32(C.wgpuTextureGetHeight(cTexture))
}

func (t *Texture) GetDepthOrArrayLayers() uint32 {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))

	// Call and return
	return uint32(C.wgpuTextureGetDepthOrArrayLayers(cTexture))
}

func (t *Texture) GetMipLevelCount() uint32 {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))

	// Call and return
	return uint32(C.wgpuTextureGetMipLevelCount(cTexture))
}

func (t *Texture) GetSampleCount() uint32 {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))

	// Call and return
	return uint32(C.wgpuTextureGetSampleCount(cTexture))
}

func (t *Texture) GetDimension() TextureDimension {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))

	// Call and return
	return TextureDimension(C.wgpuTextureGetDimension(cTexture))
}

func (t *Texture) GetFormat() TextureFormat {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))

	// Call and return
	return TextureFormat(C.wgpuTextureGetFormat(cTexture))
}

func (t *Texture) GetUsage() TextureUsage {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))

	// Call and return
	return TextureUsage(C.wgpuTextureGetUsage(cTexture))
}

func (t *Texture) GetTextureBindingViewDimension() TextureViewDimension {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))

	// Call and return
	return TextureViewDimension(C.wgpuTextureGetTextureBindingViewDimension(cTexture))
}

func (t *Texture) Destroy() {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))

	// Call and return
	C.wgpuTextureDestroy(cTexture)
}

type TextureView struct {
	ref uintptr
}

func (t *TextureView) SetLabel(label string) {
	cTextureView := C.WGPUTextureView(unsafe.Pointer(t.ref))

	cLabelStr := C.CString(label)
	defer C.free(unsafe.Pointer(cLabelStr))
	var cLabel C.WGPUStringView
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	// Call and return
	C.wgpuTextureViewSetLabel(cTextureView, cLabel)
}

func CreateInstance(descriptor *InstanceDescriptor) Instance {
	var pDescriptor C.WGPUInstanceDescriptor
	if descriptor != nil {

		requiredFeatureCount := len(descriptor.RequiredFeatures)
		if requiredFeatureCount > 0 {
			requiredFeatures := C.malloc(C.size_t(requiredFeatureCount) * C.size_t(unsafe.Sizeof(C.WGPUInstanceFeatureName(0))))
			defer C.free(requiredFeatures)

			slice := unsafe.Slice((*InstanceFeatureName)(requiredFeatures), requiredFeatureCount)
			copy(slice, descriptor.RequiredFeatures)

			pDescriptor.requiredFeatures = (*C.WGPUInstanceFeatureName)(requiredFeatures)
			pDescriptor.requiredFeatureCount = C.size_t(requiredFeatureCount)
		}

		if descriptor.RequiredLimits != nil {
			pDescriptor.requiredLimits.timedWaitAnyMaxCount = C.size_t(descriptor.RequiredLimits.TimedWaitAnyMaxCount)
		}
	}

	return Instance{ref: uintptr(unsafe.Pointer(C.wgpuCreateInstance(&pDescriptor)))}
}

func GetInstanceFeatures() []InstanceFeatureName {
	var cFeatures C.WGPUSupportedInstanceFeatures
	C.wgpuGetInstanceFeatures(&cFeatures)

	return unsafe.Slice((*InstanceFeatureName)(unsafe.Pointer(cFeatures.features)), cFeatures.featureCount)
}

func GetInstanceLimits() InstanceLimits {
	limits, err := TryGetInstanceLimits()
	if err != nil {
		panic(err)
	}
	return limits
}

func TryGetInstanceLimits() (InstanceLimits, error) {
	var cLimits C.WGPUInstanceLimits

	status := statusCode(C.wgpuGetInstanceLimits(&cLimits))

	if status != statusCodeSuccess {
		return InstanceLimits{}, fmt.Errorf("error getting instance limits")
	}

	var limits InstanceLimits
	limits.TimedWaitAnyMaxCount = int(cLimits.timedWaitAnyMaxCount)

	return limits, nil
}

// func GetProcAddress(procName string) uintptr {
// 	cProcNameStr := C.CString(procName)
// 	defer C.free(unsafe.Pointer(cProcNameStr))
// 	var cProcName C.WGPUStringView
// 	cProcName.data = cProcNameStr
// 	cProcName.length = C.size_t(len(procName))
// 	// Call and return
// 	return uintptr(C.wgpuGetProcAddress(cProcName))
// }

func HasInstanceFeature(feature InstanceFeatureName) bool {
	cFeature := C.WGPUInstanceFeatureName(feature)
	return bool(C.wgpuHasInstanceFeature(cFeature) != 0)
}
