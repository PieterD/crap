package glimmer

import "github.com/go-gl/gl/v3.3-core/gl"

type Shader struct {
	id         uint32
	shaderType ShaderType
}

type ShaderType uint32

const (
	VertexShader   ShaderType = gl.VERTEX_SHADER
	GeometryShader ShaderType = gl.GEOMETRY_SHADER
	FragmentShader ShaderType = gl.FRAGMENT_SHADER
)

func CreateShader(shaderType ShaderType, source []byte) (*Shader, error) {
	shader := new(Shader)
	shader.id = gl.CreateShader(uint32(shaderType))
	if shader.id == 0 {
		return nil, GetError()
	}
	var ptr *uint8 = &source[0]
	var length int32 = int32(len(source))
	gl.ShaderSource(shader.id, 1, &ptr, &length)
	gl.CompileShader(shader.id)
	var status int32
	gl.GetShaderiv(shader.id, gl.INFO_LOG_LENGTH, &status)
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
