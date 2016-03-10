package gli

import "github.com/go-gl/gl/v3.3-core/gl"

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

func (coll AttributeCollection) ByIndex(index uint32) (ProgramAttribute, bool) {
	i, ok := coll.byIndex[index]
	if ok {
		return coll.list[i], true
	}
	return ProgramAttribute{}, false
}

func (coll AttributeCollection) ByName(name string) (ProgramAttribute, bool) {
	i, ok := coll.byName[name]
	if ok {
		return coll.list[i], true
	}
	return ProgramAttribute{}, false
}
