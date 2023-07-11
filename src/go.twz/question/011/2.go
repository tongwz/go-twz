package main

import (
	"fmt"
	"sync"
	"time"
)

type tMap struct {
	mm map[string]string
	sync.RWMutex
}

func (t *tMap) Write(ip string) {
	t.Lock()
	defer func() {
		fmt.Printf("解锁了 外部Lock \n")
		t.Unlock()
	}()
	if v, ok := t.mm[ip]; ok {
		fmt.Printf("已经有值了%s \n", v)
	}
	t.mm[ip] = ip
	go t.delTimeOut(ip)
	go t.delTimeOut(ip)
}

func (t *tMap) delTimeOut(ip string) {

	t.Lock()
	time.Sleep(time.Second * 2)

	defer t.Unlock()
	if _, ok := t.mm[ip]; ok {
		delete(t.mm, ip)
		fmt.Printf("我执行了删除key：%s \n", ip)
	} else {
		fmt.Printf("我没有执行删除key：%s \n", ip)
	}

}

func main() {

	ban := new(tMap)
	ban.mm = make(map[string]string)
	ban.Write("192.168.11.10")

	time.Sleep(time.Second * 5)
}
