package gl330

import (
	"errors"
	"fmt"
	"github.com/PieterD/gl/v3.3-core/gl"
	"github.com/PieterD/glimmer/gli"
)

type gl330 struct{}

func init() {
	gli.Register(gl330{})
}

func (_ gl330) Name() string {
	return "gl330"
}

func getError() error {
	id := gl.GetError()
	var desc string
	switch id {
	case gl.NO_ERROR:
		return nil
	case gl.INVALID_ENUM:
		desc = "Invalid enum"
	case gl.INVALID_VALUE:
		desc = "Invalid value"
	case gl.INVALID_OPERATION:
		desc = "Invalid operation"
	case gl.INVALID_FRAMEBUFFER_OPERATION:
		desc = "Invalid framebuffer operation"
	case gl.OUT_OF_MEMORY:
		desc = "Out of memory"
	case gl.STACK_UNDERFLOW:
		desc = "Stack underflow"
	case gl.STACK_OVERFLOW:
		desc = "Stack overflow"
	default:
		desc = fmt.Sprintf("Unknown error %d", id)
	}
	next := getError()
	if next != nil {
		desc = ", " + next.Error()
	}
	return errors.New(desc)
}
