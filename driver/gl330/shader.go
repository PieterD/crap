package gl330

import (
	"fmt"
	"github.com/PieterD/gl/v3.3-core/gl"
	"github.com/PieterD/glimmer/convc"
	"github.com/PieterD/glimmer/gli"
)

func (_ gl330) ShaderCreate(typ gli.ShaderType, sources ...string) (gli.ShaderId, error) {
	shaderid := gl.CreateShader(toUint(typ))
	if shaderid == 0 {
		return 0, fmt.Errorf("Failed to create shader: %v", getError())
	}
	ptr, free := convc.MultiStringToC(sources...)
	defer free()
	gl.ShaderSource(shaderid, int32(len(sources)), ptr, nil)
	gl.CompileShader(shaderid)
	var status int32
	gl.GetShaderiv(shaderid, gl.COMPILE_STATUS, &status)
	var err error
	if status == gl.TRUE {
		return gli.ShaderId(shaderid), nil
	} else if status == gl.FALSE {
		var loglength int32
		gl.GetShaderiv(shaderid, gl.INFO_LOG_LENGTH, &loglength)
		log := make([]byte, loglength)
		gl.GetShaderInfoLog(shaderid, int32(cap(log)), &loglength, &log[0])
		log = log[:loglength]
		err = fmt.Errorf("Failed to compile shader: %s", log)
	} else {
		err = fmt.Errorf("Failed to create new shader: COMPILE_STATUS was neither TRUE or FALSE, but %d", status)
	}
	gl.DeleteShader(shaderid)
	return 0, err
}

func (_ gl330) ShaderDelete(id gli.ShaderId) {
	gl.DeleteShader(uint32(id))
}
