package mat

import (
	"sync"

	"github.com/go-gl/mathgl/mgl32"
)

type stackNode struct {
	m    mgl32.Mat4
	next *stackNode
}

type Stack struct {
	node *stackNode
}

var stackPool = sync.Pool{New: func() interface{} { return new(stackNode) }}

func newStackNode() *stackNode {
	return stackPool.Get().(*stackNode)
}

func (n *stackNode) free() {
	stackPool.Put(n)
}

func NewStack() *Stack {
	n := newStackNode()
	n.m = mgl32.Ident4()
	return &Stack{node: n}
}

func (s *Stack) Push() {
	n := newStackNode()
	n.m = s.node.m
	n.next = s.node
	s.node = n
}

func (s *Stack) Pop() bool {
	if s.node.next == nil {
		s.node.m = mgl32.Ident4()
		return false
	}
	s.node = s.node.next
	return true
}

func (s *Stack) Peek() *mgl32.Mat4 {
	return &s.node.m
}
