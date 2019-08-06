package spiders

import "sync"

var mux sync.WaitGroup

func WaitFinish() {
	mux.Wait()
}
