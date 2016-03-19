package term

import "strings"

var termMap = make(map[string]Term)

func Register(name string, t Term) {
	termMap[strings.ToLower(name)] = t
}

func Resolve(name string) Term {
	return termMap[strings.ToLower(name)]
}

type Term interface {
	Clear(b []byte) []byte
	Pos(b []byte, x, y int) []byte
	Attr(b []byte) AttrBuilder
}

type Color int

const (
	Default Color = iota
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

type AttrBuilder interface {
	Reset() AttrBuilder
	Fore(color Color) AttrBuilder
	Back(color Color) AttrBuilder
	Bright() AttrBuilder
	Dim() AttrBuilder
	Underscore() AttrBuilder
	Blink() AttrBuilder
	Reverse() AttrBuilder
	Done() ([]byte, error)
}
