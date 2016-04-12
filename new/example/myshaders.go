// +build generate

package main

import (
	"fmt"

	"github.com/PieterD/glimmer/new/gli"
)

var myVertexShader = gli.CorrectShader(gli.NewVertexShaderString(`
in vec4 position;
out vec4 color;

void main() {
	gl_Position = position;
	color = vec4(1.0, 0.0, 0.0, 1.0)
}
`))

var myFragmentShader = gli.CorrectShader(gli.NewFragmentShaderString(`
in vec4 color;
out vec4 outputColor;
uniform float height;

void main() {
	float lerpValue = gl_FragCoord.y / height;
	outputColor = mix(vec4(1.0f, 1.0f, 1.0f, 1.0f), color, lerpValue);
}
`))

var myProgram = gli.CorrectProgram(gli.NewProgram(myVertexShader, myFragmentShader))

//go:generate go run myshaders.go
func main() {
	fmt.Printf("moo\n")
}
