//go:build js

package wgpu

import "syscall/js"

func (f TextureFormat) toJS() string {
	switch f {
	case TextureFormatR8Unorm:
		return "r8unorm"
	case TextureFormatR8Snorm:
		return "r8snorm"
	case TextureFormatR8Uint:
		return "r8uint"
	case TextureFormatR8Sint:
		return "r8sint"
	case TextureFormatR16Unorm:
		return "r16unorm"
	case TextureFormatR16Snorm:
		return "r16snorm"
	case TextureFormatR16Uint:
		return "r16uint"
	case TextureFormatR16Sint:
		return "r16sint"
	case TextureFormatR16Float:
		return "r16float"
	case TextureFormatRG8Unorm:
		return "rg8unorm"
	case TextureFormatRG8Snorm:
		return "rg8snorm"
	case TextureFormatRG8Uint:
		return "rg8uint"
	case TextureFormatRG8Sint:
		return "rg8sint"
	case TextureFormatR32Float:
		return "r32float"
	case TextureFormatR32Uint:
		return "r32uint"
	case TextureFormatR32Sint:
		return "r32sint"
	case TextureFormatRG16Unorm:
		return "rg16unorm"
	case TextureFormatRG16Snorm:
		return "rg16snorm"
	case TextureFormatRG16Uint:
		return "rg16uint"
	case TextureFormatRG16Sint:
		return "rg16sint"
	case TextureFormatRG16Float:
		return "rg16float"
	case TextureFormatRGBA8Unorm:
		return "rgba8unorm"
	case TextureFormatRGBA8UnormSRGB:
		return "rgba8unorm-srgb"
	case TextureFormatRGBA8Snorm:
		return "rgba8snorm"
	case TextureFormatRGBA8Uint:
		return "rgba8uint"
	case TextureFormatRGBA8Sint:
		return "rgba8sint"
	case TextureFormatBGRA8Unorm:
		return "bgra8unorm"
	case TextureFormatBGRA8UnormSRGB:
		return "bgra8unorm-srgb"
	case TextureFormatRGB10A2Uint:
		return "rgb10a2uint"
	case TextureFormatRGB10A2Unorm:
		return "rgb10a2unorm"
	case TextureFormatRG11B10Ufloat:
		return "rg11b10ufloat"
	case TextureFormatRGB9E5Ufloat:
		return "rgb9e5ufloat"
	case TextureFormatRG32Float:
		return "rg32float"
	case TextureFormatRG32Uint:
		return "rg32uint"
	case TextureFormatRG32Sint:
		return "rg32sint"
	case TextureFormatRGBA16Unorm:
		return "rgba16unorm"
	case TextureFormatRGBA16Snorm:
		return "rgba16snorm"
	case TextureFormatRGBA16Uint:
		return "rgba16uint"
	case TextureFormatRGBA16Sint:
		return "rgba16sint"
	case TextureFormatRGBA16Float:
		return "rgba16float"
	case TextureFormatRGBA32Float:
		return "rgba32float"
	case TextureFormatRGBA32Uint:
		return "rgba32uint"
	case TextureFormatRGBA32Sint:
		return "rgba32sint"
	case TextureFormatStencil8:
		return "stencil8"
	case TextureFormatDepth16Unorm:
		return "depth16unorm"
	case TextureFormatDepth24Plus:
		return "depth24plus"
	case TextureFormatDepth24PlusStencil8:
		return "depth24plus-stencil8"
	case TextureFormatDepth32Float:
		return "depth32float"
	case TextureFormatDepth32FloatStencil8:
		return "depth32float-stencil8"
	case TextureFormatBC1RGBAUnorm:
		return "bc1-rgba-unorm"
	case TextureFormatBC1RGBAUnormSRGB:
		return "bc1-rgba-unorm-srgb"
	case TextureFormatBC2RGBAUnorm:
		return "bc2-rgba-unorm"
	case TextureFormatBC2RGBAUnormSRGB:
		return "bc2-rgba-unorm-srgb"
	case TextureFormatBC3RGBAUnorm:
		return "bc3-rgba-unorm"
	case TextureFormatBC3RGBAUnormSRGB:
		return "bc3-rgba-unorm-srgb"
	case TextureFormatBC4RUnorm:
		return "bc4-r-unorm"
	case TextureFormatBC4RSnorm:
		return "bc4-r-snorm"
	case TextureFormatBC5RGUnorm:
		return "bc5-rg-unorm"
	case TextureFormatBC5RGSnorm:
		return "bc5-rg-snorm"
	case TextureFormatBC6HRGBUfloat:
		return "bc6h-rgb-ufloat"
	case TextureFormatBC6HRGBFloat:
		return "bc6h-rgb-float"
	case TextureFormatBC7RGBAUnorm:
		return "bc7-rgba-unorm"
	case TextureFormatBC7RGBAUnormSRGB:
		return "bc7-rgba-unorm-srgb"
	case TextureFormatETC2RGB8Unorm:
		return "etc2-rgb8unorm"
	case TextureFormatETC2RGB8UnormSRGB:
		return "etc2-rgb8unorm-srgb"
	case TextureFormatETC2RGB8A1Unorm:
		return "etc2-rgb8a1unorm"
	case TextureFormatETC2RGB8A1UnormSRGB:
		return "etc2-rgb8a1unorm-srgb"
	case TextureFormatETC2RGBA8Unorm:
		return "etc2-rgba8unorm"
	case TextureFormatETC2RGBA8UnormSRGB:
		return "etc2-rgba8unorm-srgb"
	case TextureFormatEACR11Unorm:
		return "eac-r11unorm"
	case TextureFormatEACR11Snorm:
		return "eac-r11snorm"
	case TextureFormatEACRG11Unorm:
		return "eac-rg11unorm"
	case TextureFormatEACRG11Snorm:
		return "eac-rg11snorm"
	case TextureFormatASTC4x4Unorm:
		return "astc-4x4-unorm"
	case TextureFormatASTC4x4UnormSRGB:
		return "astc-4x4-unorm-srgb"
	case TextureFormatASTC5x4Unorm:
		return "astc-5x4-unorm"
	case TextureFormatASTC5x4UnormSRGB:
		return "astc-5x4-unorm-srgb"
	case TextureFormatASTC5x5Unorm:
		return "astc-5x5-unorm"
	case TextureFormatASTC5x5UnormSRGB:
		return "astc-5x5-unorm-srgb"
	case TextureFormatASTC6x5Unorm:
		return "astc-6x5-unorm"
	case TextureFormatASTC6x5UnormSRGB:
		return "astc-6x5-unorm-srgb"
	case TextureFormatASTC6x6Unorm:
		return "astc-6x6-unorm"
	case TextureFormatASTC6x6UnormSRGB:
		return "astc-6x6-unorm-srgb"
	case TextureFormatASTC8x5Unorm:
		return "astc-8x5-unorm"
	case TextureFormatASTC8x5UnormSRGB:
		return "astc-8x5-unorm-srgb"
	case TextureFormatASTC8x6Unorm:
		return "astc-8x6-unorm"
	case TextureFormatASTC8x6UnormSRGB:
		return "astc-8x6-unorm-srgb"
	case TextureFormatASTC8x8Unorm:
		return "astc-8x8-unorm"
	case TextureFormatASTC8x8UnormSRGB:
		return "astc-8x8-unorm-srgb"
	case TextureFormatASTC10x5Unorm:
		return "astc-10x5-unorm"
	case TextureFormatASTC10x5UnormSRGB:
		return "astc-10x5-unorm-srgb"
	case TextureFormatASTC10x6Unorm:
		return "astc-10x6-unorm"
	case TextureFormatASTC10x6UnormSRGB:
		return "astc-10x6-unorm-srgb"
	case TextureFormatASTC10x8Unorm:
		return "astc-10x8-unorm"
	case TextureFormatASTC10x8UnormSRGB:
		return "astc-10x8-unorm-srgb"
	case TextureFormatASTC10x10Unorm:
		return "astc-10x10-unorm"
	case TextureFormatASTC10x10UnormSRGB:
		return "astc-10x10-unorm-srgb"
	case TextureFormatASTC12x10Unorm:
		return "astc-12x10-unorm"
	case TextureFormatASTC12x10UnormSRGB:
		return "astc-12x10-unorm-srgb"
	case TextureFormatASTC12x12Unorm:
		return "astc-12x12-unorm"
	case TextureFormatASTC12x12UnormSRGB:
		return "astc-12x12-unorm-srgb"
	default:
		return ""
	}
}

func textureFormatFromJS(s string) TextureFormat {
	switch s {
	case "r8unorm":
		return TextureFormatR8Unorm
	case "r8snorm":
		return TextureFormatR8Snorm
	case "r8uint":
		return TextureFormatR8Uint
	case "r8sint":
		return TextureFormatR8Sint
	case "r16unorm":
		return TextureFormatR16Unorm
	case "r16snorm":
		return TextureFormatR16Snorm
	case "r16uint":
		return TextureFormatR16Uint
	case "r16sint":
		return TextureFormatR16Sint
	case "r16float":
		return TextureFormatR16Float
	case "rg8unorm":
		return TextureFormatRG8Unorm
	case "rg8snorm":
		return TextureFormatRG8Snorm
	case "rg8uint":
		return TextureFormatRG8Uint
	case "rg8sint":
		return TextureFormatRG8Sint
	case "r32float":
		return TextureFormatR32Float
	case "r32uint":
		return TextureFormatR32Uint
	case "r32sint":
		return TextureFormatR32Sint
	case "rg16unorm":
		return TextureFormatRG16Unorm
	case "rg16snorm":
		return TextureFormatRG16Snorm
	case "rg16uint":
		return TextureFormatRG16Uint
	case "rg16sint":
		return TextureFormatRG16Sint
	case "rg16float":
		return TextureFormatRG16Float
	case "rgba8unorm":
		return TextureFormatRGBA8Unorm
	case "rgba8unorm-srgb":
		return TextureFormatRGBA8UnormSRGB
	case "rgba8snorm":
		return TextureFormatRGBA8Snorm
	case "rgba8uint":
		return TextureFormatRGBA8Uint
	case "rgba8sint":
		return TextureFormatRGBA8Sint
	case "bgra8unorm":
		return TextureFormatBGRA8Unorm
	case "bgra8unorm-srgb":
		return TextureFormatBGRA8UnormSRGB
	case "rgb10a2uint":
		return TextureFormatRGB10A2Uint
	case "rgb10a2unorm":
		return TextureFormatRGB10A2Unorm
	case "rg11b10ufloat":
		return TextureFormatRG11B10Ufloat
	case "rgb9e5ufloat":
		return TextureFormatRGB9E5Ufloat
	case "rg32float":
		return TextureFormatRG32Float
	case "rg32uint":
		return TextureFormatRG32Uint
	case "rg32sint":
		return TextureFormatRG32Sint
	case "rgba16unorm":
		return TextureFormatRGBA16Unorm
	case "rgba16snorm":
		return TextureFormatRGBA16Snorm
	case "rgba16uint":
		return TextureFormatRGBA16Uint
	case "rgba16sint":
		return TextureFormatRGBA16Sint
	case "rgba16float":
		return TextureFormatRGBA16Float
	case "rgba32float":
		return TextureFormatRGBA32Float
	case "rgba32uint":
		return TextureFormatRGBA32Uint
	case "rgba32sint":
		return TextureFormatRGBA32Sint
	case "stencil8":
		return TextureFormatStencil8
	case "depth16unorm":
		return TextureFormatDepth16Unorm
	case "depth24plus":
		return TextureFormatDepth24Plus
	case "depth24plus-stencil8":
		return TextureFormatDepth24PlusStencil8
	case "depth32float":
		return TextureFormatDepth32Float
	case "depth32float-stencil8":
		return TextureFormatDepth32FloatStencil8
	case "bc1-rgba-unorm":
		return TextureFormatBC1RGBAUnorm
	case "bc1-rgba-unorm-srgb":
		return TextureFormatBC1RGBAUnormSRGB
	case "bc2-rgba-unorm":
		return TextureFormatBC2RGBAUnorm
	case "bc2-rgba-unorm-srgb":
		return TextureFormatBC2RGBAUnormSRGB
	case "bc3-rgba-unorm":
		return TextureFormatBC3RGBAUnorm
	case "bc3-rgba-unorm-srgb":
		return TextureFormatBC3RGBAUnormSRGB
	case "bc4-r-unorm":
		return TextureFormatBC4RUnorm
	case "bc4-r-snorm":
		return TextureFormatBC4RSnorm
	case "bc5-rg-unorm":
		return TextureFormatBC5RGUnorm
	case "bc5-rg-snorm":
		return TextureFormatBC5RGSnorm
	case "bc6h-rgb-ufloat":
		return TextureFormatBC6HRGBUfloat
	case "bc6h-rgb-float":
		return TextureFormatBC6HRGBFloat
	case "bc7-rgba-unorm":
		return TextureFormatBC7RGBAUnorm
	case "bc7-rgba-unorm-srgb":
		return TextureFormatBC7RGBAUnormSRGB
	case "etc2-rgb8unorm":
		return TextureFormatETC2RGB8Unorm
	case "etc2-rgb8unorm-srgb":
		return TextureFormatETC2RGB8UnormSRGB
	case "etc2-rgb8a1unorm":
		return TextureFormatETC2RGB8A1Unorm
	case "etc2-rgb8a1unorm-srgb":
		return TextureFormatETC2RGB8A1UnormSRGB
	case "etc2-rgba8unorm":
		return TextureFormatETC2RGBA8Unorm
	case "etc2-rgba8unorm-srgb":
		return TextureFormatETC2RGBA8UnormSRGB
	case "eac-r11unorm":
		return TextureFormatEACR11Unorm
	case "eac-r11snorm":
		return TextureFormatEACR11Snorm
	case "eac-rg11unorm":
		return TextureFormatEACRG11Unorm
	case "eac-rg11snorm":
		return TextureFormatEACRG11Snorm
	case "astc-4x4-unorm":
		return TextureFormatASTC4x4Unorm
	case "astc-4x4-unorm-srgb":
		return TextureFormatASTC4x4UnormSRGB
	case "astc-5x4-unorm":
		return TextureFormatASTC5x4Unorm
	case "astc-5x4-unorm-srgb":
		return TextureFormatASTC5x4UnormSRGB
	case "astc-5x5-unorm":
		return TextureFormatASTC5x5Unorm
	case "astc-5x5-unorm-srgb":
		return TextureFormatASTC5x5UnormSRGB
	case "astc-6x5-unorm":
		return TextureFormatASTC6x5Unorm
	case "astc-6x5-unorm-srgb":
		return TextureFormatASTC6x5UnormSRGB
	case "astc-6x6-unorm":
		return TextureFormatASTC6x6Unorm
	case "astc-6x6-unorm-srgb":
		return TextureFormatASTC6x6UnormSRGB
	case "astc-8x5-unorm":
		return TextureFormatASTC8x5Unorm
	case "astc-8x5-unorm-srgb":
		return TextureFormatASTC8x5UnormSRGB
	case "astc-8x6-unorm":
		return TextureFormatASTC8x6Unorm
	case "astc-8x6-unorm-srgb":
		return TextureFormatASTC8x6UnormSRGB
	case "astc-8x8-unorm":
		return TextureFormatASTC8x8Unorm
	case "astc-8x8-unorm-srgb":
		return TextureFormatASTC8x8UnormSRGB
	case "astc-10x5-unorm":
		return TextureFormatASTC10x5Unorm
	case "astc-10x5-unorm-srgb":
		return TextureFormatASTC10x5UnormSRGB
	case "astc-10x6-unorm":
		return TextureFormatASTC10x6Unorm
	case "astc-10x6-unorm-srgb":
		return TextureFormatASTC10x6UnormSRGB
	case "astc-10x8-unorm":
		return TextureFormatASTC10x8Unorm
	case "astc-10x8-unorm-srgb":
		return TextureFormatASTC10x8UnormSRGB
	case "astc-10x10-unorm":
		return TextureFormatASTC10x10Unorm
	case "astc-10x10-unorm-srgb":
		return TextureFormatASTC10x10UnormSRGB
	case "astc-12x10-unorm":
		return TextureFormatASTC12x10Unorm
	case "astc-12x10-unorm-srgb":
		return TextureFormatASTC12x10UnormSRGB
	case "astc-12x12-unorm":
		return TextureFormatASTC12x12Unorm
	case "astc-12x12-unorm-srgb":
		return TextureFormatASTC12x12UnormSRGB
	default:
		return TextureFormatUndefined
	}
}

func (f VertexFormat) toJS() string {
	switch f {
	case VertexFormatUint8:
		return "uint8"
	case VertexFormatUint8x2:
		return "uint8x2"
	case VertexFormatUint8x4:
		return "uint8x4"
	case VertexFormatSint8:
		return "sint8"
	case VertexFormatSint8x2:
		return "sint8x2"
	case VertexFormatSint8x4:
		return "sint8x4"
	case VertexFormatUnorm8:
		return "unorm8"
	case VertexFormatUnorm8x2:
		return "unorm8x2"
	case VertexFormatUnorm8x4:
		return "unorm8x4"
	case VertexFormatSnorm8:
		return "snorm8"
	case VertexFormatSnorm8x2:
		return "snorm8x2"
	case VertexFormatSnorm8x4:
		return "snorm8x4"
	case VertexFormatUint16:
		return "uint16"
	case VertexFormatUint16x2:
		return "uint16x2"
	case VertexFormatUint16x4:
		return "uint16x4"
	case VertexFormatSint16:
		return "sint16"
	case VertexFormatSint16x2:
		return "sint16x2"
	case VertexFormatSint16x4:
		return "sint16x4"
	case VertexFormatUnorm16:
		return "unorm16"
	case VertexFormatUnorm16x2:
		return "unorm16x2"
	case VertexFormatUnorm16x4:
		return "unorm16x4"
	case VertexFormatSnorm16:
		return "snorm16"
	case VertexFormatSnorm16x2:
		return "snorm16x2"
	case VertexFormatSnorm16x4:
		return "snorm16x4"
	case VertexFormatFloat16:
		return "float16"
	case VertexFormatFloat16x2:
		return "float16x2"
	case VertexFormatFloat16x4:
		return "float16x4"
	case VertexFormatFloat32:
		return "float32"
	case VertexFormatFloat32x2:
		return "float32x2"
	case VertexFormatFloat32x3:
		return "float32x3"
	case VertexFormatFloat32x4:
		return "float32x4"
	case VertexFormatUint32:
		return "uint32"
	case VertexFormatUint32x2:
		return "uint32x2"
	case VertexFormatUint32x3:
		return "uint32x3"
	case VertexFormatUint32x4:
		return "uint32x4"
	case VertexFormatSint32:
		return "sint32"
	case VertexFormatSint32x2:
		return "sint32x2"
	case VertexFormatSint32x3:
		return "sint32x3"
	case VertexFormatSint32x4:
		return "sint32x4"
	case VertexFormatUnorm10_10_10_2:
		return "unorm10-10-10-2"
	case VertexFormatUnorm8x4BGRA:
		return "unorm8x4-bgra"
	default:
		return ""
	}
}

func (t PrimitiveTopology) toJS() string {
	switch t {
	case PrimitiveTopologyPointList:
		return "point-list"
	case PrimitiveTopologyLineList:
		return "line-list"
	case PrimitiveTopologyLineStrip:
		return "line-strip"
	case PrimitiveTopologyTriangleList:
		return "triangle-list"
	case PrimitiveTopologyTriangleStrip:
		return "triangle-strip"
	default:
		return "triangle-list"
	}
}

func (f IndexFormat) toJS() string {
	switch f {
	case IndexFormatUint16:
		return "uint16"
	case IndexFormatUint32:
		return "uint32"
	default:
		return "undefined"
	}
}

func (f FrontFace) toJS() string {
	switch f {
	case FrontFaceCCW:
		return "ccw"
	case FrontFaceCW:
		return "cw"
	default:
		return "ccw"
	}
}

func (c CullMode) toJS() string {
	switch c {
	case CullModeNone:
		return "none"
	case CullModeFront:
		return "front"
	case CullModeBack:
		return "back"
	default:
		return "none"
	}
}

func (op BlendOperation) toJS() string {
	switch op {
	case BlendOperationAdd:
		return "add"
	case BlendOperationSubtract:
		return "subtract"
	case BlendOperationReverseSubtract:
		return "reverse-subtract"
	case BlendOperationMin:
		return "min"
	case BlendOperationMax:
		return "max"
	default:
		return "add"
	}
}

func (f BlendFactor) toJS() string {
	switch f {
	case BlendFactorZero:
		return "zero"
	case BlendFactorOne:
		return "one"
	case BlendFactorSrc:
		return "src"
	case BlendFactorOneMinusSrc:
		return "one-minus-src"
	case BlendFactorSrcAlpha:
		return "src-alpha"
	case BlendFactorOneMinusSrcAlpha:
		return "one-minus-src-alpha"
	case BlendFactorDst:
		return "dst"
	case BlendFactorOneMinusDst:
		return "one-minus-dst"
	case BlendFactorDstAlpha:
		return "dst-alpha"
	case BlendFactorOneMinusDstAlpha:
		return "one-minus-dst-alpha"
	case BlendFactorSrcAlphaSaturated:
		return "src-alpha-saturated"
	case BlendFactorConstant:
		return "constant"
	case BlendFactorOneMinusConstant:
		return "one-minus-constant"
	case BlendFactorSrc1:
		return "src1"
	case BlendFactorOneMinusSrc1:
		return "one-minus-src1"
	case BlendFactorSrc1Alpha:
		return "src1-alpha"
	case BlendFactorOneMinusSrc1Alpha:
		return "one-minus-src1-alpha"
	default:
		return "one"
	}
}

func (f CompareFunction) toJS() string {
	switch f {
	case CompareFunctionNever:
		return "never"
	case CompareFunctionLess:
		return "less"
	case CompareFunctionEqual:
		return "equal"
	case CompareFunctionLessEqual:
		return "less-equal"
	case CompareFunctionGreater:
		return "greater"
	case CompareFunctionNotEqual:
		return "not-equal"
	case CompareFunctionGreaterEqual:
		return "greater-equal"
	case CompareFunctionAlways:
		return "always"
	default:
		return "always"
	}
}

func (op StencilOperation) toJS() string {
	switch op {
	case StencilOperationKeep:
		return "keep"
	case StencilOperationZero:
		return "zero"
	case StencilOperationReplace:
		return "replace"
	case StencilOperationInvert:
		return "invert"
	case StencilOperationIncrementClamp:
		return "increment-clamp"
	case StencilOperationDecrementClamp:
		return "decrement-clamp"
	case StencilOperationIncrementWrap:
		return "increment-wrap"
	case StencilOperationDecrementWrap:
		return "decrement-wrap"
	default:
		return "keep"
	}
}

func (f FilterMode) toJS() string {
	switch f {
	case FilterModeNearest:
		return "nearest"
	case FilterModeLinear:
		return "linear"
	default:
		return "nearest"
	}
}

func (f MipmapFilterMode) toJS() string {
	switch f {
	case MipmapFilterModeNearest:
		return "nearest"
	case MipmapFilterModeLinear:
		return "linear"
	default:
		return "nearest"
	}
}

func (m AddressMode) toJS() string {
	switch m {
	case AddressModeClampToEdge:
		return "clamp-to-edge"
	case AddressModeRepeat:
		return "repeat"
	case AddressModeMirrorRepeat:
		return "mirror-repeat"
	default:
		return "clamp-to-edge"
	}
}

func (op LoadOp) toJS() string {
	switch op {
	case LoadOpLoad:
		return "load"
	case LoadOpClear:
		return "clear"
	default:
		return "load"
	}
}

func (op StoreOp) toJS() string {
	switch op {
	case StoreOpStore:
		return "store"
	case StoreOpDiscard:
		return "discard"
	default:
		return "store"
	}
}

func (d TextureDimension) toJS() string {
	switch d {
	case TextureDimension1D:
		return "1d"
	case TextureDimension2D:
		return "2d"
	case TextureDimension3D:
		return "3d"
	default:
		return "2d"
	}
}

func (d TextureViewDimension) toJS() string {
	switch d {
	case TextureViewDimension1D:
		return "1d"
	case TextureViewDimension2D:
		return "2d"
	case TextureViewDimension2DArray:
		return "2d-array"
	case TextureViewDimensionCube:
		return "cube"
	case TextureViewDimensionCubeArray:
		return "cube-array"
	case TextureViewDimension3D:
		return "3d"
	default:
		return "2d"
	}
}

func (a TextureAspect) toJS() string {
	switch a {
	case TextureAspectAll:
		return "all"
	case TextureAspectStencilOnly:
		return "stencil-only"
	case TextureAspectDepthOnly:
		return "depth-only"
	default:
		return "all"
	}
}

func (m VertexStepMode) toJS() string {
	switch m {
	case VertexStepModeVertex:
		return "vertex"
	case VertexStepModeInstance:
		return "instance"
	default:
		return "vertex"
	}
}

func (m PresentMode) toJS() string {
	switch m {
	case PresentModeFifo:
		return "fifo"
	case PresentModeFifoRelaxed:
		return "fifo-relaxed"
	case PresentModeImmediate:
		return "immediate"
	case PresentModeMailbox:
		return "mailbox"
	default:
		return "fifo"
	}
}

func (m CompositeAlphaMode) toJS() string {
	switch m {
	case CompositeAlphaModeOpaque:
		return "opaque"
	case CompositeAlphaModePremultiplied:
		return "premultiplied"
	default:
		// Auto, Unpremultiplied, and Inherit have no browser WebGPU equivalent;
		// return "" so callers can omit the field and let the browser default to opaque.
		return ""
	}
}

func (t QueryType) toJS() string {
	switch t {
	case QueryTypeOcclusion:
		return "occlusion"
	case QueryTypeTimestamp:
		return "timestamp"
	default:
		return "occlusion"
	}
}

func (p PowerPreference) toJS() string {
	switch p {
	case PowerPreferenceLowPower:
		return "low-power"
	case PowerPreferenceHighPerformance:
		return "high-performance"
	default:
		return ""
	}
}

func (f ErrorFilter) toJS() string {
	switch f {
	case ErrorFilterValidation:
		return "validation"
	case ErrorFilterOutOfMemory:
		return "out-of-memory"
	case ErrorFilterInternal:
		return "internal"
	default:
		return "validation"
	}
}

func (t BufferBindingType) toJS() string {
	switch t {
	case BufferBindingTypeUniform:
		return "uniform"
	case BufferBindingTypeStorage:
		return "storage"
	case BufferBindingTypeReadOnlyStorage:
		return "read-only-storage"
	default:
		return ""
	}
}

func (t SamplerBindingType) toJS() string {
	switch t {
	case SamplerBindingTypeFiltering:
		return "filtering"
	case SamplerBindingTypeNonFiltering:
		return "non-filtering"
	case SamplerBindingTypeComparison:
		return "comparison"
	default:
		return ""
	}
}

func (t TextureSampleType) toJS() string {
	switch t {
	case TextureSampleTypeFloat:
		return "float"
	case TextureSampleTypeUnfilterableFloat:
		return "unfilterable-float"
	case TextureSampleTypeDepth:
		return "depth"
	case TextureSampleTypeSint:
		return "sint"
	case TextureSampleTypeUint:
		return "uint"
	default:
		return ""
	}
}

func (a StorageTextureAccess) toJS() string {
	switch a {
	case StorageTextureAccessWriteOnly:
		return "write-only"
	case StorageTextureAccessReadOnly:
		return "read-only"
	case StorageTextureAccessReadWrite:
		return "read-write"
	default:
		return ""
	}
}

func (b OptionalBool) toJS() any {
	switch b {
	case OptionalBoolTrue:
		return true
	case OptionalBoolFalse:
		return false
	default:
		return js.Undefined()
	}
}

func (f FeatureName) toJS() string {
	switch f {
	case FeatureNameCoreFeaturesAndLimits:
		return "core-features-and-limits"
	case FeatureNameDepthClipControl:
		return "depth-clip-control"
	case FeatureNameDepth32FloatStencil8:
		return "depth32float-stencil8"
	case FeatureNameTextureCompressionBC:
		return "texture-compression-bc"
	case FeatureNameTextureCompressionBCSliced3D:
		return "texture-compression-bc-sliced-3d"
	case FeatureNameTextureCompressionETC2:
		return "texture-compression-etc2"
	case FeatureNameTextureCompressionASTC:
		return "texture-compression-astc"
	case FeatureNameTextureCompressionASTCSliced3D:
		return "texture-compression-astc-sliced-3d"
	case FeatureNameTimestampQuery:
		return "timestamp-query"
	case FeatureNameIndirectFirstInstance:
		return "indirect-first-instance"
	case FeatureNameShaderF16:
		return "shader-f16"
	case FeatureNameRG11B10UfloatRenderable:
		return "rg11b10ufloat-renderable"
	case FeatureNameBGRA8UnormStorage:
		return "bgra8unorm-storage"
	case FeatureNameFloat32Filterable:
		return "float32-filterable"
	case FeatureNameFloat32Blendable:
		return "float32-blendable"
	case FeatureNameClipDistances:
		return "clip-distances"
	case FeatureNameDualSourceBlending:
		return "dual-source-blending"
	case FeatureNameSubgroups:
		return "subgroups"
	case FeatureNameTextureFormatsTier1:
		return "texture-formats-tier-1"
	case FeatureNameTextureFormatsTier2:
		return "texture-formats-tier-2"
	case FeatureNamePrimitiveIndex:
		return "primitive-index"
	case FeatureNameTextureComponentSwizzle:
		return "texture-component-swizzle"
	default:
		return ""
	}
}

func featureNameFromJS(s string) (FeatureName, bool) {
	switch s {
	case "core-features-and-limits":
		return FeatureNameCoreFeaturesAndLimits, true
	case "depth-clip-control":
		return FeatureNameDepthClipControl, true
	case "depth32float-stencil8":
		return FeatureNameDepth32FloatStencil8, true
	case "texture-compression-bc":
		return FeatureNameTextureCompressionBC, true
	case "texture-compression-bc-sliced-3d":
		return FeatureNameTextureCompressionBCSliced3D, true
	case "texture-compression-etc2":
		return FeatureNameTextureCompressionETC2, true
	case "texture-compression-astc":
		return FeatureNameTextureCompressionASTC, true
	case "texture-compression-astc-sliced-3d":
		return FeatureNameTextureCompressionASTCSliced3D, true
	case "timestamp-query":
		return FeatureNameTimestampQuery, true
	case "indirect-first-instance":
		return FeatureNameIndirectFirstInstance, true
	case "shader-f16":
		return FeatureNameShaderF16, true
	case "rg11b10ufloat-renderable":
		return FeatureNameRG11B10UfloatRenderable, true
	case "bgra8unorm-storage":
		return FeatureNameBGRA8UnormStorage, true
	case "float32-filterable":
		return FeatureNameFloat32Filterable, true
	case "float32-blendable":
		return FeatureNameFloat32Blendable, true
	case "clip-distances":
		return FeatureNameClipDistances, true
	case "dual-source-blending":
		return FeatureNameDualSourceBlending, true
	case "subgroups":
		return FeatureNameSubgroups, true
	case "texture-formats-tier-1":
		return FeatureNameTextureFormatsTier1, true
	case "texture-formats-tier-2":
		return FeatureNameTextureFormatsTier2, true
	case "primitive-index":
		return FeatureNamePrimitiveIndex, true
	case "texture-component-swizzle":
		return FeatureNameTextureComponentSwizzle, true
	default:
		return 0, false
	}
}

func ErrorTypeFromJS(jsErr js.Value) (ErrorType, string) {
	if jsErr.IsNull() || jsErr.IsUndefined() {
		return ErrorTypeNoError, ""
	}
	typ := jsErr.Get("constructor").Get("name").String()
	msg := jsErr.Get("message").String()
	switch typ {
	case "GPUValidationError":
		return ErrorTypeValidation, msg
	case "GPUOutOfMemoryError":
		return ErrorTypeOutOfMemory, msg
	case "GPUInternalError":
		return ErrorTypeInternal, msg
	default:
		return ErrorTypeUnknown, msg
	}
}
