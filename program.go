package glimmer

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Program struct {
	id          uint32
	attributes  []programAttribute
	indexByName map[string]uint32
	vaoByIndex  map[uint32]*vertexArray

	uniformIndexByName map[string]uint32
	uniformByIndex     map[uint32]programUniform
}

type programAttribute struct {
	name  string
	index uint32
	vao   *vertexArray
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
	program.indexByName = make(map[string]uint32)
	program.vaoByIndex = make(map[uint32]*vertexArray)
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
			for j := 0; j < i; j++ {
				program.attributes[j].vao.delete()
			}
			return nil, &ShaderError{Desc: fmt.Sprintf("Attribute location for '%s' not found", name)}
		}
		index := uint32(location)
		vao := createVertexArray()
		program.attributes[i].name = name
		program.attributes[i].index = index
		program.attributes[i].vao = vao
		program.indexByName[name] = index
		program.vaoByIndex[index] = vao
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
			for j := int32(0); j < i; j++ {
				program.attributes[j].vao.delete()
			}
			return nil, &ShaderError{Desc: fmt.Sprintf("Uniform location for '%s' not found", name)}
		}
		index := uint32(location)
		fmt.Printf("name: '%s' %d\n", name, index)

		var datatype int32
		gl.GetActiveUniformsiv(program.id, 1, &index, gl.UNIFORM_TYPE, &datatype)
		var arraysize int32
		gl.GetActiveUniformsiv(program.id, 1, &index, gl.UNIFORM_SIZE, &arraysize)
		fmt.Printf("datatype: %d size: %d\n", datatype, arraysize)
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
	index, ok := program.indexByName[name]
	if !ok {
		return false
	}
	return program.AttributeByIndex(index, pointer)
}

func (program *Program) AttributeByIndex(index uint32, pointer *ArrayPointer) bool {
	va, ok := program.vaoByIndex[index]
	if !ok {
		return false
	}
	va.enable(index, pointer)
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
	gl.DeleteProgram(program.id)
	for _, vao := range program.vaoByIndex {
		vao.delete()
	}
}

func (program *Program) Use() {
	gl.UseProgram(program.id)
}
