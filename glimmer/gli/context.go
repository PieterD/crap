package gli

import "github.com/go-gl/gl/v3.3-core/gl"

func Init() {
	gl.Init()
}

type iContext struct {
}

type Context interface {
	CreateShader(shaderType ShaderType) Shader
	CreateProgram() Program
	UseProgram(Program)
	UseNoProgram()
}

func CreateContext() Context {
	return iContext{}
}
