package gli_test

import (
	"testing"

	"github.com/PieterD/glimmer/new/gli"
)

var shader_vert = gli.CorrectShader(gli.NewVertexShaderFile("testfiles/shader.vert"))
var shader_geom = gli.CorrectShader(gli.NewGeometryShaderFile("testfiles/shader.geom"))
var shader_frag = gli.CorrectShader(gli.NewFragmentShaderFile("testfiles/shader.frag"))
var program = gli.CorrectProgram(gli.NewProgram(shader_vert, shader_geom, shader_frag))

func TestShaderString(t *testing.T) {
	_, err := gli.NewVertexShaderString("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	_, err = gli.NewGeometryShaderString("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	_, err = gli.NewFragmentShaderString("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestNonexistentShaderFile(t *testing.T) {
	_, err := gli.NewVertexShaderFile("testfiles/doesnotexist")
	if err == nil {
		t.Fatalf("Expected error, got nothing")
	}
}

func TestShaderError(t *testing.T) {
	shader, err := gli.NewVertexShaderFile("testfiles/badshader")
	if err == nil {
		t.Fatalf("Expected error, was nil")
	}
	if shader != nil {
		t.Fatalf("Expected nil shader")
	}
}

func TestProgram(t *testing.T) {
}
