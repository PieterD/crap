package transaction

import "sync"

type syncTransaction struct {
	l sync.Mutex
	t Transaction
}

func NewSyncTransaction(t Transaction) Transaction {
	return &syncTransaction{t: t}
}

func (t *syncTransaction) Add(key, value string) (ok bool) {
	t.l.Lock()
	defer t.l.Unlock()
	return t.t.Add(key, value)
}

func (t *syncTransaction) Set(key, value string) (ok bool) {
	t.l.Lock()
	defer t.l.Unlock()
	return t.t.Set(key, value)
}

func (t *syncTransaction) Del(key string) (value string, ok bool) {
	t.l.Lock()
	defer t.l.Unlock()
	return t.t.Del(key)
}

func (t *syncTransaction) Get(key string) (value string, ok bool) {
	t.l.Lock()
	defer t.l.Unlock()
	return t.t.Get(key)
}
