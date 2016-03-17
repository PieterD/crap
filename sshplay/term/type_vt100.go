package term

import (
	"io"
	"strconv"
)

func init() {
	Register("vt100", factoryVT100{})
}

type factoryVT100 struct{}

func (_ factoryVT100) Create(w io.Writer) Term {
	return &termVT100{w: w, b: make([]byte, 0, 50)}
}

type termVT100 struct {
	b []byte
	w io.Writer
}

func (t *termVT100) Clear() error {
	s := t.b
	s = append(s, 27, '[', '2', 'J')
	_, err := t.w.Write(s)
	return err
}

func (t *termVT100) Pos(x, y int) error {
	s := t.b
	s = append(s, 27, '[')
	s = strconv.AppendUint(s, uint64(y), 10)
	s = append(s, ';')
	s = strconv.AppendUint(s, uint64(x), 10)
	s = append(s, 'f')
	_, err := t.w.Write(s)
	return err
}

func (t *termVT100) Fore(color Color) error {
	s := t.b
	s = append(s, 27, '[', '3', byte(color), 'm')
	_, err := t.w.Write(s)
	return err
}

func (t *termVT100) Back(color Color) error {
	s := t.b
	s = append(s, 27, '[', '4', byte(color), 'm')
	_, err := t.w.Write(s)
	return err
}
