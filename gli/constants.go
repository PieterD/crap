package gli

import "github.com/PieterD/glimmer/raw"

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
	StaticDraw  = iShaderType{raw.AccessTypeStaticDraw}
	StaticRead  = iShaderType{raw.AccessTypeStaticRead}
	StaticCopy  = iShaderType{raw.AccessTypeStaticCopy}
	StreamDraw  = iShaderType{raw.AccessTypeStreamDraw}
	StreamRead  = iShaderType{raw.AccessTypeStreamRead}
	StreamCopy  = iShaderType{raw.AccessTypeStreamCopy}
	DynamicDraw = iShaderType{raw.AccessTypeDynamicDraw}
	DynamicRead = iShaderType{raw.AccessTypeDynamicRead}
	DynamicCopy = iShaderType{raw.AccessTypeDynamicCopy}
)

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
