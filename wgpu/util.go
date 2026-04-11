package wgpu

import (
	"unsafe"
)

func ToBytes[T any, E []T](data E) []byte {
	l := uintptr(len(data))
	if l == 0 {
		return nil
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), l*unsafe.Sizeof(data[0]))
}
