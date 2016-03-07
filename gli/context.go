package gli

import "github.com/go-gl/gl/v3.3-core/gl"

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
