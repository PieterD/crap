package trog

import "sort"

type Flag string

type listFlag []Flag

func sortFlags(list []Flag)         { sort.Sort(listFlag(list)) }
func (list listFlag) Swap(i, j int) { list[i], list[j] = list[j], list[i] }
func (list listFlag) Len() int      { return len(list) }
func (list listFlag) Less(i, j int) bool {
	a := list[i]
	b := list[j]
	return a < b
}
