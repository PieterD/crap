package raw

const (
	ShaderTypeVertex int = iota + 1
	ShaderTypeGeometry
	ShaderTypeFragment
)

type Raw interface {
	Init() error
	Viewport(x, y, width, height int)

	ShaderCreate(iShadertype int) (shaderid uint32, err error)
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
}
