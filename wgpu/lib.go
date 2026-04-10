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

extern void cgo_callback_CompilationInfoCallback(WGPUCompilationInfoRequestStatus status, WGPUCompilationInfo compilationInfo, void *userData1, void *userData2);
*/
import "C"

import (
	"fmt"
	"runtime/cgo"
	"unsafe"
)

type ComputePassEncoder struct {
	ref C.WGPUComputePassEncoder
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
	ref C.WGPURenderPassEncoder
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
	pBuffer := buffer.ref
	cFormat := C.WGPUIndexFormat(format)
	cOffset := C.uint64_t(offset)
	cSize := C.uint64_t(size)
	// Call and return
	C.wgpuRenderPassEncoderSetIndexBuffer(r.ref, pBuffer, cFormat, cOffset, cSize)
}

func (r *RenderPassEncoder) BeginOcclusionQuery(queryIndex uint32) {
	cQueryIndex := C.uint32_t(queryIndex)
	C.wgpuRenderPassEncoderBeginOcclusionQuery(r.ref, cQueryIndex)
}

func (r *RenderPassEncoder) EndOcclusionQuery() {
	C.wgpuRenderPassEncoderEndOcclusionQuery(r.ref)
}

func (r *RenderPassEncoder) End() {
	C.wgpuRenderPassEncoderEnd(r.ref)
}

func (r *RenderPassEncoder) SetLabel(label string) {
	cLabel := toCStr(label)
	C.wgpuRenderPassEncoderSetLabel(r.ref, cLabel)
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
