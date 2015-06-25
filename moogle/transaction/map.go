package transaction

type mapTransaction struct {
	m map[string]string
}

func NewMapTransaction() Transaction {
	return &mapTransaction{
		m: make(map[string]string),
	}
}

func (t *mapTransaction) Add(key, value string) (ok bool) {
	_, ok = t.m[key]
	if ok {
		return false
	}
	t.m[key] = value
	return true
}

func (t *mapTransaction) Set(key, value string) (ok bool) {
	_, ok = t.m[key]
	if !ok {
		return false
	}
	t.m[key] = value
	return true
}

func (t *mapTransaction) Del(key string) (value string, ok bool) {
	value, ok = t.m[key]
	delete(t.m, key)
	return
}

func (t *mapTransaction) Get(key string) (value string, ok bool) {
	value, ok = t.m[key]
	return
}
