package gli

import (
	"fmt"

	"github.com/PieterD/crap/roguelike/gli/convc"
	"github.com/go-gl/gl/v2.1/gl"
)

type Program struct {
	id       uint32
	vertId   uint32
	fragId   uint32
	uniforms map[string]Uniform
}

func (program *Program) Id() uint32 {
	return program.id
}

func (program *Program) Use() {
	gl.UseProgram(program.id)
}

func (program *Program) Delete() {
	gl.DeleteProgram(program.id)
	gl.DeleteShader(program.vertId)
	gl.DeleteShader(program.fragId)
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
		id:       id,
		vertId:   vertexId,
		fragId:   fragmentId,
		uniforms: uniforms(id),
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

func uniforms(id uint32) map[string]Uniform {
	var maxlength int32
	gl.GetProgramiv(id, gl.ACTIVE_UNIFORM_MAX_LENGTH, &maxlength)
	var numuniforms int32
	gl.GetProgramiv(id, gl.ACTIVE_UNIFORMS, &numuniforms)
	buf := make([]byte, maxlength)
	m := make(map[string]Uniform)
	for index := uint32(0); index < uint32(numuniforms); index++ {
		name, dt, siz := uniform(id, index, buf)
		loc := gl.GetUniformLocation(id, &buf[0])
		if loc == -1 {
			panic(fmt.Errorf("Expected location for indexed uniform '%s'", name))
		}
		m[name] = Uniform{
			name:     name,
			location: loc,
			typ:      dt,
			siz:      siz,
		}
	}
	return m
}

func uniform(id uint32, index uint32, buf []byte) (name string, datatype uint32, size int32) {
	var length int32
	var isize int32
	var idatatype uint32
	gl.GetActiveUniform(id, index, int32(len(buf)), &length, &isize, &idatatype, &buf[0])
	return string(buf[:length : length+1]), idatatype, isize
}

type Uniform struct {
	program  *Program
	name     string
	location int32
	typ      uint32
	siz      int32
}

func (program *Program) Uniform(uniformname string) Uniform {
	uniform, ok := program.uniforms[uniformname]
	if !ok {
		return Uniform{
			program:  program,
			name:     uniformname,
			location: -1,
		}
	}
	uniform.program = program
	return uniform
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

func (uniform Uniform) SetSampler(data ...int32) {
	fmt.Printf("set %#v\n", uniform)
	switch uniform.typ {
	case gl.SAMPLER, gl.SAMPLER_1D, gl.SAMPLER_2D, gl.SAMPLER_3D:
		gl.ProgramUniform1iv(uniform.program.id, uniform.Location(), uniform.siz, &data[0])
	default:
		panic(fmt.Errorf("Unusable gl type '%04X", uniform.typ))
	}
}

func (uniform Uniform) SetInt(data ...int32) {
	fmt.Printf("set %#v\n", uniform)
	switch uniform.typ {
	case gl.INT:
		gl.ProgramUniform1iv(uniform.program.id, uniform.Location(), uniform.siz, &data[0])
	case gl.INT_VEC2:
		gl.ProgramUniform2iv(uniform.program.id, uniform.Location(), uniform.siz, &data[0])
	case gl.INT_VEC3:
		gl.ProgramUniform3iv(uniform.program.id, uniform.Location(), uniform.siz, &data[0])
	case gl.INT_VEC4:
		gl.ProgramUniform4iv(uniform.program.id, uniform.Location(), uniform.siz, &data[0])
	default:
		panic(fmt.Errorf("Unusable gl type '%04X", uniform.typ))
	}
}

func (uniform Uniform) SetFloat(data ...float32) {
	fmt.Printf("set %#v\n", uniform)
	switch uniform.typ {
	case gl.FLOAT:
		gl.ProgramUniform1fv(uniform.program.id, uniform.Location(), uniform.siz, &data[0])
	case gl.FLOAT_VEC2:
		gl.ProgramUniform2fv(uniform.program.id, uniform.Location(), uniform.siz, &data[0])
	case gl.FLOAT_VEC3:
		gl.ProgramUniform3fv(uniform.program.id, uniform.Location(), uniform.siz, &data[0])
	case gl.FLOAT_VEC4:
		gl.ProgramUniform4fv(uniform.program.id, uniform.Location(), uniform.siz, &data[0])
	default:
		panic(fmt.Errorf("Unusable gl type '%04X'", uniform.typ))
	}
}
