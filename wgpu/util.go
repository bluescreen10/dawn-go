package wgpu

/*
#include "webgpu.h"
*/
import "C"
import (
	"unsafe"
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

func ToBytes[T any, E []T](data E) []byte {
	l := uintptr(len(data))
	if l == 0 {
		return nil
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), l*unsafe.Sizeof(data[0]))
}
