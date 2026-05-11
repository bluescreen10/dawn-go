//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"
import "runtime"

// QuerySet represents a query set that can be used to collect timestamp and occlusion query results.
// Query sets are created from a device and have a specific type and count.
type QuerySet struct {
	ref C.WGPUQuerySet
}

// SetLabel sets the debug label for the query set.
// This label appears in debuggers and validation layers.
func (q *QuerySet) SetLabel(label string) {
	var pinner runtime.Pinner
	defer pinner.Unpin()

	cLabel := toCStr(label)
	pinner.Pin(cLabel.data)

	C.wgpuQuerySetSetLabel(q.ref, cLabel)
}

// GetType returns the type of the query set (timestamp or occlusion).
func (q *QuerySet) GetType() QueryType {
	return QueryType(C.wgpuQuerySetGetType(q.ref))
}

// GetCount returns the number of queries in the query set.
func (q *QuerySet) GetCount() uint32 {
	return uint32(C.wgpuQuerySetGetCount(q.ref))
}

// Destroy destroys the query set and frees all associated GPU resources.
func (q *QuerySet) Destroy() {
	C.wgpuQuerySetDestroy(q.ref)
}
