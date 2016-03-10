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
	attr, ok := program.attributes.ByName(name)
	if !ok {
		return false
	}
	return program.AttributeByIndex(attr.Index, pointer)
}

func (program *Program) AttributeByIndex(index uint32, pointer gli.DataPointer) bool {
	gli.BindBuffer(gli.ArrayBuffer, pointer.Buffer)
	program.vao.EnableAttrib(index)
	program.vao.AttribPointer(index, int32(pointer.Components), gli.DataType(pointer.Type), pointer.Normalize, int32(pointer.Stride), int32(pointer.Start))
	gli.UnbindBuffer(gli.ArrayBuffer)
	return true
}

func (program *Program) UniformFloat(name string, value float32) bool {
	attr, ok := program.uniforms.ByName(name)
	if !ok {
		return false
	}
	gl.ProgramUniform1f(program.program.Id(), int32(attr.Index), value)
	return true
}

func (program *Program) UniformFloat2(name string, x float32, y float32) bool {
	attr, ok := program.uniforms.ByName(name)
	if !ok {
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
