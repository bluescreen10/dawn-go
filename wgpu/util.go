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

func FromBytes[T any](data []byte) []T {
	if len(data) == 0 {
		return nil
	}
	var v T
	return unsafe.Slice((*T)(unsafe.Pointer(&data[0])), uintptr(len(data))/unsafe.Sizeof(v))
}
