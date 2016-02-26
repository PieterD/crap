package glimmer

import (
	"fmt"

	"github.com/PieterD/crap/glimmer/convc"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Shader struct {
	id         uint32
	shaderType ShaderType
}

type ShaderType uint32

const (
	VertexShader         ShaderType = gl.VERTEX_SHADER
	GeometryShader       ShaderType = gl.GEOMETRY_SHADER
	FragmentShader       ShaderType = gl.FRAGMENT_SHADER
	ComputeShader        ShaderType = gl.COMPUTE_SHADER
	TessControlShader    ShaderType = gl.TESS_CONTROL_SHADER
	TessEvaluationShader ShaderType = gl.TESS_EVALUATION_SHADER
)

func CreateShader(shaderType ShaderType, source ...string) (*Shader, error) {
	shader := new(Shader)
	shader.id = gl.CreateShader(uint32(shaderType))
	if shader.id == 0 {
		return nil, GetError()
	}
	// ptr, free := gl.Strs(source...)
	ptr, free := convc.MultiStringToC(source...)
	defer free()
	gl.ShaderSource(shader.id, 1, ptr, nil)
	gl.CompileShader(shader.id)
	var status int32
	gl.GetShaderiv(shader.id, gl.COMPILE_STATUS, &status)
	fmt.Printf("status: %v\n", status)
	if status == gl.FALSE {
		var loglength int32
		gl.GetShaderiv(shader.id, gl.INFO_LOG_LENGTH, &loglength)
		log := make([]byte, loglength)
		var logptr *uint8 = &log[0]
		gl.GetShaderInfoLog(shader.id, loglength, nil, logptr)
		gl.DeleteShader(shader.id)
		return nil, &ShaderError{Desc: string(log[:len(log)-1])}
	}
	return shader, nil
}

func (shader *Shader) Delete() {
	gl.DeleteShader(shader.id)
}
