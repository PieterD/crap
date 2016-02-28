package glimmer

import (
	"github.com/PieterD/crap/glimmer/convc"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Shader struct {
	id         uint32
	shaderType ShaderType
}

type ShaderType uint32

const (
	VertexShader         ShaderType = gl.VERTEX_SHADER
	GeometryShader       ShaderType = gl.GEOMETRY_SHADER
	FragmentShader       ShaderType = gl.FRAGMENT_SHADER
	ComputeShader        ShaderType = gl.COMPUTE_SHADER
	TessControlShader    ShaderType = gl.TESS_CONTROL_SHADER
	TessEvaluationShader ShaderType = gl.TESS_EVALUATION_SHADER
)

func CreateShader(shaderType ShaderType, source ...string) (*Shader, error) {
	shader := new(Shader)
	shader.id = gl.CreateShader(uint32(shaderType))
	if shader.id == 0 {
		return nil, GetError()
	}
	ptr, free := convc.MultiStringToC(source...)
	defer free()
	gl.ShaderSource(shader.id, 1, ptr, nil)
	gl.CompileShader(shader.id)
	var status int32
	gl.GetShaderiv(shader.id, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var loglength int32
		gl.GetShaderiv(shader.id, gl.INFO_LOG_LENGTH, &loglength)
		log := make([]byte, loglength)
		var logptr *uint8 = &log[0]
		gl.GetShaderInfoLog(shader.id, loglength, nil, logptr)
		gl.DeleteShader(shader.id)
		return nil, &ShaderError{Desc: string(log[:len(log)-1])}
	}
	return shader, nil
}

func (shader *Shader) Delete() {
	gl.DeleteShader(shader.id)
}

type Program struct {
	id uint32
}

func CreateProgram(shaders ...*Shader) (*Program, error) {
	program := new(Program)
	program.id = gl.CreateProgram()
	if program.id == 0 {
		return nil, GetError()
	}
	for _, shader := range shaders {
		gl.AttachShader(program.id, shader.id)
	}
	gl.LinkProgram(program.id)

	var status int32
	gl.GetProgramiv(program.id, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var loglength int32
		gl.GetProgramiv(program.id, gl.INFO_LOG_LENGTH, &loglength)
		log := make([]byte, loglength)
		var logptr *uint8 = &log[0]
		gl.GetProgramInfoLog(program.id, loglength, nil, logptr)
		gl.DeleteProgram(program.id)
		return nil, &ShaderError{Desc: string(log[:len(log)-1])}
	}
	return program, nil
}

func (program *Program) Delete() {
	gl.DeleteProgram(program.id)
}

func (program *Program) Use() {
	gl.UseProgram(program.id)
}
