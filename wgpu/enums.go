package wgpu

/*
#include "webgpu.h"
*/
import "C"

type WGSLLanguageFeatureName C.WGPUWGSLLanguageFeatureName

const (
	WGSLLanguageFeatureNameReadonlyAndReadwriteStorageTextures WGSLLanguageFeatureName = 1
	WGSLLanguageFeatureNamePacked4x8IntegerDotProduct          WGSLLanguageFeatureName = 2
	WGSLLanguageFeatureNameUnrestrictedPointerParameters       WGSLLanguageFeatureName = 3
	WGSLLanguageFeatureNamePointerCompositeAccess              WGSLLanguageFeatureName = 4
	WGSLLanguageFeatureNameUniformBufferStandardLayout         WGSLLanguageFeatureName = 5
	WGSLLanguageFeatureNameSubgroupId                          WGSLLanguageFeatureName = 6
	WGSLLanguageFeatureNameTextureAndSamplerLet                WGSLLanguageFeatureName = 7
	WGSLLanguageFeatureNameSubgroupUniformity                  WGSLLanguageFeatureName = 8
	WGSLLanguageFeatureNameTextureFormatsTier1                 WGSLLanguageFeatureName = 9
)

func (f WGSLLanguageFeatureName) String() string {
	switch f {
	case WGSLLanguageFeatureNameReadonlyAndReadwriteStorageTextures:
		return "Readonly And Readwrite Storage Textures"
	case WGSLLanguageFeatureNamePacked4x8IntegerDotProduct:
		return "Packed4x8 Integer Dot Product"
	case WGSLLanguageFeatureNameUnrestrictedPointerParameters:
		return "Unrestricted Pointer Parameters"
	case WGSLLanguageFeatureNamePointerCompositeAccess:
		return "Pointer Composite Access"
	case WGSLLanguageFeatureNameUniformBufferStandardLayout:
		return "Uniform Buffer Standard Layout"
	case WGSLLanguageFeatureNameSubgroupId:
		return "Subgroup Id"
	case WGSLLanguageFeatureNameTextureAndSamplerLet:
		return "Texture And Sampler Let"
	case WGSLLanguageFeatureNameSubgroupUniformity:
		return "Subgroup Uniformity"
	case WGSLLanguageFeatureNameTextureFormatsTier1:
		return "TextureFormatsTier1"
	default:
		return "Unknown"
	}
}

type AdapterType C.WGPUAdapterType

const (
	AdapterTypeDiscreteGPU   AdapterType = 1
	AdapterTypeIntegratedGPU AdapterType = 2
	AdapterTypeCPU           AdapterType = 3
	AdapterTypeUnknown       AdapterType = 4
)

func (a AdapterType) String() string {
	switch a {
	case AdapterTypeDiscreteGPU:
		return "Discrete GPU"
	case AdapterTypeIntegratedGPU:
		return "Integrated GPU"
	case AdapterTypeCPU:
		return "CPU"
	default:
		return "Unknown"
	}
}

type AddressMode C.WGPUAddressMode

const (
	AddressModeUndefined    AddressMode = 0
	AddressModeClampToEdge  AddressMode = 1
	AddressModeRepeat       AddressMode = 2
	AddressModeMirrorRepeat AddressMode = 3
)

type BackendType C.WGPUBackendType

const (
	BackendTypeUndefined BackendType = 0
	BackendTypeNull      BackendType = 1
	BackendTypeWebGPU    BackendType = 2
	BackendTypeD3D11     BackendType = 3
	BackendTypeD3D12     BackendType = 4
	BackendTypeMetal     BackendType = 5
	BackendTypeVulkan    BackendType = 6
	BackendTypeOpenGL    BackendType = 7
	BackendTypeOpenGLES  BackendType = 8
)

func (b BackendType) String() string {
	switch b {
	case BackendTypeNull:
		return "null"
	case BackendTypeWebGPU:
		return "Webgpu"
	case BackendTypeD3D11:
		return "D3D11"
	case BackendTypeD3D12:
		return "D3D12"
	case BackendTypeMetal:
		return "Metal"
	case BackendTypeVulkan:
		return "Vulkan"
	case BackendTypeOpenGL:
		return "OpenGL"
	case BackendTypeOpenGLES:
		return "OpenGLES"
	default:
		return "undefined"
	}
}

type BlendFactor C.WGPUBlendFactor

const (
	BlendFactorUndefined         BlendFactor = 0
	BlendFactorZero              BlendFactor = 1
	BlendFactorOne               BlendFactor = 2
	BlendFactorSrc               BlendFactor = 3
	BlendFactorOneMinusSrc       BlendFactor = 4
	BlendFactorSrcAlpha          BlendFactor = 5
	BlendFactorOneMinusSrcAlpha  BlendFactor = 6
	BlendFactorDst               BlendFactor = 7
	BlendFactorOneMinusDst       BlendFactor = 8
	BlendFactorDstAlpha          BlendFactor = 9
	BlendFactorOneMinusDstAlpha  BlendFactor = 10
	BlendFactorSrcAlphaSaturated BlendFactor = 11
	BlendFactorConstant          BlendFactor = 12
	BlendFactorOneMinusConstant  BlendFactor = 13
	BlendFactorSrc1              BlendFactor = 14
	BlendFactorOneMinusSrc1      BlendFactor = 15
	BlendFactorSrc1Alpha         BlendFactor = 16
	BlendFactorOneMinusSrc1Alpha BlendFactor = 17
)

type BlendOperation C.WGPUBlendOperation

const (
	BlendOperationUndefined       BlendOperation = 0
	BlendOperationAdd             BlendOperation = 1
	BlendOperationSubtract        BlendOperation = 2
	BlendOperationReverseSubtract BlendOperation = 3
	BlendOperationMin             BlendOperation = 4
	BlendOperationMax             BlendOperation = 5
)

type BufferBindingType C.WGPUBufferBindingType

const (
	BufferBindingTypeBindingNotUsed  BufferBindingType = 0
	BufferBindingTypeUndefined       BufferBindingType = 1
	BufferBindingTypeUniform         BufferBindingType = 2
	BufferBindingTypeStorage         BufferBindingType = 3
	BufferBindingTypeReadOnlyStorage BufferBindingType = 4
)

type BufferMapState C.WGPUBufferMapState

const (
	BufferMapStateUnmapped BufferMapState = 1
	BufferMapStatePending  BufferMapState = 2
	BufferMapStateMapped   BufferMapState = 3
)

type BufferUsage C.WGPUBufferUsage

const (
	BufferUsageNone         BufferUsage = 0
	BufferUsageMapRead      BufferUsage = 1
	BufferUsageMapWrite     BufferUsage = 2
	BufferUsageCopySrc      BufferUsage = 4
	BufferUsageCopyDst      BufferUsage = 8
	BufferUsageIndex        BufferUsage = 16
	BufferUsageVertex       BufferUsage = 32
	BufferUsageUniform      BufferUsage = 64
	BufferUsageStorage      BufferUsage = 128
	BufferUsageIndirect     BufferUsage = 256
	BufferUsageQueryResolve BufferUsage = 512
	BufferUsageTexelBuffer  BufferUsage = 1024
)

type callbackMode C.WGPUCallbackMode

const (
	callbackModeWaitAnyOnly        callbackMode = 1
	callbackModeAllowProcessEvents callbackMode = 2
	callbackModeAllowSpontaneous   callbackMode = 3
)

type ColorWriteMask C.WGPUColorWriteMask

const (
	ColorWriteMaskNone  ColorWriteMask = 0
	ColorWriteMaskRed   ColorWriteMask = 1
	ColorWriteMaskGreen ColorWriteMask = 2
	ColorWriteMaskBlue  ColorWriteMask = 4
	ColorWriteMaskAlpha ColorWriteMask = 8
	ColorWriteMaskAll   ColorWriteMask = 15
)

type CompareFunction C.WGPUCompareFunction

const (
	CompareFunctionUndefined    CompareFunction = 0
	CompareFunctionNever        CompareFunction = 1
	CompareFunctionLess         CompareFunction = 2
	CompareFunctionEqual        CompareFunction = 3
	CompareFunctionLessEqual    CompareFunction = 4
	CompareFunctionGreater      CompareFunction = 5
	CompareFunctionNotEqual     CompareFunction = 6
	CompareFunctionGreaterEqual CompareFunction = 7
	CompareFunctionAlways       CompareFunction = 8
)

type compilationInfoRequestStatus C.WGPUCompilationInfoRequestStatus

const (
	compilationInfoRequestStatusSuccess           compilationInfoRequestStatus = 1
	compilationInfoRequestStatusCallbackCancelled compilationInfoRequestStatus = 2
)

type CompilationMessageType C.WGPUCompilationMessageType

const (
	CompilationMessageTypeError   CompilationMessageType = 1
	CompilationMessageTypeWarning CompilationMessageType = 2
	CompilationMessageTypeInfo    CompilationMessageType = 3
)

type ComponentSwizzle C.WGPUComponentSwizzle

const (
	ComponentSwizzleUndefined ComponentSwizzle = 0
	ComponentSwizzleZero      ComponentSwizzle = 1
	ComponentSwizzleOne       ComponentSwizzle = 2
	ComponentSwizzleR         ComponentSwizzle = 3
	ComponentSwizzleG         ComponentSwizzle = 4
	ComponentSwizzleB         ComponentSwizzle = 5
	ComponentSwizzleA         ComponentSwizzle = 6
)

type CompositeAlphaMode C.WGPUCompositeAlphaMode

const (
	CompositeAlphaModeAuto            CompositeAlphaMode = 0
	CompositeAlphaModeOpaque          CompositeAlphaMode = 1
	CompositeAlphaModePremultiplied   CompositeAlphaMode = 2
	CompositeAlphaModeUnpremultiplied CompositeAlphaMode = 3
	CompositeAlphaModeInherit         CompositeAlphaMode = 4
)

type CreatePipelineAsyncStatus C.WGPUCreatePipelineAsyncStatus

const (
	CreatePipelineAsyncStatusSuccess           CreatePipelineAsyncStatus = 1
	CreatePipelineAsyncStatusCallbackCancelled CreatePipelineAsyncStatus = 2
	CreatePipelineAsyncStatusValidationError   CreatePipelineAsyncStatus = 3
	CreatePipelineAsyncStatusInternalError     CreatePipelineAsyncStatus = 4
)

type CullMode C.WGPUCullMode

const (
	CullModeUndefined CullMode = 0
	CullModeNone      CullMode = 1
	CullModeFront     CullMode = 2
	CullModeBack      CullMode = 3
)

type DeviceLostReason C.WGPUDeviceLostReason

const (
	DeviceLostReasonUnknown           DeviceLostReason = 1
	DeviceLostReasonDestroyed         DeviceLostReason = 2
	DeviceLostReasonCallbackCancelled DeviceLostReason = 3
	DeviceLostReasonFailedCreation    DeviceLostReason = 4
)

type ErrorFilter C.WGPUErrorType

const (
	ErrorFilterValidation  ErrorFilter = 1
	ErrorFilterOutOfMemory ErrorFilter = 2
	ErrorFilterInternal    ErrorFilter = 3
)

func (e ErrorFilter) String() string {
	switch e {
	case ErrorFilterValidation:
		return "validation error"
	case ErrorFilterOutOfMemory:
		return "out of memory error"
	case ErrorFilterInternal:
		return "internal error"
	default:
		return ""
	}
}

type ErrorType C.WGPUErrorType

const (
	ErrorTypeNoError     ErrorType = 1
	ErrorTypeValidation  ErrorType = 2
	ErrorTypeOutOfMemory ErrorType = 3
	ErrorTypeInternal    ErrorType = 4
	ErrorTypeUnknown     ErrorType = 5
)

func (e ErrorType) String() string {
	switch e {
	case ErrorTypeValidation:
		return "validation"
	case ErrorTypeOutOfMemory:
		return "out of memory"
	case ErrorTypeInternal:
		return "internal"
	case ErrorTypeUnknown:
		return "unknown"
	default:
		return "no error"
	}
}

type FeatureLevel C.WGPUFeatureLevel

const (
	FeatureLevelUndefined     FeatureLevel = 0
	FeatureLevelCompatibility FeatureLevel = 1
	FeatureLevelCore          FeatureLevel = 2
)

type FeatureName C.WGPUFeatureName

const (
	FeatureNameCoreFeaturesAndLimits          FeatureName = 1
	FeatureNameDepthClipControl               FeatureName = 2
	FeatureNameDepth32FloatStencil8           FeatureName = 3
	FeatureNameTextureCompressionBC           FeatureName = 4
	FeatureNameTextureCompressionBCSliced3D   FeatureName = 5
	FeatureNameTextureCompressionETC2         FeatureName = 6
	FeatureNameTextureCompressionASTC         FeatureName = 7
	FeatureNameTextureCompressionASTCSliced3D FeatureName = 8
	FeatureNameTimestampQuery                 FeatureName = 9
	FeatureNameIndirectFirstInstance          FeatureName = 10
	FeatureNameShaderF16                      FeatureName = 11
	FeatureNameRG11B10UfloatRenderable        FeatureName = 12
	FeatureNameBGRA8UnormStorage              FeatureName = 13
	FeatureNameFloat32Filterable              FeatureName = 14
	FeatureNameFloat32Blendable               FeatureName = 15
	FeatureNameClipDistances                  FeatureName = 16
	FeatureNameDualSourceBlending             FeatureName = 17
	FeatureNameSubgroups                      FeatureName = 18
	FeatureNameTextureFormatsTier1            FeatureName = 19
	FeatureNameTextureFormatsTier2            FeatureName = 20
	FeatureNamePrimitiveIndex                 FeatureName = 21
	FeatureNameTextureComponentSwizzle        FeatureName = 22
)

type FilterMode C.WGPUFilterMode

const (
	FilterModeUndefined FilterMode = 0
	FilterModeNearest   FilterMode = 1
	FilterModeLinear    FilterMode = 2
)

type FrontFace C.WGPUFrontFace

const (
	FrontFaceUndefined FrontFace = 0
	FrontFaceCCW       FrontFace = 1
	FrontFaceCW        FrontFace = 2
)

type IndexFormat C.WGPUIndexFormat

const (
	IndexFormatUndefined IndexFormat = 0
	IndexFormatUint16    IndexFormat = 1
	IndexFormatUint32    IndexFormat = 2
)

type InstanceFeatureName C.WGPUInstanceFeatureName

const (
	InstanceFeatureNameTimedWaitAny              InstanceFeatureName = 1
	InstanceFeatureNameShaderSourceSPIRV         InstanceFeatureName = 2
	InstanceFeatureNameMultipleDevicesPerAdapter InstanceFeatureName = 3
)

type LoadOp C.WGPULoadOp

const (
	LoadOpUndefined            LoadOp = 0
	LoadOpLoad                 LoadOp = 1
	LoadOpClear                LoadOp = 2
	LoadOpExpandResolveTexture LoadOp = 3
)

type MapAsyncStatus C.WGPUMapAsyncStatus

const (
	MapAsyncStatusSuccess           MapAsyncStatus = 1
	MapAsyncStatusCallbackCancelled MapAsyncStatus = 2
	MapAsyncStatusError             MapAsyncStatus = 3
	MapAsyncStatusAborted           MapAsyncStatus = 4
)

type MapMode C.WGPUMapMode

const (
	MapModeNone  MapMode = 0
	MapModeRead  MapMode = 1
	MapModeWrite MapMode = 2
)

type MipmapFilterMode C.WGPUMipmapFilterMode

const (
	MipmapFilterModeUndefined MipmapFilterMode = 0
	MipmapFilterModeNearest   MipmapFilterMode = 1
	MipmapFilterModeLinear    MipmapFilterMode = 2
)

type OptionalBool C.WGPUOptionalBool

const (
	OptionalBoolFalse     OptionalBool = 0
	OptionalBoolTrue      OptionalBool = 1
	OptionalBoolUndefined OptionalBool = 2
)

type popErrorScopeStatus C.WGPUPopErrorScopeStatus

const (
	popErrorScopeStatusSuccess           popErrorScopeStatus = 1
	popErrorScopeStatusCallbackCancelled popErrorScopeStatus = 2
	popErrorScopeStatusError             popErrorScopeStatus = 3
)

type PowerPreference C.WGPUPowerPreference

const (
	PowerPreferenceUndefined       PowerPreference = 0
	PowerPreferenceLowPower        PowerPreference = 1
	PowerPreferenceHighPerformance PowerPreference = 2
)

type PredefinedColorSpace C.WGPUPredefinedColorSpace

const (
	PredefinedColorSpaceSRGB            PredefinedColorSpace = 1
	PredefinedColorSpaceDisplayP3       PredefinedColorSpace = 2
	PredefinedColorSpaceSRGBLinear      PredefinedColorSpace = 3
	PredefinedColorSpaceDisplayP3Linear PredefinedColorSpace = 4
)

type PresentMode C.WGPUPresentMode

const (
	PresentModeUndefined   PresentMode = 0
	PresentModeFifo        PresentMode = 1
	PresentModeFifoRelaxed PresentMode = 2
	PresentModeImmediate   PresentMode = 3
	PresentModeMailbox     PresentMode = 4
)

func (p PresentMode) String() string {
	switch p {
	case PresentModeFifo:
		return "Fifo"
	case PresentModeFifoRelaxed:
		return "Fifo Relaxed"
	case PresentModeImmediate:
		return "Immediate"
	case PresentModeMailbox:
		return "Mailbox"
	default:
		return "Undefined"
	}
}

type PrimitiveTopology C.WGPUPrimitiveTopology

const (
	PrimitiveTopologyUndefined     PrimitiveTopology = 0
	PrimitiveTopologyPointList     PrimitiveTopology = 1
	PrimitiveTopologyLineList      PrimitiveTopology = 2
	PrimitiveTopologyLineStrip     PrimitiveTopology = 3
	PrimitiveTopologyTriangleList  PrimitiveTopology = 4
	PrimitiveTopologyTriangleStrip PrimitiveTopology = 5
)

type QueryType C.WGPUQueryType

const (
	QueryTypeOcclusion QueryType = 1
	QueryTypeTimestamp QueryType = 2
)

type QueueWorkDoneStatus C.WGPUQueueWorkDoneStatus

const (
	QueueWorkDoneStatusSuccess           QueueWorkDoneStatus = 1
	QueueWorkDoneStatusCallbackCancelled QueueWorkDoneStatus = 2
	QueueWorkDoneStatusError             QueueWorkDoneStatus = 3
)

type requestAdapterStatus C.WGPURequestAdapterStatus

const (
	requestAdapterStatusSuccess           requestAdapterStatus = 1
	requestAdapterStatusCallbackCancelled requestAdapterStatus = 2
	requestAdapterStatusUnavailable       requestAdapterStatus = 3
	requestAdapterStatusError             requestAdapterStatus = 4
)

type requestDeviceStatus C.WGPURequestDeviceStatus

const (
	requestDeviceStatusSuccess           requestDeviceStatus = 1
	requestDeviceStatusCallbackCancelled requestDeviceStatus = 2
	requestDeviceStatusError             requestDeviceStatus = 3
)

type SType C.WGPUSType

const (
	STypeShaderSourceSPIRV                 SType = 1
	STypeShaderSourceWGSL                  SType = 2
	STypeRenderPassMaxDrawCount            SType = 3
	STypeSurfaceSourceMetalLayer           SType = 4
	STypeSurfaceSourceWindowsHWND          SType = 5
	STypeSurfaceSourceXlibWindow           SType = 6
	STypeSurfaceSourceWaylandSurface       SType = 7
	STypeSurfaceSourceAndroidNativeWindow  SType = 8
	STypeSurfaceSourceXCBWindow            SType = 9
	STypeSurfaceColorManagement            SType = 10
	STypeRequestAdapterWebXROptions        SType = 11
	STypeTextureComponentSwizzleDescriptor SType = 12
	STypeExternalTextureBindingLayout      SType = 13
	STypeExternalTextureBindingEntry       SType = 14
	STypeCompatibilityModeLimits           SType = 15
	STypeTextureBindingViewDimension       SType = 16
)

type SamplerBindingType C.WGPUSamplerBindingType

const (
	SamplerBindingTypeBindingNotUsed SamplerBindingType = 0
	SamplerBindingTypeUndefined      SamplerBindingType = 1
	SamplerBindingTypeFiltering      SamplerBindingType = 2
	SamplerBindingTypeNonFiltering   SamplerBindingType = 3
	SamplerBindingTypeComparison     SamplerBindingType = 4
)

type ShaderStage C.WGPUShaderStage

const (
	ShaderStageNone     ShaderStage = 0
	ShaderStageVertex   ShaderStage = 1
	ShaderStageFragment ShaderStage = 2
	ShaderStageCompute  ShaderStage = 4
)

type statusCode C.WGPUStatus

const (
	statusCodeSuccess statusCode = 1
	statusCodeError   statusCode = 2
)

type StencilOperation C.WGPUStencilOperation

const (
	StencilOperationUndefined      StencilOperation = 0
	StencilOperationKeep           StencilOperation = 1
	StencilOperationZero           StencilOperation = 2
	StencilOperationReplace        StencilOperation = 3
	StencilOperationInvert         StencilOperation = 4
	StencilOperationIncrementClamp StencilOperation = 5
	StencilOperationDecrementClamp StencilOperation = 6
	StencilOperationIncrementWrap  StencilOperation = 7
	StencilOperationDecrementWrap  StencilOperation = 8
)

type StorageTextureAccess C.WGPUStorageTextureAccess

const (
	StorageTextureAccessBindingNotUsed StorageTextureAccess = 0
	StorageTextureAccessUndefined      StorageTextureAccess = 1
	StorageTextureAccessWriteOnly      StorageTextureAccess = 2
	StorageTextureAccessReadOnly       StorageTextureAccess = 3
	StorageTextureAccessReadWrite      StorageTextureAccess = 4
)

type StoreOp C.WGPUStoreOp

const (
	StoreOpUndefined StoreOp = 0
	StoreOpStore     StoreOp = 1
	StoreOpDiscard   StoreOp = 2
)

type SurfaceGetCurrentTextureStatus C.WGPUSurfaceGetCurrentTextureStatus

const (
	SurfaceGetCurrentTextureStatusSuccessOptimal    SurfaceGetCurrentTextureStatus = 1
	SurfaceGetCurrentTextureStatusSuccessSuboptimal SurfaceGetCurrentTextureStatus = 2
	SurfaceGetCurrentTextureStatusTimeout           SurfaceGetCurrentTextureStatus = 3
	SurfaceGetCurrentTextureStatusOutdated          SurfaceGetCurrentTextureStatus = 4
	SurfaceGetCurrentTextureStatusLost              SurfaceGetCurrentTextureStatus = 5
	SurfaceGetCurrentTextureStatusError             SurfaceGetCurrentTextureStatus = 6
)

type TextureAspect C.WGPUTextureAspect

const (
	TextureAspectUndefined   TextureAspect = 0
	TextureAspectAll         TextureAspect = 1
	TextureAspectStencilOnly TextureAspect = 2
	TextureAspectDepthOnly   TextureAspect = 3
	TextureAspectPlane0Only  TextureAspect = 0
	TextureAspectPlane1Only  TextureAspect = 1
	TextureAspectPlane2Only  TextureAspect = 2
)

type TextureDimension C.WGPUTextureDimension

const (
	TextureDimensionUndefined TextureDimension = 0
	TextureDimension1D        TextureDimension = 1
	TextureDimension2D        TextureDimension = 2
	TextureDimension3D        TextureDimension = 3
)

type TextureFormat C.WGPUTextureFormat

const (
	TextureFormatUndefined            TextureFormat = 0
	TextureFormatR8Unorm              TextureFormat = 1
	TextureFormatR8Snorm              TextureFormat = 2
	TextureFormatR8Uint               TextureFormat = 3
	TextureFormatR8Sint               TextureFormat = 4
	TextureFormatR16Unorm             TextureFormat = 5
	TextureFormatR16Snorm             TextureFormat = 6
	TextureFormatR16Uint              TextureFormat = 7
	TextureFormatR16Sint              TextureFormat = 8
	TextureFormatR16Float             TextureFormat = 9
	TextureFormatRG8Unorm             TextureFormat = 10
	TextureFormatRG8Snorm             TextureFormat = 11
	TextureFormatRG8Uint              TextureFormat = 12
	TextureFormatRG8Sint              TextureFormat = 13
	TextureFormatR32Float             TextureFormat = 14
	TextureFormatR32Uint              TextureFormat = 15
	TextureFormatR32Sint              TextureFormat = 16
	TextureFormatRG16Unorm            TextureFormat = 17
	TextureFormatRG16Snorm            TextureFormat = 18
	TextureFormatRG16Uint             TextureFormat = 19
	TextureFormatRG16Sint             TextureFormat = 20
	TextureFormatRG16Float            TextureFormat = 21
	TextureFormatRGBA8Unorm           TextureFormat = 22
	TextureFormatRGBA8UnormSRGB       TextureFormat = 23
	TextureFormatRGBA8Snorm           TextureFormat = 24
	TextureFormatRGBA8Uint            TextureFormat = 25
	TextureFormatRGBA8Sint            TextureFormat = 26
	TextureFormatBGRA8Unorm           TextureFormat = 27
	TextureFormatBGRA8UnormSRGB       TextureFormat = 28
	TextureFormatRGB10A2Uint          TextureFormat = 29
	TextureFormatRGB10A2Unorm         TextureFormat = 30
	TextureFormatRG11B10Ufloat        TextureFormat = 31
	TextureFormatRGB9E5Ufloat         TextureFormat = 32
	TextureFormatRG32Float            TextureFormat = 33
	TextureFormatRG32Uint             TextureFormat = 34
	TextureFormatRG32Sint             TextureFormat = 35
	TextureFormatRGBA16Unorm          TextureFormat = 36
	TextureFormatRGBA16Snorm          TextureFormat = 37
	TextureFormatRGBA16Uint           TextureFormat = 38
	TextureFormatRGBA16Sint           TextureFormat = 39
	TextureFormatRGBA16Float          TextureFormat = 40
	TextureFormatRGBA32Float          TextureFormat = 41
	TextureFormatRGBA32Uint           TextureFormat = 42
	TextureFormatRGBA32Sint           TextureFormat = 43
	TextureFormatStencil8             TextureFormat = 44
	TextureFormatDepth16Unorm         TextureFormat = 45
	TextureFormatDepth24Plus          TextureFormat = 46
	TextureFormatDepth24PlusStencil8  TextureFormat = 47
	TextureFormatDepth32Float         TextureFormat = 48
	TextureFormatDepth32FloatStencil8 TextureFormat = 49
	TextureFormatBC1RGBAUnorm         TextureFormat = 50
	TextureFormatBC1RGBAUnormSRGB     TextureFormat = 51
	TextureFormatBC2RGBAUnorm         TextureFormat = 52
	TextureFormatBC2RGBAUnormSRGB     TextureFormat = 53
	TextureFormatBC3RGBAUnorm         TextureFormat = 54
	TextureFormatBC3RGBAUnormSRGB     TextureFormat = 55
	TextureFormatBC4RUnorm            TextureFormat = 56
	TextureFormatBC4RSnorm            TextureFormat = 57
	TextureFormatBC5RGUnorm           TextureFormat = 58
	TextureFormatBC5RGSnorm           TextureFormat = 59
	TextureFormatBC6HRGBUfloat        TextureFormat = 60
	TextureFormatBC6HRGBFloat         TextureFormat = 61
	TextureFormatBC7RGBAUnorm         TextureFormat = 62
	TextureFormatBC7RGBAUnormSRGB     TextureFormat = 63
	TextureFormatETC2RGB8Unorm        TextureFormat = 64
	TextureFormatETC2RGB8UnormSRGB    TextureFormat = 65
	TextureFormatETC2RGB8A1Unorm      TextureFormat = 66
	TextureFormatETC2RGB8A1UnormSRGB  TextureFormat = 67
	TextureFormatETC2RGBA8Unorm       TextureFormat = 68
	TextureFormatETC2RGBA8UnormSRGB   TextureFormat = 69
	TextureFormatEACR11Unorm          TextureFormat = 70
	TextureFormatEACR11Snorm          TextureFormat = 71
	TextureFormatEACRG11Unorm         TextureFormat = 72
	TextureFormatEACRG11Snorm         TextureFormat = 73
	TextureFormatASTC4x4Unorm         TextureFormat = 74
	TextureFormatASTC4x4UnormSRGB     TextureFormat = 75
	TextureFormatASTC5x4Unorm         TextureFormat = 76
	TextureFormatASTC5x4UnormSRGB     TextureFormat = 77
	TextureFormatASTC5x5Unorm         TextureFormat = 78
	TextureFormatASTC5x5UnormSRGB     TextureFormat = 79
	TextureFormatASTC6x5Unorm         TextureFormat = 80
	TextureFormatASTC6x5UnormSRGB     TextureFormat = 81
	TextureFormatASTC6x6Unorm         TextureFormat = 82
	TextureFormatASTC6x6UnormSRGB     TextureFormat = 83
	TextureFormatASTC8x5Unorm         TextureFormat = 84
	TextureFormatASTC8x5UnormSRGB     TextureFormat = 85
	TextureFormatASTC8x6Unorm         TextureFormat = 86
	TextureFormatASTC8x6UnormSRGB     TextureFormat = 87
	TextureFormatASTC8x8Unorm         TextureFormat = 88
	TextureFormatASTC8x8UnormSRGB     TextureFormat = 89
	TextureFormatASTC10x5Unorm        TextureFormat = 90
	TextureFormatASTC10x5UnormSRGB    TextureFormat = 91
	TextureFormatASTC10x6Unorm        TextureFormat = 92
	TextureFormatASTC10x6UnormSRGB    TextureFormat = 93
	TextureFormatASTC10x8Unorm        TextureFormat = 94
	TextureFormatASTC10x8UnormSRGB    TextureFormat = 95
	TextureFormatASTC10x10Unorm       TextureFormat = 96
	TextureFormatASTC10x10UnormSRGB   TextureFormat = 97
	TextureFormatASTC12x10Unorm       TextureFormat = 98
	TextureFormatASTC12x10UnormSRGB   TextureFormat = 99
	TextureFormatASTC12x12Unorm       TextureFormat = 100
	TextureFormatASTC12x12UnormSRGB   TextureFormat = 101
)

func (t TextureFormat) String() string {
	switch t {
	case TextureFormatR8Unorm:
		return "R8Unorm"
	case TextureFormatR8Snorm:
		return "R8Snorm"
	case TextureFormatR8Uint:
		return "R8Uint"
	case TextureFormatR8Sint:
		return "R8Sint"
	case TextureFormatR16Unorm:
		return "R16Unorm"
	case TextureFormatR16Snorm:
		return "R16Snorm"
	case TextureFormatR16Uint:
		return "R16Uint"
	case TextureFormatR16Sint:
		return "R16Sint"
	case TextureFormatR16Float:
		return "R16Float"
	case TextureFormatRG8Unorm:
		return "RG8Unorm"
	case TextureFormatRG8Snorm:
		return "RG8Snorm"
	case TextureFormatRG8Uint:
		return "RG8Uint"
	case TextureFormatRG8Sint:
		return "RG8Sint"
	case TextureFormatR32Float:
		return "R32Float"
	case TextureFormatR32Uint:
		return "R32Uint"
	case TextureFormatR32Sint:
		return "R32Sint"
	case TextureFormatRG16Unorm:
		return "RG16Unorm"
	case TextureFormatRG16Snorm:
		return "RG16Snorm"
	case TextureFormatRG16Uint:
		return "RG16Uint"
	case TextureFormatRG16Sint:
		return "RG16Sint"
	case TextureFormatRG16Float:
		return "RG16Float"
	case TextureFormatRGBA8Unorm:
		return "RGBA8Unorm"
	case TextureFormatRGBA8UnormSRGB:
		return "RGBA8UnormSRGB"
	case TextureFormatRGBA8Snorm:
		return "RGBA8Snorm"
	case TextureFormatRGBA8Uint:
		return "RGBA8Uint"
	case TextureFormatRGBA8Sint:
		return "RGBA8Sint"
	case TextureFormatBGRA8Unorm:
		return "BGRA8Unorm"
	case TextureFormatBGRA8UnormSRGB:
		return "BGRA8UnormSRGB"
	case TextureFormatRGB10A2Uint:
		return "RGB10A2Uint"
	case TextureFormatRGB10A2Unorm:
		return "RGB10A2Unorm"
	case TextureFormatRG11B10Ufloat:
		return "RG11B10Ufloat"
	case TextureFormatRGB9E5Ufloat:
		return "RGB9E5Ufloat"
	case TextureFormatRG32Float:
		return "RG32Float"
	case TextureFormatRG32Uint:
		return "RG32Uint"
	case TextureFormatRG32Sint:
		return "RG32Sint"
	case TextureFormatRGBA16Unorm:
		return "RGBA16Unorm"
	case TextureFormatRGBA16Snorm:
		return "RGBA16Snorm"
	case TextureFormatRGBA16Uint:
		return "RGBA16Uint"
	case TextureFormatRGBA16Sint:
		return "RGBA16Sint"
	case TextureFormatRGBA16Float:
		return "RGBA16Float"
	case TextureFormatRGBA32Float:
		return "RGBA32Float"
	case TextureFormatRGBA32Uint:
		return "RGBA32Uint"
	case TextureFormatRGBA32Sint:
		return "RGBA32Sint"
	case TextureFormatStencil8:
		return "Stencil8"
	case TextureFormatDepth16Unorm:
		return "Depth16Unorm"
	case TextureFormatDepth24Plus:
		return "Depth24Plus"
	case TextureFormatDepth24PlusStencil8:
		return "Depth24PlusStencil8"
	case TextureFormatDepth32Float:
		return "Depth32Float"
	case TextureFormatDepth32FloatStencil8:
		return "Depth32FloatStencil8"
	case TextureFormatBC1RGBAUnorm:
		return "BC1RGBAUnorm"
	case TextureFormatBC1RGBAUnormSRGB:
		return "BC1RGBAUnormSRGB"
	case TextureFormatBC2RGBAUnorm:
		return "BC2RGBAUnorm"
	case TextureFormatBC2RGBAUnormSRGB:
		return "BC2RGBAUnormSRGB"
	case TextureFormatBC3RGBAUnorm:
		return "BC3RGBAUnorm"
	case TextureFormatBC3RGBAUnormSRGB:
		return "BC3RGBAUnormSRGB"
	case TextureFormatBC4RUnorm:
		return "BC4RUnorm"
	case TextureFormatBC4RSnorm:
		return "BC4RSnorm"
	case TextureFormatBC5RGUnorm:
		return "BC5RGUnorm"
	case TextureFormatBC5RGSnorm:
		return "BC5RGSnorm"
	case TextureFormatBC6HRGBUfloat:
		return "BC6HRGBUfloat"
	case TextureFormatBC6HRGBFloat:
		return "BC6HRGBFloat"
	case TextureFormatBC7RGBAUnorm:
		return "BC7RGBAUnorm"
	case TextureFormatBC7RGBAUnormSRGB:
		return "BC7RGBAUnormSRGB"
	case TextureFormatETC2RGB8Unorm:
		return "ETC2RGB8Unorm"
	case TextureFormatETC2RGB8UnormSRGB:
		return "ETC2RGB8UnormSRGB"
	case TextureFormatETC2RGB8A1Unorm:
		return "ETC2RGB8A1Unorm"
	case TextureFormatETC2RGB8A1UnormSRGB:
		return "ETC2RGB8A1UnormSRGB"
	case TextureFormatETC2RGBA8Unorm:
		return "ETC2RGBA8Unorm"
	case TextureFormatETC2RGBA8UnormSRGB:
		return "ETC2RGBA8UnormSRGB"
	case TextureFormatEACR11Unorm:
		return "EACR11Unorm"
	case TextureFormatEACR11Snorm:
		return "EACR11Snorm"
	case TextureFormatEACRG11Unorm:
		return "EACRG11Unorm"
	case TextureFormatEACRG11Snorm:
		return "EACRG11Snorm"
	case TextureFormatASTC4x4Unorm:
		return "ASTC4x4Unorm"
	case TextureFormatASTC4x4UnormSRGB:
		return "ASTC4x4UnormSRGB"
	case TextureFormatASTC5x4Unorm:
		return "ASTC5x4Unorm"
	case TextureFormatASTC5x4UnormSRGB:
		return "ASTC5x4UnormSRGB"
	case TextureFormatASTC5x5Unorm:
		return "ASTC5x5Unorm"
	case TextureFormatASTC5x5UnormSRGB:
		return "ASTC5x5UnormSRGB"
	case TextureFormatASTC6x5Unorm:
		return "ASTC6x5Unorm"
	case TextureFormatASTC6x5UnormSRGB:
		return "ASTC6x5UnormSRGB"
	case TextureFormatASTC6x6Unorm:
		return "ASTC6x6Unorm"
	case TextureFormatASTC6x6UnormSRGB:
		return "ASTC6x6UnormSRGB"
	case TextureFormatASTC8x5Unorm:
		return "ASTC8x5Unorm"
	case TextureFormatASTC8x5UnormSRGB:
		return "ASTC8x5UnormSRGB"
	case TextureFormatASTC8x6Unorm:
		return "ASTC8x6Unorm"
	case TextureFormatASTC8x6UnormSRGB:
		return "ASTC8x6UnormSRGB"
	case TextureFormatASTC8x8Unorm:
		return "ASTC8x8Unorm"
	case TextureFormatASTC8x8UnormSRGB:
		return "ASTC8x8UnormSRGB"
	case TextureFormatASTC10x5Unorm:
		return "ASTC10x5Unorm"
	case TextureFormatASTC10x5UnormSRGB:
		return "ASTC10x5UnormSRGB"
	case TextureFormatASTC10x6Unorm:
		return "ASTC10x6Unorm"
	case TextureFormatASTC10x6UnormSRGB:
		return "ASTC10x6UnormSRGB"
	case TextureFormatASTC10x8Unorm:
		return "ASTC10x8Unorm"
	case TextureFormatASTC10x8UnormSRGB:
		return "ASTC10x8UnormSRGB"
	case TextureFormatASTC10x10Unorm:
		return "ASTC10x10Unorm"
	case TextureFormatASTC10x10UnormSRGB:
		return "ASTC10x10UnormSRGB"
	case TextureFormatASTC12x10Unorm:
		return "ASTC12x10Unorm"
	case TextureFormatASTC12x10UnormSRGB:
		return "ASTC12x10UnormSRGB"
	case TextureFormatASTC12x12Unorm:
		return "ASTC12x12Unorm"
	case TextureFormatASTC12x12UnormSRGB:
		return "ASTC12x12UnormSRGB"
	default:
		return "Undefined"
	}
}

type TextureSampleType C.WGPUTextureSampleType

const (
	TextureSampleTypeBindingNotUsed    TextureSampleType = 0
	TextureSampleTypeUndefined         TextureSampleType = 1
	TextureSampleTypeFloat             TextureSampleType = 2
	TextureSampleTypeUnfilterableFloat TextureSampleType = 3
	TextureSampleTypeDepth             TextureSampleType = 4
	TextureSampleTypeSint              TextureSampleType = 5
	TextureSampleTypeUint              TextureSampleType = 6
)

type TextureUsage C.WGPUTextureUsage

const (
	TextureUsageNone                TextureUsage = 0
	TextureUsageCopySrc             TextureUsage = 1
	TextureUsageCopyDst             TextureUsage = 2
	TextureUsageTextureBinding      TextureUsage = 4
	TextureUsageStorageBinding      TextureUsage = 8
	TextureUsageRenderAttachment    TextureUsage = 16
	TextureUsageTransientAttachment TextureUsage = 32
	TextureUsageStorageAttachment   TextureUsage = 64
)

type TextureViewDimension C.WGPUTextureViewDimension

const (
	TextureViewDimensionUndefined TextureViewDimension = 0
	TextureViewDimension1D        TextureViewDimension = 1
	TextureViewDimension2D        TextureViewDimension = 2
	TextureViewDimension2DArray   TextureViewDimension = 3
	TextureViewDimensionCube      TextureViewDimension = 4
	TextureViewDimensionCubeArray TextureViewDimension = 5
	TextureViewDimension3D        TextureViewDimension = 6
)

type ToneMappingMode C.WGPUToneMappingMode

const (
	ToneMappingModeStandard ToneMappingMode = 1
	ToneMappingModeExtended ToneMappingMode = 2
)

type VertexFormat C.WGPUVertexFormat

const (
	VertexFormatUint8           VertexFormat = 1
	VertexFormatUint8x2         VertexFormat = 2
	VertexFormatUint8x4         VertexFormat = 3
	VertexFormatSint8           VertexFormat = 4
	VertexFormatSint8x2         VertexFormat = 5
	VertexFormatSint8x4         VertexFormat = 6
	VertexFormatUnorm8          VertexFormat = 7
	VertexFormatUnorm8x2        VertexFormat = 8
	VertexFormatUnorm8x4        VertexFormat = 9
	VertexFormatSnorm8          VertexFormat = 10
	VertexFormatSnorm8x2        VertexFormat = 11
	VertexFormatSnorm8x4        VertexFormat = 12
	VertexFormatUint16          VertexFormat = 13
	VertexFormatUint16x2        VertexFormat = 14
	VertexFormatUint16x4        VertexFormat = 15
	VertexFormatSint16          VertexFormat = 16
	VertexFormatSint16x2        VertexFormat = 17
	VertexFormatSint16x4        VertexFormat = 18
	VertexFormatUnorm16         VertexFormat = 19
	VertexFormatUnorm16x2       VertexFormat = 20
	VertexFormatUnorm16x4       VertexFormat = 21
	VertexFormatSnorm16         VertexFormat = 22
	VertexFormatSnorm16x2       VertexFormat = 23
	VertexFormatSnorm16x4       VertexFormat = 24
	VertexFormatFloat16         VertexFormat = 25
	VertexFormatFloat16x2       VertexFormat = 26
	VertexFormatFloat16x4       VertexFormat = 27
	VertexFormatFloat32         VertexFormat = 28
	VertexFormatFloat32x2       VertexFormat = 29
	VertexFormatFloat32x3       VertexFormat = 30
	VertexFormatFloat32x4       VertexFormat = 31
	VertexFormatUint32          VertexFormat = 32
	VertexFormatUint32x2        VertexFormat = 33
	VertexFormatUint32x3        VertexFormat = 34
	VertexFormatUint32x4        VertexFormat = 35
	VertexFormatSint32          VertexFormat = 36
	VertexFormatSint32x2        VertexFormat = 37
	VertexFormatSint32x3        VertexFormat = 38
	VertexFormatSint32x4        VertexFormat = 39
	VertexFormatUnorm10_10_10_2 VertexFormat = 40
	VertexFormatUnorm8x4BGRA    VertexFormat = 41
)

type VertexStepMode C.WGPUVertexStepMode

const (
	VertexStepModeUndefined VertexStepMode = 0
	VertexStepModeVertex    VertexStepMode = 1
	VertexStepModeInstance  VertexStepMode = 2
)

type waitStatus C.WGPUWaitStatus

const (
	waitStatusSuccess  waitStatus = 1
	waitStatusTimedOut waitStatus = 2
	waitStatusError    waitStatus = 3
)

const (
	ArrayLayerCountUndefined        = 0xffffffff
	CopyStrideUndefined             = 0xffffffff
	LimitU32Undefined        uint32 = 0xffffffff
	LimitU64Undefined        uint64 = 0xffffffffffffffff
	MipLevelCountUndefined          = 0xffffffff
	WholeMapSize                    = ^uint(0)
	WholeSize                       = 0xffffffffffffffff
)
