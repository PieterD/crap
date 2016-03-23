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
	node *stackNode
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
	return &Stack{}
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

func (s *Stack) Multiply() {
	r := s.Pop()
	l := s.Pop()
	p := l.Mul4(r)
	s.Push(p)
}

func (s *Stack) Safe(f func(*Stack)) {
	n := newStackNode(s.Peek())
	ns := &Stack{node: n}
	f(ns)
}

func (s *Stack) RightMul(m mgl32.Mat4) {
	p := s.Ptr()
	*p = p.Mul4(m)
}
