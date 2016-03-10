package gli

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type UniformCollection struct {
	program Program
	byName  map[string]int
	byIndex map[uint32]int
	list    []ProgramUniform
}

type ProgramUniform struct {
	Program Program
	Name    string
	Index   uint32
	Type    DataType
	Size    uint32
}

func (program iProgram) getActiveUniform(index uint32, buf []byte) (name []byte, datatype DataType, size int) {
	var length int32
	var isize int32
	var idatatype uint32
	gl.GetActiveUniform(program.id, index, int32(len(buf)), &length, &isize, &idatatype, &buf[0])
	return buf[:length : length+1], DataType(idatatype), int(isize)
}

func (program iProgram) Uniforms() UniformCollection {
	max := int(program.GetIV(ACTIVE_UNIFORMS))
	byname := make(map[string]int, max)
	byindex := make(map[uint32]int, max)
	uniforms := make([]ProgramUniform, 0, max)
	buf := make([]byte, program.GetIV(ACTIVE_UNIFORM_MAX_LENGTH))
	for i := 0; i < max; i++ {
		namebytes, datatype, arraysize := program.getActiveUniform(uint32(i), buf)
		name := string(namebytes)
		location := gl.GetUniformLocation(program.id, &namebytes[0])
		if location <= -1 {
			continue
		}
		index := uint32(location)
		byname[name] = len(uniforms)
		byindex[index] = len(uniforms)
		uniforms = append(uniforms, ProgramUniform{
			Program: program,
			Name:    name,
			Index:   index,
			Type:    DataType(datatype),
			Size:    uint32(arraysize),
		})
	}
	return UniformCollection{
		program: program,
		byName:  byname,
		byIndex: byindex,
		list:    uniforms,
	}
}

func (coll UniformCollection) List() []ProgramUniform {
	return coll.list
}

func (coll UniformCollection) ByIndex(index uint32) ProgramUniform {
	i, ok := coll.byIndex[index]
	if ok {
		return coll.list[i]
	}
	return ProgramUniform{}
}

func (coll UniformCollection) ByName(name string) ProgramUniform {
	i, ok := coll.byName[name]
	if ok {
		return coll.list[i]
	}
	return ProgramUniform{}
}

func (uni ProgramUniform) Valid() bool {
	return uni.Size > 0
}

func (uni ProgramUniform) Float(f ...float32) {
	switch len(f) {
	case 1:
		gl.ProgramUniform1f(uni.Program.Id(), int32(uni.Index), f[0])
	case 2:
		gl.ProgramUniform2f(uni.Program.Id(), int32(uni.Index), f[0], f[1])
	case 3:
		gl.ProgramUniform3f(uni.Program.Id(), int32(uni.Index), f[0], f[1], f[2])
	case 4:
		gl.ProgramUniform4f(uni.Program.Id(), int32(uni.Index), f[0], f[1], f[2], f[3])
	default:
		panic(fmt.Errorf("ProgramUniform.Float must be passed between 1 to 4 floats inclusive"))
	}
}
