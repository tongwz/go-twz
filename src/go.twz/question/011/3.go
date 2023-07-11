package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
	sync.RWMutex
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}
func (o *Ban) visit(ip string) bool {
	o.Lock()
	defer o.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	go o.delMap(ip)
	return false
}

func (o *Ban) delMap(ip string) {
	timer := time.NewTicker(time.Minute * 3)
	select {
	case <-timer.C:
		o.Lock()
		defer o.Unlock()
		delete(o.visitIPs, ip)
		fmt.Printf("我们删除了key : %s \n", ip)
	}
}
func main() {
	var success int64 = 0
	ban := NewBan()
	wait := sync.WaitGroup{}
	wait.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wait.Done()
				waitTime := time.Duration(rand.Int63n(3))
				time.Sleep(time.Second * waitTime)
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}

	}
	wait.Wait()
	fmt.Println("success:", success)

}
