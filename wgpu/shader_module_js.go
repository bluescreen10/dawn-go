//go:build js

package wgpu

import (
	"fmt"
	"syscall/js"
)

// ShaderModule represents a compiled shader module that can be used in GPU pipelines.
// Shader modules are created from WGSL or SPIR-V source code.
type ShaderModule struct {
	ref js.Value
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
	val, err := awaitPromise(s.ref.Call("getCompilationInfo"))
	if err != nil {
		return nil, fmt.Errorf("getCompilationInfo: %w", err)
	}

	jsMessages := val.Get("messages")
	length := jsMessages.Length()
	messages := make([]CompilationMessage, length)
	for i := 0; i < length; i++ {
		m := jsMessages.Index(i)
		var typ CompilationMessageType
		switch m.Get("type").String() {
		case "error":
			typ = CompilationMessageTypeError
		case "warning":
			typ = CompilationMessageTypeWarning
		default:
			typ = CompilationMessageTypeInfo
		}
		messages[i] = CompilationMessage{
			Message: m.Get("message").String(),
			Type:    typ,
			LineNum: uint64(m.Get("lineNum").Int()),
			LinePos: uint64(m.Get("linePos").Int()),
			Offset:  uint64(m.Get("offset").Int()),
			Length:  uint64(m.Get("length").Int()),
		}
	}
	return messages, nil
}

// SetLabel sets the debug label for the shader module.
// This label appears in debuggers and validation layers.
func (s *ShaderModule) SetLabel(label string) {
	s.ref.Set("label", label)
}

// Release releases the shader module and all associated resources.
// In JS, this is a no-op as GC handles cleanup.
func (s *ShaderModule) Release() {}
