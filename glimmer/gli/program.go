package gli

import (
	"fmt"

	"github.com/PieterD/crap/glimmer/convc"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type ProgramParameter uint32

const (
	PROGRAM_DELETE_STATUS       ProgramParameter = gl.DELETE_STATUS
	LINK_STATUS                 ProgramParameter = gl.LINK_STATUS
	VALIDATE_STATUS             ProgramParameter = gl.VALIDATE_STATUS
	PROGRAM_INFO_LOG_LENGTH     ProgramParameter = gl.INFO_LOG_LENGTH
	ATTACHED_SHADERS            ProgramParameter = gl.ATTACHED_SHADERS
	ACTIVE_ATTRIBUTES           ProgramParameter = gl.ACTIVE_ATTRIBUTES
	ACTIVE_ATTRIBUTE_MAX_LENGTH ProgramParameter = gl.ACTIVE_ATTRIBUTE_MAX_LENGTH
	ACTIVE_UNIFORMS             ProgramParameter = gl.ACTIVE_UNIFORMS
	ACTIVE_UNIFORM_MAX_LENGTH   ProgramParameter = gl.ACTIVE_UNIFORM_MAX_LENGTH
)

type UniformParameter uint32

const (
	UNIFORM_TYPE                        UniformParameter = gl.UNIFORM_TYPE
	UNIFORM_SIZE                        UniformParameter = gl.UNIFORM_SIZE
	UNIFORM_NAME_LENGTH                 UniformParameter = gl.UNIFORM_NAME_LENGTH
	UNIFORM_BLOCK_INDEX                 UniformParameter = gl.UNIFORM_BLOCK_INDEX
	UNIFORM_OFFSET                      UniformParameter = gl.UNIFORM_OFFSET
	UNIFORM_ARRAY_STRIDE                UniformParameter = gl.UNIFORM_ARRAY_STRIDE
	UNIFORM_MATRIX_STRIDE               UniformParameter = gl.UNIFORM_MATRIX_STRIDE
	UNIFORM_IS_ROW_MAJOR                UniformParameter = gl.UNIFORM_IS_ROW_MAJOR
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX UniformParameter = gl.UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX
)

type DataType uint32

type iProgram struct {
	id uint32
}

type Program interface {
	Id() uint32
	Delete()
	GetIV(param ProgramParameter) int32
	GetActiveAttrib(index uint32, buf []byte) (name []byte, datatype DataType, size int)
	GetAttribLocationBytes(name []byte) int32
	GetAttribLocation(name string) int32
	GetActiveUniformName(index uint32, buf []byte) []byte
	GetUniformLocationBytes(name []byte) int32
	GetUniformLocation(name string) int32
	GetActiveUniformIV(param UniformParameter, index uint32) int32
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

func (program iProgram) GetActiveAttrib(index uint32, buf []byte) (name []byte, datatype DataType, size int) {
	var length int32
	var isize int32
	var idatatype uint32
	gl.GetActiveAttrib(program.id, index, int32(len(buf)), &length, &isize, &idatatype, &buf[0])
	return buf[:length : length+1], DataType(idatatype), int(isize)
}

func (program iProgram) GetAttribLocationBytes(name []byte) int32 {
	return gl.GetAttribLocation(program.id, &name[0])
}

func (program iProgram) GetAttribLocation(name string) int32 {
	ptr, free := convc.StringToC(name)
	defer free()
	return gl.GetAttribLocation(program.id, ptr)
}

func (program iProgram) GetActiveUniformName(index uint32, buf []byte) []byte {
	var length int32
	gl.GetActiveUniformName(program.id, index, int32(len(buf)), &length, &buf[0])
	return buf[:length : length+1]
}

func (program iProgram) GetUniformLocationBytes(name []byte) int32 {
	return gl.GetUniformLocation(program.id, &name[0])
}

func (program iProgram) GetUniformLocation(name string) int32 {
	ptr, free := convc.StringToC(name)
	defer free()
	return gl.GetUniformLocation(program.id, ptr)
}

func (program iProgram) GetActiveUniformIV(param UniformParameter, index uint32) int32 {
	var pi int32
	gl.GetActiveUniformsiv(program.id, 1, &index, uint32(param), &pi)
	return pi
}
