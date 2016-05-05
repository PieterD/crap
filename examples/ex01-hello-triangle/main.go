package main

import (
	"fmt"

	_ "github.com/PieterD/glimmer/driver/gl330"
	"github.com/PieterD/glimmer/win"
)

type Profile struct {
	win.DefaultHandler
}

func (p *Profile) Init() error {
	return nil
}

func (p *Profile) Draw() error {
	return nil
}

func main() {
	err := win.New(win.DefaultConfig(), &Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
