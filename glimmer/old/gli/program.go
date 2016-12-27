package gli

import "fmt"

type Program struct {
	ctx  *Context
	id   uint32
	attr *programAttributeCollection
}

func (ctx *Context) NewProgram(shaders ...*Shader) (*Program, error) {
	programid, err := ctx.r.ProgramCreate()
	if err != nil {
		return nil, fmt.Errorf("Program creation error: %v", err)
	}
	var buf []byte
	for name, attr := range ctx.attributeIndexMap {
		buf = append(buf[:0], name...)
		buf = append(buf, 0)
		ctx.r.ProgramAttributeLocationBind(programid, attr.index, buf)
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
	prog := &Program{
		ctx: ctx,
		id:  programid,
	}
	prog.attributes()
	return prog, nil
}

func (program *Program) Delete() {
	program.ctx.r.ProgramDelete(program.id)
}

type programAttributeCollection struct {
	program *Program
	nameMap map[string]int
	list    []*programAttribute
}

type programAttribute struct {
	program   *Program
	name      string
	location  int
	datatype  iDataType
	arraysize int
}

func (program *Program) attributes() {
	max := program.ctx.r.ProgramAttributeNum(program.id)
	nameMap := make(map[string]int, max)
	attributes := make([]*programAttribute, 0, max)
	buf := make([]byte, program.ctx.r.ProgramAttributeMaxLength(program.id))
	for i := 0; i < max; i++ {
		namebytes, datatype, size := program.ctx.r.ProgramAttribute(program.id, i, buf)
		location, ok := program.ctx.r.ProgramAttributeLocation(program.id, namebytes)
		if !ok {
			continue
		}
		name := string(namebytes)
		nameMap[name] = len(attributes)
		attributes = append(attributes, &programAttribute{
			program:   program,
			name:      name,
			location:  location,
			datatype:  iDataType{datatype},
			arraysize: size,
		})
	}
	program.attr = &programAttributeCollection{
		program: program,
		list:    attributes,
		nameMap: nameMap,
	}
}

func (coll *programAttributeCollection) byName(name string) *programAttribute {
	i, ok := coll.nameMap[name]
	if ok {
		return coll.list[i]
	}
	return nil
}

func (attr *programAttribute) typ() (datatype iDataType, arraysize int) {
	return attr.datatype, attr.arraysize
}

type ProgramUniformCollection struct {
	program *Program
	byName  map[string]int
	list    []*ProgramUniform
}

type ProgramUniform struct {
	program   *Program
	name      string
	location  int
	datatype  iDataType
	arraysize int
}

func (program *Program) Uniforms() *ProgramUniformCollection {
	max := program.ctx.r.ProgramUniformNum(program.id)
	byname := make(map[string]int, max)
	uniforms := make([]*ProgramUniform, 0, max)
	buf := make([]byte, program.ctx.r.ProgramUniformMaxLength(program.id))
	for i := 0; i < max; i++ {
		namebytes, datatype, arraysize := program.ctx.r.ProgramUniform(program.id, i, buf)
		location, ok := program.ctx.r.ProgramUniformLocation(program.id, namebytes)
		if !ok {
			continue
		}
		name := string(namebytes)
		byname[name] = len(uniforms)
		uniforms = append(uniforms, &ProgramUniform{
			program:   program,
			name:      name,
			location:  location,
			datatype:  iDataType{datatype},
			arraysize: arraysize,
		})
	}
	return &ProgramUniformCollection{
		program: program,
		list:    uniforms,
		byName:  byname,
	}
}

func (coll *ProgramUniformCollection) ByName(name string) *ProgramUniform {
	i, ok := coll.byName[name]
	if ok {
		return coll.list[i]
	}
	return nil
}

func (attr *ProgramUniform) Type() (datatype iDataType, arraysize int) {
	return attr.datatype, attr.arraysize
}
