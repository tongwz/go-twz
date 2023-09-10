package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var pool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	}}

func main() {
	// runtime.GOMAXPROCS(1)
	go func() {
		for {
			processRequest(1 << 28) // 256MiB
		}
	}()
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				processRequest(1 << 10) // 1KiB
			}
		}()
	}
	var stats runtime.MemStats
	for i := 0; ; i++ {
		runtime.ReadMemStats(&stats)
		fmt.Printf("Cycle %d: %dB\n", i, stats.Alloc)
		time.Sleep(time.Second)
		runtime.GC()
	}
}
func processRequest(size int) {
	b := pool.Get().(*bytes.Buffer)
	time.Sleep(500 * time.Millisecond)
	b.Grow(size)
	pool.Put(b)
	time.Sleep(1 * time.Millisecond)
}

// A Pool is a set of temporary objects that may be individually saved and retrieved.
//
// Any item stored in the Pool may be removed automatically at any time without notification. If the Pool holds the only reference when this happens, the item might be deallocated.
//
// A Pool is safe for use by multiple goroutines simultaneously.
//
// Pool's purpose is to cache allocated but unused items for later reuse, relieving pressure on the garbage collector. That is, it makes it easy to build efficient, thread-safe free lists. However, it is not suitable for all free lists.
//
// An appropriate use of a Pool is to manage a group of temporary items silently shared among and potentially reused by concurrent independent clients of a package. Pool provides a way to amortize allocation overhead across many clients.
//
// An example of good use of a Pool is in the fmt package, which maintains a dynamically-sized store of temporary output buffers. The store scales under load (when many goroutines are actively printing) and shrinks when quiescent.
//
// On the other hand, a free list maintained as part of a short-lived object is not a suitable use for a Pool, since the overhead does not amortize well in that scenario. It is more efficient to have such objects implement their own free list.
//
// A Pool must not be copied after first use.
//
// In the terminology of the Go memory model, a call to Put(x) “synchronizes before” a call to Get returning that same value x. Similarly, a call to New returning x “synchronizes before” a call to Get returning that same value x.

// Grow grows the buffer's capacity, if necessary, to guarantee space for another n bytes. After Grow(n), at least n bytes can be written to the buffer without another allocation.
// If n is negative, Grow will panic.
// If the buffer can't grow it will panic with ErrTooLarge.
