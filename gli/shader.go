package gli

import (
	"fmt"

	"github.com/PieterD/glimmer/raw"
)

type Shader struct {
	ctx *Context
	id  uint32
	typ iShaderType
}

type iShaderType struct {
	t int
}

var (
	VertexShader   = iShaderType{raw.ShaderTypeVertex}
	GeometryShader = iShaderType{raw.ShaderTypeGeometry}
	FragmentShader = iShaderType{raw.ShaderTypeFragment}
)

func (ctx *Context) NewShader(typ iShaderType, sources ...string) (*Shader, error) {
	shaderid, err := ctx.r.ShaderCreate(typ.t)
	if err != nil {
		return nil, fmt.Errorf("Shader creation error: %v", err)
	}
	ctx.r.ShaderSource(shaderid, sources...)
	ctx.r.ShaderCompile(shaderid)
	if !ctx.r.ShaderCompileStatus(shaderid) {
		loglength := ctx.r.ShaderInfoLogLength(shaderid)
		log := make([]byte, loglength)
		log = ctx.r.ShaderInfoLog(shaderid, log)
		ctx.r.ShaderDelete(shaderid)
		return nil, fmt.Errorf("Shader compilation error: %s", log)
	}
	return &Shader{
		ctx: ctx,
		id:  shaderid,
		typ: typ,
	}, nil
}

func (shader *Shader) Delete() {
	shader.ctx.r.ShaderDelete(shader.id)
}
