package gligen_test

import (
	"testing"

	"github.com/PieterD/gligenmmer/new/gligen"
)

var shader_vert = gligen.CorrectShader(gligen.NewVertexShaderFile("testfiles/shader.vert"))
var shader_geom = gligen.CorrectShader(gligen.NewGeometryShaderFile("testfiles/shader.geom"))
var shader_frag = gligen.CorrectShader(gligen.NewFragmentShaderFile("testfiles/shader.frag"))
var program = gligen.CorrectProgram(gligen.NewProgram(shader_vert, shader_geom, shader_frag))

func TestShaderString(t *testing.T) {
	_, err := gligen.NewVertexShaderString("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	_, err = gligen.NewGeometryShaderString("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	_, err = gligen.NewFragmentShaderString("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestNonexistentShaderFile(t *testing.T) {
	_, err := gligen.NewVertexShaderFile("testfiles/doesnotexist")
	if err == nil {
		t.Fatalf("Expected error, got nothing")
	}
}

func TestShaderError(t *testing.T) {
	shader, err := gligen.NewVertexShaderFile("testfiles/badshader")
	if err == nil {
		t.Fatalf("Expected error, was nil")
	}
	if shader != nil {
		t.Fatalf("Expected nil shader")
	}
}

func TestProgram(t *testing.T) {
}
