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

// Buffer represents a GPU buffer, which is a region of memory that can be used to store data.
// Buffers are created from a device and can be mapped for reading or writing.
type Buffer struct {
	ref C.WGPUBuffer
}

// AsImageCopyBuffer returns a TexelCopyBufferInfo for use in image copy operations.
// The bytesPerRow and rowsPerImage parameters define the layout of the data.
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

	go fn(
		MapAsyncStatus(status),
		msg,
	)
}

// MapAsync maps the buffer for reading or writing asynchronously.
// The mode specifies whether to map for reading or writing.
// The offset and size specify the range of the buffer to map.
// The callback is called when the mapping is complete.
// Returns a Future that can be used to wait for the mapping to complete.
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
// GetMappedRange returns a slice of bytes representing the mapped range of the buffer for writing.
// The buffer must have been mapped with MapModeWrite.
func (b *Buffer) GetMappedRange(offset int, size int) []byte {
	buf := C.wgpuBufferGetMappedRange(b.ref, C.size_t(offset), C.size_t(size))
	return unsafe.Slice((*byte)(buf), size)
}

// MapModeRead
// GetConstMappedRange returns a slice of bytes representing the mapped range of the buffer for reading.
// The buffer must have been mapped with MapModeRead.
func (b *Buffer) GetConstMappedRange(offset int, size int) []byte {
	buf := C.wgpuBufferGetConstMappedRange(b.ref, C.size_t(offset), C.size_t(size))
	return unsafe.Slice((*byte)(buf), size)
}

// SetLabel sets the debug label for the buffer.
// This label appears in debuggers and validation layers.
func (b *Buffer) SetLabel(label string) {
	C.wgpuBufferSetLabel(b.ref, toCStr(label))
}

// GetUsage returns the usage flags for the buffer.
func (b *Buffer) GetUsage() BufferUsage {
	return BufferUsage(C.wgpuBufferGetUsage(b.ref))
}

// GetSize returns the size of the buffer in bytes.
func (b *Buffer) GetSize() uint64 {
	return uint64(C.wgpuBufferGetSize(b.ref))
}

// GetMapState returns the current mapping state of the buffer.
func (b *Buffer) GetMapState() BufferMapState {
	return BufferMapState(C.wgpuBufferGetMapState(b.ref))
}

// Unmap unmaps the buffer, flushing any writes if it was mapped for writing.
// After unmapping, the mapped range is no longer valid.
func (b *Buffer) Unmap() {
	C.wgpuBufferUnmap(b.ref)
}

// Release releases the buffer and all associated resources.
// After calling this method, the buffer should no longer be used.
func (b *Buffer) Release() {
	C.wgpuBufferRelease(b.ref)
}

// Destroy destroys the buffer and frees all associated GPU resources.
func (b *Buffer) Destroy() {
	C.wgpuBufferDestroy(b.ref)
}
