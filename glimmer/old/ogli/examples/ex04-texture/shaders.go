package main

var vertexShaderText = `
#version 330

in vec3 position;
in vec2 vertTexCoord;
out vec2 fragTexCoord;

void main() {
	fragTexCoord = vertTexCoord;
	gl_Position = vec4(position, 1);
}
`

var fragmentShaderText = `
#version 330

in vec2 fragTexCoord;
out vec4 outputColor;
uniform sampler2D sampler;

void main() {
	outputColor = texture(sampler, fragTexCoord);
}
`
