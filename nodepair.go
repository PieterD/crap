package trog

import "sort"

type Pair struct {
	Key string
	Val string
}

type listPair []Pair

func sortPairs(list []Pair)         { sort.Sort(listPair(list)) }
func (list listPair) Swap(i, j int) { list[i], list[j] = list[j], list[i] }
func (list listPair) Len() int      { return len(list) }
func (list listPair) Less(i, j int) bool {
	a := list[i]
	b := list[j]
	if a.Key < b.Key {
		return true
	}
	if a.Val < b.Val {
		return true
	}
	return false
}
