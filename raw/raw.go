package raw

type Enum byte

const (
	ShaderTypeVertex Enum = iota + 1
	ShaderTypeGeometry
	ShaderTypeFragment
)

const (
	DataTypeFloat Enum = iota + 1
	DataTypeFloatVec2
	DataTypeFloatVec3
	DataTypeFloatVec4
	DataTypeFloatMat2
	DataTypeFloatMat3
	DataTypeFloatMat4
	DataTypeFloatMat2x3
	DataTypeFloatMat2x4
	DataTypeFloatMat3x2
	DataTypeFloatMat3x4
	DataTypeFloatMat4x2
	DataTypeFloatMat4x3
	DataTypeInt
	DataTypeIntVec2
	DataTypeIntVec3
	DataTypeIntVec4
	DataTypeUnsignedInt
	DataTypeUnsignedIntVec2
	DataTypeUnsignedIntVec3
	DataTypeUnsignedIntVec4
	DataTypeDouble
	DataTypeDoubleVec2
	DataTypeDoubleVec3
	DataTypeDoubleVec4
	DataTypeDoubleMat2
	DataTypeDoubleMat3
	DataTypeDoubleMat4
	DataTypeDoubleMat2x3
	DataTypeDoubleMat2x4
	DataTypeDoubleMat3x2
	DataTypeDoubleMat3x4
	DataTypeDoubleMat4x2
	DataTypeDoubleMat4x3
)

type Raw interface {
	Init() error
	Viewport(x, y, width, height int)
	ClearColor(r, g, b, a float32)

	ShaderCreate(iShadertype Enum) (shaderid uint32, err error)
	ShaderDelete(shaderid uint32)
	ShaderSource(shaderid uint32, source ...string)
	ShaderCompile(shaderid uint32)
	ShaderCompileStatus(shaderid uint32) bool
	ShaderInfoLogLength(shaderid uint32) int
	ShaderInfoLog(shaderid uint32, buf []byte) []byte

	ProgramCreate() (programid uint32, err error)
	ProgramDelete(programid uint32)
	ProgramAttachShader(programid uint32, shaderid uint32)
	ProgramLink(programid uint32)
	ProgramLinkStatus(programid uint32) bool
	ProgramInfoLogLength(programid uint32) int
	ProgramInfoLog(programid uint32, buf []byte) []byte

	ProgramAttributeNum(programid uint32) int
	ProgramAttributeMaxLength(programid uint32) int
	ProgramAttribute(programid uint32, index int, buf []byte) (namebytes []byte, datatype Enum, size int)
	ProgramAttributeLocation(programid uint32, namebytes []byte) (location int, ok bool)

	ProgramUniformNum(programid uint32) int
	ProgramUniformMaxLength(programid uint32) int
	ProgramUniform(programid uint32, index int, buf []byte) (namebytes []byte, datatype Enum, size int)
	ProgramUniformLocation(programid uint32, namebytes []byte) (location int, ok bool)
}
