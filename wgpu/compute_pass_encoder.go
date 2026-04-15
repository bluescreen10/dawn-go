//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"
import "unsafe"

type ComputePassEncoder struct {
	ref C.WGPUComputePassEncoder
}

func (c *ComputePassEncoder) InsertDebugMarker(markerLabel string) {
	C.wgpuComputePassEncoderInsertDebugMarker(c.ref, toCStr(markerLabel))
}

func (c *ComputePassEncoder) PopDebugGroup() {
	C.wgpuComputePassEncoderPopDebugGroup(c.ref)
}

func (c *ComputePassEncoder) PushDebugGroup(groupLabel string) {
	C.wgpuComputePassEncoderPushDebugGroup(c.ref, toCStr(groupLabel))
}

func (c *ComputePassEncoder) SetPipeline(pipeline *ComputePipeline) {
	C.wgpuComputePassEncoderSetPipeline(c.ref, pipeline.ref)
}

func (c *ComputePassEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsets []uint32) {
	cDynamicOffsetCount := C.size_t(len(dynamicOffsets))
	var cDynamicOffsets *C.uint32_t

	if cDynamicOffsetCount > 0 {
		cDynamicOffsets = (*C.uint32_t)(unsafe.Pointer(&dynamicOffsets[0]))
	}
	C.wgpuComputePassEncoderSetBindGroup(c.ref, C.uint32_t(groupIndex), group.ref, cDynamicOffsetCount, cDynamicOffsets)
}

func (c *ComputePassEncoder) DispatchWorkgroups(workgroupCountX uint32, workgroupCountY uint32, workgroupCountZ uint32) {
	C.wgpuComputePassEncoderDispatchWorkgroups(c.ref, C.uint32_t(workgroupCountX), C.uint32_t(workgroupCountY), C.uint32_t(workgroupCountZ))
}

func (c *ComputePassEncoder) DispatchWorkgroupsIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	C.wgpuComputePassEncoderDispatchWorkgroupsIndirect(c.ref, indirectBuffer.ref, C.uint64_t(indirectOffset))
}

func (c *ComputePassEncoder) End() {
	C.wgpuComputePassEncoderEnd(c.ref)
}

func (c *ComputePassEncoder) SetLabel(label string) {
	C.wgpuComputePassEncoderSetLabel(c.ref, toCStr(label))
}

func (c *ComputePassEncoder) Release() {
	C.wgpuComputePassEncoderRelease(c.ref)
}
