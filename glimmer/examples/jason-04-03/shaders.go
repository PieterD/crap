package main

var vertexShaderText = `
#version 330
layout(location = 0) in vec4 position;
layout(location = 1) in vec4 color;
uniform vec2 offset;
uniform mat4 perspectiveMatrix;

smooth out vec4 theColor;

void main() {
	vec4 cameraPos = position + vec4(offset.x, offset.y, 0.0, 0.0);
	gl_Position = perspectiveMatrix * cameraPos;
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
