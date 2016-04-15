package gli

import "fmt"

type Program struct {
	ctx *Context
	id  uint32
}

func (ctx *Context) NewProgram(shaders ...*Shader) (*Program, error) {
	programid, err := ctx.r.ProgramCreate()
	if err != nil {
		return nil, fmt.Errorf("Program creation error: %v", err)
	}
	for _, shader := range shaders {
		ctx.r.ProgramAttachShader(programid, shader.id)
	}
	ctx.r.ProgramLink(programid)
	if !ctx.r.ProgramLinkStatus(programid) {
		loglength := ctx.r.ProgramInfoLogLength(programid)
		log := make([]byte, loglength)
		log = ctx.r.ProgramInfoLog(programid, log)
		ctx.r.ProgramDelete(programid)
		return nil, fmt.Errorf("Program link error: %s", log)
	}
	return &Program{
		ctx: ctx,
		id:  programid,
	}, nil
}
