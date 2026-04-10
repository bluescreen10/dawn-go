//go:build darwin

package wgpuglfw

import (
	"unsafe"

	"github.com/bluescreen10/dawn-go/wgpu"
	"github.com/go-gl/glfw/v3.4/glfw"
)

/*

#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework QuartzCore

#import <Cocoa/Cocoa.h>
#import <QuartzCore/CAMetalLayer.h>

CFTypeRef metalLayerFromNSWindow(CFTypeRef nsWindowRef) {
	NSWindow *ns_window = (__bridge NSWindow *)nsWindowRef;
    [ns_window.contentView setWantsLayer:YES];

    CAMetalLayer *layer = [CAMetalLayer layer];

    layer.contentsScale = ns_window.backingScaleFactor;

    [ns_window.contentView setLayer:layer];
    return (__bridge void *)layer;
}

*/
import "C"

func GetSurfaceDescriptor(w *glfw.Window) wgpu.SurfaceDescriptor {

	cocoaWindow := w.GetCocoaWindow()
	layer := C.metalLayerFromNSWindow((C.CFTypeRef)(unsafe.Pointer(cocoaWindow)))

	return wgpu.SurfaceDescriptor{
		MetalLayer: &wgpu.SurfaceSourceMetalLayer{
			Layer: unsafe.Pointer(layer),
		},
	}
}
