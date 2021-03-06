package mat

import (
	"errors"
	"sync"

	"github.com/go-gl/mathgl/mgl32"
)

var PopError = errors.New("Attempted to Pop an empty stack")

type stackNode struct {
	m    mgl32.Mat4
	next *stackNode
}

type Stack struct {
	angle   float32
	automul bool
	node    *stackNode
}

var stackPool = sync.Pool{New: func() interface{} { return new(stackNode) }}

func newStackNode(m mgl32.Mat4) *stackNode {
	n := stackPool.Get().(*stackNode)
	n.m = m
	return n
}

func (n *stackNode) free() {
	stackPool.Put(n)
}

func NewStack() *Stack {
	return &Stack{angle: 1.0, automul: true}
}

func (s *Stack) AngleDeg() *Stack {
	return s.AngleFactor(3.14159 * 2 / 360)
}

func (s *Stack) AngleRad() *Stack {
	return s.AngleFactor(1.0)
}

func (s *Stack) AngleFactor(a float32) *Stack {
	s.angle = a
	return s
}

func (s *Stack) AutoMultiply(mul bool) *Stack {
	s.automul = mul
	return s
}

func (s *Stack) Ident() {
	s.Push(mgl32.Ident4())
}

func (s *Stack) Copy() {
	s.Push(s.node.m)
}

func (s *Stack) Push(m mgl32.Mat4) {
	n := newStackNode(m)
	n.m = m
	n.next = s.node
	s.node = n
}

func (s *Stack) Pop() mgl32.Mat4 {
	n := s.node
	if n == nil {
		panic(PopError)
	}
	s.node = n.next
	n.next = nil
	m := n.m
	n.free()
	return m
}

func (s *Stack) Peek() mgl32.Mat4 {
	return s.node.m
}

func (s *Stack) Ptr() *mgl32.Mat4 {
	return &s.node.m
}

func (s *Stack) Swap() *Stack {
	a := s.Pop()
	b := s.Pop()
	s.Push(a)
	s.Push(b)
	return s
}

func (s *Stack) Multiply() *Stack {
	r := s.Pop()
	l := s.Pop()
	p := l.Mul4(r)
	s.Push(p)
	return s
}

func (s *Stack) Safe(f func(*Stack)) {
	n := newStackNode(s.Peek())
	ns := &Stack{node: n, angle: s.angle, automul: s.automul}
	f(ns)
}

func (s *Stack) Translate(x, y, z float32) *Stack {
	s.Push(mgl32.Translate3D(x, y, z))
	if s.automul {
		s.Multiply()
	}
	return s
}

func (s *Stack) TranslateV(v mgl32.Vec3) *Stack {
	return s.Translate(v[0], v[1], v[2])
}

func (s *Stack) Scale(x, y, z float32) *Stack {
	s.Push(mgl32.Scale3D(x, y, z))
	if s.automul {
		s.Multiply()
	}
	return s
}

func (s *Stack) ScaleV(v mgl32.Vec3) *Stack {
	return s.Scale(v[0], v[1], v[2])
}

func (s *Stack) RotateX(a float32) *Stack {
	s.Push(mgl32.HomogRotate3DX(a * s.angle))
	if s.automul {
		s.Multiply()
	}
	return s
}

func (s *Stack) RotateY(a float32) *Stack {
	s.Push(mgl32.HomogRotate3DY(a * s.angle))
	if s.automul {
		s.Multiply()
	}
	return s
}

func (s *Stack) RotateZ(a float32) *Stack {
	s.Push(mgl32.HomogRotate3DZ(a * s.angle))
	if s.automul {
		s.Multiply()
	}
	return s
}
