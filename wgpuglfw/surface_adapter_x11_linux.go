//go:build linux && !android && !wayland

package wgpuglfw

import (
	"unsafe"

	"github.com/bluescreen10/dawn-go/wgpu"
	"github.com/go-gl/glfw/v3.4/glfw"
)

func GetSurfaceDescriptor(w *glfw.Window) wgpu.SurfaceDescriptor {
	return wgpu.SurfaceDescriptor{
		XlibWindow: &wgpu.SurfaceSourceXlibWindow{
			Display: unsafe.Pointer(glfw.GetX11Display()),
			Window:  uint64(w.GetX11Window()),
		},
	}
}
