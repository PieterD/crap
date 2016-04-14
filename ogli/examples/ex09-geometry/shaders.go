package main

var vertexShaderText = `
#version 330

in vec2 position;

void main() {
	gl_Position = vec4(position, 0.0, 1.0);
}
`

var geometryShaderText = `
#version 330
layout(points) in;
layout(triangle_strip, max_vertices = 4) out;

void main() {
	gl_Position = gl_in[0].gl_Position + vec4(-0.1, -0.1, 0.0, 0.0);
	EmitVertex();
	gl_Position = gl_in[0].gl_Position + vec4(+0.1, -0.1, 0.0, 0.0);
	EmitVertex();
	gl_Position = gl_in[0].gl_Position + vec4(-0.1, +0.1, 0.0, 0.0);
	EmitVertex();
	gl_Position = gl_in[0].gl_Position + vec4(+0.1, +0.1, 0.0, 0.0);
	EmitVertex();
	EndPrimitive();
}
`

var fragmentShaderText = `
#version 330

out vec4 outputColor;

void main() {
	outputColor = vec4(1.0, 0.0, 0.0, 1.0);
}
`
