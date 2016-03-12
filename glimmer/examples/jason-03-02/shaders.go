package main

var vertexShaderText = `
#version 330
layout(location = 0) in vec4 position;
uniform vec2 offset;

void main() {
	vec4 totalOffset = vec4(offset.x, offset.y, 0.0, 0.0);
	gl_Position = position + totalOffset;
}
`

var fragmentShaderText = `
#version 330
out vec4 outputColor;
void main() {
	outputColor = vec4(1.0f, 1.0f, 1.0f, 1.0f);
}
`
