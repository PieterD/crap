// Code generated by "stringer -type=ShaderType"; DO NOT EDIT

package gli

import "fmt"

const _ShaderType_name = "ShaderTypeVertexShaderTypeGeometryShaderTypeFragment"

var _ShaderType_index = [...]uint8{0, 16, 34, 52}

func (i ShaderType) String() string {
	i -= 1
	if i >= ShaderType(len(_ShaderType_index)-1) {
		return fmt.Sprintf("ShaderType(%d)", i+1)
	}
	return _ShaderType_name[_ShaderType_index[i]:_ShaderType_index[i+1]]
}
