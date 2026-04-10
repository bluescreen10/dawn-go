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
func (b *Buffer) MapAsync(mode MapMode, offset int, size int, callback BufferMapCallback) {

	cMode := C.WGPUMapMode(mode)
	cOffset := C.size_t(offset)
	cSize := C.size_t(size)

	handle := cgo.NewHandle(callback)

	cCallbackInfo := C.WGPUBufferMapCallbackInfo{
		mode:      C.WGPUCallbackMode(callbackModeAllowProcessEvents),
		callback:  C.WGPUBufferMapCallback(C.cgo_callback_BufferMapCallback),
		userdata1: unsafe.Pointer(handle),
	}

	C.wgpuBufferMapAsync(b.ref, cMode, cOffset, cSize, cCallbackInfo)
}

// MapModeWrite
func (b *Buffer) GetMappedRange(offset int, size int) []byte {
	cOffset := C.size_t(offset)
	cSize := C.size_t(size)
	buf := C.wgpuBufferGetMappedRange(b.ref, cOffset, cSize)
	return unsafe.Slice((*byte)(buf), size)
}

// MapModeRead
func (b *Buffer) GetConstMappedRange(offset int, size int) []byte {
	cOffset := C.size_t(offset)
	cSize := C.size_t(size)
	buf := C.wgpuBufferGetConstMappedRange(b.ref, cOffset, cSize)
	return unsafe.Slice((*byte)(buf), size)
}

func (b *Buffer) SetLabel(label string) {
	cLabel := toCStr(label)
	C.wgpuBufferSetLabel(b.ref, cLabel)
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
