package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ban1 struct {
	visitIPs map[string]time.Time
	sync.RWMutex
}

func NewBan1() *Ban1 {
	return &Ban1{visitIPs: make(map[string]time.Time)}
}
func (o *Ban1) visit(ip string) bool {
	o.Lock()
	defer o.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	go o.delTimeOut(ip)
	return false
}

func (o *Ban1) delTimeOut(ip string) {
	ticker := time.NewTicker(time.Second * 180)
	select {
	case <-ticker.C:
		o.Lock()
		defer o.Unlock()
		fmt.Printf("我执行了删除key：%s \n", ip)
		delete(o.visitIPs, ip)
	}
}
func main() {
	var success int64 = 0
	Ban1 := NewBan1()
	wait := sync.WaitGroup{}
	wait.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wait.Done()
				// randd := rand.Int63n(2)
				// time.Sleep(time.Millisecond * time.Duration(randd))
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !Ban1.visit(ip) {
					// fmt.Printf("查看randd的值：%d \n", randd)
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}

	}
	wait.Wait()
	fmt.Println("success:", success)
}
