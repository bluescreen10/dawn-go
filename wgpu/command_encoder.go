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
	var pinner runtime.Pinner
	defer pinner.Unpin()

	cDescriptor := C.WGPURenderPassDescriptor{
		label: toCStr(descriptor.Label),
	}

	if count := C.size_t(len(descriptor.ColorAttachments)); count > 0 {
		colorAttachments := make([]C.WGPURenderPassColorAttachment, count)
		pinner.Pin(&colorAttachments[0])

		cDescriptor.colorAttachments = (*C.WGPURenderPassColorAttachment)(unsafe.Pointer(&colorAttachments[0]))
		cDescriptor.colorAttachmentCount = count

		for i, a := range descriptor.ColorAttachments {
			colorAttachments[i].depthSlice = C.WGPU_DEPTH_SLICE_UNDEFINED
			colorAttachments[i].loadOp = C.WGPULoadOp(a.LoadOp)
			colorAttachments[i].storeOp = C.WGPUStoreOp(a.StoreOp)
			colorAttachments[i].clearValue.r = C.double(a.ClearValue.R)
			colorAttachments[i].clearValue.g = C.double(a.ClearValue.G)
			colorAttachments[i].clearValue.b = C.double(a.ClearValue.B)
			colorAttachments[i].clearValue.a = C.double(a.ClearValue.A)

			if a.View != nil {
				colorAttachments[i].view = a.View.ref
			}

			if a.ResolveTarget != nil {
				colorAttachments[i].resolveTarget = a.ResolveTarget.ref
			}
		}
	}

	if descriptor.DepthStencilAttachment != nil {
		cDescriptor.depthStencilAttachment = &C.WGPURenderPassDepthStencilAttachment{
			depthLoadOp:       C.WGPULoadOp(descriptor.DepthStencilAttachment.DepthLoadOp),
			depthStoreOp:      C.WGPUStoreOp(descriptor.DepthStencilAttachment.DepthStoreOp),
			depthClearValue:   C.float(descriptor.DepthStencilAttachment.DepthClearValue),
			depthReadOnly:     toCBool(descriptor.DepthStencilAttachment.DepthReadOnly),
			stencilLoadOp:     C.WGPULoadOp(descriptor.DepthStencilAttachment.StencilLoadOp),
			stencilStoreOp:    C.WGPUStoreOp(descriptor.DepthStencilAttachment.StencilStoreOp),
			stencilClearValue: C.uint32_t(descriptor.DepthStencilAttachment.StencilClearValue),
			stencilReadOnly:   toCBool(descriptor.DepthStencilAttachment.StencilReadOnly),
		}
		pinner.Pin(cDescriptor.depthStencilAttachment)

		if descriptor.DepthStencilAttachment.View != nil {
			cDescriptor.depthStencilAttachment.view = descriptor.DepthStencilAttachment.View.ref
		}
	}

	if descriptor.OcclusionQuerySet != nil {
		cDescriptor.occlusionQuerySet = descriptor.OcclusionQuerySet.ref
	}

	if descriptor.TimestampWrites != nil {
		cDescriptor.timestampWrites.querySet = descriptor.TimestampWrites.QuerySet.ref
		cDescriptor.timestampWrites.beginningOfPassWriteIndex = C.uint32_t(descriptor.TimestampWrites.BeginningOfPassWriteIndex)
		cDescriptor.timestampWrites.endOfPassWriteIndex = C.uint32_t(descriptor.TimestampWrites.EndOfPassWriteIndex)
	}

	return &RenderPassEncoder{ref: C.wgpuCommandEncoderBeginRenderPass(c.ref, &cDescriptor)}
}

func (c *CommandEncoder) CopyBufferToBuffer(source *Buffer, sourceOffset uint64, destination *Buffer, destinationOffset uint64, size uint64) {
	C.wgpuCommandEncoderCopyBufferToBuffer(c.ref, source.ref, C.uint64_t(sourceOffset), destination.ref, C.uint64_t(destinationOffset), C.uint64_t(size))
}

func (c *CommandEncoder) CopyBufferToTexture(source TexelCopyBufferInfo, destination TexelCopyTextureInfo, copySize Extent3D) {

	cSource := C.WGPUTexelCopyBufferInfo{
		layout: C.WGPUTexelCopyBufferLayout{
			offset:       C.uint64_t(source.Layout.Offset),
			bytesPerRow:  C.uint32_t(source.Layout.BytesPerRow),
			rowsPerImage: C.uint32_t(source.Layout.RowsPerImage),
		},
		buffer: source.Buffer.ref,
	}

	cDestination := C.WGPUTexelCopyTextureInfo{
		texture:  destination.Texture.ref,
		mipLevel: C.uint32_t(destination.MipLevel),
		origin: C.WGPUOrigin3D{
			x: C.uint32_t(destination.Origin.X),
			y: C.uint32_t(destination.Origin.Y),
			z: C.uint32_t(destination.Origin.Z),
		},
		aspect: C.WGPUTextureAspect(destination.Aspect),
	}

	cCopySize := C.WGPUExtent3D{
		width:              C.uint32_t(copySize.Width),
		height:             C.uint32_t(copySize.Height),
		depthOrArrayLayers: C.uint32_t(copySize.DepthOrArrayLayers),
	}

	C.wgpuCommandEncoderCopyBufferToTexture(c.ref, &cSource, &cDestination, &cCopySize)
}

func (c *CommandEncoder) CopyTextureToBuffer(source TexelCopyTextureInfo, destination TexelCopyBufferInfo, copySize Extent3D) {

	cSource := C.WGPUTexelCopyTextureInfo{
		texture:  source.Texture.ref,
		mipLevel: C.uint32_t(source.MipLevel),
		origin: C.WGPUOrigin3D{
			x: C.uint32_t(source.Origin.X),
			y: C.uint32_t(source.Origin.Y),
			z: C.uint32_t(source.Origin.Z),
		},
		aspect: C.WGPUTextureAspect(source.Aspect),
	}

	cDestination := C.WGPUTexelCopyBufferInfo{
		layout: C.WGPUTexelCopyBufferLayout{
			offset:       C.uint64_t(destination.Layout.Offset),
			bytesPerRow:  C.uint32_t(destination.Layout.BytesPerRow),
			rowsPerImage: C.uint32_t(destination.Layout.RowsPerImage),
		},
		buffer: destination.Buffer.ref,
	}

	cCopySize := C.WGPUExtent3D{
		width:              C.uint32_t(copySize.Width),
		height:             C.uint32_t(copySize.Height),
		depthOrArrayLayers: C.uint32_t(copySize.DepthOrArrayLayers),
	}

	C.wgpuCommandEncoderCopyTextureToBuffer(c.ref, &cSource, &cDestination, &cCopySize)
}

func (c *CommandEncoder) CopyTextureToTexture(source TexelCopyTextureInfo, destination TexelCopyTextureInfo, copySize Extent3D) {

	cSource := C.WGPUTexelCopyTextureInfo{
		texture:  source.Texture.ref,
		mipLevel: C.uint32_t(source.MipLevel),
		origin: C.WGPUOrigin3D{
			x: C.uint32_t(source.Origin.X),
			y: C.uint32_t(source.Origin.Y),
			z: C.uint32_t(source.Origin.Z),
		},
		aspect: C.WGPUTextureAspect(source.Aspect),
	}

	cDestination := C.WGPUTexelCopyTextureInfo{
		texture:  destination.Texture.ref,
		mipLevel: C.uint32_t(destination.MipLevel),
		origin: C.WGPUOrigin3D{
			x: C.uint32_t(destination.Origin.X),
			y: C.uint32_t(destination.Origin.Y),
			z: C.uint32_t(destination.Origin.Z),
		},
		aspect: C.WGPUTextureAspect(destination.Aspect),
	}

	cCopySize := C.WGPUExtent3D{
		width:              C.uint32_t(copySize.Width),
		height:             C.uint32_t(copySize.Height),
		depthOrArrayLayers: C.uint32_t(copySize.DepthOrArrayLayers),
	}

	C.wgpuCommandEncoderCopyTextureToTexture(c.ref, &cSource, &cDestination, &cCopySize)
}

func (c *CommandEncoder) ClearBuffer(buffer *Buffer, offset uint64, size uint64) {
	C.wgpuCommandEncoderClearBuffer(c.ref, buffer.ref, C.uint64_t(offset), C.uint64_t(size))
}

func (c *CommandEncoder) InsertDebugMarker(markerLabel string) {
	C.wgpuCommandEncoderInsertDebugMarker(c.ref, toCStr(markerLabel))
}

func (c *CommandEncoder) PopDebugGroup() {
	C.wgpuCommandEncoderPopDebugGroup(c.ref)
}

func (c *CommandEncoder) PushDebugGroup(groupLabel string) {
	C.wgpuCommandEncoderPushDebugGroup(c.ref, toCStr(groupLabel))
}

func (c *CommandEncoder) ResolveQuerySet(querySet *QuerySet, firstQuery uint32, queryCount uint32, destination *Buffer, destinationOffset uint64) {
	C.wgpuCommandEncoderResolveQuerySet(c.ref, querySet.ref, C.uint32_t(firstQuery), C.uint32_t(queryCount), destination.ref, C.uint64_t(destinationOffset))
}

func (c *CommandEncoder) WriteTimestamp(querySet *QuerySet, queryIndex uint32) {
	C.wgpuCommandEncoderWriteTimestamp(c.ref, querySet.ref, C.uint32_t(queryIndex))
}

func (c *CommandEncoder) SetLabel(label string) {
	C.wgpuCommandEncoderSetLabel(c.ref, toCStr(label))
}

func (c *CommandEncoder) Release() {
	C.wgpuCommandEncoderRelease(c.ref)
}
