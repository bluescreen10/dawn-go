//go:build !js

package wgpu

/*
#include "webgpu.h"

extern void cgo_callback_BufferMapCallback(WGPUMapAsyncStatus status, WGPUStringView message, void *userData1, void *userData2);
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

type Buffer struct {
	ref C.WGPUBuffer
}

func (b *Buffer) AsImageCopyBuffer(bytesPerRow, rowsPerImage uint32) TexelCopyBufferInfo {
	return TexelCopyBufferInfo{
		Layout: TexelCopyBufferLayout{
			Offset:       0,
			BytesPerRow:  bytesPerRow,
			RowsPerImage: rowsPerImage,
		},
		Buffer: b,
	}
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
func (b *Buffer) MapAsync(mode MapMode, offset int, size int, callback BufferMapCallback) Future {
	handle := cgo.NewHandle(callback)

	cCallbackInfo := C.WGPUBufferMapCallbackInfo{
		mode:      C.WGPUCallbackMode(callbackModeAllowSpontaneous),
		callback:  C.WGPUBufferMapCallback(C.cgo_callback_BufferMapCallback),
		userdata1: unsafe.Pointer(handle),
	}

	future := C.wgpuBufferMapAsync(b.ref, C.WGPUMapMode(mode), C.size_t(offset), C.size_t(size), cCallbackInfo)
	return Future{id: uint64(future.id)}
}

// MapModeWrite
func (b *Buffer) GetMappedRange(offset int, size int) []byte {
	buf := C.wgpuBufferGetMappedRange(b.ref, C.size_t(offset), C.size_t(size))
	return unsafe.Slice((*byte)(buf), size)
}

// MapModeRead
func (b *Buffer) GetConstMappedRange(offset int, size int) []byte {
	buf := C.wgpuBufferGetConstMappedRange(b.ref, C.size_t(offset), C.size_t(size))
	return unsafe.Slice((*byte)(buf), size)
}

func (b *Buffer) SetLabel(label string) {
	C.wgpuBufferSetLabel(b.ref, toCStr(label))
}

func (b *Buffer) GetUsage() BufferUsage {
	return BufferUsage(C.wgpuBufferGetUsage(b.ref))
}

func (b *Buffer) GetSize() uint64 {
	return uint64(C.wgpuBufferGetSize(b.ref))
}

func (b *Buffer) GetMapState() BufferMapState {
	return BufferMapState(C.wgpuBufferGetMapState(b.ref))
}

func (b *Buffer) Unmap() {
	C.wgpuBufferUnmap(b.ref)
}

func (b *Buffer) Destroy() {
	C.wgpuBufferDestroy(b.ref)
}
