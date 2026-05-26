//go:build js

package wgpu

import "syscall/js"

// CommandEncoder encodes a sequence of GPU commands that can be submitted to a queue.
// Command encoders are created from a device and are used to build command buffers.
type CommandEncoder struct {
	ref js.Value
}

// Finish finishes recording commands and returns a command buffer.
// The descriptor can be used to set the label of the command buffer.
func (c *CommandEncoder) Finish(descriptor *CommandBufferDescriptor) *CommandBuffer {
	var desc map[string]any
	if descriptor != nil {
		desc = map[string]any{"label": descriptor.Label}
	}
	return &CommandBuffer{ref: c.ref.Call("finish", desc)}
}

// BeginComputePass begins a compute pass and returns a compute pass encoder.
// The descriptor can be used to set the label and timestamp writes for the pass.
func (c *CommandEncoder) BeginComputePass(descriptor *ComputePassDescriptor) *ComputePassEncoder {
	var desc map[string]any
	if descriptor != nil {
		desc = map[string]any{"label": descriptor.Label}
		if descriptor.TimestampWrites != nil {
			desc["timestampWrites"] = map[string]any{
				"querySet":                  descriptor.TimestampWrites.QuerySet.ref,
				"beginningOfPassWriteIndex": descriptor.TimestampWrites.BeginningOfPassWriteIndex,
				"endOfPassWriteIndex":       descriptor.TimestampWrites.EndOfPassWriteIndex,
			}
		}
	}
	return &ComputePassEncoder{ref: c.ref.Call("beginComputePass", desc)}
}

// BeginRenderPass begins a render pass and returns a render pass encoder.
// The descriptor defines the color attachments, depth stencil attachment, and other settings for the pass.
func (c *CommandEncoder) BeginRenderPass(descriptor RenderPassDescriptor) *RenderPassEncoder {
	colorAttachments := make([]any, len(descriptor.ColorAttachments))
	for i, a := range descriptor.ColorAttachments {
		att := map[string]any{
			"loadOp":  a.LoadOp.toJS(),
			"storeOp": a.StoreOp.toJS(),
			"clearValue": map[string]any{
				"r": a.ClearValue.R,
				"g": a.ClearValue.G,
				"b": a.ClearValue.B,
				"a": a.ClearValue.A,
			},
		}
		if a.View != nil {
			att["view"] = a.View.ref
		}
		if a.ResolveTarget != nil {
			att["resolveTarget"] = a.ResolveTarget.ref
		}
		colorAttachments[i] = att
	}

	desc := map[string]any{
		"label":            descriptor.Label,
		"colorAttachments": colorAttachments,
	}

	if descriptor.DepthStencilAttachment != nil {
		dsa := descriptor.DepthStencilAttachment
		dsDesc := map[string]any{
			"depthLoadOp":       dsa.DepthLoadOp.toJS(),
			"depthStoreOp":      dsa.DepthStoreOp.toJS(),
			"depthClearValue":   dsa.DepthClearValue,
			"depthReadOnly":     dsa.DepthReadOnly,
			"stencilLoadOp":     dsa.StencilLoadOp.toJS(),
			"stencilStoreOp":    dsa.StencilStoreOp.toJS(),
			"stencilClearValue": dsa.StencilClearValue,
			"stencilReadOnly":   dsa.StencilReadOnly,
		}
		if dsa.View != nil {
			dsDesc["view"] = dsa.View.ref
		}
		desc["depthStencilAttachment"] = dsDesc
	}

	if descriptor.OcclusionQuerySet != nil {
		desc["occlusionQuerySet"] = descriptor.OcclusionQuerySet.ref
	}

	if descriptor.TimestampWrites != nil {
		desc["timestampWrites"] = map[string]any{
			"querySet":                  descriptor.TimestampWrites.QuerySet.ref,
			"beginningOfPassWriteIndex": descriptor.TimestampWrites.BeginningOfPassWriteIndex,
			"endOfPassWriteIndex":       descriptor.TimestampWrites.EndOfPassWriteIndex,
		}
	}

	return &RenderPassEncoder{ref: c.ref.Call("beginRenderPass", desc)}
}

// CopyBufferToBuffer copies data from one buffer to another.
// The source and destination buffers must have the CopySrc and CopyDst usage flags respectively.
func (c *CommandEncoder) CopyBufferToBuffer(source *Buffer, sourceOffset uint64, destination *Buffer, destinationOffset uint64, size uint64) {
	c.ref.Call("copyBufferToBuffer", source.ref, sourceOffset, destination.ref, destinationOffset, size)
}

// CopyBufferToTexture copies data from a buffer to a texture.
// The source defines the buffer and layout, and the destination defines the texture and region.
func (c *CommandEncoder) CopyBufferToTexture(source TexelCopyBufferInfo, destination TexelCopyTextureInfo, copySize Extent3D) {
	src := map[string]any{
		"buffer":       source.Buffer.ref,
		"offset":       source.Layout.Offset,
		"bytesPerRow":  source.Layout.BytesPerRow,
		"rowsPerImage": source.Layout.RowsPerImage,
	}
	dst := map[string]any{
		"texture":  destination.Texture.ref,
		"mipLevel": destination.MipLevel,
		"origin": map[string]any{
			"x": destination.Origin.X,
			"y": destination.Origin.Y,
			"z": destination.Origin.Z,
		},
		"aspect": destination.Aspect.toJS(),
	}
	size := map[string]any{
		"width":              copySize.Width,
		"height":             copySize.Height,
		"depthOrArrayLayers": copySize.DepthOrArrayLayers,
	}
	c.ref.Call("copyBufferToTexture", src, dst, size)
}

// CopyTextureToBuffer copies data from a texture to a buffer.
// The source defines the texture and region, and the destination defines the buffer and layout.
func (c *CommandEncoder) CopyTextureToBuffer(source TexelCopyTextureInfo, destination TexelCopyBufferInfo, copySize Extent3D) {
	src := map[string]any{
		"texture":  source.Texture.ref,
		"mipLevel": source.MipLevel,
		"origin": map[string]any{
			"x": source.Origin.X,
			"y": source.Origin.Y,
			"z": source.Origin.Z,
		},
		"aspect": source.Aspect.toJS(),
	}
	dst := map[string]any{
		"buffer":       destination.Buffer.ref,
		"offset":       destination.Layout.Offset,
		"bytesPerRow":  destination.Layout.BytesPerRow,
		"rowsPerImage": destination.Layout.RowsPerImage,
	}
	size := map[string]any{
		"width":              copySize.Width,
		"height":             copySize.Height,
		"depthOrArrayLayers": copySize.DepthOrArrayLayers,
	}
	c.ref.Call("copyTextureToBuffer", src, dst, size)
}

// CopyTextureToTexture copies data from one texture to another.
// The source and destination define their respective textures and regions.
func (c *CommandEncoder) CopyTextureToTexture(source TexelCopyTextureInfo, destination TexelCopyTextureInfo, copySize Extent3D) {
	src := map[string]any{
		"texture":  source.Texture.ref,
		"mipLevel": source.MipLevel,
		"origin": map[string]any{
			"x": source.Origin.X,
			"y": source.Origin.Y,
			"z": source.Origin.Z,
		},
		"aspect": source.Aspect.toJS(),
	}
	dst := map[string]any{
		"texture":  destination.Texture.ref,
		"mipLevel": destination.MipLevel,
		"origin": map[string]any{
			"x": destination.Origin.X,
			"y": destination.Origin.Y,
			"z": destination.Origin.Z,
		},
		"aspect": destination.Aspect.toJS(),
	}
	size := map[string]any{
		"width":              copySize.Width,
		"height":             copySize.Height,
		"depthOrArrayLayers": copySize.DepthOrArrayLayers,
	}
	c.ref.Call("copyTextureToTexture", src, dst, size)
}

// ClearBuffer fills a buffer with zeros or a specific value.
// The offset and size specify the range of the buffer to clear.
// If size is 0, the whole buffer is cleared.
func (c *CommandEncoder) ClearBuffer(buffer *Buffer, offset uint64, size uint64) {
	c.ref.Call("clearBuffer", buffer.ref, offset, size)
}

// InsertDebugMarker inserts a debug marker into the command encoder.
// The marker label is used to identify the marker in debuggers and profilers.
func (c *CommandEncoder) InsertDebugMarker(markerLabel string) {
	c.ref.Call("insertDebugMarker", markerLabel)
}

// PopDebugGroup pops the most recently pushed debug group from the command encoder.
func (c *CommandEncoder) PopDebugGroup() {
	c.ref.Call("popDebugGroup")
}

// PushDebugGroup pushes a debug group into the command encoder with the given label.
// Debug groups can be nested and are used to group commands in debuggers and profilers.
func (c *CommandEncoder) PushDebugGroup(groupLabel string) {
	c.ref.Call("pushDebugGroup", groupLabel)
}

// ResolveQuerySet resolves an occlusion or timestamp query set to a buffer.
// The query results are written to the destination buffer starting at destinationOffset.
func (c *CommandEncoder) ResolveQuerySet(querySet *QuerySet, firstQuery uint32, queryCount uint32, destination *Buffer, destinationOffset uint64) {
	c.ref.Call("resolveQuerySet", querySet.ref, firstQuery, queryCount, destination.ref, destinationOffset)
}

// WriteTimestamp writes a timestamp to a query set at the current point in the command encoder.
// Timestamps can be used to measure GPU execution times.
func (c *CommandEncoder) WriteTimestamp(querySet *QuerySet, queryIndex uint32) {
	c.ref.Call("writeTimestamp", querySet.ref, queryIndex)
}

// SetLabel sets the debug label for the command encoder.
// This label appears in debuggers and validation layers.
func (c *CommandEncoder) SetLabel(label string) {
	c.ref.Set("label", label)
}

// Release releases the command encoder and all associated resources.
// After calling this method, the encoder should no longer be used.
// Note: In JS, this is a no-op as GC handles cleanup.
func (c *CommandEncoder) Release() {}
