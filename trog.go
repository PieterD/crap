package trog

import "encoding/base64"

type rawNode struct {
	Id        NodeId
	ParentIds []NodeId // Every ancestor id, all the way up to the root
	ChildIds  []NodeId // Every direct child id of this node
	Pairs     []Pair   // Key/value pairs
	Flags     []Flag   // Value-only flags
}

type Node struct {
	rawNode
	Parents  []*Node // Every ancestor, all the way up to the root
	Children []*Node // Every direct child of this node
}

type NodeId string

func (id NodeId) fromHash(h []byte) string {
	return base64.StdEncoding.EncodeToString(h)
}

type Pair struct {
	Key string
	Val string
}

type Flag string
