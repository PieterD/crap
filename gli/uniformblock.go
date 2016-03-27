package gli

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type UniformBlockCollection struct {
	program Program
	byName  map[string]int
	byIndex map[uint32]int
	list    []UniformBlock
}

type UniformBlock struct {
	Program Program
	Name    string
	Index   uint32
	Binding uint32
}

func (program iProgram) getActiveUniformBlockName(index uint32, buf []byte) []byte {
	var length int32
	gl.GetActiveUniformBlockName(program.id, index, int32(len(buf)), &length, &buf[0])
	return buf[:length : length+1]
}

func (program iProgram) UniformBlocks() UniformBlockCollection {
	max := int(program.GetIV(ACTIVE_UNIFORM_BLOCKS))
	byname := make(map[string]int, max)
	byindex := make(map[uint32]int, max)
	uniformblocks := make([]UniformBlock, 0, max)
	buf := make([]byte, program.GetIV(ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH))
	for i := 0; i < max; i++ {
		namebytes := program.getActiveUniformBlockName(uint32(i), buf)
		fmt.Printf("name: '%s'\n", namebytes)
		name := string(namebytes)
		location := gl.GetUniformBlockIndex(program.id, &namebytes[0])
		if location == INVALID_INDEX {
			continue
		}
		index := uint32(location)
		binding := uint32(i)
		gl.UniformBlockBinding(program.id, index, binding)
		byname[name] = len(uniformblocks)
		byindex[index] = len(uniformblocks)
		uniformblocks = append(uniformblocks, UniformBlock{
			Program: program,
			Name:    name,
			Index:   index,
			Binding: binding,
		})
	}
	return UniformBlockCollection{
		program: program,
		byName:  byname,
		byIndex: byindex,
		list:    uniformblocks,
	}
}

func (coll UniformBlockCollection) List() []UniformBlock {
	return coll.list
}

func (coll UniformBlockCollection) ByIndex(index uint32) UniformBlock {
	i, ok := coll.byIndex[index]
	if ok {
		return coll.list[i]
	}
	return UniformBlock{}
}

func (coll UniformBlockCollection) ByName(name string) UniformBlock {
	i, ok := coll.byName[name]
	if ok {
		return coll.list[i]
	}
	return UniformBlock{}
}
