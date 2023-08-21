package main

import "fmt"

func main() {
	const (
		mutexLocked = 1 << iota // mutex is locked
		mutexWoken
		mutexStarving
		mutexWaiterShift      = iota
		starvationThresholdNs = 1e6
	)

	fmt.Printf("mutexLocked = %d, mutexWoken = %d , mutexStarving = %d, mutexWaiterShift = %d, starvationThresholdNs = %d", mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift, starvationThresholdNs)
}
