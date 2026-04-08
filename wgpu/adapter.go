package wgpu

import (
	"fmt"
	"runtime/cgo"
	"unsafe"
)

/*
#include <stdlib.h>
#include "webgpu.h"
extern void cgo_callback_RequestDeviceCallback(WGPURequestDeviceStatus status, WGPUDevice device, WGPUStringView message, void *userData1, void *userData2);
extern void cgo_callback_DeviceLostCallback(WGPUDevice device, WGPUDeviceLostReason reason, WGPUStringView message, void *userData1, void *userData2);
extern void cgo_callback_UncapturedErrorCallback(WGPUDevice device, WGPUErrorType typ, WGPUStringView message, void *userData1, void *userData2);

*/
import "C"

type Adapter struct {
	ref uintptr
}

func (a *Adapter) GetLimits() Limits {
	limits, err := a.TryGetLimits()
	if err != nil {
		panic(err)
	}
	return limits
}

func (a *Adapter) TryGetLimits() (Limits, error) {
	cAdapter := C.WGPUAdapter(unsafe.Pointer(a.ref))

	var cLimits C.WGPULimits
	var limits Limits

	status := statusCode(C.wgpuAdapterGetLimits(cAdapter, &cLimits))
	if status == statusCodeError {
		return limits, fmt.Errorf("error getting adapter limits")
	}

	limits.MaxTextureDimension1D = uint32(cLimits.maxTextureDimension1D)
	limits.MaxTextureDimension2D = uint32(cLimits.maxTextureDimension2D)
	limits.MaxTextureDimension3D = uint32(cLimits.maxTextureDimension3D)
	limits.MaxTextureArrayLayers = uint32(cLimits.maxTextureArrayLayers)
	limits.MaxBindGroups = uint32(cLimits.maxBindGroups)
	limits.MaxBindGroupsPlusVertexBuffers = uint32(cLimits.maxBindGroupsPlusVertexBuffers)
	limits.MaxBindingsPerBindGroup = uint32(cLimits.maxBindingsPerBindGroup)
	limits.MaxDynamicUniformBuffersPerPipelineLayout = uint32(cLimits.maxDynamicUniformBuffersPerPipelineLayout)
	limits.MaxDynamicStorageBuffersPerPipelineLayout = uint32(cLimits.maxDynamicStorageBuffersPerPipelineLayout)
	limits.MaxSampledTexturesPerShaderStage = uint32(cLimits.maxSampledTexturesPerShaderStage)
	limits.MaxSamplersPerShaderStage = uint32(cLimits.maxSamplersPerShaderStage)
	limits.MaxStorageBuffersPerShaderStage = uint32(cLimits.maxStorageBuffersPerShaderStage)
	limits.MaxStorageTexturesPerShaderStage = uint32(cLimits.maxStorageTexturesPerShaderStage)
	limits.MaxUniformBuffersPerShaderStage = uint32(cLimits.maxUniformBuffersPerShaderStage)
	limits.MaxUniformBufferBindingSize = uint64(cLimits.maxUniformBufferBindingSize)
	limits.MaxStorageBufferBindingSize = uint64(cLimits.maxStorageBufferBindingSize)
	limits.MinUniformBufferOffsetAlignment = uint32(cLimits.minUniformBufferOffsetAlignment)
	limits.MinStorageBufferOffsetAlignment = uint32(cLimits.minStorageBufferOffsetAlignment)
	limits.MaxVertexBuffers = uint32(cLimits.maxVertexBuffers)
	limits.MaxBufferSize = uint64(cLimits.maxBufferSize)
	limits.MaxVertexAttributes = uint32(cLimits.maxVertexAttributes)
	limits.MaxVertexBufferArrayStride = uint32(cLimits.maxVertexBufferArrayStride)
	limits.MaxInterStageShaderVariables = uint32(cLimits.maxInterStageShaderVariables)
	limits.MaxColorAttachments = uint32(cLimits.maxColorAttachments)
	limits.MaxColorAttachmentBytesPerSample = uint32(cLimits.maxColorAttachmentBytesPerSample)
	limits.MaxComputeWorkgroupStorageSize = uint32(cLimits.maxComputeWorkgroupStorageSize)
	limits.MaxComputeInvocationsPerWorkgroup = uint32(cLimits.maxComputeInvocationsPerWorkgroup)
	limits.MaxComputeWorkgroupSizeX = uint32(cLimits.maxComputeWorkgroupSizeX)
	limits.MaxComputeWorkgroupSizeY = uint32(cLimits.maxComputeWorkgroupSizeY)
	limits.MaxComputeWorkgroupSizeZ = uint32(cLimits.maxComputeWorkgroupSizeZ)
	limits.MaxComputeWorkgroupsPerDimension = uint32(cLimits.maxComputeWorkgroupsPerDimension)
	limits.MaxImmediateSize = uint32(cLimits.maxImmediateSize)

	return limits, nil
}

func (a *Adapter) GetInfo() AdapterInfo {
	info, err := a.TryGetInfo()
	if err != nil {
		panic(err)
	}
	return info
}

func (a *Adapter) TryGetInfo() (AdapterInfo, error) {
	cAdapter := C.WGPUAdapter(unsafe.Pointer(a.ref))

	var cInfo C.WGPUAdapterInfo
	var info AdapterInfo

	status := statusCode(C.wgpuAdapterGetInfo(cAdapter, &cInfo))
	if status == statusCodeError {
		return info, fmt.Errorf("error getting adapter info")
	}

	info.Vendor = C.GoStringN(cInfo.vendor.data, C.int(cInfo.vendor.length))
	C.free(unsafe.Pointer(cInfo.vendor.data))

	info.Architecture = C.GoStringN(cInfo.architecture.data, C.int(cInfo.architecture.length))
	C.free(unsafe.Pointer(cInfo.architecture.data))

	info.Device = C.GoStringN(cInfo.device.data, C.int(cInfo.device.length))
	C.free(unsafe.Pointer(cInfo.device.data))

	info.Description = C.GoStringN(cInfo.description.data, C.int(cInfo.description.length))
	C.free(unsafe.Pointer(cInfo.description.data))

	info.BackendType = BackendType(cInfo.backendType)
	info.AdapterType = AdapterType(cInfo.adapterType)
	info.VendorID = uint32(cInfo.vendorID)
	info.DeviceID = uint32(cInfo.deviceID)
	info.SubgroupMinSize = uint32(cInfo.subgroupMinSize)
	info.SubgroupMaxSize = uint32(cInfo.subgroupMaxSize)

	return info, nil
}

func (a *Adapter) HasFeature(feature FeatureName) bool {
	cAdapter := C.WGPUAdapter(unsafe.Pointer(a.ref))
	cFeature := C.WGPUFeatureName(feature)

	return bool(C.wgpuAdapterHasFeature(cAdapter, cFeature) != 0)
}

func (a *Adapter) GetFeatures() []FeatureName {
	cAdapter := C.WGPUAdapter(unsafe.Pointer(a.ref))

	var cFeatures C.WGPUSupportedFeatures
	C.wgpuAdapterGetFeatures(cAdapter, &cFeatures)
	defer C.free(unsafe.Pointer(cFeatures.features))

	if cFeatures.featureCount == 0 {
		return nil
	}

	features := make([]FeatureName, int(cFeatures.featureCount))

	for i, val := range unsafe.Slice(cFeatures.features, int(cFeatures.featureCount)) {
		features[i] = FeatureName(val)
	}

	return features
}

//export goRequestDeviceCallbackHandler
func goRequestDeviceCallbackHandler(status C.WGPURequestDeviceStatus, device C.WGPUDevice, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	defer handle.Delete()
	fn := handle.Value().(requestDeviceCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}

	fn(
		RequestDeviceStatus(status),
		&Device{ref: uintptr(unsafe.Pointer(device))},
		msg,
	)
}

//export goDeviceLostCallbackHandler
func goDeviceLostCallbackHandler(device C.WGPUDevice, reason C.WGPUDeviceLostReason, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	defer handle.Delete()
	fn := handle.Value().(DeviceLostCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}

	fn(
		&Device{ref: uintptr(unsafe.Pointer(device))},
		DeviceLostReason(reason),
		msg,
	)
}

//export goUncapturedErrorCallbackHandler
func goUncapturedErrorCallbackHandler(device C.WGPUDevice, typ C.WGPUErrorType, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	defer handle.Delete()
	fn := handle.Value().(UncapturedErrorCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}

	fn(
		&Device{ref: uintptr(unsafe.Pointer(device))},
		ErrorType(typ),
		msg,
	)
}

func (a *Adapter) RequestDevice(descriptor *DeviceDescriptor) *Device {
	device, err := a.TryRequestDevice(descriptor)
	if err != nil {
		panic(err)
	}
	return device
}

func (a *Adapter) TryRequestDevice(descriptor *DeviceDescriptor) (*Device, error) {
	cAdapter := C.WGPUAdapter(unsafe.Pointer(a.ref))

	var pDescriptor C.WGPUDeviceDescriptor

	if descriptor != nil {
		pDescriptorlabelStr := C.CString(descriptor.Label)
		defer C.free(unsafe.Pointer(pDescriptorlabelStr))
		pDescriptor.label.data = pDescriptorlabelStr
		pDescriptor.label.length = C.size_t(len(descriptor.Label))

		featuresCount := len(descriptor.RequiredFeatures)
		if featuresCount > 0 {
			requiredFeatures := C.malloc(C.size_t(featuresCount) * C.size_t(unsafe.Sizeof(C.WGPUFeatureName(0))))
			defer C.free(requiredFeatures)

			slice := unsafe.Slice((*FeatureName)(requiredFeatures), featuresCount)
			copy(slice, descriptor.RequiredFeatures)

			pDescriptor.requiredFeatures = (*C.WGPUFeatureName)(requiredFeatures)
			pDescriptor.requiredFeatureCount = C.size_t(featuresCount)
		}

		pDescriptor.requiredLimits = (*C.WGPULimits)(C.malloc(C.size_t(unsafe.Sizeof(C.WGPULimits{}))))
		defer C.free(unsafe.Pointer(pDescriptor.requiredLimits))

		pDescriptor.requiredLimits.maxTextureDimension1D = C.uint32_t(descriptor.RequiredLimits.MaxTextureDimension1D)
		pDescriptor.requiredLimits.maxTextureDimension2D = C.uint32_t(descriptor.RequiredLimits.MaxTextureDimension2D)
		pDescriptor.requiredLimits.maxTextureDimension3D = C.uint32_t(descriptor.RequiredLimits.MaxTextureDimension3D)
		pDescriptor.requiredLimits.maxTextureArrayLayers = C.uint32_t(descriptor.RequiredLimits.MaxTextureArrayLayers)
		pDescriptor.requiredLimits.maxBindGroups = C.uint32_t(descriptor.RequiredLimits.MaxBindGroups)
		pDescriptor.requiredLimits.maxBindGroupsPlusVertexBuffers = C.uint32_t(descriptor.RequiredLimits.MaxBindGroupsPlusVertexBuffers)
		pDescriptor.requiredLimits.maxBindingsPerBindGroup = C.uint32_t(descriptor.RequiredLimits.MaxBindingsPerBindGroup)
		pDescriptor.requiredLimits.maxDynamicUniformBuffersPerPipelineLayout = C.uint32_t(descriptor.RequiredLimits.MaxDynamicUniformBuffersPerPipelineLayout)
		pDescriptor.requiredLimits.maxDynamicStorageBuffersPerPipelineLayout = C.uint32_t(descriptor.RequiredLimits.MaxDynamicStorageBuffersPerPipelineLayout)
		pDescriptor.requiredLimits.maxSampledTexturesPerShaderStage = C.uint32_t(descriptor.RequiredLimits.MaxSampledTexturesPerShaderStage)
		pDescriptor.requiredLimits.maxSamplersPerShaderStage = C.uint32_t(descriptor.RequiredLimits.MaxSamplersPerShaderStage)
		pDescriptor.requiredLimits.maxStorageBuffersPerShaderStage = C.uint32_t(descriptor.RequiredLimits.MaxStorageBuffersPerShaderStage)
		pDescriptor.requiredLimits.maxStorageTexturesPerShaderStage = C.uint32_t(descriptor.RequiredLimits.MaxStorageTexturesPerShaderStage)
		pDescriptor.requiredLimits.maxUniformBuffersPerShaderStage = C.uint32_t(descriptor.RequiredLimits.MaxUniformBuffersPerShaderStage)
		pDescriptor.requiredLimits.maxUniformBufferBindingSize = C.uint64_t(descriptor.RequiredLimits.MaxUniformBufferBindingSize)
		pDescriptor.requiredLimits.maxStorageBufferBindingSize = C.uint64_t(descriptor.RequiredLimits.MaxStorageBufferBindingSize)
		pDescriptor.requiredLimits.minUniformBufferOffsetAlignment = C.uint32_t(descriptor.RequiredLimits.MinUniformBufferOffsetAlignment)
		pDescriptor.requiredLimits.minStorageBufferOffsetAlignment = C.uint32_t(descriptor.RequiredLimits.MinStorageBufferOffsetAlignment)
		pDescriptor.requiredLimits.maxVertexBuffers = C.uint32_t(descriptor.RequiredLimits.MaxVertexBuffers)
		pDescriptor.requiredLimits.maxBufferSize = C.uint64_t(descriptor.RequiredLimits.MaxBufferSize)
		pDescriptor.requiredLimits.maxVertexAttributes = C.uint32_t(descriptor.RequiredLimits.MaxVertexAttributes)
		pDescriptor.requiredLimits.maxVertexBufferArrayStride = C.uint32_t(descriptor.RequiredLimits.MaxVertexBufferArrayStride)
		pDescriptor.requiredLimits.maxInterStageShaderVariables = C.uint32_t(descriptor.RequiredLimits.MaxInterStageShaderVariables)
		pDescriptor.requiredLimits.maxColorAttachments = C.uint32_t(descriptor.RequiredLimits.MaxColorAttachments)
		pDescriptor.requiredLimits.maxColorAttachmentBytesPerSample = C.uint32_t(descriptor.RequiredLimits.MaxColorAttachmentBytesPerSample)
		pDescriptor.requiredLimits.maxComputeWorkgroupStorageSize = C.uint32_t(descriptor.RequiredLimits.MaxComputeWorkgroupStorageSize)
		pDescriptor.requiredLimits.maxComputeInvocationsPerWorkgroup = C.uint32_t(descriptor.RequiredLimits.MaxComputeInvocationsPerWorkgroup)
		pDescriptor.requiredLimits.maxComputeWorkgroupSizeX = C.uint32_t(descriptor.RequiredLimits.MaxComputeWorkgroupSizeX)
		pDescriptor.requiredLimits.maxComputeWorkgroupSizeY = C.uint32_t(descriptor.RequiredLimits.MaxComputeWorkgroupSizeY)
		pDescriptor.requiredLimits.maxComputeWorkgroupSizeZ = C.uint32_t(descriptor.RequiredLimits.MaxComputeWorkgroupSizeZ)
		pDescriptor.requiredLimits.maxComputeWorkgroupsPerDimension = C.uint32_t(descriptor.RequiredLimits.MaxComputeWorkgroupsPerDimension)
		pDescriptor.requiredLimits.maxImmediateSize = C.uint32_t(descriptor.RequiredLimits.MaxImmediateSize)

		pDescriptordefaultQueuelabelStr := C.CString(descriptor.DefaultQueue.Label)
		defer C.free(unsafe.Pointer(pDescriptordefaultQueuelabelStr))
		pDescriptor.defaultQueue.label.data = pDescriptordefaultQueuelabelStr
		pDescriptor.defaultQueue.label.length = C.size_t(len(descriptor.DefaultQueue.Label))

		if descriptor.DeviceLostCallback != nil {
			handle := cgo.NewHandle(descriptor.DeviceLostCallback)
			pDescriptor.deviceLostCallbackInfo.mode = C.WGPUCallbackMode(callbackModeAllowProcessEvents)
			pDescriptor.deviceLostCallbackInfo.callback = C.WGPUDeviceLostCallback(C.cgo_callback_DeviceLostCallback)
			pDescriptor.deviceLostCallbackInfo.userdata1 = unsafe.Pointer(handle)
			pDescriptor.deviceLostCallbackInfo.userdata2 = nil
		}

		if descriptor.UncapturedErrorCallback != nil {
			handle := cgo.NewHandle(descriptor.UncapturedErrorCallback)
			pDescriptor.uncapturedErrorCallbackInfo.callback = C.WGPUUncapturedErrorCallback(C.cgo_callback_UncapturedErrorCallback)
			pDescriptor.uncapturedErrorCallbackInfo.userdata1 = unsafe.Pointer(handle)
			pDescriptor.uncapturedErrorCallbackInfo.userdata2 = nil
		}
	}

	var status requestAdapterStatus
	var message string
	var device *Device

	callback := func(s requestAdapterStatus, d *Device, m string) {
		status = s
		device = d
		message = m
	}

	handle := cgo.NewHandle(callback)

	var cCallbackInfo C.WGPURequestDeviceCallbackInfo
	cCallbackInfo.mode = C.WGPUCallbackMode(callbackModeAllowSpontaneous)
	cCallbackInfo.callback = C.WGPURequestDeviceCallback(C.cgo_callback_RequestDeviceCallback)
	cCallbackInfo.userdata1 = unsafe.Pointer(handle)
	cCallbackInfo.userdata2 = nil

	C.wgpuAdapterRequestDevice(cAdapter, &pDescriptor, cCallbackInfo)

	if status != requestAdapterStatusSuccess {
		return nil, fmt.Errorf("error requesting adapter: %s", message)
	}

	return device, nil
}
