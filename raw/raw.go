package raw

import "unsafe"

type Raw interface {
	Init() error
	Viewport(x, y, width, height int)
	ClearColor(r, g, b, a float32)

	ShaderCreate(iShadertype Enum) (shaderid uint32, err error)
	ShaderDelete(shaderid uint32)
	ShaderSource(shaderid uint32, source ...string)
	ShaderCompile(shaderid uint32)
	ShaderCompileStatus(shaderid uint32) (ok bool)
	ShaderInfoLogLength(shaderid uint32) int
	ShaderInfoLog(shaderid uint32, buf []byte) []byte

	ProgramCreate() (programid uint32, err error)
	ProgramDelete(programid uint32)
	ProgramAttachShader(programid uint32, shaderid uint32)
	ProgramLink(programid uint32)
	ProgramLinkStatus(programid uint32) (ok bool)
	ProgramInfoLogLength(programid uint32) int
	ProgramInfoLog(programid uint32, buf []byte) []byte

	ProgramAttributeNum(programid uint32) int
	ProgramAttributeMaxLength(programid uint32) int
	ProgramAttribute(programid uint32, index int, buf []byte) (namebytes []byte, datatype Enum, size int)
	ProgramAttributeLocation(programid uint32, namebytes []byte) (location int, ok bool)
	ProgramAttributeLocationBind(programid uint32, index int, namebytes []byte)

	ProgramUniformNum(programid uint32) int
	ProgramUniformMaxLength(programid uint32) int
	ProgramUniform(programid uint32, index int, buf []byte) (namebytes []byte, datatype Enum, size int)
	ProgramUniformLocation(programid uint32, namebytes []byte) (location int, ok bool)

	ArrayBufferCreate() (bufferid uint32)
	ArrayBufferDelete(bufferid uint32)
	ArrayBufferBind(bufferid uint32)
	ArrayBufferData(bufferid uint32, bytenum int, ptr unsafe.Pointer, accesstype Enum)
	ArrayBufferSubData(bufferid uint32, offset int, bytes int, ptr unsafe.Pointer)
	ArrayBufferUnbind()

	VertexArrayCreate() (vaoid uint32)
	VertexArrayDelete(vaoid uint32)
	VertexArrayBind(vaoid uint32)
	VertexArrayEnable(idx int)
	VertexArrayDisable(idx int)
}
