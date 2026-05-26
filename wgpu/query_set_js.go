//go:build js

package wgpu

import "syscall/js"

// QuerySet represents a query set that can be used to collect timestamp and occlusion query results.
// Query sets are created from a device and have a specific type and count.
type QuerySet struct {
	ref js.Value
}

// SetLabel sets the debug label for the query set.
// This label appears in debuggers and validation layers.
func (q *QuerySet) SetLabel(label string) {
	q.ref.Set("label", label)
}

// GetType returns the type of the query set (timestamp or occlusion).
func (q *QuerySet) GetType() QueryType {
	switch q.ref.Get("type").String() {
	case "occlusion":
		return QueryTypeOcclusion
	default:
		return QueryTypeTimestamp
	}
}

// GetCount returns the number of queries in the query set.
func (q *QuerySet) GetCount() uint32 {
	return uint32(q.ref.Get("count").Int())
}

// Destroy destroys the query set and frees all associated GPU resources.
func (q *QuerySet) Destroy() {
	q.ref.Call("destroy")
}
