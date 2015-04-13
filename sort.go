package trog

import "sort"

type listNodeId []NodeId

func sortNodeIds(list []NodeId) {
	sort.Sort(listNodeId(list))
}

func (list listNodeId) Less(i, j int) bool {
	a := list[i]
	b := list[j]
	for idx := 0; idx < 32; idx++ {
		if a[idx] < b[idx] {
			return true
		}
		if a[idx] > b[idx] {
			return false
		}
	}
	return false
}

func (list listNodeId) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list listNodeId) Len() int {
	return len(list)
}

type listPair []Pair

func sortPairs(list []Pair) {
	sort.Sort(listPair(list))
}

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

func (list listPair) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list listPair) Len() int {
	return len(list)
}

type listFlag []Flag

func sortFlags(list []Flag) {
	sort.Sort(listFlag(list))
}

func (list listFlag) Less(i, j int) bool {
	a := list[i]
	b := list[j]
	if a < b {
		return true
	}
	return false
}

func (list listFlag) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list listFlag) Len() int {
	return len(list)
}
