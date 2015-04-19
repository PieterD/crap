package killchan

import "sync/atomic"

// A Killchan is a chan struct{} that can be closed multiple times without panicing.
type Killchan struct {
	isclosed  int32
	closechan chan struct{}
}

// Return a new killchan.
func New() *Killchan {
	return &Killchan{
		closechan: make(chan struct{}),
	}
}

// Kill the channel. Returns true if the channel has been killed,
// false if it had already been killed previously.
func (kc *Killchan) Kill() bool {
	if atomic.CompareAndSwapInt32(&kc.isclosed, 0, 1) {
		close(kc.closechan)
		return true
	}
	return false
}

// Return the channel.
func (kc *Killchan) Chan() <-chan struct{} {
	return kc.closechan
}

// Wait for the channel to be killed.
func (kc *Killchan) Wait() {
	<-kc.closechan
}

// Return true if the channel is still alive, false if it has been killed.
func (kc *Killchan) Alive() bool {
	select {
	case <-kc.Chan():
		return false
	default:
		return true
	}
}
