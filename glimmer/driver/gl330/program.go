package gl330

import (
	"fmt"
	"github.com/PieterD/glimmer/gli"
	"github.com/go-gl/gl/v3.3-core/gl"
)

func (_ gl330) ProgramCreate(locations []gli.AttributeLocation, shaders ...gli.ShaderId) (gli.ProgramId, error) {
	programid := gl.CreateProgram()
	if programid == 0 {
		return 0, fmt.Errorf("Failed to create program: %v", getError())
	}
	bindAttributeLocations(programid, locations)
	for _, shaderid := range shaders {
		gl.AttachShader(programid, uint32(shaderid))
	}
	gl.LinkProgram(programid)
	var status int32
	gl.GetProgramiv(programid, gl.LINK_STATUS, &status)
	var err error
	if status == gl.TRUE {
		return gli.ProgramId(programid), nil
	} else if status == gl.FALSE {
		var loglength int32
		gl.GetProgramiv(programid, gl.INFO_LOG_LENGTH, &loglength)
		log := make([]byte, loglength)
		gl.GetProgramInfoLog(programid, int32(cap(log)), &loglength, &log[0])
		log = log[:loglength]
		err = fmt.Errorf("Failed to link program: %s", log)
	} else {
		err = fmt.Errorf("Failed to create new program: LINK_STATUS was neither TRUE or FALSE, but %d", status)
	}
	gl.DeleteProgram(programid)
	return 0, err
}

func bindAttributeLocations(programid uint32, locations []gli.AttributeLocation) {
	size := 0
	for _, location := range locations {
		if len(location.Name) > size {
			size = len(location.Name)
		}
	}
	size++
	buf := make([]byte, size)
	for _, location := range locations {
		buf = append(buf[:0], location.Name...)
		buf = append(buf, 0)
		gl.BindAttribLocation(programid, uint32(location.Location), &buf[0])
	}
}

func (_ gl330) ProgramDelete(id gli.ProgramId) {
	gl.DeleteProgram(uint32(id))
}

func (_ gl330) ProgramAttributes(id gli.ProgramId) ([]gli.ProgramResource, error) {
	var attrnum int32
	gl.GetProgramiv(uint32(id), gl.ACTIVE_ATTRIBUTES, &attrnum)
	var namemax int32
	gl.GetProgramiv(uint32(id), gl.ACTIVE_ATTRIBUTE_MAX_LENGTH, &namemax)
	namebuf := make([]byte, namemax)
	attrs := make([]gli.ProgramResource, 0, attrnum)
	for i := 0; i < int(attrnum); i++ {
		var namelength, arraysize int32
		var typeenum uint32
		gl.GetActiveAttrib(uint32(id), uint32(i), namemax, &namelength, &arraysize, &typeenum, &namebuf[0])
		if arraysize == 0 {
			continue
		}
		location := gl.GetAttribLocation(uint32(id), &namebuf[0])
		if location == -1 {
			continue
		}
		dt, ok := attrDataMap[typeenum]
		if !ok {
			return nil, fmt.Errorf("Unknown type enum for attribute '%s': %d", namebuf[:namelength], typeenum)
		}
		err := dt.IsValid()
		if err != nil {
			return nil, err
		}
		dt.Size = uint(arraysize)
		attrs = append(attrs, gli.ProgramResource{
			Resource: gli.ResourceTypeAttribute,
			Name:     string(namebuf[:namelength]),
			Type:     dt,
			Index:    uint(location),
		})
	}
	return attrs, nil
}

var attrDataMap = map[uint32]gli.DataType{
	gl.FLOAT:      gli.DataType{Base: gli.BaseTypeFloat},
	gl.FLOAT_VEC2: gli.DataType{Base: gli.BaseTypeFloat, Cols: 2},
	gl.FLOAT_VEC3: gli.DataType{Base: gli.BaseTypeFloat, Cols: 3},
	gl.FLOAT_VEC4: gli.DataType{Base: gli.BaseTypeFloat, Cols: 4},
	gl.FLOAT_MAT2: gli.DataType{Base: gli.BaseTypeFloat, Cols: 2, Matrix: true},
	gl.FLOAT_MAT3: gli.DataType{Base: gli.BaseTypeFloat, Cols: 3, Matrix: true},
	gl.FLOAT_MAT4: gli.DataType{Base: gli.BaseTypeFloat, Cols: 4, Matrix: true},

	gl.INT:               gli.DataType{Base: gli.BaseTypeInt},
	gl.INT_VEC2:          gli.DataType{Base: gli.BaseTypeInt, Cols: 2},
	gl.INT_VEC3:          gli.DataType{Base: gli.BaseTypeInt, Cols: 3},
	gl.INT_VEC4:          gli.DataType{Base: gli.BaseTypeInt, Cols: 4},
	gl.UNSIGNED_INT:      gli.DataType{Base: gli.BaseTypeUnsignedInt},
	gl.UNSIGNED_INT_VEC2: gli.DataType{Base: gli.BaseTypeUnsignedInt, Cols: 2},
	gl.UNSIGNED_INT_VEC3: gli.DataType{Base: gli.BaseTypeUnsignedInt, Cols: 3},
	gl.UNSIGNED_INT_VEC4: gli.DataType{Base: gli.BaseTypeUnsignedInt, Cols: 4},
}
