package main

var vertexShaderText = `
#version 330

in vec4 position;
in vec4 color;
smooth out vec4 theColor;
uniform vec3 offset;
uniform mat4 perspectiveMatrix;

void main() {
	vec4 cameraPos = position + vec4(offset.x, offset.y, offset.z, 0.0);
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
