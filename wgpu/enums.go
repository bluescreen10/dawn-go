// CODE GENERATED. DO NOT EDIT
//
//go:generate go run ./cmd/wrapper/.
package wgpu

type WGSLLanguageFeatureName int

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

type AdapterType int

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

type AddressMode int

const (
	AddressModeUndefined    AddressMode = 0
	AddressModeClampToEdge              = 1
	AddressModeRepeat                   = 2
	AddressModeMirrorRepeat             = 3
)

type BackendType int

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

type BlendFactor int

const (
	BlendFactorUndefined         BlendFactor = 0
	BlendFactorZero                          = 1
	BlendFactorOne                           = 2
	BlendFactorSrc                           = 3
	BlendFactorOneMinusSrc                   = 4
	BlendFactorSrcAlpha                      = 5
	BlendFactorOneMinusSrcAlpha              = 6
	BlendFactorDst                           = 7
	BlendFactorOneMinusDst                   = 8
	BlendFactorDstAlpha                      = 9
	BlendFactorOneMinusDstAlpha              = 10
	BlendFactorSrcAlphaSaturated             = 11
	BlendFactorConstant                      = 12
	BlendFactorOneMinusConstant              = 13
	BlendFactorSrc1                          = 14
	BlendFactorOneMinusSrc1                  = 15
	BlendFactorSrc1Alpha                     = 16
	BlendFactorOneMinusSrc1Alpha             = 17
)

type BlendOperation int

const (
	BlendOperationUndefined       BlendOperation = 0
	BlendOperationAdd                            = 1
	BlendOperationSubtract                       = 2
	BlendOperationReverseSubtract                = 3
	BlendOperationMin                            = 4
	BlendOperationMax                            = 5
)

type BufferBindingType int

const (
	BufferBindingTypeBindingNotUsed  BufferBindingType = 0
	BufferBindingTypeUndefined                         = 1
	BufferBindingTypeUniform                           = 2
	BufferBindingTypeStorage                           = 3
	BufferBindingTypeReadOnlyStorage                   = 4
)

type BufferMapState int

const (
	BufferMapStateUnmapped BufferMapState = 1
	BufferMapStatePending                 = 2
	BufferMapStateMapped                  = 3
)

type BufferUsage int

const (
	BufferUsageNone         BufferUsage = 0
	BufferUsageMapRead                  = 1
	BufferUsageMapWrite                 = 2
	BufferUsageCopySrc                  = 4
	BufferUsageCopyDst                  = 8
	BufferUsageIndex                    = 16
	BufferUsageVertex                   = 32
	BufferUsageUniform                  = 64
	BufferUsageStorage                  = 128
	BufferUsageIndirect                 = 256
	BufferUsageQueryResolve             = 512
	BufferUsageTexelBuffer              = 1024
)

type callbackMode int

const (
	callbackModeWaitAnyOnly        callbackMode = 1
	callbackModeAllowProcessEvents              = 2
	callbackModeAllowSpontaneous                = 3
)

type ColorWriteMask int

const (
	ColorWriteMaskNone  ColorWriteMask = 0
	ColorWriteMaskRed                  = 1
	ColorWriteMaskGreen                = 2
	ColorWriteMaskBlue                 = 4
	ColorWriteMaskAlpha                = 8
	ColorWriteMaskAll                  = 15
)

type CompareFunction int

const (
	CompareFunctionUndefined    CompareFunction = 0
	CompareFunctionNever                        = 1
	CompareFunctionLess                         = 2
	CompareFunctionEqual                        = 3
	CompareFunctionLessEqual                    = 4
	CompareFunctionGreater                      = 5
	CompareFunctionNotEqual                     = 6
	CompareFunctionGreaterEqual                 = 7
	CompareFunctionAlways                       = 8
)

type compilationInfoRequestStatus int

const (
	compilationInfoRequestStatusSuccess           compilationInfoRequestStatus = 1
	compilationInfoRequestStatusCallbackCancelled                              = 2
)

type CompilationMessageType int

const (
	CompilationMessageTypeError   CompilationMessageType = 1
	CompilationMessageTypeWarning                        = 2
	CompilationMessageTypeInfo                           = 3
)

type ComponentSwizzle int

const (
	ComponentSwizzleUndefined ComponentSwizzle = 0
	ComponentSwizzleZero                       = 1
	ComponentSwizzleOne                        = 2
	ComponentSwizzleR                          = 3
	ComponentSwizzleG                          = 4
	ComponentSwizzleB                          = 5
	ComponentSwizzleA                          = 6
)

type CompositeAlphaMode int

const (
	CompositeAlphaModeAuto            CompositeAlphaMode = 0
	CompositeAlphaModeOpaque                             = 1
	CompositeAlphaModePremultiplied                      = 2
	CompositeAlphaModeUnpremultiplied                    = 3
	CompositeAlphaModeInherit                            = 4
)

type CreatePipelineAsyncStatus int

const (
	CreatePipelineAsyncStatusSuccess           CreatePipelineAsyncStatus = 1
	CreatePipelineAsyncStatusCallbackCancelled                           = 2
	CreatePipelineAsyncStatusValidationError                             = 3
	CreatePipelineAsyncStatusInternalError                               = 4
)

type CullMode int

const (
	CullModeUndefined CullMode = 0
	CullModeNone               = 1
	CullModeFront              = 2
	CullModeBack               = 3
)

type DeviceLostReason int

const (
	DeviceLostReasonUnknown           DeviceLostReason = 1
	DeviceLostReasonDestroyed                          = 2
	DeviceLostReasonCallbackCancelled                  = 3
	DeviceLostReasonFailedCreation                     = 4
)

type ErrorFilter int

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

type ErrorType int

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

type FeatureLevel int

const (
	FeatureLevelUndefined     FeatureLevel = 0
	FeatureLevelCompatibility              = 1
	FeatureLevelCore                       = 2
)

type FeatureName int

const (
	FeatureNameCoreFeaturesAndLimits                                FeatureName = 1
	FeatureNameDepthClipControl                                                 = 2
	FeatureNameDepth32FloatStencil8                                             = 3
	FeatureNameTextureCompressionBC                                             = 4
	FeatureNameTextureCompressionBCSliced3D                                     = 5
	FeatureNameTextureCompressionETC2                                           = 6
	FeatureNameTextureCompressionASTC                                           = 7
	FeatureNameTextureCompressionASTCSliced3D                                   = 8
	FeatureNameTimestampQuery                                                   = 9
	FeatureNameIndirectFirstInstance                                            = 10
	FeatureNameShaderF16                                                        = 11
	FeatureNameRG11B10UfloatRenderable                                          = 12
	FeatureNameBGRA8UnormStorage                                                = 13
	FeatureNameFloat32Filterable                                                = 14
	FeatureNameFloat32Blendable                                                 = 15
	FeatureNameClipDistances                                                    = 16
	FeatureNameDualSourceBlending                                               = 17
	FeatureNameSubgroups                                                        = 18
	FeatureNameTextureFormatsTier1                                              = 19
	FeatureNameTextureFormatsTier2                                              = 20
	FeatureNamePrimitiveIndex                                                   = 21
	FeatureNameTextureComponentSwizzle                                          = 22
	FeatureNameDawnInternalUsages                                               = 0
	FeatureNameDawnMultiPlanarFormats                                           = 1
	FeatureNameDawnNative                                                       = 2
	FeatureNameChromiumExperimentalTimestampQueryInsidePasses                   = 3
	FeatureNameImplicitDeviceSynchronization                                    = 4
	FeatureNameTransientAttachments                                             = 6
	FeatureNameMSAARenderToSingleSampled                                        = 7
	FeatureNameD3D11MultithreadProtected                                        = 8
	FeatureNameANGLETextureSharing                                              = 9
	FeatureNamePixelLocalStorageCoherent                                        = 10
	FeatureNamePixelLocalStorageNonCoherent                                     = 11
	FeatureNameUnorm16TextureFormats                                            = 12
	FeatureNameMultiPlanarFormatExtendedUsages                                  = 13
	FeatureNameMultiPlanarFormatP010                                            = 14
	FeatureNameHostMappedPointer                                                = 15
	FeatureNameMultiPlanarRenderTargets                                         = 16
	FeatureNameMultiPlanarFormatNv12a                                           = 17
	FeatureNameFramebufferFetch                                                 = 18
	FeatureNameBufferMapExtendedUsages                                          = 19
	FeatureNameAdapterPropertiesMemoryHeaps                                     = 20
	FeatureNameAdapterPropertiesD3D                                             = 21
	FeatureNameAdapterPropertiesVk                                              = 22
	FeatureNameDawnFormatCapabilities                                           = 23
	FeatureNameDawnDrmFormatCapabilities                                        = 24
	FeatureNameMultiPlanarFormatNv16                                            = 25
	FeatureNameMultiPlanarFormatNv24                                            = 26
	FeatureNameMultiPlanarFormatP210                                            = 27
	FeatureNameMultiPlanarFormatP410                                            = 28
	FeatureNameSharedTextureMemoryVkDedicatedAllocation                         = 29
	FeatureNameSharedTextureMemoryAHardwareBuffer                               = 30
	FeatureNameSharedTextureMemoryDmaBuf                                        = 31
	FeatureNameSharedTextureMemoryOpaqueFD                                      = 32
	FeatureNameSharedTextureMemoryZirconHandle                                  = 33
	FeatureNameSharedTextureMemoryDXGISharedHandle                              = 34
	FeatureNameSharedTextureMemoryD3D11Texture2D                                = 35
	FeatureNameSharedTextureMemoryIOSurface                                     = 36
	FeatureNameSharedTextureMemoryEGLImage                                      = 37
	FeatureNameSharedFenceVkSemaphoreOpaqueFD                                   = 38
	FeatureNameSharedFenceSyncFD                                                = 39
	FeatureNameSharedFenceVkSemaphoreZirconHandle                               = 40
	FeatureNameSharedFenceDXGISharedHandle                                      = 41
	FeatureNameSharedFenceMTLSharedEvent                                        = 42
	FeatureNameSharedBufferMemoryD3D12Resource                                  = 43
	FeatureNameStaticSamplers                                                   = 44
	FeatureNameYCbCrVulkanSamplers                                              = 45
	FeatureNameShaderModuleCompilationOptions                                   = 46
	FeatureNameDawnLoadResolveTexture                                           = 47
	FeatureNameDawnPartialLoadResolveTexture                                    = 48
	FeatureNameMultiDrawIndirect                                                = 49
	FeatureNameDawnTexelCopyBufferRowAlignment                                  = 50
	FeatureNameFlexibleTextureViews                                             = 51
	FeatureNameChromiumExperimentalSubgroupMatrix                               = 52
	FeatureNameSharedFenceEGLSync                                               = 53
	FeatureNameDawnDeviceAllocatorControl                                       = 54
	FeatureNameAdapterPropertiesWGPU                                            = 55
	FeatureNameSharedBufferMemoryD3D12SharedMemoryFileMappingHandle             = 56
	FeatureNameSharedTextureMemoryD3D12Resource                                 = 57
	FeatureNameChromiumExperimentalSamplingResourceTable                        = 58
	FeatureNameChromiumExperimentalSubgroupSizeControl                          = 59
	FeatureNameAtomicVec2uMinMax                                                = 60
	FeatureNameUnorm16FormatsForExternalTexture                                 = 61
	FeatureNameOpaqueYCbCrAndroidForExternalTexture                             = 62
	FeatureNameUnorm16Filterable                                                = 63
	FeatureNameRenderPassRenderArea                                             = 64
	FeatureNameDawnNativeSpontaneousQueueEvents                                 = 65
	FeatureNameAdapterPropertiesDrm                                             = 66
)

type FilterMode int

const (
	FilterModeUndefined FilterMode = 0
	FilterModeNearest              = 1
	FilterModeLinear               = 2
)

type FrontFace int

const (
	FrontFaceUndefined FrontFace = 0
	FrontFaceCCW                 = 1
	FrontFaceCW                  = 2
)

type IndexFormat int

const (
	IndexFormatUndefined IndexFormat = 0
	IndexFormatUint16                = 1
	IndexFormatUint32                = 2
)

type InstanceFeatureName int

const (
	InstanceFeatureNameTimedWaitAny              InstanceFeatureName = 1
	InstanceFeatureNameShaderSourceSPIRV                             = 2
	InstanceFeatureNameMultipleDevicesPerAdapter                     = 3
)

type LoadOp int

const (
	LoadOpUndefined            LoadOp = 0
	LoadOpLoad                        = 1
	LoadOpClear                       = 2
	LoadOpExpandResolveTexture        = 3
)

type MapAsyncStatus int

const (
	MapAsyncStatusSuccess           MapAsyncStatus = 1
	MapAsyncStatusCallbackCancelled                = 2
	MapAsyncStatusError                            = 3
	MapAsyncStatusAborted                          = 4
)

type MapMode int

const (
	MapModeNone  MapMode = 0
	MapModeRead          = 1
	MapModeWrite         = 2
)

type MipmapFilterMode int

const (
	MipmapFilterModeUndefined MipmapFilterMode = 0
	MipmapFilterModeNearest                    = 1
	MipmapFilterModeLinear                     = 2
)

type OptionalBool int

const (
	OptionalBoolFalse     OptionalBool = 0
	OptionalBoolTrue                   = 1
	OptionalBoolUndefined              = 2
)

type popErrorScopeStatus int

const (
	popErrorScopeStatusSuccess           popErrorScopeStatus = 1
	popErrorScopeStatusCallbackCancelled                     = 2
	popErrorScopeStatusError                                 = 3
)

type PowerPreference int

const (
	PowerPreferenceUndefined       PowerPreference = 0
	PowerPreferenceLowPower                        = 1
	PowerPreferenceHighPerformance                 = 2
)

type PredefinedColorSpace int

const (
	PredefinedColorSpaceSRGB            PredefinedColorSpace = 1
	PredefinedColorSpaceDisplayP3                            = 2
	PredefinedColorSpaceSRGBLinear                           = 3
	PredefinedColorSpaceDisplayP3Linear                      = 4
)

type PresentMode int

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

type PrimitiveTopology int

const (
	PrimitiveTopologyUndefined     PrimitiveTopology = 0
	PrimitiveTopologyPointList                       = 1
	PrimitiveTopologyLineList                        = 2
	PrimitiveTopologyLineStrip                       = 3
	PrimitiveTopologyTriangleList                    = 4
	PrimitiveTopologyTriangleStrip                   = 5
)

type QueryType int

const (
	QueryTypeOcclusion QueryType = 1
	QueryTypeTimestamp           = 2
)

type QueueWorkDoneStatus int

const (
	QueueWorkDoneStatusSuccess           QueueWorkDoneStatus = 1
	QueueWorkDoneStatusCallbackCancelled                     = 2
	QueueWorkDoneStatusError                                 = 3
)

type requestAdapterStatus int

const (
	requestAdapterStatusSuccess           requestAdapterStatus = 1
	requestAdapterStatusCallbackCancelled                      = 2
	requestAdapterStatusUnavailable                            = 3
	requestAdapterStatusError                                  = 4
)

type requestDeviceStatus int

const (
	requestDeviceStatusSuccess           requestDeviceStatus = 1
	requestDeviceStatusCallbackCancelled                     = 2
	requestDeviceStatusError                                 = 3
)

type SType int

const (
	STypeShaderSourceSPIRV                                              SType = 1
	STypeShaderSourceWGSL                                                     = 2
	STypeRenderPassMaxDrawCount                                               = 3
	STypeSurfaceSourceMetalLayer                                              = 4
	STypeSurfaceSourceWindowsHWND                                             = 5
	STypeSurfaceSourceXlibWindow                                              = 6
	STypeSurfaceSourceWaylandSurface                                          = 7
	STypeSurfaceSourceAndroidNativeWindow                                     = 8
	STypeSurfaceSourceXCBWindow                                               = 9
	STypeSurfaceColorManagement                                               = 10
	STypeRequestAdapterWebXROptions                                           = 11
	STypeTextureComponentSwizzleDescriptor                                    = 12
	STypeExternalTextureBindingLayout                                         = 13
	STypeExternalTextureBindingEntry                                          = 14
	STypeCompatibilityModeLimits                                              = 15
	STypeTextureBindingViewDimension                                          = 16
	STypeEmscriptenSurfaceSourceCanvasHTMLSelector                            = 0
	STypeSurfaceDescriptorFromWindowsCoreWindow                               = 0
	STypeSurfaceDescriptorFromWindowsUWPSwapChainPanel                        = 3
	STypeDawnTextureInternalUsageDescriptor                                   = 4
	STypeDawnEncoderInternalUsageDescriptor                                   = 5
	STypeDawnInstanceDescriptor                                               = 6
	STypeDawnCacheDeviceDescriptor                                            = 7
	STypeDawnAdapterPropertiesPowerPreference                                 = 8
	STypeDawnBufferDescriptorErrorInfoFromWireClient                          = 9
	STypeDawnTogglesDescriptor                                                = 10
	STypeDawnShaderModuleSPIRVOptionsDescriptor                               = 11
	STypeRequestAdapterOptionsLUID                                            = 12
	STypeRequestAdapterOptionsGetGLProc                                       = 13
	STypeRequestAdapterOptionsD3D11Device                                     = 14
	STypeDawnRenderPassSampleCount                                            = 15
	STypeRenderPassPixelLocalStorage                                          = 16
	STypePipelineLayoutPixelLocalStorage                                      = 17
	STypeBufferHostMappedPointer                                              = 18
	STypeAdapterPropertiesMemoryHeaps                                         = 19
	STypeAdapterPropertiesD3D                                                 = 20
	STypeAdapterPropertiesVk                                                  = 21
	STypeDawnWireWGSLControl                                                  = 22
	STypeDawnWGSLBlocklist                                                    = 23
	STypeDawnDrmFormatCapabilities                                            = 24
	STypeShaderModuleCompilationOptions                                       = 25
	STypeColorTargetStateExpandResolveTextureDawn                             = 26
	STypeRenderPassRenderAreaRect                                             = 27
	STypeSharedTextureMemoryVkDedicatedAllocationDescriptor                   = 28
	STypeSharedTextureMemoryAHardwareBufferDescriptor                         = 29
	STypeSharedTextureMemoryDmaBufDescriptor                                  = 30
	STypeSharedTextureMemoryOpaqueFDDescriptor                                = 31
	STypeSharedTextureMemoryZirconHandleDescriptor                            = 32
	STypeSharedTextureMemoryDXGISharedHandleDescriptor                        = 33
	STypeSharedTextureMemoryD3D11Texture2DDescriptor                          = 34
	STypeSharedTextureMemoryIOSurfaceDescriptor                               = 35
	STypeSharedTextureMemoryEGLImageDescriptor                                = 36
	STypeSharedTextureMemoryInitializedBeginState                             = 37
	STypeSharedTextureMemoryInitializedEndState                               = 38
	STypeSharedTextureMemoryVkImageLayoutBeginState                           = 39
	STypeSharedTextureMemoryVkImageLayoutEndState                             = 40
	STypeSharedTextureMemoryD3DSwapchainBeginState                            = 41
	STypeSharedFenceVkSemaphoreOpaqueFDDescriptor                             = 42
	STypeSharedFenceVkSemaphoreOpaqueFDExportInfo                             = 43
	STypeSharedFenceSyncFDDescriptor                                          = 44
	STypeSharedFenceSyncFDExportInfo                                          = 45
	STypeSharedFenceVkSemaphoreZirconHandleDescriptor                         = 46
	STypeSharedFenceVkSemaphoreZirconHandleExportInfo                         = 47
	STypeSharedFenceDXGISharedHandleDescriptor                                = 48
	STypeSharedFenceDXGISharedHandleExportInfo                                = 49
	STypeSharedFenceMTLSharedEventDescriptor                                  = 50
	STypeSharedFenceMTLSharedEventExportInfo                                  = 51
	STypeSharedBufferMemoryD3D12ResourceDescriptor                            = 52
	STypeStaticSamplerBindingLayout                                           = 53
	STypeYCbCrVkDescriptor                                                    = 54
	STypeSharedTextureMemoryAHardwareBufferProperties                         = 55
	STypeAHardwareBufferProperties                                            = 56
	STypeDawnTexelCopyBufferRowAlignmentLimits                                = 58
	STypeAdapterPropertiesSubgroupMatrixConfigs                               = 59
	STypeSharedFenceEGLSyncDescriptor                                         = 60
	STypeSharedFenceEGLSyncExportInfo                                         = 61
	STypeDawnInjectedInvalidSType                                             = 62
	STypeDawnCompilationMessageUtf16                                          = 63
	STypeDawnFakeBufferOOMForTesting                                          = 64
	STypeSurfaceDescriptorFromWindowsWinUISwapChainPanel                      = 65
	STypeDawnDeviceAllocatorControl                                           = 66
	STypeDawnHostMappedPointerLimits                                          = 67
	STypeRenderPassDescriptorResolveRect                                      = 68
	STypeRequestAdapterWebGPUBackendOptions                                   = 69
	STypeDawnFakeDeviceInitializeErrorForTesting                              = 70
	STypeSharedTextureMemoryD3D11BeginState                                   = 71
	STypeDawnConsumeAdapterDescriptor                                         = 72
	STypeTexelBufferBindingEntry                                              = 73
	STypeTexelBufferBindingLayout                                             = 74
	STypeSharedTextureMemoryMetalEndAccessState                               = 75
	STypeAdapterPropertiesWGPU                                                = 76
	STypeSharedBufferMemoryD3D12SharedMemoryFileMappingHandleDescriptor       = 77
	STypeSharedTextureMemoryD3D12ResourceDescriptor                           = 78
	STypeRequestAdapterOptionsAngleVirtualizationGroup                        = 79
	STypePipelineLayoutResourceTable                                          = 80
	STypeAdapterPropertiesExplicitComputeSubgroupSizeConfigs                  = 81
	STypeAdapterPropertiesDrm                                                 = 82
)

type SamplerBindingType int

const (
	SamplerBindingTypeBindingNotUsed SamplerBindingType = 0
	SamplerBindingTypeUndefined                         = 1
	SamplerBindingTypeFiltering                         = 2
	SamplerBindingTypeNonFiltering                      = 3
	SamplerBindingTypeComparison                        = 4
)

type ShaderStage int

const (
	ShaderStageNone     ShaderStage = 0
	ShaderStageVertex               = 1
	ShaderStageFragment             = 2
	ShaderStageCompute              = 4
)

type statusCode int

const (
	statusCodeSuccess statusCode = 1
	statusCodeError              = 2
)

type StencilOperation int

const (
	StencilOperationUndefined      StencilOperation = 0
	StencilOperationKeep                            = 1
	StencilOperationZero                            = 2
	StencilOperationReplace                         = 3
	StencilOperationInvert                          = 4
	StencilOperationIncrementClamp                  = 5
	StencilOperationDecrementClamp                  = 6
	StencilOperationIncrementWrap                   = 7
	StencilOperationDecrementWrap                   = 8
)

type StorageTextureAccess int

const (
	StorageTextureAccessBindingNotUsed StorageTextureAccess = 0
	StorageTextureAccessUndefined                           = 1
	StorageTextureAccessWriteOnly                           = 2
	StorageTextureAccessReadOnly                            = 3
	StorageTextureAccessReadWrite                           = 4
)

type StoreOp int

const (
	StoreOpUndefined StoreOp = 0
	StoreOpStore             = 1
	StoreOpDiscard           = 2
)

type SurfaceGetCurrentTextureStatus int

const (
	SurfaceGetCurrentTextureStatusSuccessOptimal    SurfaceGetCurrentTextureStatus = 1
	SurfaceGetCurrentTextureStatusSuccessSuboptimal                                = 2
	SurfaceGetCurrentTextureStatusTimeout                                          = 3
	SurfaceGetCurrentTextureStatusOutdated                                         = 4
	SurfaceGetCurrentTextureStatusLost                                             = 5
	SurfaceGetCurrentTextureStatusError                                            = 6
)

type TextureAspect int

const (
	TextureAspectUndefined   TextureAspect = 0
	TextureAspectAll                       = 1
	TextureAspectStencilOnly               = 2
	TextureAspectDepthOnly                 = 3
	TextureAspectPlane0Only                = 0
	TextureAspectPlane1Only                = 1
	TextureAspectPlane2Only                = 2
)

type TextureDimension int

const (
	TextureDimensionUndefined TextureDimension = 0
	TextureDimension1D                         = 1
	TextureDimension2D                         = 2
	TextureDimension3D                         = 3
)

type TextureFormat int

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

type TextureSampleType int

const (
	TextureSampleTypeBindingNotUsed    TextureSampleType = 0
	TextureSampleTypeUndefined                           = 1
	TextureSampleTypeFloat                               = 2
	TextureSampleTypeUnfilterableFloat                   = 3
	TextureSampleTypeDepth                               = 4
	TextureSampleTypeSint                                = 5
	TextureSampleTypeUint                                = 6
)

type TextureUsage int

const (
	TextureUsageNone                TextureUsage = 0
	TextureUsageCopySrc                          = 1
	TextureUsageCopyDst                          = 2
	TextureUsageTextureBinding                   = 4
	TextureUsageStorageBinding                   = 8
	TextureUsageRenderAttachment                 = 16
	TextureUsageTransientAttachment              = 32
	TextureUsageStorageAttachment                = 64
)

type TextureViewDimension int

const (
	TextureViewDimensionUndefined TextureViewDimension = 0
	TextureViewDimension1D                             = 1
	TextureViewDimension2D                             = 2
	TextureViewDimension2DArray                        = 3
	TextureViewDimensionCube                           = 4
	TextureViewDimensionCubeArray                      = 5
	TextureViewDimension3D                             = 6
)

type ToneMappingMode int

const (
	ToneMappingModeStandard ToneMappingMode = 1
	ToneMappingModeExtended                 = 2
)

type VertexFormat int

const (
	VertexFormatUint8           VertexFormat = 1
	VertexFormatUint8x2                      = 2
	VertexFormatUint8x4                      = 3
	VertexFormatSint8                        = 4
	VertexFormatSint8x2                      = 5
	VertexFormatSint8x4                      = 6
	VertexFormatUnorm8                       = 7
	VertexFormatUnorm8x2                     = 8
	VertexFormatUnorm8x4                     = 9
	VertexFormatSnorm8                       = 10
	VertexFormatSnorm8x2                     = 11
	VertexFormatSnorm8x4                     = 12
	VertexFormatUint16                       = 13
	VertexFormatUint16x2                     = 14
	VertexFormatUint16x4                     = 15
	VertexFormatSint16                       = 16
	VertexFormatSint16x2                     = 17
	VertexFormatSint16x4                     = 18
	VertexFormatUnorm16                      = 19
	VertexFormatUnorm16x2                    = 20
	VertexFormatUnorm16x4                    = 21
	VertexFormatSnorm16                      = 22
	VertexFormatSnorm16x2                    = 23
	VertexFormatSnorm16x4                    = 24
	VertexFormatFloat16                      = 25
	VertexFormatFloat16x2                    = 26
	VertexFormatFloat16x4                    = 27
	VertexFormatFloat32                      = 28
	VertexFormatFloat32x2                    = 29
	VertexFormatFloat32x3                    = 30
	VertexFormatFloat32x4                    = 31
	VertexFormatUint32                       = 32
	VertexFormatUint32x2                     = 33
	VertexFormatUint32x3                     = 34
	VertexFormatUint32x4                     = 35
	VertexFormatSint32                       = 36
	VertexFormatSint32x2                     = 37
	VertexFormatSint32x3                     = 38
	VertexFormatSint32x4                     = 39
	VertexFormatUnorm10_10_10_2              = 40
	VertexFormatUnorm8x4BGRA                 = 41
)

type VertexStepMode int

const (
	VertexStepModeUndefined VertexStepMode = 0
	VertexStepModeVertex                   = 1
	VertexStepModeInstance                 = 2
)

type WaitStatus int

const (
	WaitStatusSuccess  WaitStatus = 1
	WaitStatusTimedOut            = 2
	WaitStatusError               = 3
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
