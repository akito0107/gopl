package ex03

import (
	"log"
	"net/url"
	"strings"
)

type request struct {
	key      string
	response chan<- result
	done     <-chan struct{}
}

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	requests chan request
}

type Done <-chan struct{}

func (d Done) canceled() bool {
	select {
	case <-d:
		return true
	default:
		return false
	}
}

type Func func(key string, done Done) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func, done <-chan struct{}) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done Done) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, done}
	if done.canceled() {
		return nil, nil
	}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key, req.done)

			<-e.ready

			if IsCanceled(e.res.err) {
				delete(cache, req.key)
			}
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string, done Done) {
	value, err := f(key, done)
	e.res.value, e.res.err = value, err
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}

func IsCanceled(err error) bool {
	if err == nil {
		return false
	}
	log.Printf("called: canceled %+v\n", err.Error())
	e, ok := err.(*url.Error)
	if !ok {
		log.Printf("called: canceled ok=false %+v\n", err.Error())
		return false
	}
	if strings.Contains(e.Error(), "cancel") {
		return true
	}
	return false
}
