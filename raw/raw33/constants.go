package raw33

import (
	"fmt"

	"github.com/PieterD/glimmer/raw"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type glEnum int64

type glEnumList []int64

func (l glEnumList) valid(i raw.Enum) {
	if i == 0 {
		panic(fmt.Errorf("Zero enum index %d", i))
	}
	if int(i) >= len(l) {
		panic(fmt.Errorf("High enum index %d", i))
	}
}

func (l glEnumList) unsigned(i raw.Enum) uint32 {
	l.valid(i)
	return uint32(l[i])
}

func (l glEnumList) signed(i raw.Enum) int32 {
	l.valid(i)
	return int32(l[i])
}

func (l glEnumList) reverse(v int64) (raw.Enum, bool) {
	for i, j := range l {
		if j == v {
			return raw.Enum(i), true
		}
	}
	return 0, false
}

var aShaderType = glEnumList{
	0,
	gl.VERTEX_SHADER,
	gl.GEOMETRY_SHADER,
	gl.FRAGMENT_SHADER,
}

var aAccessType = glEnumList{
	0,
	gl.STATIC_DRAW,
	gl.STATIC_READ,
	gl.STATIC_COPY,
	gl.STREAM_DRAW,
	gl.STREAM_READ,
	gl.STREAM_COPY,
	gl.DYNAMIC_DRAW,
	gl.DYNAMIC_READ,
	gl.DYNAMIC_COPY,
}

var aDataType = glEnumList{
	0,
	gl.FLOAT,
	gl.FLOAT_VEC2,
	gl.FLOAT_VEC3,
	gl.FLOAT_VEC4,
	gl.FLOAT_MAT2,
	gl.FLOAT_MAT3,
	gl.FLOAT_MAT4,
	gl.FLOAT_MAT2x3,
	gl.FLOAT_MAT2x4,
	gl.FLOAT_MAT3x2,
	gl.FLOAT_MAT3x4,
	gl.FLOAT_MAT4x2,
	gl.FLOAT_MAT4x3,
	gl.INT,
	gl.INT_VEC2,
	gl.INT_VEC3,
	gl.INT_VEC4,
	gl.UNSIGNED_INT,
	gl.UNSIGNED_INT_VEC2,
	gl.UNSIGNED_INT_VEC3,
	gl.UNSIGNED_INT_VEC4,
	gl.DOUBLE,
	gl.DOUBLE_VEC2,
	gl.DOUBLE_VEC3,
	gl.DOUBLE_VEC4,
	gl.DOUBLE_MAT2,
	gl.DOUBLE_MAT3,
	gl.DOUBLE_MAT4,
	gl.DOUBLE_MAT2x3,
	gl.DOUBLE_MAT2x4,
	gl.DOUBLE_MAT3x2,
	gl.DOUBLE_MAT3x4,
	gl.DOUBLE_MAT4x2,
	gl.DOUBLE_MAT4x3,

	gl.BOOL,
	gl.BOOL_VEC2,
	gl.BOOL_VEC3,
	gl.BOOL_VEC4,
	gl.SAMPLER_1D,
	gl.SAMPLER_2D,
	gl.SAMPLER_3D,
	gl.SAMPLER_CUBE,
	gl.SAMPLER_CUBE_SHADOW,
	gl.SAMPLER_1D_SHADOW,
	gl.SAMPLER_2D_SHADOW,
	gl.SAMPLER_1D_ARRAY,
	gl.SAMPLER_2D_ARRAY,
	gl.SAMPLER_1D_ARRAY_SHADOW,
	gl.SAMPLER_2D_ARRAY_SHADOW,
	gl.SAMPLER_2D_MULTISAMPLE,
	gl.SAMPLER_2D_MULTISAMPLE_ARRAY,
	gl.SAMPLER_BUFFER,
	gl.SAMPLER_2D_RECT,
	gl.SAMPLER_2D_RECT_SHADOW,
	gl.INT_SAMPLER_1D,
	gl.INT_SAMPLER_2D,
	gl.INT_SAMPLER_3D,
	gl.INT_SAMPLER_CUBE,
	gl.INT_SAMPLER_1D_ARRAY,
	gl.INT_SAMPLER_2D_ARRAY,
	gl.INT_SAMPLER_2D_MULTISAMPLE,
	gl.INT_SAMPLER_2D_MULTISAMPLE_ARRAY,
	gl.INT_SAMPLER_BUFFER,
	gl.INT_SAMPLER_2D_RECT,
	gl.UNSIGNED_INT_SAMPLER_1D,
	gl.UNSIGNED_INT_SAMPLER_2D,
	gl.UNSIGNED_INT_SAMPLER_3D,
	gl.UNSIGNED_INT_SAMPLER_CUBE,
	gl.UNSIGNED_INT_SAMPLER_1D_ARRAY,
	gl.UNSIGNED_INT_SAMPLER_2D_ARRAY,
	gl.UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE,
	gl.UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY,
	gl.UNSIGNED_INT_SAMPLER_BUFFER,
	gl.UNSIGNED_INT_SAMPLER_2D_RECT,
	gl.IMAGE_1D,
	gl.IMAGE_2D,
	gl.IMAGE_3D,
	gl.IMAGE_CUBE,
	gl.IMAGE_1D_ARRAY,
	gl.IMAGE_2D_ARRAY,
	gl.IMAGE_2D_MULTISAMPLE,
	gl.IMAGE_2D_MULTISAMPLE_ARRAY,
	gl.IMAGE_BUFFER,
	gl.IMAGE_2D_RECT,
	gl.INT_IMAGE_1D,
	gl.INT_IMAGE_2D,
	gl.INT_IMAGE_3D,
	gl.INT_IMAGE_CUBE,
	gl.INT_IMAGE_1D_ARRAY,
	gl.INT_IMAGE_2D_ARRAY,
	gl.INT_IMAGE_2D_MULTISAMPLE,
	gl.INT_IMAGE_2D_MULTISAMPLE_ARRAY,
	gl.INT_IMAGE_BUFFER,
	gl.INT_IMAGE_2D_RECT,
	gl.UNSIGNED_INT_IMAGE_1D,
	gl.UNSIGNED_INT_IMAGE_2D,
	gl.UNSIGNED_INT_IMAGE_3D,
	gl.UNSIGNED_INT_IMAGE_CUBE,
	gl.UNSIGNED_INT_IMAGE_1D_ARRAY,
	gl.UNSIGNED_INT_IMAGE_2D_ARRAY,
	gl.UNSIGNED_INT_IMAGE_2D_MULTISAMPLE,
	gl.UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY,
	gl.UNSIGNED_INT_IMAGE_BUFFER,
	gl.UNSIGNED_INT_IMAGE_2D_RECT,
	gl.UNSIGNED_INT_ATOMIC_COUNTER,
}
