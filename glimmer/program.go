package glimmer

import (
	"github.com/PieterD/crap/glimmer/gli"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Program struct {
	program    gli.Program
	attributes []gli.ProgramAttribute
	uniforms   []gli.ProgramUniform
	vao        gli.VertexArrayObject

	attributeIndexByName map[string]uint32
	uniformIndexByName   map[string]uint32
}

func CreateProgram(shaders ...gli.Shader) (*Program, error) {
	p, err := gli.CreateProgram(shaders...)
	if err != nil {
		return nil, err
	}

	program := &Program{program: p}
	program.attributes, err = p.Attributes()
	if err != nil {
		p.Delete()
		return nil, err
	}
	program.attributeIndexByName = make(map[string]uint32)
	for _, attribute := range program.attributes {
		program.attributeIndexByName[attribute.Name] = attribute.Index
	}

	program.uniforms, err = p.Uniforms()
	if err != nil {
		p.Delete()
		return nil, err
	}
	program.uniformIndexByName = make(map[string]uint32)
	for _, uniform := range program.uniforms {
		program.uniformIndexByName[uniform.Name] = uniform.Index
	}

	program.vao = gli.CreateVertexArrayObject()

	return program, nil
}

func (program *Program) AttributeByName(name string, pointer gli.DataPointer) bool {
	index, ok := program.attributeIndexByName[name]
	if !ok {
		return false
	}
	return program.AttributeByIndex(index, pointer)
}

func (program *Program) AttributeByIndex(index uint32, pointer gli.DataPointer) bool {
	gli.BindBuffer(gli.ArrayBuffer, pointer.Buffer)
	program.vao.EnableAttrib(index)
	program.vao.AttribPointer(index, int32(pointer.Components), gli.DataType(pointer.Type), pointer.Normalize, int32(pointer.Stride), int32(pointer.Start))
	gli.UnbindBuffer(gli.ArrayBuffer)
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
	gli.BindProgram(program.program)
}

func (program *Program) Unbind() {
	gli.UnbindProgram()
	gli.UnbindVertexArrayObject()
}
