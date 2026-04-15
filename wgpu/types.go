package wgpu

import "unsafe"

// AdapterInfo contains information about a GPU adapter, including vendor, device, and backend details.
type AdapterInfo struct {
	Vendor          string
	Architecture    string
	Device          string
	Description     string
	BackendType     BackendType
	AdapterType     AdapterType
	VendorID        uint32
	DeviceID        uint32
	SubgroupMinSize uint32
	SubgroupMaxSize uint32
}

// BindGroupDescriptor describes a bind group, which is a collection of resources bound together for rendering.
type BindGroupDescriptor struct {
	Label   string
	Layout  *BindGroupLayout
	Entries []BindGroupEntry
}

// BindGroupEntry represents a single entry in a bind group, binding a resource to a specific binding point.
type BindGroupEntry struct {
	Binding     uint32
	Buffer      *Buffer
	Offset      uint64
	Size        uint64
	Sampler     *Sampler
	TextureView *TextureView
}

// BindGroupLayoutDescriptor describes a bind group layout, which defines the interface for a bind group.
type BindGroupLayoutDescriptor struct {
	Label   string
	Entries []BindGroupLayoutEntry
}

// BindGroupLayoutEntry describes a single entry in a bind group layout, defining the type and access pattern for a binding.
type BindGroupLayoutEntry struct {
	Binding          uint32
	Visibility       ShaderStage
	BindingArraySize uint32
	Buffer           BufferBindingLayout
	Sampler          SamplerBindingLayout
	Texture          TextureBindingLayout
	StorageTexture   StorageTextureBindingLayout
}

// BlendComponent describes how to blend a single color component (red, green, blue, or alpha) during rendering.
type BlendComponent struct {
	Operation BlendOperation
	SrcFactor BlendFactor
	DstFactor BlendFactor
}

// BlendState describes how to blend color and alpha components during rendering.
type BlendState struct {
	Color BlendComponent
	Alpha BlendComponent
}

// BufferBindingLayout describes the layout for a buffer binding, including type and whether it has a dynamic offset.
type BufferBindingLayout struct {
	Type             BufferBindingType
	HasDynamicOffset bool
	MinBindingSize   uint64
}

// BufferDescriptor describes a buffer, including its size, usage flags, and whether it should be mapped at creation.
type BufferDescriptor struct {
	Label            string
	Usage            BufferUsage
	Size             uint64
	MappedAtCreation bool
}

// BufferInitDescriptor describes a buffer to be created and initialized with data in a single operation.
type BufferInitDescriptor struct {
	Label    string
	Usage    BufferUsage
	Contents []byte
}

type bufferMapCallbackInfo struct {
	Mode     callbackMode
	Callback BufferMapCallback
}

// Color represents an RGBA color with double-precision floating-point components.
type Color struct {
	R float64
	G float64
	B float64
	A float64
}

// ColorTargetState describes the state of a color target in a render pipeline, including format, blend, and write mask.
type ColorTargetState struct {
	Format    TextureFormat
	Blend     *BlendState
	WriteMask ColorWriteMask
}

// CommandBufferDescriptor describes a command buffer, primarily for setting its label.
type CommandBufferDescriptor struct {
	Label string
}

// CommandEncoderDescriptor describes a command encoder, primarily for setting its label.
type CommandEncoderDescriptor struct {
	Label string
}

// CompatibilityModeLimits contains limits specific to WebGPU compatibility mode.
type CompatibilityModeLimits struct {
	MaxStorageBuffersInVertexStage    uint32
	MaxStorageTexturesInVertexStage   uint32
	MaxStorageBuffersInFragmentStage  uint32
	MaxStorageTexturesInFragmentStage uint32
}

// CompilationMessage contains a message from shader compilation, including type, text, and location information.
type CompilationMessage struct {
	Message string
	Type    CompilationMessageType
	LineNum uint64
	LinePos uint64
	Offset  uint64
	Length  uint64
}

// ComputePassDescriptor describes a compute pass, including label and timestamp writes.
type ComputePassDescriptor struct {
	Label           string
	TimestampWrites *PassTimestampWrites
}

// ComputePipelineDescriptor describes a compute pipeline, including label, layout, and compute stage.
type ComputePipelineDescriptor struct {
	Label   string
	Layout  *PipelineLayout
	Compute ComputeState
}

// ComputeState describes the compute stage of a compute pipeline, including the shader module and entry point.
type ComputeState struct {
	Module     *ShaderModule
	EntryPoint string
	Constants  []ConstantEntry
}

// ConstantEntry defines a constant value to be used in a shader module, keyed by name.
type ConstantEntry struct {
	Key   string
	Value float64
}

type createComputePipelineAsyncCallbackInfo struct {
	Mode     callbackMode
	Callback CreateComputePipelineAsyncCallback
}

type createRenderPipelineAsyncCallbackInfo struct {
	Mode     callbackMode
	Callback CreateRenderPipelineAsyncCallback
}

// DepthStencilState describes the depth-stencil state for a render pipeline, including format, compare function, and stencil operations.
type DepthStencilState struct {
	Format              TextureFormat
	DepthWriteEnabled   OptionalBool
	DepthCompare        CompareFunction
	StencilFront        StencilFaceState
	StencilBack         StencilFaceState
	StencilReadMask     uint32
	StencilWriteMask    uint32
	DepthBias           int32
	DepthBiasSlopeScale float32
	DepthBiasClamp      float32
}

// DeviceDescriptor describes a logical device, including required features, limits, and callbacks for errors and device loss.
type DeviceDescriptor struct {
	Label                   string
	RequiredFeatures        []FeatureName
	RequiredLimits          *Limits
	DefaultQueue            QueueDescriptor
	DeviceLostCallback      DeviceLostCallback
	UncapturedErrorCallback UncapturedErrorCallback
}

type deviceLostCallbackInfo struct {
	Mode     callbackMode
	Callback DeviceLostCallback
}

// Extent3D defines the size of a resource in three dimensions, used for textures and other resources.
type Extent3D struct {
	Width              uint32
	Height             uint32
	DepthOrArrayLayers uint32
}

// FragmentState describes the fragment stage of a render pipeline, including shader module, entry point, constants, and color targets.
type FragmentState struct {
	Module     *ShaderModule
	EntryPoint string
	Constants  []ConstantEntry
	Targets    []ColorTargetState
}

// Future represents a handle to an asynchronous operation that can be waited on.
type Future struct {
	id uint64
}

// InstanceDescriptor describes an instance, including required features and limits.
type InstanceDescriptor struct {
	RequiredFeatures []InstanceFeatureName
	RequiredLimits   *InstanceLimits
}

// InstanceLimits contains instance-level limits, such as the maximum number of futures that can be waited on concurrently.
type InstanceLimits struct {
	TimedWaitAnyMaxCount int
}

// Limits contains device-level limits for various GPU capabilities, such as maximum texture dimensions and buffer sizes.
type Limits struct {
	MaxTextureDimension1D                     uint32
	MaxTextureDimension2D                     uint32
	MaxTextureDimension3D                     uint32
	MaxTextureArrayLayers                     uint32
	MaxBindGroups                             uint32
	MaxBindGroupsPlusVertexBuffers            uint32
	MaxBindingsPerBindGroup                   uint32
	MaxDynamicUniformBuffersPerPipelineLayout uint32
	MaxDynamicStorageBuffersPerPipelineLayout uint32
	MaxSampledTexturesPerShaderStage          uint32
	MaxSamplersPerShaderStage                 uint32
	MaxStorageBuffersPerShaderStage           uint32
	MaxStorageTexturesPerShaderStage          uint32
	MaxUniformBuffersPerShaderStage           uint32
	MaxUniformBufferBindingSize               uint64
	MaxStorageBufferBindingSize               uint64
	MinUniformBufferOffsetAlignment           uint32
	MinStorageBufferOffsetAlignment           uint32
	MaxVertexBuffers                          uint32
	MaxBufferSize                             uint64
	MaxVertexAttributes                       uint32
	MaxVertexBufferArrayStride                uint32
	MaxInterStageShaderVariables              uint32
	MaxColorAttachments                       uint32
	MaxColorAttachmentBytesPerSample          uint32
	MaxComputeWorkgroupStorageSize            uint32
	MaxComputeInvocationsPerWorkgroup         uint32
	MaxComputeWorkgroupSizeX                  uint32
	MaxComputeWorkgroupSizeY                  uint32
	MaxComputeWorkgroupSizeZ                  uint32
	MaxComputeWorkgroupsPerDimension          uint32
	MaxImmediateSize                          uint32
}

// MultisampleState describes multisampling settings for a render pipeline, including sample count, mask, and alpha-to-coverage.
type MultisampleState struct {
	Count                  uint32
	Mask                   uint32
	AlphaToCoverageEnabled bool
}

// Origin3D defines a three-dimensional origin point, used for texture copy origins and viewport origins.
type Origin3D struct {
	X uint32
	Y uint32
	Z uint32
}

// PassTimestampWrites defines where to write timestamps in a render or compute pass.
type PassTimestampWrites struct {
	QuerySet                  *QuerySet
	BeginningOfPassWriteIndex uint32
	EndOfPassWriteIndex       uint32
}

// PipelineLayoutDescriptor describes a pipeline layout, which defines the bind group layouts used by a pipeline.
type PipelineLayoutDescriptor struct {
	Label            string
	BindGroupLayouts []*BindGroupLayout
	ImmediateSize    uint32
}

type popErrorScopeCallbackInfo struct {
	Mode     callbackMode
	Callback PopErrorScopeCallback
}

// PrimitiveState describes the primitive topology and rasterization settings for a render pipeline.
type PrimitiveState struct {
	Topology         PrimitiveTopology
	StripIndexFormat IndexFormat
	FrontFace        FrontFace
	CullMode         CullMode
	UnclippedDepth   bool
}

// QuerySetDescriptor describes a query set, which is used to collect timestamp and occlusion query results.
type QuerySetDescriptor struct {
	Label string
	Type  QueryType
	Count uint32
}

// QueueDescriptor describes a queue, primarily for setting its label.
type QueueDescriptor struct {
	Label string
}

type QueueWorkDoneCallbackInfo struct {
	Mode     callbackMode
	Callback QueueWorkDoneCallback
}

// RenderBundleDescriptor describes a render bundle, primarily for setting its label.
type RenderBundleDescriptor struct {
	Label string
}

// RenderBundleEncoderDescriptor describes a render bundle encoder, including color formats, depth stencil format, and sample count.
type RenderBundleEncoderDescriptor struct {
	Label              string
	ColorFormats       []TextureFormat
	DepthStencilFormat TextureFormat
	SampleCount        uint32
	DepthReadOnly      bool
	StencilReadOnly    bool
}

// RenderPassColorAttachment describes a color attachment for a render pass, including the texture view, load/store operations, and clear value.
type RenderPassColorAttachment struct {
	View          *TextureView
	DepthSlice    uint32
	ResolveTarget *TextureView
	LoadOp        LoadOp
	StoreOp       StoreOp
	ClearValue    Color
}

// RenderPassDepthStencilAttachment describes a depth-stencil attachment for a render pass, including format, load/store operations, and clear values.
type RenderPassDepthStencilAttachment struct {
	View              *TextureView
	DepthLoadOp       LoadOp
	DepthStoreOp      StoreOp
	DepthClearValue   float32
	DepthReadOnly     bool
	StencilLoadOp     LoadOp
	StencilStoreOp    StoreOp
	StencilClearValue uint32
	StencilReadOnly   bool
}

// RenderPassDescriptor describes a render pass, including label, color attachments, depth stencil attachment, and timestamp writes.
type RenderPassDescriptor struct {
	Label                  string
	ColorAttachments       []RenderPassColorAttachment
	DepthStencilAttachment *RenderPassDepthStencilAttachment
	OcclusionQuerySet      *QuerySet
	TimestampWrites        *PassTimestampWrites
}

// RenderPassMaxDrawCount defines the maximum number of draws allowed in a render pass.
type RenderPassMaxDrawCount struct {
	MaxDrawCount uint64
}

// RenderPipelineDescriptor describes a render pipeline, including label, layout, vertex, primitive, depth stencil, multisample, and fragment states.
type RenderPipelineDescriptor struct {
	Label        string
	Layout       *PipelineLayout
	Vertex       VertexState
	Primitive    PrimitiveState
	DepthStencil *DepthStencilState
	Multisample  MultisampleState
	Fragment     *FragmentState
}

// RequestAdapterWebXROptions contains WebXR-specific options for requesting an adapter.
//type RequestAdapterWebXROptions struct {
//	XrCompatible bool
//}

type requestAdapterCallbackInfo struct {
	Mode     callbackMode
	Callback requestAdapterCallback
}

// RequestAdapterOptions contains options for requesting an adapter, such as power preference and compatible surface.
type RequestAdapterOptions struct {
	FeatureLevel         FeatureLevel
	PowerPreference      PowerPreference
	ForceFallbackAdapter bool
	BackendType          BackendType
	CompatibleSurface    *Surface
}

type requestDeviceCallbackInfo struct {
	Mode     callbackMode
	Callback requestDeviceCallback
}

// SamplerBindingLayout describes the layout for a sampler binding, including the sampler type.
type SamplerBindingLayout struct {
	Type SamplerBindingType
}

// SamplerDescriptor describes a sampler, including addressing modes, filtering modes, and comparison function.
type SamplerDescriptor struct {
	Label         string
	AddressModeU  AddressMode
	AddressModeV  AddressMode
	AddressModeW  AddressMode
	MagFilter     FilterMode
	MinFilter     FilterMode
	MipmapFilter  MipmapFilterMode
	LodMinClamp   float32
	LodMaxClamp   float32
	Compare       CompareFunction
	MaxAnisotropy uint16
}

// ShaderModuleDescriptor describes a shader module, which can contain either WGSL or SPIR-V source code.
type ShaderModuleDescriptor struct {
	Label string

	SPIRVSource *ShaderSourceSPIRV
	WGSLSource  *ShaderSourceWGSL
}

// ShaderSourceSPIRV contains SPIR-V shader source code as a slice of uint32 words.
type ShaderSourceSPIRV struct {
	Code []uint32
}

// ShaderSourceWGSL contains WGSL (WebGPU Shading Language) shader source code as a string.
type ShaderSourceWGSL struct {
	Code string
}

// StencilFaceState describes the stencil state for a single face (front or back) of a depth-stencil attachment.
type StencilFaceState struct {
	Compare     CompareFunction
	FailOp      StencilOperation
	DepthFailOp StencilOperation
	PassOp      StencilOperation
}

// StorageTextureBindingLayout describes the layout for a storage texture binding, including access mode, format, and view dimension.
type StorageTextureBindingLayout struct {
	Access        StorageTextureAccess
	Format        TextureFormat
	ViewDimension TextureViewDimension
}

// SurfaceCapabilities contains the capabilities of a surface, including supported usages, formats, present modes, and alpha modes.
type SurfaceCapabilities struct {
	Usages       TextureUsage
	Formats      []TextureFormat
	PresentModes []PresentMode
	AlphaModes   []CompositeAlphaMode
}

// SurfaceColorManagement defines color management settings for a surface, including color space and tone mapping mode.
type SurfaceColorManagement struct {
	ColorSpace      PredefinedColorSpace
	ToneMappingMode ToneMappingMode
}

// SurfaceConfiguration configures a surface for rendering, including the device, format, usage, dimensions, and present mode.
type SurfaceConfiguration struct {
	Device      *Device
	Format      TextureFormat
	Usage       TextureUsage
	Width       uint32
	Height      uint32
	ViewFormats []TextureFormat
	AlphaMode   CompositeAlphaMode
	PresentMode PresentMode
}

// SurfaceDescriptor describes a surface, which can be created from various platform-specific sources like Metal layer, Windows HWND, or Wayland surface.
type SurfaceDescriptor struct {
	Label          string
	MetalLayer     *SurfaceSourceMetalLayer
	WaylandSurface *SurfaceSourceWaylandSurface
	XlibWindow     *SurfaceSourceXlibWindow
	WindowsHWND    *SurfaceSourceWindowsHWND
}

// SurfaceSourceMetalLayer contains a pointer to a Metal layer for creating a surface on macOS and iOS.
type SurfaceSourceMetalLayer struct {
	Layer unsafe.Pointer
}

// SurfaceSourceWaylandSurface contains pointers to a Wayland display and surface for creating a surface on Linux.
type SurfaceSourceWaylandSurface struct {
	Display unsafe.Pointer
	Surface unsafe.Pointer
}

// SurfaceSourceXlibWindow contains display and window handles for creating a surface using Xlib on Linux.
type SurfaceSourceXlibWindow struct {
	Display unsafe.Pointer
	Window  uint64
}

// SurfaceSourceWindowsHWND contains window handle and instance for creating a surface on Windows.
type SurfaceSourceWindowsHWND struct {
	Hwnd      unsafe.Pointer
	Hinstance unsafe.Pointer
}

// TexelCopyBufferInfo contains information about a buffer for texel copy operations, including layout and the buffer itself.
type TexelCopyBufferInfo struct {
	Layout TexelCopyBufferLayout
	Buffer *Buffer
}

// TexelCopyBufferLayout defines the layout of data in a buffer for texel copy operations, including offset, bytes per row, and rows per image.
type TexelCopyBufferLayout struct {
	Offset       uint64
	BytesPerRow  uint32
	RowsPerImage uint32
}

// TexelCopyTextureInfo contains information about a texture for texel copy operations, including the texture, mip level, origin, and aspect.
type TexelCopyTextureInfo struct {
	Texture  *Texture
	MipLevel uint32
	Origin   Origin3D
	Aspect   TextureAspect
}

// TextureBindingLayout describes the layout for a texture binding, including sample type, view dimension, and whether it is multisampled.
type TextureBindingLayout struct {
	SampleType    TextureSampleType
	ViewDimension TextureViewDimension
	Multisampled  bool
}

// TextureBindingViewDimension is used to specify the expected view dimension for a texture binding.
type TextureBindingViewDimension struct {
	TextureBindingViewDimension TextureViewDimension
}

// TextureComponentSwizzle describes how to swizzle the components of a texture view.
type TextureComponentSwizzle struct {
	R ComponentSwizzle
	G ComponentSwizzle
	B ComponentSwizzle
	A ComponentSwizzle
}

// TextureComponentSwizzleDescriptor describes a swizzle to be applied to a texture view.
type TextureComponentSwizzleDescriptor struct {
	Swizzle TextureComponentSwizzle
}

// TextureDescriptor describes a texture, including its dimensions, format, usage, mip level count, and sample count.
type TextureDescriptor struct {
	Label         string
	Usage         TextureUsage
	Dimension     TextureDimension
	Size          Extent3D
	Format        TextureFormat
	MipLevelCount uint32
	SampleCount   uint32
	ViewFormats   []TextureFormat
}

// TextureViewDescriptor describes a texture view, including format, dimension, mip level range, array layer range, and aspect.
type TextureViewDescriptor struct {
	Label           string
	Format          TextureFormat
	Dimension       TextureViewDimension
	BaseMipLevel    uint32
	MipLevelCount   uint32
	BaseArrayLayer  uint32
	ArrayLayerCount uint32
	Aspect          TextureAspect
	Usage           TextureUsage
}

type UncapturedErrorCallbackInfo struct {
	Callback UncapturedErrorCallback
}

// VertexAttribute describes a vertex attribute, including format, offset, and shader location.
type VertexAttribute struct {
	Format         VertexFormat
	Offset         uint64
	ShaderLocation uint32
}

// VertexBufferLayout describes the layout of a vertex buffer, including step mode, array stride, and attributes.
type VertexBufferLayout struct {
	StepMode    VertexStepMode
	ArrayStride uint64
	Attributes  []VertexAttribute
}

// VertexState describes the vertex stage of a render pipeline, including shader module, entry point, constants, and buffer layouts.
type VertexState struct {
	Module     *ShaderModule
	EntryPoint string
	Constants  []ConstantEntry
	Buffers    []VertexBufferLayout
}

// BufferMapCallback is called when a buffer mapping operation completes.
// The callback receives the status of the operation and any error message.
type BufferMapCallback func(status MapAsyncStatus, message string)

// CreateComputePipelineAsyncCallback is called when an asynchronous compute pipeline creation completes.
// The callback receives the status, the created pipeline (if successful), and any error message.
type CreateComputePipelineAsyncCallback func(status CreatePipelineAsyncStatus, pipeline *ComputePipeline, message string)

// CreateRenderPipelineAsyncCallback is called when an asynchronous render pipeline creation completes.
// The callback receives the status, the created pipeline (if successful), and any error message.
type CreateRenderPipelineAsyncCallback func(status CreatePipelineAsyncStatus, pipeline *RenderPipeline, message string)

// DeviceLostCallback is called when a device is lost.
// The callback receives the device, the reason for loss, and a message describing what happened.
type DeviceLostCallback func(device *Device, reason DeviceLostReason, message string)

// PopErrorScopeCallback is called when popping an error scope from the error scope stack.
// The callback receives the error type and any error message.
type PopErrorScopeCallback func(typ ErrorType, message string)

// QueueWorkDoneCallback is called when work submitted to a queue completes.
// The callback receives the status of the work and any error message.
type QueueWorkDoneCallback func(status QueueWorkDoneStatus, message string)

// UncapturedErrorCallback is called when an uncaptured error occurs on a device.
// The callback receives the device, the error type, and the error message.
type UncapturedErrorCallback func(device *Device, typ ErrorType, message string)

type compilationInfoCallback func(status compilationInfoRequestStatus, messages []CompilationMessage)
type requestAdapterCallback func(status requestAdapterStatus, adapter *Adapter, message string)
type requestDeviceCallback func(status requestDeviceStatus, device *Device, message string)
