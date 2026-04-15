package main

import (
	"fmt"
	"math"
	"runtime"
	"unsafe"

	"github.com/bluescreen10/dawn-go/examples/glm"
	"github.com/bluescreen10/dawn-go/wgpu"
	"github.com/bluescreen10/dawn-go/wgpuglfw"
	"github.com/go-gl/glfw/v3.4/glfw"

	_ "embed"
)

func init() {
	runtime.LockOSThread()
}

type Vertex struct {
	pos      [4]float32
	texCoord [2]float32
}

var VertexBufferLayout = wgpu.VertexBufferLayout{
	ArrayStride: uint64(unsafe.Sizeof(Vertex{})),
	StepMode:    wgpu.VertexStepModeVertex,
	Attributes: []wgpu.VertexAttribute{
		{
			Format:         wgpu.VertexFormatFloat32x4,
			Offset:         0,
			ShaderLocation: 0,
		},
		{
			Format:         wgpu.VertexFormatFloat32x2,
			Offset:         4 * 4,
			ShaderLocation: 1,
		},
	},
}

func vertex(pos1, pos2, pos3, tc1, tc2 float32) Vertex {
	return Vertex{
		pos:      [4]float32{pos1, pos2, pos3, 1},
		texCoord: [2]float32{tc1, tc2},
	}
}

var vertexData = [...]Vertex{
	// top (0, 0, 1)
	vertex(-1, -1, 1, 0, 0),
	vertex(1, -1, 1, 1, 0),
	vertex(1, 1, 1, 1, 1),
	vertex(-1, 1, 1, 0, 1),
	// bottom (0, 0, -1)
	vertex(-1, 1, -1, 1, 0),
	vertex(1, 1, -1, 0, 0),
	vertex(1, -1, -1, 0, 1),
	vertex(-1, -1, -1, 1, 1),
	// right (1, 0, 0)
	vertex(1, -1, -1, 0, 0),
	vertex(1, 1, -1, 1, 0),
	vertex(1, 1, 1, 1, 1),
	vertex(1, -1, 1, 0, 1),
	// left (-1, 0, 0)
	vertex(-1, -1, 1, 1, 0),
	vertex(-1, 1, 1, 0, 0),
	vertex(-1, 1, -1, 0, 1),
	vertex(-1, -1, -1, 1, 1),
	// front (0, 1, 0)
	vertex(1, 1, -1, 1, 0),
	vertex(-1, 1, -1, 0, 0),
	vertex(-1, 1, 1, 0, 1),
	vertex(1, 1, 1, 1, 1),
	// back (0, -1, 0)
	vertex(1, -1, 1, 0, 0),
	vertex(-1, -1, 1, 1, 0),
	vertex(-1, -1, -1, 1, 1),
	vertex(1, -1, -1, 0, 1),
}

var indexData = [...]uint16{
	0, 1, 2, 2, 3, 0, // top
	4, 5, 6, 6, 7, 4, // bottom
	8, 9, 10, 10, 11, 8, // right
	12, 13, 14, 14, 15, 12, // left
	16, 17, 18, 18, 19, 16, // front
	20, 21, 22, 22, 23, 20, // back
}

const texelsSize = 256

func createTexels() (texels [texelsSize * texelsSize]uint8) {
	for id := 0; id < (texelsSize * texelsSize); id++ {
		cx := 3.0*float32(id%texelsSize)/float32(texelsSize-1) - 2.0
		cy := 2.0*float32(id/texelsSize)/float32(texelsSize-1) - 1.0
		x, y, count := float32(cx), float32(cy), uint8(0)
		for count < 0xFF && x*x+y*y < 4.0 {
			oldX := x
			x = x*x - y*y + cx
			y = 2.0*oldX*y + cy
			count += 1
		}
		texels[id] = count
	}

	return texels
}

func generateMatrix(aspectRatio float32) glm.Mat4[float32] {
	projection := glm.PerspectiveRH(math.Pi/4, aspectRatio, 1, 10)
	view := glm.LookAtRH(
		glm.Vec3[float32]{1.5, -5, 3},
		glm.Vec3[float32]{0, 0, 0},
		glm.Vec3[float32]{0, 0, 1},
	)

	return projection.Mul4(view)
}

//go:embed shader.wgsl
var shader string

type State struct {
	surface    *wgpu.Surface
	adapter    *wgpu.Adapter
	device     *wgpu.Device
	queue      *wgpu.Queue
	config     wgpu.SurfaceConfiguration
	vertexBuf  *wgpu.Buffer
	indexBuf   *wgpu.Buffer
	uniformBuf *wgpu.Buffer
	pipeline   *wgpu.RenderPipeline
	bindGroup  *wgpu.BindGroup
}

func (s *State) Init(window *glfw.Window) (err error) {
	defer func() {
		if err != nil {
			s.Destroy()
			s = nil
		}
	}()

	instance := wgpu.CreateInstance(nil)
	defer instance.Release()

	s.surface = instance.CreateSurface(wgpuglfw.GetSurfaceDescriptor(window))

	s.adapter, err = instance.RequestAdapter(&wgpu.RequestAdapterOptions{
		ForceFallbackAdapter: false,
		CompatibleSurface:    s.surface,
	})

	if err != nil {
		panic(err)
	}

	s.device = s.adapter.RequestDevice(&wgpu.DeviceDescriptor{
		UncapturedErrorCallback: wgpu.UncapturedErrorCallback(func(_ *wgpu.Device, typ wgpu.ErrorType, message string) {
			panic(fmt.Sprintf("Type: %s Msg: %s\n", typ, message))
		}),
	})

	s.queue = s.device.GetQueue()

	caps, err := s.surface.GetCapabilities(s.adapter)
	if err != nil {
		panic(err)
	}

	width, height := window.GetSize()
	s.config = wgpu.SurfaceConfiguration{
		Device:      s.device,
		Usage:       wgpu.TextureUsageRenderAttachment,
		Format:      caps.Formats[0],
		Width:       uint32(width),
		Height:      uint32(height),
		PresentMode: wgpu.PresentModeFifo,
		AlphaMode:   caps.AlphaModes[0],
	}

	s.surface.Configure(s.config)

	s.vertexBuf = s.device.CreateBufferInit(wgpu.BufferInitDescriptor{
		Label:    "Vertex Buffer",
		Contents: wgpu.ToBytes(vertexData[:]),
		Usage:    wgpu.BufferUsageVertex,
	})

	s.indexBuf = s.device.CreateBufferInit(wgpu.BufferInitDescriptor{
		Label:    "Index Buffer",
		Contents: wgpu.ToBytes(indexData[:]),
		Usage:    wgpu.BufferUsageIndex,
	})

	textureExtent := wgpu.Extent3D{
		Width:              texelsSize,
		Height:             texelsSize,
		DepthOrArrayLayers: 1,
	}
	texture := s.device.CreateTexture(&wgpu.TextureDescriptor{
		Size:          textureExtent,
		MipLevelCount: 1,
		SampleCount:   1,
		Dimension:     wgpu.TextureDimension2D,
		Format:        wgpu.TextureFormatR8Uint,
		Usage:         wgpu.TextureUsageTextureBinding | wgpu.TextureUsageCopyDst,
	})

	textureView := texture.CreateView(nil)

	texels := createTexels()
	s.queue.WriteTexture(
		texture.AsImageCopy(),
		wgpu.ToBytes(texels[:]),
		wgpu.TexelCopyBufferLayout{
			Offset:       0,
			BytesPerRow:  texelsSize,
			RowsPerImage: wgpu.CopyStrideUndefined,
		},
		textureExtent,
	)

	mxTotal := generateMatrix(float32(s.config.Width) / float32(s.config.Height))
	s.uniformBuf = s.device.CreateBufferInit(wgpu.BufferInitDescriptor{
		Label:    "Uniform Buffer",
		Contents: wgpu.ToBytes(mxTotal[:]),
		Usage:    wgpu.BufferUsageUniform | wgpu.BufferUsageCopyDst,
	})
	if err != nil {
		return err
	}

	shader := s.device.CreateShaderModule(wgpu.ShaderModuleDescriptor{
		Label:      "shader.wgsl",
		WGSLSource: &wgpu.ShaderSourceWGSL{Code: shader},
	})

	defer shader.Release()

	s.pipeline = s.device.CreateRenderPipeline(wgpu.RenderPipelineDescriptor{
		Vertex: wgpu.VertexState{
			Module:     shader,
			EntryPoint: "vs_main",
			Buffers:    []wgpu.VertexBufferLayout{VertexBufferLayout},
		},
		Fragment: &wgpu.FragmentState{
			Module:     shader,
			EntryPoint: "fs_main",
			Targets: []wgpu.ColorTargetState{
				{
					Format:    s.config.Format,
					Blend:     nil,
					WriteMask: wgpu.ColorWriteMaskAll,
				},
			},
		},
		Primitive: wgpu.PrimitiveState{
			Topology:  wgpu.PrimitiveTopologyTriangleList,
			FrontFace: wgpu.FrontFaceCCW,
			CullMode:  wgpu.CullModeBack,
		},
		DepthStencil: nil,
		Multisample: wgpu.MultisampleState{
			Count:                  1,
			Mask:                   0xFFFFFFFF,
			AlphaToCoverageEnabled: false,
		},
	})

	bindGroupLayout := s.pipeline.GetBindGroupLayout(0)
	defer bindGroupLayout.Release()

	s.bindGroup = s.device.CreateBindGroup(wgpu.BindGroupDescriptor{
		Layout: bindGroupLayout,
		Entries: []wgpu.BindGroupEntry{
			{
				Binding: 0,
				Buffer:  s.uniformBuf,
				Size:    wgpu.WholeSize,
			},
			{
				Binding:     1,
				TextureView: textureView,
				Size:        wgpu.WholeSize,
			},
		},
	})

	return nil
}

func (s *State) Resize(width, height int) {
	if width > 0 && height > 0 {
		s.config.Width = uint32(width)
		s.config.Height = uint32(height)

		mxTotal := generateMatrix(float32(width) / float32(height))
		s.queue.WriteBuffer(s.uniformBuf, 0, wgpu.ToBytes(mxTotal[:]))

		s.surface.Configure(s.config)
	}
}

func (s *State) Render() error {
	nextTexture := s.surface.GetCurrentTexture()
	view := nextTexture.CreateView(nil)
	defer view.Release()

	encoder := s.device.CreateCommandEncoder(nil)
	defer encoder.Release()

	renderPass := encoder.BeginRenderPass(wgpu.RenderPassDescriptor{
		ColorAttachments: []wgpu.RenderPassColorAttachment{
			{
				View:       view,
				LoadOp:     wgpu.LoadOpClear,
				StoreOp:    wgpu.StoreOpStore,
				ClearValue: wgpu.Color{R: 0.1, G: 0.2, B: 0.3, A: 1.0},
			},
		},
	})

	renderPass.SetPipeline(s.pipeline)
	renderPass.SetBindGroup(0, s.bindGroup, nil)
	renderPass.SetIndexBuffer(s.indexBuf, wgpu.IndexFormatUint16, 0, wgpu.WholeSize)
	renderPass.SetVertexBuffer(0, s.vertexBuf, 0, wgpu.WholeSize)
	renderPass.DrawIndexed(uint32(len(indexData)), 1, 0, 0, 0)
	renderPass.End()

	cmdBuffer := encoder.Finish(nil)
	defer cmdBuffer.Release()

	s.queue.Submit(cmdBuffer)
	s.surface.Present()

	return nil
}

func (s *State) Destroy() {
	if s.bindGroup != nil {
		s.bindGroup.Release()
		s.bindGroup = nil
	}
	if s.pipeline != nil {
		s.pipeline.Release()
		s.pipeline = nil
	}
	if s.uniformBuf != nil {
		s.uniformBuf.Destroy()
		s.uniformBuf = nil
	}
	if s.indexBuf != nil {
		s.indexBuf.Destroy()
		s.indexBuf = nil
	}
	if s.vertexBuf != nil {
		s.vertexBuf.Destroy()
		s.vertexBuf = nil
	}

	s.config = wgpu.SurfaceConfiguration{}

	if s.queue != nil {
		s.queue.Release()
		s.queue = nil
	}
	if s.device != nil {
		s.device.Destroy()
		s.device = nil
	}

	if s.adapter != nil {
		s.adapter.Release()
		s.adapter = nil
	}

	if s.surface != nil {
		s.surface.Release()
		s.surface = nil
	}
}

func main() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	window, err := glfw.CreateWindow(640, 480, "dawn-go", nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	s := &State{}
	err = s.Init(window)
	if err != nil {
		panic(err)
	}
	defer s.Destroy()

	window.SetSizeCallback(func(w *glfw.Window, width, height int) {
		s.Resize(width, height)
	})

	for !window.ShouldClose() {
		glfw.PollEvents()
		s.Render()
	}
}
