package term

import (
	"fmt"
	"io"
	"sync/atomic"
)

type Full struct {
	term Term
	b    []byte
	w    io.Writer
	dim  uint64
	err  error
}

func New(term Term, w io.Writer) *Full {
	if term == nil {
		return nil
	}
	return &Full{
		term: term,
		b:    make([]byte, 50),
		w:    w,
	}
}

func (full *Full) SetDimensions(width, height uint32) {
	dim := uint64(width)<<32 | uint64(height)
	atomic.StoreUint64(&full.dim, dim)
}

func (full *Full) GetDimensions() (width, height uint32) {
	dim := atomic.LoadUint64(&full.dim)
	width = uint32(dim >> 32)
	height = uint32(dim)
	return
}

func (full *Full) Error() error {
	return full.err
}

func (full *Full) Printf(format string, args ...interface{}) *Full {
	if full.Error() != nil {
		return full
	}
	_, full.err = fmt.Fprintf(full.w, format, args...)
	return full
}

func (full *Full) Write(b []byte) (n int, err error) {
	if full.Error() != nil {
		return 0, err
	}
	n, err = full.w.Write(b)
	full.err = err
	return n, err
}

func (full *Full) write() {
	if full.err != nil {
		return
	}
	_, full.err = full.w.Write(full.b)
	full.b = full.b[:0]
}

func (full *Full) Clear() *Full {
	if full.Error() != nil {
		return full
	}
	full.b = full.term.Clear(full.b[:0])
	full.write()
	return full
}

func (full *Full) Pos(x, y int) *Full {
	if full.Error() != nil {
		return full
	}
	full.b = full.term.Pos(full.b[:0], x, y)
	full.write()
	return full
}

func (full *Full) Attr() FullAttr {
	return FullAttr{ab: full.term.Attr(full.b[:0]), full: full}
}

type FullAttr struct {
	ab   AttrBuilder
	full *Full
}

func (fa FullAttr) Reset() FullAttr {
	fa.ab = fa.ab.Reset()
	return fa
}

func (fa FullAttr) Fore(color Color) FullAttr {
	fa.ab = fa.ab.Fore(color)
	return fa
}

func (fa FullAttr) Back(color Color) FullAttr {
	fa.ab = fa.ab.Back(color)
	return fa
}

func (fa FullAttr) Bright() FullAttr {
	fa.ab = fa.ab.Bright()
	return fa
}

func (fa FullAttr) Dim() FullAttr {
	fa.ab = fa.ab.Dim()
	return fa
}

func (fa FullAttr) Underscore() FullAttr {
	fa.ab = fa.ab.Underscore()
	return fa
}

func (fa FullAttr) Blink() FullAttr {
	fa.ab = fa.ab.Blink()
	return fa
}

func (fa FullAttr) Reverse() FullAttr {
	fa.ab = fa.ab.Reverse()
	return fa
}

func (fa FullAttr) Done() *Full {
	if fa.full.Error() != nil {
		return fa.full
	}
	fa.full.b, fa.full.err = fa.ab.Done()
	if fa.full.err == nil && len(fa.full.b) > 0 {
		fa.full.write()
	}
	return fa.full
}
