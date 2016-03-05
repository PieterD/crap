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
	CreateShader(shaderType ShaderType) Shader
	CreateProgram() Program
	CreateVertexArrayObject() VertexArrayObject
	UseProgram(program Program)
	UseNoProgram()
	BindVertexArrayObject(vao VertexArrayObject)
	UnbindVertexArrayObject()
}

func CreateShader(shaderType ShaderType) Shader {
	return Current.CreateShader(shaderType)
}
func CreateProgram() Program {
	return Current.CreateProgram()
}
func CreateVertexArrayObject() VertexArrayObject {
	return Current.CreateVertexArrayObject()
}
func UseProgram(program Program) {
	Current.UseProgram(program)
}
func UseNoProgram() {
	Current.UseNoProgram()
}
func BindVertexArrayObject(vao VertexArrayObject) {
	Current.BindVertexArrayObject(vao)
}
func UnbindVertexArrayObject() {
	Current.UnbindVertexArrayObject()
}
