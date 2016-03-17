package term

import (
	"io"
	"strings"
)

var termMap = make(map[string]Factory)

func Register(name string, factory Factory) {
	termMap[strings.ToLower(name)] = factory
}

func Get(name string, w io.Writer) Term {
	f, ok := termMap[strings.ToLower(name)]
	if !ok {
		return nil
	}
	return f.Create(w)
}

type Factory interface {
	Create(w io.Writer) Term
}

type Term interface {
	Clear() error
	Pos(x, y int) error
	Fore(color Color) error
	Back(color Color) error
	Attr(attr Attribute) error
	Reset() error
}

type Color int

const (
	Black   Color = '0'
	Red     Color = '1'
	Green   Color = '2'
	Yellow  Color = '3'
	Blue    Color = '4'
	Magenta Color = '5'
	Cyan    Color = '6'
	White   Color = '7'
	Default Color = '9'
)

type Attribute int

const (
	Bright     Attribute = '1'
	Dim        Attribute = '2'
	Underscore Attribute = '4'
	Blink      Attribute = '5'
)
