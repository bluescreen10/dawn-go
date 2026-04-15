//go:build !js

package wgpu

/*
#include "webgpu.h"

extern void cgo_callback_QueueWorkDoneCallback(WGPUQueueWorkDoneStatus status, WGPUStringView message, void *userData1, void *userData2);
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

// Queue represents a command queue that is used to submit commands to the GPU.
// Queues are obtained from a device and are used to execute command buffers.
type Queue struct {
	ref C.WGPUQueue
}

// Submit submits command buffers to the queue for execution.
// The command buffers contain recorded commands that will be executed on the GPU.
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

	go fn(
		QueueWorkDoneStatus(status),
		msg,
	)
}

// OnSubmittedWorkDone registers a callback that is called when all previously submitted commands complete.
// The callback is called when the GPU has finished executing all commands submitted up to this point.
func (q *Queue) OnSubmittedWorkDone(callback QueueWorkDoneCallback) {
	handle := cgo.NewHandle(callback)

	cCallbackInfo := C.WGPUQueueWorkDoneCallbackInfo{
		mode:      C.WGPUCallbackMode(callbackModeAllowSpontaneous),
		callback:  C.WGPUQueueWorkDoneCallback(C.cgo_callback_QueueWorkDoneCallback),
		userdata1: unsafe.Pointer(handle),
	}

	C.wgpuQueueOnSubmittedWorkDone(q.ref, cCallbackInfo)
}

// WriteBuffer writes data from the CPU to a buffer on the GPU.
// The offset specifies where to start writing in the buffer.
func (q *Queue) WriteBuffer(buffer *Buffer, offset uint64, data []byte) {
	cSize := C.size_t(len(data))

	var cData unsafe.Pointer
	if cSize > 0 {
		cData = unsafe.Pointer(&data[0])
	}

	C.wgpuQueueWriteBuffer(q.ref, buffer.ref, C.uint64_t(offset), cData, cSize)
}

// WriteTexture writes data from the CPU to a texture on the GPU.
// The destination defines the texture region to write to, and the dataLayout defines the layout of the source data.
func (q *Queue) WriteTexture(destination TexelCopyTextureInfo, data []byte, dataLayout TexelCopyBufferLayout, writeSize Extent3D) {
	cDestination := C.WGPUTexelCopyTextureInfo{
		texture:  C.WGPUTexture(destination.Texture.ref),
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

// SetLabel sets the debug label for the queue.
// This label appears in debuggers and validation layers.
func (q *Queue) SetLabel(label string) {
	C.wgpuQueueSetLabel(q.ref, toCStr(label))
}

// Release releases the queue and all associated resources.
// After calling this method, the queue should no longer be used.
func (q *Queue) Release() {
	C.wgpuQueueRelease(q.ref)
}
