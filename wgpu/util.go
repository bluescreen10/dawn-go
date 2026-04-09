package wgpu

/*
#include "./lib/webgpu.h"

extern void cgo_callback_PopErrorScopeCallback(WGPUPopErrorScopeStatus status, WGPUErrorType typ, WGPUStringView message, void *userData1, void *userData2);
*/
import "C"
import (
	"fmt"
	"runtime/cgo"
	"unsafe"
)

func boolToWGPUBool(value bool) C.WGPUBool {
	if value {
		return 1
	} else {
		return 0
	}
}

func makeErrorCallback(err *error) C.WGPUPopErrorScopeCallbackInfo {
	callback := popErrorScopeCallback(func(status popErrorScopeStatus, typ ErrorType, message string) {
		if status != popErrorScopeStatusSuccess {
			*err = fmt.Errorf("error: %s", message)
		}
	})

	handle := cgo.NewHandle(callback)

	return C.WGPUPopErrorScopeCallbackInfo{
		mode:      C.WGPUCallbackMode(callbackModeAllowSpontaneous),
		callback:  C.WGPUPopErrorScopeCallback(C.cgo_callback_PopErrorScopeCallback),
		userdata1: unsafe.Pointer(handle),
		userdata2: nil,
	}
}

func ToBytes[T any, E []T](data E) []byte {
	l := uintptr(len(data))
	if l == 0 {
		return nil
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), l*unsafe.Sizeof(data[0]))
}
