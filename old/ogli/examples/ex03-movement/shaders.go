package main

var vertexShaderText = `
#version 330

in vec4 position;
uniform vec2 offset;

void main() {
	vec4 totalOffset = vec4(offset.x, offset.y, 0.0, 0.0);
	gl_Position = position + totalOffset;
}
`

var fragmentShaderText = `
#version 330

out vec4 outputColor;
uniform float fragDuration;
uniform float time;

const vec4 firstColor = vec4(1.0f, 1.0f, 1.0f, 1.0f);
const vec4 secondColor = vec4(0.0f, 1.0f, 0.0f, 1.0f);
void main() {
	float currTime = mod(time, fragDuration);
	float currLerp = currTime / fragDuration;
	outputColor = mix(firstColor, secondColor, currLerp);
}
`
