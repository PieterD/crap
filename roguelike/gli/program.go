package gli

import (
	"fmt"

	"github.com/PieterD/crap/roguelike/gli/convc"
	"github.com/go-gl/gl/v2.1/gl"
)

type Program struct {
	id uint32
}

func (program *Program) Id() uint32 {
	return program.id
}

func (program *Program) Use() {
	gl.UseProgram(program.id)
}

func (program *Program) Delete() {
	gl.DeleteProgram(program.id)
}

func NewProgram(vertexSource, fragmentSource string) (*Program, error) {
	vertexId, err := newShader(vertexSource, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}
	fragmentId, err := newShader(fragmentSource, gl.FRAGMENT_SHADER)
	if err != nil {
		gl.DeleteShader(vertexId)
		return nil, err
	}
	id := gl.CreateProgram()
	if id == 0 {
		gl.DeleteShader(vertexId)
		gl.DeleteShader(fragmentId)
		return nil, fmt.Errorf("Unable to allocate program")
	}
	gl.AttachShader(id, vertexId)
	gl.AttachShader(id, fragmentId)
	gl.LinkProgram(id)
	var result int32
	gl.GetProgramiv(id, gl.LINK_STATUS, &result)
	if result == int32(gl.FALSE) {
		var loglength int32
		gl.GetProgramiv(id, gl.INFO_LOG_LENGTH, &loglength)
		log := make([]byte, loglength)
		var length int32
		gl.GetProgramInfoLog(id, loglength, &length, &log[0])
		gl.DeleteShader(vertexId)
		gl.DeleteShader(fragmentId)
		gl.DeleteProgram(id)
		return nil, fmt.Errorf("Unable to link program: %s", log[:length])
	}

	return &Program{
		id: id,
	}, nil
}

func newShader(source string, shaderType uint32) (uint32, error) {
	id := gl.CreateShader(shaderType)
	if id == 0 {
		return 0, fmt.Errorf("Unable to allocate shader")
	}
	ptr, free := convc.StringToC(source)
	defer free()
	gl.ShaderSource(id, 1, &ptr, nil)
	gl.CompileShader(id)

	var result int32
	gl.GetShaderiv(id, gl.COMPILE_STATUS, &result)
	if result == int32(gl.FALSE) {
		var loglength int32
		gl.GetShaderiv(id, gl.INFO_LOG_LENGTH, &loglength)
		log := make([]byte, loglength)
		var length int32
		gl.GetShaderInfoLog(id, loglength, &length, &log[0])
		gl.DeleteShader(id)
		return 0, fmt.Errorf("Unable to compile shader: %s", log[:length])
	}

	return id, nil
}

type Attrib struct {
	name     string
	location int32
}

func (program *Program) Attrib(attrname string) Attrib {
	location := gl.GetAttribLocation(program.id, gl.Str(attrname+"\x00"))
	return Attrib{
		name:     attrname,
		location: location,
	}
}

func (attrib Attrib) Name() string {
	return attrib.name
}

func (attrib Attrib) Valid() bool {
	return attrib.location != -1
}

func (attrib Attrib) Location() uint32 {
	if !attrib.Valid() {
		panic(fmt.Errorf("Could not find location for attribute '%s'", attrib.name))
	}
	return uint32(attrib.location)
}

type Uniform struct {
	program  *Program
	name     string
	location int32
}

func (program *Program) Uniform(uniformname string) Uniform {
	location := gl.GetUniformLocation(program.id, gl.Str(uniformname+"\x00"))
	return Uniform{
		program:  program,
		name:     uniformname,
		location: location,
	}
}

func (uniform Uniform) Name() string {
	return uniform.name
}

func (uniform Uniform) Valid() bool {
	return uniform.location != -1
}

func (uniform Uniform) Location() int32 {
	if !uniform.Valid() {
		panic(fmt.Errorf("Could not find location for uniform '%s'", uniform.name))
	}
	return int32(uniform.location)
}

func (uniform Uniform) SetInt(data ...int32) {
	gl.ProgramUniform1iv(uniform.program.id, uniform.Location(), 1, &data[0])
}
