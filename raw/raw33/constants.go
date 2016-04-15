package raw33

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type glEnum int64

type glEnumList []int64

func (l glEnumList) valid(i int) {
	if i == 0 {
		panic(fmt.Errorf("Zero enum index %d", i))
	}
	if i >= len(l) {
		panic(fmt.Errorf("High enum index %d", i))
	}
}

func (l glEnumList) unsigned(i int) uint32 {
	l.valid(i)
	return uint32(l[i])
}

func (l glEnumList) signed(i int) int32 {
	l.valid(i)
	return int32(l[i])
}

var aShaderType = glEnumList{
	0,
	gl.VERTEX_SHADER,
	gl.GEOMETRY_SHADER,
	gl.FRAGMENT_SHADER,
}
