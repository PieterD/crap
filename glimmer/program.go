package glimmer

import (
	"fmt"

	"github.com/PieterD/crap/glimmer/gli"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Program struct {
	program    gli.Program
	attributes []programAttribute
	vao        gli.VertexArrayObject

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
	p := gli.CreateProgram()
	program.program = p
	if !p.Valid() {
		return nil, GetError()
	}
	for _, shader := range shaders {
		p.AttachShader(shader.shader)
	}
	p.Link()

	if !p.GetLinkSuccess() {
		loglength := p.GetInfoLogLength()
		log := make([]byte, loglength)
		log = p.GetInfoLog(log)
		return nil, &ShaderError{Desc: string(log)}
	}

	attributes := p.GetIV(gli.ACTIVE_ATTRIBUTES)
	program.attributes = make([]programAttribute, attributes)
	program.attributeIndexByName = make(map[string]uint32)
	program.vao = gli.CreateVertexArrayObject()
	for i := 0; i < int(attributes); i++ {
		buf := make([]byte, p.GetIV(gli.ACTIVE_ATTRIBUTE_MAX_LENGTH))
		namebytes, _, _ := p.GetActiveAttrib(uint32(i), buf)
		location := p.GetAttribLocationBytes(namebytes)
		name := string(namebytes)
		if location == -1 {
			p.Delete()
			program.vao.Delete()
			return nil, &ShaderError{Desc: fmt.Sprintf("Attribute location for '%s' not found", name)}
		}
		index := uint32(location)
		program.attributes[i].name = name
		program.attributes[i].index = index
		program.attributeIndexByName[name] = index
	}

	uniforms := p.GetIV(gli.ACTIVE_UNIFORMS)
	program.uniformIndexByName = make(map[string]uint32)
	program.uniformByIndex = make(map[uint32]programUniform)
	buf := make([]byte, p.GetIV(gli.ACTIVE_UNIFORM_MAX_LENGTH))
	for i := int32(0); i < uniforms; i++ {
		namebytes := p.GetActiveUniformName(uint32(i), buf)
		location := p.GetUniformLocationBytes(namebytes)
		name := string(namebytes)
		if location == -1 {
			p.Delete()
			program.vao.Delete()
			return nil, &ShaderError{Desc: fmt.Sprintf("Uniform location for '%s' not found", name)}
		}
		index := uint32(location)

		datatype := p.GetActiveUniformIV(gli.UNIFORM_TYPE, uint32(i))
		arraysize := p.GetActiveUniformIV(gli.UNIFORM_SIZE, uint32(i))
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
	pointer.buffer.bind()
	program.vao.EnableAttrib(index)
	program.vao.AttribPointer(index, int32(pointer.datasize), gli.DataType(pointer.buffer.datatype), pointer.normalize, int32(pointer.stride), int32(pointer.start))
	pointer.buffer.unbind()
	return true
}

func (program *Program) UniformFloat(name string, value float32) bool {
	index, ok := program.uniformIndexByName[name]
	if !ok {
		return false
	}
	gl.ProgramUniform1f(program.program.Id(), int32(index), value)
	return true
}

func (program *Program) Delete() {
	if program == nil {
		return
	}
	program.program.Delete()
	program.vao.Delete()
}

func (program *Program) Bind() {
	gli.BindVertexArrayObject(program.vao)
	gli.UseProgram(program.program)
}

func (program *Program) Unbind() {
	gli.UseNoProgram()
	gli.UnbindVertexArrayObject()
}
