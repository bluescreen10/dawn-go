// CODE GENERATED. DO NOT EDIT
//
//go:generate go run ./cmd/wrapper/.
package wgpu

import "unsafe"

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

type BindGroupDescriptor struct {
	Label   string
	Layout  *BindGroupLayout
	Entries []BindGroupEntry
}

type BindGroupEntry struct {
	Binding     uint32
	Buffer      *Buffer
	Offset      uint64
	Size        uint64
	Sampler     *Sampler
	TextureView *TextureView
}

type BindGroupLayoutDescriptor struct {
	Label   string
	Entries []BindGroupLayoutEntry
}

type BindGroupLayoutEntry struct {
	Binding          uint32
	Visibility       ShaderStage
	BindingArraySize uint32
	Buffer           BufferBindingLayout
	Sampler          SamplerBindingLayout
	Texture          TextureBindingLayout
	StorageTexture   StorageTextureBindingLayout
}

type BlendComponent struct {
	Operation BlendOperation
	SrcFactor BlendFactor
	DstFactor BlendFactor
}

type BlendState struct {
	Color BlendComponent
	Alpha BlendComponent
}

type BufferBindingLayout struct {
	Type             BufferBindingType
	HasDynamicOffset bool
	MinBindingSize   uint64
}

type BufferDescriptor struct {
	Label            string
	Usage            BufferUsage
	Size             uint64
	MappedAtCreation bool
}

type BufferInitDescriptor struct {
	Label    string
	Usage    BufferUsage
	Contents []byte
}

type bufferMapCallbackInfo struct {
	Mode     callbackMode
	Callback BufferMapCallback
}

type Color struct {
	R float64
	G float64
	B float64
	A float64
}

type ColorTargetState struct {
	Format    TextureFormat
	Blend     *BlendState
	WriteMask ColorWriteMask
}

type CommandBufferDescriptor struct {
	Label string
}

type CommandEncoderDescriptor struct {
	Label string
}

type CompatibilityModeLimits struct {
	MaxStorageBuffersInVertexStage    uint32
	MaxStorageTexturesInVertexStage   uint32
	MaxStorageBuffersInFragmentStage  uint32
	MaxStorageTexturesInFragmentStage uint32
}

type CompilationMessage struct {
	Message string
	Type    CompilationMessageType
	LineNum uint64
	LinePos uint64
	Offset  uint64
	Length  uint64
}

type ComputePassDescriptor struct {
	Label           string
	TimestampWrites *PassTimestampWrites
}

type ComputePipelineDescriptor struct {
	Label   string
	Layout  *PipelineLayout
	Compute ComputeState
}

type ComputeState struct {
	Module        *ShaderModule
	EntryPoint    string
	ConstantCount int
	Constants     ConstantEntry
}

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

type Extent3D struct {
	Width              uint32
	Height             uint32
	DepthOrArrayLayers uint32
}

type ExternalTextureBindingEntry struct {
	ExternalTexture *ExternalTexture
}

type ExternalTextureBindingLayout struct {
}

type FragmentState struct {
	Module     *ShaderModule
	EntryPoint string
	Constants  []ConstantEntry
	Targets    []ColorTargetState
}

type Future struct {
	Id uint64
}

type FutureWaitInfo struct {
	Future    Future
	Completed bool
}

type InstanceDescriptor struct {
	RequiredFeatures []InstanceFeatureName
	RequiredLimits   *InstanceLimits
}

type InstanceLimits struct {
	TimedWaitAnyMaxCount int
}

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

type MultisampleState struct {
	Count                  uint32
	Mask                   uint32
	AlphaToCoverageEnabled bool
}

type Origin3D struct {
	X uint32
	Y uint32
	Z uint32
}

type PassTimestampWrites struct {
	QuerySet                  *QuerySet
	BeginningOfPassWriteIndex uint32
	EndOfPassWriteIndex       uint32
}

type PipelineLayoutDescriptor struct {
	Label                string
	BindGroupLayoutCount int
	BindGroupLayouts     *BindGroupLayout
	ImmediateSize        uint32
}

type popErrorScopeCallbackInfo struct {
	Mode     callbackMode
	Callback popErrorScopeCallback
}

type PrimitiveState struct {
	Topology         PrimitiveTopology
	StripIndexFormat IndexFormat
	FrontFace        FrontFace
	CullMode         CullMode
	UnclippedDepth   bool
}

type QuerySetDescriptor struct {
	Label string
	Type  QueryType
	Count uint32
}

type QueueDescriptor struct {
	Label string
}

type QueueWorkDoneCallbackInfo struct {
	Mode     callbackMode
	Callback QueueWorkDoneCallback
}

type RenderBundleDescriptor struct {
	Label string
}

type RenderBundleEncoderDescriptor struct {
	Label              string
	ColorFormats       []TextureFormat
	DepthStencilFormat TextureFormat
	SampleCount        uint32
	DepthReadOnly      bool
	StencilReadOnly    bool
}

type RenderPassColorAttachment struct {
	View          *TextureView
	DepthSlice    uint32
	ResolveTarget *TextureView
	LoadOp        LoadOp
	StoreOp       StoreOp
	ClearValue    Color
}

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

type RenderPassDescriptor struct {
	Label                  string
	ColorAttachments       []RenderPassColorAttachment
	DepthStencilAttachment *RenderPassDepthStencilAttachment
	OcclusionQuerySet      *QuerySet
	TimestampWrites        *PassTimestampWrites
}

type RenderPassMaxDrawCount struct {
	MaxDrawCount uint64
}

type RenderPipelineDescriptor struct {
	Label        string
	Layout       *PipelineLayout
	Vertex       VertexState
	Primitive    PrimitiveState
	DepthStencil *DepthStencilState
	Multisample  MultisampleState
	Fragment     *FragmentState
}

type RequestAdapterWebXROptions struct {
	XrCompatible bool
}

type requestAdapterCallbackInfo struct {
	Mode     callbackMode
	Callback requestAdapterCallback
}

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

type SamplerBindingLayout struct {
	Type SamplerBindingType
}

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

type ShaderModuleDescriptor struct {
	Label string

	SPIRVSource *ShaderSourceSPIRV
	WGSLSource  *ShaderSourceWGSL
}

type ShaderSourceSPIRV struct {
	Code []uint32
}

type ShaderSourceWGSL struct {
	Code string
}

type StencilFaceState struct {
	Compare     CompareFunction
	FailOp      StencilOperation
	DepthFailOp StencilOperation
	PassOp      StencilOperation
}

type StorageTextureBindingLayout struct {
	Access        StorageTextureAccess
	Format        TextureFormat
	ViewDimension TextureViewDimension
}

type SurfaceCapabilities struct {
	Usages       TextureUsage
	Formats      []TextureFormat
	PresentModes []PresentMode
	AlphaModes   []CompositeAlphaMode
}

type SurfaceColorManagement struct {
	ColorSpace      PredefinedColorSpace
	ToneMappingMode ToneMappingMode
}

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

type SurfaceDescriptor struct {
	Label      string
	MetalLayer *SurfaceSourceMetalLayer
}

type SurfaceSourceMetalLayer struct {
	Layer unsafe.Pointer
}

type TexelCopyBufferInfo struct {
	Layout TexelCopyBufferLayout
	Buffer *Buffer
}

type TexelCopyBufferLayout struct {
	Offset       uint64
	BytesPerRow  uint32
	RowsPerImage uint32
}

type TexelCopyTextureInfo struct {
	Texture  *Texture
	MipLevel uint32
	Origin   Origin3D
	Aspect   TextureAspect
}

type TextureBindingLayout struct {
	SampleType    TextureSampleType
	ViewDimension TextureViewDimension
	Multisampled  bool
}

type TextureBindingViewDimension struct {
	TextureBindingViewDimension TextureViewDimension
}

type TextureComponentSwizzle struct {
	R ComponentSwizzle
	G ComponentSwizzle
	B ComponentSwizzle
	A ComponentSwizzle
}

type TextureComponentSwizzleDescriptor struct {
	Swizzle TextureComponentSwizzle
}

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

type VertexAttribute struct {
	Format         VertexFormat
	Offset         uint64
	ShaderLocation uint32
}

type VertexBufferLayout struct {
	StepMode    VertexStepMode
	ArrayStride uint64
	Attributes  []VertexAttribute
}

type VertexState struct {
	Module     *ShaderModule
	EntryPoint string
	Constants  []ConstantEntry
	Buffers    []VertexBufferLayout
}

type BufferMapCallback func(status MapAsyncStatus, message string)

type CreateComputePipelineAsyncCallback func(status CreatePipelineAsyncStatus, pipeline *ComputePipeline, message string)
type CreateRenderPipelineAsyncCallback func(status CreatePipelineAsyncStatus, pipeline *RenderPipeline, message string)
type DeviceLostCallback func(device *Device, reason DeviceLostReason, message string)
type QueueWorkDoneCallback func(status QueueWorkDoneStatus, message string)
type UncapturedErrorCallback func(device *Device, typ ErrorType, message string)

type compilationInfoCallback func(status compilationInfoRequestStatus, messages []CompilationMessage)
type popErrorScopeCallback func(status popErrorScopeStatus, typ ErrorType, message string)
type requestAdapterCallback func(status requestAdapterStatus, adapter *Adapter, message string)
type requestDeviceCallback func(status requestDeviceStatus, device *Device, message string)
