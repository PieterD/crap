package gli

import (
	"fmt"

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
	id         uint32
	shadertype ShaderType
}

type Shader interface {
	Id() uint32
	Delete()
	GetIV(param ShaderParameter) int32
}

func (context iContext) CreateShader(shaderType ShaderType, source ...string) (Shader, error) {
	id := gl.CreateShader(uint32(shaderType))
	shader := iShader{id: id, shadertype: shaderType}
	if id == 0 {
		// TODO: GetError
		return nil, fmt.Errorf("Unable to allocate shader")
	}
	ptr, free := convc.MultiStringToC(source...)
	defer free()
	gl.ShaderSource(id, int32(len(source)), ptr, nil)
	gl.CompileShader(shader.id)
	result := shader.GetIV(COMPILE_STATUS)
	if result == int32(FALSE) {
		loglength := shader.GetIV(INFO_LOG_LENGTH)
		log := make([]byte, loglength)
		var length int32
		gl.GetShaderInfoLog(id, loglength, &length, &log[0])
		shader.Delete()
		// TODO: ShaderError
		return nil, fmt.Errorf("Unsable to compile shader: %s", log[:length])
	}

	return shader, nil
}

func (shader iShader) Delete() {
	gl.DeleteShader(shader.id)
}

func (shader iShader) Id() uint32 {
	return shader.id
}

func (shader iShader) GetIV(param ShaderParameter) int32 {
	var pi int32
	gl.GetShaderiv(shader.id, uint32(param), &pi)
	return pi
}

func (shader iShader) GetInfoLogLength() int32 {
	return shader.GetIV(INFO_LOG_LENGTH)
}
