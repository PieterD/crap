package gli

import (
	"fmt"

	"github.com/PieterD/glimmer/raw"
)

type iSyncResult struct {
	t raw.Enum
}

var (
	SyncAlreadySignaled    = iSyncResult{raw.SyncAlreadySignaled}
	SyncTimeoutExpired     = iSyncResult{raw.SyncTimeoutExpired}
	SyncConditionSatisfied = iSyncResult{raw.SyncConditionSatisfied}
	SyncWaitFailed         = iSyncResult{raw.SyncWaitFailed}
)

func (t iSyncResult) String() string {
	switch t {
	case SyncAlreadySignaled:
		return "SyncAlreadySignaled"
	case SyncTimeoutExpired:
		return "SyncTimeoutExpired"
	case SyncConditionSatisfied:
		return "SyncConditionSatisfied"
	case SyncWaitFailed:
		return "SyncWaitFailed"
	}
	return "Unknown SyncResult"
}

type iBindTarget struct {
	t raw.Enum
}

var (
	BindArrayBuffer             = iBindTarget{raw.BindTargetArrayBuffer}
	BindCopyReadBuffer          = iBindTarget{raw.BindTargetCopyReadBuffer}
	BindCopyWriteBuffer         = iBindTarget{raw.BindTargetCopyWriteBuffer}
	BindDrawIndirectBuffer      = iBindTarget{raw.BindTargetDrawIndirectBuffer}
	BindElementArrayBuffer      = iBindTarget{raw.BindTargetElementArrayBuffer}
	BindPixelPackBuffer         = iBindTarget{raw.BindTargetPixelPackBuffer}
	BindPixelUnpackBuffer       = iBindTarget{raw.BindTargetPixelUnpackBuffer}
	BindTextureBuffer           = iBindTarget{raw.BindTargetTextureBuffer}
	BindTransformFeedbackBuffer = iBindTarget{raw.BindTargetTransformFeedbackBuffer}
	BindUniformBuffer           = iBindTarget{raw.BindTargetUniformBuffer}
)

func (t iBindTarget) String() string {
	switch t {
	case BindArrayBuffer:
		return "BindArrayBuffer"
	case BindCopyReadBuffer:
		return "BindCopyReadBuffer"
	case BindCopyWriteBuffer:
		return "BindCopyWriteBuffer"
	case BindDrawIndirectBuffer:
		return "BindDrawIndirectBuffer"
	case BindElementArrayBuffer:
		return "BindElementArrayBuffer"
	case BindPixelPackBuffer:
		return "BindPixelPackBuffer"
	case BindPixelUnpackBuffer:
		return "BindPixelUnpackBuffer"
	case BindTextureBuffer:
		return "BindTextureBuffer"
	case BindTransformFeedbackBuffer:
		return "BindTransformFeedbackBuffer"
	case BindUniformBuffer:
		return "BindUniformBuffer"
	}
	return "Unknown BindTarget"
}

type iDrawMode struct {
	t raw.Enum
}

var (
	DrawPoints                 = iDrawMode{raw.DrawPoints}
	DrawLineStrip              = iDrawMode{raw.DrawLineStrip}
	DrawLineLoop               = iDrawMode{raw.DrawLineLoop}
	DrawLines                  = iDrawMode{raw.DrawLines}
	DrawLineStripAdjacency     = iDrawMode{raw.DrawLineStripAdjacency}
	DrawLinesAdjacency         = iDrawMode{raw.DrawLinesAdjacency}
	DrawTriangleStrip          = iDrawMode{raw.DrawTriangleStrip}
	DrawTriangleFan            = iDrawMode{raw.DrawTriangleFan}
	DrawTriangles              = iDrawMode{raw.DrawTriangles}
	DrawTriangleStripAdjacency = iDrawMode{raw.DrawTriangleStripAdjacency}
	DrawTrianglesAdjacency     = iDrawMode{raw.DrawTrianglesAdjacency}
	DrawPatches                = iDrawMode{raw.DrawPatches}
)

func (t iDrawMode) String() string {
	switch t {
	case DrawPoints:
		return "DrawPoints"
	case DrawLineStrip:
		return "DrawLineStrip"
	case DrawLineLoop:
		return "DrawLineLoop"
	case DrawLines:
		return "DrawLines"
	case DrawLineStripAdjacency:
		return "DrawLineStripAdjacency"
	case DrawLinesAdjacency:
		return "DrawLinesAdjacency"
	case DrawTriangleStrip:
		return "DrawTriangleStrip"
	case DrawTriangleFan:
		return "DrawTriangleFan"
	case DrawTriangles:
		return "DrawTriangles"
	case DrawTriangleStripAdjacency:
		return "DrawTriangleStripAdjacency"
	case DrawTrianglesAdjacency:
		return "DrawTrianglesAdjacency"
	case DrawPatches:
		return "DrawPatches"
	}
	return "Unknown DrawMode"
}

type iShaderType struct {
	t raw.Enum
}

var (
	VertexShader   = iShaderType{raw.ShaderTypeVertex}
	GeometryShader = iShaderType{raw.ShaderTypeGeometry}
	FragmentShader = iShaderType{raw.ShaderTypeFragment}
)

func (t iShaderType) String() string {
	switch t {
	case VertexShader:
		return "VertexShader"
	case GeometryShader:
		return "GeometryShader"
	case FragmentShader:
		return "FragmentShader"
	}
	return "Unknown Shader"
}

type iAccessType struct {
	t raw.Enum
}

var (
	StaticDraw  = iAccessType{raw.AccessTypeStaticDraw}
	StaticRead  = iAccessType{raw.AccessTypeStaticRead}
	StaticCopy  = iAccessType{raw.AccessTypeStaticCopy}
	StreamDraw  = iAccessType{raw.AccessTypeStreamDraw}
	StreamRead  = iAccessType{raw.AccessTypeStreamRead}
	StreamCopy  = iAccessType{raw.AccessTypeStreamCopy}
	DynamicDraw = iAccessType{raw.AccessTypeDynamicDraw}
	DynamicRead = iAccessType{raw.AccessTypeDynamicRead}
	DynamicCopy = iAccessType{raw.AccessTypeDynamicCopy}
)

func (t iAccessType) String() string {
	switch t {
	case StaticDraw:
		return "StaticDraw"
	case StaticRead:
		return "StaticRead"
	case StaticCopy:
		return "StaticCopy"
	case StreamDraw:
		return "StreamDraw"
	case StreamRead:
		return "StreamRead"
	case StreamCopy:
		return "StreamCopy"
	case DynamicDraw:
		return "DynamicDraw"
	case DynamicRead:
		return "DynamicRead"
	case DynamicCopy:
		return "DynamicCopy"
	}
	return "Unknown AccessType"
}

type iDataFormat struct {
	t raw.Enum
}

var (
	FmByte                = iDataFormat{raw.DataFormatByte}
	FmUByte               = iDataFormat{raw.DataFormatUnsignedByte}
	FmShort               = iDataFormat{raw.DataFormatShort}
	FmUShort              = iDataFormat{raw.DataFormatUnsignedShort}
	FmInt                 = iDataFormat{raw.DataFormatInt}
	FmUInt                = iDataFormat{raw.DataFormatUnsignedInt}
	FmHalfFloat           = iDataFormat{raw.DataFormatHalfFloat}
	FmFloat               = iDataFormat{raw.DataFormatFloat}
	FmDouble              = iDataFormat{raw.DataFormatDouble}
	FmFixed               = iDataFormat{raw.DataFormatFixed}
	FmIntRev_2_10_10_10   = iDataFormat{raw.DataFormatIntRev_2_10_10_10}
	FmUIntRev_2_10_10_10  = iDataFormat{raw.DataFormatUnsignedIntRev_2_10_10_10}
	FmUIntRev_10F_11F_11F = iDataFormat{raw.DataFormatUnsignedIntRev_10F_11F_11F}
)

type FullFormat struct {
	DataFormat iDataFormat
	Components int
}

func (f iDataFormat) Full(components int) FullFormat {
	if components < 1 {
		components = 1
	}
	if components > 4 {
		components = 4
	}
	return FullFormat{
		DataFormat: f,
		Components: components,
	}
}

func (f iDataFormat) Size() int {
	switch f {
	case FmByte, FmUByte:
		return 1
	case FmShort, FmUShort, FmHalfFloat:
		return 2
	case FmInt, FmUInt, FmFloat, FmFixed, FmIntRev_2_10_10_10, FmUIntRev_2_10_10_10, FmUIntRev_10F_11F_11F:
		return 4
	case FmDouble:
		return 8
	default:
		return -1
	}
}

func (ff FullFormat) String() string {
	if ff.Components <= 1 {
		return ff.DataFormat.String()
	}
	return fmt.Sprintf("%s[%d]", ff.DataFormat.String(), ff.Components)
}

func (t iDataFormat) String() string {
	switch t {
	case FmByte:
		return "FmByte"
	case FmUByte:
		return "FmUByte"
	case FmShort:
		return "FmShort"
	case FmUShort:
		return "FmUShort"
	case FmInt:
		return "FmInt"
	case FmUInt:
		return "FmUInt"
	case FmHalfFloat:
		return "FmHalfFloat"
	case FmFloat:
		return "FmFloat"
	case FmDouble:
		return "FmDouble"
	case FmFixed:
		return "FmFixed"
	case FmIntRev_2_10_10_10:
		return "FmIntRev_2_10_10_10"
	case FmUIntRev_2_10_10_10:
		return "FmUIntRev_2_10_10_10"
	case FmUIntRev_10F_11F_11F:
		return "FmUIntRev_10F_11F_11F"
	}
	return "Unknown DataFormat"
}

type iIndexType struct {
	t raw.Enum
}

var (
	IndexByte  = iIndexType{raw.IndexFormatByte}
	IndexShort = iIndexType{raw.IndexFormatShort}
	IndexInt   = iIndexType{raw.IndexFormatInt}
)

func (t iIndexType) String() string {
	switch t {
	case IndexByte:
		return "IndexByte"
	case IndexShort:
		return "IndexShort"
	case IndexInt:
		return "IndexInt"
	}
	return "Unknown IndexType"
}

type iDataType struct {
	t raw.Enum
}

// Valid for attributes and uniforms
var (
	Float        = iDataType{raw.DataTypeFloat}
	Float2       = iDataType{raw.DataTypeFloatVec2}
	Float3       = iDataType{raw.DataTypeFloatVec3}
	Float4       = iDataType{raw.DataTypeFloatVec4}
	FloatMat2    = iDataType{raw.DataTypeFloatMat2}
	FloatMat3    = iDataType{raw.DataTypeFloatMat3}
	FloatMat4    = iDataType{raw.DataTypeFloatMat4}
	FloatMat2x3  = iDataType{raw.DataTypeFloatMat2x3}
	FloatMat2x4  = iDataType{raw.DataTypeFloatMat2x4}
	FloatMat3x2  = iDataType{raw.DataTypeFloatMat3x2}
	FloatMat3x4  = iDataType{raw.DataTypeFloatMat3x4}
	FloatMat4x2  = iDataType{raw.DataTypeFloatMat4x2}
	FloatMat4x3  = iDataType{raw.DataTypeFloatMat4x3}
	Int          = iDataType{raw.DataTypeInt}
	Int2         = iDataType{raw.DataTypeIntVec2}
	Int3         = iDataType{raw.DataTypeIntVec3}
	Int4         = iDataType{raw.DataTypeIntVec4}
	UInt         = iDataType{raw.DataTypeUnsignedInt}
	UInt2        = iDataType{raw.DataTypeUnsignedIntVec2}
	UInt3        = iDataType{raw.DataTypeUnsignedIntVec3}
	UInt4        = iDataType{raw.DataTypeUnsignedIntVec4}
	Double       = iDataType{raw.DataTypeDouble}
	Double2      = iDataType{raw.DataTypeDoubleVec2}
	Double3      = iDataType{raw.DataTypeDoubleVec3}
	Double4      = iDataType{raw.DataTypeDoubleVec4}
	DoubleMat2   = iDataType{raw.DataTypeDoubleMat2}
	DoubleMat3   = iDataType{raw.DataTypeDoubleMat3}
	DoubleMat4   = iDataType{raw.DataTypeDoubleMat4}
	DoubleMat2x3 = iDataType{raw.DataTypeDoubleMat2x3}
	DoubleMat2x4 = iDataType{raw.DataTypeDoubleMat2x4}
	DoubleMat3x2 = iDataType{raw.DataTypeDoubleMat3x2}
	DoubleMat3x4 = iDataType{raw.DataTypeDoubleMat3x4}
	DoubleMat4x2 = iDataType{raw.DataTypeDoubleMat4x2}
	DoubleMat4x3 = iDataType{raw.DataTypeDoubleMat4x3}
)

func (dt iDataType) ValidAttribute() bool {
	return dt.t < raw.DataTypeAttributeLimit
}

// Valid only for uniforms
var (
	Bool                          = iDataType{raw.DataTypeBool}
	Bool2                         = iDataType{raw.DataTypeBool2}
	Bool3                         = iDataType{raw.DataTypeBool3}
	Bool4                         = iDataType{raw.DataTypeBool4}
	Sampler1d                     = iDataType{raw.DataTypeSampler1d}
	Sampler2d                     = iDataType{raw.DataTypeSampler2d}
	Sampler3d                     = iDataType{raw.DataTypeSampler3d}
	SamplerCube                   = iDataType{raw.DataTypeSamplerCube}
	SamplerCubeShadow             = iDataType{raw.DataTypeSamplerCubeShadow}
	Sampler1dShadow               = iDataType{raw.DataTypeSampler1dShadow}
	Sampler2dShadow               = iDataType{raw.DataTypeSampler2dShadow}
	Sampler1dArray                = iDataType{raw.DataTypeSampler1dArray}
	Sampler2dArray                = iDataType{raw.DataTypeSampler2dArray}
	Sampler1dArrayShadow          = iDataType{raw.DataTypeSampler1dArrayShadow}
	Sampler2dArrayShadow          = iDataType{raw.DataTypeSampler2dArrayShadow}
	Sampler2dMultisample          = iDataType{raw.DataTypeSampler2dMultisample}
	Sampler2dMultisampleArray     = iDataType{raw.DataTypeSampler2dMultisampleArray}
	SamplerBuffer                 = iDataType{raw.DataTypeSamplerBuffer}
	Sampler2dRect                 = iDataType{raw.DataTypeSampler2dRect}
	Sampler2dRectShadow           = iDataType{raw.DataTypeSampler2dRectShadow}
	IntSampler1d                  = iDataType{raw.DataTypeIntSampler1d}
	IntSampler2d                  = iDataType{raw.DataTypeIntSampler2d}
	IntSampler3d                  = iDataType{raw.DataTypeIntSampler3d}
	IntSamplerCube                = iDataType{raw.DataTypeIntSamplerCube}
	IntSampler1dArray             = iDataType{raw.DataTypeIntSampler1dArray}
	IntSampler2dArray             = iDataType{raw.DataTypeIntSampler2dArray}
	IntSampler2dMultisample       = iDataType{raw.DataTypeIntSampler2dMultisample}
	IntSampler2dMultisampleArray  = iDataType{raw.DataTypeIntSampler2dMultisampleArray}
	IntSamplerBuffer              = iDataType{raw.DataTypeIntSamplerBuffer}
	IntSampler2dRect              = iDataType{raw.DataTypeIntSampler2dRect}
	UIntSampler1d                 = iDataType{raw.DataTypeUIntSampler1d}
	UIntSampler2d                 = iDataType{raw.DataTypeUIntSampler2d}
	UIntSampler3d                 = iDataType{raw.DataTypeUIntSampler3d}
	UIntSamplerCube               = iDataType{raw.DataTypeUIntSamplerCube}
	UIntSampler1dArray            = iDataType{raw.DataTypeUIntSampler1dArray}
	UIntSampler2dArray            = iDataType{raw.DataTypeUIntSampler2dArray}
	UIntSampler2dMultisample      = iDataType{raw.DataTypeUIntSampler2dMultisample}
	UIntSampler2dMultisampleArray = iDataType{raw.DataTypeUIntSampler2dMultisampleArray}
	UIntSamplerBuffer             = iDataType{raw.DataTypeUIntSamplerBuffer}
	UIntSampler2dRect             = iDataType{raw.DataTypeUIntSampler2dRect}
	Image1d                       = iDataType{raw.DataTypeImage1d}
	Image2d                       = iDataType{raw.DataTypeImage2d}
	Image3d                       = iDataType{raw.DataTypeImage3d}
	ImageCube                     = iDataType{raw.DataTypeImageCube}
	Image1dArray                  = iDataType{raw.DataTypeImage1dArray}
	Image2dArray                  = iDataType{raw.DataTypeImage2dArray}
	Image2dMultisample            = iDataType{raw.DataTypeImage2dMultisample}
	Image2dMultisampleArray       = iDataType{raw.DataTypeImage2dMultisampleArray}
	ImageBuffer                   = iDataType{raw.DataTypeImageBuffer}
	Image2dRect                   = iDataType{raw.DataTypeImage2dRect}
	Int1d                         = iDataType{raw.DataTypeInt1d}
	Int2d                         = iDataType{raw.DataTypeInt2d}
	Int3d                         = iDataType{raw.DataTypeInt3d}
	IntCube                       = iDataType{raw.DataTypeIntCube}
	Int1dArray                    = iDataType{raw.DataTypeInt1dArray}
	Int2dArray                    = iDataType{raw.DataTypeInt2dArray}
	Int2dMultisample              = iDataType{raw.DataTypeInt2dMultisample}
	Int2dMultisampleArray         = iDataType{raw.DataTypeInt2dMultisampleArray}
	IntBuffer                     = iDataType{raw.DataTypeIntBuffer}
	Int2dRect                     = iDataType{raw.DataTypeInt2dRect}
	UInt1d                        = iDataType{raw.DataTypeUInt1d}
	UInt2d                        = iDataType{raw.DataTypeUInt2d}
	UInt3d                        = iDataType{raw.DataTypeUInt3d}
	UIntCube                      = iDataType{raw.DataTypeUIntCube}
	UInt1dArray                   = iDataType{raw.DataTypeUInt1dArray}
	UInt2dArray                   = iDataType{raw.DataTypeUInt2dArray}
	UInt2dMultisample             = iDataType{raw.DataTypeUInt2dMultisample}
	UInt2dMultisampleArray        = iDataType{raw.DataTypeUInt2dMultisampleArray}
	UIntBuffer                    = iDataType{raw.DataTypeUIntBuffer}
	UInt2dRect                    = iDataType{raw.DataTypeUInt2dRect}
	UIntAtomicCounter             = iDataType{raw.DataTypeUIntAtomicCounter}
)

func (t iDataType) String() string {
	switch t {
	case Float:
		return "Float"
	case Float2:
		return "Float2"
	case Float3:
		return "Float3"
	case Float4:
		return "Float4"
	case FloatMat2:
		return "FloatMat2"
	case FloatMat3:
		return "FloatMat3"
	case FloatMat4:
		return "FloatMat4"
	case FloatMat2x3:
		return "FloatMat2x3"
	case FloatMat2x4:
		return "FloatMat2x4"
	case FloatMat3x2:
		return "FloatMat3x2"
	case FloatMat3x4:
		return "FloatMat3x4"
	case FloatMat4x2:
		return "FloatMat4x2"
	case FloatMat4x3:
		return "FloatMat4x3"
	case Int:
		return "Int"
	case Int2:
		return "Int2"
	case Int3:
		return "Int3"
	case Int4:
		return "Int4"
	case UInt:
		return "UInt"
	case UInt2:
		return "UInt2"
	case UInt3:
		return "UInt3"
	case UInt4:
		return "UInt4"
	case Double:
		return "Double"
	case Double2:
		return "Double2"
	case Double3:
		return "Double3"
	case Double4:
		return "Double4"
	case DoubleMat2:
		return "DoubleMat2"
	case DoubleMat3:
		return "DoubleMat3"
	case DoubleMat4:
		return "DoubleMat4"
	case DoubleMat2x3:
		return "DoubleMat2x3"
	case DoubleMat2x4:
		return "DoubleMat2x4"
	case DoubleMat3x2:
		return "DoubleMat3x2"
	case DoubleMat3x4:
		return "DoubleMat3x4"
	case DoubleMat4x2:
		return "DoubleMat4x2"
	case DoubleMat4x3:
		return "DoubleMat4x3"

	case Bool:
		return "Bool"
	case Bool2:
		return "Bool2"
	case Bool3:
		return "Bool3"
	case Bool4:
		return "Bool4"
	case Sampler1d:
		return "Sampler1d"
	case Sampler2d:
		return "Sampler2d"
	case Sampler3d:
		return "Sampler3d"
	case SamplerCube:
		return "SamplerCube"
	case SamplerCubeShadow:
		return "SamplerCubeShadow"
	case Sampler1dShadow:
		return "Sampler1dShadow"
	case Sampler2dShadow:
		return "Sampler2dShadow"
	case Sampler1dArray:
		return "Sampler1dArray"
	case Sampler2dArray:
		return "Sampler2dArray"
	case Sampler1dArrayShadow:
		return "Sampler1dArrayShadow"
	case Sampler2dArrayShadow:
		return "Sampler2dArrayShadow"
	case Sampler2dMultisample:
		return "Sampler2dMultisample"
	case Sampler2dMultisampleArray:
		return "Sampler2dMultisampleArray"
	case SamplerBuffer:
		return "SamplerBuffer"
	case Sampler2dRect:
		return "Sampler2dRect"
	case Sampler2dRectShadow:
		return "Sampler2dRectShadow"
	case IntSampler1d:
		return "IntSampler1d"
	case IntSampler2d:
		return "IntSampler2d"
	case IntSampler3d:
		return "IntSampler3d"
	case IntSamplerCube:
		return "IntSamplerCube"
	case IntSampler1dArray:
		return "IntSampler1dArray"
	case IntSampler2dArray:
		return "IntSampler2dArray"
	case IntSampler2dMultisample:
		return "IntSampler2dMultisample"
	case IntSampler2dMultisampleArray:
		return "IntSampler2dMultisampleArray"
	case IntSamplerBuffer:
		return "IntSamplerBuffer"
	case IntSampler2dRect:
		return "IntSampler2dRect"
	case UIntSampler1d:
		return "UIntSampler1d"
	case UIntSampler2d:
		return "UIntSampler2d"
	case UIntSampler3d:
		return "UIntSampler3d"
	case UIntSamplerCube:
		return "UIntSamplerCube"
	case UIntSampler1dArray:
		return "UIntSampler1dArray"
	case UIntSampler2dArray:
		return "UIntSampler2dArray"
	case UIntSampler2dMultisample:
		return "UIntSampler2dMultisample"
	case UIntSampler2dMultisampleArray:
		return "UIntSampler2dMultisampleArray"
	case UIntSamplerBuffer:
		return "UIntSamplerBuffer"
	case UIntSampler2dRect:
		return "UIntSampler2dRect"
	case Image1d:
		return "Image1d"
	case Image2d:
		return "Image2d"
	case Image3d:
		return "Image3d"
	case ImageCube:
		return "ImageCube"
	case Image1dArray:
		return "Image1dArray"
	case Image2dArray:
		return "Image2dArray"
	case Image2dMultisample:
		return "Image2dMultisample"
	case Image2dMultisampleArray:
		return "Image2dMultisampleArray"
	case ImageBuffer:
		return "ImageBuffer"
	case Image2dRect:
		return "Image2dRect"
	case Int1d:
		return "Int1d"
	case Int2d:
		return "Int2d"
	case Int3d:
		return "Int3d"
	case IntCube:
		return "IntCube"
	case Int1dArray:
		return "Int1dArray"
	case Int2dArray:
		return "Int2dArray"
	case Int2dMultisample:
		return "Int2dMultisample"
	case Int2dMultisampleArray:
		return "Int2dMultisampleArray"
	case IntBuffer:
		return "IntBuffer"
	case Int2dRect:
		return "Int2dRect"
	case UInt1d:
		return "UInt1d"
	case UInt2d:
		return "UInt2d"
	case UInt3d:
		return "UInt3d"
	case UIntCube:
		return "UIntCube"
	case UInt1dArray:
		return "UInt1dArray"
	case UInt2dArray:
		return "UInt2dArray"
	case UInt2dMultisample:
		return "UInt2dMultisample"
	case UInt2dMultisampleArray:
		return "UInt2dMultisampleArray"
	case UIntBuffer:
		return "UIntBuffer"
	case UInt2dRect:
		return "UInt2dRect"
	case UIntAtomicCounter:
		return "UIntAtomicCounter"

	}
	return "Unknown DataType"
}

func (typ iDataType) Full() FullType {
	return FullType{DataType: typ, ArraySize: -1}
}

func (typ iDataType) Array(size uint) FullType {
	return FullType{DataType: typ, ArraySize: int(size)}
}

type FullType struct {
	DataType  iDataType
	ArraySize int
}

func (ft FullType) String() string {
	if ft.ArraySize <= -1 {
		return ft.DataType.String()
	}
	if ft.ArraySize == 0 {
		return ft.DataType.String() + "[]"
	}
	return fmt.Sprintf("%s[%d]", ft.DataType.String(), ft.ArraySize)
}
