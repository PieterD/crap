package main

var vertexShaderText = `
#version 330
layout(location = 0) in vec4 position;
layout(location = 1) in vec4 color;
uniform vec2 offset;
uniform float zNear;
uniform float zFar;
uniform float frustumScale;

smooth out vec4 theColor;

void main() {
	vec4 cameraPos = position + vec4(offset.x, offset.y, 0.0, 0.0);
	vec4 clipPos;
	clipPos.xy = cameraPos.xy * frustumScale;
	clipPos.z = cameraPos.z * (zNear + zFar) / (zNear - zFar);
	clipPos.z += 2.0 * zNear * zFar / (zNear - zFar);
	clipPos.w = -cameraPos.z;

	gl_Position = clipPos;
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
