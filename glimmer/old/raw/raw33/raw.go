package raw33

import (
	"fmt"
	"unsafe"

	"github.com/PieterD/glimmer/convc"
	"github.com/PieterD/glimmer/raw"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Raw struct{}

func (_ Raw) Init() error {
	return gl.Init()
}

func (_ Raw) Viewport(x, y, width, height int) {
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
}

func (_ Raw) ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

func (_ Raw) ShaderCreate(iShadertype raw.Enum) (uint32, error) {
	shadertype := aShaderType.unsigned(iShadertype)
	id := gl.CreateShader(shadertype)
	if id == 0 {
		err := getError()
		return 0, fmt.Errorf("Failed to create shader: %v", err)
	}
	return id, nil
}

func (_ Raw) ShaderDelete(shaderid uint32) {
	gl.DeleteShader(shaderid)
}

func (_ Raw) ShaderSource(shaderid uint32, source ...string) {
	ptr, free := convc.MultiStringToC(source...)
	defer free()
	gl.ShaderSource(shaderid, int32(len(source)), ptr, nil)
}

func (_ Raw) ShaderCompile(shaderid uint32) {
	gl.CompileShader(shaderid)
}

func (_ Raw) ShaderCompileStatus(shaderid uint32) bool {
	var pi int32
	gl.GetShaderiv(shaderid, gl.COMPILE_STATUS, &pi)
	if pi == gl.TRUE {
		return true
	}
	return false
}

func (_ Raw) ShaderInfoLogLength(shaderid uint32) int {
	var pi int32
	gl.GetShaderiv(shaderid, gl.INFO_LOG_LENGTH, &pi)
	return int(pi)
}

func (_ Raw) ShaderInfoLog(shaderid uint32, buf []byte) []byte {
	var length int32
	gl.GetShaderInfoLog(shaderid, int32(cap(buf)), &length, &buf[0])
	if int(length) > cap(buf) {
		return buf[:cap(buf)]
	}
	return buf[:length]
}

func (_ Raw) ProgramCreate() (programid uint32, err error) {
	id := gl.CreateProgram()
	if id == 0 {
		err := getError()
		return 0, fmt.Errorf("Failed to create program: %v", err)
	}
	return id, nil
}

func (_ Raw) ProgramDelete(programid uint32) {
	gl.DeleteProgram(programid)
}

func (_ Raw) ProgramAttachShader(programid uint32, shaderid uint32) {
	gl.AttachShader(programid, shaderid)
}

func (_ Raw) ProgramLink(programid uint32) {
	gl.LinkProgram(programid)
}

func (_ Raw) ProgramLinkStatus(programid uint32) bool {
	var pi int32
	gl.GetProgramiv(programid, gl.LINK_STATUS, &pi)
	if pi == gl.TRUE {
		return true
	}
	return false
}

func (_ Raw) ProgramInfoLogLength(programid uint32) int {
	var pi int32
	gl.GetProgramiv(programid, gl.INFO_LOG_LENGTH, &pi)
	return int(pi)
}

func (_ Raw) ProgramInfoLog(programid uint32, buf []byte) []byte {
	var length int32
	gl.GetProgramInfoLog(programid, int32(cap(buf)), &length, &buf[0])
	if int(length) > cap(buf) {
		return buf[:cap(buf)]
	}
	return buf[:length]
}

func (_ Raw) ProgramAttributeNum(programid uint32) int {
	var pi int32
	gl.GetProgramiv(programid, gl.ACTIVE_ATTRIBUTES, &pi)
	return int(pi)
}

func (_ Raw) ProgramAttributeMaxLength(programid uint32) int {
	var pi int32
	gl.GetProgramiv(programid, gl.ACTIVE_ATTRIBUTE_MAX_LENGTH, &pi)
	return int(pi)
}

func (_ Raw) ProgramAttribute(programid uint32, index int, buf []byte) (namebytes []byte, datatype raw.Enum, size int) {
	var length int32
	var isize int32
	var idatatype uint32
	gl.GetActiveAttrib(programid, uint32(index), int32(len(buf)), &length, &isize, &idatatype, &buf[0])
	dt, ok := aDataType.reverse(int64(idatatype))
	if !ok {
		panic(fmt.Errorf("Failed to reverse map gl data type %d", idatatype))
	}
	return buf[:length : length+1], dt, int(isize)
}

func (_ Raw) ProgramAttributeLocation(programid uint32, namebytes []byte) (location int, ok bool) {
	location = int(gl.GetAttribLocation(programid, &namebytes[0]))
	if location <= -1 {
		return -1, false
	}
	return location, true
}

func (_ Raw) ProgramAttributeLocationBind(programid uint32, index int, namebytes []byte) {
	gl.BindAttribLocation(programid, uint32(index), &namebytes[0])
}

func (_ Raw) ProgramUniformNum(programid uint32) int {
	var pi int32
	gl.GetProgramiv(programid, gl.ACTIVE_UNIFORMS, &pi)
	return int(pi)
}

func (_ Raw) ProgramUniformMaxLength(programid uint32) int {
	var pi int32
	gl.GetProgramiv(programid, gl.ACTIVE_UNIFORM_MAX_LENGTH, &pi)
	return int(pi)
}

func (_ Raw) ProgramUniform(programid uint32, index int, buf []byte) (namebytes []byte, datatype raw.Enum, size int) {
	var length int32
	var isize int32
	var idatatype uint32
	gl.GetActiveUniform(programid, uint32(index), int32(len(buf)), &length, &isize, &idatatype, &buf[0])
	dt, ok := aDataType.reverse(int64(idatatype))
	if !ok {
		panic(fmt.Errorf("Failed to reverse map gl data type %d", idatatype))
	}
	return buf[:length : length+1], dt, int(isize)
}

func (_ Raw) ProgramUniformLocation(programid uint32, namebytes []byte) (location int, ok bool) {
	location = int(gl.GetUniformLocation(programid, &namebytes[0]))
	if location <= -1 {
		return -1, false
	}
	return location, true
}

func (_ Raw) BufferCreate() (id uint32) {
	gl.GenBuffers(1, &id)
	return id
}

func (_ Raw) BufferDelete(bufferid uint32) {
	gl.DeleteBuffers(1, &bufferid)
}

func (_ Raw) BufferBind(bufferid uint32, target raw.Enum) {
	gl.BindBuffer(uint32(target), bufferid)
}

func (_ Raw) BufferData(target raw.Enum, bytes int, ptr unsafe.Pointer, accesstype raw.Enum) {
	gl.BufferData(uint32(target), bytes, ptr, uint32(accesstype))
}

func (_ Raw) BufferSubData(target raw.Enum, offset int, bytes int, ptr unsafe.Pointer) {
	gl.BufferSubData(uint32(target), offset, bytes, ptr)
}

func (_ Raw) VertexArrayCreate() (vaoid uint32) {
	var id uint32
	gl.GenVertexArrays(1, &id)
	return id
}

func (_ Raw) VertexArrayDelete(vaoid uint32) {
	gl.DeleteVertexArrays(1, &vaoid)
}

func (_ Raw) VertexArrayBind(vaoid uint32) {
	gl.BindVertexArray(vaoid)
}

func (_ Raw) VertexArrayEnable(idx int) {
	gl.EnableVertexAttribArray(uint32(idx))
}

func (_ Raw) VertexArrayDisable(idx int) {
	gl.DisableVertexAttribArray(uint32(idx))
}

func (_ Raw) SyncFence() unsafe.Pointer {
	return gl.FenceSync(gl.SYNC_GPU_COMMANDS_COMPLETE, 0)
}

func (_ Raw) SyncClientWait(s unsafe.Pointer, flush bool, timeout uint64) raw.Enum {
	var bits uint32
	if flush {
		bits = gl.SYNC_FLUSH_COMMANDS_BIT
	}
	rv := gl.ClientWaitSync(s, bits, timeout)
	enum, ok := aSyncResult.reverse(int64(rv))
	if !ok {
		panic(fmt.Errorf("Impossible return value from glClientWaitSync: %d", rv))
	}
	return enum
}
