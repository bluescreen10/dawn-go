package wgpu

/*
#include "./lib/webgpu.h"
*/
import "C"

func boolToWGPUBool(value bool) C.WGPUBool {
	if value {
		return 1
	} else {
		return 0
	}
}
