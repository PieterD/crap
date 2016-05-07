package gl330

import (
	"fmt"
	"github.com/PieterD/glimmer/gli"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type enum int64

var enumMap = map[fmt.Stringer]enum{
	gli.ShaderTypeVertex:   gl.VERTEX_SHADER,
	gli.ShaderTypeGeometry: gl.GEOMETRY_SHADER,
	gli.ShaderTypeFragment: gl.FRAGMENT_SHADER,
}

func toEnum(iface fmt.Stringer) enum {
	e, ok := enumMap[iface]
	if !ok {
		panic(fmt.Errorf("Unsupported enum: %s", iface.String()))
	}
	return e
}

func toUint(iface fmt.Stringer) uint32 {
	return uint32(toEnum(iface))
}

func toInt(iface fmt.Stringer) int32 {
	return int32(toEnum(iface))
}
