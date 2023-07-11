package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type sp1 interface {
	Out(key string, val interface{})                  // 存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} // 读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type safeMap struct {
	info map[string]*realInfo
	sync.RWMutex
}

type realInfo struct {
	value interface{}
	ch    chan bool
}

func (s *safeMap) Out(key string, val interface{}) {
	s.Lock()
	defer s.Unlock()
	if pointer, ok := s.info[key]; ok {
		pointer.value = val
		// 放入值
		pointer.ch <- true
	} else {
		s.info[key] = &realInfo{
			value: val,
			ch:    make(chan bool, 1),
		}
		s.info[key].ch <- true
	}
}

func (s *safeMap) Rd(key string, timeout time.Duration) interface{} {
	// 读取数据
	s.Lock()
	fmt.Println("我们获取到了读锁", key)
	if pointer, ok := s.info[key]; ok {
		s.Unlock()
		return pointer.value
	} else {
		// 如果没有值
		r := &realInfo{
			value: nil,
			ch:    make(chan bool, 1),
		}
		s.info[key] = r

		select {
		case <-r.ch:
			return r.value
		case <-time.After(timeout):
			fmt.Printf("超时没有查询到 key = %s 的值 \n", key)
		}
		s.Unlock()
		return nil
	}

}

func main() {
	safeV := &safeMap{
		info: make(map[string]*realInfo),
	}

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("读取的值是：%#v \n", safeV.Rd(strconv.Itoa(i), time.Second))
		}
	}()

	go func() {
		fmt.Printf("读取的值是：%#v \n", safeV.Rd(strconv.Itoa(11), time.Second))
	}()
	go func() {
		fmt.Printf("读取的值是：%#v \n", safeV.Rd(strconv.Itoa(11), time.Second))
	}()
	go func() {
		fmt.Printf("读取的值是：%#v \n", safeV.Rd(strconv.Itoa(11), time.Second))
	}()
	go func() {
		fmt.Printf("读取的值是：%#v \n", safeV.Rd(strconv.Itoa(11), time.Second))
	}()
	go func() {
		fmt.Printf("读取的值是：%#v \n", safeV.Rd(strconv.Itoa(11), time.Second))
	}()

	go func() {
		for i := 9; i >= 0; i-- {
			safeV.Out(strconv.Itoa(i), strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 11)
}
