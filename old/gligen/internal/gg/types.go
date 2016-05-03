package gg

import "fmt"

type DataType struct {
	raw  rawType
	cols byte
	rows byte
	name string
}

func (dt DataType) String() string {
	return dt.name
}

type rawType byte

const (
	rawFloat rawType = iota + 1
	rawInt
	rawUint
)

var GlFloat = DataType{raw: rawFloat, cols: 0, rows: 0, name: "float"}
var GlVec2 = DataType{raw: rawFloat, cols: 2, rows: 0, name: "vec2"}
var GlVec3 = DataType{raw: rawFloat, cols: 3, rows: 0, name: "vec3"}
var GlVec4 = DataType{raw: rawFloat, cols: 4, rows: 0, name: "vec4"}

var typeNameMap = make(map[string]DataType)

func init() {
	addToNameMap(GlFloat)
	addToNameMap(GlVec2)
	addToNameMap(GlVec3)
	addToNameMap(GlVec4)
}

func addToNameMap(dt DataType) {
	typeNameMap[dt.name] = dt
}

func typeFromString(str string) (DataType, error) {
	dt, ok := typeNameMap[str]
	if !ok {
		return DataType{}, fmt.Errorf("Unknown datatype: '%s'", str)
	}
	return dt, nil
}
