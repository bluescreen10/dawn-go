// CODE GENERATED. DO NOT EDIT
//
//go:generate go run ./cmd/wrapper/.
package wgpu

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L./lib -lwebgpu_dawn -framework Metal -framework IOKit -framework QuartzCore -framework Foundation -framework IOSurface -lc++
#include <webgpu/webgpu.h>
#include <stdio.h>
#include <stdlib.h>

// This is the C-side function that WebGPU will actually call.
// It just forwards the call to our Go function "goCallbackHandler".
extern void c_callback_trampoline(WGPURequestAdapterStatus status, WGPUAdapter adapter, WGPUStringView message, void * userdata1, void * userdata2);

*/
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

func boolToWGPUBool(in bool) C.WGPUBool {
	if in {
		return 1
	} else {
		return 0
	}
}

type BindGroup struct {
	ref uintptr
}

func (b *BindGroup) SetLabel(label string) {
	cBindGroup := C.WGPUBindGroup(unsafe.Pointer(b.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuBindGroupSetLabel(cBindGroup, cLabel)
}

type Texture struct {
	ref uintptr
}

func (t *Texture) CreateView(descriptor *TextureViewDescriptor) TextureView {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))
	// Convert descriptor to C.WGPUTextureViewDescriptor
	var pDescriptor *C.WGPUTextureViewDescriptor
	if descriptor != nil {
		var cDescriptor C.WGPUTextureViewDescriptor
		cLabelStr0 := C.CString(descriptor.Label)
		cDescriptor.label.data = cLabelStr0
		cDescriptor.label.length = C.size_t(len(descriptor.Label))
		defer C.free(unsafe.Pointer(cLabelStr0))
		pDescriptor = &cDescriptor
	}
	return TextureView{ref: uintptr(unsafe.Pointer(C.wgpuTextureCreateView(cTexture, pDescriptor)))}
}

func (t *Texture) SetLabel(label string) {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuTextureSetLabel(cTexture, cLabel)
}

func (t *Texture) GetWidth() uint32 {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))
	return uint32(C.wgpuTextureGetWidth(cTexture))
}

func (t *Texture) GetHeight() uint32 {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))
	return uint32(C.wgpuTextureGetHeight(cTexture))
}

func (t *Texture) GetDepthOrArrayLayers() uint32 {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))
	return uint32(C.wgpuTextureGetDepthOrArrayLayers(cTexture))
}

func (t *Texture) GetMipLevelCount() uint32 {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))
	return uint32(C.wgpuTextureGetMipLevelCount(cTexture))
}

func (t *Texture) GetSampleCount() uint32 {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))
	return uint32(C.wgpuTextureGetSampleCount(cTexture))
}

func (t *Texture) GetDimension() TextureDimension {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))
	return TextureDimension(C.wgpuTextureGetDimension(cTexture))
}

func (t *Texture) GetFormat() TextureFormat {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))
	return TextureFormat(C.wgpuTextureGetFormat(cTexture))
}

func (t *Texture) GetUsage() TextureUsage {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))
	return TextureUsage(C.wgpuTextureGetUsage(cTexture))
}

func (t *Texture) GetTextureBindingViewDimension() TextureViewDimension {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))
	return TextureViewDimension(C.wgpuTextureGetTextureBindingViewDimension(cTexture))
}

func (t *Texture) Destroy() {
	cTexture := C.WGPUTexture(unsafe.Pointer(t.ref))
	C.wgpuTextureDestroy(cTexture)
}

type Sampler struct {
	ref uintptr
}

func (s *Sampler) SetLabel(label string) {
	cSampler := C.WGPUSampler(unsafe.Pointer(s.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuSamplerSetLabel(cSampler, cLabel)
}

type ComputePassEncoder struct {
	ref uintptr
}

func (c *ComputePassEncoder) InsertDebugMarker(markerLabel string) {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))
	// Convert markerLabel to C.WGPUStringView
	var cMarkerLabel C.WGPUStringView
	cMarkerLabelStr := C.CString(markerLabel)
	cMarkerLabel.data = cMarkerLabelStr
	cMarkerLabel.length = C.size_t(len(markerLabel))
	defer C.free(unsafe.Pointer(cMarkerLabelStr))
	C.wgpuComputePassEncoderInsertDebugMarker(cComputePassEncoder, cMarkerLabel)
}

func (c *ComputePassEncoder) PopDebugGroup() {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))
	C.wgpuComputePassEncoderPopDebugGroup(cComputePassEncoder)
}

func (c *ComputePassEncoder) PushDebugGroup(groupLabel string) {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))
	// Convert groupLabel to C.WGPUStringView
	var cGroupLabel C.WGPUStringView
	cGroupLabelStr := C.CString(groupLabel)
	cGroupLabel.data = cGroupLabelStr
	cGroupLabel.length = C.size_t(len(groupLabel))
	defer C.free(unsafe.Pointer(cGroupLabelStr))
	C.wgpuComputePassEncoderPushDebugGroup(cComputePassEncoder, cGroupLabel)
}

func (c *ComputePassEncoder) SetPipeline(pipeline ComputePipeline) {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))
	cPipeline := C.WGPUComputePipeline(unsafe.Pointer(pipeline.ref))
	C.wgpuComputePassEncoderSetPipeline(cComputePassEncoder, cPipeline)
}

func (c *ComputePassEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsetCount int, dynamicOffsets uint32) {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))
	cGroupIndex := C.uint32_t(groupIndex)
	pGroup := C.WGPUBindGroup(unsafe.Pointer(group.ref))
	cDynamicOffsetCount := C.size_t(dynamicOffsetCount)
	cDynamicOffsets := C.uint32_t(dynamicOffsets)
	C.wgpuComputePassEncoderSetBindGroup(cComputePassEncoder, cGroupIndex, pGroup, cDynamicOffsetCount, &cDynamicOffsets)
}

func (c *ComputePassEncoder) DispatchWorkgroups(workgroupCountX uint32, workgroupCountY uint32, workgroupCountZ uint32) {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))
	cWorkgroupCountX := C.uint32_t(workgroupCountX)
	cWorkgroupCountY := C.uint32_t(workgroupCountY)
	cWorkgroupCountZ := C.uint32_t(workgroupCountZ)
	C.wgpuComputePassEncoderDispatchWorkgroups(cComputePassEncoder, cWorkgroupCountX, cWorkgroupCountY, cWorkgroupCountZ)
}

func (c *ComputePassEncoder) DispatchWorkgroupsIndirect(indirectBuffer Buffer, indirectOffset uint64) {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))
	cIndirectBuffer := C.WGPUBuffer(unsafe.Pointer(indirectBuffer.ref))
	cIndirectOffset := C.uint64_t(indirectOffset)
	C.wgpuComputePassEncoderDispatchWorkgroupsIndirect(cComputePassEncoder, cIndirectBuffer, cIndirectOffset)
}

func (c *ComputePassEncoder) End() {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))
	C.wgpuComputePassEncoderEnd(cComputePassEncoder)
}

func (c *ComputePassEncoder) SetLabel(label string) {
	cComputePassEncoder := C.WGPUComputePassEncoder(unsafe.Pointer(c.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuComputePassEncoderSetLabel(cComputePassEncoder, cLabel)
}

type RenderBundleEncoder struct {
	ref uintptr
}

func (r *RenderBundleEncoder) SetPipeline(pipeline RenderPipeline) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))
	cPipeline := C.WGPURenderPipeline(unsafe.Pointer(pipeline.ref))
	C.wgpuRenderBundleEncoderSetPipeline(cRenderBundleEncoder, cPipeline)
}

func (r *RenderBundleEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsetCount int, dynamicOffsets uint32) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))
	cGroupIndex := C.uint32_t(groupIndex)
	pGroup := C.WGPUBindGroup(unsafe.Pointer(group.ref))
	cDynamicOffsetCount := C.size_t(dynamicOffsetCount)
	cDynamicOffsets := C.uint32_t(dynamicOffsets)
	C.wgpuRenderBundleEncoderSetBindGroup(cRenderBundleEncoder, cGroupIndex, pGroup, cDynamicOffsetCount, &cDynamicOffsets)
}

func (r *RenderBundleEncoder) Draw(vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))
	cVertexCount := C.uint32_t(vertexCount)
	cInstanceCount := C.uint32_t(instanceCount)
	cFirstVertex := C.uint32_t(firstVertex)
	cFirstInstance := C.uint32_t(firstInstance)
	C.wgpuRenderBundleEncoderDraw(cRenderBundleEncoder, cVertexCount, cInstanceCount, cFirstVertex, cFirstInstance)
}

func (r *RenderBundleEncoder) DrawIndexed(indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))
	cIndexCount := C.uint32_t(indexCount)
	cInstanceCount := C.uint32_t(instanceCount)
	cFirstIndex := C.uint32_t(firstIndex)
	cBaseVertex := C.int32_t(baseVertex)
	cFirstInstance := C.uint32_t(firstInstance)
	C.wgpuRenderBundleEncoderDrawIndexed(cRenderBundleEncoder, cIndexCount, cInstanceCount, cFirstIndex, cBaseVertex, cFirstInstance)
}

func (r *RenderBundleEncoder) DrawIndirect(indirectBuffer Buffer, indirectOffset uint64) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))
	cIndirectBuffer := C.WGPUBuffer(unsafe.Pointer(indirectBuffer.ref))
	cIndirectOffset := C.uint64_t(indirectOffset)
	C.wgpuRenderBundleEncoderDrawIndirect(cRenderBundleEncoder, cIndirectBuffer, cIndirectOffset)
}

func (r *RenderBundleEncoder) DrawIndexedIndirect(indirectBuffer Buffer, indirectOffset uint64) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))
	cIndirectBuffer := C.WGPUBuffer(unsafe.Pointer(indirectBuffer.ref))
	cIndirectOffset := C.uint64_t(indirectOffset)
	C.wgpuRenderBundleEncoderDrawIndexedIndirect(cRenderBundleEncoder, cIndirectBuffer, cIndirectOffset)
}

func (r *RenderBundleEncoder) InsertDebugMarker(markerLabel string) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))
	// Convert markerLabel to C.WGPUStringView
	var cMarkerLabel C.WGPUStringView
	cMarkerLabelStr := C.CString(markerLabel)
	cMarkerLabel.data = cMarkerLabelStr
	cMarkerLabel.length = C.size_t(len(markerLabel))
	defer C.free(unsafe.Pointer(cMarkerLabelStr))
	C.wgpuRenderBundleEncoderInsertDebugMarker(cRenderBundleEncoder, cMarkerLabel)
}

func (r *RenderBundleEncoder) PopDebugGroup() {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))
	C.wgpuRenderBundleEncoderPopDebugGroup(cRenderBundleEncoder)
}

func (r *RenderBundleEncoder) PushDebugGroup(groupLabel string) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))
	// Convert groupLabel to C.WGPUStringView
	var cGroupLabel C.WGPUStringView
	cGroupLabelStr := C.CString(groupLabel)
	cGroupLabel.data = cGroupLabelStr
	cGroupLabel.length = C.size_t(len(groupLabel))
	defer C.free(unsafe.Pointer(cGroupLabelStr))
	C.wgpuRenderBundleEncoderPushDebugGroup(cRenderBundleEncoder, cGroupLabel)
}

func (r *RenderBundleEncoder) SetVertexBuffer(slot uint32, buffer *Buffer, offset uint64, size uint64) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))
	cSlot := C.uint32_t(slot)
	pBuffer := C.WGPUBuffer(unsafe.Pointer(buffer.ref))
	cOffset := C.uint64_t(offset)
	cSize := C.uint64_t(size)
	C.wgpuRenderBundleEncoderSetVertexBuffer(cRenderBundleEncoder, cSlot, pBuffer, cOffset, cSize)
}

func (r *RenderBundleEncoder) SetIndexBuffer(buffer Buffer, format IndexFormat, offset uint64, size uint64) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))
	cBuffer := C.WGPUBuffer(unsafe.Pointer(buffer.ref))
	cFormat := C.WGPUIndexFormat(format)
	cOffset := C.uint64_t(offset)
	cSize := C.uint64_t(size)
	C.wgpuRenderBundleEncoderSetIndexBuffer(cRenderBundleEncoder, cBuffer, cFormat, cOffset, cSize)
}

func (r *RenderBundleEncoder) Finish(descriptor *RenderBundleDescriptor) RenderBundle {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))
	// Convert descriptor to C.WGPURenderBundleDescriptor
	var pDescriptor *C.WGPURenderBundleDescriptor
	if descriptor != nil {
		var cDescriptor C.WGPURenderBundleDescriptor
		cLabelStr0 := C.CString(descriptor.Label)
		cDescriptor.label.data = cLabelStr0
		cDescriptor.label.length = C.size_t(len(descriptor.Label))
		defer C.free(unsafe.Pointer(cLabelStr0))
		pDescriptor = &cDescriptor
	}
	return RenderBundle{ref: uintptr(unsafe.Pointer(C.wgpuRenderBundleEncoderFinish(cRenderBundleEncoder, pDescriptor)))}
}

func (r *RenderBundleEncoder) SetLabel(label string) {
	cRenderBundleEncoder := C.WGPURenderBundleEncoder(unsafe.Pointer(r.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuRenderBundleEncoderSetLabel(cRenderBundleEncoder, cLabel)
}

type CommandEncoder struct {
	ref uintptr
}

func (c *CommandEncoder) Finish(descriptor *CommandBufferDescriptor) CommandBuffer {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	// Convert descriptor to C.WGPUCommandBufferDescriptor
	var pDescriptor *C.WGPUCommandBufferDescriptor
	if descriptor != nil {
		var cDescriptor C.WGPUCommandBufferDescriptor
		cLabelStr0 := C.CString(descriptor.Label)
		cDescriptor.label.data = cLabelStr0
		cDescriptor.label.length = C.size_t(len(descriptor.Label))
		defer C.free(unsafe.Pointer(cLabelStr0))
		pDescriptor = &cDescriptor
	}
	return CommandBuffer{ref: uintptr(unsafe.Pointer(C.wgpuCommandEncoderFinish(cCommandEncoder, pDescriptor)))}
}

func (c *CommandEncoder) BeginComputePass(descriptor *ComputePassDescriptor) ComputePassEncoder {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	// Convert descriptor to C.WGPUComputePassDescriptor
	var pDescriptor *C.WGPUComputePassDescriptor
	if descriptor != nil {
		var cDescriptor C.WGPUComputePassDescriptor
		cLabelStr0 := C.CString(descriptor.Label)
		cDescriptor.label.data = cLabelStr0
		cDescriptor.label.length = C.size_t(len(descriptor.Label))
		defer C.free(unsafe.Pointer(cLabelStr0))
		pDescriptor = &cDescriptor
	}
	return ComputePassEncoder{ref: uintptr(unsafe.Pointer(C.wgpuCommandEncoderBeginComputePass(cCommandEncoder, pDescriptor)))}
}

func (c *CommandEncoder) BeginRenderPass(descriptor RenderPassDescriptor) RenderPassEncoder {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	// Convert descriptor to C.WGPURenderPassDescriptor
	var cDescriptor C.WGPURenderPassDescriptor
	cLabelStr0 := C.CString(descriptor.Label)
	cDescriptor.label.data = cLabelStr0
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	defer C.free(unsafe.Pointer(cLabelStr0))
	cDescriptor.colorAttachmentCount = C.size_t(descriptor.ColorAttachmentCount)
	if descriptor.OcclusionQuerySet != nil {
		cDescriptor.occlusionQuerySet = C.WGPUQuerySet(unsafe.Pointer(descriptor.OcclusionQuerySet.ref))
	}
	return RenderPassEncoder{ref: uintptr(unsafe.Pointer(C.wgpuCommandEncoderBeginRenderPass(cCommandEncoder, &cDescriptor)))}
}

func (c *CommandEncoder) CopyBufferToBuffer(source Buffer, sourceOffset uint64, destination Buffer, destinationOffset uint64, size uint64) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	cSource := C.WGPUBuffer(unsafe.Pointer(source.ref))
	cSourceOffset := C.uint64_t(sourceOffset)
	cDestination := C.WGPUBuffer(unsafe.Pointer(destination.ref))
	cDestinationOffset := C.uint64_t(destinationOffset)
	cSize := C.uint64_t(size)
	C.wgpuCommandEncoderCopyBufferToBuffer(cCommandEncoder, cSource, cSourceOffset, cDestination, cDestinationOffset, cSize)
}

func (c *CommandEncoder) CopyBufferToTexture(source TexelCopyBufferInfo, destination TexelCopyTextureInfo, copySize Extent3D) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	// Convert source to C.WGPUTexelCopyBufferInfo
	var cSource C.WGPUTexelCopyBufferInfo
	if source.Buffer != nil {
		cSource.buffer = C.WGPUBuffer(unsafe.Pointer(source.Buffer.ref))
	}
	// Convert destination to C.WGPUTexelCopyTextureInfo
	var cDestination C.WGPUTexelCopyTextureInfo
	if destination.Texture != nil {
		cDestination.texture = C.WGPUTexture(unsafe.Pointer(destination.Texture.ref))
	}
	// Convert copySize to C.WGPUExtent3D
	var cCopySize C.WGPUExtent3D
	C.wgpuCommandEncoderCopyBufferToTexture(cCommandEncoder, &cSource, &cDestination, &cCopySize)
}

func (c *CommandEncoder) CopyTextureToBuffer(source TexelCopyTextureInfo, destination TexelCopyBufferInfo, copySize Extent3D) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	// Convert source to C.WGPUTexelCopyTextureInfo
	var cSource C.WGPUTexelCopyTextureInfo
	if source.Texture != nil {
		cSource.texture = C.WGPUTexture(unsafe.Pointer(source.Texture.ref))
	}
	// Convert destination to C.WGPUTexelCopyBufferInfo
	var cDestination C.WGPUTexelCopyBufferInfo
	if destination.Buffer != nil {
		cDestination.buffer = C.WGPUBuffer(unsafe.Pointer(destination.Buffer.ref))
	}
	// Convert copySize to C.WGPUExtent3D
	var cCopySize C.WGPUExtent3D
	C.wgpuCommandEncoderCopyTextureToBuffer(cCommandEncoder, &cSource, &cDestination, &cCopySize)
}

func (c *CommandEncoder) CopyTextureToTexture(source TexelCopyTextureInfo, destination TexelCopyTextureInfo, copySize Extent3D) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	// Convert source to C.WGPUTexelCopyTextureInfo
	var cSource C.WGPUTexelCopyTextureInfo
	if source.Texture != nil {
		cSource.texture = C.WGPUTexture(unsafe.Pointer(source.Texture.ref))
	}
	// Convert destination to C.WGPUTexelCopyTextureInfo
	var cDestination C.WGPUTexelCopyTextureInfo
	if destination.Texture != nil {
		cDestination.texture = C.WGPUTexture(unsafe.Pointer(destination.Texture.ref))
	}
	// Convert copySize to C.WGPUExtent3D
	var cCopySize C.WGPUExtent3D
	C.wgpuCommandEncoderCopyTextureToTexture(cCommandEncoder, &cSource, &cDestination, &cCopySize)
}

func (c *CommandEncoder) ClearBuffer(buffer Buffer, offset uint64, size uint64) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	cBuffer := C.WGPUBuffer(unsafe.Pointer(buffer.ref))
	cOffset := C.uint64_t(offset)
	cSize := C.uint64_t(size)
	C.wgpuCommandEncoderClearBuffer(cCommandEncoder, cBuffer, cOffset, cSize)
}

func (c *CommandEncoder) InsertDebugMarker(markerLabel string) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	// Convert markerLabel to C.WGPUStringView
	var cMarkerLabel C.WGPUStringView
	cMarkerLabelStr := C.CString(markerLabel)
	cMarkerLabel.data = cMarkerLabelStr
	cMarkerLabel.length = C.size_t(len(markerLabel))
	defer C.free(unsafe.Pointer(cMarkerLabelStr))
	C.wgpuCommandEncoderInsertDebugMarker(cCommandEncoder, cMarkerLabel)
}

func (c *CommandEncoder) PopDebugGroup() {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	C.wgpuCommandEncoderPopDebugGroup(cCommandEncoder)
}

func (c *CommandEncoder) PushDebugGroup(groupLabel string) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	// Convert groupLabel to C.WGPUStringView
	var cGroupLabel C.WGPUStringView
	cGroupLabelStr := C.CString(groupLabel)
	cGroupLabel.data = cGroupLabelStr
	cGroupLabel.length = C.size_t(len(groupLabel))
	defer C.free(unsafe.Pointer(cGroupLabelStr))
	C.wgpuCommandEncoderPushDebugGroup(cCommandEncoder, cGroupLabel)
}

func (c *CommandEncoder) ResolveQuerySet(querySet QuerySet, firstQuery uint32, queryCount uint32, destination Buffer, destinationOffset uint64) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	cQuerySet := C.WGPUQuerySet(unsafe.Pointer(querySet.ref))
	cFirstQuery := C.uint32_t(firstQuery)
	cQueryCount := C.uint32_t(queryCount)
	cDestination := C.WGPUBuffer(unsafe.Pointer(destination.ref))
	cDestinationOffset := C.uint64_t(destinationOffset)
	C.wgpuCommandEncoderResolveQuerySet(cCommandEncoder, cQuerySet, cFirstQuery, cQueryCount, cDestination, cDestinationOffset)
}

func (c *CommandEncoder) WriteTimestamp(querySet QuerySet, queryIndex uint32) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	cQuerySet := C.WGPUQuerySet(unsafe.Pointer(querySet.ref))
	cQueryIndex := C.uint32_t(queryIndex)
	C.wgpuCommandEncoderWriteTimestamp(cCommandEncoder, cQuerySet, cQueryIndex)
}

func (c *CommandEncoder) SetLabel(label string) {
	cCommandEncoder := C.WGPUCommandEncoder(unsafe.Pointer(c.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuCommandEncoderSetLabel(cCommandEncoder, cLabel)
}

type Buffer struct {
	ref uintptr
}

func (b *Buffer) MapAsync(mode MapMode, offset int, size int, callbackInfo BufferMapCallbackInfo) Future {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))
	cMode := C.WGPUMapMode(mode)
	cOffset := C.size_t(offset)
	cSize := C.size_t(size)
	// Convert callbackInfo to C.WGPUBufferMapCallbackInfo
	var cCallbackInfo C.WGPUBufferMapCallbackInfo
	return Future{Id: uint64(C.wgpuBufferMapAsync(cBuffer, cMode, cOffset, cSize, cCallbackInfo).id)}
}

func (b *Buffer) GetMappedRange(offset int, size int) uintptr {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))
	cOffset := C.size_t(offset)
	cSize := C.size_t(size)
	return uintptr(C.wgpuBufferGetMappedRange(cBuffer, cOffset, cSize))
}

func (b *Buffer) WriteMappedRange(offset int, data []byte, size int) Status {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))
	cOffset := C.size_t(offset)
	cData := unsafe.Pointer(&data[0])
	cSize := C.size_t(size)
	return Status(C.wgpuBufferWriteMappedRange(cBuffer, cOffset, cData, cSize))
}

func (b *Buffer) ReadMappedRange(offset int, data []byte, size int) Status {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))
	cOffset := C.size_t(offset)
	cData := unsafe.Pointer(&data[0])
	cSize := C.size_t(size)
	return Status(C.wgpuBufferReadMappedRange(cBuffer, cOffset, cData, cSize))
}

func (b *Buffer) SetLabel(label string) {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuBufferSetLabel(cBuffer, cLabel)
}

func (b *Buffer) GetUsage() BufferUsage {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))
	return BufferUsage(C.wgpuBufferGetUsage(cBuffer))
}

func (b *Buffer) GetSize() uint64 {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))
	return uint64(C.wgpuBufferGetSize(cBuffer))
}

func (b *Buffer) GetMapState() BufferMapState {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))
	return BufferMapState(C.wgpuBufferGetMapState(cBuffer))
}

func (b *Buffer) Unmap() {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))
	C.wgpuBufferUnmap(cBuffer)
}

func (b *Buffer) Destroy() {
	cBuffer := C.WGPUBuffer(unsafe.Pointer(b.ref))
	C.wgpuBufferDestroy(cBuffer)
}

type Queue struct {
	ref uintptr
}

func (q *Queue) Submit(commandCount int, commands CommandBuffer) {
	cQueue := C.WGPUQueue(unsafe.Pointer(q.ref))
	cCommandCount := C.size_t(commandCount)
	cCommands := C.WGPUCommandBuffer(unsafe.Pointer(commands.ref))
	C.wgpuQueueSubmit(cQueue, cCommandCount, &cCommands)
}

func (q *Queue) OnSubmittedWorkDone(callbackInfo QueueWorkDoneCallbackInfo) Future {
	cQueue := C.WGPUQueue(unsafe.Pointer(q.ref))
	// Convert callbackInfo to C.WGPUQueueWorkDoneCallbackInfo
	var cCallbackInfo C.WGPUQueueWorkDoneCallbackInfo
	return Future{Id: uint64(C.wgpuQueueOnSubmittedWorkDone(cQueue, cCallbackInfo).id)}
}

func (q *Queue) WriteBuffer(buffer Buffer, bufferOffset uint64, data []byte, size int) {
	cQueue := C.WGPUQueue(unsafe.Pointer(q.ref))
	cBuffer := C.WGPUBuffer(unsafe.Pointer(buffer.ref))
	cBufferOffset := C.uint64_t(bufferOffset)
	cData := unsafe.Pointer(&data[0])
	cSize := C.size_t(size)
	C.wgpuQueueWriteBuffer(cQueue, cBuffer, cBufferOffset, cData, cSize)
}

func (q *Queue) WriteTexture(destination TexelCopyTextureInfo, data []byte, dataSize int, dataLayout TexelCopyBufferLayout, writeSize Extent3D) {
	cQueue := C.WGPUQueue(unsafe.Pointer(q.ref))
	// Convert destination to C.WGPUTexelCopyTextureInfo
	var cDestination C.WGPUTexelCopyTextureInfo
	if destination.Texture != nil {
		cDestination.texture = C.WGPUTexture(unsafe.Pointer(destination.Texture.ref))
	}
	cData := unsafe.Pointer(&data[0])
	cDataSize := C.size_t(dataSize)
	// Convert dataLayout to C.WGPUTexelCopyBufferLayout
	var cDataLayout C.WGPUTexelCopyBufferLayout
	// Convert writeSize to C.WGPUExtent3D
	var cWriteSize C.WGPUExtent3D
	C.wgpuQueueWriteTexture(cQueue, &cDestination, cData, cDataSize, &cDataLayout, &cWriteSize)
}

func (q *Queue) SetLabel(label string) {
	cQueue := C.WGPUQueue(unsafe.Pointer(q.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuQueueSetLabel(cQueue, cLabel)
}

type Instance struct {
	ref uintptr
}

func (i *Instance) CreateSurface(descriptor SurfaceDescriptor) Surface {
	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))
	// Convert descriptor to C.WGPUSurfaceDescriptor
	var cDescriptor C.WGPUSurfaceDescriptor
	cLabelStr0 := C.CString(descriptor.Label)
	cDescriptor.label.data = cLabelStr0
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	defer C.free(unsafe.Pointer(cLabelStr0))
	return Surface{ref: uintptr(unsafe.Pointer(C.wgpuInstanceCreateSurface(cInstance, &cDescriptor)))}
}

func (i *Instance) ProcessEvents() {
	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))
	C.wgpuInstanceProcessEvents(cInstance)
}

func (i *Instance) WaitAny(futures []FutureWaitInfo, timeoutNS uint64) WaitStatus {
	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))
	count := len(futures)

	if count == 0 {
		return WaitStatus(C.wgpuInstanceWaitAny(cInstance, 0, nil, C.uint64_t(timeoutNS)))
	}

	// 1. Create a slice of C structs.
	// This ensures the memory layout is a contiguous array that C can read.
	cFutures := make([]C.WGPUFutureWaitInfo, count)

	for idx, f := range futures {
		// 2. Map BOTH the ID and the completed status
		// Note: WGPUFuture is usually a struct with an 'id' field in Dawn
		cFutures[idx].future = C.WGPUFuture{id: C.uint64_t(f.Future.Id)}
		cFutures[idx].completed = boolToWGPUBool(f.Completed)
	}

	// 3. Pass the pointer to the first element of our C-compatible slice
	return WaitStatus(C.wgpuInstanceWaitAny(
		cInstance,
		C.size_t(count),
		&cFutures[0],
		C.uint64_t(timeoutNS),
	))
}

func (i *Instance) RequestAdapter(options *RequestAdapterOptions, callbackInfo RequestAdapterCallbackInfo) Future {
	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))
	// Convert options to C.WGPURequestAdapterOptions
	var pOptions *C.WGPURequestAdapterOptions
	if options != nil {
		var cOptions C.WGPURequestAdapterOptions
		cOptions.forceFallbackAdapter = boolToWGPUBool(options.ForceFallbackAdapter)
		if options.CompatibleSurface != nil {
			cOptions.compatibleSurface = C.WGPUSurface(unsafe.Pointer(options.CompatibleSurface.ref))
		}
		pOptions = &cOptions
	}
	// Convert callbackInfo to C.WGPURequestAdapterCallbackInfo
	var cCallbackInfo C.WGPURequestAdapterCallbackInfo
	cCallbackInfo.mode = C.WGPUCallbackMode(callbackInfo.Mode)
	handle := cgo.NewHandle(callbackInfo.Callback)
	cCallbackInfo.nextInChain = nil
	cCallbackInfo.callback = C.WGPURequestAdapterCallback(C.c_callback_trampoline)
	cCallbackInfo.userdata1 = unsafe.Pointer(handle)
	cCallbackInfo.userdata2 = unsafe.Pointer(handle)
	return Future{Id: uint64(C.wgpuInstanceRequestAdapter(cInstance, pOptions, cCallbackInfo).id)}
}

//export goCallbackHandler
func goCallbackHandler(status C.WGPURequestAdapterStatus, cAdapter C.WGPUAdapter, cMessage C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	handleID := uintptr(userdata1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	defer handle.Delete()

	// THE FIX: Use the named type defined in your package
	// Instead of: .(func(RequestAdapterStatus, *Adapter, string))
	fn := handle.Value().(RequestAdapterCallback)

	// Safety check for the string (WGPUStringView is not null-terminated)
	var message string
	if cMessage.data != nil && cMessage.length > 0 {
		message = C.GoStringN(cMessage.data, C.int(cMessage.length))
	}

	// Call the function
	fn(
		RequestAdapterStatus(status),
		&Adapter{ref: uintptr(unsafe.Pointer(cAdapter))},
		message,
	)
}

func (i *Instance) HasWGSLLanguageFeature(feature WGSLLanguageFeatureName) bool {
	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))
	cFeature := C.WGPUWGSLLanguageFeatureName(feature)
	return bool(C.wgpuInstanceHasWGSLLanguageFeature(cInstance, cFeature) != 0)
}

func (i *Instance) GetWGSLLanguageFeatures(features SupportedWGSLLanguageFeatures) {
	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))
	// Convert features to C.WGPUSupportedWGSLLanguageFeatures
	var cFeatures C.WGPUSupportedWGSLLanguageFeatures
	cFeatures.featureCount = C.size_t(features.FeatureCount)
	C.wgpuInstanceGetWGSLLanguageFeatures(cInstance, &cFeatures)
}

type RenderBundle struct {
	ref uintptr
}

func (r *RenderBundle) SetLabel(label string) {
	cRenderBundle := C.WGPURenderBundle(unsafe.Pointer(r.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuRenderBundleSetLabel(cRenderBundle, cLabel)
}

type TextureView struct {
	ref uintptr
}

func (t *TextureView) SetLabel(label string) {
	cTextureView := C.WGPUTextureView(unsafe.Pointer(t.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuTextureViewSetLabel(cTextureView, cLabel)
}

type BindGroupLayout struct {
	ref uintptr
}

func (b *BindGroupLayout) SetLabel(label string) {
	cBindGroupLayout := C.WGPUBindGroupLayout(unsafe.Pointer(b.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuBindGroupLayoutSetLabel(cBindGroupLayout, cLabel)
}

type Surface struct {
	ref uintptr
}

func (s *Surface) Configure(config SurfaceConfiguration) {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))
	// Convert config to C.WGPUSurfaceConfiguration
	var cConfig C.WGPUSurfaceConfiguration
	if config.Device != nil {
		cConfig.device = C.WGPUDevice(unsafe.Pointer(config.Device.ref))
	}
	cConfig.viewFormatCount = C.size_t(config.ViewFormatCount)
	C.wgpuSurfaceConfigure(cSurface, &cConfig)
}

func (s *Surface) GetCapabilities(adapter Adapter, capabilities SurfaceCapabilities) Status {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))
	cAdapter := C.WGPUAdapter(unsafe.Pointer(adapter.ref))
	// Convert capabilities to C.WGPUSurfaceCapabilities
	var cCapabilities C.WGPUSurfaceCapabilities
	cCapabilities.formatCount = C.size_t(capabilities.FormatCount)
	cCapabilities.presentModeCount = C.size_t(capabilities.PresentModeCount)
	cCapabilities.alphaModeCount = C.size_t(capabilities.AlphaModeCount)
	return Status(C.wgpuSurfaceGetCapabilities(cSurface, cAdapter, &cCapabilities))
}

func (s *Surface) GetCurrentTexture(surfaceTexture SurfaceTexture) {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))
	// Convert surfaceTexture to C.WGPUSurfaceTexture
	var cSurfaceTexture C.WGPUSurfaceTexture
	if surfaceTexture.Texture != nil {
		cSurfaceTexture.texture = C.WGPUTexture(unsafe.Pointer(surfaceTexture.Texture.ref))
	}
	C.wgpuSurfaceGetCurrentTexture(cSurface, &cSurfaceTexture)
}

func (s *Surface) Present() Status {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))
	return Status(C.wgpuSurfacePresent(cSurface))
}

func (s *Surface) Unconfigure() {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))
	C.wgpuSurfaceUnconfigure(cSurface)
}

func (s *Surface) SetLabel(label string) {
	cSurface := C.WGPUSurface(unsafe.Pointer(s.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuSurfaceSetLabel(cSurface, cLabel)
}

type RenderPipeline struct {
	ref uintptr
}

func (r *RenderPipeline) GetBindGroupLayout(groupIndex uint32) BindGroupLayout {
	cRenderPipeline := C.WGPURenderPipeline(unsafe.Pointer(r.ref))
	cGroupIndex := C.uint32_t(groupIndex)
	return BindGroupLayout{ref: uintptr(unsafe.Pointer(C.wgpuRenderPipelineGetBindGroupLayout(cRenderPipeline, cGroupIndex)))}
}

func (r *RenderPipeline) SetLabel(label string) {
	cRenderPipeline := C.WGPURenderPipeline(unsafe.Pointer(r.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuRenderPipelineSetLabel(cRenderPipeline, cLabel)
}

type Device struct {
	ref uintptr
}

func (d *Device) CreateBindGroup(descriptor BindGroupDescriptor) BindGroup {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert descriptor to C.WGPUBindGroupDescriptor
	var cDescriptor C.WGPUBindGroupDescriptor
	cLabelStr0 := C.CString(descriptor.Label)
	cDescriptor.label.data = cLabelStr0
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	defer C.free(unsafe.Pointer(cLabelStr0))
	if descriptor.Layout != nil {
		cDescriptor.layout = C.WGPUBindGroupLayout(unsafe.Pointer(descriptor.Layout.ref))
	}
	cDescriptor.entryCount = C.size_t(descriptor.EntryCount)
	return BindGroup{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateBindGroup(cDevice, &cDescriptor)))}
}

func (d *Device) CreateBindGroupLayout(descriptor BindGroupLayoutDescriptor) BindGroupLayout {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert descriptor to C.WGPUBindGroupLayoutDescriptor
	var cDescriptor C.WGPUBindGroupLayoutDescriptor
	cLabelStr0 := C.CString(descriptor.Label)
	cDescriptor.label.data = cLabelStr0
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	defer C.free(unsafe.Pointer(cLabelStr0))
	cDescriptor.entryCount = C.size_t(descriptor.EntryCount)
	return BindGroupLayout{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateBindGroupLayout(cDevice, &cDescriptor)))}
}

func (d *Device) CreateBuffer(descriptor BufferDescriptor) (*Buffer, error) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert descriptor to C.WGPUBufferDescriptor
	var cDescriptor C.WGPUBufferDescriptor
	cLabelStr0 := C.CString(descriptor.Label)
	cDescriptor.label.data = cLabelStr0
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	defer C.free(unsafe.Pointer(cLabelStr0))
	cDescriptor.mappedAtCreation = boolToWGPUBool(descriptor.MappedAtCreation)
	_ = C.wgpuDeviceCreateBuffer(cDevice, &cDescriptor) // TODO: Implement async/error logic
	return nil, nil
}

func (d *Device) CreateCommandEncoder(descriptor *CommandEncoderDescriptor) CommandEncoder {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert descriptor to C.WGPUCommandEncoderDescriptor
	var pDescriptor *C.WGPUCommandEncoderDescriptor
	if descriptor != nil {
		var cDescriptor C.WGPUCommandEncoderDescriptor
		cLabelStr0 := C.CString(descriptor.Label)
		cDescriptor.label.data = cLabelStr0
		cDescriptor.label.length = C.size_t(len(descriptor.Label))
		defer C.free(unsafe.Pointer(cLabelStr0))
		pDescriptor = &cDescriptor
	}
	return CommandEncoder{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateCommandEncoder(cDevice, pDescriptor)))}
}

func (d *Device) CreateComputePipeline(descriptor ComputePipelineDescriptor) ComputePipeline {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert descriptor to C.WGPUComputePipelineDescriptor
	var cDescriptor C.WGPUComputePipelineDescriptor
	cLabelStr0 := C.CString(descriptor.Label)
	cDescriptor.label.data = cLabelStr0
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	defer C.free(unsafe.Pointer(cLabelStr0))
	if descriptor.Layout != nil {
		cDescriptor.layout = C.WGPUPipelineLayout(unsafe.Pointer(descriptor.Layout.ref))
	}
	return ComputePipeline{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateComputePipeline(cDevice, &cDescriptor)))}
}

func (d *Device) CreateComputePipelineAsync(descriptor ComputePipelineDescriptor, callbackInfo CreateComputePipelineAsyncCallbackInfo) Future {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert descriptor to C.WGPUComputePipelineDescriptor
	var cDescriptor C.WGPUComputePipelineDescriptor
	cLabelStr0 := C.CString(descriptor.Label)
	cDescriptor.label.data = cLabelStr0
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	defer C.free(unsafe.Pointer(cLabelStr0))
	if descriptor.Layout != nil {
		cDescriptor.layout = C.WGPUPipelineLayout(unsafe.Pointer(descriptor.Layout.ref))
	}
	// Convert callbackInfo to C.WGPUCreateComputePipelineAsyncCallbackInfo
	var cCallbackInfo C.WGPUCreateComputePipelineAsyncCallbackInfo
	return Future{Id: uint64(C.wgpuDeviceCreateComputePipelineAsync(cDevice, &cDescriptor, cCallbackInfo).id)}
}

func (d *Device) CreateQuerySet(descriptor QuerySetDescriptor) QuerySet {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert descriptor to C.WGPUQuerySetDescriptor
	var cDescriptor C.WGPUQuerySetDescriptor
	cLabelStr0 := C.CString(descriptor.Label)
	cDescriptor.label.data = cLabelStr0
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	defer C.free(unsafe.Pointer(cLabelStr0))
	return QuerySet{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateQuerySet(cDevice, &cDescriptor)))}
}

func (d *Device) CreateRenderPipelineAsync(descriptor RenderPipelineDescriptor, callbackInfo CreateRenderPipelineAsyncCallbackInfo) Future {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert descriptor to C.WGPURenderPipelineDescriptor
	var cDescriptor C.WGPURenderPipelineDescriptor
	cLabelStr0 := C.CString(descriptor.Label)
	cDescriptor.label.data = cLabelStr0
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	defer C.free(unsafe.Pointer(cLabelStr0))
	if descriptor.Layout != nil {
		cDescriptor.layout = C.WGPUPipelineLayout(unsafe.Pointer(descriptor.Layout.ref))
	}
	// Convert callbackInfo to C.WGPUCreateRenderPipelineAsyncCallbackInfo
	var cCallbackInfo C.WGPUCreateRenderPipelineAsyncCallbackInfo
	return Future{Id: uint64(C.wgpuDeviceCreateRenderPipelineAsync(cDevice, &cDescriptor, cCallbackInfo).id)}
}

func (d *Device) CreateRenderBundleEncoder(descriptor RenderBundleEncoderDescriptor) RenderBundleEncoder {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert descriptor to C.WGPURenderBundleEncoderDescriptor
	var cDescriptor C.WGPURenderBundleEncoderDescriptor
	cLabelStr0 := C.CString(descriptor.Label)
	cDescriptor.label.data = cLabelStr0
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	defer C.free(unsafe.Pointer(cLabelStr0))
	cDescriptor.colorFormatCount = C.size_t(descriptor.ColorFormatCount)
	cDescriptor.depthReadOnly = boolToWGPUBool(descriptor.DepthReadOnly)
	cDescriptor.stencilReadOnly = boolToWGPUBool(descriptor.StencilReadOnly)
	return RenderBundleEncoder{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateRenderBundleEncoder(cDevice, &cDescriptor)))}
}

func (d *Device) CreateRenderPipeline(descriptor RenderPipelineDescriptor) RenderPipeline {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert descriptor to C.WGPURenderPipelineDescriptor
	var cDescriptor C.WGPURenderPipelineDescriptor
	cLabelStr0 := C.CString(descriptor.Label)
	cDescriptor.label.data = cLabelStr0
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	defer C.free(unsafe.Pointer(cLabelStr0))
	if descriptor.Layout != nil {
		cDescriptor.layout = C.WGPUPipelineLayout(unsafe.Pointer(descriptor.Layout.ref))
	}
	return RenderPipeline{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateRenderPipeline(cDevice, &cDescriptor)))}
}

func (d *Device) CreateSampler(descriptor *SamplerDescriptor) Sampler {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert descriptor to C.WGPUSamplerDescriptor
	var pDescriptor *C.WGPUSamplerDescriptor
	if descriptor != nil {
		var cDescriptor C.WGPUSamplerDescriptor
		cLabelStr0 := C.CString(descriptor.Label)
		cDescriptor.label.data = cLabelStr0
		cDescriptor.label.length = C.size_t(len(descriptor.Label))
		defer C.free(unsafe.Pointer(cLabelStr0))
		pDescriptor = &cDescriptor
	}
	return Sampler{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateSampler(cDevice, pDescriptor)))}
}

func (d *Device) CreateShaderModule(descriptor ShaderModuleDescriptor) ShaderModule {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert descriptor to C.WGPUShaderModuleDescriptor
	var cDescriptor C.WGPUShaderModuleDescriptor
	cLabelStr0 := C.CString(descriptor.Label)
	cDescriptor.label.data = cLabelStr0
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	defer C.free(unsafe.Pointer(cLabelStr0))
	return ShaderModule{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateShaderModule(cDevice, &cDescriptor)))}
}

func (d *Device) CreateTexture(descriptor TextureDescriptor) Texture {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert descriptor to C.WGPUTextureDescriptor
	var cDescriptor C.WGPUTextureDescriptor
	cLabelStr0 := C.CString(descriptor.Label)
	cDescriptor.label.data = cLabelStr0
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	defer C.free(unsafe.Pointer(cLabelStr0))
	cDescriptor.viewFormatCount = C.size_t(descriptor.ViewFormatCount)
	return Texture{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateTexture(cDevice, &cDescriptor)))}
}

func (d *Device) Destroy() {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	C.wgpuDeviceDestroy(cDevice)
}

func (d *Device) GetLimits(limits Limits) Status {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert limits to C.WGPULimits
	var cLimits C.WGPULimits
	return Status(C.wgpuDeviceGetLimits(cDevice, &cLimits))
}

func (d *Device) GetLostFuture() Future {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	return Future{Id: uint64(C.wgpuDeviceGetLostFuture(cDevice).id)}
}

func (d *Device) HasFeature(feature FeatureName) bool {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	cFeature := C.WGPUFeatureName(feature)
	return bool(C.wgpuDeviceHasFeature(cDevice, cFeature) != 0)
}

func (d *Device) GetFeatures(features SupportedFeatures) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert features to C.WGPUSupportedFeatures
	var cFeatures C.WGPUSupportedFeatures
	cFeatures.featureCount = C.size_t(features.FeatureCount)
	C.wgpuDeviceGetFeatures(cDevice, &cFeatures)
}

func (d *Device) GetAdapterInfo(adapterInfo AdapterInfo) Status {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert adapterInfo to C.WGPUAdapterInfo
	var cAdapterInfo C.WGPUAdapterInfo
	cVendorStr0 := C.CString(adapterInfo.Vendor)
	cAdapterInfo.vendor.data = cVendorStr0
	cAdapterInfo.vendor.length = C.size_t(len(adapterInfo.Vendor))
	defer C.free(unsafe.Pointer(cVendorStr0))
	cArchitectureStr1 := C.CString(adapterInfo.Architecture)
	cAdapterInfo.architecture.data = cArchitectureStr1
	cAdapterInfo.architecture.length = C.size_t(len(adapterInfo.Architecture))
	defer C.free(unsafe.Pointer(cArchitectureStr1))
	cDeviceStr2 := C.CString(adapterInfo.Device)
	cAdapterInfo.device.data = cDeviceStr2
	cAdapterInfo.device.length = C.size_t(len(adapterInfo.Device))
	defer C.free(unsafe.Pointer(cDeviceStr2))
	cDescriptionStr3 := C.CString(adapterInfo.Description)
	cAdapterInfo.description.data = cDescriptionStr3
	cAdapterInfo.description.length = C.size_t(len(adapterInfo.Description))
	defer C.free(unsafe.Pointer(cDescriptionStr3))
	return Status(C.wgpuDeviceGetAdapterInfo(cDevice, &cAdapterInfo))
}

func (d *Device) GetQueue() Queue {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	return Queue{ref: uintptr(unsafe.Pointer(C.wgpuDeviceGetQueue(cDevice)))}
}

func (d *Device) PushErrorScope(filter ErrorFilter) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	cFilter := C.WGPUErrorFilter(filter)
	C.wgpuDevicePushErrorScope(cDevice, cFilter)
}

func (d *Device) PopErrorScope(callbackInfo PopErrorScopeCallbackInfo) Future {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert callbackInfo to C.WGPUPopErrorScopeCallbackInfo
	var cCallbackInfo C.WGPUPopErrorScopeCallbackInfo
	return Future{Id: uint64(C.wgpuDevicePopErrorScope(cDevice, cCallbackInfo).id)}
}

func (d *Device) SetLabel(label string) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuDeviceSetLabel(cDevice, cLabel)
}

type QuerySet struct {
	ref uintptr
}

func (q *QuerySet) SetLabel(label string) {
	cQuerySet := C.WGPUQuerySet(unsafe.Pointer(q.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuQuerySetSetLabel(cQuerySet, cLabel)
}

func (q *QuerySet) GetType() QueryType {
	cQuerySet := C.WGPUQuerySet(unsafe.Pointer(q.ref))
	return QueryType(C.wgpuQuerySetGetType(cQuerySet))
}

func (q *QuerySet) GetCount() uint32 {
	cQuerySet := C.WGPUQuerySet(unsafe.Pointer(q.ref))
	return uint32(C.wgpuQuerySetGetCount(cQuerySet))
}

func (q *QuerySet) Destroy() {
	cQuerySet := C.WGPUQuerySet(unsafe.Pointer(q.ref))
	C.wgpuQuerySetDestroy(cQuerySet)
}

type ExternalTexture struct {
	ref uintptr
}

func (e *ExternalTexture) SetLabel(label string) {
	cExternalTexture := C.WGPUExternalTexture(unsafe.Pointer(e.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuExternalTextureSetLabel(cExternalTexture, cLabel)
}

type CommandBuffer struct {
	ref uintptr
}

func (c *CommandBuffer) SetLabel(label string) {
	cCommandBuffer := C.WGPUCommandBuffer(unsafe.Pointer(c.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuCommandBufferSetLabel(cCommandBuffer, cLabel)
}

type ShaderModule struct {
	ref uintptr
}

func (s *ShaderModule) GetCompilationInfo(callbackInfo CompilationInfoCallbackInfo) Future {
	cShaderModule := C.WGPUShaderModule(unsafe.Pointer(s.ref))
	// Convert callbackInfo to C.WGPUCompilationInfoCallbackInfo
	var cCallbackInfo C.WGPUCompilationInfoCallbackInfo
	return Future{Id: uint64(C.wgpuShaderModuleGetCompilationInfo(cShaderModule, cCallbackInfo).id)}
}

func (s *ShaderModule) SetLabel(label string) {
	cShaderModule := C.WGPUShaderModule(unsafe.Pointer(s.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuShaderModuleSetLabel(cShaderModule, cLabel)
}

type Adapter struct {
	ref uintptr
}

func (a *Adapter) GetLimits(limits Limits) Status {
	cAdapter := C.WGPUAdapter(unsafe.Pointer(a.ref))
	// Convert limits to C.WGPULimits
	var cLimits C.WGPULimits
	return Status(C.wgpuAdapterGetLimits(cAdapter, &cLimits))
}

func (a *Adapter) GetInfo(info AdapterInfo) Status {
	cAdapter := C.WGPUAdapter(unsafe.Pointer(a.ref))
	// Convert info to C.WGPUAdapterInfo
	var cInfo C.WGPUAdapterInfo
	cVendorStr0 := C.CString(info.Vendor)
	cInfo.vendor.data = cVendorStr0
	cInfo.vendor.length = C.size_t(len(info.Vendor))
	defer C.free(unsafe.Pointer(cVendorStr0))
	cArchitectureStr1 := C.CString(info.Architecture)
	cInfo.architecture.data = cArchitectureStr1
	cInfo.architecture.length = C.size_t(len(info.Architecture))
	defer C.free(unsafe.Pointer(cArchitectureStr1))
	cDeviceStr2 := C.CString(info.Device)
	cInfo.device.data = cDeviceStr2
	cInfo.device.length = C.size_t(len(info.Device))
	defer C.free(unsafe.Pointer(cDeviceStr2))
	cDescriptionStr3 := C.CString(info.Description)
	cInfo.description.data = cDescriptionStr3
	cInfo.description.length = C.size_t(len(info.Description))
	defer C.free(unsafe.Pointer(cDescriptionStr3))
	return Status(C.wgpuAdapterGetInfo(cAdapter, &cInfo))
}

func (a *Adapter) HasFeature(feature FeatureName) bool {
	cAdapter := C.WGPUAdapter(unsafe.Pointer(a.ref))
	cFeature := C.WGPUFeatureName(feature)
	return bool(C.wgpuAdapterHasFeature(cAdapter, cFeature) != 0)
}

func (a *Adapter) GetFeatures(features SupportedFeatures) {
	cAdapter := C.WGPUAdapter(unsafe.Pointer(a.ref))
	// Convert features to C.WGPUSupportedFeatures
	var cFeatures C.WGPUSupportedFeatures
	cFeatures.featureCount = C.size_t(features.FeatureCount)
	C.wgpuAdapterGetFeatures(cAdapter, &cFeatures)
}

func (a *Adapter) RequestDevice(descriptor *DeviceDescriptor, callbackInfo RequestDeviceCallbackInfo) Future {
	cAdapter := C.WGPUAdapter(unsafe.Pointer(a.ref))
	// Convert descriptor to C.WGPUDeviceDescriptor
	var pDescriptor *C.WGPUDeviceDescriptor
	if descriptor != nil {
		var cDescriptor C.WGPUDeviceDescriptor
		cLabelStr0 := C.CString(descriptor.Label)
		cDescriptor.label.data = cLabelStr0
		cDescriptor.label.length = C.size_t(len(descriptor.Label))
		defer C.free(unsafe.Pointer(cLabelStr0))
		cDescriptor.requiredFeatureCount = C.size_t(descriptor.RequiredFeatureCount)
		pDescriptor = &cDescriptor
	}
	// Convert callbackInfo to C.WGPURequestDeviceCallbackInfo
	var cCallbackInfo C.WGPURequestDeviceCallbackInfo
	return Future{Id: uint64(C.wgpuAdapterRequestDevice(cAdapter, pDescriptor, cCallbackInfo).id)}
}

type ComputePipeline struct {
	ref uintptr
}

func (c *ComputePipeline) GetBindGroupLayout(groupIndex uint32) BindGroupLayout {
	cComputePipeline := C.WGPUComputePipeline(unsafe.Pointer(c.ref))
	cGroupIndex := C.uint32_t(groupIndex)
	return BindGroupLayout{ref: uintptr(unsafe.Pointer(C.wgpuComputePipelineGetBindGroupLayout(cComputePipeline, cGroupIndex)))}
}

func (c *ComputePipeline) SetLabel(label string) {
	cComputePipeline := C.WGPUComputePipeline(unsafe.Pointer(c.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuComputePipelineSetLabel(cComputePipeline, cLabel)
}

type PipelineLayout struct {
	ref uintptr
}

func (p *PipelineLayout) SetLabel(label string) {
	cPipelineLayout := C.WGPUPipelineLayout(unsafe.Pointer(p.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuPipelineLayoutSetLabel(cPipelineLayout, cLabel)
}

type RenderPassEncoder struct {
	ref uintptr
}

func (r *RenderPassEncoder) SetPipeline(pipeline RenderPipeline) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	cPipeline := C.WGPURenderPipeline(unsafe.Pointer(pipeline.ref))
	C.wgpuRenderPassEncoderSetPipeline(cRenderPassEncoder, cPipeline)
}

func (r *RenderPassEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsetCount int, dynamicOffsets uint32) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	cGroupIndex := C.uint32_t(groupIndex)
	pGroup := C.WGPUBindGroup(unsafe.Pointer(group.ref))
	cDynamicOffsetCount := C.size_t(dynamicOffsetCount)
	cDynamicOffsets := C.uint32_t(dynamicOffsets)
	C.wgpuRenderPassEncoderSetBindGroup(cRenderPassEncoder, cGroupIndex, pGroup, cDynamicOffsetCount, &cDynamicOffsets)
}

func (r *RenderPassEncoder) Draw(vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	cVertexCount := C.uint32_t(vertexCount)
	cInstanceCount := C.uint32_t(instanceCount)
	cFirstVertex := C.uint32_t(firstVertex)
	cFirstInstance := C.uint32_t(firstInstance)
	C.wgpuRenderPassEncoderDraw(cRenderPassEncoder, cVertexCount, cInstanceCount, cFirstVertex, cFirstInstance)
}

func (r *RenderPassEncoder) DrawIndexed(indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	cIndexCount := C.uint32_t(indexCount)
	cInstanceCount := C.uint32_t(instanceCount)
	cFirstIndex := C.uint32_t(firstIndex)
	cBaseVertex := C.int32_t(baseVertex)
	cFirstInstance := C.uint32_t(firstInstance)
	C.wgpuRenderPassEncoderDrawIndexed(cRenderPassEncoder, cIndexCount, cInstanceCount, cFirstIndex, cBaseVertex, cFirstInstance)
}

func (r *RenderPassEncoder) DrawIndirect(indirectBuffer Buffer, indirectOffset uint64) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	cIndirectBuffer := C.WGPUBuffer(unsafe.Pointer(indirectBuffer.ref))
	cIndirectOffset := C.uint64_t(indirectOffset)
	C.wgpuRenderPassEncoderDrawIndirect(cRenderPassEncoder, cIndirectBuffer, cIndirectOffset)
}

func (r *RenderPassEncoder) DrawIndexedIndirect(indirectBuffer Buffer, indirectOffset uint64) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	cIndirectBuffer := C.WGPUBuffer(unsafe.Pointer(indirectBuffer.ref))
	cIndirectOffset := C.uint64_t(indirectOffset)
	C.wgpuRenderPassEncoderDrawIndexedIndirect(cRenderPassEncoder, cIndirectBuffer, cIndirectOffset)
}

func (r *RenderPassEncoder) ExecuteBundles(bundleCount int, bundles RenderBundle) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	cBundleCount := C.size_t(bundleCount)
	cBundles := C.WGPURenderBundle(unsafe.Pointer(bundles.ref))
	C.wgpuRenderPassEncoderExecuteBundles(cRenderPassEncoder, cBundleCount, &cBundles)
}

func (r *RenderPassEncoder) InsertDebugMarker(markerLabel string) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	// Convert markerLabel to C.WGPUStringView
	var cMarkerLabel C.WGPUStringView
	cMarkerLabelStr := C.CString(markerLabel)
	cMarkerLabel.data = cMarkerLabelStr
	cMarkerLabel.length = C.size_t(len(markerLabel))
	defer C.free(unsafe.Pointer(cMarkerLabelStr))
	C.wgpuRenderPassEncoderInsertDebugMarker(cRenderPassEncoder, cMarkerLabel)
}

func (r *RenderPassEncoder) PopDebugGroup() {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	C.wgpuRenderPassEncoderPopDebugGroup(cRenderPassEncoder)
}

func (r *RenderPassEncoder) PushDebugGroup(groupLabel string) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	// Convert groupLabel to C.WGPUStringView
	var cGroupLabel C.WGPUStringView
	cGroupLabelStr := C.CString(groupLabel)
	cGroupLabel.data = cGroupLabelStr
	cGroupLabel.length = C.size_t(len(groupLabel))
	defer C.free(unsafe.Pointer(cGroupLabelStr))
	C.wgpuRenderPassEncoderPushDebugGroup(cRenderPassEncoder, cGroupLabel)
}

func (r *RenderPassEncoder) SetStencilReference(reference uint32) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	cReference := C.uint32_t(reference)
	C.wgpuRenderPassEncoderSetStencilReference(cRenderPassEncoder, cReference)
}

func (r *RenderPassEncoder) SetBlendConstant(color Color) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	// Convert color to C.WGPUColor
	var cColor C.WGPUColor
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
	C.wgpuRenderPassEncoderSetViewport(cRenderPassEncoder, cX, cY, cWidth, cHeight, cMinDepth, cMaxDepth)
}

func (r *RenderPassEncoder) SetScissorRect(x uint32, y uint32, width uint32, height uint32) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	cX := C.uint32_t(x)
	cY := C.uint32_t(y)
	cWidth := C.uint32_t(width)
	cHeight := C.uint32_t(height)
	C.wgpuRenderPassEncoderSetScissorRect(cRenderPassEncoder, cX, cY, cWidth, cHeight)
}

func (r *RenderPassEncoder) SetVertexBuffer(slot uint32, buffer *Buffer, offset uint64, size uint64) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	cSlot := C.uint32_t(slot)
	pBuffer := C.WGPUBuffer(unsafe.Pointer(buffer.ref))
	cOffset := C.uint64_t(offset)
	cSize := C.uint64_t(size)
	C.wgpuRenderPassEncoderSetVertexBuffer(cRenderPassEncoder, cSlot, pBuffer, cOffset, cSize)
}

func (r *RenderPassEncoder) SetIndexBuffer(buffer Buffer, format IndexFormat, offset uint64, size uint64) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	cBuffer := C.WGPUBuffer(unsafe.Pointer(buffer.ref))
	cFormat := C.WGPUIndexFormat(format)
	cOffset := C.uint64_t(offset)
	cSize := C.uint64_t(size)
	C.wgpuRenderPassEncoderSetIndexBuffer(cRenderPassEncoder, cBuffer, cFormat, cOffset, cSize)
}

func (r *RenderPassEncoder) BeginOcclusionQuery(queryIndex uint32) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	cQueryIndex := C.uint32_t(queryIndex)
	C.wgpuRenderPassEncoderBeginOcclusionQuery(cRenderPassEncoder, cQueryIndex)
}

func (r *RenderPassEncoder) EndOcclusionQuery() {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	C.wgpuRenderPassEncoderEndOcclusionQuery(cRenderPassEncoder)
}

func (r *RenderPassEncoder) End() {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	C.wgpuRenderPassEncoderEnd(cRenderPassEncoder)
}

func (r *RenderPassEncoder) SetLabel(label string) {
	cRenderPassEncoder := C.WGPURenderPassEncoder(unsafe.Pointer(r.ref))
	// Convert label to C.WGPUStringView
	var cLabel C.WGPUStringView
	cLabelStr := C.CString(label)
	cLabel.data = cLabelStr
	cLabel.length = C.size_t(len(label))
	defer C.free(unsafe.Pointer(cLabelStr))
	C.wgpuRenderPassEncoderSetLabel(cRenderPassEncoder, cLabel)
}

func GetInstanceLimits(limits InstanceLimits) Status {
	// Convert limits to C.WGPUInstanceLimits
	var cLimits C.WGPUInstanceLimits
	cLimits.timedWaitAnyMaxCount = C.size_t(limits.TimedWaitAnyMaxCount)
	return Status(C.wgpuGetInstanceLimits(&cLimits))
}

func CreateInstance(descriptor *InstanceDescriptor) Instance {
	// Convert descriptor to C.WGPUInstanceDescriptor
	var pDescriptor *C.WGPUInstanceDescriptor
	if descriptor != nil {
		var cDescriptor C.WGPUInstanceDescriptor
		cDescriptor.requiredFeatureCount = C.size_t(descriptor.RequiredFeatureCount)
		pDescriptor = &cDescriptor
	}
	return Instance{ref: uintptr(unsafe.Pointer(C.wgpuCreateInstance(pDescriptor)))}
}

func GetInstanceFeatures(features SupportedInstanceFeatures) {
	// Convert features to C.WGPUSupportedInstanceFeatures
	var cFeatures C.WGPUSupportedInstanceFeatures
	cFeatures.featureCount = C.size_t(features.FeatureCount)
	C.wgpuGetInstanceFeatures(&cFeatures)
}

func HasInstanceFeature(feature InstanceFeatureName) bool {
	cFeature := C.WGPUInstanceFeatureName(feature)
	return bool(C.wgpuHasInstanceFeature(cFeature) != 0)
}
