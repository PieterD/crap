package trog

import "sort"

type FlagMap map[string]struct{}

func (fm FlagMap) Set(flag string) {
	fm[flag] = struct{}{}
}

func (fm FlagMap) Del(flag string) {
	delete(fm, flag)
}

func (fm FlagMap) Has(flag string) bool {
	_, ok := fm[flag]
	return ok
}

func (fm FlagMap) List() []string {
	list := make([]string, 0, len(fm))
	for key := range fm {
		list = append(list, key)
	}
	sort.Strings(list)
	return list
}
