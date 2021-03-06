package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PieterD/glimmer/gligen/internal/gg"
)

var usage = `
Usage:
  gligen <ProgramName> <shaderfile> <shaderfile> [shaderfile] [...]
Where:
  ProgramName:
    The name of the program.
  shaderfile:
    A file containing a shader.
    Must have one of the following extensions:
      .vert
      .geom
      .frag
`

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, usage)
		return
	}
	pkg := os.Getenv("GOPACKAGE")
	name := os.Args[1]
	paths := os.Args[2:]
	fmt.Printf("generating: %s.%s\n", pkg, name)
	for _, path := range paths {
		fmt.Printf("  shader: %s\n", path)
	}
	_, err := getProgram(paths)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get program %s: %v\n", name, err)
		return
	}
}

func getProgram(paths []string) (*gg.DefProgram, error) {
	shaders := make([]*gg.DefShader, 0, len(paths))
	for _, path := range paths {
		var shader *gg.DefShader
		var err error
		if strings.HasSuffix(path, ".vert") {
			shader, err = gg.NewVertexShaderFile(path)
		} else if strings.HasSuffix(path, ".geom") {
			shader, err = gg.NewGeometryShaderFile(path)
		} else if strings.HasSuffix(path, ".frag") {
			shader, err = gg.NewFragmentShaderFile(path)
		} else {
			err = fmt.Errorf("Unknown extension")
		}
		if err != nil {
			return nil, fmt.Errorf("Failed to read shader file %s: %v", path, err)
		}
		shaders = append(shaders, shader)
	}
	return gg.NewProgram(shaders...)
}
