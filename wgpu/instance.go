//go:build !js

package wgpu

/*
#include "webgpu.h"

extern void cgo_callback_RequestAdapterCallback(WGPURequestAdapterStatus status, WGPUAdapter adapter, WGPUStringView message, void *userData1, void *userData2);
*/
import "C"
import (
	"fmt"
	"runtime"
	"runtime/cgo"
	"time"
	"unsafe"
)

// Instance is a WebGPU instance, which serves as the entry point for all WebGPU operations.
// It manages the underlying GPU backend and is used to request adapters.
type Instance struct {
	ref C.WGPUInstance
}

// CreateSurface creates a surface for rendering to a window.
// The surface is configured for the specified descriptor, which defines how the surface should be presented.
func (i *Instance) CreateSurface(descriptor SurfaceDescriptor) *Surface {
	var pinner runtime.Pinner

	cDescriptor := C.WGPUSurfaceDescriptor{
		label: toCStr(descriptor.Label),
	}

	if descriptor.MetalLayer != nil {
		metalSource := &C.WGPUSurfaceSourceMetalLayer{
			chain: C.WGPUChainedStruct{
				next:  nil,
				sType: C.WGPUSType_SurfaceSourceMetalLayer,
			},
			layer: descriptor.MetalLayer.Layer,
		}

		pinner.Pin(metalSource)
		cDescriptor.nextInChain = (*C.WGPUChainedStruct)(unsafe.Pointer(metalSource))
	} else if descriptor.WindowsHWND != nil {
		windowsSource := &C.WGPUSurfaceSourceWindowsHWND{
			chain: C.WGPUChainedStruct{
				next:  nil,
				sType: C.WGPUSType_SurfaceSourceWindowsHWND,
			},
			hinstance: descriptor.WindowsHWND.Hinstance,
			hwnd:      descriptor.WindowsHWND.Hwnd,
		}

		pinner.Pin(windowsSource)
		cDescriptor.nextInChain = (*C.WGPUChainedStruct)(unsafe.Pointer(windowsSource))
	} else if descriptor.WaylandSurface != nil {
		waylandSource := &C.WGPUSurfaceSourceWaylandSurface{
			chain: C.WGPUChainedStruct{
				next:  nil,
				sType: C.WGPUSType_SurfaceSourceWaylandSurface,
			},
			display: descriptor.WaylandSurface.Display,
			surface: descriptor.WaylandSurface.Surface,
		}

		pinner.Pin(waylandSource)
		cDescriptor.nextInChain = (*C.WGPUChainedStruct)(unsafe.Pointer(waylandSource))
	} else if descriptor.XlibWindow != nil {
		xlibSource := &C.WGPUSurfaceSourceXlibWindow{
			chain: C.WGPUChainedStruct{
				next:  nil,
				sType: C.WGPUSType_SurfaceSourceXlibWindow,
			},
			display: descriptor.XlibWindow.Display,
			window:  C.uint64_t(descriptor.XlibWindow.Window),
		}

		pinner.Pin(xlibSource)
		cDescriptor.nextInChain = (*C.WGPUChainedStruct)(unsafe.Pointer(&xlibSource))
	}

	return &Surface{ref: C.wgpuInstanceCreateSurface(i.ref, &cDescriptor)}
}

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

// RequestAdapter requests a GPU adapter from the instance based on the provided options.
// Returns the adapter and an error if the request fails.
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

// HasWGSLLanguageFeature checks if the instance supports the given WGSL language feature.
// Returns true if the feature is supported, false otherwise.
func (i *Instance) HasWGSLLanguageFeature(feature WGSLLanguageFeatureName) bool {
	cFeature := C.WGPUWGSLLanguageFeatureName(feature)
	return bool(C.wgpuInstanceHasWGSLLanguageFeature(i.ref, cFeature) != 0)
}

// GetWGSLLanguageFeatures returns all WGSL language features supported by the instance.
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

// ProcessEvents processes any pending events in the instance, such as callbacks and device lost notifications.
// This must be called periodically to ensure callbacks are executed.
func (i *Instance) ProcessEvents() {
	C.wgpuInstanceProcessEvents(i.ref)
}

// Release releases the instance and all associated resources.
// After calling this method, the instance should no longer be used.
func (i *Instance) Release() {
	C.wgpuInstanceRelease(i.ref)
}

// Wait blocks the current thread until the specified future completes or the timeout expires.
// Returns an error if the wait fails or times out.
func (i *Instance) Wait(future Future, timeout time.Duration) error {
	return i.WaitAny([]Future{future}, timeout)
}

// WaitAny blocks the current thread until any of the specified futures completes or the timeout expires.
// Returns an error if the wait fails or times out.
func (i *Instance) WaitAny(futures []Future, timeout time.Duration) error {
	cTimeoutNs := C.uint64_t(timeout.Nanoseconds())

	cCount := C.size_t(len(futures))
	var cFutureWaitInfo *C.WGPUFutureWaitInfo
	if cCount > 0 {
		futureWaitInfo := make([]C.WGPUFutureWaitInfo, cCount)
		cFutureWaitInfo = (*C.WGPUFutureWaitInfo)(unsafe.Pointer(&futureWaitInfo[0]))

		for i, f := range futures {
			futureWaitInfo[i].future = C.WGPUFuture{id: C.uint64_t(f.id)}
		}
	}

	status := C.wgpuInstanceWaitAny(i.ref, cCount, cFutureWaitInfo, cTimeoutNs)

	switch waitStatus(status) {
	case waitStatusTimedOut:
		return fmt.Errorf("waitAny status timeout error")
	case waitStatusError:
		return fmt.Errorf("waitAny status error")
	default:
		return nil
	}
}

// CreateInstance creates a new WebGPU instance with optional descriptor.
// The instance is the entry point for all WebGPU operations.
func CreateInstance(descriptor *InstanceDescriptor) *Instance {
	var cDescriptor *C.WGPUInstanceDescriptor

	if descriptor != nil {
		var pinner runtime.Pinner
		defer pinner.Unpin()

		cDescriptor = &C.WGPUInstanceDescriptor{}

		if count := len(descriptor.RequiredFeatures); count > 0 {
			pinner.Pin(&descriptor.RequiredFeatures[0])
			cDescriptor.requiredFeatureCount = C.size_t(count)
			cDescriptor.requiredFeatures = (*C.WGPUInstanceFeatureName)(&descriptor.RequiredFeatures[0])
		}

		if descriptor.RequiredLimits != nil {
			limits := &C.WGPUInstanceLimits{timedWaitAnyMaxCount: C.size_t(descriptor.RequiredLimits.TimedWaitAnyMaxCount)}
			pinner.Pin(limits)
		}
	}

	return &Instance{ref: C.wgpuCreateInstance(cDescriptor)}
}

// GetInstanceFeatures returns a list of all instance-level features supported by the WebGPU implementation.
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

// GetInstanceLimits returns the instance-level limits supported by the WebGPU implementation.
// Returns the limits and an error if they cannot be retrieved.
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

// HasInstanceFeature checks if the given instance feature is supported by the WebGPU implementation.
// Returns true if the feature is supported, false otherwise.
func HasInstanceFeature(feature InstanceFeatureName) bool {
	return bool(C.wgpuHasInstanceFeature(C.WGPUInstanceFeatureName(feature)) != 0)
}
