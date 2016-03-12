package main

var vertexShaderText = `
#version 330
layout(location = 0) in vec4 position;
uniform float duration;
uniform float time;

void main() {
	float timeScale = 3.14159f * 2.0f / duration;
	float currTime = mod(time, duration);
	vec4 totalOffset = vec4(
		cos(currTime * timeScale) * 0.5f,
		sin(currTime * timeScale) * 0.5f,
		0.0f,
		0.0f);
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
