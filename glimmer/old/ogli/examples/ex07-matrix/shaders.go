package main

var vertexShaderText = `
#version 330

in vec4 position;
in vec4 color;
smooth out vec4 theColor;
uniform mat4 modelToCameraMatrix;
uniform mat4 cameraToClipMatrix;

void main() {
	vec4 cameraPos = modelToCameraMatrix * position;
	gl_Position = cameraToClipMatrix * cameraPos;
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
