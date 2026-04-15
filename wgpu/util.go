package wgpu

import (
	"unsafe"
)

// ToBytes converts a slice of any type to a slice of bytes.
// This is useful for passing data to C functions that expect byte arrays.
// Returns nil if the input slice is empty.
func ToBytes[T any, E ~[]T](data E) []byte {
	l := uintptr(len(data))
	if l == 0 {
		return nil
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), l*unsafe.Sizeof(data[0]))
}

// FromBytes converts a slice of bytes to a slice of a specific type.
// This is useful for reading data from C functions that return byte arrays.
// Returns nil if the input slice is empty.
func FromBytes[T any](data []byte) []T {
	if len(data) == 0 {
		return nil
	}
	var v T
	return unsafe.Slice((*T)(unsafe.Pointer(&data[0])), uintptr(len(data))/unsafe.Sizeof(v))
}
