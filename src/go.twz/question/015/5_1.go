package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	const (
		mutexLocked = 1 << iota // mutex is locked
		mutexWoken
		mutexStarving
		mutexWaiterShift      = iota
		starvationThresholdNs = 1e6
	)

	fmt.Printf("mutexLocked = %d, mutexWoken = %d , mutexStarving = %d, mutexWaiterShift = %d, starvationThresholdNs = %d \n", mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift, starvationThresholdNs)

	var state int32 = 0

	// 这个函数通过原子性操作 先看state值是否 == old值 如果等于 就将state = new 之后返回true 反之返回false 不改变 state的值
	if atomic.CompareAndSwapInt32(&state, 1, mutexWoken) {
		fmt.Printf("进入 atomic.CompareAndSwapInt32  state = %d \n", state)
	} else {
		fmt.Printf("没进入 atomic.CompareAndSwapInt32  state = %d \n", state)
	}

}
