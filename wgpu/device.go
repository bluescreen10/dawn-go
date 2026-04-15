//go:build !js

package wgpu

/*
#include "webgpu.h"

extern void cgo_callback_CreateComputePipelineAsyncCallback(WGPUCreatePipelineAsyncStatus status, WGPUComputePipeline pipeline, WGPUStringView message, void *userData1, void *userData2);
extern void cgo_callback_CreateRenderPipelineAsyncCallback(WGPUCreatePipelineAsyncStatus status, WGPURenderPipeline pipeline, WGPUStringView message, void *userData1, void *userData2);
extern void cgo_callback_PopErrorScopeCallback(WGPUPopErrorScopeStatus status, WGPUErrorType typ, WGPUStringView message, void *userData1, void *userData2);
*/
import "C"

import (
	"errors"
	"fmt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

// Device represents a logical GPU device that can be used to create resources and execute commands.
// It is the main interface for interacting with the GPU.
type Device struct {
	ref     C.WGPUDevice
	handles []cgo.Handle
}

// CreateBindGroup creates a bind group from the given descriptor, which defines a set of resources to be bound together.
func (d *Device) CreateBindGroup(descriptor BindGroupDescriptor) *BindGroup {
	var pinner runtime.Pinner
	defer pinner.Unpin()

	cDescriptor := C.WGPUBindGroupDescriptor{
		label: toCStr(descriptor.Label),
	}

	if descriptor.Layout != nil {
		cDescriptor.layout = descriptor.Layout.ref
	}

	if count := C.size_t(len(descriptor.Entries)); count > 0 {
		entries := make([]C.WGPUBindGroupEntry, count)
		pinner.Pin(&entries[0])

		cDescriptor.entries = (*C.WGPUBindGroupEntry)(unsafe.Pointer(&entries[0]))
		cDescriptor.entryCount = count

		for i, e := range descriptor.Entries {
			entries[i].binding = C.uint32_t(e.Binding)
			entries[i].offset = C.uint64_t(e.Offset)
			entries[i].size = C.uint64_t(e.Size)

			if e.Buffer != nil {
				entries[i].buffer = e.Buffer.ref
			}

			if e.Sampler != nil {
				entries[i].sampler = e.Sampler.ref
			}

			if e.TextureView != nil {
				entries[i].textureView = e.TextureView.ref
			}
		}
	}

	return &BindGroup{ref: C.wgpuDeviceCreateBindGroup(d.ref, &cDescriptor)}
}

// CreateBindGroupLayout creates a bind group layout from the given descriptor, which defines the interface for a bind group.
func (d *Device) CreateBindGroupLayout(descriptor BindGroupLayoutDescriptor) *BindGroupLayout {
	var pinner runtime.Pinner
	defer pinner.Unpin()

	cDescriptor := C.WGPUBindGroupLayoutDescriptor{
		label: toCStr(descriptor.Label),
	}

	if count := C.size_t(len(descriptor.Entries)); count > 0 {
		entries := make([]C.WGPUBindGroupLayoutEntry, count)
		pinner.Pin(&entries[0])

		cDescriptor.entries = (*C.WGPUBindGroupLayoutEntry)(unsafe.Pointer(&entries[0]))
		cDescriptor.entryCount = count

		for i, e := range descriptor.Entries {
			entries[i] = C.WGPUBindGroupLayoutEntry{
				binding:          C.uint32_t(e.Binding),
				visibility:       C.WGPUShaderStage(e.Visibility),
				bindingArraySize: C.uint32_t(e.BindingArraySize),
				buffer: C.WGPUBufferBindingLayout{
					_type:            C.WGPUBufferBindingType(e.Buffer.Type),
					hasDynamicOffset: toCBool(e.Buffer.HasDynamicOffset),
					minBindingSize:   C.uint64_t(e.Buffer.MinBindingSize),
				},
				sampler: C.WGPUSamplerBindingLayout{
					_type: C.WGPUSamplerBindingType(e.Sampler.Type),
				},
				texture: C.WGPUTextureBindingLayout{
					sampleType:    C.WGPUTextureSampleType(e.Texture.SampleType),
					viewDimension: C.WGPUTextureViewDimension(e.Texture.ViewDimension),
					multisampled:  toCBool(e.Texture.Multisampled),
				},
				storageTexture: C.WGPUStorageTextureBindingLayout{
					access:        C.WGPUStorageTextureAccess(e.StorageTexture.Access),
					format:        C.WGPUTextureFormat(e.StorageTexture.Format),
					viewDimension: C.WGPUTextureViewDimension(e.StorageTexture.ViewDimension),
				},
			}
		}
	}

	return &BindGroupLayout{ref: C.wgpuDeviceCreateBindGroupLayout(d.ref, &cDescriptor)}
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
	var pinner runtime.Pinner
	defer pinner.Unpin()

	cDescriptor := C.WGPUBufferDescriptor{
		label:            toCStr(descriptor.Label),
		usage:            C.WGPUBufferUsage(descriptor.Usage),
		size:             C.uint64_t(descriptor.Size),
		mappedAtCreation: toCBool(descriptor.MappedAtCreation),
	}

	pinner.Pin(cDescriptor.label.data)

	return &Buffer{ref: C.wgpuDeviceCreateBuffer(d.ref, &cDescriptor)}
}

// CreateCommandEncoder creates a command encoder from the given descriptor.
// Command encoders are used to record commands that will be submitted to the GPU queue.
func (d *Device) CreateCommandEncoder(descriptor *CommandEncoderDescriptor) *CommandEncoder {

	var cDescriptor *C.WGPUCommandEncoderDescriptor

	if descriptor != nil {
		cDescriptor = &C.WGPUCommandEncoderDescriptor{
			label: toCStr(descriptor.Label),
		}
	}

	return &CommandEncoder{ref: C.wgpuDeviceCreateCommandEncoder(d.ref, cDescriptor)}
}

// CreateComputePipeline creates a compute pipeline from the given descriptor.
// Compute pipelines execute compute shaders on the GPU.
func (d *Device) CreateComputePipeline(descriptor ComputePipelineDescriptor) *ComputePipeline {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()

	cDescriptor := toCComputePipelineDescriptor(pinner, descriptor)
	return &ComputePipeline{ref: C.wgpuDeviceCreateComputePipeline(d.ref, &cDescriptor)}
}

//export goCreateComputePipelineAsyncCallbackHandler
func goCreateComputePipelineAsyncCallbackHandler(status C.WGPUCreatePipelineAsyncStatus, pipeline C.WGPUComputePipeline, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	defer handle.Delete()
	fn := handle.Value().(CreateComputePipelineAsyncCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}

	go fn(
		CreatePipelineAsyncStatus(status),
		&ComputePipeline{ref: pipeline},
		msg,
	)
}

// CreateComputePipelineAsync creates a compute pipeline asynchronously from the given descriptor.
// Returns a Future that can be used to wait for the pipeline to be created.
// The callback is called when the pipeline is ready or an error occurs.
func (d *Device) CreateComputePipelineAsync(descriptor ComputePipelineDescriptor, callback CreateComputePipelineAsyncCallback) Future {

	pinner := &runtime.Pinner{}
	defer pinner.Unpin()

	cDescriptor := toCComputePipelineDescriptor(pinner, descriptor)

	handle := cgo.NewHandle(callback)

	cCallbackInfo := C.WGPUCreateComputePipelineAsyncCallbackInfo{
		mode:      C.WGPUCallbackMode(callbackModeAllowSpontaneous),
		callback:  C.WGPUCreateComputePipelineAsyncCallback(C.cgo_callback_CreateComputePipelineAsyncCallback),
		userdata1: unsafe.Pointer(handle),
	}

	future := C.wgpuDeviceCreateComputePipelineAsync(d.ref, &cDescriptor, cCallbackInfo)
	return Future{id: uint64(future.id)}
}

// CreatePipelineLayout creates a pipeline layout from the given descriptor.
// Pipeline layouts define the resource bindings used by pipelines.
func (d *Device) CreatePipelineLayout(descriptor PipelineLayoutDescriptor) PipelineLayout {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	cDescriptor := C.WGPUPipelineLayoutDescriptor{
		label: toCStr(descriptor.Label),
	}

	if count := C.size_t(len(descriptor.BindGroupLayouts)); count > 0 {
		layouts := make([]C.WGPUBindGroupLayout, count)
		pinner.Pin(&layouts[0])

		cDescriptor.bindGroupLayouts = (*C.WGPUBindGroupLayout)(unsafe.Pointer(&layouts[0]))
		cDescriptor.bindGroupLayoutCount = count

		for i, l := range descriptor.BindGroupLayouts {
			layouts[i] = l.ref
		}
	}

	return PipelineLayout{ref: C.wgpuDeviceCreatePipelineLayout(d.ref, &cDescriptor)}
}

// CreateQuerySet creates a query set from the given descriptor.
// Query sets are used to collect timestamp and occlusion query results.
func (d *Device) CreateQuerySet(descriptor QuerySetDescriptor) QuerySet {
	cDescriptor := C.WGPUQuerySetDescriptor{
		label: toCStr(descriptor.Label),
		count: C.uint32_t(descriptor.Count),
		_type: C.WGPUQueryType(descriptor.Type),
	}

	return QuerySet{ref: C.wgpuDeviceCreateQuerySet(d.ref, &cDescriptor)}
}

//export goCreateRenderPipelineAsyncCallbackHandler
func goCreateRenderPipelineAsyncCallbackHandler(status C.WGPUCreatePipelineAsyncStatus, pipeline C.WGPURenderPipeline, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	defer handle.Delete()
	fn := handle.Value().(CreateRenderPipelineAsyncCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}

	go fn(
		CreatePipelineAsyncStatus(status),
		&RenderPipeline{ref: pipeline},
		msg,
	)
}

// CreateRenderPipelineAsync creates a render pipeline asynchronously from the given descriptor.
// Returns a Future that can be used to wait for the pipeline to be created.
// The callback is called when the pipeline is ready or an error occurs.
func (d *Device) CreateRenderPipelineAsync(descriptor RenderPipelineDescriptor, callback CreateRenderPipelineAsyncCallback) Future {

	pinner := &runtime.Pinner{}
	cDescriptor := toCRenderPipelineDescriptor(pinner, descriptor)

	handle := cgo.NewHandle(callback)

	cCallbackInfo := C.WGPUCreateRenderPipelineAsyncCallbackInfo{
		mode:      C.WGPUCallbackMode(callbackModeAllowSpontaneous),
		callback:  C.WGPUCreateRenderPipelineAsyncCallback(C.cgo_callback_CreateRenderPipelineAsyncCallback),
		userdata1: unsafe.Pointer(handle),
	}

	future := C.wgpuDeviceCreateRenderPipelineAsync(d.ref, &cDescriptor, cCallbackInfo)
	return Future{id: uint64(future.id)}
}

// CreateRenderBundleEncoder creates a render bundle encoder from the given descriptor.
// Render bundles are pre-recorded render commands that can be executed efficiently multiple times.
func (d *Device) CreateRenderBundleEncoder(descriptor RenderBundleEncoderDescriptor) RenderBundleEncoder {
	var pinner runtime.Pinner
	defer pinner.Unpin()

	cDescriptor := C.WGPURenderBundleEncoderDescriptor{
		label:              toCStr(descriptor.Label),
		depthStencilFormat: C.WGPUTextureFormat(descriptor.DepthStencilFormat),
		sampleCount:        C.uint32_t(descriptor.SampleCount),
		depthReadOnly:      toCBool(descriptor.DepthReadOnly),
		stencilReadOnly:    toCBool(descriptor.StencilReadOnly),
	}

	if count := C.size_t(len(descriptor.ColorFormats)); count > 0 {
		colorFormats := make([]C.WGPUTextureFormat, count)
		pinner.Pin(&colorFormats[0])

		cDescriptor.colorFormats = (*C.WGPUTextureFormat)(unsafe.Pointer(&colorFormats[0]))
		cDescriptor.colorFormatCount = count

		for i, f := range descriptor.ColorFormats {
			colorFormats[i] = C.WGPUTextureFormat(f)
		}
	}

	return RenderBundleEncoder{ref: C.wgpuDeviceCreateRenderBundleEncoder(d.ref, &cDescriptor)}
}

// CreateRenderPipeline creates a render pipeline from the given descriptor.
// Render pipelines define how graphics are rendered.
func (d *Device) CreateRenderPipeline(descriptor RenderPipelineDescriptor) *RenderPipeline {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()

	cDescriptor := toCRenderPipelineDescriptor(pinner, descriptor)
	return &RenderPipeline{ref: C.wgpuDeviceCreateRenderPipeline(d.ref, &cDescriptor)}
}

// CreateSampler creates a sampler from the given descriptor.
// Samplers define how textures are sampled in shaders.
func (d *Device) CreateSampler(descriptor *SamplerDescriptor) *Sampler {
	var cDescriptor *C.WGPUSamplerDescriptor

	if descriptor != nil {
		cDescriptor = &C.WGPUSamplerDescriptor{
			label:         toCStr(descriptor.Label),
			addressModeU:  C.WGPUAddressMode(descriptor.AddressModeU),
			addressModeV:  C.WGPUAddressMode(descriptor.AddressModeV),
			addressModeW:  C.WGPUAddressMode(descriptor.AddressModeW),
			magFilter:     C.WGPUFilterMode(descriptor.MagFilter),
			minFilter:     C.WGPUFilterMode(descriptor.MinFilter),
			mipmapFilter:  C.WGPUMipmapFilterMode(descriptor.MipmapFilter),
			lodMinClamp:   C.float(descriptor.LodMinClamp),
			lodMaxClamp:   C.float(descriptor.LodMaxClamp),
			compare:       C.WGPUCompareFunction(descriptor.Compare),
			maxAnisotropy: C.uint16_t(descriptor.MaxAnisotropy),
		}
	}

	return &Sampler{ref: C.wgpuDeviceCreateSampler(d.ref, cDescriptor)}
}

// CreateShaderModule creates a shader module from the given descriptor.
// Shader modules contain shader code (WGSL or SPIR-V) that can be used in pipelines.
func (d *Device) CreateShaderModule(descriptor ShaderModuleDescriptor) *ShaderModule {
	var pinner runtime.Pinner
	defer pinner.Unpin()

	cDescriptor := C.WGPUShaderModuleDescriptor{
		label: toCStr(descriptor.Label),
	}

	if descriptor.WGSLSource != nil {
		wgslSource := C.WGPUShaderSourceWGSL{
			chain: C.WGPUChainedStruct{
				next:  nil,
				sType: C.WGPUSType_ShaderSourceWGSL,
			},
			code: toCStr(descriptor.WGSLSource.Code),
		}
		pinner.Pin(&wgslSource)
		cDescriptor.nextInChain = (*C.WGPUChainedStruct)(unsafe.Pointer(&wgslSource))
	}

	if descriptor.SPIRVSource != nil {
		spirvSource := C.WGPUShaderSourceSPIRV{
			chain: C.WGPUChainedStruct{
				next:  nil,
				sType: C.WGPUSType_ShaderSourceSPIRV,
			},
			code:     (*C.uint32_t)(&descriptor.SPIRVSource.Code[0]),
			codeSize: C.uint32_t(len(descriptor.SPIRVSource.Code)),
		}

		pinner.Pin(&spirvSource)
		cDescriptor.nextInChain = (*C.WGPUChainedStruct)(unsafe.Pointer(&spirvSource))
	}

	return &ShaderModule{ref: C.wgpuDeviceCreateShaderModule(d.ref, &cDescriptor)}
}

// CreateTexture creates a new texture with the given descriptor.
// Textures are used to store image data that can be sampled by shaders.
func (d *Device) CreateTexture(descriptor *TextureDescriptor) *Texture {
	var pinner runtime.Pinner
	defer pinner.Unpin()

	var cDescriptor *C.WGPUTextureDescriptor

	if descriptor != nil {
		cDescriptor = &C.WGPUTextureDescriptor{
			usage:     C.WGPUTextureUsage(descriptor.Usage),
			dimension: C.WGPUTextureDimension(descriptor.Dimension),
			size: C.WGPUExtent3D{
				width:              C.uint32_t(descriptor.Size.Width),
				height:             C.uint32_t(descriptor.Size.Height),
				depthOrArrayLayers: C.uint32_t(descriptor.Size.DepthOrArrayLayers),
			},
			format:        C.WGPUTextureFormat(descriptor.Format),
			mipLevelCount: C.uint32_t(descriptor.MipLevelCount),
			sampleCount:   C.uint32_t(descriptor.SampleCount),
		}

		cDescriptor.label = toCStr(descriptor.Label)

		if count := C.size_t(len(descriptor.ViewFormats)); count > 0 {
			viewFormats := make([]C.WGPUTextureFormat, count)
			pinner.Pin(&viewFormats[0])

			cDescriptor.viewFormats = (*C.WGPUTextureFormat)(unsafe.Pointer(&viewFormats[0]))
			cDescriptor.viewFormatCount = count

			for i, f := range descriptor.ViewFormats {
				viewFormats[i] = C.WGPUTextureFormat(f)
			}
		}
	}

	return &Texture{ref: C.wgpuDeviceCreateTexture(d.ref, cDescriptor)}
}

// Release releases the device and all associated resources.
// After calling this method, the device should no longer be used.
func (d *Device) Release() {
	for _, h := range d.handles {
		h.Delete()
	}
	d.handles = nil

	C.wgpuDeviceRelease(d.ref)
}

// Destroy destroys the device and all associated resources.
// This is similar to Release but also frees all GPU resources associated with the device.
func (d *Device) Destroy() {
	for _, h := range d.handles {
		h.Delete()
	}
	d.handles = nil

	C.wgpuDeviceDestroy(d.ref)
}

// GetLimits returns the limits supported by the device.
// Returns the limits and an error if they cannot be retrieved.
func (d *Device) GetLimits() (Limits, error) {

	var cLimits C.WGPULimits
	status := C.wgpuDeviceGetLimits(d.ref, &cLimits)

	if statusCode(status) != statusCodeSuccess {
		return Limits{}, fmt.Errorf("error getting device limits")
	}

	limits := Limits{
		MaxTextureDimension1D:                     uint32(cLimits.maxTextureDimension1D),
		MaxTextureDimension2D:                     uint32(cLimits.maxTextureDimension2D),
		MaxTextureDimension3D:                     uint32(cLimits.maxTextureDimension3D),
		MaxTextureArrayLayers:                     uint32(cLimits.maxTextureArrayLayers),
		MaxBindGroups:                             uint32(cLimits.maxBindGroups),
		MaxBindGroupsPlusVertexBuffers:            uint32(cLimits.maxBindGroupsPlusVertexBuffers),
		MaxBindingsPerBindGroup:                   uint32(cLimits.maxBindingsPerBindGroup),
		MaxDynamicUniformBuffersPerPipelineLayout: uint32(cLimits.maxDynamicUniformBuffersPerPipelineLayout),
		MaxDynamicStorageBuffersPerPipelineLayout: uint32(cLimits.maxDynamicStorageBuffersPerPipelineLayout),
		MaxSampledTexturesPerShaderStage:          uint32(cLimits.maxSampledTexturesPerShaderStage),
		MaxSamplersPerShaderStage:                 uint32(cLimits.maxSamplersPerShaderStage),
		MaxStorageBuffersPerShaderStage:           uint32(cLimits.maxStorageBuffersPerShaderStage),
		MaxStorageTexturesPerShaderStage:          uint32(cLimits.maxStorageTexturesPerShaderStage),
		MaxUniformBuffersPerShaderStage:           uint32(cLimits.maxUniformBuffersPerShaderStage),
		MaxUniformBufferBindingSize:               uint64(cLimits.maxUniformBufferBindingSize),
		MaxStorageBufferBindingSize:               uint64(cLimits.maxStorageBufferBindingSize),
		MinUniformBufferOffsetAlignment:           uint32(cLimits.minUniformBufferOffsetAlignment),
		MinStorageBufferOffsetAlignment:           uint32(cLimits.minStorageBufferOffsetAlignment),
		MaxVertexBuffers:                          uint32(cLimits.maxVertexBuffers),
		MaxBufferSize:                             uint64(cLimits.maxBufferSize),
		MaxVertexAttributes:                       uint32(cLimits.maxVertexAttributes),
		MaxVertexBufferArrayStride:                uint32(cLimits.maxVertexBufferArrayStride),
		MaxInterStageShaderVariables:              uint32(cLimits.maxInterStageShaderVariables),
		MaxColorAttachments:                       uint32(cLimits.maxColorAttachments),
		MaxColorAttachmentBytesPerSample:          uint32(cLimits.maxColorAttachmentBytesPerSample),
		MaxComputeWorkgroupStorageSize:            uint32(cLimits.maxComputeWorkgroupStorageSize),
		MaxComputeInvocationsPerWorkgroup:         uint32(cLimits.maxComputeInvocationsPerWorkgroup),
		MaxComputeWorkgroupSizeX:                  uint32(cLimits.maxComputeWorkgroupSizeX),
		MaxComputeWorkgroupSizeY:                  uint32(cLimits.maxComputeWorkgroupSizeY),
		MaxComputeWorkgroupSizeZ:                  uint32(cLimits.maxComputeWorkgroupSizeZ),
		MaxComputeWorkgroupsPerDimension:          uint32(cLimits.maxComputeWorkgroupsPerDimension),
		MaxImmediateSize:                          uint32(cLimits.maxImmediateSize),
	}

	return limits, nil
}

// HasFeature checks if the device supports the given feature.
// Returns true if the feature is supported, false otherwise.
func (d *Device) HasFeature(feature FeatureName) bool {
	return bool(C.wgpuDeviceHasFeature(d.ref, C.WGPUFeatureName(feature)) != 0)
}

// GetFeatures returns a list of all features supported by the device.
func (d *Device) GetFeatures() []FeatureName {
	var cFeatures C.WGPUSupportedFeatures
	C.wgpuDeviceGetFeatures(d.ref, &cFeatures)
	defer C.wgpuSupportedFeaturesFreeMembers(cFeatures)

	if cFeatures.featureCount == 0 {
		return nil
	}

	features := make([]FeatureName, int(cFeatures.featureCount))

	for i, val := range unsafe.Slice(cFeatures.features, int(cFeatures.featureCount)) {
		features[i] = FeatureName(val)
	}

	return features
}

// GetAdapterInfo returns information about the adapter that created this device.
// Returns the adapter info and an error if it cannot be retrieved.
func (d *Device) GetAdapterInfo() (AdapterInfo, error) {
	var cAdapterInfo C.WGPUAdapterInfo

	status := C.wgpuDeviceGetAdapterInfo(d.ref, &cAdapterInfo)

	if statusCode(status) != statusCodeSuccess {
		return AdapterInfo{}, fmt.Errorf("error getting adapter info")
	}

	return AdapterInfo{
		Vendor:          C.GoStringN(cAdapterInfo.vendor.data, C.int(cAdapterInfo.vendor.length)),
		Architecture:    C.GoStringN(cAdapterInfo.architecture.data, C.int(cAdapterInfo.architecture.length)),
		Device:          C.GoStringN(cAdapterInfo.device.data, C.int(cAdapterInfo.device.length)),
		Description:     C.GoStringN(cAdapterInfo.description.data, C.int(cAdapterInfo.description.length)),
		BackendType:     BackendType(cAdapterInfo.backendType),
		AdapterType:     AdapterType(cAdapterInfo.adapterType),
		VendorID:        uint32(cAdapterInfo.vendorID),
		DeviceID:        uint32(cAdapterInfo.deviceID),
		SubgroupMinSize: uint32(cAdapterInfo.subgroupMinSize),
		SubgroupMaxSize: uint32(cAdapterInfo.subgroupMaxSize),
	}, nil
}

// GetQueue returns the default queue for this device.
// The queue is used to submit commands to the GPU.
func (d *Device) GetQueue() *Queue {
	return &Queue{ref: C.wgpuDeviceGetQueue(d.ref)}
}

// SetLabel sets the debug label for the device.
// This label appears in debuggers and validation layers.
func (d *Device) SetLabel(label string) {
	cLabel := toCStr(label)
	C.wgpuDeviceSetLabel(d.ref, cLabel)
}

// PushErrorScope pushes an error scope onto the device's error scope stack.
// Errors that match the filter will be captured in the scope.
func (d *Device) PushErrorScope(filter ErrorFilter) {
	C.wgpuDevicePushErrorScope(d.ref, C.WGPUErrorFilter(filter))
}

//export goPopErrorScopeCallbackHandler
func goPopErrorScopeCallbackHandler(status C.WGPUPopErrorScopeStatus, typ C.WGPUErrorType, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	fn := handle.Value().(PopErrorScopeCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}

	if popErrorScopeStatus(status) == popErrorScopeStatusSuccess {
		go fn(ErrorType(typ), msg)
	}
}

// PopErrorScope pops an error scope from the device's error scope stack and calls the callback with the result.
// The callback is called with the error type and message, or no error if the scope was empty.
func (d *Device) PopErrorScope(callback PopErrorScopeCallback) {
	handle := cgo.NewHandle(callback)
	defer handle.Delete()

	cCallbackInfo := C.WGPUPopErrorScopeCallbackInfo{
		mode:      C.WGPUCallbackMode(callbackModeAllowSpontaneous),
		callback:  C.WGPUPopErrorScopeCallback(C.cgo_callback_PopErrorScopeCallback),
		userdata1: unsafe.Pointer(handle),
	}
	C.wgpuDevicePopErrorScope(d.ref, cCallbackInfo)
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

	var err error
	callback := PopErrorScopeCallback(func(typ ErrorType, message string) {
		if typ != ErrorTypeNoError {
			err = errors.Join(fmt.Errorf("%s error: %s", typ, message))
		}
	})

	for range filters {
		d.PopErrorScope(callback)
	}

	return err
}

func toCComputePipelineDescriptor(pinner *runtime.Pinner, descriptor ComputePipelineDescriptor) C.WGPUComputePipelineDescriptor {
	cDescriptor := C.WGPUComputePipelineDescriptor{
		label: toCStr(descriptor.Label),
	}

	if descriptor.Layout != nil {
		cDescriptor.layout = descriptor.Layout.ref
	}

	cDescriptor.compute.module = descriptor.Compute.Module.ref
	cDescriptor.compute.entryPoint = toCStr(descriptor.Compute.EntryPoint)

	if count := C.size_t(len(descriptor.Compute.Constants)); count > 0 {
		constants := make([]C.WGPUConstantEntry, count)
		pinner.Pin(&constants[0])

		cDescriptor.compute.constants = (*C.WGPUConstantEntry)(unsafe.Pointer(&constants[0]))
		cDescriptor.compute.constantCount = count

		for i, c := range descriptor.Compute.Constants {
			constants[i] = C.WGPUConstantEntry{
				key:   toCStr(c.Key),
				value: C.double(c.Value),
			}
		}
	}

	return cDescriptor
}

func toCRenderPipelineDescriptor(pinner *runtime.Pinner, descriptor RenderPipelineDescriptor) C.WGPURenderPipelineDescriptor {
	cDescriptor := C.WGPURenderPipelineDescriptor{
		label: toCStr(descriptor.Label),
	}

	if descriptor.Layout != nil {
		cDescriptor.layout = descriptor.Layout.ref
	}

	if descriptor.Vertex.Module != nil {
		cDescriptor.vertex.module = descriptor.Vertex.Module.ref
	}

	cDescriptor.vertex.entryPoint = toCStr(descriptor.Vertex.EntryPoint)

	if count := C.size_t(len(descriptor.Vertex.Constants)); count > 0 {
		constants := make([]C.WGPUConstantEntry, count)
		pinner.Pin(&constants[0])

		cDescriptor.vertex.constants = (*C.WGPUConstantEntry)(unsafe.Pointer(&constants[0]))
		cDescriptor.vertex.constantCount = count

		for i, c := range descriptor.Vertex.Constants {
			constants[i] = C.WGPUConstantEntry{
				key:   toCStr(c.Key),
				value: C.double(c.Value),
			}
		}
	}

	if count := C.size_t(len(descriptor.Vertex.Buffers)); count > 0 {
		buffers := make([]C.WGPUVertexBufferLayout, count)
		pinner.Pin(&buffers[0])

		cDescriptor.vertex.buffers = (*C.WGPUVertexBufferLayout)(unsafe.Pointer(&buffers[0]))
		cDescriptor.vertex.bufferCount = count

		for i, buffer := range descriptor.Vertex.Buffers {
			buffers[i].stepMode = C.WGPUVertexStepMode(buffer.StepMode)
			buffers[i].arrayStride = C.uint64_t(buffer.ArrayStride)

			if count := C.size_t(len(buffer.Attributes)); count > 0 {
				attributes := make([]C.WGPUVertexAttribute, count)
				pinner.Pin(&attributes[0])

				buffers[i].attributes = (*C.WGPUVertexAttribute)(unsafe.Pointer(&attributes[0]))
				buffers[i].attributeCount = count

				for j, attribute := range buffer.Attributes {
					attributes[j] = C.WGPUVertexAttribute{
						format:         C.WGPUVertexFormat(attribute.Format),
						offset:         C.uint64_t(attribute.Offset),
						shaderLocation: C.uint32_t(attribute.ShaderLocation),
					}
				}

			}
		}
	}

	cDescriptor.primitive = C.WGPUPrimitiveState{
		topology:         C.WGPUPrimitiveTopology(descriptor.Primitive.Topology),
		stripIndexFormat: C.WGPUIndexFormat(descriptor.Primitive.StripIndexFormat),
		frontFace:        C.WGPUFrontFace(descriptor.Primitive.FrontFace),
		cullMode:         C.WGPUCullMode(descriptor.Primitive.CullMode),
		unclippedDepth:   toCBool(descriptor.Primitive.UnclippedDepth),
	}

	if descriptor.DepthStencil != nil {
		cDescriptor.depthStencil = &C.WGPUDepthStencilState{
			format:            C.WGPUTextureFormat(descriptor.DepthStencil.Format),
			depthWriteEnabled: C.WGPUOptionalBool(descriptor.DepthStencil.DepthWriteEnabled),
			depthCompare:      C.WGPUCompareFunction(descriptor.DepthStencil.DepthCompare),
			stencilFront: C.WGPUStencilFaceState{
				compare:     C.WGPUCompareFunction(descriptor.DepthStencil.StencilFront.Compare),
				failOp:      C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.FailOp),
				depthFailOp: C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.DepthFailOp),
				passOp:      C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.PassOp),
			},
			stencilBack: C.WGPUStencilFaceState{
				compare:     C.WGPUCompareFunction(descriptor.DepthStencil.StencilBack.Compare),
				failOp:      C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.FailOp),
				depthFailOp: C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.DepthFailOp),
				passOp:      C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.PassOp),
			},
			stencilReadMask:     C.uint32_t(descriptor.DepthStencil.StencilReadMask),
			stencilWriteMask:    C.uint32_t(descriptor.DepthStencil.StencilWriteMask),
			depthBias:           C.int32_t(descriptor.DepthStencil.DepthBias),
			depthBiasSlopeScale: C.float(descriptor.DepthStencil.DepthBiasSlopeScale),
			depthBiasClamp:      C.float(descriptor.DepthStencil.DepthBiasClamp),
		}
		pinner.Pin(cDescriptor.depthStencil)
	}

	cDescriptor.multisample = C.WGPUMultisampleState{
		count:                  C.uint32_t(descriptor.Multisample.Count),
		mask:                   C.uint32_t(descriptor.Multisample.Mask),
		alphaToCoverageEnabled: toCBool(descriptor.Multisample.AlphaToCoverageEnabled),
	}

	if descriptor.Fragment != nil {
		fragment := &C.WGPUFragmentState{}
		pinner.Pin(fragment)

		cDescriptor.fragment = fragment

		fragment.entryPoint = toCStr(descriptor.Fragment.EntryPoint)

		if descriptor.Fragment.Module != nil {
			fragment.module = descriptor.Fragment.Module.ref
		}

		// Constants
		if count := C.size_t(len(descriptor.Fragment.Constants)); count > 0 {
			constants := make([]C.WGPUConstantEntry, count)
			pinner.Pin(&constants[0])

			cDescriptor.fragment.constants = (*C.WGPUConstantEntry)(unsafe.Pointer(&constants[0]))
			cDescriptor.fragment.constantCount = count

			for i, c := range descriptor.Fragment.Constants {
				constants[i] = C.WGPUConstantEntry{
					key:   toCStr(c.Key),
					value: C.double(c.Value),
				}
			}

		}

		// Targets
		if count := C.size_t(len(descriptor.Fragment.Targets)); count > 0 {
			targets := make([]C.WGPUColorTargetState, count)
			pinner.Pin(&targets[0])

			cDescriptor.fragment.targets = (*C.WGPUColorTargetState)(unsafe.Pointer(&targets[0]))
			cDescriptor.fragment.targetCount = count

			for i, t := range descriptor.Fragment.Targets {
				targets[i].format = C.WGPUTextureFormat(t.Format)
				targets[i].writeMask = C.WGPUColorWriteMask(t.WriteMask)
				if t.Blend != nil {
					blend := C.WGPUBlendState{
						color: C.WGPUBlendComponent{
							operation: C.WGPUBlendOperation(t.Blend.Color.Operation),
							srcFactor: C.WGPUBlendFactor(t.Blend.Color.SrcFactor),
							dstFactor: C.WGPUBlendFactor(t.Blend.Color.DstFactor),
						},
						alpha: C.WGPUBlendComponent{
							operation: C.WGPUBlendOperation(t.Blend.Alpha.Operation),
							srcFactor: C.WGPUBlendFactor(t.Blend.Alpha.SrcFactor),
							dstFactor: C.WGPUBlendFactor(t.Blend.Alpha.DstFactor),
						},
					}

					targets[i].blend = &blend
					pinner.Pin(targets[i].blend)
				}
			}
		}
	}

	return cDescriptor
}
