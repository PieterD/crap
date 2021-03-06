package gli

import (
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func Draw(program Program, vao VertexArrayObject, object Object) {
	Current.Draw(program, vao, object)
}
func (context iContext) Draw(program Program, vao VertexArrayObject, object Object) {
	context.BindVertexArrayObject(vao)
	context.BindProgram(program)
	if object.IndexType != 0 {
		if object.IndexBase > 0 {
			gl.DrawElementsBaseVertex(uint32(object.Mode), int32(object.Vertices), uint32(object.IndexType), unsafe.Pointer(uintptr(object.Start)), int32(object.IndexBase))
		} else {
			gl.DrawElements(uint32(object.Mode), int32(object.Vertices), uint32(object.IndexType), unsafe.Pointer(uintptr(object.Start)))
		}
	} else {
		gl.DrawArrays(uint32(object.Mode), int32(object.Start), int32(object.Vertices))
	}
	context.UnbindProgram()
	context.UnbindVertexArrayObject()
}
