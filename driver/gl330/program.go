package gl330

import (
	"fmt"
	"github.com/PieterD/gl/v3.3-core/gl"
	"github.com/PieterD/glimmer/gli"
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
