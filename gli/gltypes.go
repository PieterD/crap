package gli

import "fmt"

//go:generate stringer -type=ShaderType
type ShaderType uint32

const (
	ShaderTypeVertex ShaderType = iota + 1
	ShaderTypeGeometry
	ShaderTypeFragment
)

type ShaderId uint32

type ProgramId uint32

//go:generate stringer -type=BaseType
type BaseType uint32

const (
	BaseTypeFloat BaseType = iota + 1
	BaseTypeInt
	BaseTypeUnsignedInt
	BaseTypeDouble
	BaseTypeBool
)

//go:generate stringer -type=ResourceType
type ResourceType uint32

const (
	ResourceTypeUniform ResourceType = iota + 1
	ResourceTypeAttribute
)

//go:generate stringer -type=SamplerType
type SamplerType uint32

const (
	SamplerType1d SamplerType = iota + 1
	SamplerType2d
	SamplerType3d
	SamplerTypeCude
	SamplerType2dRect
	SamplerType1dArray
	SamplerType2dArray
	SamplerTypeCubeArray
	SamplerTypeBuffer
	SamplerType2dMS
	SamplerType2dMSArray

	SamplerType1dShadow
	SamplerType2dShadow
	SamplerTypeCubeShadow
	SamplerType2dRectShadow
	SamplerType1dArrayShadow
	SamplerType2dArrayShadow
	SamplerTypeCubeArrayShadow
)

type DataType struct {
	Base    BaseType
	Sampler SamplerType
	Cols    byte
	Rows    byte
	Size    uint
}

func (dt DataType) GoString() string {
	return fmt.Sprintf("gli.DataType{Base: %s, Sampler: %s, Cols: %d, Rows: %d, Size: %d}", dt.Base.String(), dt.Sampler.String(), dt.Cols, dt.Rows, dt.Size)
}

func (dt DataType) IsValid() error {
	if dt.Cols == 0 && dt.Rows > 0 {
		return fmt.Errorf("Invalid DataType %#v: Rows without Columns", dt)
	}
	if dt.Base > BaseTypeBool {
		return fmt.Errorf("Invalid DataType %#v: Unknown BaseType", dt)
	}
	if dt.Base == 0 {
		return fmt.Errorf("Invalid DataType %#v: No base type", dt)
	}
	if dt.Sampler > 0 {
		return dt.isValidSampler()
	}
	if dt.Rows > 0 {
		return dt.isValidMatrix()
	}
	if dt.Cols > 0 {
		return dt.isValidVector()
	}
	return nil
}

func (dt DataType) isValidSampler() error {
	if dt.Sampler > SamplerTypeCubeArrayShadow {
		return fmt.Errorf("Invalid DataType %#v: Unknown SamplerType", dt)
	}
	if dt.Cols > 0 || dt.Rows > 0 {
		return fmt.Errorf("Invalid DataType %#v: Sampler types not allowed to have columns or rows", dt)
	}
	if dt.Sampler > SamplerType1dShadow {
		if dt.Base != BaseTypeFloat {
			return fmt.Errorf("Invalid DataType %#v: Shadow samplers can only have BaseType Float", dt)
		}
	}
	if dt.Base != BaseTypeFloat && dt.Base != BaseTypeInt && dt.Base != BaseTypeUnsignedInt {
		return fmt.Errorf("Invalid DataTyp %#v: Only float, int and uint samplers are allowed", dt)
	}
	return nil
}
func (dt DataType) isValidMatrix() error {
	if dt.Rows > 4 || dt.Cols > 4 {
		return fmt.Errorf("Invalid DataType %#v: Matrix dimensions too large", dt)
	}
	if dt.Rows == 1 || dt.Cols == 1 {
		return fmt.Errorf("Invalid DataType %#v: Matrix dimensions too small", dt)
	}
	if dt.Base != BaseTypeFloat && dt.Base != BaseTypeDouble {
		return fmt.Errorf("Invalid DataType %#v: Only Float or Double matrices are allowed", dt)
	}
	return nil
}
func (dt DataType) isValidVector() error {
	if dt.Cols > 4 {
		return fmt.Errorf("Invalid DataType %#v: Vector dimensions too large", dt)
	}
	if dt.Cols == 1 {
		return fmt.Errorf("Invalid DataType %#v: Vector dimensions too small", dt)
	}
	return nil
}
