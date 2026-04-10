//go:build !js

package wgpu

/*
#include <stdlib.h>
#include "webgpu.h"

extern void cgo_callback_RequestAdapterCallback(WGPURequestAdapterStatus status, WGPUAdapter adapter, WGPUStringView message, void *userData1, void *userData2);
*/
import "C"
import (
	"fmt"
	"runtime/cgo"
	"unsafe"
)

type Instance struct {
	ref C.WGPUInstance
}

func (i *Instance) CreateSurface(descriptor SurfaceDescriptor) *Surface {

	cDescriptor := C.WGPUSurfaceDescriptor{
		label: toCStr(descriptor.Label),
	}

	if descriptor.MetalLayer != nil {
		metalSourcePtr := C.malloc(C.size_t(unsafe.Sizeof(C.WGPUSurfaceSourceMetalLayer{})))
		metalSource := (*C.WGPUSurfaceSourceMetalLayer)(metalSourcePtr)
		metalSource.chain.next = nil
		metalSource.chain.sType = C.WGPUSType_SurfaceSourceMetalLayer
		metalSource.layer = descriptor.MetalLayer.Layer
		cDescriptor.nextInChain = (*C.WGPUChainedStruct)(metalSourcePtr)
		defer C.free(metalSourcePtr)
	}

	return &Surface{ref: C.wgpuInstanceCreateSurface(i.ref, &cDescriptor)}
}

// Not needed
// func (i *Instance) ProcessEvents() {
// 	C.wgpuInstanceProcessEvents(i.ref)
// }

// Not needed
// func (i *Instance) WaitAny(futureCount int, futures *FutureWaitInfo, timeoutNS uint64) WaitStatus {
// 	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))

// 	cFutureCount := C.size_t(futureCount)
// 	if futures != nil {
// 		var pFutures C.WGPUFutureWaitInfo
// 		pFutures.future.id = C.uint64_t(futures.Future.Id)
// 		pFutures.completed = toCBool(futures.Completed)
// 	}
// 	cTimeoutNS := C.uint64_t(timeoutNS)
// 	// Call and return
// 	return WaitStatus(C.wgpuInstanceWaitAny(cInstance, cFutureCount, pFutures, cTimeoutNS))
// }

//export goRequestAdapterCallbackHandler
func goRequestAdapterCallbackHandler(status C.WGPURequestAdapterStatus, adapter C.WGPUAdapter, message C.WGPUStringView, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	fn := handle.Value().(requestAdapterCallback)

	var msg string
	if message.data != nil && message.length > 0 {
		msg = C.GoStringN(message.data, C.int(message.length))
	}
	fn(
		requestAdapterStatus(status),
		&Adapter{ref: adapter},
		msg,
	)
}

func (i *Instance) RequestAdapter(options *RequestAdapterOptions) (*Adapter, error) {

	var cOptions *C.WGPURequestAdapterOptions
	if options != nil {
		cOptions = &C.WGPURequestAdapterOptions{
			featureLevel:         C.WGPUFeatureLevel(options.FeatureLevel),
			powerPreference:      C.WGPUPowerPreference(options.PowerPreference),
			forceFallbackAdapter: toCBool(options.ForceFallbackAdapter),
			backendType:          C.WGPUBackendType(options.BackendType),
			compatibleSurface:    C.WGPUSurface(unsafe.Pointer(options.CompatibleSurface.ref)),
		}
	}

	var status requestAdapterStatus
	var adapter *Adapter
	var message string

	callback := requestAdapterCallback(func(s requestAdapterStatus, a *Adapter, m string) {
		status = s
		adapter = a
		message = m
	})

	handle := cgo.NewHandle(callback)
	defer handle.Delete()

	cCallbackInfo := C.WGPURequestAdapterCallbackInfo{
		mode:      C.WGPUCallbackMode(callbackModeAllowSpontaneous),
		callback:  C.WGPURequestAdapterCallback(C.cgo_callback_RequestAdapterCallback),
		userdata1: unsafe.Pointer(handle),
	}

	C.wgpuInstanceRequestAdapter(i.ref, cOptions, cCallbackInfo)

	if status != requestAdapterStatusSuccess {
		return nil, fmt.Errorf("error request adapter: %s", message)
	}

	return adapter, nil
}

func (i *Instance) HasWGSLLanguageFeature(feature WGSLLanguageFeatureName) bool {
	cFeature := C.WGPUWGSLLanguageFeatureName(feature)
	return bool(C.wgpuInstanceHasWGSLLanguageFeature(i.ref, cFeature) != 0)
}

func (i *Instance) GetWGSLLanguageFeatures() []WGSLLanguageFeatureName {
	var cFeatures C.WGPUSupportedWGSLLanguageFeatures
	C.wgpuInstanceGetWGSLLanguageFeatures(i.ref, &cFeatures)
	defer C.wgpuSupportedWGSLLanguageFeaturesFreeMembers(cFeatures)

	count := cFeatures.featureCount
	if count == 0 {
		return nil
	}

	features := make([]WGSLLanguageFeatureName, count)
	slice := unsafe.Slice((*C.WGPUWGSLLanguageFeatureName)(cFeatures.features), count)

	for i, f := range slice {
		features[i] = WGSLLanguageFeatureName(f)
	}

	return features
}

func (i *Instance) Release() {
	C.wgpuInstanceRelease(i.ref)
}

func CreateInstance(descriptor *InstanceDescriptor) Instance {
	var cDescriptor *C.WGPUInstanceDescriptor
	if descriptor != nil {
		cDescriptor := &C.WGPUInstanceDescriptor{}

		if count := len(descriptor.RequiredFeatures); count > 0 {
			cDescriptor.requiredFeatureCount = C.size_t(count)
			cDescriptor.requiredFeatures = (*C.WGPUInstanceFeatureName)(&descriptor.RequiredFeatures[0])
		}

		if descriptor.RequiredLimits != nil {
			cDescriptor.requiredLimits.timedWaitAnyMaxCount = C.size_t(descriptor.RequiredLimits.TimedWaitAnyMaxCount)
		}
	}

	return Instance{ref: C.wgpuCreateInstance(cDescriptor)}
}

func GetInstanceFeatures() []InstanceFeatureName {
	var cFeatures C.WGPUSupportedInstanceFeatures
	C.wgpuGetInstanceFeatures(&cFeatures)
	defer C.wgpuSupportedInstanceFeaturesFreeMembers(cFeatures)

	features := make([]InstanceFeatureName, cFeatures.featureCount)
	slice := unsafe.Slice((*C.WGPUInstanceFeatureName)(unsafe.Pointer(cFeatures.features)), cFeatures.featureCount)

	for i, f := range slice {
		features[i] = InstanceFeatureName(f)
	}

	return features
}

func GetInstanceLimits() (InstanceLimits, error) {
	var cLimits C.WGPUInstanceLimits

	status := C.wgpuGetInstanceLimits(&cLimits)

	if statusCode(status) != statusCodeSuccess {
		return InstanceLimits{}, fmt.Errorf("error getting instance limits")
	}

	limits := InstanceLimits{
		TimedWaitAnyMaxCount: int(cLimits.timedWaitAnyMaxCount),
	}

	return limits, nil
}

// func GetProcAddress(procName string) uintptr {
// 	cProcNameStr := C.CString(procName)
// 	defer C.free(unsafe.Pointer(cProcNameStr))
// 	var cProcName C.WGPUStringView
// 	cProcName.data = cProcNameStr
// 	cProcName.length = C.size_t(len(procName))
// 	// Call and return
// 	return uintptr(C.wgpuGetProcAddress(cProcName))
// }

func HasInstanceFeature(feature InstanceFeatureName) bool {
	cFeature := C.WGPUInstanceFeatureName(feature)
	return bool(C.wgpuHasInstanceFeature(cFeature) != 0)
}
