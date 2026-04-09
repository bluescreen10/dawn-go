//go:build !js

package wgpu

/*
#include <stdlib.h>
#include "webgpu.h"

extern void cgo_callback_CreateComputePipelineAsyncCallback(WGPUCreatePipelineAsyncStatus status, WGPUComputePipeline pipeline, WGPUStringView message, void *userData1, void *userData2);
extern void cgo_callback_CreateRenderPipelineAsyncCallback(WGPUCreatePipelineAsyncStatus status, WGPURenderPipeline pipeline, WGPUStringView message, void *userData1, void *userData2);
extern void cgo_callback_PopErrorScopeCallback(WGPUPopErrorScopeStatus status, WGPUErrorType typ, WGPUStringView message, void *userData1, void *userData2);
*/
import "C"

import (
	"errors"
	"fmt"
	"runtime/cgo"
	"unsafe"
)

type Device struct {
	ref     uintptr
	handles []cgo.Handle
}

func (d *Device) CreateBindGroup(descriptor BindGroupDescriptor) *BindGroup {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUBindGroupDescriptor

	if descriptor.Label != "" {
		cDescriptor.label.data = C.CString(descriptor.Label)
		cDescriptor.label.length = C.size_t(len(descriptor.Label))
		defer C.free(unsafe.Pointer(cDescriptor.label.data))
	}

	if descriptor.Layout != nil {
		cDescriptor.layout = C.WGPUBindGroupLayout(unsafe.Pointer(descriptor.Layout.ref))
	}

	entriesCount := len(descriptor.Entries)
	if entriesCount > 0 {
		entries := C.malloc(C.size_t(entriesCount) * C.size_t(unsafe.Sizeof(C.WGPUBindGroupEntry{})))
		slice := unsafe.Slice((*C.WGPUBindGroupEntry)(entries), entriesCount)
		defer C.free(unsafe.Pointer(entries))

		for i, e := range descriptor.Entries {
			slice[i].binding = C.uint32_t(e.Binding)
			slice[i].offset = C.uint64_t(e.Offset)
			slice[i].size = C.uint64_t(e.Size)

			if e.Buffer != nil {
				slice[i].buffer = C.WGPUBuffer(unsafe.Pointer(e.Buffer.ref))
			}

			if e.Sampler != nil {
				slice[i].sampler = C.WGPUSampler(unsafe.Pointer(e.Sampler.ref))
			}

			if e.TextureView != nil {
				slice[i].sampler = C.WGPUSampler(unsafe.Pointer(e.TextureView.ref))
			}
		}

		cDescriptor.entries = (*C.WGPUBindGroupEntry)(entries)
		cDescriptor.entryCount = C.size_t(entriesCount)
	}

	ptr := unsafe.Pointer(C.wgpuDeviceCreateBindGroup(cDevice, &cDescriptor))
	return &BindGroup{ref: uintptr(ptr)}
}

func (d *Device) CreateBindGroupLayout(descriptor BindGroupLayoutDescriptor) BindGroupLayout {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUBindGroupLayoutDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor.entryCount = C.size_t(descriptor.EntryCount)
	cDescriptor.entries.binding = C.uint32_t(descriptor.Entries.Binding)
	cDescriptor.entries.visibility = C.WGPUShaderStage(descriptor.Entries.Visibility)
	cDescriptor.entries.bindingArraySize = C.uint32_t(descriptor.Entries.BindingArraySize)
	cDescriptor.entries.buffer._type = C.WGPUBufferBindingType(descriptor.Entries.Buffer.Type)
	cDescriptor.entries.buffer.hasDynamicOffset = boolToWGPUBool(descriptor.Entries.Buffer.HasDynamicOffset)
	cDescriptor.entries.buffer.minBindingSize = C.uint64_t(descriptor.Entries.Buffer.MinBindingSize)
	cDescriptor.entries.sampler._type = C.WGPUSamplerBindingType(descriptor.Entries.Sampler.Type)
	cDescriptor.entries.texture.sampleType = C.WGPUTextureSampleType(descriptor.Entries.Texture.SampleType)
	cDescriptor.entries.texture.viewDimension = C.WGPUTextureViewDimension(descriptor.Entries.Texture.ViewDimension)
	cDescriptor.entries.texture.multisampled = boolToWGPUBool(descriptor.Entries.Texture.Multisampled)
	cDescriptor.entries.storageTexture.access = C.WGPUStorageTextureAccess(descriptor.Entries.StorageTexture.Access)
	cDescriptor.entries.storageTexture.format = C.WGPUTextureFormat(descriptor.Entries.StorageTexture.Format)
	cDescriptor.entries.storageTexture.viewDimension = C.WGPUTextureViewDimension(descriptor.Entries.StorageTexture.ViewDimension)
	// Call and return
	return BindGroupLayout{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateBindGroupLayout(cDevice, &cDescriptor)))}
}

func (d *Device) CreateBufferInit(descriptor BufferInitDescriptor) *Buffer {
	buffer, err := d.TryCreateBufferInit(descriptor)
	if err != nil {
		panic(err)
	}
	return buffer
}

func (d *Device) TryCreateBufferInit(descriptor BufferInitDescriptor) (*Buffer, error) {
	buffer, err := d.TryCreateBuffer(BufferDescriptor{
		Label: descriptor.Label,
		Size:  uint64(len(descriptor.Contents)),
		Usage: descriptor.Usage | BufferUsageCopyDst,
	})
	if err != nil {
		return nil, err
	}

	d.GetQueue().WriteBuffer(buffer, 0, descriptor.Contents)

	return buffer, nil
}

func (d *Device) CreateBuffer(descriptor BufferDescriptor) *Buffer {
	buffer, err := d.TryCreateBuffer(descriptor)
	if err != nil {
		panic(err)
	}
	return buffer
}

func (d *Device) TryCreateBuffer(descriptor BufferDescriptor) (*Buffer, error) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUBufferDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor.usage = C.WGPUBufferUsage(descriptor.Usage)
	cDescriptor.size = C.uint64_t(descriptor.Size)
	cDescriptor.mappedAtCreation = boolToWGPUBool(descriptor.MappedAtCreation)

	ptr := unsafe.Pointer(C.wgpuDeviceCreateBuffer(cDevice, &cDescriptor))

	if ptr == nil {
		var err error
		cCallbackInfo := makeErrorCallback(&err)
		C.wgpuDevicePopErrorScope(cDevice, cCallbackInfo)
		return nil, fmt.Errorf("error creating buffer: %v", err)
	}

	return &Buffer{ref: uintptr(ptr)}, nil
}

func (d *Device) CreateCommandEncoder(descriptor *CommandEncoderDescriptor) *CommandEncoder {
	commandEncoder, err := d.TryCreateCommandEncoder(descriptor)
	if err != nil {
		panic(err)
	}
	return commandEncoder
}

func (d *Device) TryCreateCommandEncoder(descriptor *CommandEncoderDescriptor) (*CommandEncoder, error) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var pDescriptor C.WGPUCommandEncoderDescriptor
	if descriptor != nil && descriptor.Label != "" {
		pDescriptor.label.data = C.CString(descriptor.Label)
		pDescriptor.label.length = C.size_t(len(descriptor.Label))
		defer C.free(unsafe.Pointer(C.CString(descriptor.Label)))
	}

	ptr := unsafe.Pointer(C.wgpuDeviceCreateCommandEncoder(cDevice, &pDescriptor))
	if ptr == nil {
		var err error
		cCallbackInfo := makeErrorCallback(&err)
		C.wgpuDevicePopErrorScope(cDevice, cCallbackInfo)
		fmt.Errorf("error creating command encoder: %v", err)
	}

	return &CommandEncoder{ref: uintptr(ptr)}, nil
}

func (d *Device) CreateComputePipeline(descriptor ComputePipelineDescriptor) ComputePipeline {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUComputePipelineDescriptor

	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))

	cDescriptor.layout = C.WGPUPipelineLayout(unsafe.Pointer(descriptor.Layout.ref))
	cDescriptor.compute.module = C.WGPUShaderModule(unsafe.Pointer(descriptor.Compute.Module.ref))

	cDescriptorcomputeentryPointStr := C.CString(descriptor.Compute.EntryPoint)
	defer C.free(unsafe.Pointer(cDescriptorcomputeentryPointStr))
	cDescriptor.compute.entryPoint.data = cDescriptorcomputeentryPointStr
	cDescriptor.compute.entryPoint.length = C.size_t(len(descriptor.Compute.EntryPoint))
	cDescriptor.compute.constantCount = C.size_t(descriptor.Compute.ConstantCount)
	cDescriptorcomputeconstantskeyStr := C.CString(descriptor.Compute.Constants.Key)
	defer C.free(unsafe.Pointer(cDescriptorcomputeconstantskeyStr))
	cDescriptor.compute.constants.key.data = cDescriptorcomputeconstantskeyStr
	cDescriptor.compute.constants.key.length = C.size_t(len(descriptor.Compute.Constants.Key))
	cDescriptor.compute.constants.value = C.double(descriptor.Compute.Constants.Value)

	return ComputePipeline{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateComputePipeline(cDevice, &cDescriptor)))}
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

	fn(
		CreatePipelineAsyncStatus(status),
		&ComputePipeline{ref: uintptr(unsafe.Pointer(pipeline))},
		msg,
	)
}
func (d *Device) CreateComputePipelineAsync(descriptor ComputePipelineDescriptor, callback CreateComputePipelineAsyncCallback) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUComputePipelineDescriptor
	cDescriptor.layout = C.WGPUPipelineLayout(unsafe.Pointer(descriptor.Layout.ref))
	cDescriptor.compute.module = C.WGPUShaderModule(unsafe.Pointer(descriptor.Compute.Module.ref))
	cDescriptor.compute.constantCount = C.size_t(descriptor.Compute.ConstantCount)
	cDescriptorcomputeconstantskeyStr := C.CString(descriptor.Compute.Constants.Key)
	cDescriptor.compute.constants.value = C.double(descriptor.Compute.Constants.Value)

	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))

	cDescriptorcomputeentryPointStr := C.CString(descriptor.Compute.EntryPoint)
	defer C.free(unsafe.Pointer(cDescriptorcomputeentryPointStr))
	cDescriptor.compute.entryPoint.data = cDescriptorcomputeentryPointStr
	cDescriptor.compute.entryPoint.length = C.size_t(len(descriptor.Compute.EntryPoint))

	defer C.free(unsafe.Pointer(cDescriptorcomputeconstantskeyStr))
	cDescriptor.compute.constants.key.data = cDescriptorcomputeconstantskeyStr
	cDescriptor.compute.constants.key.length = C.size_t(len(descriptor.Compute.Constants.Key))

	// Setup callback
	handle := cgo.NewHandle(callback)
	var cCallbackInfo C.WGPUCreateComputePipelineAsyncCallbackInfo
	cCallbackInfo.mode = C.WGPUCallbackMode(callbackModeAllowProcessEvents)
	cCallbackInfo.callback = C.WGPUCreateComputePipelineAsyncCallback(C.cgo_callback_CreateComputePipelineAsyncCallback)
	cCallbackInfo.userdata1 = unsafe.Pointer(handle)
	cCallbackInfo.userdata2 = nil

	C.wgpuDeviceCreateComputePipelineAsync(cDevice, &cDescriptor, cCallbackInfo)
}

func (d *Device) CreatePipelineLayout(descriptor PipelineLayoutDescriptor) PipelineLayout {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUPipelineLayoutDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor.bindGroupLayoutCount = C.size_t(descriptor.BindGroupLayoutCount)
	cDescriptor.bindGroupLayouts = (*C.WGPUBindGroupLayout)(unsafe.Pointer(descriptor.BindGroupLayouts.ref))
	cDescriptor.immediateSize = C.uint32_t(descriptor.ImmediateSize)
	// Call and return
	return PipelineLayout{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreatePipelineLayout(cDevice, &cDescriptor)))}
}

func (d *Device) CreateQuerySet(descriptor QuerySetDescriptor) QuerySet {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUQuerySetDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor._type = C.WGPUQueryType(descriptor.Type)
	cDescriptor.count = C.uint32_t(descriptor.Count)
	// Call and return
	return QuerySet{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateQuerySet(cDevice, &cDescriptor)))}
}

//export goCreateRenderPipelineAsyncCallbackHandler
func goCreateRenderPipelineAsyncCallbackHandler(status C.WGPUCreatePipelineAsyncStatus, pipeline C.WGPURenderPipeline, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	devRef := uintptr(userData2)
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
	fn(
		CreatePipelineAsyncStatus(status),
		&RenderPipeline{ref: uintptr(unsafe.Pointer(pipeline)), devRef: devRef},
		msg,
	)
}
func (d *Device) CreateRenderPipelineAsync(descriptor RenderPipelineDescriptor, callback CreateRenderPipelineAsyncCallback) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPURenderPipelineDescriptor
	cDescriptor.label.data = C.CString(descriptor.Label)
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	defer C.free(unsafe.Pointer(cDescriptor.label.data))

	cDescriptor.layout = C.WGPUPipelineLayout(unsafe.Pointer(descriptor.Layout.ref))
	cDescriptor.vertex.module = C.WGPUShaderModule(unsafe.Pointer(descriptor.Vertex.Module.ref))

	if descriptor.Vertex.EntryPoint != "" {
		cDescriptor.vertex.entryPoint.data = C.CString(descriptor.Vertex.EntryPoint)
		cDescriptor.vertex.entryPoint.length = C.size_t(len(descriptor.Vertex.EntryPoint))
		defer C.free(unsafe.Pointer(cDescriptor.vertex.entryPoint.data))
	}

	constansCount := len(descriptor.Vertex.Constants)
	if constansCount > 0 {
		constants := C.malloc(C.size_t(constansCount) * C.size_t(unsafe.Sizeof(C.WGPUConstantEntry{})))
		slice := unsafe.Slice((*C.WGPUConstantEntry)(constants), constansCount)
		defer C.free(unsafe.Pointer(constants))

		for i, c := range descriptor.Vertex.Constants {
			slice[i].value = C.double(c.Value)
			slice[i].key = C.WGPUStringView{
				data:   C.CString(c.Key),
				length: C.size_t(len(c.Key)),
			}
			defer C.free(unsafe.Pointer(slice[i].key.data))
		}

		cDescriptor.vertex.constants = (*C.WGPUConstantEntry)(constants)
		cDescriptor.vertex.constantCount = C.size_t(constansCount)
	}

	buffersCount := len(descriptor.Vertex.Buffers)
	if buffersCount > 0 {
		buffers := C.malloc(C.size_t(buffersCount) * C.size_t(unsafe.Sizeof(C.WGPUVertexBufferLayout{})))
		slice := unsafe.Slice((*C.WGPUVertexBufferLayout)(buffers), constansCount)

		for i, buffer := range descriptor.Vertex.Buffers {
			slice[i].stepMode = C.WGPUVertexStepMode(buffer.StepMode)
			slice[i].arrayStride = C.uint64_t(buffer.ArrayStride)

			attributesCount := len(buffer.Attributes)
			if attributesCount > 0 {
				attributes := C.malloc(C.size_t(attributesCount) * C.size_t(unsafe.Sizeof(C.WGPUVertexAttribute{})))
				attributesSlice := unsafe.Slice((*C.WGPUVertexAttribute)(attributes), attributesCount)

				for j, attribute := range buffer.Attributes {
					attributesSlice[j].format = C.WGPUVertexFormat(attribute.Format)
					attributesSlice[j].offset = C.uint64_t(attribute.Offset)
					attributesSlice[j].shaderLocation = C.uint32_t(attribute.ShaderLocation)
				}

				slice[i].attributes = (*C.WGPUVertexAttribute)(attributes)
				slice[i].attributeCount = C.size_t(attributesCount)
			}
		}

		cDescriptor.vertex.buffers = (*C.WGPUVertexBufferLayout)(buffers)
		cDescriptor.vertex.bufferCount = C.size_t(buffersCount)
	}

	cDescriptor.primitive.topology = C.WGPUPrimitiveTopology(descriptor.Primitive.Topology)
	cDescriptor.primitive.stripIndexFormat = C.WGPUIndexFormat(descriptor.Primitive.StripIndexFormat)
	cDescriptor.primitive.frontFace = C.WGPUFrontFace(descriptor.Primitive.FrontFace)
	cDescriptor.primitive.cullMode = C.WGPUCullMode(descriptor.Primitive.CullMode)
	cDescriptor.primitive.unclippedDepth = boolToWGPUBool(descriptor.Primitive.UnclippedDepth)

	if descriptor.DepthStencil != nil {
		cDescriptor.depthStencil.format = C.WGPUTextureFormat(descriptor.DepthStencil.Format)
		cDescriptor.depthStencil.depthWriteEnabled = C.WGPUOptionalBool(descriptor.DepthStencil.DepthWriteEnabled)
		cDescriptor.depthStencil.depthCompare = C.WGPUCompareFunction(descriptor.DepthStencil.DepthCompare)
		cDescriptor.depthStencil.stencilFront.compare = C.WGPUCompareFunction(descriptor.DepthStencil.StencilFront.Compare)
		cDescriptor.depthStencil.stencilFront.failOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.FailOp)
		cDescriptor.depthStencil.stencilFront.depthFailOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.DepthFailOp)
		cDescriptor.depthStencil.stencilFront.passOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.PassOp)
		cDescriptor.depthStencil.stencilBack.compare = C.WGPUCompareFunction(descriptor.DepthStencil.StencilBack.Compare)
		cDescriptor.depthStencil.stencilBack.failOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.FailOp)
		cDescriptor.depthStencil.stencilBack.depthFailOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.DepthFailOp)
		cDescriptor.depthStencil.stencilBack.passOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.PassOp)
		cDescriptor.depthStencil.stencilReadMask = C.uint32_t(descriptor.DepthStencil.StencilReadMask)
		cDescriptor.depthStencil.stencilWriteMask = C.uint32_t(descriptor.DepthStencil.StencilWriteMask)
		cDescriptor.depthStencil.depthBias = C.int32_t(descriptor.DepthStencil.DepthBias)
		cDescriptor.depthStencil.depthBiasSlopeScale = C.float(descriptor.DepthStencil.DepthBiasSlopeScale)
		cDescriptor.depthStencil.depthBiasClamp = C.float(descriptor.DepthStencil.DepthBiasClamp)
	}

	cDescriptor.multisample.count = C.uint32_t(descriptor.Multisample.Count)
	cDescriptor.multisample.mask = C.uint32_t(descriptor.Multisample.Mask)
	cDescriptor.multisample.alphaToCoverageEnabled = boolToWGPUBool(descriptor.Multisample.AlphaToCoverageEnabled)

	if descriptor.Fragment != nil {
		cDescriptor.fragment.module = C.WGPUShaderModule(unsafe.Pointer(descriptor.Fragment.Module.ref))
		cDescriptorfragmententryPointStr := C.CString(descriptor.Fragment.EntryPoint)
		defer C.free(unsafe.Pointer(cDescriptorfragmententryPointStr))
		cDescriptor.fragment.entryPoint.data = cDescriptorfragmententryPointStr
		cDescriptor.fragment.entryPoint.length = C.size_t(len(descriptor.Fragment.EntryPoint))

		constansCount := len(descriptor.Fragment.Constants)
		if constansCount > 0 {
			constants := C.malloc(C.size_t(constansCount) * C.size_t(unsafe.Sizeof(C.WGPUConstantEntry{})))
			slice := unsafe.Slice((*C.WGPUConstantEntry)(constants), constansCount)
			defer C.free(unsafe.Pointer(constants))

			for i, c := range descriptor.Fragment.Constants {
				slice[i].value = C.double(c.Value)
				slice[i].key = C.WGPUStringView{
					data:   C.CString(c.Key),
					length: C.size_t(len(c.Key)),
				}
				defer C.free(unsafe.Pointer(slice[i].key.data))
			}

			cDescriptor.fragment.constants = (*C.WGPUConstantEntry)(constants)
			cDescriptor.fragment.constantCount = C.size_t(constansCount)
		}

		targetsCount := len(descriptor.Fragment.Targets)
		if targetsCount > 0 {
			targets := C.malloc(C.size_t(targetsCount) * C.size_t(unsafe.Sizeof(C.WGPUColorTargetState{})))
			slice := unsafe.Slice((*C.WGPUColorTargetState)(targets), targetsCount)
			defer C.free(unsafe.Pointer(targets))

			for i, t := range descriptor.Fragment.Targets {
				slice[i].format = C.WGPUTextureFormat(t.Format)
				slice[i].writeMask = C.WGPUColorWriteMask(t.WriteMask)
				if t.Blend != nil {
					slice[i].blend = &C.WGPUBlendState{
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
				}
			}

			cDescriptor.fragment.targets = (*C.WGPUColorTargetState)(targets)
			cDescriptor.fragment.targetCount = C.size_t(targetsCount)
		}
	}

	handle := cgo.NewHandle(callback)

	var cCallbackInfo C.WGPUCreateRenderPipelineAsyncCallbackInfo
	cCallbackInfo.mode = C.WGPUCallbackMode(callbackModeAllowProcessEvents)
	cCallbackInfo.callback = C.WGPUCreateRenderPipelineAsyncCallback(C.cgo_callback_CreateRenderPipelineAsyncCallback)
	cCallbackInfo.userdata1 = unsafe.Pointer(handle)
	cCallbackInfo.userdata2 = unsafe.Pointer(cDevice)

	C.wgpuDeviceCreateRenderPipelineAsync(cDevice, &cDescriptor, cCallbackInfo)
}

func (d *Device) CreateRenderBundleEncoder(descriptor RenderBundleEncoderDescriptor) RenderBundleEncoder {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPURenderBundleEncoderDescriptor
	cDescriptorlabelStr := C.CString(descriptor.Label)
	defer C.free(unsafe.Pointer(cDescriptorlabelStr))
	cDescriptor.label.data = cDescriptorlabelStr
	cDescriptor.label.length = C.size_t(len(descriptor.Label))
	cDescriptor.depthStencilFormat = C.WGPUTextureFormat(descriptor.DepthStencilFormat)
	cDescriptor.sampleCount = C.uint32_t(descriptor.SampleCount)
	cDescriptor.depthReadOnly = boolToWGPUBool(descriptor.DepthReadOnly)
	cDescriptor.stencilReadOnly = boolToWGPUBool(descriptor.StencilReadOnly)

	colorFormatCount := len(descriptor.ColorFormats)
	if colorFormatCount > 0 {
		colorFormats := C.malloc(C.size_t(colorFormatCount) * C.size_t(unsafe.Sizeof(C.WGPUTextureFormat(0))))
		defer C.free(colorFormats)

		slice := unsafe.Slice((*TextureFormat)(colorFormats), colorFormatCount)
		copy(slice, descriptor.ColorFormats)

		cDescriptor.colorFormats = (*C.WGPUTextureFormat)(colorFormats)
		cDescriptor.colorFormatCount = C.size_t(colorFormatCount)
	}

	return RenderBundleEncoder{ref: uintptr(unsafe.Pointer(C.wgpuDeviceCreateRenderBundleEncoder(cDevice, &cDescriptor)))}
}

func (d *Device) CreateRenderPipeline(descriptor RenderPipelineDescriptor) *RenderPipeline {
	renderPipeline, err := d.TryCreateRenderPipeline(descriptor)
	if err != nil {
		panic(err)
	}
	return renderPipeline
}

func (d *Device) TryCreateRenderPipeline(descriptor RenderPipelineDescriptor) (*RenderPipeline, error) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPURenderPipelineDescriptor

	if descriptor.Label != "" {
		cDescriptor.label.length = C.size_t(len(descriptor.Label))
		cDescriptor.label.data = C.CString(descriptor.Label)
		defer C.free(unsafe.Pointer(cDescriptor.label.data))
	}

	if descriptor.Layout != nil {
		cDescriptor.layout = C.WGPUPipelineLayout(unsafe.Pointer(descriptor.Layout.ref))
	}

	if descriptor.Vertex.Module != nil {
		cDescriptor.vertex.module = C.WGPUShaderModule(unsafe.Pointer(descriptor.Vertex.Module.ref))
	}

	if descriptor.Vertex.EntryPoint != "" {
		cDescriptor.vertex.entryPoint.data = C.CString(descriptor.Vertex.EntryPoint)
		cDescriptor.vertex.entryPoint.length = C.size_t(len(descriptor.Vertex.EntryPoint))
		defer C.free(unsafe.Pointer(cDescriptor.vertex.entryPoint.data))
	}

	constansCount := len(descriptor.Vertex.Constants)
	if constansCount > 0 {
		constants := C.malloc(C.size_t(constansCount) * C.size_t(unsafe.Sizeof(C.WGPUConstantEntry{})))
		slice := unsafe.Slice((*C.WGPUConstantEntry)(constants), constansCount)

		for i, c := range descriptor.Vertex.Constants {
			slice[i].value = C.double(c.Value)
			slice[i].key = C.WGPUStringView{
				data:   C.CString(c.Key),
				length: C.size_t(len(c.Key)),
			}
			defer C.free(unsafe.Pointer(slice[i].key.data))
		}

		cDescriptor.vertex.constants = (*C.WGPUConstantEntry)(constants)
		cDescriptor.vertex.constantCount = C.size_t(constansCount)
	}

	buffersCount := len(descriptor.Vertex.Buffers)
	if buffersCount > 0 {
		buffers := C.malloc(C.size_t(buffersCount) * C.size_t(unsafe.Sizeof(C.WGPUVertexBufferLayout{})))
		slice := unsafe.Slice((*C.WGPUVertexBufferLayout)(buffers), buffersCount)

		for i, buffer := range descriptor.Vertex.Buffers {
			slice[i].stepMode = C.WGPUVertexStepMode(buffer.StepMode)
			slice[i].arrayStride = C.uint64_t(buffer.ArrayStride)

			attributesCount := len(buffer.Attributes)
			if attributesCount > 0 {
				attributes := C.malloc(C.size_t(attributesCount) * C.size_t(unsafe.Sizeof(C.WGPUVertexAttribute{})))
				attributesSlice := unsafe.Slice((*C.WGPUVertexAttribute)(attributes), attributesCount)

				for j, attribute := range buffer.Attributes {
					attributesSlice[j].format = C.WGPUVertexFormat(attribute.Format)
					attributesSlice[j].offset = C.uint64_t(attribute.Offset)
					attributesSlice[j].shaderLocation = C.uint32_t(attribute.ShaderLocation)
				}

				slice[i].attributes = (*C.WGPUVertexAttribute)(attributes)
				slice[i].attributeCount = C.size_t(attributesCount)
			}
		}

		cDescriptor.vertex.buffers = (*C.WGPUVertexBufferLayout)(buffers)
		cDescriptor.vertex.bufferCount = C.size_t(buffersCount)
	}

	cDescriptor.primitive.topology = C.WGPUPrimitiveTopology(descriptor.Primitive.Topology)
	cDescriptor.primitive.stripIndexFormat = C.WGPUIndexFormat(descriptor.Primitive.StripIndexFormat)
	cDescriptor.primitive.frontFace = C.WGPUFrontFace(descriptor.Primitive.FrontFace)
	cDescriptor.primitive.cullMode = C.WGPUCullMode(descriptor.Primitive.CullMode)
	cDescriptor.primitive.unclippedDepth = boolToWGPUBool(descriptor.Primitive.UnclippedDepth)

	if descriptor.DepthStencil != nil {
		cDescriptor.depthStencil.format = C.WGPUTextureFormat(descriptor.DepthStencil.Format)
		cDescriptor.depthStencil.depthWriteEnabled = C.WGPUOptionalBool(descriptor.DepthStencil.DepthWriteEnabled)
		cDescriptor.depthStencil.depthCompare = C.WGPUCompareFunction(descriptor.DepthStencil.DepthCompare)
		cDescriptor.depthStencil.stencilFront.compare = C.WGPUCompareFunction(descriptor.DepthStencil.StencilFront.Compare)
		cDescriptor.depthStencil.stencilFront.failOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.FailOp)
		cDescriptor.depthStencil.stencilFront.depthFailOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.DepthFailOp)
		cDescriptor.depthStencil.stencilFront.passOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilFront.PassOp)
		cDescriptor.depthStencil.stencilBack.compare = C.WGPUCompareFunction(descriptor.DepthStencil.StencilBack.Compare)
		cDescriptor.depthStencil.stencilBack.failOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.FailOp)
		cDescriptor.depthStencil.stencilBack.depthFailOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.DepthFailOp)
		cDescriptor.depthStencil.stencilBack.passOp = C.WGPUStencilOperation(descriptor.DepthStencil.StencilBack.PassOp)
		cDescriptor.depthStencil.stencilReadMask = C.uint32_t(descriptor.DepthStencil.StencilReadMask)
		cDescriptor.depthStencil.stencilWriteMask = C.uint32_t(descriptor.DepthStencil.StencilWriteMask)
		cDescriptor.depthStencil.depthBias = C.int32_t(descriptor.DepthStencil.DepthBias)
		cDescriptor.depthStencil.depthBiasSlopeScale = C.float(descriptor.DepthStencil.DepthBiasSlopeScale)
		cDescriptor.depthStencil.depthBiasClamp = C.float(descriptor.DepthStencil.DepthBiasClamp)
	}

	cDescriptor.multisample.count = C.uint32_t(descriptor.Multisample.Count)
	cDescriptor.multisample.mask = C.uint32_t(descriptor.Multisample.Mask)
	cDescriptor.multisample.alphaToCoverageEnabled = boolToWGPUBool(descriptor.Multisample.AlphaToCoverageEnabled)

	if descriptor.Fragment != nil {
		cDescriptor.fragment = (*C.WGPUFragmentState)(C.malloc(C.size_t(unsafe.Sizeof(C.WGPUFragmentState{}))))
		defer C.free(unsafe.Pointer(cDescriptor.fragment))

		if descriptor.Fragment.Module != nil {
			cDescriptor.fragment.module = C.WGPUShaderModule(unsafe.Pointer(descriptor.Fragment.Module.ref))
		}

		cDescriptor.fragment.entryPoint.length = C.size_t(len(descriptor.Fragment.EntryPoint))
		cDescriptor.fragment.entryPoint.data = C.CString(descriptor.Fragment.EntryPoint)
		defer C.free(unsafe.Pointer(cDescriptor.fragment.entryPoint.data))

		constansCount := len(descriptor.Fragment.Constants)
		if constansCount > 0 {
			constants := C.malloc(C.size_t(constansCount) * C.size_t(unsafe.Sizeof(C.WGPUConstantEntry{})))
			defer C.free(unsafe.Pointer(constants))

			slice := unsafe.Slice((*C.WGPUConstantEntry)(constants), constansCount)

			for i, c := range descriptor.Fragment.Constants {
				slice[i].value = C.double(c.Value)
				slice[i].key = C.WGPUStringView{
					data:   C.CString(c.Key),
					length: C.size_t(len(c.Key)),
				}
				defer C.free(unsafe.Pointer(slice[i].key.data))
			}

			cDescriptor.fragment.constants = (*C.WGPUConstantEntry)(constants)
			cDescriptor.fragment.constantCount = C.size_t(constansCount)
		}

		targetsCount := len(descriptor.Fragment.Targets)
		if targetsCount > 0 {
			targets := C.malloc(C.size_t(targetsCount) * C.size_t(unsafe.Sizeof(C.WGPUColorTargetState{})))
			slice := unsafe.Slice((*C.WGPUColorTargetState)(targets), targetsCount)
			defer C.free(unsafe.Pointer(targets))

			for i, t := range descriptor.Fragment.Targets {
				slice[i].format = C.WGPUTextureFormat(t.Format)
				slice[i].writeMask = C.WGPUColorWriteMask(t.WriteMask)
				if t.Blend != nil {
					blendPtr := C.malloc(C.size_t(unsafe.Sizeof(C.WGPUBlendState{})))
					blend := (*C.WGPUBlendState)(blendPtr)
					blend.color.operation = C.WGPUBlendOperation(t.Blend.Color.Operation)
					blend.color.srcFactor = C.WGPUBlendFactor(t.Blend.Color.SrcFactor)
					blend.color.dstFactor = C.WGPUBlendFactor(t.Blend.Color.DstFactor)
					blend.alpha.operation = C.WGPUBlendOperation(t.Blend.Alpha.Operation)
					blend.alpha.srcFactor = C.WGPUBlendFactor(t.Blend.Alpha.SrcFactor)
					blend.alpha.dstFactor = C.WGPUBlendFactor(t.Blend.Alpha.DstFactor)
					slice[i].blend = blend
				}
			}

			cDescriptor.fragment.targets = (*C.WGPUColorTargetState)(targets)
			cDescriptor.fragment.targetCount = C.size_t(targetsCount)
		}
	}

	ptr := unsafe.Pointer(C.wgpuDeviceCreateRenderPipeline(cDevice, &cDescriptor))

	if ptr == nil {
		var err error
		cCallbackInfo := makeErrorCallback(&err)
		C.wgpuDevicePopErrorScope(cDevice, cCallbackInfo)
		return nil, fmt.Errorf("error creating pipeline: %v", err)
	}

	return &RenderPipeline{ref: uintptr(ptr), devRef: d.ref}, nil
}

func (d *Device) CreateSampler(descriptor *SamplerDescriptor) *Sampler {
	sampler, err := d.TryCreateSampler(descriptor)
	if err != nil {
		panic(err)
	}
	return sampler
}

func (d *Device) TryCreateSampler(descriptor *SamplerDescriptor) (*Sampler, error) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var pDescriptor C.WGPUSamplerDescriptor
	if descriptor != nil {
		pDescriptor.label.data = C.CString(descriptor.Label)
		pDescriptor.label.length = C.size_t(len(descriptor.Label))
		defer C.free(unsafe.Pointer(pDescriptor.label.data))

		pDescriptor.addressModeU = C.WGPUAddressMode(descriptor.AddressModeU)
		pDescriptor.addressModeV = C.WGPUAddressMode(descriptor.AddressModeV)
		pDescriptor.addressModeW = C.WGPUAddressMode(descriptor.AddressModeW)
		pDescriptor.magFilter = C.WGPUFilterMode(descriptor.MagFilter)
		pDescriptor.minFilter = C.WGPUFilterMode(descriptor.MinFilter)
		pDescriptor.mipmapFilter = C.WGPUMipmapFilterMode(descriptor.MipmapFilter)
		pDescriptor.lodMinClamp = C.float(descriptor.LodMinClamp)
		pDescriptor.lodMaxClamp = C.float(descriptor.LodMaxClamp)
		pDescriptor.compare = C.WGPUCompareFunction(descriptor.Compare)
		pDescriptor.maxAnisotropy = C.uint16_t(descriptor.MaxAnisotropy)
	}

	ptr := unsafe.Pointer(C.wgpuDeviceCreateSampler(cDevice, &pDescriptor))

	if ptr == nil {
		var err error
		cCallbackInfo := makeErrorCallback(&err)
		C.wgpuDevicePopErrorScope(cDevice, cCallbackInfo)
		return nil, fmt.Errorf("error creating sampler: %v", err)
	}

	return &Sampler{ref: uintptr(ptr)}, nil
}

func (d *Device) CreateShaderModule(descriptor ShaderModuleDescriptor) *ShaderModule {
	shaderModule, err := d.TryCreateShaderModule(descriptor)
	if err != nil {
		panic(err)
	}
	return shaderModule
}

func (d *Device) TryCreateShaderModule(descriptor ShaderModuleDescriptor) (*ShaderModule, error) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cDescriptor C.WGPUShaderModuleDescriptor

	if descriptor.Label != "" {
		cDescriptor.label.data = C.CString(descriptor.Label)
		cDescriptor.label.length = C.size_t(len(descriptor.Label))
		defer C.free(unsafe.Pointer(cDescriptor.label.data))
	}

	if descriptor.WGSLSource != nil {
		source := C.malloc(C.size_t(unsafe.Sizeof(C.WGPUShaderSourceWGSL{})))
		defer C.free(source)

		wgslSource := (*C.WGPUShaderSourceWGSL)(source)
		wgslSource.chain.next = nil
		wgslSource.chain.sType = C.WGPUSType_ShaderSourceWGSL

		wgslSource.code.length = C.size_t(len(descriptor.WGSLSource.Code))
		wgslSource.code.data = C.CString(descriptor.WGSLSource.Code)
		defer C.free(unsafe.Pointer(wgslSource.code.data))

		cDescriptor.nextInChain = (*C.WGPUChainedStruct)(source)
	}

	if descriptor.SPIRVSource != nil {
		source := C.malloc(C.size_t(unsafe.Sizeof(C.WGPUShaderSourceSPIRV{})))
		defer C.free(source)

		spirvSource := (*C.WGPUShaderSourceSPIRV)(source)
		spirvSource.chain.next = nil
		spirvSource.chain.sType = C.WGPUSType_ShaderSourceSPIRV

		codeSize := len(descriptor.SPIRVSource.Code)
		if codeSize > 0 {
			code := C.malloc(C.size_t(codeSize) * C.size_t(unsafe.Sizeof(uint32(0))))
			slice := unsafe.Slice((*uint32)(code), codeSize)
			copy(slice, descriptor.SPIRVSource.Code)
			spirvSource.code = (*C.uint32_t)(code)
			spirvSource.codeSize = C.uint32_t(codeSize)
			defer C.free(code)
		}
	}

	ptr := unsafe.Pointer(C.wgpuDeviceCreateShaderModule(cDevice, &cDescriptor))

	if ptr == nil {
		var err error
		cCallbackInfo := makeErrorCallback(&err)
		C.wgpuDevicePopErrorScope(cDevice, cCallbackInfo)
		return nil, fmt.Errorf("error creating shader module: %v", err)
	}

	return &ShaderModule{ref: uintptr(ptr)}, nil
}

func (d *Device) CreateTexture(descriptor *TextureDescriptor) *Texture {
	texture, err := d.TryCreateTexture(descriptor)
	if err != nil {
		panic(err)
	}
	return texture
}

func (d *Device) TryCreateTexture(descriptor *TextureDescriptor) (*Texture, error) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

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

		if descriptor.Label != "" {
			cDescriptor.label.length = C.size_t(len(descriptor.Label))
			cDescriptor.label.data = C.CString(descriptor.Label)
			defer C.free(unsafe.Pointer(&cDescriptor.label.data))
		}

		viewFormatsCount := len(descriptor.ViewFormats)
		if viewFormatsCount > 0 {
			viewFormats := C.malloc(C.size_t(viewFormatsCount) * C.size_t(unsafe.Sizeof(C.WGPUTextureFormat(0))))
			defer C.free(viewFormats)

			slice := unsafe.Slice((*TextureFormat)(viewFormats), viewFormatsCount)
			copy(slice, descriptor.ViewFormats)

			cDescriptor.viewFormats = (*C.WGPUTextureFormat)(viewFormats)
			cDescriptor.viewFormatCount = C.size_t(viewFormatsCount)
		}
	}

	var err error
	cCallbackInfo := makeErrorCallback(&err)
	C.wgpuDevicePushErrorScope(cDevice, C.WGPUErrorFilter(ErrorFilterValidation))
	ptr := unsafe.Pointer(C.wgpuDeviceCreateTexture(cDevice, cDescriptor))
	C.wgpuDevicePopErrorScope(cDevice, cCallbackInfo)

	if err != nil {
		C.wgpuTextureRelease(C.WGPUTexture(ptr))
		return nil, err
	}

	return &Texture{ref: uintptr(ptr), devRef: d.ref}, nil
}

func (d *Device) Release() {
	for _, h := range d.handles {
		h.Delete()
	}
	d.handles = nil

	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	C.wgpuDeviceRelease(cDevice)
}

func (d *Device) Destroy() {
	for _, h := range d.handles {
		h.Delete()
	}
	d.handles = nil

	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	C.wgpuDeviceDestroy(cDevice)
}

func (d *Device) GetLimits(limits Limits) statusCode {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cLimits C.WGPULimits
	cLimits.maxTextureDimension1D = C.uint32_t(limits.MaxTextureDimension1D)
	cLimits.maxTextureDimension2D = C.uint32_t(limits.MaxTextureDimension2D)
	cLimits.maxTextureDimension3D = C.uint32_t(limits.MaxTextureDimension3D)
	cLimits.maxTextureArrayLayers = C.uint32_t(limits.MaxTextureArrayLayers)
	cLimits.maxBindGroups = C.uint32_t(limits.MaxBindGroups)
	cLimits.maxBindGroupsPlusVertexBuffers = C.uint32_t(limits.MaxBindGroupsPlusVertexBuffers)
	cLimits.maxBindingsPerBindGroup = C.uint32_t(limits.MaxBindingsPerBindGroup)
	cLimits.maxDynamicUniformBuffersPerPipelineLayout = C.uint32_t(limits.MaxDynamicUniformBuffersPerPipelineLayout)
	cLimits.maxDynamicStorageBuffersPerPipelineLayout = C.uint32_t(limits.MaxDynamicStorageBuffersPerPipelineLayout)
	cLimits.maxSampledTexturesPerShaderStage = C.uint32_t(limits.MaxSampledTexturesPerShaderStage)
	cLimits.maxSamplersPerShaderStage = C.uint32_t(limits.MaxSamplersPerShaderStage)
	cLimits.maxStorageBuffersPerShaderStage = C.uint32_t(limits.MaxStorageBuffersPerShaderStage)
	cLimits.maxStorageTexturesPerShaderStage = C.uint32_t(limits.MaxStorageTexturesPerShaderStage)
	cLimits.maxUniformBuffersPerShaderStage = C.uint32_t(limits.MaxUniformBuffersPerShaderStage)
	cLimits.maxUniformBufferBindingSize = C.uint64_t(limits.MaxUniformBufferBindingSize)
	cLimits.maxStorageBufferBindingSize = C.uint64_t(limits.MaxStorageBufferBindingSize)
	cLimits.minUniformBufferOffsetAlignment = C.uint32_t(limits.MinUniformBufferOffsetAlignment)
	cLimits.minStorageBufferOffsetAlignment = C.uint32_t(limits.MinStorageBufferOffsetAlignment)
	cLimits.maxVertexBuffers = C.uint32_t(limits.MaxVertexBuffers)
	cLimits.maxBufferSize = C.uint64_t(limits.MaxBufferSize)
	cLimits.maxVertexAttributes = C.uint32_t(limits.MaxVertexAttributes)
	cLimits.maxVertexBufferArrayStride = C.uint32_t(limits.MaxVertexBufferArrayStride)
	cLimits.maxInterStageShaderVariables = C.uint32_t(limits.MaxInterStageShaderVariables)
	cLimits.maxColorAttachments = C.uint32_t(limits.MaxColorAttachments)
	cLimits.maxColorAttachmentBytesPerSample = C.uint32_t(limits.MaxColorAttachmentBytesPerSample)
	cLimits.maxComputeWorkgroupStorageSize = C.uint32_t(limits.MaxComputeWorkgroupStorageSize)
	cLimits.maxComputeInvocationsPerWorkgroup = C.uint32_t(limits.MaxComputeInvocationsPerWorkgroup)
	cLimits.maxComputeWorkgroupSizeX = C.uint32_t(limits.MaxComputeWorkgroupSizeX)
	cLimits.maxComputeWorkgroupSizeY = C.uint32_t(limits.MaxComputeWorkgroupSizeY)
	cLimits.maxComputeWorkgroupSizeZ = C.uint32_t(limits.MaxComputeWorkgroupSizeZ)
	cLimits.maxComputeWorkgroupsPerDimension = C.uint32_t(limits.MaxComputeWorkgroupsPerDimension)
	cLimits.maxImmediateSize = C.uint32_t(limits.MaxImmediateSize)
	// Call and return
	return statusCode(C.wgpuDeviceGetLimits(cDevice, &cLimits))
}

func (d *Device) HasFeature(feature FeatureName) bool {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	cFeature := C.WGPUFeatureName(feature)
	// Call and return
	return bool(C.wgpuDeviceHasFeature(cDevice, cFeature) != 0)
}

func (d *Device) GetFeatures() []FeatureName {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cFeatures C.WGPUSupportedFeatures
	//TODO:here
	// cFeatures.featureCount = C.size_t(features.FeatureCount)
	// cFeatures.features = C.WGPUFeatureName(features.Features)
	// Call and return
	C.wgpuDeviceGetFeatures(cDevice, &cFeatures)

	return nil
}

func (d *Device) GetAdapterInfo() (AdapterInfo, error) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	var cAdapterInfo C.WGPUAdapterInfo

	status := C.wgpuDeviceGetAdapterInfo(cDevice, &cAdapterInfo)

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

func (d *Device) GetQueue() *Queue {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))
	return &Queue{ref: uintptr(unsafe.Pointer(C.wgpuDeviceGetQueue(cDevice)))}
}

func (d *Device) SetLabel(label string) {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	cLabel := C.WGPUStringView{
		data:   C.CString(label),
		length: C.size_t(len(label)),
	}
	defer C.free(unsafe.Pointer(cLabel.data))

	C.wgpuDeviceSetLabel(cDevice, cLabel)
}

func (d *Device) Try(fn func()) error {
	cDevice := C.WGPUDevice(unsafe.Pointer(d.ref))

	//TODO: Make scopes optional
	// Push error scopes
	C.wgpuDevicePushErrorScope(cDevice, C.WGPUErrorFilter(ErrorFilterValidation))
	C.wgpuDevicePushErrorScope(cDevice, C.WGPUErrorFilter(ErrorFilterOutOfMemory))
	C.wgpuDevicePushErrorScope(cDevice, C.WGPUErrorFilter(ErrorFilterInternal))

	// Call function
	fn()

	var err error
	callback := popErrorScopeCallback(func(_ popErrorScopeStatus, typ ErrorType, message string) {
		if typ != ErrorTypeNoError {
			err = errors.Join(fmt.Errorf("%s error: %s", typ, message))
		}
	})

	handle := cgo.NewHandle(callback)
	defer handle.Delete()

	cCallbackInfo := C.WGPUPopErrorScopeCallbackInfo{
		mode:      C.WGPUCallbackMode(callbackModeAllowSpontaneous),
		callback:  C.WGPUPopErrorScopeCallback(C.cgo_callback_PopErrorScopeCallback),
		userdata1: unsafe.Pointer(handle),
		userdata2: nil,
	}

	// Pop Error Scopes
	C.wgpuDevicePopErrorScope(cDevice, cCallbackInfo)
	C.wgpuDevicePopErrorScope(cDevice, cCallbackInfo)
	C.wgpuDevicePopErrorScope(cDevice, cCallbackInfo)
	return err
}

//export goPopErrorScopeCallbackHandler
func goPopErrorScopeCallbackHandler(status C.WGPUPopErrorScopeStatus, typ C.WGPUErrorType, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	fn := handle.Value().(popErrorScopeCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}
	fn(
		popErrorScopeStatus(status),
		ErrorType(typ),
		msg,
	)
}
