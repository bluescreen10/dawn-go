//go:build !js

package wgpu

/*
#include <stdlib.h>
#include "webgpu.h"
*/
import "C"
import "unsafe"

type CommandEncoder struct {
	ref C.WGPUCommandEncoder
}

func (c *CommandEncoder) Finish(descriptor *CommandBufferDescriptor) *CommandBuffer {

	var cDescriptor *C.WGPUCommandBufferDescriptor

	if descriptor != nil {
		cDescriptor = &C.WGPUCommandBufferDescriptor{
			label: toCStr(descriptor.Label),
		}
	}

	return &CommandBuffer{ref: C.wgpuCommandEncoderFinish(c.ref, cDescriptor)}
}

func (c *CommandEncoder) BeginComputePass(descriptor *ComputePassDescriptor) ComputePassEncoder {

	var cDescriptor *C.WGPUComputePassDescriptor
	if descriptor != nil {
		cDescriptor = &C.WGPUComputePassDescriptor{
			label: toCStr(descriptor.Label),
		}

		if descriptor.TimestampWrites != nil {
			cDescriptor.timestampWrites = &C.WGPUPassTimestampWrites{
				querySet:                  C.WGPUQuerySet(descriptor.TimestampWrites.QuerySet.ref),
				beginningOfPassWriteIndex: C.uint32_t(descriptor.TimestampWrites.BeginningOfPassWriteIndex),
				endOfPassWriteIndex:       C.uint32_t(descriptor.TimestampWrites.EndOfPassWriteIndex),
			}
		}
	}

	return ComputePassEncoder{ref: C.wgpuCommandEncoderBeginComputePass(c.ref, cDescriptor)}
}

func (c *CommandEncoder) BeginRenderPass(descriptor RenderPassDescriptor) *RenderPassEncoder {

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
			slice[i].depthSlice = C.WGPU_DEPTH_SLICE_UNDEFINED
			slice[i].loadOp = C.WGPULoadOp(a.LoadOp)
			slice[i].storeOp = C.WGPUStoreOp(a.StoreOp)
			slice[i].clearValue.r = C.double(a.ClearValue.R)
			slice[i].clearValue.g = C.double(a.ClearValue.G)
			slice[i].clearValue.b = C.double(a.ClearValue.B)
			slice[i].clearValue.a = C.double(a.ClearValue.A)

			if a.View != nil {
				slice[i].view = C.WGPUTextureView(unsafe.Pointer(a.View.ref))
			}

			if a.DepthSlice != 0 {
				slice[i].depthSlice = C.uint32_t(a.DepthSlice)
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

	return &RenderPassEncoder{ref: C.wgpuCommandEncoderBeginRenderPass(c.ref, &cDescriptor)}
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
	cLabel := toCStr(label)
	C.wgpuCommandEncoderSetLabel(c.ref, cLabel)
}

func (c *CommandEncoder) Release() {
	C.wgpuCommandEncoderRelease(c.ref)
}
