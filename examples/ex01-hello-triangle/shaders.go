package main

var vertexShaderText = `
#version 330
layout(location = 0) in vec4 position;
layout(location = 1) in vec4 color;
smooth out vec4 theColor;
void main() {
	gl_Position = position;
	theColor = color;
}
`

var fragmentShaderText = `
#version 330
smooth in vec4 theColor;
out vec4 outputColor;
void main() {
	outputColor = theColor;
}
`
