package main

import (
	"fmt"

	"github.com/bluescreen10/dawn-go/wgpu"
)

func main() {
	instance := wgpu.CreateInstance(nil)
	adapter := instance.RequestAdapter(nil)

	fmt.Println(adapter)
	//instance.WaitAny([]wgpu.FutureWaitInfo{{Future: future}}, 0)
	fmt.Println("bye")
	// var wg sync.WaitGroup
	// wg.Add(1)
	// wg.Wait()
}

// func (i *Instance) RequestAdapter(options *RequestAdapterOptions, callbackInfo RequestAdapterCallbackInfo) Future {
// 	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))
// 	// Convert options to C.WGPURequestAdapterOptions
// 	var pOptions *C.WGPURequestAdapterOptions
// 	if options != nil {
// 		var cOptions C.WGPURequestAdapterOptions
// 		cOptions.forceFallbackAdapter = boolToWGPUBool(options.ForceFallbackAdapter)
// 		if options.CompatibleSurface != nil {
// 			cOptions.compatibleSurface = C.WGPUSurface(unsafe.Pointer(options.CompatibleSurface.ref))
// 		}
// 		pOptions = &cOptions
// 	}

// 	handle := cgo.NewHandle(callbackInfo.Callback)

// 	// Convert callbackInfo to C.WGPURequestAdapterCallbackInfo
// 	var cCallbackInfo C.WGPURequestAdapterCallbackInfo
// 	cCallbackInfo.mode = C.WGPUCallbackMode(callbackInfo.Mode)
// 	cCallbackInfo.callback = C.WGPURequestAdapterCallback(C.cgo_callback_RequestAdapterCallback)
// 	cCallbackInfo.userdata1 = unsafe.Pointer(handle)
// 	cCallbackInfo.userdata2 = nil
// 	return Future{Id: uint64(C.wgpuInstanceRequestAdapter(cInstance, pOptions, cCallbackInfo).id)}
// }

// //export goRequestAdapterCallbackHandler
// func goRequestAdapterCallbackHandler(status C.WGPURequestAdapterStatus, adapter C.WGPUAdapter, message C.WGPUStringView, userData1, userData2 unsafe.Pointer) {
// 	handleID := uintptr(userData1)
// 	if handleID == 0 {
// 		return
// 	}

// 	handle := cgo.Handle(handleID)
// 	defer handle.Delete()

// 	// THE FIX: Use the named type defined in your package
// 	// Instead of: .(func(RequestAdapterStatus, *Adapter, string))
// 	fn := handle.Value().(RequestAdapterCallback)

// 	// Safety check for the string (WGPUStringView is not null-terminated)
// 	// var message string
// 	// if cMessage.data != nil && cMessage.length > 0 {
// 	// 	message = C.GoStringN(cMessage.data, C.int(cMessage.length))
// 	// }

// 	// Call the function
// 	fn(
// 		RequestAdapterStatus(status),
// 		&Adapter{ref: uintptr(unsafe.Pointer(adapter))},
// 		"message",
// 	)
// }
