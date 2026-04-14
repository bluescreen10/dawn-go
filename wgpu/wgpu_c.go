//go:build !js

package wgpu

/*
#cgo CFLAGS: -I ${SRCDIR}/lib

#include "webgpu.h"
*/
import "C"
import (
	"unsafe"

	_ "github.com/bluescreen10/dawn-go/wgpu/lib"
)

func toCBool(value bool) C.WGPUBool {
	if value {
		return 1
	} else {
		return 0
	}
}

func toCStr(val string) C.WGPUStringView {
	return C.WGPUStringView{
		data:   (*C.char)(unsafe.Pointer(unsafe.StringData(val))),
		length: C.size_t(len(val)),
	}
}
