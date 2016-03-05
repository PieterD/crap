package glimmer

import "github.com/PieterD/crap/glimmer/gli"

type Shader struct {
	shader     gli.Shader
	shaderType gli.ShaderType
}

func CreateVertexShader(source ...string) (*Shader, error) {
	return createShader(gli.VertexShader, source...)
}

func CreateGeometryShader(source ...string) (*Shader, error) {
	return createShader(gli.GeometryShader, source...)
}

func CreateFragmentShader(source ...string) (*Shader, error) {
	return createShader(gli.FragmentShader, source...)
}

func CreateComputeShader(source ...string) (*Shader, error) {
	return createShader(gli.ComputeShader, source...)
}

func CreateTessControlShader(source ...string) (*Shader, error) {
	return createShader(gli.TessControlShader, source...)
}

func CreateTessEvaluationShader(source ...string) (*Shader, error) {
	return createShader(gli.TessEvaluationShader, source...)
}

func createShader(shaderType gli.ShaderType, source ...string) (*Shader, error) {
	shader := gli.CreateShader(shaderType)
	if !shader.Valid() {
		return nil, GetError()
	}
	shader.Source(source)
	shader.Compile()
	if !shader.GetCompileSuccess() {
		loglength := shader.GetInfoLogLength()
		log := make([]byte, loglength)
		log = shader.GetInfoLog(log)
		shader.Delete()
		return nil, &ShaderError{Desc: string(log)}
	}
	return &Shader{shader: shader, shaderType: shaderType}, nil
}

func (shader *Shader) Delete() {
	if shader == nil {
		return
	}
	shader.shader.Delete()
}
