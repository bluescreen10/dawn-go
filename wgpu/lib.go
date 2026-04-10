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
	ref C.WGPUBindGroup
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

func (b *BindGroup) Release() {
	cBindGroup := C.WGPUBindGroup(unsafe.Pointer(b.ref))
	C.wgpuBindGroupRelease(cBindGroup)
}

type BindGroupLayout struct {
	ref C.WGPUBindGroupLayout
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

func (b *BindGroupLayout) Release() {
	cBindGroupLayout := C.WGPUBindGroupLayout(unsafe.Pointer(b.ref))
	C.wgpuBindGroupLayoutRelease(cBindGroupLayout)
}

type Buffer struct {
	ref C.WGPUBuffer
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

func (c *CommandBuffer) Release() {
	cCommandBuffer := C.WGPUCommandBuffer(unsafe.Pointer(c.ref))
	C.wgpuCommandBufferRelease(cCommandBuffer)
}

type CommandEncoder struct {
	ref C.WGPUCommandEncoder
}

func (c *CommandEncoder) Finish(descriptor *CommandBufferDescriptor) *CommandBuffer {
	commandBuffer, err := c.TryFinish(descriptor)
	if err != nil {
		panic(err)
	}
	return commandBuffer
}

func (c *CommandEncoder) TryFinish(descriptor *CommandBufferDescriptor) (*CommandBuffer, error) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	var pDescriptor C.WGPUCommandBufferDescriptor
	if descriptor != nil && descriptor.Label != "" {
		pDescriptor.label.data = C.CString(descriptor.Label)
		pDescriptor.label.length = C.size_t(len(descriptor.Label))
		defer C.free(unsafe.Pointer(pDescriptor.label.data))
	}

	ptr := unsafe.Pointer(C.wgpuCommandEncoderFinish(cCommandEncoder, &pDescriptor))
	if ptr == nil {
		return nil, fmt.Errorf("error finishing command encoder")
	}

	return &CommandBuffer{ref: uintptr(ptr)}, nil
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

func (c *CommandEncoder) BeginRenderPass(descriptor RenderPassDescriptor) *RenderPassEncoder {
	renderPassEncoder, err := c.TryBeginRenderPass(descriptor)
	if err != nil {
		panic(err)
	}
	return renderPassEncoder
}

func (c *CommandEncoder) TryBeginRenderPass(descriptor RenderPassDescriptor) (*RenderPassEncoder, error) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))

	var cDescriptor C.WGPURenderPassDescriptor

	if descriptor.Label != "" {
		cDescriptor.label.data = C.CString(descriptor.Label)
		cDescriptor.label.length = C.size_t(len(descriptor.Label))
		defer C.free(unsafe.Pointer(cDescriptor.label.data))
	}

	colorAttachmentsCount := len(descriptor.ColorAttachments)
	if colorAttachmentsCount > 0 {
		colorAttachments := C.malloc(C.size_t(colorAttachmentsCount) * C.size_t(unsafe.Sizeof(C.WGPURenderPassColorAttachment{})))
		slice := unsafe.Slice((*C.WGPURenderPassColorAttachment)(colorAttachments), colorAttachmentsCount)
		defer C.free(unsafe.Pointer(colorAttachments))

		for i, a := range descriptor.ColorAttachments {
			slice[i].depthSlice = C.uint32_t(a.DepthSlice)
			slice[i].loadOp = C.WGPULoadOp(a.LoadOp)
			slice[i].storeOp = C.WGPUStoreOp(a.StoreOp)
			slice[i].clearValue.r = C.double(a.ClearValue.R)
			slice[i].clearValue.g = C.double(a.ClearValue.G)
			slice[i].clearValue.b = C.double(a.ClearValue.B)
			slice[i].clearValue.a = C.double(a.ClearValue.A)

			if a.View != nil {
				slice[i].view = C.WGPUTextureView(unsafe.Pointer(a.View.ref))
			}

			if a.ResolveTarget != nil {
				slice[i].resolveTarget = C.WGPUTextureView(unsafe.Pointer(a.ResolveTarget.ref))
			}
		}

		cDescriptor.colorAttachments = (*C.WGPURenderPassColorAttachment)(colorAttachments)
		cDescriptor.colorAttachmentCount = C.size_t(colorAttachmentsCount)
	}

	if descriptor.DepthStencilAttachment != nil {
		if descriptor.DepthStencilAttachment.View != nil {
			cDescriptor.depthStencilAttachment.view = C.WGPUTextureView(unsafe.Pointer(descriptor.DepthStencilAttachment.View.ref))
		}
		cDescriptor.depthStencilAttachment.depthLoadOp = C.WGPULoadOp(descriptor.DepthStencilAttachment.DepthLoadOp)
		cDescriptor.depthStencilAttachment.depthStoreOp = C.WGPUStoreOp(descriptor.DepthStencilAttachment.DepthStoreOp)
		cDescriptor.depthStencilAttachment.depthClearValue = C.float(descriptor.DepthStencilAttachment.DepthClearValue)
		cDescriptor.depthStencilAttachment.depthReadOnly = toCBool(descriptor.DepthStencilAttachment.DepthReadOnly)
		cDescriptor.depthStencilAttachment.stencilLoadOp = C.WGPULoadOp(descriptor.DepthStencilAttachment.StencilLoadOp)
		cDescriptor.depthStencilAttachment.stencilStoreOp = C.WGPUStoreOp(descriptor.DepthStencilAttachment.StencilStoreOp)
		cDescriptor.depthStencilAttachment.stencilClearValue = C.uint32_t(descriptor.DepthStencilAttachment.StencilClearValue)
		cDescriptor.depthStencilAttachment.stencilReadOnly = toCBool(descriptor.DepthStencilAttachment.StencilReadOnly)
	}

	if descriptor.OcclusionQuerySet != nil {
		cDescriptor.occlusionQuerySet = C.WGPUQuerySet(unsafe.Pointer(descriptor.OcclusionQuerySet.ref))
	}

	if descriptor.TimestampWrites != nil {
		cDescriptor.timestampWrites.querySet = C.WGPUQuerySet(unsafe.Pointer(descriptor.TimestampWrites.QuerySet.ref))
		cDescriptor.timestampWrites.beginningOfPassWriteIndex = C.uint32_t(descriptor.TimestampWrites.BeginningOfPassWriteIndex)
		cDescriptor.timestampWrites.endOfPassWriteIndex = C.uint32_t(descriptor.TimestampWrites.EndOfPassWriteIndex)
	}

	ptr := unsafe.Pointer(C.wgpuCommandEncoderBeginRenderPass(cCommandEncoder, &cDescriptor))
	if ptr == nil {
		return nil, fmt.Errorf("error beginning render pass")
	}

	return &RenderPassEncoder{ref: uintptr(ptr)}, nil
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

func (c *CommandEncoder) Release() {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	C.wgpuCommandEncoderRelease(cCommandEncoder)
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
	ref C.WGPUComputePipeline
}

func (c *ComputePipeline) GetBindGroupLayout(groupIndex uint32) BindGroupLayout {
	cGroupIndex := C.uint32_t(groupIndex)
	return BindGroupLayout{ref: C.wgpuComputePipelineGetBindGroupLayout(c.ref, cGroupIndex)}
}

func (c *ComputePipeline) SetLabel(label string) {
	cLabel := toCStr(label)
	C.wgpuComputePipelineSetLabel(c.ref, cLabel)
}

type ExternalTexture struct {
	ref C.WGPUExternalTexture
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
	ref C.WGPUInstance
}

func (i *Instance) CreateSurface(descriptor SurfaceDescriptor) *Surface {
	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))

	var cDescriptor C.WGPUSurfaceDescriptor

	if descriptor.Label != "" {
		cDescriptor.label.data = C.CString(descriptor.Label)
		cDescriptor.label.length = C.size_t(len(descriptor.Label))
		defer C.free(unsafe.Pointer(cDescriptor.label.data))
	}

	if descriptor.MetalLayer != nil {
		metalSourcePtr := C.malloc(C.size_t(unsafe.Sizeof(C.WGPUSurfaceSourceMetalLayer{})))
		metalSource := (*C.WGPUSurfaceSourceMetalLayer)(metalSourcePtr)
		metalSource.chain.next = nil
		metalSource.chain.sType = C.WGPUSType_SurfaceSourceMetalLayer
		metalSource.layer = descriptor.MetalLayer.Layer
		cDescriptor.nextInChain = (*C.WGPUChainedStruct)(metalSourcePtr)
		defer C.free(metalSourcePtr)
	}

	ptr := unsafe.Pointer(C.wgpuInstanceCreateSurface(cInstance, &cDescriptor))
	return &Surface{ref: uintptr(ptr)}
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
// 		pFutures.completed = toCBool(futures.Completed)
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
		&Adapter{ref: adapter},
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

	var cOptions *C.WGPURequestAdapterOptions
	if options != nil {
		cOptions = &C.WGPURequestAdapterOptions{
			featureLevel:         C.WGPUFeatureLevel(options.FeatureLevel),
			powerPreference:      C.WGPUPowerPreference(options.PowerPreference),
			forceFallbackAdapter: toCBool(options.ForceFallbackAdapter),
			backendType:          C.WGPUBackendType(options.BackendType),
			compatibleSurface:    C.WGPUSurface(unsafe.Pointer(options.CompatibleSurface.ref)),
		}
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

	cCallbackInfo := C.WGPURequestAdapterCallbackInfo{
		mode:      C.WGPUCallbackMode(callbackModeAllowSpontaneous),
		callback:  C.WGPURequestAdapterCallback(C.cgo_callback_RequestAdapterCallback),
		userdata1: unsafe.Pointer(handle),
	}

	C.wgpuInstanceRequestAdapter(i.ref, cOptions, cCallbackInfo)

	if status != requestAdapterStatusSuccess {
		return nil, fmt.Errorf("error request adapter: %s", message)
	}

	return adapter, nil
}

func (i *Instance) HasWGSLLanguageFeature(feature WGSLLanguageFeatureName) bool {
	cFeature := C.WGPUWGSLLanguageFeatureName(feature)
	return bool(C.wgpuInstanceHasWGSLLanguageFeature(i.ref, cFeature) != 0)
}

func (i *Instance) GetWGSLLanguageFeatures() []WGSLLanguageFeatureName {
	var cFeatures C.WGPUSupportedWGSLLanguageFeatures
	C.wgpuInstanceGetWGSLLanguageFeatures(i.ref, &cFeatures)
	defer C.wgpuSupportedWGSLLanguageFeaturesFreeMembers(cFeatures)

	count := cFeatures.featureCount
	if count == 0 {
		return nil
	}

	features := make([]WGSLLanguageFeatureName, count)
	slice := unsafe.Slice((*C.WGPUWGSLLanguageFeatureName)(cFeatures.features), count)

	for i, f := range slice {
		features[i] = WGSLLanguageFeatureName(f)
	}

	return features
}

func (i *Instance) Release() {
	C.wgpuInstanceRelease(i.ref)
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

type Queue struct {
	ref C.WGPUQueue
}

func (q *Queue) Submit(commands ...*CommandBuffer) {
	commandsCount := len(commands)

	cCommandsCount := C.size_t(commandsCount)
	cCommands := make([]C.WGPUCommandBuffer, commandsCount)

	for i, c := range commands {
		cCommands[i] = C.WGPUCommandBuffer(unsafe.Pointer(c.ref))
	}

	C.wgpuQueueSubmit(q.ref, cCommandsCount, &cCommands[0])
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
	handle := cgo.NewHandle(callback)

	cCallbackInfo := C.WGPUQueueWorkDoneCallbackInfo{
		mode:      C.WGPUCallbackMode(callbackModeAllowProcessEvents),
		callback:  C.WGPUQueueWorkDoneCallback(C.cgo_callback_QueueWorkDoneCallback),
		userdata1: unsafe.Pointer(handle),
	}

	C.wgpuQueueOnSubmittedWorkDone(q.ref, cCallbackInfo)
}

func (q *Queue) WriteBuffer(buffer *Buffer, offset uint64, data []byte) {

	cBufferOffset := C.uint64_t(offset)
	cSize := C.size_t(len(data))

	var cData unsafe.Pointer
	if cSize > 0 {
		cData = unsafe.Pointer(&data[0])
	}

	C.wgpuQueueWriteBuffer(q.ref, buffer.ref, cBufferOffset, cData, cSize)
}

func (q *Queue) WriteTexture(destination TexelCopyTextureInfo, data []byte, dataLayout TexelCopyBufferLayout, writeSize Extent3D) {
	cDestination := C.WGPUTexelCopyTextureInfo{
		texture:  C.WGPUTexture(unsafe.Pointer(destination.Texture.ref)),
		mipLevel: C.uint32_t(destination.MipLevel),
		origin: C.WGPUOrigin3D{
			x: C.uint32_t(destination.Origin.X),
			y: C.uint32_t(destination.Origin.Y),
			z: C.uint32_t(destination.Origin.Z),
		},
		aspect: C.WGPUTextureAspect(destination.Aspect),
	}

	var cData unsafe.Pointer
	cDataSize := C.size_t(len(data))

	if cDataSize > 0 {
		cData = unsafe.Pointer(&data[0])
	}

	cDataLayout := C.WGPUTexelCopyBufferLayout{
		offset:       C.uint64_t(dataLayout.Offset),
		bytesPerRow:  C.uint32_t(dataLayout.BytesPerRow),
		rowsPerImage: C.uint32_t(dataLayout.RowsPerImage),
	}

	cWriteSize := C.WGPUExtent3D{
		width:              C.uint32_t(writeSize.Width),
		height:             C.uint32_t(writeSize.Height),
		depthOrArrayLayers: C.uint32_t(writeSize.DepthOrArrayLayers),
	}

	C.wgpuQueueWriteTexture(q.ref, &cDestination, cData, cDataSize, &cDataLayout, &cWriteSize)
}

func (q *Queue) SetLabel(label string) {
	cQueue := C.WGPUQueue(unsafe.Pointer(q.ref))
	cLabel := toCStr(label)
	C.wgpuQueueSetLabel(cQueue, cLabel)
}

func (q *Queue) Release() {
	cQueue := C.WGPUQueue(unsafe.Pointer(q.ref))
	C.wgpuQueueRelease(cQueue)
}

type RenderBundle struct {
	ref uintptr
}

func (r *RenderBundle) SetLabel(label string) {
	cRenderBundle := C.WGPURenderBundle(unsafe.Pointer(r.ref))
	cLabel := toCStr(label)
	C.wgpuRenderBundleSetLabel(cRenderBundle, cLabel)
}

type RenderBundleEncoder struct {
	ref C.WGPURenderBundleEncoder
}

func (r *RenderBundleEncoder) SetPipeline(pipeline *RenderPipeline) {
	cPipeline := C.WGPURenderPipeline(unsafe.Pointer(pipeline.ref))
	C.wgpuRenderBundleEncoderSetPipeline(r.ref, cPipeline)
}

func (r *RenderBundleEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsetCount int, dynamicOffsets uint32) {

	cGroupIndex := C.uint32_t(groupIndex)
	cGroup := C.WGPUBindGroup(unsafe.Pointer(group.ref))
	cDynamicOffsetCount := C.size_t(dynamicOffsetCount)
	cDynamicOffsets := C.uint32_t(dynamicOffsets)

	C.wgpuRenderBundleEncoderSetBindGroup(r.ref, cGroupIndex, cGroup, cDynamicOffsetCount, &cDynamicOffsets)
}

func (r *RenderBundleEncoder) Draw(vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	cVertexCount := C.uint32_t(vertexCount)
	cInstanceCount := C.uint32_t(instanceCount)
	cFirstVertex := C.uint32_t(firstVertex)
	cFirstInstance := C.uint32_t(firstInstance)

	C.wgpuRenderBundleEncoderDraw(r.ref, cVertexCount, cInstanceCount, cFirstVertex, cFirstInstance)
}

func (r *RenderBundleEncoder) DrawIndexed(indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	cIndexCount := C.uint32_t(indexCount)
	cInstanceCount := C.uint32_t(instanceCount)
	cFirstIndex := C.uint32_t(firstIndex)
	cBaseVertex := C.int32_t(baseVertex)
	cFirstInstance := C.uint32_t(firstInstance)

	C.wgpuRenderBundleEncoderDrawIndexed(r.ref, cIndexCount, cInstanceCount, cFirstIndex, cBaseVertex, cFirstInstance)
}

func (r *RenderBundleEncoder) DrawIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))
	cIndirectOffset := C.uint64_t(indirectOffset)
	cIndirectBuffer := C.WGPUBuffer(unsafe.Pointer(indirectBuffer.ref))

	C.wgpuRenderBundleEncoderDrawIndirect(cRenderBundleEncoder, cIndirectBuffer, cIndirectOffset)
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
	ref C.WGPURenderPipeline
}

func (r *RenderPipeline) GetBindGroupLayout(groupIndex uint32) *BindGroupLayout {
	cGroupIndex := C.uint32_t(groupIndex)
	return &BindGroupLayout{ref: C.wgpuRenderPipelineGetBindGroupLayout(r.ref, cGroupIndex)}
}

func (r *RenderPipeline) SetLabel(label string) {
	cLabel := toCStr(label)
	C.wgpuRenderPipelineSetLabel(r.ref, cLabel)
}

func (r *RenderPipeline) Release() {
	cRenderPipeline := C.WGPURenderPipeline(unsafe.Pointer(r.ref))
	C.wgpuRenderPipelineRelease(cRenderPipeline)
}

type Sampler struct {
	ref C.WGPUSampler
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
	ref C.WGPUShaderModule
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

	callback := compilationInfoCallback(func(s compilationInfoRequestStatus, i CompilationInfo) {
		status = s
		info = i
	})

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

func (s *ShaderModule) Release() {
	cShaderModule := C.WGPUShaderModule(unsafe.Pointer(s.ref))
	C.wgpuShaderModuleRelease(cShaderModule)
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

func (s *Surface) GetCapabilities(adapter *Adapter) (SurfaceCapabilities, error) {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))
	pAdapter := C.WGPUAdapter(unsafe.Pointer(adapter.ref))

	var cCapabilities C.WGPUSurfaceCapabilities

	status := statusCode(C.wgpuSurfaceGetCapabilities(cSurface, pAdapter, &cCapabilities))
	defer C.wgpuSurfaceCapabilitiesFreeMembers(cCapabilities)

	if status != statusCodeSuccess {
		return SurfaceCapabilities{}, fmt.Errorf("error getting surface capabilities: %v", status)
	}

	var capabilities SurfaceCapabilities
	capabilities.Usages = TextureUsage(cCapabilities.usages)

	if count := cCapabilities.formatCount; count > 0 {
		cFormats := unsafe.Slice((*C.WGPUTextureFormat)(cCapabilities.formats), count)
		capabilities.Formats = make([]TextureFormat, count)
		for i := range cFormats {
			capabilities.Formats[i] = TextureFormat(cFormats[i])
		}
	}

	if count := cCapabilities.presentModeCount; count > 0 {
		cPresentModes := unsafe.Slice((*C.WGPUPresentMode)(cCapabilities.presentModes), count)
		capabilities.PresentModes = make([]PresentMode, count)
		for i := range cPresentModes {
			capabilities.PresentModes[i] = PresentMode(cPresentModes[i])
		}
	}

	if count := cCapabilities.alphaModeCount; count > 0 {
		cAlphaModes := unsafe.Slice((*C.WGPUCompositeAlphaMode)(cCapabilities.alphaModes), count)
		capabilities.AlphaModes = make([]CompositeAlphaMode, count)
		for i := range cAlphaModes {
			capabilities.AlphaModes[i] = CompositeAlphaMode(cAlphaModes[i])
		}
	}

	return capabilities, nil
}

func (s *Surface) GetCurrentTexture() *SurfaceTexture {
	surfaceTexture, err := s.TryGetCurrentTexture()
	if err != nil {
		panic(err)
	}
	return surfaceTexture
}

func (s *Surface) TryGetCurrentTexture() (*SurfaceTexture, error) {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))

	var cSurfaceTexture C.WGPUSurfaceTexture
	C.wgpuSurfaceGetCurrentTexture(cSurface, &cSurfaceTexture)

	status := SurfaceGetCurrentTextureStatus(cSurfaceTexture.status)
	if status != SurfaceGetCurrentTextureStatusSuccessOptimal && status != SurfaceGetCurrentTextureStatusSuccessSuboptimal {
		return nil, fmt.Errorf("error getting current texture")
	}

	surfaceTexture := &SurfaceTexture{
		Texture: &Texture{ref: cSurfaceTexture.texture},
		Status:  status,
	}

	return surfaceTexture, nil
}

func (s *Surface) Release() {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))
	C.wgpuSurfaceRelease(cSurface)
}

type SurfaceTexture struct {
	Texture *Texture
	Status  SurfaceGetCurrentTextureStatus
}

func (s *SurfaceTexture) CreateView(descriptor *TextureViewDescriptor) *TextureView {
	return s.Texture.CreateView(descriptor)
}

func (s *SurfaceTexture) TryCreateView(descriptor *TextureViewDescriptor) (*TextureView, error) {
	return s.Texture.TryCreateView(descriptor)
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
	ref C.WGPUTexture
}

func (t *Texture) AsImageCopy() TexelCopyTextureInfo {
	return TexelCopyTextureInfo{
		Texture:  t,
		MipLevel: 0,
		Origin:   Origin3D{},
		Aspect:   TextureAspectAll,
	}
}

func (t *Texture) CreateView(descriptor *TextureViewDescriptor) *TextureView {
	textureView, err := t.TryCreateView(descriptor)
	if err != nil {
		panic(err)
	}
	return textureView
}

func (t *Texture) TryCreateView(descriptor *TextureViewDescriptor) (*TextureView, error) {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))

	var pDescriptor *C.WGPUTextureViewDescriptor
	if descriptor != nil {
		pDescriptor = &C.WGPUTextureViewDescriptor{
			format:          C.WGPUTextureFormat(descriptor.Format),
			dimension:       C.WGPUTextureViewDimension(descriptor.Dimension),
			baseMipLevel:    C.uint32_t(descriptor.BaseMipLevel),
			mipLevelCount:   C.uint32_t(descriptor.MipLevelCount),
			baseArrayLayer:  C.uint32_t(descriptor.BaseArrayLayer),
			arrayLayerCount: C.uint32_t(descriptor.ArrayLayerCount),
			aspect:          C.WGPUTextureAspect(descriptor.Aspect),
			usage:           C.WGPUTextureUsage(descriptor.Usage),
		}

		if descriptor.Label != "" {
			pDescriptor.label.length = C.size_t(len(descriptor.Label))
			pDescriptor.label.data = C.CString(descriptor.Label)
			defer C.free(unsafe.Pointer(pDescriptor.label.data))
		}
	}

	ptr := unsafe.Pointer(C.wgpuTextureCreateView(cTexture, pDescriptor))
	if ptr == nil {
		return nil, fmt.Errorf("error creating view")
	}

	return &TextureView{ref: uintptr(ptr)}, nil
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

func (t *TextureView) Release() {
	cTextureView := C.WGPUTextureView(unsafe.Pointer(t.ref))
	C.wgpuTextureViewRelease(cTextureView)
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

	return Instance{ref: C.wgpuCreateInstance(&pDescriptor)}
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
