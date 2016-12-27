package gli

type ProgramCollection interface {
	Program(name string) (ProgramDef, error)
}

type ShaderDef struct {
	Type   ShaderType
	Source []string
}

type BufferDef struct {
	Name       string
	Usage      BufferUsage
	Attributes []ProgramResource
}

type ProgramDef struct {
	Name      string
	Group     string
	Locations []AttributeLocation
	Buffers   []BufferDef
	Shaders   []ShaderDef
}
