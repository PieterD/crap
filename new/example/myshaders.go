package main

import "fmt"

//go:generate gligen MyProgram shaders/shader.frag shaders/shader.geom shaders/shader.vert
func main() {
	fmt.Printf("moo\n")
}
