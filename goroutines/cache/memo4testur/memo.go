package memo4testur

import (
	"log"
	"sync"
)

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//!+
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.RWMutex // guards cache
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {

	memo.mu.RLock()
	e := memo.cache[key]
	if e != nil && e.res.value != nil && e.res.err == nil {
		memo.mu.RUnlock()
		return e.res.value, e.res.err
	}
	memo.mu.RUnlock()

	memo.mu.Lock()
	e = memo.cache[key]
	if e == nil || (e != nil && e.res.err != nil) {
		if e != nil && e.res.err != nil {
			log.Printf("retry  %v\n", e.res.err)
		}
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		value, err = memo.f(key)
		memo.mu.Lock()
		e.res.value, e.res.err = value, err
		memo.mu.Unlock()

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}
