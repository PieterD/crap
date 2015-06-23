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

type PairMap map[string]string

func (pm PairMap) Get(key string) string {
	return pm[key]
}

func (pm PairMap) Set(key, val string) {
	pm[key] = val
}

func (pm PairMap) Del(key string) {
	delete(pm, key)
}

func (pm PairMap) Has(key string) bool {
	_, ok := pm[key]
	return ok
}

func (pm PairMap) List() []Pair {
	list := make([]Pair, 0, len(pm))
	for key, val := range pm {
		list = append(list, Pair{Key: key, Val: val})
	}
	sortPairs(list)
	return list
}
