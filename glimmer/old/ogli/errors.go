package gli

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type ShaderError struct {
	Desc string
}

func (se *ShaderError) Error() string {
	return "GL Shader Error: " + se.Desc
}

type GlError struct {
	Id   uint32
	Desc string
	More *GlError
}

func (gle *GlError) Error() string {
	return gle.message(0)
}

func (gle *GlError) message(depth int) string {
	var str string
	if depth == 0 {
		str = "GL Errors: "
	} else {
		str = ", "
	}
	str += gle.Desc
	if gle.More == nil {
		return str
	}
	return str + gle.More.message(depth+1)
}

func GetError() error {
	gle := getError()
	if gle == nil {
		return nil
	}
	return gle
}

func getError() *GlError {
	id := gl.GetError()
	var desc string
	switch id {
	case gl.NO_ERROR:
		return nil
	case gl.INVALID_ENUM:
		desc = "GL_INVALID_ENUM"
	case gl.INVALID_VALUE:
		desc = "GL_INVALID_VALUE"
	case gl.INVALID_OPERATION:
		desc = "GL_INVALID_OPERATION"
	case gl.INVALID_FRAMEBUFFER_OPERATION:
		desc = "GL_INVALID_FRAMEBUFFER_OPERATION"
	case gl.OUT_OF_MEMORY:
		desc = "GL_OUT_OF_MEMORY"
	case gl.STACK_UNDERFLOW:
		desc = "GL_STACK_UNDERFLOW"
	case gl.STACK_OVERFLOW:
		desc = "GL_STACK_OVERFLOW"
	default:
		desc = fmt.Sprintf("UNKNOWN_ERROR(%d)", id)
	}
	return &GlError{Id: id, Desc: desc, More: getError()}
}
