package gli

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type iProgram struct {
	id uint32
}

type Program interface {
	Id() uint32
	Delete()
	GetIV(param ProgramParameter) int32
	GetActiveUniformIV(param UniformParameter, index uint32) int32
	Attributes() AttributeCollection
	Uniforms() UniformCollection
}

func (context iContext) CreateProgram(shaders ...Shader) (Program, error) {
	id := gl.CreateProgram()
	var program iProgram
	program.id = id
	if id == 0 {
		// TODO: GetError
		return nil, fmt.Errorf("Unable to create program")
	}
	for _, shader := range shaders {
		gl.AttachShader(program.Id(), shader.Id())
	}
	gl.LinkProgram(program.Id())
	status := program.GetIV(LINK_STATUS)
	if status == int32(FALSE) {
		loglength := program.GetIV(PROGRAM_INFO_LOG_LENGTH)
		log := make([]byte, loglength)
		var length int32
		gl.GetProgramInfoLog(program.Id(), loglength, &length, &log[0])
		program.Delete()
		// TODO: ShaderError
		return nil, fmt.Errorf("Unable to link program: %s", log[:length])
	}

	return program, nil
}

func (context iContext) BindProgram(program Program) {
	gl.UseProgram(program.Id())
}

func (context iContext) UnbindProgram() {
	gl.UseProgram(0)
}

func (program iProgram) Id() uint32 {
	return program.id
}

func (program iProgram) Delete() {
	gl.DeleteProgram(program.id)
}

func (program iProgram) GetIV(param ProgramParameter) int32 {
	var pi int32
	gl.GetProgramiv(program.id, uint32(param), &pi)
	return pi
}

func (program iProgram) GetActiveUniformIV(param UniformParameter, index uint32) int32 {
	var pi int32
	gl.GetActiveUniformsiv(program.id, 1, &index, uint32(param), &pi)
	return pi
}
