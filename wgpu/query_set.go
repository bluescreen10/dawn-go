//go:build !js

package wgpu

/*
#include "webgpu.h"
*/
import "C"

type QuerySet struct {
	ref C.WGPUQuerySet
}

func (q *QuerySet) SetLabel(label string) {
	C.wgpuQuerySetSetLabel(q.ref, toCStr(label))
}

func (q *QuerySet) GetType() QueryType {
	return QueryType(C.wgpuQuerySetGetType(q.ref))
}

func (q *QuerySet) GetCount() uint32 {
	return uint32(C.wgpuQuerySetGetCount(q.ref))
}

func (q *QuerySet) Destroy() {
	C.wgpuQuerySetDestroy(q.ref)
}
