package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strconv"
	"unsafe"

	"github.com/bluescreen10/dawn-go/wgpu"
	"github.com/bluescreen10/dawn-go/wgpuglfw"
	"github.com/go-gl/glfw/v3.4/glfw"

	_ "embed"
)

const (
	// number of boid particles to simulate
	NumParticles = 1500
	// number of single-particle calculations (invocations) in each gpu work group
	ParticlesPerGroup = 64
)

//go:embed compute.wgsl
var compute string

//go:embed draw.wgsl
var draw string

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
			Offset:         uint64(wgpu.VertexFormatFloat32x2.Size()),
			ShaderLocation: 1,
		},
	},
}

type State struct {
	surface            *wgpu.Surface
	adapter            *wgpu.Adapter
	device             *wgpu.Device
	queue              *wgpu.Queue
	config             wgpu.SurfaceConfiguration
	renderPipeline     *wgpu.RenderPipeline
	computePipeline    *wgpu.ComputePipeline
	vertexBuffer       *wgpu.Buffer
	simParamBuffer     *wgpu.Buffer
	particleBindGroups []*wgpu.BindGroup
	particleBuffers    []*wgpu.Buffer
	frameNum           uint64
	workGroupCount     uint32
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

	computeShader := s.device.CreateShaderModule(wgpu.ShaderModuleDescriptor{
		Label: "compute.wgsl",
		WGSLSource: &wgpu.ShaderSourceWGSL{
			Code: compute,
		},
	})

	defer computeShader.Release()

	drawShader := s.device.CreateShaderModule(wgpu.ShaderModuleDescriptor{
		Label: "draw.wgsl",
		WGSLSource: &wgpu.ShaderSourceWGSL{
			Code: draw,
		},
	})

	defer drawShader.Release()

	simParamData := [...]float32{
		0.04,  // deltaT
		0.1,   // rule1Distance
		0.025, // rule2Distance
		0.025, // rule3Distance
		0.02,  // rule1Scale
		0.05,  // rule2Scale
		0.005, // rule3Scale
	}

	simParamBuffer := s.device.CreateBufferInit(wgpu.BufferInitDescriptor{
		Label:    "Simulation Param Buffer",
		Contents: wgpu.ToBytes(simParamData[:]),
		Usage:    wgpu.BufferUsageUniform | wgpu.BufferUsageCopyDst,
	})

	s.renderPipeline = s.device.CreateRenderPipeline(wgpu.RenderPipelineDescriptor{
		Vertex: wgpu.VertexState{
			Module:     drawShader,
			EntryPoint: "main_vs",
			Buffers: []wgpu.VertexBufferLayout{
				{
					ArrayStride: 4 * 4,
					StepMode:    wgpu.VertexStepModeInstance,
					Attributes: []wgpu.VertexAttribute{
						{
							Format:         wgpu.VertexFormatFloat32x2,
							Offset:         0,
							ShaderLocation: 0,
						},
						{
							Format:         wgpu.VertexFormatFloat32x2,
							Offset:         uint64(wgpu.VertexFormatFloat32x2.Size()),
							ShaderLocation: 1,
						},
					},
				},
				{
					ArrayStride: 2 * 4,
					StepMode:    wgpu.VertexStepModeVertex,
					Attributes: []wgpu.VertexAttribute{
						{
							Format:         wgpu.VertexFormatFloat32x2,
							Offset:         0,
							ShaderLocation: 2,
						},
					},
				},
			},
		},
		Fragment: &wgpu.FragmentState{
			Module:     drawShader,
			EntryPoint: "main_fs",
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
		},
		Multisample: wgpu.MultisampleState{
			Count:                  1,
			Mask:                   0xFFFFFFFF,
			AlphaToCoverageEnabled: false,
		},
	})

	s.computePipeline = s.device.CreateComputePipeline(wgpu.ComputePipelineDescriptor{
		Label: "Compute pipeline",
		Compute: wgpu.ComputeState{
			Module:     computeShader,
			EntryPoint: "main",
		},
	})

	vertexBufferData := [...]float32{-0.01, -0.02, 0.01, -0.02, 0.00, 0.02}
	s.vertexBuffer = s.device.CreateBufferInit(wgpu.BufferInitDescriptor{
		Label:    "Vertex Buffer",
		Contents: wgpu.ToBytes(vertexBufferData[:]),
		Usage:    wgpu.BufferUsageVertex | wgpu.BufferUsageCopyDst,
	})

	var initialParticleData [4 * NumParticles]float32
	rng := rand.NewSource(42)

	for i := 0; i < len(initialParticleData); i += 4 {
		initialParticleData[i+0] = float32(rng.Int63())/math.MaxInt64*2 - 1
		initialParticleData[i+1] = float32(rng.Int63())/math.MaxInt64*2 - 1
		initialParticleData[i+2] = (float32(rng.Int63())/math.MaxInt64*2 - 1) * 0.1
		initialParticleData[i+3] = (float32(rng.Int63())/math.MaxInt64*2 - 1) * 0.1
	}

	for i := 0; i < 2; i++ {
		particleBuffer := s.device.CreateBufferInit(wgpu.BufferInitDescriptor{
			Label:    "Particle Buffer " + strconv.Itoa(i),
			Contents: wgpu.ToBytes(initialParticleData[:]),
			Usage: wgpu.BufferUsageVertex |
				wgpu.BufferUsageStorage |
				wgpu.BufferUsageCopyDst,
		})

		s.particleBuffers = append(s.particleBuffers, particleBuffer)
	}

	computeBindGroupLayout := s.computePipeline.GetBindGroupLayout(0)
	defer computeBindGroupLayout.Release()

	for i := 0; i < 2; i++ {
		particleBindGroup := s.device.CreateBindGroup(wgpu.BindGroupDescriptor{
			Layout: computeBindGroupLayout,
			Entries: []wgpu.BindGroupEntry{
				{
					Binding: 0,
					Buffer:  simParamBuffer,
					Size:    wgpu.WholeSize,
				},
				{
					Binding: 1,
					Buffer:  s.particleBuffers[i],
					Size:    wgpu.WholeSize,
				},
				{
					Binding: 2,
					Buffer:  s.particleBuffers[(i+1)%2],
					Size:    wgpu.WholeSize,
				},
			},
		})

		s.particleBindGroups = append(s.particleBindGroups, particleBindGroup)
	}

	s.workGroupCount = uint32(math.Ceil(float64(NumParticles) / float64(ParticlesPerGroup)))
	s.frameNum = uint64(0)

	return nil
}

func (s *State) Resize(width, height int) {
	if width > 0 && height > 0 {
		s.config.Width = uint32(width)
		s.config.Height = uint32(height)
		s.surface.Configure(s.config)
	}
}

func (s *State) Render() error {
	nextTexture := s.surface.GetCurrentTexture()

	view := nextTexture.CreateView(nil)
	defer view.Release()

	commandEncoder := s.device.CreateCommandEncoder(nil)
	defer commandEncoder.Release()

	computePass := commandEncoder.BeginComputePass(nil)
	computePass.SetPipeline(s.computePipeline)
	computePass.SetBindGroup(0, s.particleBindGroups[s.frameNum%2], nil)
	computePass.DispatchWorkgroups(s.workGroupCount, 1, 1)
	computePass.End()
	computePass.Release() // must release immediately

	renderPass := commandEncoder.BeginRenderPass(wgpu.RenderPassDescriptor{
		ColorAttachments: []wgpu.RenderPassColorAttachment{
			{
				View:    view,
				LoadOp:  wgpu.LoadOpLoad,
				StoreOp: wgpu.StoreOpStore,
			},
		},
	})
	renderPass.SetPipeline(s.renderPipeline)
	renderPass.SetVertexBuffer(0, s.particleBuffers[(s.frameNum+1)%2], 0, wgpu.WholeSize)
	renderPass.SetVertexBuffer(1, s.vertexBuffer, 0, wgpu.WholeSize)
	renderPass.Draw(3, NumParticles, 0, 0)
	renderPass.End()

	s.frameNum += 1

	cmdBuffer := commandEncoder.Finish(nil)
	defer cmdBuffer.Release()

	s.queue.Submit(cmdBuffer)
	s.surface.Present()

	return nil
}

func (s *State) Destroy() {
	if s.particleBindGroups != nil {
		for _, bg := range s.particleBindGroups {
			bg.Release()
		}
		s.particleBindGroups = nil
	}

	if s.particleBuffers != nil {
		for _, buffer := range s.particleBuffers {
			buffer.Destroy()
		}
		s.particleBuffers = nil
	}
	if s.vertexBuffer != nil {
		s.vertexBuffer.Release()
		s.vertexBuffer = nil
	}
	if s.computePipeline != nil {
		s.computePipeline.Release()
		s.computePipeline = nil
	}

	if s.simParamBuffer != nil {
		s.simParamBuffer.Destroy()
		s.simParamBuffer = nil
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
