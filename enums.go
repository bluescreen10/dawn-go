// CODE GENERATED. DO NOT EDIT
//go:generate go run ./cmd/wrapper/. 
package wgpu

type DeviceLostReason int

const (
	DeviceLostReasonUnknown DeviceLostReason =  1
	DeviceLostReasonDestroyed = 2
	DeviceLostReasonCallbackCancelled = 3
	DeviceLostReasonFailedCreation = 4
)

type ErrorFilter int

const (
	ErrorFilterValidation ErrorFilter =  1
	ErrorFilterOutOfMemory = 2
	ErrorFilterInternal = 3
)

type TextureFormat int

const (
	TextureFormatUndefined TextureFormat =  0
	TextureFormatR8Unorm = 1
	TextureFormatR8Snorm = 2
	TextureFormatR8Uint = 3
	TextureFormatR8Sint = 4
	TextureFormatR16Unorm = 5
	TextureFormatR16Snorm = 6
	TextureFormatR16Uint = 7
	TextureFormatR16Sint = 8
	TextureFormatR16Float = 9
	TextureFormatRG8Unorm = 10
	TextureFormatRG8Snorm = 11
	TextureFormatRG8Uint = 12
	TextureFormatRG8Sint = 13
	TextureFormatR32Float = 14
	TextureFormatR32Uint = 15
	TextureFormatR32Sint = 16
	TextureFormatRG16Unorm = 17
	TextureFormatRG16Snorm = 18
	TextureFormatRG16Uint = 19
	TextureFormatRG16Sint = 20
	TextureFormatRG16Float = 21
	TextureFormatRGBA8Unorm = 22
	TextureFormatRGBA8UnormSrgb = 23
	TextureFormatRGBA8Snorm = 24
	TextureFormatRGBA8Uint = 25
	TextureFormatRGBA8Sint = 26
	TextureFormatBGRA8Unorm = 27
	TextureFormatBGRA8UnormSrgb = 28
	TextureFormatRGB10A2Uint = 29
	TextureFormatRGB10A2Unorm = 30
	TextureFormatRG11B10Ufloat = 31
	TextureFormatRGB9E5Ufloat = 32
	TextureFormatRG32Float = 33
	TextureFormatRG32Uint = 34
	TextureFormatRG32Sint = 35
	TextureFormatRGBA16Unorm = 36
	TextureFormatRGBA16Snorm = 37
	TextureFormatRGBA16Uint = 38
	TextureFormatRGBA16Sint = 39
	TextureFormatRGBA16Float = 40
	TextureFormatRGBA32Float = 41
	TextureFormatRGBA32Uint = 42
	TextureFormatRGBA32Sint = 43
	TextureFormatStencil8 = 44
	TextureFormatDepth16Unorm = 45
	TextureFormatDepth24Plus = 46
	TextureFormatDepth24PlusStencil8 = 47
	TextureFormatDepth32Float = 48
	TextureFormatDepth32FloatStencil8 = 49
	TextureFormatBC1RGBAUnorm = 50
	TextureFormatBC1RGBAUnormSrgb = 51
	TextureFormatBC2RGBAUnorm = 52
	TextureFormatBC2RGBAUnormSrgb = 53
	TextureFormatBC3RGBAUnorm = 54
	TextureFormatBC3RGBAUnormSrgb = 55
	TextureFormatBC4RUnorm = 56
	TextureFormatBC4RSnorm = 57
	TextureFormatBC5RGUnorm = 58
	TextureFormatBC5RGSnorm = 59
	TextureFormatBC6HRGBUfloat = 60
	TextureFormatBC6HRGBFloat = 61
	TextureFormatBC7RGBAUnorm = 62
	TextureFormatBC7RGBAUnormSrgb = 63
	TextureFormatETC2RGB8Unorm = 64
	TextureFormatETC2RGB8UnormSrgb = 65
	TextureFormatETC2RGB8A1Unorm = 66
	TextureFormatETC2RGB8A1UnormSrgb = 67
	TextureFormatETC2RGBA8Unorm = 68
	TextureFormatETC2RGBA8UnormSrgb = 69
	TextureFormatEACR11Unorm = 70
	TextureFormatEACR11Snorm = 71
	TextureFormatEACRG11Unorm = 72
	TextureFormatEACRG11Snorm = 73
	TextureFormatASTC4x4Unorm = 74
	TextureFormatASTC4x4UnormSrgb = 75
	TextureFormatASTC5x4Unorm = 76
	TextureFormatASTC5x4UnormSrgb = 77
	TextureFormatASTC5x5Unorm = 78
	TextureFormatASTC5x5UnormSrgb = 79
	TextureFormatASTC6x5Unorm = 80
	TextureFormatASTC6x5UnormSrgb = 81
	TextureFormatASTC6x6Unorm = 82
	TextureFormatASTC6x6UnormSrgb = 83
	TextureFormatASTC8x5Unorm = 84
	TextureFormatASTC8x5UnormSrgb = 85
	TextureFormatASTC8x6Unorm = 86
	TextureFormatASTC8x6UnormSrgb = 87
	TextureFormatASTC8x8Unorm = 88
	TextureFormatASTC8x8UnormSrgb = 89
	TextureFormatASTC10x5Unorm = 90
	TextureFormatASTC10x5UnormSrgb = 91
	TextureFormatASTC10x6Unorm = 92
	TextureFormatASTC10x6UnormSrgb = 93
	TextureFormatASTC10x8Unorm = 94
	TextureFormatASTC10x8UnormSrgb = 95
	TextureFormatASTC10x10Unorm = 96
	TextureFormatASTC10x10UnormSrgb = 97
	TextureFormatASTC12x10Unorm = 98
	TextureFormatASTC12x10UnormSrgb = 99
	TextureFormatASTC12x12Unorm = 100
	TextureFormatASTC12x12UnormSrgb = 101
	TextureFormatR8BG8Biplanar420Unorm = 0
	TextureFormatR10X6BG10X6Biplanar420Unorm = 1
	TextureFormatR8BG8A8Triplanar420Unorm = 2
	TextureFormatR8BG8Biplanar422Unorm = 3
	TextureFormatR8BG8Biplanar444Unorm = 4
	TextureFormatR10X6BG10X6Biplanar422Unorm = 5
	TextureFormatR10X6BG10X6Biplanar444Unorm = 6
	TextureFormatOpaqueYCbCrAndroid = 7
)

type WGSLLanguageFeatureName int

const (
	WGSLLanguageFeatureNameReadonlyAndReadwriteStorageTextures WGSLLanguageFeatureName =  1
	WGSLLanguageFeatureNamePacked4x8IntegerDotProduct = 2
	WGSLLanguageFeatureNameUnrestrictedPointerParameters = 3
	WGSLLanguageFeatureNamePointerCompositeAccess = 4
	WGSLLanguageFeatureNameUniformBufferStandardLayout = 5
	WGSLLanguageFeatureNameSubgroupId = 6
	WGSLLanguageFeatureNameTextureAndSamplerLet = 7
	WGSLLanguageFeatureNameSubgroupUniformity = 8
	WGSLLanguageFeatureNameTextureFormatsTier1 = 9
	WGSLLanguageFeatureNameChromiumTestingUnimplemented = 0
	WGSLLanguageFeatureNameChromiumTestingUnsafeExperimental = 1
	WGSLLanguageFeatureNameChromiumTestingExperimental = 2
	WGSLLanguageFeatureNameChromiumTestingShippedWithKillswitch = 3
	WGSLLanguageFeatureNameChromiumTestingShipped = 4
	WGSLLanguageFeatureNameSizedBindingArray = 5
	WGSLLanguageFeatureNameTexelBuffers = 6
	WGSLLanguageFeatureNameChromiumPrint = 7
	WGSLLanguageFeatureNameFragmentDepth = 8
	WGSLLanguageFeatureNameImmediateAddressSpace = 9
	WGSLLanguageFeatureNameBufferView = 11
	WGSLLanguageFeatureNameFilteringParameters = 12
	WGSLLanguageFeatureNameSwizzleAssignment = 13
	WGSLLanguageFeatureNameLinearIndexing = 14
)

type PredefinedColorSpace int

const (
	PredefinedColorSpaceSRGB PredefinedColorSpace =  1
	PredefinedColorSpaceDisplayP3 = 2
	PredefinedColorSpaceSRGBLinear = 3
	PredefinedColorSpaceDisplayP3Linear = 4
)

type CompilationInfoRequestStatus int

const (
	CompilationInfoRequestStatusSuccess CompilationInfoRequestStatus =  1
	CompilationInfoRequestStatusCallbackCancelled = 2
)

type ColorWriteMask int

const (
	ColorWriteMaskNone ColorWriteMask =  0
	ColorWriteMaskRed = 1
	ColorWriteMaskGreen = 2
	ColorWriteMaskBlue = 4
	ColorWriteMaskAlpha = 8
	ColorWriteMaskAll = 15
)

type AddressMode int

const (
	AddressModeUndefined AddressMode =  0
	AddressModeClampToEdge = 1
	AddressModeRepeat = 2
	AddressModeMirrorRepeat = 3
)

type StoreOp int

const (
	StoreOpUndefined StoreOp =  0
	StoreOpStore = 1
	StoreOpDiscard = 2
)

type ComponentSwizzle int

const (
	ComponentSwizzleUndefined ComponentSwizzle =  0
	ComponentSwizzleZero = 1
	ComponentSwizzleOne = 2
	ComponentSwizzleR = 3
	ComponentSwizzleG = 4
	ComponentSwizzleB = 5
	ComponentSwizzleA = 6
)

type BlendOperation int

const (
	BlendOperationUndefined BlendOperation =  0
	BlendOperationAdd = 1
	BlendOperationSubtract = 2
	BlendOperationReverseSubtract = 3
	BlendOperationMin = 4
	BlendOperationMax = 5
)

type ShaderStage int

const (
	ShaderStageNone ShaderStage =  0
	ShaderStageVertex = 1
	ShaderStageFragment = 2
	ShaderStageCompute = 4
)

type CreatePipelineAsyncStatus int

const (
	CreatePipelineAsyncStatusSuccess CreatePipelineAsyncStatus =  1
	CreatePipelineAsyncStatusCallbackCancelled = 2
	CreatePipelineAsyncStatusValidationError = 3
	CreatePipelineAsyncStatusInternalError = 4
)

type VertexStepMode int

const (
	VertexStepModeUndefined VertexStepMode =  0
	VertexStepModeVertex = 1
	VertexStepModeInstance = 2
)

type ErrorType int

const (
	ErrorTypeNoError ErrorType =  1
	ErrorTypeValidation = 2
	ErrorTypeOutOfMemory = 3
	ErrorTypeInternal = 4
	ErrorTypeUnknown = 5
)

type BlendFactor int

const (
	BlendFactorUndefined BlendFactor =  0
	BlendFactorZero = 1
	BlendFactorOne = 2
	BlendFactorSrc = 3
	BlendFactorOneMinusSrc = 4
	BlendFactorSrcAlpha = 5
	BlendFactorOneMinusSrcAlpha = 6
	BlendFactorDst = 7
	BlendFactorOneMinusDst = 8
	BlendFactorDstAlpha = 9
	BlendFactorOneMinusDstAlpha = 10
	BlendFactorSrcAlphaSaturated = 11
	BlendFactorConstant = 12
	BlendFactorOneMinusConstant = 13
	BlendFactorSrc1 = 14
	BlendFactorOneMinusSrc1 = 15
	BlendFactorSrc1Alpha = 16
	BlendFactorOneMinusSrc1Alpha = 17
)

type IndexFormat int

const (
	IndexFormatUndefined IndexFormat =  0
	IndexFormatUint16 = 1
	IndexFormatUint32 = 2
)

type SamplerBindingType int

const (
	SamplerBindingTypeBindingNotUsed SamplerBindingType =  0
	SamplerBindingTypeUndefined = 1
	SamplerBindingTypeFiltering = 2
	SamplerBindingTypeNonFiltering = 3
	SamplerBindingTypeComparison = 4
)

type FilterMode int

const (
	FilterModeUndefined FilterMode =  0
	FilterModeNearest = 1
	FilterModeLinear = 2
)

type CallbackMode int

const (
	CallbackModeWaitAnyOnly CallbackMode =  1
	CallbackModeAllowProcessEvents = 2
	CallbackModeAllowSpontaneous = 3
)

type BufferBindingType int

const (
	BufferBindingTypeBindingNotUsed BufferBindingType =  0
	BufferBindingTypeUndefined = 1
	BufferBindingTypeUniform = 2
	BufferBindingTypeStorage = 3
	BufferBindingTypeReadOnlyStorage = 4
)

type PopErrorScopeStatus int

const (
	PopErrorScopeStatusSuccess PopErrorScopeStatus =  1
	PopErrorScopeStatusCallbackCancelled = 2
	PopErrorScopeStatusError = 3
)

type MapAsyncStatus int

const (
	MapAsyncStatusSuccess MapAsyncStatus =  1
	MapAsyncStatusCallbackCancelled = 2
	MapAsyncStatusError = 3
	MapAsyncStatusAborted = 4
)

type TextureAspect int

const (
	TextureAspectUndefined TextureAspect =  0
	TextureAspectAll = 1
	TextureAspectStencilOnly = 2
	TextureAspectDepthOnly = 3
	TextureAspectPlane0Only = 0
	TextureAspectPlane1Only = 1
	TextureAspectPlane2Only = 2
)

type AdapterType int

const (
	AdapterTypeDiscreteGPU AdapterType =  1
	AdapterTypeIntegratedGPU = 2
	AdapterTypeCPU = 3
	AdapterTypeUnknown = 4
)

type RequestDeviceStatus int

const (
	RequestDeviceStatusSuccess RequestDeviceStatus =  1
	RequestDeviceStatusCallbackCancelled = 2
	RequestDeviceStatusError = 3
)

type StencilOperation int

const (
	StencilOperationUndefined StencilOperation =  0
	StencilOperationKeep = 1
	StencilOperationZero = 2
	StencilOperationReplace = 3
	StencilOperationInvert = 4
	StencilOperationIncrementClamp = 5
	StencilOperationDecrementClamp = 6
	StencilOperationIncrementWrap = 7
	StencilOperationDecrementWrap = 8
)

type TextureViewDimension int

const (
	TextureViewDimensionUndefined TextureViewDimension =  0
	TextureViewDimension1D = 1
	TextureViewDimension2D = 2
	TextureViewDimension2DArray = 3
	TextureViewDimensionCube = 4
	TextureViewDimensionCubeArray = 5
	TextureViewDimension3D = 6
)

type BufferMapState int

const (
	BufferMapStateUnmapped BufferMapState =  1
	BufferMapStatePending = 2
	BufferMapStateMapped = 3
)

type CompositeAlphaMode int

const (
	CompositeAlphaModeAuto CompositeAlphaMode =  0
	CompositeAlphaModeOpaque = 1
	CompositeAlphaModePremultiplied = 2
	CompositeAlphaModeUnpremultiplied = 3
	CompositeAlphaModeInherit = 4
)

type CompilationMessageType int

const (
	CompilationMessageTypeError CompilationMessageType =  1
	CompilationMessageTypeWarning = 2
	CompilationMessageTypeInfo = 3
)

type QueryType int

const (
	QueryTypeOcclusion QueryType =  1
	QueryTypeTimestamp = 2
)

type TextureUsage int

const (
	TextureUsageNone TextureUsage =  0
	TextureUsageCopySrc = 1
	TextureUsageCopyDst = 2
	TextureUsageTextureBinding = 4
	TextureUsageStorageBinding = 8
	TextureUsageRenderAttachment = 16
	TextureUsageTransientAttachment = 32
	TextureUsageStorageAttachment = 64
)

type TextureDimension int

const (
	TextureDimensionUndefined TextureDimension =  0
	TextureDimension1D = 1
	TextureDimension2D = 2
	TextureDimension3D = 3
)

type OptionalBool int

const (
	OptionalBoolFalse OptionalBool =  0
	OptionalBoolTrue = 1
	OptionalBoolUndefined = 2
)

type Status int

const (
	StatusSuccess Status =  1
	StatusError = 2
)

type CompareFunction int

const (
	CompareFunctionUndefined CompareFunction =  0
	CompareFunctionNever = 1
	CompareFunctionLess = 2
	CompareFunctionEqual = 3
	CompareFunctionLessEqual = 4
	CompareFunctionGreater = 5
	CompareFunctionNotEqual = 6
	CompareFunctionGreaterEqual = 7
	CompareFunctionAlways = 8
)

type StorageTextureAccess int

const (
	StorageTextureAccessBindingNotUsed StorageTextureAccess =  0
	StorageTextureAccessUndefined = 1
	StorageTextureAccessWriteOnly = 2
	StorageTextureAccessReadOnly = 3
	StorageTextureAccessReadWrite = 4
)

type BufferUsage int

const (
	BufferUsageNone BufferUsage =  0
	BufferUsageMapRead = 1
	BufferUsageMapWrite = 2
	BufferUsageCopySrc = 4
	BufferUsageCopyDst = 8
	BufferUsageIndex = 16
	BufferUsageVertex = 32
	BufferUsageUniform = 64
	BufferUsageStorage = 128
	BufferUsageIndirect = 256
	BufferUsageQueryResolve = 512
	BufferUsageTexelBuffer = 1024
)

type SType int

const (
	STypeShaderSourceSPIRV SType =  1
	STypeShaderSourceWGSL = 2
	STypeRenderPassMaxDrawCount = 3
	STypeSurfaceSourceMetalLayer = 4
	STypeSurfaceSourceWindowsHWND = 5
	STypeSurfaceSourceXlibWindow = 6
	STypeSurfaceSourceWaylandSurface = 7
	STypeSurfaceSourceAndroidNativeWindow = 8
	STypeSurfaceSourceXCBWindow = 9
	STypeSurfaceColorManagement = 10
	STypeRequestAdapterWebXROptions = 11
	STypeTextureComponentSwizzleDescriptor = 12
	STypeExternalTextureBindingLayout = 13
	STypeExternalTextureBindingEntry = 14
	STypeCompatibilityModeLimits = 15
	STypeTextureBindingViewDimension = 16
	STypeEmscriptenSurfaceSourceCanvasHTMLSelector = 0
	STypeSurfaceDescriptorFromWindowsCoreWindow = 0
	STypeSurfaceDescriptorFromWindowsUWPSwapChainPanel = 3
	STypeDawnTextureInternalUsageDescriptor = 4
	STypeDawnEncoderInternalUsageDescriptor = 5
	STypeDawnInstanceDescriptor = 6
	STypeDawnCacheDeviceDescriptor = 7
	STypeDawnAdapterPropertiesPowerPreference = 8
	STypeDawnBufferDescriptorErrorInfoFromWireClient = 9
	STypeDawnTogglesDescriptor = 10
	STypeDawnShaderModuleSPIRVOptionsDescriptor = 11
	STypeRequestAdapterOptionsLUID = 12
	STypeRequestAdapterOptionsGetGLProc = 13
	STypeRequestAdapterOptionsD3D11Device = 14
	STypeDawnRenderPassSampleCount = 15
	STypeRenderPassPixelLocalStorage = 16
	STypePipelineLayoutPixelLocalStorage = 17
	STypeBufferHostMappedPointer = 18
	STypeAdapterPropertiesMemoryHeaps = 19
	STypeAdapterPropertiesD3D = 20
	STypeAdapterPropertiesVk = 21
	STypeDawnWireWGSLControl = 22
	STypeDawnWGSLBlocklist = 23
	STypeDawnDrmFormatCapabilities = 24
	STypeShaderModuleCompilationOptions = 25
	STypeColorTargetStateExpandResolveTextureDawn = 26
	STypeRenderPassRenderAreaRect = 27
	STypeSharedTextureMemoryVkDedicatedAllocationDescriptor = 28
	STypeSharedTextureMemoryAHardwareBufferDescriptor = 29
	STypeSharedTextureMemoryDmaBufDescriptor = 30
	STypeSharedTextureMemoryOpaqueFDDescriptor = 31
	STypeSharedTextureMemoryZirconHandleDescriptor = 32
	STypeSharedTextureMemoryDXGISharedHandleDescriptor = 33
	STypeSharedTextureMemoryD3D11Texture2DDescriptor = 34
	STypeSharedTextureMemoryIOSurfaceDescriptor = 35
	STypeSharedTextureMemoryEGLImageDescriptor = 36
	STypeSharedTextureMemoryInitializedBeginState = 37
	STypeSharedTextureMemoryInitializedEndState = 38
	STypeSharedTextureMemoryVkImageLayoutBeginState = 39
	STypeSharedTextureMemoryVkImageLayoutEndState = 40
	STypeSharedTextureMemoryD3DSwapchainBeginState = 41
	STypeSharedFenceVkSemaphoreOpaqueFDDescriptor = 42
	STypeSharedFenceVkSemaphoreOpaqueFDExportInfo = 43
	STypeSharedFenceSyncFDDescriptor = 44
	STypeSharedFenceSyncFDExportInfo = 45
	STypeSharedFenceVkSemaphoreZirconHandleDescriptor = 46
	STypeSharedFenceVkSemaphoreZirconHandleExportInfo = 47
	STypeSharedFenceDXGISharedHandleDescriptor = 48
	STypeSharedFenceDXGISharedHandleExportInfo = 49
	STypeSharedFenceMTLSharedEventDescriptor = 50
	STypeSharedFenceMTLSharedEventExportInfo = 51
	STypeSharedBufferMemoryD3D12ResourceDescriptor = 52
	STypeStaticSamplerBindingLayout = 53
	STypeYCbCrVkDescriptor = 54
	STypeSharedTextureMemoryAHardwareBufferProperties = 55
	STypeAHardwareBufferProperties = 56
	STypeDawnTexelCopyBufferRowAlignmentLimits = 58
	STypeAdapterPropertiesSubgroupMatrixConfigs = 59
	STypeSharedFenceEGLSyncDescriptor = 60
	STypeSharedFenceEGLSyncExportInfo = 61
	STypeDawnInjectedInvalidSType = 62
	STypeDawnCompilationMessageUtf16 = 63
	STypeDawnFakeBufferOOMForTesting = 64
	STypeSurfaceDescriptorFromWindowsWinUISwapChainPanel = 65
	STypeDawnDeviceAllocatorControl = 66
	STypeDawnHostMappedPointerLimits = 67
	STypeRenderPassDescriptorResolveRect = 68
	STypeRequestAdapterWebGPUBackendOptions = 69
	STypeDawnFakeDeviceInitializeErrorForTesting = 70
	STypeSharedTextureMemoryD3D11BeginState = 71
	STypeDawnConsumeAdapterDescriptor = 72
	STypeTexelBufferBindingEntry = 73
	STypeTexelBufferBindingLayout = 74
	STypeSharedTextureMemoryMetalEndAccessState = 75
	STypeAdapterPropertiesWGPU = 76
	STypeSharedBufferMemoryD3D12SharedMemoryFileMappingHandleDescriptor = 77
	STypeSharedTextureMemoryD3D12ResourceDescriptor = 78
	STypeRequestAdapterOptionsAngleVirtualizationGroup = 79
	STypePipelineLayoutResourceTable = 80
	STypeAdapterPropertiesExplicitComputeSubgroupSizeConfigs = 81
	STypeAdapterPropertiesDrm = 82
)

type QueueWorkDoneStatus int

const (
	QueueWorkDoneStatusSuccess QueueWorkDoneStatus =  1
	QueueWorkDoneStatusCallbackCancelled = 2
	QueueWorkDoneStatusError = 3
)

type CullMode int

const (
	CullModeUndefined CullMode =  0
	CullModeNone = 1
	CullModeFront = 2
	CullModeBack = 3
)

type MipmapFilterMode int

const (
	MipmapFilterModeUndefined MipmapFilterMode =  0
	MipmapFilterModeNearest = 1
	MipmapFilterModeLinear = 2
)

type PowerPreference int

const (
	PowerPreferenceUndefined PowerPreference =  0
	PowerPreferenceLowPower = 1
	PowerPreferenceHighPerformance = 2
)

type SurfaceGetCurrentTextureStatus int

const (
	SurfaceGetCurrentTextureStatusSuccessOptimal SurfaceGetCurrentTextureStatus =  1
	SurfaceGetCurrentTextureStatusSuccessSuboptimal = 2
	SurfaceGetCurrentTextureStatusTimeout = 3
	SurfaceGetCurrentTextureStatusOutdated = 4
	SurfaceGetCurrentTextureStatusLost = 5
	SurfaceGetCurrentTextureStatusError = 6
)

type WaitStatus int

const (
	WaitStatusSuccess WaitStatus =  1
	WaitStatusTimedOut = 2
	WaitStatusError = 3
)

type ToneMappingMode int

const (
	ToneMappingModeStandard ToneMappingMode =  1
	ToneMappingModeExtended = 2
)

type RequestAdapterStatus int

const (
	RequestAdapterStatusSuccess RequestAdapterStatus =  1
	RequestAdapterStatusCallbackCancelled = 2
	RequestAdapterStatusUnavailable = 3
	RequestAdapterStatusError = 4
)

type VertexFormat int

const (
	VertexFormatUint8 VertexFormat =  1
	VertexFormatUint8x2 = 2
	VertexFormatUint8x4 = 3
	VertexFormatSint8 = 4
	VertexFormatSint8x2 = 5
	VertexFormatSint8x4 = 6
	VertexFormatUnorm8 = 7
	VertexFormatUnorm8x2 = 8
	VertexFormatUnorm8x4 = 9
	VertexFormatSnorm8 = 10
	VertexFormatSnorm8x2 = 11
	VertexFormatSnorm8x4 = 12
	VertexFormatUint16 = 13
	VertexFormatUint16x2 = 14
	VertexFormatUint16x4 = 15
	VertexFormatSint16 = 16
	VertexFormatSint16x2 = 17
	VertexFormatSint16x4 = 18
	VertexFormatUnorm16 = 19
	VertexFormatUnorm16x2 = 20
	VertexFormatUnorm16x4 = 21
	VertexFormatSnorm16 = 22
	VertexFormatSnorm16x2 = 23
	VertexFormatSnorm16x4 = 24
	VertexFormatFloat16 = 25
	VertexFormatFloat16x2 = 26
	VertexFormatFloat16x4 = 27
	VertexFormatFloat32 = 28
	VertexFormatFloat32x2 = 29
	VertexFormatFloat32x3 = 30
	VertexFormatFloat32x4 = 31
	VertexFormatUint32 = 32
	VertexFormatUint32x2 = 33
	VertexFormatUint32x3 = 34
	VertexFormatUint32x4 = 35
	VertexFormatSint32 = 36
	VertexFormatSint32x2 = 37
	VertexFormatSint32x3 = 38
	VertexFormatSint32x4 = 39
	VertexFormatUnorm10_10_10_2 = 40
	VertexFormatUnorm8x4BGRA = 41
)

type FrontFace int

const (
	FrontFaceUndefined FrontFace =  0
	FrontFaceCCW = 1
	FrontFaceCW = 2
)

type LoadOp int

const (
	LoadOpUndefined LoadOp =  0
	LoadOpLoad = 1
	LoadOpClear = 2
	LoadOpExpandResolveTexture = 3
)

type BackendType int

const (
	BackendTypeUndefined BackendType =  0
	BackendTypeNull = 1
	BackendTypeWebGPU = 2
	BackendTypeD3D11 = 3
	BackendTypeD3D12 = 4
	BackendTypeMetal = 5
	BackendTypeVulkan = 6
	BackendTypeOpenGL = 7
	BackendTypeOpenGLES = 8
)

type FeatureName int

const (
	FeatureNameCoreFeaturesAndLimits FeatureName =  1
	FeatureNameDepthClipControl = 2
	FeatureNameDepth32FloatStencil8 = 3
	FeatureNameTextureCompressionBC = 4
	FeatureNameTextureCompressionBCSliced3D = 5
	FeatureNameTextureCompressionETC2 = 6
	FeatureNameTextureCompressionASTC = 7
	FeatureNameTextureCompressionASTCSliced3D = 8
	FeatureNameTimestampQuery = 9
	FeatureNameIndirectFirstInstance = 10
	FeatureNameShaderF16 = 11
	FeatureNameRG11B10UfloatRenderable = 12
	FeatureNameBGRA8UnormStorage = 13
	FeatureNameFloat32Filterable = 14
	FeatureNameFloat32Blendable = 15
	FeatureNameClipDistances = 16
	FeatureNameDualSourceBlending = 17
	FeatureNameSubgroups = 18
	FeatureNameTextureFormatsTier1 = 19
	FeatureNameTextureFormatsTier2 = 20
	FeatureNamePrimitiveIndex = 21
	FeatureNameTextureComponentSwizzle = 22
	FeatureNameDawnInternalUsages = 0
	FeatureNameDawnMultiPlanarFormats = 1
	FeatureNameDawnNative = 2
	FeatureNameChromiumExperimentalTimestampQueryInsidePasses = 3
	FeatureNameImplicitDeviceSynchronization = 4
	FeatureNameTransientAttachments = 6
	FeatureNameMSAARenderToSingleSampled = 7
	FeatureNameD3D11MultithreadProtected = 8
	FeatureNameANGLETextureSharing = 9
	FeatureNamePixelLocalStorageCoherent = 10
	FeatureNamePixelLocalStorageNonCoherent = 11
	FeatureNameUnorm16TextureFormats = 12
	FeatureNameMultiPlanarFormatExtendedUsages = 13
	FeatureNameMultiPlanarFormatP010 = 14
	FeatureNameHostMappedPointer = 15
	FeatureNameMultiPlanarRenderTargets = 16
	FeatureNameMultiPlanarFormatNv12a = 17
	FeatureNameFramebufferFetch = 18
	FeatureNameBufferMapExtendedUsages = 19
	FeatureNameAdapterPropertiesMemoryHeaps = 20
	FeatureNameAdapterPropertiesD3D = 21
	FeatureNameAdapterPropertiesVk = 22
	FeatureNameDawnFormatCapabilities = 23
	FeatureNameDawnDrmFormatCapabilities = 24
	FeatureNameMultiPlanarFormatNv16 = 25
	FeatureNameMultiPlanarFormatNv24 = 26
	FeatureNameMultiPlanarFormatP210 = 27
	FeatureNameMultiPlanarFormatP410 = 28
	FeatureNameSharedTextureMemoryVkDedicatedAllocation = 29
	FeatureNameSharedTextureMemoryAHardwareBuffer = 30
	FeatureNameSharedTextureMemoryDmaBuf = 31
	FeatureNameSharedTextureMemoryOpaqueFD = 32
	FeatureNameSharedTextureMemoryZirconHandle = 33
	FeatureNameSharedTextureMemoryDXGISharedHandle = 34
	FeatureNameSharedTextureMemoryD3D11Texture2D = 35
	FeatureNameSharedTextureMemoryIOSurface = 36
	FeatureNameSharedTextureMemoryEGLImage = 37
	FeatureNameSharedFenceVkSemaphoreOpaqueFD = 38
	FeatureNameSharedFenceSyncFD = 39
	FeatureNameSharedFenceVkSemaphoreZirconHandle = 40
	FeatureNameSharedFenceDXGISharedHandle = 41
	FeatureNameSharedFenceMTLSharedEvent = 42
	FeatureNameSharedBufferMemoryD3D12Resource = 43
	FeatureNameStaticSamplers = 44
	FeatureNameYCbCrVulkanSamplers = 45
	FeatureNameShaderModuleCompilationOptions = 46
	FeatureNameDawnLoadResolveTexture = 47
	FeatureNameDawnPartialLoadResolveTexture = 48
	FeatureNameMultiDrawIndirect = 49
	FeatureNameDawnTexelCopyBufferRowAlignment = 50
	FeatureNameFlexibleTextureViews = 51
	FeatureNameChromiumExperimentalSubgroupMatrix = 52
	FeatureNameSharedFenceEGLSync = 53
	FeatureNameDawnDeviceAllocatorControl = 54
	FeatureNameAdapterPropertiesWGPU = 55
	FeatureNameSharedBufferMemoryD3D12SharedMemoryFileMappingHandle = 56
	FeatureNameSharedTextureMemoryD3D12Resource = 57
	FeatureNameChromiumExperimentalSamplingResourceTable = 58
	FeatureNameChromiumExperimentalSubgroupSizeControl = 59
	FeatureNameAtomicVec2uMinMax = 60
	FeatureNameUnorm16FormatsForExternalTexture = 61
	FeatureNameOpaqueYCbCrAndroidForExternalTexture = 62
	FeatureNameUnorm16Filterable = 63
	FeatureNameRenderPassRenderArea = 64
	FeatureNameDawnNativeSpontaneousQueueEvents = 65
	FeatureNameAdapterPropertiesDrm = 66
)

type TextureSampleType int

const (
	TextureSampleTypeBindingNotUsed TextureSampleType =  0
	TextureSampleTypeUndefined = 1
	TextureSampleTypeFloat = 2
	TextureSampleTypeUnfilterableFloat = 3
	TextureSampleTypeDepth = 4
	TextureSampleTypeSint = 5
	TextureSampleTypeUint = 6
)

type InstanceFeatureName int

const (
	InstanceFeatureNameTimedWaitAny InstanceFeatureName =  1
	InstanceFeatureNameShaderSourceSPIRV = 2
	InstanceFeatureNameMultipleDevicesPerAdapter = 3
)

type FeatureLevel int

const (
	FeatureLevelUndefined FeatureLevel =  0
	FeatureLevelCompatibility = 1
	FeatureLevelCore = 2
)

type PresentMode int

const (
	PresentModeUndefined PresentMode =  0
	PresentModeFifo = 1
	PresentModeFifoRelaxed = 2
	PresentModeImmediate = 3
	PresentModeMailbox = 4
)

type PrimitiveTopology int

const (
	PrimitiveTopologyUndefined PrimitiveTopology =  0
	PrimitiveTopologyPointList = 1
	PrimitiveTopologyLineList = 2
	PrimitiveTopologyLineStrip = 3
	PrimitiveTopologyTriangleList = 4
	PrimitiveTopologyTriangleStrip = 5
)

type MapMode int

const (
	MapModeNone MapMode =  0
	MapModeRead = 1
	MapModeWrite = 2
)

