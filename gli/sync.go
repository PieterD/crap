package gli

import "unsafe"

type Sync struct {
	ctx *Context
	p   unsafe.Pointer
}

func (ctx *Context) Fence() *Sync {
	p := ctx.r.SyncFence()
	return &Sync{ctx: ctx, p: p}
}

func (sync *Sync) ClientWait(flush bool, nanos uint64) iSyncResult {
	rv := sync.ctx.r.SyncClientWait(sync.p, flush, nanos)
	return iSyncResult{rv}
}
