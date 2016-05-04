package gli

//go:generate stringer -type=ShaderType
type ShaderType uint32

const (
	ShaderTypeVertex ShaderType = iota + 1
	ShaderTypeGeometry
	ShaderTypeFragment
)

type ShaderId uint32

type ProgramId uint32
