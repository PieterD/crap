package gli

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type DefShader struct {
	version  int
	text     string
	stype    defShaderType
	inputs   []defVar
	outputs  []defVar
	uniforms []defVar
}

type defVar struct {
	datatype DataType
	name     string
}

type defShaderType byte

const (
	vertexShaderType defShaderType = iota + 1
	geometryShaderType
	fragmentShaderType
)

func NewVertexShaderFile(path string) (*DefShader, error) {
	return newShaderFile(vertexShaderType, "vertex", path)
}

func NewGeometryShaderFile(path string) (*DefShader, error) {
	return newShaderFile(geometryShaderType, "geometry", path)
}

func NewFragmentShaderFile(path string) (*DefShader, error) {
	return newShaderFile(fragmentShaderType, "fragment", path)
}

func NewVertexShaderString(text string) (*DefShader, error) {
	return newShaderString(vertexShaderType, "vertex", text)
}

func NewGeometryShaderString(text string) (*DefShader, error) {
	return newShaderString(geometryShaderType, "geometry", text)
}

func NewFragmentShaderString(text string) (*DefShader, error) {
	return newShaderString(fragmentShaderType, "fragment", text)
}

func newShaderFile(stype defShaderType, sname string, path string) (*DefShader, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read shader file %s: %v", path, err)
	}
	return newShaderString(stype, sname, string(bs))
}

func newShaderString(stype defShaderType, sname string, text string) (*DefShader, error) {
	inputs, outputs, uniforms, err := parseVariables(text)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse %s shader: %v", sname, err)
	}
	return &DefShader{
		text:     text,
		stype:    stype,
		inputs:   inputs,
		outputs:  outputs,
		uniforms: uniforms,
	}, nil
}

func CorrectShader(shader *DefShader, err error) *DefShader {
	if err != nil {
		panic(err)
	}
	return shader
}

var varRegexp = regexp.MustCompile(`^\s*(in|out|uniform)\s+([a-z][a-z0-9]+)\s+([_a-zA-Z][_a-zA-Z0-9]*)\s*;\s*(//.*)?$`)

func parseVariables(text string) (inputs []defVar, outputs []defVar, uniforms []defVar, err error) {
	for num, line := range strings.Split(text, "\n") {
		matches := varRegexp.FindStringSubmatch(line)
		if matches != nil {
			iou := matches[1]
			typ, err := typeFromString(matches[2])
			if err != nil {
				return nil, nil, nil, fmt.Errorf("Line %d: %v", num, err)
			}
			nam := matches[3]
			switch iou {
			case "in":
				inputs = append(inputs, defVar{datatype: typ, name: nam})
			case "out":
				outputs = append(outputs, defVar{datatype: typ, name: nam})
			case "uniform":
				uniforms = append(uniforms, defVar{datatype: typ, name: nam})
			}
		}
	}
	return
}
