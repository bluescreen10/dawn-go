// CODE GENERATED. DO NOT EDIT
//go:generate go run ./cmd/wrapper/. 
package wgpu

type ComputeState struct{
	Module *ShaderModule
	EntryPoint string
	ConstantCount int
	Constants ConstantEntry
}

type BufferMapCallbackInfo struct{
	Mode CallbackMode
	Callback BufferMapCallback
}

type Future struct{
	Id uint64
}

type InstanceDescriptor struct{
	RequiredFeatureCount int
	RequiredFeatures InstanceFeatureName
	RequiredLimits InstanceLimits
}

type FragmentState struct{
	Module *ShaderModule
	EntryPoint string
	ConstantCount int
	Constants ConstantEntry
	TargetCount int
	Targets ColorTargetState
}

type StorageTextureBindingLayout struct{
	Access StorageTextureAccess
	Format TextureFormat
	ViewDimension TextureViewDimension
}

type ConstantEntry struct{
	Key string
	Value float64
}

type CommandBufferDescriptor struct{
	Label string
}

type ComputePassDescriptor struct{
	Label string
	TimestampWrites PassTimestampWrites
}

type VertexBufferLayout struct{
	StepMode VertexStepMode
	ArrayStride uint64
	AttributeCount int
	Attributes VertexAttribute
}

type StencilFaceState struct{
	Compare CompareFunction
	FailOp StencilOperation
	DepthFailOp StencilOperation
	PassOp StencilOperation
}

type Limits struct{
	MaxTextureDimension1D uint32
	MaxTextureDimension2D uint32
	MaxTextureDimension3D uint32
	MaxTextureArrayLayers uint32
	MaxBindGroups uint32
	MaxBindGroupsPlusVertexBuffers uint32
	MaxBindingsPerBindGroup uint32
	MaxDynamicUniformBuffersPerPipelineLayout uint32
	MaxDynamicStorageBuffersPerPipelineLayout uint32
	MaxSampledTexturesPerShaderStage uint32
	MaxSamplersPerShaderStage uint32
	MaxStorageBuffersPerShaderStage uint32
	MaxStorageTexturesPerShaderStage uint32
	MaxUniformBuffersPerShaderStage uint32
	MaxUniformBufferBindingSize uint64
	MaxStorageBufferBindingSize uint64
	MinUniformBufferOffsetAlignment uint32
	MinStorageBufferOffsetAlignment uint32
	MaxVertexBuffers uint32
	MaxBufferSize uint64
	MaxVertexAttributes uint32
	MaxVertexBufferArrayStride uint32
	MaxInterStageShaderVariables uint32
	MaxColorAttachments uint32
	MaxColorAttachmentBytesPerSample uint32
	MaxComputeWorkgroupStorageSize uint32
	MaxComputeInvocationsPerWorkgroup uint32
	MaxComputeWorkgroupSizeX uint32
	MaxComputeWorkgroupSizeY uint32
	MaxComputeWorkgroupSizeZ uint32
	MaxComputeWorkgroupsPerDimension uint32
	MaxImmediateSize uint32
}

type TexelCopyTextureInfo struct{
	Texture *Texture
	MipLevel uint32
	Origin Origin3D
	Aspect TextureAspect
}

type BlendState struct{
	Color BlendComponent
	Alpha BlendComponent
}

type SurfaceColorManagement struct{
	ColorSpace PredefinedColorSpace
	ToneMappingMode ToneMappingMode
}

type CompilationInfo struct{
	MessageCount int
	Messages CompilationMessage
}

type ShaderModuleDescriptor struct{
	Label string
}

type ColorTargetState struct{
	Format TextureFormat
	Blend BlendState
	WriteMask ColorWriteMask
}

type SurfaceCapabilities struct{
	Usages TextureUsage
	FormatCount int
	Formats TextureFormat
	PresentModeCount int
	PresentModes PresentMode
	AlphaModeCount int
	AlphaModes CompositeAlphaMode
}

type PopErrorScopeCallbackInfo struct{
	Mode CallbackMode
	Callback PopErrorScopeCallback
}

type SupportedWGSLLanguageFeatures struct{
	FeatureCount int
	Features WGSLLanguageFeatureName
}

type FutureWaitInfo struct{
	Future Future
	Completed bool
}

type CompilationInfoCallbackInfo struct{
	Mode CallbackMode
	Callback CompilationInfoCallback
}

type TextureBindingViewDimension struct{
	TextureBindingViewDimension TextureViewDimension
}

type BindGroupLayoutDescriptor struct{
	Label string
	EntryCount int
	Entries BindGroupLayoutEntry
}

type TextureViewDescriptor struct{
	Label string
	Format TextureFormat
	Dimension TextureViewDimension
	BaseMipLevel uint32
	MipLevelCount uint32
	BaseArrayLayer uint32
	ArrayLayerCount uint32
	Aspect TextureAspect
	Usage TextureUsage
}

type ShaderSourceWGSL struct{
	Code string
}

type SurfaceDescriptor struct{
	Label string
}

type VertexAttribute struct{
	Format VertexFormat
	Offset uint64
	ShaderLocation uint32
}

type DepthStencilState struct{
	Format TextureFormat
	DepthWriteEnabled OptionalBool
	DepthCompare CompareFunction
	StencilFront StencilFaceState
	StencilBack StencilFaceState
	StencilReadMask uint32
	StencilWriteMask uint32
	DepthBias int32
	DepthBiasSlopeScale float32
	DepthBiasClamp float32
}

type Origin3D struct{
	X uint32
	Y uint32
	Z uint32
}

type ExternalTextureBindingEntry struct{
	ExternalTexture *ExternalTexture
}

type BufferDescriptor struct{
	Label string
	Usage BufferUsage
	Size uint64
	MappedAtCreation bool
}

type Extent3D struct{
	Width uint32
	Height uint32
	DepthOrArrayLayers uint32
}

type ComputePipelineDescriptor struct{
	Label string
	Layout *PipelineLayout
	Compute ComputeState
}

type SurfaceConfiguration struct{
	Device *Device
	Format TextureFormat
	Usage TextureUsage
	Width uint32
	Height uint32
	ViewFormatCount int
	ViewFormats TextureFormat
	AlphaMode CompositeAlphaMode
	PresentMode PresentMode
}

type PrimitiveState struct{
	Topology PrimitiveTopology
	StripIndexFormat IndexFormat
	FrontFace FrontFace
	CullMode CullMode
	UnclippedDepth bool
}

type SurfaceTexture struct{
	Texture *Texture
	Status SurfaceGetCurrentTextureStatus
}

type RequestAdapterWebXROptions struct{
	XrCompatible bool
}

type RequestDeviceCallbackInfo struct{
	Mode CallbackMode
	Callback RequestDeviceCallback
}

type RenderBundleEncoderDescriptor struct{
	Label string
	ColorFormatCount int
	ColorFormats TextureFormat
	DepthStencilFormat TextureFormat
	SampleCount uint32
	DepthReadOnly bool
	StencilReadOnly bool
}

type CreateComputePipelineAsyncCallbackInfo struct{
	Mode CallbackMode
	Callback CreateComputePipelineAsyncCallback
}

type CompilationMessage struct{
	Message string
	Type CompilationMessageType
	LineNum uint64
	LinePos uint64
	Offset uint64
	Length uint64
}

type AdapterInfo struct{
	Vendor string
	Architecture string
	Device string
	Description string
	BackendType BackendType
	AdapterType AdapterType
	VendorID uint32
	DeviceID uint32
	SubgroupMinSize uint32
	SubgroupMaxSize uint32
}

type SamplerDescriptor struct{
	Label string
	AddressModeU AddressMode
	AddressModeV AddressMode
	AddressModeW AddressMode
	MagFilter FilterMode
	MinFilter FilterMode
	MipmapFilter MipmapFilterMode
	LodMinClamp float32
	LodMaxClamp float32
	Compare CompareFunction
	MaxAnisotropy uint16
}

type CreateRenderPipelineAsyncCallbackInfo struct{
	Mode CallbackMode
	Callback CreateRenderPipelineAsyncCallback
}

type QueueWorkDoneCallbackInfo struct{
	Mode CallbackMode
	Callback QueueWorkDoneCallback
}

type RequestAdapterOptions struct{
	FeatureLevel FeatureLevel
	PowerPreference PowerPreference
	ForceFallbackAdapter bool
	BackendType BackendType
	CompatibleSurface *Surface
}

type ShaderSourceSPIRV struct{
	CodeSize uint32
	Code uint32
}

type PassTimestampWrites struct{
	QuerySet *QuerySet
	BeginningOfPassWriteIndex uint32
	EndOfPassWriteIndex uint32
}

type TextureComponentSwizzleDescriptor struct{
	Swizzle TextureComponentSwizzle
}

type ExternalTextureBindingLayout struct{
}

type TexelCopyBufferInfo struct{
	Layout TexelCopyBufferLayout
	Buffer *Buffer
}

type PipelineLayoutDescriptor struct{
	Label string
	BindGroupLayoutCount int
	BindGroupLayouts *BindGroupLayout
	ImmediateSize uint32
}

type SamplerBindingLayout struct{
	Type SamplerBindingType
}

type QuerySetDescriptor struct{
	Label string
	Type QueryType
	Count uint32
}

type BufferBindingLayout struct{
	Type BufferBindingType
	HasDynamicOffset bool
	MinBindingSize uint64
}

type CommandEncoderDescriptor struct{
	Label string
}

type DeviceLostCallbackInfo struct{
	Mode CallbackMode
	Callback DeviceLostCallback
}

type Color struct{
	R float64
	G float64
	B float64
	A float64
}

type SupportedInstanceFeatures struct{
	FeatureCount int
	Features InstanceFeatureName
}

type VertexState struct{
	Module *ShaderModule
	EntryPoint string
	ConstantCount int
	Constants ConstantEntry
	BufferCount int
	Buffers VertexBufferLayout
}

type TexelCopyBufferLayout struct{
	Offset uint64
	BytesPerRow uint32
	RowsPerImage uint32
}

type RenderPassColorAttachment struct{
	View *TextureView
	DepthSlice uint32
	ResolveTarget *TextureView
	LoadOp LoadOp
	StoreOp StoreOp
	ClearValue Color
}

type BindGroupEntry struct{
	Binding uint32
	Buffer *Buffer
	Offset uint64
	Size uint64
	Sampler *Sampler
	TextureView *TextureView
}

type TextureComponentSwizzle struct{
	R ComponentSwizzle
	G ComponentSwizzle
	B ComponentSwizzle
	A ComponentSwizzle
}

type BlendComponent struct{
	Operation BlendOperation
	SrcFactor BlendFactor
	DstFactor BlendFactor
}

type BindGroupLayoutEntry struct{
	Binding uint32
	Visibility ShaderStage
	BindingArraySize uint32
	Buffer BufferBindingLayout
	Sampler SamplerBindingLayout
	Texture TextureBindingLayout
	StorageTexture StorageTextureBindingLayout
}

type TextureBindingLayout struct{
	SampleType TextureSampleType
	ViewDimension TextureViewDimension
	Multisampled bool
}

type SupportedFeatures struct{
	FeatureCount int
	Features FeatureName
}

type RenderPassMaxDrawCount struct{
	MaxDrawCount uint64
}

type RenderBundleDescriptor struct{
	Label string
}

type UncapturedErrorCallbackInfo struct{
	Callback UncapturedErrorCallback
}

type BindGroupDescriptor struct{
	Label string
	Layout *BindGroupLayout
	EntryCount int
	Entries BindGroupEntry
}

type MultisampleState struct{
	Count uint32
	Mask uint32
	AlphaToCoverageEnabled bool
}

type CompatibilityModeLimits struct{
	MaxStorageBuffersInVertexStage uint32
	MaxStorageTexturesInVertexStage uint32
	MaxStorageBuffersInFragmentStage uint32
	MaxStorageTexturesInFragmentStage uint32
}

type RenderPassDescriptor struct{
	Label string
	ColorAttachmentCount int
	ColorAttachments RenderPassColorAttachment
	DepthStencilAttachment RenderPassDepthStencilAttachment
	OcclusionQuerySet *QuerySet
	TimestampWrites PassTimestampWrites
}

type InstanceLimits struct{
	TimedWaitAnyMaxCount int
}

type RequestAdapterCallbackInfo struct{
	Mode CallbackMode
	Callback RequestAdapterCallback
}

type RenderPipelineDescriptor struct{
	Label string
	Layout *PipelineLayout
	Vertex VertexState
	Primitive PrimitiveState
	DepthStencil DepthStencilState
	Multisample MultisampleState
	Fragment FragmentState
}

type TextureDescriptor struct{
	Label string
	Usage TextureUsage
	Dimension TextureDimension
	Size Extent3D
	Format TextureFormat
	MipLevelCount uint32
	SampleCount uint32
	ViewFormatCount int
	ViewFormats TextureFormat
}

type RenderPassDepthStencilAttachment struct{
	View *TextureView
	DepthLoadOp LoadOp
	DepthStoreOp StoreOp
	DepthClearValue float32
	DepthReadOnly bool
	StencilLoadOp LoadOp
	StencilStoreOp StoreOp
	StencilClearValue uint32
	StencilReadOnly bool
}

type QueueDescriptor struct{
	Label string
}

type DeviceDescriptor struct{
	Label string
	RequiredFeatureCount int
	RequiredFeatures FeatureName
	RequiredLimits Limits
	DefaultQueue QueueDescriptor
	DeviceLostCallbackInfo DeviceLostCallbackInfo
	UncapturedErrorCallbackInfo UncapturedErrorCallbackInfo
}

type CreateRenderPipelineAsyncCallback func (status CreatePipelineAsyncStatus, pipeline *RenderPipeline, message string)
type RequestDeviceCallback func (status RequestDeviceStatus, device *Device, message string)
type CompilationInfoCallback func (status CompilationInfoRequestStatus, compilationInfo CompilationInfo)
type QueueWorkDoneCallback func (status QueueWorkDoneStatus, message string)
type PopErrorScopeCallback func (status PopErrorScopeStatus, typ ErrorType, message string)
type BufferMapCallback func (status MapAsyncStatus, message string)
type DeviceLostCallback func (device *Device, reason DeviceLostReason, message string)
type CreateComputePipelineAsyncCallback func (status CreatePipelineAsyncStatus, pipeline *ComputePipeline, message string)
type RequestAdapterCallback func (status RequestAdapterStatus, adapter *Adapter, message string)
type UncapturedErrorCallback func (device *Device, typ ErrorType, message string)
