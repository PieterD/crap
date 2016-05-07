package prog

import (
	"fmt"
	"github.com/PieterD/glimmer/gli"
	"regexp"
	"strconv"
	"strings"
)

var inRegexpBasic = regexp.MustCompile(`^\s*in\s+`)
var inRegexp = regexp.MustCompile(`^(?:\s*in)(?:\s+([a-z][a-z0-9]+))(?:\s+([_a-zA-Z][_a-zA-Z0-9]*))(?:\s*\[([0-9]+)\])?(?:\s+BUFFER\(([a-zA-Z0-9_]+)\))?(?:\s*;)(?:\s*\/\/.*)?(?:\s*)$`)

func (coll *programCollection) parseAttrs() error {
	for _, group := range coll.groups {
		for _, program := range group.programs {
			program.buffers = make(map[string]*programAttributeBuffer)
			for _, filename := range program.vertexShader.files {
				content := coll.text[filename]
				err := program.parseAttrs(filename, content)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (p *program) parseAttrs(path string, content string) error {
	for num, line := range strings.Split(content, "\n") {
		if inRegexpBasic.MatchString(line) {
			matches := inRegexp.FindStringSubmatch(line)
			if matches == nil {
				return fmt.Errorf("Shader %s:%d: Failed to parse '%s'", path, num+1, line)
			}
			typ := matches[1]
			nam := matches[2]
			arr := matches[3]
			buf := matches[4]
			if buf == "" {
				buf = "DEFAULT"
			}
			if arr == "" {
				arr = "1"
			}
			size, err := strconv.ParseUint(arr, 10, 32)
			if err != nil {
				return fmt.Errorf("Shader %s:%d: Failed to parse array size '%s' for attribute '%s'", path, num+1, arr, nam)
			}
			dt, ok := attrTypeMap[typ]
			if !ok {
				return fmt.Errorf("Shader %s:%d: Failed to parse data type '%s' on line %d: unknown type", path, num+1, typ, num)
			}
			dt.Size = uint(size)
			err = dt.IsValid()
			if err != nil {
				return fmt.Errorf("Shader %s:%d: Failed to parse data type '%s' on line %d: %v", path, num+1, typ, num, err)
			}
			buffer, ok := p.buffers[buf]
			if !ok {
				buffer = &programAttributeBuffer{name: buf}
				p.buffers[buf] = buffer
			}
			buffer.attrs = append(buffer.attrs, &programAttribute{
				name:   nam,
				buffer: buf,
				typ:    dt,
			})
		}
	}
	return nil
}

var attrTypeMap = map[string]gli.DataType{
	"bool":   gli.DataType{Base: gli.BaseTypeBool},
	"bvec2":  gli.DataType{Base: gli.BaseTypeBool, Cols: 2},
	"bvec3":  gli.DataType{Base: gli.BaseTypeBool, Cols: 3},
	"bvec4":  gli.DataType{Base: gli.BaseTypeBool, Cols: 4},
	"int":    gli.DataType{Base: gli.BaseTypeInt},
	"ivec2":  gli.DataType{Base: gli.BaseTypeInt, Cols: 2},
	"ivec3":  gli.DataType{Base: gli.BaseTypeInt, Cols: 3},
	"ivec4":  gli.DataType{Base: gli.BaseTypeInt, Cols: 4},
	"uint":   gli.DataType{Base: gli.BaseTypeUnsignedInt},
	"uvec2":  gli.DataType{Base: gli.BaseTypeUnsignedInt, Cols: 2},
	"uvec3":  gli.DataType{Base: gli.BaseTypeUnsignedInt, Cols: 3},
	"uvec4":  gli.DataType{Base: gli.BaseTypeUnsignedInt, Cols: 4},
	"float":  gli.DataType{Base: gli.BaseTypeFloat},
	"vec2":   gli.DataType{Base: gli.BaseTypeFloat, Cols: 2},
	"vec3":   gli.DataType{Base: gli.BaseTypeFloat, Cols: 3},
	"vec4":   gli.DataType{Base: gli.BaseTypeFloat, Cols: 4},
	"double": gli.DataType{Base: gli.BaseTypeDouble},
	"dvec2":  gli.DataType{Base: gli.BaseTypeDouble, Cols: 2},
	"dvec3":  gli.DataType{Base: gli.BaseTypeDouble, Cols: 3},
	"dvec4":  gli.DataType{Base: gli.BaseTypeDouble, Cols: 4},
	"mat2":   gli.DataType{Base: gli.BaseTypeFloat, Cols: 2, Rows: 2},
	"mat3":   gli.DataType{Base: gli.BaseTypeFloat, Cols: 3, Rows: 3},
	"mat4":   gli.DataType{Base: gli.BaseTypeFloat, Cols: 4, Rows: 4},
	"mat2x3": gli.DataType{Base: gli.BaseTypeFloat, Cols: 2, Rows: 3},
	"mat2x4": gli.DataType{Base: gli.BaseTypeFloat, Cols: 2, Rows: 4},
	"mat3x2": gli.DataType{Base: gli.BaseTypeFloat, Cols: 3, Rows: 2},
	"mat3x4": gli.DataType{Base: gli.BaseTypeFloat, Cols: 3, Rows: 4},
	"mat4x2": gli.DataType{Base: gli.BaseTypeFloat, Cols: 4, Rows: 2},
	"mat4x3": gli.DataType{Base: gli.BaseTypeFloat, Cols: 4, Rows: 3},
}
