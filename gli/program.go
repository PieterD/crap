package gli

import (
	"fmt"

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

const (
	GlByte                 DataType = gl.BYTE
	GlUByte                DataType = gl.UNSIGNED_BYTE
	GlShort                DataType = gl.SHORT
	GlUShort               DataType = gl.UNSIGNED_SHORT
	GlInt                  DataType = gl.INT
	GlUInt                 DataType = gl.UNSIGNED_INT
	GlFloat                DataType = gl.FLOAT
	GlHalfFloat            DataType = gl.HALF_FLOAT
	GlFixed                DataType = gl.FIXED
	GlInt_2_10_10_10_REV   DataType = gl.INT_2_10_10_10_REV
	GlUInt_2_10_10_10_REV  DataType = gl.UNSIGNED_INT_2_10_10_10_REV
	GlUInt_10F_11F_11F_REV DataType = gl.UNSIGNED_INT_10F_11F_11F_REV
	GlDouble               DataType = gl.DOUBLE
)

type iProgram struct {
	id uint32
}

type Program interface {
	Id() uint32
	Delete()
	GetIV(param ProgramParameter) int32
	GetActiveUniformIV(param UniformParameter, index uint32) int32
	Attributes() ([]ProgramAttribute, error)
	Uniforms() ([]ProgramUniform, error)
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

func (program iProgram) getActiveAttrib(index uint32, buf []byte) (name []byte, datatype DataType, size int) {
	var length int32
	var isize int32
	var idatatype uint32
	gl.GetActiveAttrib(program.id, index, int32(len(buf)), &length, &isize, &idatatype, &buf[0])
	return buf[:length : length+1], DataType(idatatype), int(isize)
}

type ProgramAttribute struct {
	Name  string
	Index uint32
	Type  DataType
	Size  uint32
}

func (program iProgram) Attributes() ([]ProgramAttribute, error) {
	attributes := make([]ProgramAttribute, program.GetIV(ACTIVE_ATTRIBUTES))
	buf := make([]byte, program.GetIV(ACTIVE_ATTRIBUTE_MAX_LENGTH))
	for i := range attributes {
		namebytes, datatype, size := program.getActiveAttrib(uint32(i), buf)
		name := string(namebytes)
		location := gl.GetAttribLocation(program.id, &namebytes[0])
		if location == -1 {
			//TODO: ShaderError
			return nil, fmt.Errorf("Attribute location for '%s' not found", name)
		}
		attributes[i] = ProgramAttribute{
			Name:  name,
			Index: uint32(location),
			Type:  datatype,
			Size:  uint32(size),
		}
	}
	return attributes, nil
}

type ProgramUniform struct {
	Name  string
	Index uint32
	Type  DataType
	Size  uint32
}

func (program iProgram) getActiveUniform(index uint32, buf []byte) (name []byte, datatype DataType, size int) {
	var length int32
	var isize int32
	var idatatype uint32
	gl.GetActiveUniform(program.id, index, int32(len(buf)), &length, &isize, &idatatype, &buf[0])
	return buf[:length : length+1], DataType(idatatype), int(isize)
}

func (program iProgram) Uniforms() ([]ProgramUniform, error) {
	uniforms := make([]ProgramUniform, program.GetIV(ACTIVE_UNIFORMS))
	buf := make([]byte, program.GetIV(ACTIVE_UNIFORM_MAX_LENGTH))
	for i := range uniforms {
		namebytes, datatype, arraysize := program.getActiveUniform(uint32(i), buf)
		name := string(namebytes)
		location := gl.GetUniformLocation(program.id, &namebytes[0])
		if location == -1 {
			//TODO: ShaderError
			return nil, fmt.Errorf("Uniform location for '%s' not found", name)
		}
		uniforms[i] = ProgramUniform{
			Name:  name,
			Index: uint32(location),
			Type:  DataType(datatype),
			Size:  uint32(arraysize),
		}
	}
	return uniforms, nil
}
