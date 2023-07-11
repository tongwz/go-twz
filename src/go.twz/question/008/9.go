package main

import (
	"fmt"
	"sync"
)

type threadSafeSet struct {
	s chan interface{}
	sync.Mutex
}

func (set *threadSafeSet) RLock() {
	set.Lock()
}

func (set *threadSafeSet) RUnlock() {
	set.Unlock()
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func main() {
	th := new(threadSafeSet)
	th.s = make(chan interface{})
	go func() {
		th.s <- "ttt"
	}()

	ch := th.Iter()

	info := <-ch

	fmt.Printf("获取的值是： %#v \n", info)
}
