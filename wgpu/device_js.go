//go:build js

package wgpu

import (
	"errors"
	"fmt"
	"syscall/js"
)

// Device represents a logical GPU device that can be used to create resources and execute commands.
// It is the main interface for interacting with the GPU.
type Device struct {
	ref js.Value
}

// CreateBindGroup creates a bind group from the given descriptor, which defines a set of resources to be bound together.
func (d *Device) CreateBindGroup(descriptor BindGroupDescriptor) *BindGroup {
	entries := make([]any, len(descriptor.Entries))
	for i, e := range descriptor.Entries {
		entry := map[string]any{"binding": e.Binding}
		switch {
		case e.Buffer != nil:
			resource := map[string]any{"buffer": e.Buffer.ref}
			if e.Offset != 0 {
				resource["offset"] = e.Offset
			}
			if e.Size != 0 {
				resource["size"] = e.Size
			}
			entry["resource"] = resource
		case e.Sampler != nil:
			entry["resource"] = e.Sampler.ref
		case e.TextureView != nil:
			entry["resource"] = e.TextureView.ref
		}
		entries[i] = entry
	}

	desc := map[string]any{
		"label":   descriptor.Label,
		"entries": entries,
	}
	if descriptor.Layout != nil {
		desc["layout"] = descriptor.Layout.ref
	}

	return &BindGroup{ref: d.ref.Call("createBindGroup", desc)}
}

// CreateBindGroupLayout creates a bind group layout from the given descriptor, which defines the interface for a bind group.
func (d *Device) CreateBindGroupLayout(descriptor BindGroupLayoutDescriptor) *BindGroupLayout {
	entries := make([]any, len(descriptor.Entries))
	for i, e := range descriptor.Entries {
		entry := map[string]any{
			"binding":    e.Binding,
			"visibility": int(e.Visibility),
		}
		switch {
		case e.Buffer.Type >= BufferBindingTypeUniform:
			entry["buffer"] = map[string]any{
				"type":             e.Buffer.Type.toJS(),
				"hasDynamicOffset": e.Buffer.HasDynamicOffset,
				"minBindingSize":   e.Buffer.MinBindingSize,
			}
		case e.Sampler.Type >= SamplerBindingTypeFiltering:
			entry["sampler"] = map[string]any{
				"type": e.Sampler.Type.toJS(),
			}
		case e.Texture.SampleType >= TextureSampleTypeFloat:
			entry["texture"] = map[string]any{
				"sampleType":    e.Texture.SampleType.toJS(),
				"viewDimension": e.Texture.ViewDimension.toJS(),
				"multisampled":  e.Texture.Multisampled,
			}
		case e.StorageTexture.Access >= StorageTextureAccessWriteOnly:
			entry["storageTexture"] = map[string]any{
				"access":        e.StorageTexture.Access.toJS(),
				"format":        e.StorageTexture.Format.toJS(),
				"viewDimension": e.StorageTexture.ViewDimension.toJS(),
			}
		}
		entries[i] = entry
	}

	return &BindGroupLayout{ref: d.ref.Call("createBindGroupLayout", map[string]any{
		"label":   descriptor.Label,
		"entries": entries,
	})}
}

// CreateBufferInit creates a buffer and initializes it with the given contents in a single operation.
// This is more efficient than creating and then writing to the buffer separately.
func (d *Device) CreateBufferInit(descriptor BufferInitDescriptor) *Buffer {
	buffer := d.CreateBuffer(BufferDescriptor{
		Label: descriptor.Label,
		Size:  uint64(len(descriptor.Contents)),
		Usage: descriptor.Usage | BufferUsageCopyDst,
	})
	d.GetQueue().WriteBuffer(buffer, 0, descriptor.Contents)
	return buffer
}

// CreateBuffer creates a new buffer with the given descriptor.
// Buffers are used to store data that can be read and written by shaders.
func (d *Device) CreateBuffer(descriptor BufferDescriptor) *Buffer {
	return &Buffer{ref: d.ref.Call("createBuffer", map[string]any{
		"label":            descriptor.Label,
		"size":             descriptor.Size,
		"usage":            int(descriptor.Usage),
		"mappedAtCreation": descriptor.MappedAtCreation,
	})}
}

// CreateCommandEncoder creates a command encoder from the given descriptor.
// Command encoders are used to record commands that will be submitted to the GPU queue.
func (d *Device) CreateCommandEncoder(descriptor *CommandEncoderDescriptor) *CommandEncoder {
	var desc map[string]any
	if descriptor != nil {
		desc = map[string]any{"label": descriptor.Label}
	}
	return &CommandEncoder{ref: d.ref.Call("createCommandEncoder", desc)}
}

// CreateComputePipeline creates a compute pipeline from the given descriptor.
// Compute pipelines execute compute shaders on the GPU.
func (d *Device) CreateComputePipeline(descriptor ComputePipelineDescriptor) *ComputePipeline {
	return &ComputePipeline{ref: d.ref.Call("createComputePipeline", toJSComputePipelineDescriptor(descriptor))}
}

// CreateComputePipelineAsync creates a compute pipeline asynchronously from the given descriptor.
// Returns a Future that can be used to wait for the pipeline to be created.
// The callback is called when the pipeline is ready or an error occurs.
func (d *Device) CreateComputePipelineAsync(descriptor ComputePipelineDescriptor, callback CreateComputePipelineAsyncCallback) Future {
	desc := toJSComputePipelineDescriptor(descriptor)
	return newCallbackPromise(func() {
		val, err := awaitPromise(d.ref.Call("createComputePipelineAsync", desc))
		if err != nil {
			callback(CreatePipelineAsyncStatusInternalError, nil, err.Error())
			return
		}
		callback(CreatePipelineAsyncStatusSuccess, &ComputePipeline{ref: val}, "")
	})
}

// CreatePipelineLayout creates a pipeline layout from the given descriptor.
// Pipeline layouts define the resource bindings used by pipelines.
func (d *Device) CreatePipelineLayout(descriptor PipelineLayoutDescriptor) *PipelineLayout {
	layouts := make([]any, len(descriptor.BindGroupLayouts))
	for i, l := range descriptor.BindGroupLayouts {
		layouts[i] = l.ref
	}
	return &PipelineLayout{ref: d.ref.Call("createPipelineLayout", map[string]any{
		"label":            descriptor.Label,
		"bindGroupLayouts": layouts,
	})}
}

// CreateQuerySet creates a query set from the given descriptor.
// Query sets are used to collect timestamp and occlusion query results.
func (d *Device) CreateQuerySet(descriptor QuerySetDescriptor) QuerySet {
	return QuerySet{ref: d.ref.Call("createQuerySet", map[string]any{
		"label": descriptor.Label,
		"type":  descriptor.Type.toJS(),
		"count": descriptor.Count,
	})}
}

// CreateRenderPipeline creates a render pipeline from the given descriptor.
// Render pipelines define how graphics are rendered.
func (d *Device) CreateRenderPipeline(descriptor RenderPipelineDescriptor) *RenderPipeline {
	return &RenderPipeline{ref: d.ref.Call("createRenderPipeline", toJSRenderPipelineDescriptor(descriptor))}
}

// CreateRenderPipelineAsync creates a render pipeline asynchronously from the given descriptor.
// Returns a Future that can be used to wait for the pipeline to be created.
// The callback is called when the pipeline is ready or an error occurs.
func (d *Device) CreateRenderPipelineAsync(descriptor RenderPipelineDescriptor, callback CreateRenderPipelineAsyncCallback) Future {
	desc := toJSRenderPipelineDescriptor(descriptor)
	return newCallbackPromise(func() {
		val, err := awaitPromise(d.ref.Call("createRenderPipelineAsync", desc))
		if err != nil {
			callback(CreatePipelineAsyncStatusInternalError, nil, err.Error())
			return
		}
		callback(CreatePipelineAsyncStatusSuccess, &RenderPipeline{ref: val}, "")
	})
}

// CreateRenderBundleEncoder creates a render bundle encoder from the given descriptor.
// Render bundles are pre-recorded render commands that can be executed efficiently multiple times.
func (d *Device) CreateRenderBundleEncoder(descriptor RenderBundleEncoderDescriptor) RenderBundleEncoder {
	colorFormats := make([]any, len(descriptor.ColorFormats))
	for i, f := range descriptor.ColorFormats {
		colorFormats[i] = f.toJS()
	}
	desc := map[string]any{
		"label":           descriptor.Label,
		"colorFormats":    colorFormats,
		"sampleCount":     descriptor.SampleCount,
		"depthReadOnly":   descriptor.DepthReadOnly,
		"stencilReadOnly": descriptor.StencilReadOnly,
	}
	if descriptor.DepthStencilFormat != TextureFormatUndefined {
		desc["depthStencilFormat"] = descriptor.DepthStencilFormat.toJS()
	}
	return RenderBundleEncoder{ref: d.ref.Call("createRenderBundleEncoder", desc)}
}

// CreateSampler creates a sampler from the given descriptor.
// Samplers define how textures are sampled in shaders.
func (d *Device) CreateSampler(descriptor *SamplerDescriptor) *Sampler {
	var desc map[string]any
	if descriptor != nil {
		desc = map[string]any{
			"label":         descriptor.Label,
			"addressModeU":  descriptor.AddressModeU.toJS(),
			"addressModeV":  descriptor.AddressModeV.toJS(),
			"addressModeW":  descriptor.AddressModeW.toJS(),
			"magFilter":     descriptor.MagFilter.toJS(),
			"minFilter":     descriptor.MinFilter.toJS(),
			"mipmapFilter":  descriptor.MipmapFilter.toJS(),
			"lodMinClamp":   descriptor.LodMinClamp,
			"lodMaxClamp":   descriptor.LodMaxClamp,
			"maxAnisotropy": descriptor.MaxAnisotropy,
		}
		if descriptor.Compare != CompareFunctionUndefined {
			desc["compare"] = descriptor.Compare.toJS()
		}
	}
	return &Sampler{ref: d.ref.Call("createSampler", desc)}
}

// CreateShaderModule creates a shader module from the given descriptor.
// Shader modules contain shader code (WGSL or SPIR-V) that can be used in pipelines.
func (d *Device) CreateShaderModule(descriptor ShaderModuleDescriptor) *ShaderModule {
	desc := map[string]any{"label": descriptor.Label}
	if descriptor.WGSLSource != nil {
		desc["code"] = descriptor.WGSLSource.Code
	}
	return &ShaderModule{ref: d.ref.Call("createShaderModule", desc)}
}

// CreateTexture creates a new texture with the given descriptor.
// Textures are used to store image data that can be sampled by shaders.
func (d *Device) CreateTexture(descriptor *TextureDescriptor) *Texture {
	var desc map[string]any
	if descriptor != nil {
		desc = map[string]any{
			"label":     descriptor.Label,
			"usage":     int(descriptor.Usage),
			"dimension": descriptor.Dimension.toJS(),
			"size": map[string]any{
				"width":              descriptor.Size.Width,
				"height":             descriptor.Size.Height,
				"depthOrArrayLayers": descriptor.Size.DepthOrArrayLayers,
			},
			"format":        descriptor.Format.toJS(),
			"mipLevelCount": descriptor.MipLevelCount,
			"sampleCount":   descriptor.SampleCount,
		}
		if len(descriptor.ViewFormats) > 0 {
			vf := make([]any, len(descriptor.ViewFormats))
			for i, f := range descriptor.ViewFormats {
				vf[i] = f.toJS()
			}
			desc["viewFormats"] = vf
		}
	}
	return &Texture{ref: d.ref.Call("createTexture", desc)}
}

// GetQueue returns the default queue for this device.
// The queue is used to submit commands to the GPU.
func (d *Device) GetQueue() *Queue {
	return &Queue{ref: d.ref.Get("queue")}
}

// GetLimits returns the limits supported by the device.
// Returns the limits and an error if they cannot be retrieved.
func (d *Device) GetLimits() (Limits, error) {
	jsLimits := d.ref.Get("limits")
	if jsLimits.IsUndefined() || jsLimits.IsNull() {
		return Limits{}, fmt.Errorf("device limits unavailable")
	}
	return limitsFromJS(jsLimits), nil
}

// HasFeature checks if the device supports the given feature.
// Returns true if the feature is supported, false otherwise.
func (d *Device) HasFeature(feature FeatureName) bool {
	name := feature.toJS()
	if name == "" {
		return false
	}
	return d.ref.Get("features").Call("has", name).Bool()
}

// GetFeatures returns a list of all features supported by the device.
func (d *Device) GetFeatures() []FeatureName {
	return featuresFromJS(d.ref.Get("features"))
}

// GetAdapterInfo returns information about the adapter that created this device.
// Returns the adapter info and an error if it cannot be retrieved.
func (d *Device) GetAdapterInfo() (AdapterInfo, error) {
	jsInfo := d.ref.Get("adapterInfo")
	if jsInfo.IsUndefined() || jsInfo.IsNull() {
		return AdapterInfo{BackendType: BackendTypeWebGPU}, fmt.Errorf("error getting adapter info")
	}
	return AdapterInfo{
		Vendor:       jsInfo.Get("vendor").String(),
		Architecture: jsInfo.Get("architecture").String(),
		Device:       jsInfo.Get("device").String(),
		Description:  jsInfo.Get("description").String(),
		BackendType:  BackendTypeWebGPU,
	}, nil
}

// SetLabel sets the debug label for the device.
// This label appears in debuggers and validation layers.
func (d *Device) SetLabel(label string) {
	d.ref.Set("label", label)
}

// PushErrorScope pushes an error scope onto the device's error scope stack.
// Errors that match the filter will be captured in the scope.
func (d *Device) PushErrorScope(filter ErrorFilter) {
	d.ref.Call("pushErrorScope", filter.toJS())
}

// PopErrorScope pops an error scope from the device's error scope stack and calls the callback with the result.
// The callback is called with the error type and message, or no error if the scope was empty.
func (d *Device) PopErrorScope(callback PopErrorScopeCallback) {
	promise := d.ref.Call("popErrorScope")
	go func() {
		val, err := awaitPromise(promise)
		if err != nil {
			callback(ErrorTypeUnknown, err.Error())
			return
		}
		typ, msg := ErrorTypeFromJS(val)
		callback(typ, msg)
	}()
}

// Try executes the given function and captures any errors that occur.
// The optional filters specify which error types to capture; if not provided, all error types are captured.
// Returns the captured error, if any.
func (d *Device) Try(fn func(), filters ...ErrorFilter) error {
	if len(filters) == 0 {
		filters = []ErrorFilter{ErrorFilterValidation, ErrorFilterOutOfMemory, ErrorFilterInternal}
	}

	for _, f := range filters {
		d.PushErrorScope(f)
	}

	fn()

	errs := make([]chan error, len(filters))
	for i := range filters {
		ch := make(chan error, 1)
		errs[i] = ch
		callback := PopErrorScopeCallback(func(typ ErrorType, message string) {
			if typ != ErrorTypeNoError {
				ch <- fmt.Errorf("%s error: %s", typ, message)
			} else {
				ch <- nil
			}
		})
		d.PopErrorScope(callback)
	}

	var combined error
	for _, ch := range errs {
		if e := <-ch; e != nil {
			combined = errors.Join(combined, e)
		}
	}
	return combined
}

// Release releases the device and all associated resources.
// After calling this method, the device should no longer be used.
func (d *Device) Release() {
	d.ref.Call("destroy")
}

// Destroy destroys the device and all associated resources.
// This is similar to Release but also frees all GPU resources associated with the device.
func (d *Device) Destroy() {
	d.ref.Call("destroy")
}

// toJSComputePipelineDescriptor converts a Go ComputePipelineDescriptor to a JS object.
func toJSComputePipelineDescriptor(descriptor ComputePipelineDescriptor) map[string]any {
	compute := map[string]any{
		"entryPoint": descriptor.Compute.EntryPoint,
	}
	if descriptor.Compute.Module != nil {
		compute["module"] = descriptor.Compute.Module.ref
	}
	if len(descriptor.Compute.Constants) > 0 {
		consts := map[string]any{}
		for k, v := range descriptor.Compute.Constants {
			consts[k] = v
		}
		compute["constants"] = consts
	}

	desc := map[string]any{
		"label":   descriptor.Label,
		"compute": compute,
	}
	if descriptor.Layout != nil {
		desc["layout"] = descriptor.Layout.ref
	} else {
		desc["layout"] = "auto"
	}
	return desc
}

// toJSRenderPipelineDescriptor converts a Go RenderPipelineDescriptor to a JS object.
func toJSRenderPipelineDescriptor(descriptor RenderPipelineDescriptor) map[string]any {
	desc := map[string]any{
		"label": descriptor.Label,
	}
	if descriptor.Layout != nil {
		desc["layout"] = descriptor.Layout.ref
	} else {
		desc["layout"] = "auto"
	}

	// Vertex
	vertex := map[string]any{
		"entryPoint": descriptor.Vertex.EntryPoint,
	}
	if descriptor.Vertex.Module != nil {
		vertex["module"] = descriptor.Vertex.Module.ref
	}
	if len(descriptor.Vertex.Constants) > 0 {
		consts := map[string]any{}
		for k, v := range descriptor.Vertex.Constants {
			consts[k] = v
		}
		vertex["constants"] = consts
	}
	if len(descriptor.Vertex.Buffers) > 0 {
		buffers := make([]any, len(descriptor.Vertex.Buffers))
		for i, buf := range descriptor.Vertex.Buffers {
			attrs := make([]any, len(buf.Attributes))
			for j, attr := range buf.Attributes {
				attrs[j] = map[string]any{
					"format":         attr.Format.toJS(),
					"offset":         attr.Offset,
					"shaderLocation": attr.ShaderLocation,
				}
			}
			buffers[i] = map[string]any{
				"stepMode":    buf.StepMode.toJS(),
				"arrayStride": buf.ArrayStride,
				"attributes":  attrs,
			}
		}
		vertex["buffers"] = buffers
	}
	desc["vertex"] = vertex

	// Primitive
	primitive := map[string]any{
		"topology":  descriptor.Primitive.Topology.toJS(),
		"frontFace": descriptor.Primitive.FrontFace.toJS(),
		"cullMode":  descriptor.Primitive.CullMode.toJS(),
	}
	if descriptor.Primitive.StripIndexFormat != IndexFormatUndefined {
		primitive["stripIndexFormat"] = descriptor.Primitive.StripIndexFormat.toJS()
	}
	if descriptor.Primitive.UnclippedDepth {
		primitive["unclippedDepth"] = true
	}
	desc["primitive"] = primitive

	// Depth stencil
	if descriptor.DepthStencil != nil {
		ds := map[string]any{
			"format":            descriptor.DepthStencil.Format.toJS(),
			"depthWriteEnabled": descriptor.DepthStencil.DepthWriteEnabled.toJS(),
			"depthCompare":      descriptor.DepthStencil.DepthCompare.toJS(),
			"stencilFront": map[string]any{
				"compare":     descriptor.DepthStencil.StencilFront.Compare.toJS(),
				"failOp":      descriptor.DepthStencil.StencilFront.FailOp.toJS(),
				"depthFailOp": descriptor.DepthStencil.StencilFront.DepthFailOp.toJS(),
				"passOp":      descriptor.DepthStencil.StencilFront.PassOp.toJS(),
			},
			"stencilBack": map[string]any{
				"compare":     descriptor.DepthStencil.StencilBack.Compare.toJS(),
				"failOp":      descriptor.DepthStencil.StencilBack.FailOp.toJS(),
				"depthFailOp": descriptor.DepthStencil.StencilBack.DepthFailOp.toJS(),
				"passOp":      descriptor.DepthStencil.StencilBack.PassOp.toJS(),
			},
			"stencilReadMask":     descriptor.DepthStencil.StencilReadMask,
			"stencilWriteMask":    descriptor.DepthStencil.StencilWriteMask,
			"depthBias":           descriptor.DepthStencil.DepthBias,
			"depthBiasSlopeScale": descriptor.DepthStencil.DepthBiasSlopeScale,
			"depthBiasClamp":      descriptor.DepthStencil.DepthBiasClamp,
		}
		desc["depthStencil"] = ds
	}

	// Multisample
	desc["multisample"] = map[string]any{
		"count":                  descriptor.Multisample.Count,
		"mask":                   descriptor.Multisample.Mask,
		"alphaToCoverageEnabled": descriptor.Multisample.AlphaToCoverageEnabled,
	}

	// Fragment
	if descriptor.Fragment != nil {
		frag := map[string]any{
			"entryPoint": descriptor.Fragment.EntryPoint,
		}
		if descriptor.Fragment.Module != nil {
			frag["module"] = descriptor.Fragment.Module.ref
		}
		if len(descriptor.Fragment.Constants) > 0 {
			consts := map[string]any{}
			for k, v := range descriptor.Fragment.Constants {
				consts[k] = v
			}
			frag["constants"] = consts
		}
		targets := make([]any, len(descriptor.Fragment.Targets))
		for i, t := range descriptor.Fragment.Targets {
			target := map[string]any{
				"format":    t.Format.toJS(),
				"writeMask": int(t.WriteMask),
			}
			if t.Blend != nil {
				target["blend"] = map[string]any{
					"color": map[string]any{
						"operation": t.Blend.Color.Operation.toJS(),
						"srcFactor": t.Blend.Color.SrcFactor.toJS(),
						"dstFactor": t.Blend.Color.DstFactor.toJS(),
					},
					"alpha": map[string]any{
						"operation": t.Blend.Alpha.Operation.toJS(),
						"srcFactor": t.Blend.Alpha.SrcFactor.toJS(),
						"dstFactor": t.Blend.Alpha.DstFactor.toJS(),
					},
				}
			}
			targets[i] = target
		}
		frag["targets"] = targets
		desc["fragment"] = frag
	}

	return desc
}
