package glimmer

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Program struct {
	id         uint32
	attributes []programAttribute
	vao        *vertexArray

	attributeIndexByName map[string]uint32
	uniformIndexByName   map[string]uint32
	uniformByIndex       map[uint32]programUniform
}

type programAttribute struct {
	name  string
	index uint32
}

type programUniform struct {
	name      string
	index     uint32
	datatype  uint32
	arraysize uint32
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

	var attributes int32
	gl.GetProgramiv(program.id, gl.ACTIVE_ATTRIBUTES, &attributes)
	program.attributes = make([]programAttribute, attributes)
	program.attributeIndexByName = make(map[string]uint32)
	program.vao = createVertexArray()
	for i := 0; i < int(attributes); i++ {
		var buf [gl.ACTIVE_ATTRIBUTE_MAX_LENGTH]byte
		var length int32
		var size int32
		var xtype uint32
		gl.GetActiveAttrib(program.id, uint32(i), int32(len(buf)), &length, &size, &xtype, &buf[0])
		location := gl.GetAttribLocation(program.id, &buf[0])
		name := string(buf[:length])
		if location == -1 {
			gl.DeleteProgram(program.id)
			program.vao.Delete()
			return nil, &ShaderError{Desc: fmt.Sprintf("Attribute location for '%s' not found", name)}
		}
		index := uint32(location)
		program.attributes[i].name = name
		program.attributes[i].index = index
		program.attributeIndexByName[name] = index
	}

	var uniforms int32
	gl.GetProgramiv(program.id, gl.ACTIVE_UNIFORMS, &uniforms)
	program.uniformIndexByName = make(map[string]uint32)
	program.uniformByIndex = make(map[uint32]programUniform)
	for i := int32(0); i < uniforms; i++ {
		var buf [gl.ACTIVE_UNIFORM_MAX_LENGTH]byte
		var length int32
		gl.GetActiveUniformName(program.id, uint32(i), int32(len(buf)), &length, &buf[0])
		name := string(buf[:length])
		location := gl.GetUniformLocation(program.id, &buf[0])
		if location == -1 {
			gl.DeleteProgram(program.id)
			program.vao.Delete()
			return nil, &ShaderError{Desc: fmt.Sprintf("Uniform location for '%s' not found", name)}
		}
		index := uint32(location)

		var datatype int32
		gl.GetActiveUniformsiv(program.id, 1, &index, gl.UNIFORM_TYPE, &datatype)
		var arraysize int32
		gl.GetActiveUniformsiv(program.id, 1, &index, gl.UNIFORM_SIZE, &arraysize)
		uni := programUniform{
			name:      name,
			index:     uint32(location),
			datatype:  uint32(datatype),
			arraysize: uint32(arraysize),
		}
		program.uniformIndexByName[name] = index
		program.uniformByIndex[index] = uni
	}

	return program, nil
}

func (program *Program) AttributeByName(name string, pointer *ArrayPointer) bool {
	index, ok := program.attributeIndexByName[name]
	if !ok {
		return false
	}
	return program.AttributeByIndex(index, pointer)
}

func (program *Program) AttributeByIndex(index uint32, pointer *ArrayPointer) bool {
	program.vao.Enable(index, pointer)
	return true
}

func (program *Program) UniformFloat(name string, value float32) bool {
	index, ok := program.uniformIndexByName[name]
	if !ok {
		return false
	}
	gl.ProgramUniform1f(program.id, int32(index), value)
	return true
}

func (program *Program) Delete() {
	if program == nil {
		return
	}
	gl.DeleteProgram(program.id)
	program.vao.Delete()
}

func (program *Program) Bind() {
	program.vao.Bind()
	gl.UseProgram(program.id)
}

func (program *Program) Unbind() {
	gl.UseProgram(0)
	program.vao.Unbind()
}
