package gli

import "github.com/go-gl/gl/v3.3-core/gl"

type AttributeCollection struct {
	program Program
	byName  map[string]int
	byIndex map[uint32]int
	list    []ProgramAttribute
}

type ProgramAttribute struct {
	Program Program
	Name    string
	Index   uint32
	Type    DataType
	Size    uint32
}

func (program iProgram) getActiveAttrib(index uint32, buf []byte) (name []byte, datatype DataType, size int) {
	var length int32
	var isize int32
	var idatatype uint32
	gl.GetActiveAttrib(program.id, index, int32(len(buf)), &length, &isize, &idatatype, &buf[0])
	return buf[:length : length+1], DataType(idatatype), int(isize)
}

func (program iProgram) Attributes() AttributeCollection {
	max := int(program.GetIV(ACTIVE_ATTRIBUTES))
	byname := make(map[string]int, max)
	byindex := make(map[uint32]int, max)
	attributes := make([]ProgramAttribute, 0, max)
	buf := make([]byte, program.GetIV(ACTIVE_ATTRIBUTE_MAX_LENGTH))
	for i := 0; i < max; i++ {
		namebytes, datatype, size := program.getActiveAttrib(uint32(i), buf)
		name := string(namebytes)
		location := gl.GetAttribLocation(program.id, &namebytes[0])
		if location <= -1 {
			continue
		}
		index := uint32(location)
		byname[name] = len(attributes)
		byindex[index] = len(attributes)
		attributes = append(attributes, ProgramAttribute{
			Program: program,
			Name:    name,
			Index:   index,
			Type:    datatype,
			Size:    uint32(size),
		})
	}
	return AttributeCollection{
		program: program,
		list:    attributes,
		byName:  byname,
		byIndex: byindex,
	}
}

func (coll AttributeCollection) List() []ProgramAttribute {
	return coll.list
}

func (coll AttributeCollection) ByIndex(index uint32) ProgramAttribute {
	i, ok := coll.byIndex[index]
	if ok {
		return coll.list[i]
	}
	return ProgramAttribute{}
}

func (coll AttributeCollection) ByName(name string) ProgramAttribute {
	i, ok := coll.byName[name]
	if ok {
		return coll.list[i]
	}
	return ProgramAttribute{}
}

func (attr ProgramAttribute) Valid() bool {
	return attr.Size > 0
}
