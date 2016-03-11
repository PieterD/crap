package gli

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

func Init() {
	gl.Init()
}

type iContext struct {
}

var Current Context = CreateContext()

func CreateContext() Context {
	return iContext{}
}

type Context interface {
	CreateShader(shaderType ShaderType, source ...string) (Shader, error)
	CreateProgram(shaders ...Shader) (Program, error)
	CreateVertexArrayObject() VertexArrayObject
	CreateBuffer(accesshint BufferAccessTypeHint, targethint BufferTarget) Buffer
	BindProgram(program Program)
	UnbindProgram()
	BindVertexArrayObject(vao VertexArrayObject)
	UnbindVertexArrayObject()
	BindBuffer(target BufferTarget, buffer Buffer)
	UnbindBuffer(target BufferTarget)
	DrawArrays(mode DrawMode, program Program, array ArrayInstance)
	ClearColor(r, g, b, a float32)
	Clear(bits ...ClearBit)
}

func CreateShader(shaderType ShaderType, source ...string) (Shader, error) {
	return Current.CreateShader(shaderType, source...)
}
func CreateProgram(shaders ...Shader) (Program, error) {
	return Current.CreateProgram(shaders...)
}
func CreateVertexArrayObject() VertexArrayObject {
	return Current.CreateVertexArrayObject()
}
func CreateBuffer(accesshint BufferAccessTypeHint, targethint BufferTarget) Buffer {
	return Current.CreateBuffer(accesshint, targethint)
}
func BindProgram(program Program) {
	Current.BindProgram(program)
}
func UnbindProgram() {
	Current.UnbindProgram()
}
func BindVertexArrayObject(vao VertexArrayObject) {
	Current.BindVertexArrayObject(vao)
}
func UnbindVertexArrayObject() {
	Current.UnbindVertexArrayObject()
}
func BindBuffer(target BufferTarget, buffer Buffer) {
	Current.BindBuffer(target, buffer)
}
func UnbindBuffer(target BufferTarget) {
	Current.UnbindBuffer(target)
}
func DrawArrays(mode DrawMode, program Program, array ArrayInstance) {
	Current.DrawArrays(mode, program, array)
}
func (context iContext) DrawArrays(mode DrawMode, program Program, array ArrayInstance) {
	BindVertexArrayObject(array.Vao)
	BindProgram(program)
	gl.DrawArrays(uint32(mode), int32(array.First), int32(array.Count))
	UnbindProgram()
	UnbindVertexArrayObject()
}
func ClearColor(r, g, b, a float32) {
	Current.ClearColor(r, g, b, a)
}
func (context iContext) ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}
func Clear(bits ...ClearBit) {
	Current.Clear(bits...)
}
func (context iContext) Clear(bits ...ClearBit) {
	var b uint32
	for _, bit := range bits {
		b |= uint32(bit)
	}
	gl.Clear(b)
}
