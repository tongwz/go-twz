package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type sp interface {
	Wt(key string, val interface{})                   // 存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} // 读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type mySafeMap struct {
	info map[string]*myStruct
	sync.RWMutex
}

type myStruct struct {
	value interface{}
	ch    chan bool
	isHad bool
}

func (m *mySafeMap) Wt(key string, val interface{}) {
	m.Lock()
	defer m.Unlock()
	if v, ok := m.info[key]; ok {
		v.value = val
		if !v.isHad {
			close(v.ch)
		}
		v.isHad = true
	} else {
		mm := &myStruct{
			value: val,
			ch:    make(chan bool, 1),
			isHad: true,
		}
		m.info[key] = mm
		close(mm.ch)
	}
}

func (m *mySafeMap) Rd(key string, timeout time.Duration) interface{} {
	m.Lock()
	if v, ok := m.info[key]; ok {
		if v.isHad == true {
			m.Unlock()
			return v.value
		} else {
			m.Unlock()
			select {
			case <-v.ch:
				return v.value
			case <-time.After(timeout):
				fmt.Printf("我们读取超时了 key= %s \n", key)
				return nil
			}
		}

	} else {
		mm := &myStruct{
			value: nil,
			ch:    make(chan bool, 1),
			isHad: false,
		}
		m.info[key] = mm
		m.Unlock()
		select {
		case <-mm.ch:
			return mm.value
		case <-time.After(timeout):
			fmt.Printf("我们读取超时了 key= %s \n", key)
			return nil
		}
	}
}

func main() {
	mmm := &mySafeMap{
		info: make(map[string]*myStruct),
	}

	go func() {
		for i := 0; i < 10; i++ {
			mmm.Wt(strconv.Itoa(i), i)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			mmm.Wt(strconv.Itoa(i), i+1)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 9; i >= 0; i-- {
			fmt.Println(mmm.Rd(strconv.Itoa(i), time.Second))
		}
	}()

	time.Sleep(time.Second * 11)
}
