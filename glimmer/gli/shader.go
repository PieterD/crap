package gli

import (
	"github.com/PieterD/crap/glimmer/convc"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type ShaderType uint32

const (
	VertexShader         ShaderType = gl.VERTEX_SHADER
	GeometryShader       ShaderType = gl.GEOMETRY_SHADER
	FragmentShader       ShaderType = gl.FRAGMENT_SHADER
	ComputeShader        ShaderType = gl.COMPUTE_SHADER
	TessControlShader    ShaderType = gl.TESS_CONTROL_SHADER
	TessEvaluationShader ShaderType = gl.TESS_EVALUATION_SHADER
)

type ShaderParameter uint32

const (
	SHADER_TYPE          ShaderParameter = gl.SHADER_TYPE
	SHADER_DELETE_STATUS ShaderParameter = gl.DELETE_STATUS
	COMPILE_STATUS       ShaderParameter = gl.COMPILE_STATUS
	INFO_LOG_LENGTH      ShaderParameter = gl.INFO_LOG_LENGTH
	SHADER_SOURCE_LENGTH ShaderParameter = gl.SHADER_SOURCE_LENGTH
)

type Bool uint32

const (
	TRUE  Bool = gl.TRUE
	FALSE Bool = gl.FALSE
)

type iShader struct {
	id uint32
}

type Shader interface {
	Id() uint32
	Valid() bool
	Delete()
	Source(sources []string)
	Compile()
	GetIV(param ShaderParameter) int32
	GetCompileSuccess() bool
	GetInfoLogLength() int32
	GetInfoLog(log []byte) []byte
}

func (context iContext) CreateShader(shaderType ShaderType) Shader {
	id := gl.CreateShader(uint32(shaderType))
	return iShader{id}
}

func (shader iShader) Id() uint32 {
	return shader.id
}

func (shader iShader) Valid() bool {
	return shader.id != 0
}

func (shader iShader) Delete() {
	gl.DeleteShader(shader.id)
}

func (shader iShader) Source(sources []string) {
	ptr, free := convc.MultiStringToC(sources...)
	defer free()
	gl.ShaderSource(shader.id, int32(len(sources)), ptr, nil)
}

func (shader iShader) Compile() {
	gl.CompileShader(shader.id)
}

func (shader iShader) GetIV(param ShaderParameter) int32 {
	var pi int32
	gl.GetShaderiv(shader.id, uint32(param), &pi)
	return pi
}

func (shader iShader) GetCompileSuccess() bool {
	result := shader.GetIV(COMPILE_STATUS)
	if result == int32(FALSE) {
		return false
	}
	return true
}

func (shader iShader) GetInfoLogLength() int32 {
	return shader.GetIV(INFO_LOG_LENGTH)
}

func (shader iShader) GetInfoLog(buf []byte) []byte {
	bufsize := int32(len(buf))
	logptr := &buf[0]
	var length int32
	gl.GetShaderInfoLog(shader.id, bufsize, &length, logptr)
	return buf[:length : length+1]
}
