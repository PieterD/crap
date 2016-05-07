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

	gli.BufferUsageStreamDraw:  gl.STREAM_DRAW,
	gli.BufferUsageStreamRead:  gl.STREAM_READ,
	gli.BufferUsageStreamCopy:  gl.STREAM_COPY,
	gli.BufferUsageStaticDraw:  gl.STATIC_DRAW,
	gli.BufferUsageStaticRead:  gl.STATIC_READ,
	gli.BufferUsageStaticCopy:  gl.STATIC_COPY,
	gli.BufferUsageDynamicDraw: gl.DYNAMIC_DRAW,
	gli.BufferUsageDynamicRead: gl.DYNAMIC_READ,
	gli.BufferUsageDynamicCopy: gl.DYNAMIC_COPY,

	gli.BufferTargetArray:             gl.ARRAY_BUFFER,
	gli.BufferTargetCopyRead:          gl.COPY_READ_BUFFER,
	gli.BufferTargetCopyWrite:         gl.COPY_WRITE_BUFFER,
	gli.BufferTargetElementArray:      gl.ELEMENT_ARRAY_BUFFER,
	gli.BufferTargetPixelPack:         gl.PIXEL_PACK_BUFFER,
	gli.BufferTargetPixelUnpack:       gl.PIXEL_UNPACK_BUFFER,
	gli.BufferTargetTexture:           gl.TEXTURE_BUFFER,
	gli.BufferTargetTransformFeedback: gl.TRANSFORM_FEEDBACK_BUFFER,
	gli.BufferTargetUniform:           gl.UNIFORM_BUFFER,
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
