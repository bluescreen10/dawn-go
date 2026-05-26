//go:build !js

package wgpu

// Future represents a handle to an asynchronous operation that can be waited on.
type Future struct {
	id uint64
}
