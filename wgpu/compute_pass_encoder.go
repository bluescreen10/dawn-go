//go:build !js

package wgpu

/*
#include <stdlib.h>
#include "webgpu.h"
*/
import "C"
import "unsafe"

type ComputePassEncoder struct {
	ref C.WGPUComputePassEncoder
}

func (c *ComputePassEncoder) InsertDebugMarker(markerLabel string) {
	cMarkerLabel := toCStr(markerLabel)
	C.wgpuComputePassEncoderInsertDebugMarker(c.ref, cMarkerLabel)
}

func (c *ComputePassEncoder) PopDebugGroup() {
	C.wgpuComputePassEncoderPopDebugGroup(c.ref)
}

func (c *ComputePassEncoder) PushDebugGroup(groupLabel string) {
	cGroupLabel := toCStr(groupLabel)
	C.wgpuComputePassEncoderPushDebugGroup(c.ref, cGroupLabel)
}

func (c *ComputePassEncoder) SetPipeline(pipeline *ComputePipeline) {
	C.wgpuComputePassEncoderSetPipeline(c.ref, pipeline.ref)
}

func (c *ComputePassEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsets []uint32) {
	cGroupIndex := C.uint32_t(groupIndex)
	cDynamicOffsetCount := C.size_t(len(dynamicOffsets))
	var cDynamicOffsets *C.uint32_t

	if cDynamicOffsetCount > 0 {
		cDynamicOffsets = (*C.uint32_t)(unsafe.Pointer(&dynamicOffsets[0]))
	}
	C.wgpuComputePassEncoderSetBindGroup(c.ref, cGroupIndex, group.ref, cDynamicOffsetCount, cDynamicOffsets)
}

func (c *ComputePassEncoder) DispatchWorkgroups(workgroupCountX uint32, workgroupCountY uint32, workgroupCountZ uint32) {
	cWorkgroupCountX := C.uint32_t(workgroupCountX)
	cWorkgroupCountY := C.uint32_t(workgroupCountY)
	cWorkgroupCountZ := C.uint32_t(workgroupCountZ)
	C.wgpuComputePassEncoderDispatchWorkgroups(c.ref, cWorkgroupCountX, cWorkgroupCountY, cWorkgroupCountZ)
}

func (c *ComputePassEncoder) DispatchWorkgroupsIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	cIndirectOffset := C.uint64_t(indirectOffset)
	C.wgpuComputePassEncoderDispatchWorkgroupsIndirect(c.ref, indirectBuffer.ref, cIndirectOffset)
}

func (c *ComputePassEncoder) End() {
	C.wgpuComputePassEncoderEnd(c.ref)
}

func (c *ComputePassEncoder) SetLabel(label string) {
	cLabel := toCStr(label)
	C.wgpuComputePassEncoderSetLabel(c.ref, cLabel)
}
