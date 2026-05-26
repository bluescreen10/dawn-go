//go:build js

package wgpu

import "syscall/js"

// Buffer represents a GPU buffer, which is a region of memory that can be used to store data.
// Buffers are created from a device and can be mapped for reading or writing.
type Buffer struct {
	ref         js.Value
	mapData     []byte   // backing slice for write-mapped data
	mapJSBuffer js.Value // JS ArrayBuffer for write-back on Unmap; undefined for read mappings
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

// MapAsync maps the buffer for reading or writing asynchronously.
// The mode specifies whether to map for reading or writing.
// The offset and size specify the range of the buffer to map.
// The callback is called when the mapping is complete.
// Returns a Future that can be used to wait for the mapping to complete.
func (b *Buffer) MapAsync(mode MapMode, offset int, size int, callback BufferMapCallback) Future {
	promise := b.ref.Call("mapAsync", int(mode), offset, size)
	return newCallbackPromise(func() {
		_, err := awaitPromise(promise)
		if err != nil {
			callback(MapAsyncStatusError, err.Error())
			return
		}
		callback(MapAsyncStatusSuccess, "")
	})
}

// GetMappedRange returns a slice of bytes representing the mapped range of the buffer for writing.
// The buffer must have been mapped with MapModeWrite.
func (b *Buffer) GetMappedRange(offset int, size int) []byte {
	jsAB := b.ref.Call("getMappedRange", offset, size)
	uint8Array := js.Global().Get("Uint8Array").New(jsAB)
	data := make([]byte, size)
	js.CopyBytesToGo(data, uint8Array)
	b.mapData = data
	b.mapJSBuffer = jsAB
	return data
}

// GetConstMappedRange returns a slice of bytes representing the mapped range of the buffer for reading.
// The buffer must have been mapped with MapModeRead.
func (b *Buffer) GetConstMappedRange(offset int, size int) []byte {
	jsAB := b.ref.Call("getMappedRange", offset, size)
	uint8Array := js.Global().Get("Uint8Array").New(jsAB)
	data := make([]byte, size)
	js.CopyBytesToGo(data, uint8Array)
	b.mapData = data
	return data
}

// Unmap unmaps the buffer, flushing any writes if it was mapped for writing.
// After unmapping, the mapped range is no longer valid.
func (b *Buffer) Unmap() {
	if b.mapData != nil && !b.mapJSBuffer.IsUndefined() {
		uint8Array := js.Global().Get("Uint8Array").New(b.mapJSBuffer)
		js.CopyBytesToJS(uint8Array, b.mapData)
	}
	b.mapData = nil
	b.mapJSBuffer = js.Undefined()
	b.ref.Call("unmap")
}

// SetLabel sets the debug label for the buffer.
// This label appears in debuggers and validation layers.
func (b *Buffer) SetLabel(label string) {
	b.ref.Set("label", label)
}

// GetUsage returns the usage flags for the buffer.
func (b *Buffer) GetUsage() BufferUsage {
	return BufferUsage(b.ref.Get("usage").Int())
}

// GetSize returns the size of the buffer in bytes.
func (b *Buffer) GetSize() uint64 {
	return uint64(b.ref.Get("size").Int())
}

// GetMapState returns the current mapping state of the buffer.
func (b *Buffer) GetMapState() BufferMapState {
	switch b.ref.Get("mapState").String() {
	case "mapped":
		return BufferMapStateMapped
	case "pending":
		return BufferMapStatePending
	default:
		return BufferMapStateUnmapped
	}
}

// Release is a no-op in JS (GC handles cleanup).
func (b *Buffer) Release() {}

// Destroy destroys the buffer and frees all associated GPU resources.
func (b *Buffer) Destroy() {
	b.ref.Call("destroy")
}
