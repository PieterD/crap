package main

var vertexShaderText = `
#version 330

in vec4 Position;
in vec3 Color;

out Vertex {
	smooth vec4 Color;
} Out;

void main() {
	gl_Position = Position;
	Out.Color = vec4(Color, 1.0);
}
`

var fragmentShaderText = `
#version 330

in Vertex {
	smooth vec4 Color;
} In;

out vec4 Color;

uniform float colorshift;

void main() {
	Color = In.Color * colorshift;
}
`
