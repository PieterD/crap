package main

var vertexShaderText = `
#version 330

in vec4 position;

void main() {
	gl_Position = position;
}
`

var fragmentShaderText = `
#version 330

out vec4 outputColor;
uniform float height;

void main() {
	float lerpValue = gl_FragCoord.y / height;
	outputColor = mix(vec4(1.0f, 1.0f, 1.0f, 1.0f), vec4(0.2f, 0.2f, 0.2f, 1.0f), lerpValue);
}
`
