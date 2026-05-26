//go:build js

package wgpu

import "syscall/js"

// Queue represents a command queue that is used to submit commands to the GPU.
// Queues are obtained from a device and are used to execute command buffers.
type Queue struct {
	ref js.Value
}

// Submit submits command buffers to the queue for execution.
// The command buffers contain recorded commands that will be executed on the GPU.
func (q *Queue) Submit(commands ...*CommandBuffer) {
	cmds := make([]any, len(commands))
	for i, c := range commands {
		cmds[i] = c.ref
	}
	q.ref.Call("submit", cmds)
}

// OnSubmittedWorkDone registers a callback that is called when all previously submitted commands complete.
// The callback is called when the GPU has finished executing all commands submitted up to this point.
func (q *Queue) OnSubmittedWorkDone(callback QueueWorkDoneCallback) {
	promise := q.ref.Call("onSubmittedWorkDone")
	go func() {
		_, err := awaitPromise(promise)
		if err != nil {
			callback(QueueWorkDoneStatusError, err.Error())
			return
		}
		callback(QueueWorkDoneStatusSuccess, "")
	}()
}

// WriteBuffer writes data from the CPU to a buffer on the GPU.
// The offset specifies where to start writing in the buffer.
func (q *Queue) WriteBuffer(buffer *Buffer, offset uint64, data []byte) {
	if len(data) == 0 {
		return
	}
	uint8Array := js.Global().Get("Uint8Array").New(len(data))
	js.CopyBytesToJS(uint8Array, data)
	q.ref.Call("writeBuffer", buffer.ref, offset, uint8Array)
}

// WriteTexture writes data from the CPU to a texture on the GPU.
// The destination defines the texture region to write to, and the dataLayout defines the layout of the source data.
func (q *Queue) WriteTexture(destination TexelCopyTextureInfo, data []byte, dataLayout TexelCopyBufferLayout, writeSize Extent3D) {
	uint8Array := js.Global().Get("Uint8Array").New(len(data))
	js.CopyBytesToJS(uint8Array, data)

	dest := map[string]any{
		"texture":  destination.Texture.ref,
		"mipLevel": destination.MipLevel,
		"origin": map[string]any{
			"x": destination.Origin.X,
			"y": destination.Origin.Y,
			"z": destination.Origin.Z,
		},
		"aspect": destination.Aspect.toJS(),
	}
	layout := map[string]any{
		"offset":       dataLayout.Offset,
		"bytesPerRow":  dataLayout.BytesPerRow,
		"rowsPerImage": dataLayout.RowsPerImage,
	}
	size := map[string]any{
		"width":              writeSize.Width,
		"height":             writeSize.Height,
		"depthOrArrayLayers": writeSize.DepthOrArrayLayers,
	}
	q.ref.Call("writeTexture", dest, uint8Array, layout, size)
}

// SetLabel sets the debug label for the queue.
// This label appears in debuggers and validation layers.
func (q *Queue) SetLabel(label string) {
	q.ref.Set("label", label)
}

// Release releases the queue and all associated resources.
// In JS, this is a no-op as GC handles cleanup.
func (q *Queue) Release() {}
