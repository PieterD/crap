package gli

import (
	"fmt"
	"sort"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type UniformCollection interface {
	List() []Uniform
	ByIndex(index uint32) Uniform
	ByName(name string) Uniform
	Block(name string) UniformBlock
}

type iUniformCollection struct {
	program     Program
	byName      map[string]int
	byIndex     map[uint32]int
	blockByName map[string]int
	list        []Uniform
	blocks      []UniformBlock
	members     []UniformBlockMember
}

type Uniform struct {
	Program Program
	Name    string
	Index   uint32
	Type    DataType
	Size    uint32
}

type UniformBlock struct {
	Program  Program
	Name     string
	Index    uint32
	Uniforms []UniformBlockMember
	Size     uint32
}

type UniformBlockMember struct {
	Program  Program
	Name     string
	Type     DataType
	Size     uint32
	Block    uint32
	Offset   uint32
	AStride  uint32
	MStride  uint32
	RowMajor bool
}

func (program iProgram) Uniforms() UniformCollection {
	list, members := program.uniforms()
	blocks := program.uniformBlocks()
	coll := iUniformCollection{
		program: program,
		list:    list,
		blocks:  blocks,
		members: members,
	}
	coll.byName = make(map[string]int, len(coll.list))
	coll.byIndex = make(map[uint32]int, len(coll.list))
	coll.blockByName = make(map[string]int, len(coll.blocks))
	for i := range coll.list {
		coll.byName[coll.list[i].Name] = i
		coll.byIndex[coll.list[i].Index] = i
	}
	for i := range coll.blocks {
		coll.blockByName[coll.blocks[i].Name] = i
		index := coll.blocks[i].Index
		fst := -1
		lst := len(coll.members)
		for j := range coll.members {
			if coll.members[j].Block > index {
				lst = j
				break
			}
			if coll.members[j].Block == index {
				if fst == -1 {
					fst = j
				}
			}
		}
		coll.blocks[i].Uniforms = coll.members[fst:lst]
	}
	return coll
}

func (program iProgram) getActiveUniform(index uint32, buf []byte) (name []byte, datatype DataType, size int) {
	var length int32
	var isize int32
	var idatatype uint32
	gl.GetActiveUniform(program.id, index, int32(len(buf)), &length, &isize, &idatatype, &buf[0])
	return buf[:length : length+1], DataType(idatatype), int(isize)
}

func (program iProgram) uniforms() ([]Uniform, []UniformBlockMember) {
	max := uint32(program.GetIV(ACTIVE_UNIFORMS))
	list := make([]Uniform, 0, max)
	members := make([]UniformBlockMember, 0, max)

	buf := make([]byte, program.GetIV(ACTIVE_UNIFORM_MAX_LENGTH))
	for i := uint32(0); i < max; i++ {
		namebytes, datatype, arraysize := program.getActiveUniform(i, buf)
		location := gl.GetUniformLocation(program.id, &namebytes[0])
		name := string(namebytes)
		if location >= 0 {
			list = append(list, Uniform{
				Program: program,
				Name:    name,
				Index:   uint32(location),
				Type:    DataType(datatype),
				Size:    uint32(arraysize),
			})
			continue
		}
		var block int32
		gl.GetActiveUniformsiv(program.id, 1, &i, gl.UNIFORM_BLOCK_INDEX, &block)
		if block >= 0 {
			var offset, astride, mstride, rowmaj int32
			gl.GetActiveUniformsiv(program.id, 1, &i, gl.UNIFORM_OFFSET, &offset)
			gl.GetActiveUniformsiv(program.id, 1, &i, gl.UNIFORM_ARRAY_STRIDE, &astride)
			gl.GetActiveUniformsiv(program.id, 1, &i, gl.UNIFORM_MATRIX_STRIDE, &mstride)
			gl.GetActiveUniformsiv(program.id, 1, &i, gl.UNIFORM_IS_ROW_MAJOR, &rowmaj)
			members = append(members, UniformBlockMember{
				Program:  program,
				Name:     name,
				Type:     DataType(datatype),
				Size:     uint32(arraysize),
				Block:    uint32(block),
				Offset:   uint32(offset),
				AStride:  uint32(astride),
				MStride:  uint32(mstride),
				RowMajor: rowmaj == gl.TRUE,
			})
			continue
		}
	}
	sort.Sort(sortableMembers(members))
	return list, members
}

type sortableMembers []UniformBlockMember

func (s sortableMembers) Len() int {
	return len(s)
}

func (s sortableMembers) Less(i, j int) bool {
	if s[i].Block < s[j].Block {
		return true
	}
	if s[i].Block == s[j].Block {
		if s[i].Offset < s[j].Offset {
			return true
		}
	}
	return false
}

func (s sortableMembers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (program iProgram) getActiveUniformBlockName(index uint32, buf []byte) []byte {
	var length int32
	gl.GetActiveUniformBlockName(program.id, index, int32(len(buf)), &length, &buf[0])
	return buf[:length : length+1]
}

func (program iProgram) uniformBlocks() []UniformBlock {
	max := int(program.GetIV(ACTIVE_UNIFORM_BLOCKS))
	list := make([]UniformBlock, 0, max)
	buf := make([]byte, program.GetIV(ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH))
	for i := 0; i < max; i++ {
		namebytes := program.getActiveUniformBlockName(uint32(i), buf)
		location := gl.GetUniformBlockIndex(program.id, &namebytes[0])
		if location == INVALID_INDEX {
			continue
		}
		var size int32
		gl.GetActiveUniformBlockiv(program.id, location, gl.UNIFORM_BLOCK_DATA_SIZE, &size)
		name := string(namebytes)
		index := uint32(location)
		// binding := uint32(i)
		// gl.UniformBlockBinding(program.id, index, binding)
		list = append(list, UniformBlock{
			Program: program,
			Name:    name,
			Index:   index,
			Size:    uint32(size),
		})
	}
	sort.Sort(sortableBlocks(list))
	return list
}

type sortableBlocks []UniformBlock

func (s sortableBlocks) Len() int {
	return len(s)
}

func (s sortableBlocks) Less(i, j int) bool {
	return s[i].Index < s[j].Index
}

func (s sortableBlocks) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (coll iUniformCollection) List() []Uniform {
	return coll.list
}

func (coll iUniformCollection) ByIndex(index uint32) Uniform {
	i, ok := coll.byIndex[index]
	if ok {
		return coll.list[i]
	}
	return Uniform{}
}

func (coll iUniformCollection) ByName(name string) Uniform {
	i, ok := coll.byName[name]
	if ok {
		return coll.list[i]
	}
	return Uniform{}
}

func (coll iUniformCollection) Block(name string) UniformBlock {
	i, ok := coll.blockByName[name]
	if ok {
		return coll.blocks[i]
	}
	return UniformBlock{}
}

func (uni Uniform) Valid() bool {
	return uni.Size > 0
}

func (uni Uniform) Float(v ...float32) {
	if !uni.Valid() {
		panic(fmt.Errorf("ProgramUniform.Float: invalid uniform %#v", uni))
	}
	num := int32(len(v))
	switch uni.Type {
	case GlFloat:
		gl.ProgramUniform1fv(uni.Program.Id(), int32(uni.Index), num, &v[0])
	case GlFloatV2:
		gl.ProgramUniform2fv(uni.Program.Id(), int32(uni.Index), num/2, &v[0])
	case GlFloatV3:
		gl.ProgramUniform3fv(uni.Program.Id(), int32(uni.Index), num/3, &v[0])
	case GlFloatV4:
		gl.ProgramUniform4fv(uni.Program.Id(), int32(uni.Index), num/4, &v[0])
	case GlFloatMat2:
		gl.ProgramUniformMatrix2fv(uni.Program.Id(), int32(uni.Index), num/4, false, &v[0])
	case GlFloatMat2x3:
		gl.ProgramUniformMatrix2x3fv(uni.Program.Id(), int32(uni.Index), num/6, false, &v[0])
	case GlFloatMat2x4:
		gl.ProgramUniformMatrix2x4fv(uni.Program.Id(), int32(uni.Index), num/8, false, &v[0])
	case GlFloatMat3x2:
		gl.ProgramUniformMatrix3x2fv(uni.Program.Id(), int32(uni.Index), num/6, false, &v[0])
	case GlFloatMat3:
		gl.ProgramUniformMatrix3fv(uni.Program.Id(), int32(uni.Index), num/9, false, &v[0])
	case GlFloatMat3x4:
		gl.ProgramUniformMatrix3x4fv(uni.Program.Id(), int32(uni.Index), num/12, false, &v[0])
	case GlFloatMat4x2:
		gl.ProgramUniformMatrix4x2fv(uni.Program.Id(), int32(uni.Index), num/8, false, &v[0])
	case GlFloatMat4x3:
		gl.ProgramUniformMatrix4x3fv(uni.Program.Id(), int32(uni.Index), num/12, false, &v[0])
	case GlFloatMat4:
		gl.ProgramUniformMatrix4fv(uni.Program.Id(), int32(uni.Index), num/16, false, &v[0])
	default:
		panic(fmt.Errorf("ProgramUniform.Float: invalid type %v", uni.Type))
	}
}

func (uni Uniform) Uint(v ...uint32) {
	if !uni.Valid() {
		panic(fmt.Errorf("ProgramUniform.Uint: invalid uniform %#v", uni))
	}
	num := int32(len(v))
	switch uni.Type {
	case GlUInt:
		gl.ProgramUniform1uiv(uni.Program.Id(), int32(uni.Index), num, &v[0])
	case GlUIntV2:
		gl.ProgramUniform2uiv(uni.Program.Id(), int32(uni.Index), num/2, &v[0])
	case GlUIntV3:
		gl.ProgramUniform3uiv(uni.Program.Id(), int32(uni.Index), num/3, &v[0])
	case GlUIntV4:
		gl.ProgramUniform4uiv(uni.Program.Id(), int32(uni.Index), num/4, &v[0])
	default:
		panic(fmt.Errorf("ProgramUniform.Uint: invalid type %v", uni.Type))
	}
}

func (uni Uniform) Int(v ...int32) {
	if !uni.Valid() {
		panic(fmt.Errorf("ProgramUniform.Int: invalid uniform %#v", uni))
	}
	num := int32(len(v))
	switch uni.Type {
	case GlUInt:
		gl.ProgramUniform1iv(uni.Program.Id(), int32(uni.Index), num, &v[0])
	case GlUIntV2:
		gl.ProgramUniform2iv(uni.Program.Id(), int32(uni.Index), num/2, &v[0])
	case GlUIntV3:
		gl.ProgramUniform3iv(uni.Program.Id(), int32(uni.Index), num/3, &v[0])
	case GlUIntV4:
		gl.ProgramUniform4iv(uni.Program.Id(), int32(uni.Index), num/4, &v[0])
	default:
		panic(fmt.Errorf("ProgramUniform.Int: invalid type %v", uni.Type))
	}
}

func (uni Uniform) Sampler(v int32) {
	if !uni.Valid() {
		panic(fmt.Errorf("ProgramUniform.Int32: invalid uniform %#v", uni))
	}
	// TODO: Add the rest
	switch uni.Type {
	case GlSampler1dShadow, GlSampler2dShadow, GlSampler1dArrayShadow, GlSampler2dArrayShadow, GlSampler2dRectShadow, GlSamplerCubeShadow, GlSampler1d, GlSampler2d, GlSampler3d, GlSamplerCube, GlSampler2dArray, GlSampler2dMultisample, GlSampler2dMultisampleArray, GlSampler2dRect:
		gl.ProgramUniform1iv(uni.Program.Id(), int32(uni.Index), 1, &v)
	default:
		panic(fmt.Errorf("ProgramUniform.Sampler: invalid type %v", uni.Type))
	}
}

func (b UniformBlock) Valid() bool {
	return len(b.Uniforms) > 0
}
