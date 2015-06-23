package trog

import (
	"encoding/base64"
	"sort"
)

type NodeId string

func (id NodeId) fromHash(h []byte) string {
	return base64.StdEncoding.EncodeToString(h)
}

type listNodeId []NodeId

func sortNodeIds(list []NodeId)       { sort.Sort(listNodeId(list)) }
func (list listNodeId) Swap(i, j int) { list[i], list[j] = list[j], list[i] }
func (list listNodeId) Len() int      { return len(list) }
func (list listNodeId) Less(i, j int) bool {
	a := list[i]
	b := list[j]
	return a < b
}
