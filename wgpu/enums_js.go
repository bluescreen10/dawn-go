//go:build js

package wgpu

import "syscall/js"

func (f TextureFormat) toJS() js.Value {
	switch f {
	case TextureFormatR8Unorm:
		return js.ValueOf("r8unorm")
	case TextureFormatR8Snorm:
		return js.ValueOf("r8snorm")
	case TextureFormatR8Uint:
		return js.ValueOf("r8uint")
	case TextureFormatR8Sint:
		return js.ValueOf("r8sint")
	case TextureFormatR16Unorm:
		return js.ValueOf("r16unorm")
	case TextureFormatR16Snorm:
		return js.ValueOf("r16snorm")
	case TextureFormatR16Uint:
		return js.ValueOf("r16uint")
	case TextureFormatR16Sint:
		return js.ValueOf("r16sint")
	case TextureFormatR16Float:
		return js.ValueOf("r16float")
	case TextureFormatRG8Unorm:
		return js.ValueOf("rg8unorm")
	case TextureFormatRG8Snorm:
		return js.ValueOf("rg8snorm")
	case TextureFormatRG8Uint:
		return js.ValueOf("rg8uint")
	case TextureFormatRG8Sint:
		return js.ValueOf("rg8sint")
	case TextureFormatR32Float:
		return js.ValueOf("r32float")
	case TextureFormatR32Uint:
		return js.ValueOf("r32uint")
	case TextureFormatR32Sint:
		return js.ValueOf("r32sint")
	case TextureFormatRG16Unorm:
		return js.ValueOf("rg16unorm")
	case TextureFormatRG16Snorm:
		return js.ValueOf("rg16snorm")
	case TextureFormatRG16Uint:
		return js.ValueOf("rg16uint")
	case TextureFormatRG16Sint:
		return js.ValueOf("rg16sint")
	case TextureFormatRG16Float:
		return js.ValueOf("rg16float")
	case TextureFormatRGBA8Unorm:
		return js.ValueOf("rgba8unorm")
	case TextureFormatRGBA8UnormSRGB:
		return js.ValueOf("rgba8unorm-srgb")
	case TextureFormatRGBA8Snorm:
		return js.ValueOf("rgba8snorm")
	case TextureFormatRGBA8Uint:
		return js.ValueOf("rgba8uint")
	case TextureFormatRGBA8Sint:
		return js.ValueOf("rgba8sint")
	case TextureFormatBGRA8Unorm:
		return js.ValueOf("bgra8unorm")
	case TextureFormatBGRA8UnormSRGB:
		return js.ValueOf("bgra8unorm-srgb")
	case TextureFormatRGB10A2Uint:
		return js.ValueOf("rgb10a2uint")
	case TextureFormatRGB10A2Unorm:
		return js.ValueOf("rgb10a2unorm")
	case TextureFormatRG11B10Ufloat:
		return js.ValueOf("rg11b10ufloat")
	case TextureFormatRGB9E5Ufloat:
		return js.ValueOf("rgb9e5ufloat")
	case TextureFormatRG32Float:
		return js.ValueOf("rg32float")
	case TextureFormatRG32Uint:
		return js.ValueOf("rg32uint")
	case TextureFormatRG32Sint:
		return js.ValueOf("rg32sint")
	case TextureFormatRGBA16Unorm:
		return js.ValueOf("rgba16unorm")
	case TextureFormatRGBA16Snorm:
		return js.ValueOf("rgba16snorm")
	case TextureFormatRGBA16Uint:
		return js.ValueOf("rgba16uint")
	case TextureFormatRGBA16Sint:
		return js.ValueOf("rgba16sint")
	case TextureFormatRGBA16Float:
		return js.ValueOf("rgba16float")
	case TextureFormatRGBA32Float:
		return js.ValueOf("rgba32float")
	case TextureFormatRGBA32Uint:
		return js.ValueOf("rgba32uint")
	case TextureFormatRGBA32Sint:
		return js.ValueOf("rgba32sint")
	case TextureFormatStencil8:
		return js.ValueOf("stencil8")
	case TextureFormatDepth16Unorm:
		return js.ValueOf("depth16unorm")
	case TextureFormatDepth24Plus:
		return js.ValueOf("depth24plus")
	case TextureFormatDepth24PlusStencil8:
		return js.ValueOf("depth24plus-stencil8")
	case TextureFormatDepth32Float:
		return js.ValueOf("depth32float")
	case TextureFormatDepth32FloatStencil8:
		return js.ValueOf("depth32float-stencil8")
	case TextureFormatBC1RGBAUnorm:
		return js.ValueOf("bc1-rgba-unorm")
	case TextureFormatBC1RGBAUnormSRGB:
		return js.ValueOf("bc1-rgba-unorm-srgb")
	case TextureFormatBC2RGBAUnorm:
		return js.ValueOf("bc2-rgba-unorm")
	case TextureFormatBC2RGBAUnormSRGB:
		return js.ValueOf("bc2-rgba-unorm-srgb")
	case TextureFormatBC3RGBAUnorm:
		return js.ValueOf("bc3-rgba-unorm")
	case TextureFormatBC3RGBAUnormSRGB:
		return js.ValueOf("bc3-rgba-unorm-srgb")
	case TextureFormatBC4RUnorm:
		return js.ValueOf("bc4-r-unorm")
	case TextureFormatBC4RSnorm:
		return js.ValueOf("bc4-r-snorm")
	case TextureFormatBC5RGUnorm:
		return js.ValueOf("bc5-rg-unorm")
	case TextureFormatBC5RGSnorm:
		return js.ValueOf("bc5-rg-snorm")
	case TextureFormatBC6HRGBUfloat:
		return js.ValueOf("bc6h-rgb-ufloat")
	case TextureFormatBC6HRGBFloat:
		return js.ValueOf("bc6h-rgb-float")
	case TextureFormatBC7RGBAUnorm:
		return js.ValueOf("bc7-rgba-unorm")
	case TextureFormatBC7RGBAUnormSRGB:
		return js.ValueOf("bc7-rgba-unorm-srgb")
	case TextureFormatETC2RGB8Unorm:
		return js.ValueOf("etc2-rgb8unorm")
	case TextureFormatETC2RGB8UnormSRGB:
		return js.ValueOf("etc2-rgb8unorm-srgb")
	case TextureFormatETC2RGB8A1Unorm:
		return js.ValueOf("etc2-rgb8a1unorm")
	case TextureFormatETC2RGB8A1UnormSRGB:
		return js.ValueOf("etc2-rgb8a1unorm-srgb")
	case TextureFormatETC2RGBA8Unorm:
		return js.ValueOf("etc2-rgba8unorm")
	case TextureFormatETC2RGBA8UnormSRGB:
		return js.ValueOf("etc2-rgba8unorm-srgb")
	case TextureFormatEACR11Unorm:
		return js.ValueOf("eac-r11unorm")
	case TextureFormatEACR11Snorm:
		return js.ValueOf("eac-r11snorm")
	case TextureFormatEACRG11Unorm:
		return js.ValueOf("eac-rg11unorm")
	case TextureFormatEACRG11Snorm:
		return js.ValueOf("eac-rg11snorm")
	case TextureFormatASTC4x4Unorm:
		return js.ValueOf("astc-4x4-unorm")
	case TextureFormatASTC4x4UnormSRGB:
		return js.ValueOf("astc-4x4-unorm-srgb")
	case TextureFormatASTC5x4Unorm:
		return js.ValueOf("astc-5x4-unorm")
	case TextureFormatASTC5x4UnormSRGB:
		return js.ValueOf("astc-5x4-unorm-srgb")
	case TextureFormatASTC5x5Unorm:
		return js.ValueOf("astc-5x5-unorm")
	case TextureFormatASTC5x5UnormSRGB:
		return js.ValueOf("astc-5x5-unorm-srgb")
	case TextureFormatASTC6x5Unorm:
		return js.ValueOf("astc-6x5-unorm")
	case TextureFormatASTC6x5UnormSRGB:
		return js.ValueOf("astc-6x5-unorm-srgb")
	case TextureFormatASTC6x6Unorm:
		return js.ValueOf("astc-6x6-unorm")
	case TextureFormatASTC6x6UnormSRGB:
		return js.ValueOf("astc-6x6-unorm-srgb")
	case TextureFormatASTC8x5Unorm:
		return js.ValueOf("astc-8x5-unorm")
	case TextureFormatASTC8x5UnormSRGB:
		return js.ValueOf("astc-8x5-unorm-srgb")
	case TextureFormatASTC8x6Unorm:
		return js.ValueOf("astc-8x6-unorm")
	case TextureFormatASTC8x6UnormSRGB:
		return js.ValueOf("astc-8x6-unorm-srgb")
	case TextureFormatASTC8x8Unorm:
		return js.ValueOf("astc-8x8-unorm")
	case TextureFormatASTC8x8UnormSRGB:
		return js.ValueOf("astc-8x8-unorm-srgb")
	case TextureFormatASTC10x5Unorm:
		return js.ValueOf("astc-10x5-unorm")
	case TextureFormatASTC10x5UnormSRGB:
		return js.ValueOf("astc-10x5-unorm-srgb")
	case TextureFormatASTC10x6Unorm:
		return js.ValueOf("astc-10x6-unorm")
	case TextureFormatASTC10x6UnormSRGB:
		return js.ValueOf("astc-10x6-unorm-srgb")
	case TextureFormatASTC10x8Unorm:
		return js.ValueOf("astc-10x8-unorm")
	case TextureFormatASTC10x8UnormSRGB:
		return js.ValueOf("astc-10x8-unorm-srgb")
	case TextureFormatASTC10x10Unorm:
		return js.ValueOf("astc-10x10-unorm")
	case TextureFormatASTC10x10UnormSRGB:
		return js.ValueOf("astc-10x10-unorm-srgb")
	case TextureFormatASTC12x10Unorm:
		return js.ValueOf("astc-12x10-unorm")
	case TextureFormatASTC12x10UnormSRGB:
		return js.ValueOf("astc-12x10-unorm-srgb")
	case TextureFormatASTC12x12Unorm:
		return js.ValueOf("astc-12x12-unorm")
	case TextureFormatASTC12x12UnormSRGB:
		return js.ValueOf("astc-12x12-unorm-srgb")
	default:
		return js.Undefined()
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

func (f VertexFormat) toJS() js.Value {
	switch f {
	case VertexFormatUint8:
		return js.ValueOf("uint8")
	case VertexFormatUint8x2:
		return js.ValueOf("uint8x2")
	case VertexFormatUint8x4:
		return js.ValueOf("uint8x4")
	case VertexFormatSint8:
		return js.ValueOf("sint8")
	case VertexFormatSint8x2:
		return js.ValueOf("sint8x2")
	case VertexFormatSint8x4:
		return js.ValueOf("sint8x4")
	case VertexFormatUnorm8:
		return js.ValueOf("unorm8")
	case VertexFormatUnorm8x2:
		return js.ValueOf("unorm8x2")
	case VertexFormatUnorm8x4:
		return js.ValueOf("unorm8x4")
	case VertexFormatSnorm8:
		return js.ValueOf("snorm8")
	case VertexFormatSnorm8x2:
		return js.ValueOf("snorm8x2")
	case VertexFormatSnorm8x4:
		return js.ValueOf("snorm8x4")
	case VertexFormatUint16:
		return js.ValueOf("uint16")
	case VertexFormatUint16x2:
		return js.ValueOf("uint16x2")
	case VertexFormatUint16x4:
		return js.ValueOf("uint16x4")
	case VertexFormatSint16:
		return js.ValueOf("sint16")
	case VertexFormatSint16x2:
		return js.ValueOf("sint16x2")
	case VertexFormatSint16x4:
		return js.ValueOf("sint16x4")
	case VertexFormatUnorm16:
		return js.ValueOf("unorm16")
	case VertexFormatUnorm16x2:
		return js.ValueOf("unorm16x2")
	case VertexFormatUnorm16x4:
		return js.ValueOf("unorm16x4")
	case VertexFormatSnorm16:
		return js.ValueOf("snorm16")
	case VertexFormatSnorm16x2:
		return js.ValueOf("snorm16x2")
	case VertexFormatSnorm16x4:
		return js.ValueOf("snorm16x4")
	case VertexFormatFloat16:
		return js.ValueOf("float16")
	case VertexFormatFloat16x2:
		return js.ValueOf("float16x2")
	case VertexFormatFloat16x4:
		return js.ValueOf("float16x4")
	case VertexFormatFloat32:
		return js.ValueOf("float32")
	case VertexFormatFloat32x2:
		return js.ValueOf("float32x2")
	case VertexFormatFloat32x3:
		return js.ValueOf("float32x3")
	case VertexFormatFloat32x4:
		return js.ValueOf("float32x4")
	case VertexFormatUint32:
		return js.ValueOf("uint32")
	case VertexFormatUint32x2:
		return js.ValueOf("uint32x2")
	case VertexFormatUint32x3:
		return js.ValueOf("uint32x3")
	case VertexFormatUint32x4:
		return js.ValueOf("uint32x4")
	case VertexFormatSint32:
		return js.ValueOf("sint32")
	case VertexFormatSint32x2:
		return js.ValueOf("sint32x2")
	case VertexFormatSint32x3:
		return js.ValueOf("sint32x3")
	case VertexFormatSint32x4:
		return js.ValueOf("sint32x4")
	case VertexFormatUnorm10_10_10_2:
		return js.ValueOf("unorm10-10-10-2")
	case VertexFormatUnorm8x4BGRA:
		return js.ValueOf("unorm8x4-bgra")
	default:
		return js.Undefined()
	}
}

func (t PrimitiveTopology) toJS() js.Value {
	switch t {
	case PrimitiveTopologyPointList:
		return js.ValueOf("point-list")
	case PrimitiveTopologyLineList:
		return js.ValueOf("line-list")
	case PrimitiveTopologyLineStrip:
		return js.ValueOf("line-strip")
	case PrimitiveTopologyTriangleList:
		return js.ValueOf("triangle-list")
	case PrimitiveTopologyTriangleStrip:
		return js.ValueOf("triangle-strip")
	default:
		return js.ValueOf("triangle-list")
	}
}

func (f IndexFormat) toJS() js.Value {
	switch f {
	case IndexFormatUint16:
		return js.ValueOf("uint16")
	case IndexFormatUint32:
		return js.ValueOf("uint32")
	default:
		return js.Undefined()
	}
}

func (f FrontFace) toJS() js.Value {
	switch f {
	case FrontFaceCCW:
		return js.ValueOf("ccw")
	case FrontFaceCW:
		return js.ValueOf("cw")
	default:
		return js.ValueOf("ccw")
	}
}

func (c CullMode) toJS() js.Value {
	switch c {
	case CullModeNone:
		return js.ValueOf("none")
	case CullModeFront:
		return js.ValueOf("front")
	case CullModeBack:
		return js.ValueOf("back")
	default:
		return js.ValueOf("none")
	}
}

func (op BlendOperation) toJS() js.Value {
	switch op {
	case BlendOperationAdd:
		return js.ValueOf("add")
	case BlendOperationSubtract:
		return js.ValueOf("subtract")
	case BlendOperationReverseSubtract:
		return js.ValueOf("reverse-subtract")
	case BlendOperationMin:
		return js.ValueOf("min")
	case BlendOperationMax:
		return js.ValueOf("max")
	default:
		return js.ValueOf("add")
	}
}

func (f BlendFactor) toJS() js.Value {
	switch f {
	case BlendFactorZero:
		return js.ValueOf("zero")
	case BlendFactorOne:
		return js.ValueOf("one")
	case BlendFactorSrc:
		return js.ValueOf("src")
	case BlendFactorOneMinusSrc:
		return js.ValueOf("one-minus-src")
	case BlendFactorSrcAlpha:
		return js.ValueOf("src-alpha")
	case BlendFactorOneMinusSrcAlpha:
		return js.ValueOf("one-minus-src-alpha")
	case BlendFactorDst:
		return js.ValueOf("dst")
	case BlendFactorOneMinusDst:
		return js.ValueOf("one-minus-dst")
	case BlendFactorDstAlpha:
		return js.ValueOf("dst-alpha")
	case BlendFactorOneMinusDstAlpha:
		return js.ValueOf("one-minus-dst-alpha")
	case BlendFactorSrcAlphaSaturated:
		return js.ValueOf("src-alpha-saturated")
	case BlendFactorConstant:
		return js.ValueOf("constant")
	case BlendFactorOneMinusConstant:
		return js.ValueOf("one-minus-constant")
	case BlendFactorSrc1:
		return js.ValueOf("src1")
	case BlendFactorOneMinusSrc1:
		return js.ValueOf("one-minus-src1")
	case BlendFactorSrc1Alpha:
		return js.ValueOf("src1-alpha")
	case BlendFactorOneMinusSrc1Alpha:
		return js.ValueOf("one-minus-src1-alpha")
	default:
		return js.ValueOf("one")
	}
}

func (f CompareFunction) toJS() js.Value {
	switch f {
	case CompareFunctionNever:
		return js.ValueOf("never")
	case CompareFunctionLess:
		return js.ValueOf("less")
	case CompareFunctionEqual:
		return js.ValueOf("equal")
	case CompareFunctionLessEqual:
		return js.ValueOf("less-equal")
	case CompareFunctionGreater:
		return js.ValueOf("greater")
	case CompareFunctionNotEqual:
		return js.ValueOf("not-equal")
	case CompareFunctionGreaterEqual:
		return js.ValueOf("greater-equal")
	case CompareFunctionAlways:
		return js.ValueOf("always")
	default:
		return js.Undefined()
	}
}

func (op StencilOperation) toJS() js.Value {
	switch op {
	case StencilOperationKeep:
		return js.ValueOf("keep")
	case StencilOperationZero:
		return js.ValueOf("zero")
	case StencilOperationReplace:
		return js.ValueOf("replace")
	case StencilOperationInvert:
		return js.ValueOf("invert")
	case StencilOperationIncrementClamp:
		return js.ValueOf("increment-clamp")
	case StencilOperationDecrementClamp:
		return js.ValueOf("decrement-clamp")
	case StencilOperationIncrementWrap:
		return js.ValueOf("increment-wrap")
	case StencilOperationDecrementWrap:
		return js.ValueOf("decrement-wrap")
	default:
		return js.ValueOf("keep")
	}
}

func (f FilterMode) toJS() js.Value {
	switch f {
	case FilterModeNearest:
		return js.ValueOf("nearest")
	case FilterModeLinear:
		return js.ValueOf("linear")
	default:
		return js.ValueOf("nearest")
	}
}

func (f MipmapFilterMode) toJS() js.Value {
	switch f {
	case MipmapFilterModeNearest:
		return js.ValueOf("nearest")
	case MipmapFilterModeLinear:
		return js.ValueOf("linear")
	default:
		return js.ValueOf("nearest")
	}
}

func (m AddressMode) toJS() js.Value {
	switch m {
	case AddressModeClampToEdge:
		return js.ValueOf("clamp-to-edge")
	case AddressModeRepeat:
		return js.ValueOf("repeat")
	case AddressModeMirrorRepeat:
		return js.ValueOf("mirror-repeat")
	default:
		return js.ValueOf("clamp-to-edge")
	}
}

func (op LoadOp) toJS() js.Value {
	switch op {
	case LoadOpLoad:
		return js.ValueOf("load")
	case LoadOpClear:
		return js.ValueOf("clear")
	default:
		return js.Undefined()
	}
}

func (op StoreOp) toJS() js.Value {
	switch op {
	case StoreOpStore:
		return js.ValueOf("store")
	case StoreOpDiscard:
		return js.ValueOf("discard")
	default:
		return js.Undefined()
	}
}

func (d TextureDimension) toJS() js.Value {
	switch d {
	case TextureDimension1D:
		return js.ValueOf("1d")
	case TextureDimension2D:
		return js.ValueOf("2d")
	case TextureDimension3D:
		return js.ValueOf("3d")
	default:
		return js.ValueOf("2d")
	}
}

func (d TextureViewDimension) toJS() js.Value {
	switch d {
	case TextureViewDimension1D:
		return js.ValueOf("1d")
	case TextureViewDimension2D:
		return js.ValueOf("2d")
	case TextureViewDimension2DArray:
		return js.ValueOf("2d-array")
	case TextureViewDimensionCube:
		return js.ValueOf("cube")
	case TextureViewDimensionCubeArray:
		return js.ValueOf("cube-array")
	case TextureViewDimension3D:
		return js.ValueOf("3d")
	default:
		return js.Undefined()
	}
}

func (a TextureAspect) toJS() js.Value {
	switch a {
	case TextureAspectAll:
		return js.ValueOf("all")
	case TextureAspectStencilOnly:
		return js.ValueOf("stencil-only")
	case TextureAspectDepthOnly:
		return js.ValueOf("depth-only")
	default:
		return js.ValueOf("all")
	}
}

func (m VertexStepMode) toJS() js.Value {
	switch m {
	case VertexStepModeVertex:
		return js.ValueOf("vertex")
	case VertexStepModeInstance:
		return js.ValueOf("instance")
	default:
		return js.ValueOf("vertex")
	}
}

func (m PresentMode) toJS() js.Value {
	switch m {
	case PresentModeFifo:
		return js.ValueOf("fifo")
	case PresentModeFifoRelaxed:
		return js.ValueOf("fifo-relaxed")
	case PresentModeImmediate:
		return js.ValueOf("immediate")
	case PresentModeMailbox:
		return js.ValueOf("mailbox")
	default:
		return js.ValueOf("fifo")
	}
}

func (m CompositeAlphaMode) toJS() js.Value {
	switch m {
	case CompositeAlphaModeOpaque:
		return js.ValueOf("opaque")
	case CompositeAlphaModePremultiplied:
		return js.ValueOf("premultiplied")
	default:
		return js.Undefined()
	}
}

func (t QueryType) toJS() js.Value {
	switch t {
	case QueryTypeOcclusion:
		return js.ValueOf("occlusion")
	case QueryTypeTimestamp:
		return js.ValueOf("timestamp")
	default:
		return js.Undefined()
	}
}

func (p PowerPreference) toJS() js.Value {
	switch p {
	case PowerPreferenceLowPower:
		return js.ValueOf("low-power")
	case PowerPreferenceHighPerformance:
		return js.ValueOf("high-performance")
	default:
		return js.Undefined()
	}
}

func (f ErrorFilter) toJS() js.Value {
	switch f {
	case ErrorFilterValidation:
		return js.ValueOf("validation")
	case ErrorFilterOutOfMemory:
		return js.ValueOf("out-of-memory")
	case ErrorFilterInternal:
		return js.ValueOf("internal")
	default:
		return js.ValueOf("validation")
	}
}

func (t BufferBindingType) toJS() js.Value {
	switch t {
	case BufferBindingTypeUniform:
		return js.ValueOf("uniform")
	case BufferBindingTypeStorage:
		return js.ValueOf("storage")
	case BufferBindingTypeReadOnlyStorage:
		return js.ValueOf("read-only-storage")
	default:
		return js.Undefined()
	}
}

func (t SamplerBindingType) toJS() js.Value {
	switch t {
	case SamplerBindingTypeFiltering:
		return js.ValueOf("filtering")
	case SamplerBindingTypeNonFiltering:
		return js.ValueOf("non-filtering")
	case SamplerBindingTypeComparison:
		return js.ValueOf("comparison")
	default:
		return js.Undefined()
	}
}

func (t TextureSampleType) toJS() js.Value {
	switch t {
	case TextureSampleTypeFloat:
		return js.ValueOf("float")
	case TextureSampleTypeUnfilterableFloat:
		return js.ValueOf("unfilterable-float")
	case TextureSampleTypeDepth:
		return js.ValueOf("depth")
	case TextureSampleTypeSint:
		return js.ValueOf("sint")
	case TextureSampleTypeUint:
		return js.ValueOf("uint")
	default:
		return js.Undefined()
	}
}

func (a StorageTextureAccess) toJS() js.Value {
	switch a {
	case StorageTextureAccessWriteOnly:
		return js.ValueOf("write-only")
	case StorageTextureAccessReadOnly:
		return js.ValueOf("read-only")
	case StorageTextureAccessReadWrite:
		return js.ValueOf("read-write")
	default:
		return js.Undefined()
	}
}

func (b OptionalBool) toJS() js.Value {
	switch b {
	case OptionalBoolTrue:
		return js.ValueOf(true)
	case OptionalBoolFalse:
		return js.ValueOf(false)
	default:
		return js.Undefined()
	}
}

func (f FeatureName) toJS() js.Value {
	switch f {
	case FeatureNameCoreFeaturesAndLimits:
		return js.ValueOf("core-features-and-limits")
	case FeatureNameDepthClipControl:
		return js.ValueOf("depth-clip-control")
	case FeatureNameDepth32FloatStencil8:
		return js.ValueOf("depth32float-stencil8")
	case FeatureNameTextureCompressionBC:
		return js.ValueOf("texture-compression-bc")
	case FeatureNameTextureCompressionBCSliced3D:
		return js.ValueOf("texture-compression-bc-sliced-3d")
	case FeatureNameTextureCompressionETC2:
		return js.ValueOf("texture-compression-etc2")
	case FeatureNameTextureCompressionASTC:
		return js.ValueOf("texture-compression-astc")
	case FeatureNameTextureCompressionASTCSliced3D:
		return js.ValueOf("texture-compression-astc-sliced-3d")
	case FeatureNameTimestampQuery:
		return js.ValueOf("timestamp-query")
	case FeatureNameIndirectFirstInstance:
		return js.ValueOf("indirect-first-instance")
	case FeatureNameShaderF16:
		return js.ValueOf("shader-f16")
	case FeatureNameRG11B10UfloatRenderable:
		return js.ValueOf("rg11b10ufloat-renderable")
	case FeatureNameBGRA8UnormStorage:
		return js.ValueOf("bgra8unorm-storage")
	case FeatureNameFloat32Filterable:
		return js.ValueOf("float32-filterable")
	case FeatureNameFloat32Blendable:
		return js.ValueOf("float32-blendable")
	case FeatureNameClipDistances:
		return js.ValueOf("clip-distances")
	case FeatureNameDualSourceBlending:
		return js.ValueOf("dual-source-blending")
	case FeatureNameSubgroups:
		return js.ValueOf("subgroups")
	case FeatureNameTextureFormatsTier1:
		return js.ValueOf("texture-formats-tier-1")
	case FeatureNameTextureFormatsTier2:
		return js.ValueOf("texture-formats-tier-2")
	case FeatureNamePrimitiveIndex:
		return js.ValueOf("primitive-index")
	case FeatureNameTextureComponentSwizzle:
		return js.ValueOf("texture-component-swizzle")
	default:
		return js.Undefined()
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
