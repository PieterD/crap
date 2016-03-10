package glimmer

import (
	"github.com/PieterD/crap/glimmer/gli"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Program struct {
	program    gli.Program
	attributes gli.AttributeCollection
	uniforms   gli.UniformCollection
	vao        gli.VertexArrayObject
}

func CreateProgram(shaders ...gli.Shader) (*Program, error) {
	p, err := gli.CreateProgram(shaders...)
	if err != nil {
		return nil, err
	}

	program := &Program{program: p}
	program.attributes = p.Attributes()
	program.uniforms = p.Uniforms()
	program.vao = gli.CreateVertexArrayObject()

	return program, nil
}

func (program *Program) AttributeByName(name string, pointer gli.DataPointer) bool {
	attr := program.attributes.ByName(name)
	if !attr.Valid() {
		return false
	}
	return program.AttributeByIndex(attr.Index, pointer)
}

func (program *Program) AttributeByIndex(index uint32, pointer gli.DataPointer) bool {
	attr := program.attributes.ByIndex(index)
	program.vao.Enable(attr, pointer)
	return true
}

func (program *Program) UniformFloat(name string, value float32) bool {
	attr := program.uniforms.ByName(name)
	if !attr.Valid() {
		return false
	}
	gl.ProgramUniform1f(program.program.Id(), int32(attr.Index), value)
	return true
}

func (program *Program) UniformFloat2(name string, x float32, y float32) bool {
	attr := program.uniforms.ByName(name)
	if !attr.Valid() {
		return false
	}
	gl.ProgramUniform2f(program.program.Id(), int32(attr.Index), x, y)
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
