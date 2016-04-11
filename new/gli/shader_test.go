package gli

import "testing"

var shader_vert = CorrectShader(NewVertexShaderFile("testfiles/shader.vert"))
var shader_geom = CorrectShader(NewGeometryShaderFile("testfiles/shader.geom"))
var shader_frag = CorrectShader(NewFragmentShaderFile("testfiles/shader.frag"))

func TestShaders(t *testing.T) {
	if shader_vert.stype != vertexShaderType {
		t.Fatalf("wrong type for vertex shader")
	}
	if shader_geom.stype != geometryShaderType {
		t.Fatalf("wrong type for geometry shader")
	}
	if shader_frag.stype != fragmentShaderType {
		t.Fatalf("wrong type for fragment shader")
	}
}

func TestShaderString(t *testing.T) {
	_, err := NewVertexShaderString("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	_, err = NewGeometryShaderString("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	_, err = NewFragmentShaderString("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestNonexistentShaderFile(t *testing.T) {
	_, err := NewVertexShaderFile("testfiles/doesnotexist")
	if err == nil {
		t.Fatalf("Expected error, got nothing")
	}
}

func TestShaderError(t *testing.T) {
	shader, err := NewVertexShaderFile("testfiles/badshader")
	if err == nil {
		t.Fatalf("Expected error, was nil")
	}
	if shader != nil {
		t.Fatalf("Expected nil shader")
	}
}
