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
}
