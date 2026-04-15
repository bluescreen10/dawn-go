//go:build linux && !android && wayland

package wgpuglfw

import (
	"unsafe"

	"github.com/bluescreen10/dawn-go/wgpu"
	"github.com/go-gl/glfw/v3.4/glfw"
)

func GetSurfaceDescriptor(w *glfw.Window) wgpu.SurfaceDescriptor {
	return wgpu.SurfaceDescriptor{
		WaylandSurface: &wgpu.SurfaceSourceWaylandSurface{
			Display: unsafe.Pointer(glfw.GetWaylandDisplay()),
			Surface: unsafe.Pointer(w.GetWaylandWindow()),
		},
	}
}
