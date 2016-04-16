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

type iDataType struct {
	t raw.Enum
}

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
	}
	return "Unknown DataType"
}
