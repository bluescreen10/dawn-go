//go:build !js

package wgpu

/*
#include "webgpu.h"

extern void cgo_callback_CompilationInfoCallback(WGPUCompilationInfoRequestStatus status, WGPUCompilationInfo compilationInfo, void *userData1, void *userData2);
*/
import "C"
import (
	"fmt"
	"runtime/cgo"
	"unsafe"
)

// ShaderModule represents a compiled shader module that can be used in GPU pipelines.
// Shader modules are created from WGSL or SPIR-V source code.
type ShaderModule struct {
	ref C.WGPUShaderModule
}

//export goCompilationInfoCallbackHandler
func goCompilationInfoCallbackHandler(status C.WGPUCompilationInfoRequestStatus, compilationInfo C.WGPUCompilationInfo, userData1 unsafe.Pointer, userData2 unsafe.Pointer) {
	handleID := uintptr(userData1)
	if handleID == 0 {
		return
	}

	handle := cgo.Handle(handleID)
	fn := handle.Value().(compilationInfoCallback)

	messages := make([]CompilationMessage, compilationInfo.messageCount)
	slice := unsafe.Slice((*C.WGPUCompilationMessage)(unsafe.Pointer(compilationInfo.messages)), compilationInfo.messageCount)

	for i, m := range slice {
		messages[i] = CompilationMessage{
			Message: C.GoStringN(m.message.data, C.int(m.message.length)),
			Type:    CompilationMessageType(m._type),
			LineNum: uint64(m.lineNum),
			LinePos: uint64(m.linePos),
			Offset:  uint64(m.offset),
			Length:  uint64(m.length),
		}
	}

	fn(
		compilationInfoRequestStatus(status),
		messages,
	)
}

// GetCompilationInfo returns compilation messages from the shader module compilation.
// Panics if the information cannot be retrieved.
func (s *ShaderModule) GetCompilationInfo() []CompilationMessage {
	info, err := s.TryGetCompilationInfo()
	if err != nil {
		panic(err)
	}
	return info
}

// TryGetCompilationInfo returns compilation messages from the shader module compilation, or an error if they cannot be retrieved.
func (s *ShaderModule) TryGetCompilationInfo() ([]CompilationMessage, error) {
	var status compilationInfoRequestStatus
	var info []CompilationMessage

	callback := compilationInfoCallback(func(s compilationInfoRequestStatus, i []CompilationMessage) {
		status = s
		info = i
	})

	handle := cgo.NewHandle(callback)
	defer handle.Delete()

	cCallbackInfo := C.WGPUCompilationInfoCallbackInfo{
		mode:      C.WGPUCallbackMode(callbackModeAllowSpontaneous),
		callback:  C.WGPUCompilationInfoCallback(C.cgo_callback_CompilationInfoCallback),
		userdata1: unsafe.Pointer(handle),
	}

	C.wgpuShaderModuleGetCompilationInfo(s.ref, cCallbackInfo)

	if status != compilationInfoRequestStatusSuccess {
		return nil, fmt.Errorf("error getting compilation info")
	}

	return info, nil
}

// SetLabel sets the debug label for the shader module.
// This label appears in debuggers and validation layers.
func (s *ShaderModule) SetLabel(label string) {
	C.wgpuShaderModuleSetLabel(s.ref, toCStr(label))
}

// Release releases the shader module and all associated resources.
// After calling this method, the module should no longer be used.
func (s *ShaderModule) Release() {
	C.wgpuShaderModuleRelease(s.ref)
}
