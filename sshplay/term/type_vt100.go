package term

import (
	"fmt"
	"strconv"
)

func init() {
	Register("vt100", TermVT100{})
	Register("vt220", TermVT100{})
	Register("xterm", TermVT100{})
}

type TermVT100 struct{}

func (t TermVT100) Clear(b []byte) []byte {
	b = b[:0]
	b = append(b, 27, '[', '2', 'J')
	return b
}

func (t TermVT100) Pos(b []byte, x, y int) []byte {
	b = b[:0]
	b = append(b, 27, '[')
	b = strconv.AppendUint(b, uint64(y), 10)
	b = append(b, ';')
	b = strconv.AppendUint(b, uint64(x), 10)
	b = append(b, 'f')
	return b
}

func (t TermVT100) Attr(b []byte) AttrBuilder {
	b = b[:0]
	b = append(b, 27, '[')
	// TODO: If this allocates, use sync.Pool
	return &attrVT100{b: b}
}

type attrVT100 struct {
	b    []byte
	more bool
	err  error
	done bool
}

func (a *attrVT100) Reset() AttrBuilder {
	if a.done {
		return a
	}
	if a.more {
		a.b = a.b[:0]
		a.b = append(a.b, 27, '[')
		a.err = nil
	} else {
		a.more = true
	}
	a.b = append(a.b, '0')
	return a
}

func (a *attrVT100) color(b byte, color Color) {
	if a.err != nil || a.done {
		return
	}
	if a.more {
		a.b = append(a.b, ';')
	} else {
		a.more = true
	}
	var c byte
	switch color {
	case Black:
		c = '0'
	case Red:
		c = '1'
	case Green:
		c = '2'
	case Yellow:
		c = '3'
	case Blue:
		c = '4'
	case Magenta:
		c = '5'
	case Cyan:
		c = '6'
	case White:
		c = '7'
	case Default:
		c = '9'
	default:
		a.err = fmt.Errorf("Unknown color value: %d", int(color))
		return
	}
	a.b = append(a.b, b, c)
}

func (a *attrVT100) Fore(color Color) AttrBuilder {
	a.color('3', color)
	return a
}

func (a *attrVT100) Back(color Color) AttrBuilder {
	a.color('4', color)
	return a
}

func (a *attrVT100) Bright() AttrBuilder {
	if a.err != nil || a.done {
		return a
	}
	if a.more {
		a.b = append(a.b, ';')
	} else {
		a.more = true
	}
	a.b = append(a.b, '1')
	return a
}

func (a *attrVT100) Dim() AttrBuilder {
	if a.err != nil || a.done {
		return a
	}
	if a.more {
		a.b = append(a.b, ';')
	} else {
		a.more = true
	}
	a.b = append(a.b, '2')
	return a
}

func (a *attrVT100) Underscore() AttrBuilder {
	if a.err != nil || a.done {
		return a
	}
	if a.more {
		a.b = append(a.b, ';')
	} else {
		a.more = true
	}
	a.b = append(a.b, '4')
	return a
}

func (a *attrVT100) Blink() AttrBuilder {
	if a.err != nil || a.done {
		return a
	}
	if a.more {
		a.b = append(a.b, ';')
	} else {
		a.more = true
	}
	a.b = append(a.b, '5')
	return a
}

func (a *attrVT100) Reverse() AttrBuilder {
	if a.err != nil || a.done {
		return a
	}
	if a.more {
		a.b = append(a.b, ';')
	} else {
		a.more = true
	}
	a.b = append(a.b, '7')
	return a
}

func (a *attrVT100) Done() ([]byte, error) {
	if a.done {
	} else if a.err != nil {
		a.b = a.b[:0]
	} else if !a.more {
		a.b = a.b[:0]
	} else {
		a.done = true
		a.b = append(a.b, 'm')
	}
	return a.b, a.err
}
