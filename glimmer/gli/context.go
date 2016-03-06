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
	UseProgram(program Program)
	UseNoProgram()
	BindVertexArrayObject(vao VertexArrayObject)
	UnbindVertexArrayObject()
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
