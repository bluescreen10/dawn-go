//go:build !js

package wgpu

/*
#cgo CFLAGS: -I./lib
#cgo LDFLAGS: -lwebgpu_dawn

// Android
#cgo android,amd64 LDFLAGS: -L${SRCDIR}/lib/android/amd64
#cgo android,386 LDFLAGS: -L${SRCDIR}/lib/android/386
#cgo android,arm64 LDFLAGS: -L${SRCDIR}/lib/android/arm64
#cgo android,arm LDFLAGS: -L${SRCDIR}/lib/android/arm
#cgo android LDFLAGS: -landroid -lm -llog

// Linux
#cgo linux,!android,amd64 LDFLAGS: -L${SRCDIR}/lib/linux/amd64
#cgo linux,!android LDFLAGS: -lm -ldl

// Darwin
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/lib/darwin/amd64
#cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/lib/darwin/arm64
#cgo darwin LDFLAGS: -framework Metal -framework IOKit -framework QuartzCore -framework Foundation -framework IOSurface -lc++

// Windows
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/lib/windows/amd64
#cgo windows LDFLAGS: -ld3dcompiler_47 -lws2_32 -luserenv -lbcrypt -lntdll

#include "webgpu.h"
*/
import "C"
import "unsafe"

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
