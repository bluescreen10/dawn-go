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
	ref C.WGPUAdapter
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
	status := C.wgpuAdapterGetLimits(cAdapter, &cLimits)

	if statusCode(status) != statusCodeSuccess {
		return Limits{}, fmt.Errorf("error getting adapter limits")
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

	status := C.wgpuAdapterGetInfo(cAdapter, &cInfo)
	defer C.wgpuAdapterInfoFreeMembers(cInfo)

	if statusCode(status) != statusCodeSuccess {
		return AdapterInfo{}, fmt.Errorf("error getting adapter info")
	}

	info := AdapterInfo{
		Vendor:          C.GoStringN(cInfo.vendor.data, C.int(cInfo.vendor.length)),
		Architecture:    C.GoStringN(cInfo.architecture.data, C.int(cInfo.architecture.length)),
		Device:          C.GoStringN(cInfo.device.data, C.int(cInfo.device.length)),
		Description:     C.GoStringN(cInfo.description.data, C.int(cInfo.description.length)),
		BackendType:     BackendType(cInfo.backendType),
		AdapterType:     AdapterType(cInfo.adapterType),
		VendorID:        uint32(cInfo.vendorID),
		DeviceID:        uint32(cInfo.deviceID),
		SubgroupMinSize: uint32(cInfo.subgroupMinSize),
		SubgroupMaxSize: uint32(cInfo.subgroupMaxSize),
	}

	return info, nil
}

func (a *Adapter) HasFeature(feature FeatureName) bool {
	cFeature := C.WGPUFeatureName(feature)
	return bool(C.wgpuAdapterHasFeature(a.ref, cFeature) != 0)
}

func (a *Adapter) GetFeatures() []FeatureName {

	var cFeatures C.WGPUSupportedFeatures
	C.wgpuAdapterGetFeatures(a.ref, &cFeatures)
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
		requestDeviceStatus(status),
		&Device{ref: device},
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
		&Device{ref: device},
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
	fn := handle.Value().(UncapturedErrorCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}

	fn(
		&Device{ref: device},
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
	var cDescriptor C.WGPUDeviceDescriptor
	var handles []cgo.Handle

	if descriptor != nil {
		cDescriptor.label = toCStr(descriptor.Label)

		featuresCount := len(descriptor.RequiredFeatures)
		if featuresCount > 0 {
			requiredFeatures := C.malloc(C.size_t(featuresCount) * C.size_t(unsafe.Sizeof(C.WGPUFeatureName(0))))
			defer C.free(requiredFeatures)

			slice := unsafe.Slice((*FeatureName)(requiredFeatures), featuresCount)
			copy(slice, descriptor.RequiredFeatures)

			cDescriptor.requiredFeatures = (*C.WGPUFeatureName)(requiredFeatures)
			cDescriptor.requiredFeatureCount = C.size_t(featuresCount)
		}

		if descriptor.RequiredLimits != nil {
			cDescriptor.requiredLimits = &C.WGPULimits{
				maxTextureDimension1D:                     C.uint32_t(descriptor.RequiredLimits.MaxTextureDimension1D),
				maxTextureDimension2D:                     C.uint32_t(descriptor.RequiredLimits.MaxTextureDimension2D),
				maxTextureDimension3D:                     C.uint32_t(descriptor.RequiredLimits.MaxTextureDimension3D),
				maxTextureArrayLayers:                     C.uint32_t(descriptor.RequiredLimits.MaxTextureArrayLayers),
				maxBindGroups:                             C.uint32_t(descriptor.RequiredLimits.MaxBindGroups),
				maxBindGroupsPlusVertexBuffers:            C.uint32_t(descriptor.RequiredLimits.MaxBindGroupsPlusVertexBuffers),
				maxBindingsPerBindGroup:                   C.uint32_t(descriptor.RequiredLimits.MaxBindingsPerBindGroup),
				maxDynamicUniformBuffersPerPipelineLayout: C.uint32_t(descriptor.RequiredLimits.MaxDynamicUniformBuffersPerPipelineLayout),
				maxDynamicStorageBuffersPerPipelineLayout: C.uint32_t(descriptor.RequiredLimits.MaxDynamicStorageBuffersPerPipelineLayout),
				maxSampledTexturesPerShaderStage:          C.uint32_t(descriptor.RequiredLimits.MaxSampledTexturesPerShaderStage),
				maxSamplersPerShaderStage:                 C.uint32_t(descriptor.RequiredLimits.MaxSamplersPerShaderStage),
				maxStorageBuffersPerShaderStage:           C.uint32_t(descriptor.RequiredLimits.MaxStorageBuffersPerShaderStage),
				maxStorageTexturesPerShaderStage:          C.uint32_t(descriptor.RequiredLimits.MaxStorageTexturesPerShaderStage),
				maxUniformBuffersPerShaderStage:           C.uint32_t(descriptor.RequiredLimits.MaxUniformBuffersPerShaderStage),
				maxUniformBufferBindingSize:               C.uint64_t(descriptor.RequiredLimits.MaxUniformBufferBindingSize),
				maxStorageBufferBindingSize:               C.uint64_t(descriptor.RequiredLimits.MaxStorageBufferBindingSize),
				minUniformBufferOffsetAlignment:           C.uint32_t(descriptor.RequiredLimits.MinUniformBufferOffsetAlignment),
				minStorageBufferOffsetAlignment:           C.uint32_t(descriptor.RequiredLimits.MinStorageBufferOffsetAlignment),
				maxVertexBuffers:                          C.uint32_t(descriptor.RequiredLimits.MaxVertexBuffers),
				maxBufferSize:                             C.uint64_t(descriptor.RequiredLimits.MaxBufferSize),
				maxVertexAttributes:                       C.uint32_t(descriptor.RequiredLimits.MaxVertexAttributes),
				maxVertexBufferArrayStride:                C.uint32_t(descriptor.RequiredLimits.MaxVertexBufferArrayStride),
				maxInterStageShaderVariables:              C.uint32_t(descriptor.RequiredLimits.MaxInterStageShaderVariables),
				maxColorAttachments:                       C.uint32_t(descriptor.RequiredLimits.MaxColorAttachments),
				maxColorAttachmentBytesPerSample:          C.uint32_t(descriptor.RequiredLimits.MaxColorAttachmentBytesPerSample),
				maxComputeWorkgroupStorageSize:            C.uint32_t(descriptor.RequiredLimits.MaxComputeWorkgroupStorageSize),
				maxComputeInvocationsPerWorkgroup:         C.uint32_t(descriptor.RequiredLimits.MaxComputeInvocationsPerWorkgroup),
				maxComputeWorkgroupSizeX:                  C.uint32_t(descriptor.RequiredLimits.MaxComputeWorkgroupSizeX),
				maxComputeWorkgroupSizeY:                  C.uint32_t(descriptor.RequiredLimits.MaxComputeWorkgroupSizeY),
				maxComputeWorkgroupSizeZ:                  C.uint32_t(descriptor.RequiredLimits.MaxComputeWorkgroupSizeZ),
				maxComputeWorkgroupsPerDimension:          C.uint32_t(descriptor.RequiredLimits.MaxComputeWorkgroupsPerDimension),
				maxImmediateSize:                          C.uint32_t(descriptor.RequiredLimits.MaxImmediateSize),
			}
		}

		cDescriptor.defaultQueue.label = toCStr(descriptor.DefaultQueue.Label)

		if descriptor.DeviceLostCallback != nil {
			handle := cgo.NewHandle(descriptor.DeviceLostCallback)
			handles = append(handles, handle)
			cDescriptor.deviceLostCallbackInfo = C.WGPUDeviceLostCallbackInfo{
				mode:      C.WGPUCallbackMode(callbackModeAllowProcessEvents),
				callback:  C.WGPUDeviceLostCallback(C.cgo_callback_DeviceLostCallback),
				userdata1: unsafe.Pointer(handle),
			}
		}

		if descriptor.UncapturedErrorCallback != nil {
			handle := cgo.NewHandle(descriptor.UncapturedErrorCallback)
			handles = append(handles, handle)
			cDescriptor.uncapturedErrorCallbackInfo = C.WGPUUncapturedErrorCallbackInfo{
				callback:  C.WGPUUncapturedErrorCallback(C.cgo_callback_UncapturedErrorCallback),
				userdata1: unsafe.Pointer(handle),
			}
		}
	}

	var status requestDeviceStatus
	var message string
	var device *Device

	callback := requestDeviceCallback(func(s requestDeviceStatus, d *Device, m string) {
		status = s
		device = d
		message = m
	})

	handle := cgo.NewHandle(callback)

	cCallbackInfo := C.WGPURequestDeviceCallbackInfo{
		mode:      C.WGPUCallbackMode(callbackModeAllowSpontaneous),
		callback:  C.WGPURequestDeviceCallback(C.cgo_callback_RequestDeviceCallback),
		userdata1: unsafe.Pointer(handle),
	}

	C.wgpuAdapterRequestDevice(a.ref, &cDescriptor, cCallbackInfo)

	if status != requestDeviceStatusSuccess {
		return nil, fmt.Errorf("error requesting adapter: %s", message)
	}

	device.handles = handles

	return device, nil
}

func (a *Adapter) Release() {
	C.wgpuAdapterRelease(a.ref)
}
