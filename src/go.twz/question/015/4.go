package main

import (
	"fmt"
	"sync"
	"time"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func(), wait *sync.WaitGroup) {
	defer wait.Done()
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

func main() {
	// 内存可见性（memory visibility）
	// 内存屏障（memory barrier）
	// ww := sync.Once{}
	wait := sync.WaitGroup{}
	for j := 0; j < 10000; j++ {
		wait.Add(10000)
		d := new(Once)

		f := func() {
			fmt.Printf("我执行了   ")
		}

		for i := 0; i < 10000; i++ {
			go d.Do(f, &wait)
		}
		wait.Wait()
		fmt.Printf("结束！ \n")
		if j%50 == 0 {
			time.Sleep(time.Second)
		}
	}

	time.Sleep(time.Millisecond * 1000)
}

/**
如果o.done设置为1后，解锁，那么另外一个goroutine会从o.m.Lock()这行后面开始执行，
它会执行if o.done == 0这个判断，这个判断应该永远不会成立。但是，这样做还是有一些问题。
第一个问题是，你还是需要每次调用Do函数都加锁和解锁，这会影响性能。第二个问题是，
你的代码可能会有内存可见性（memory visibility）的问题。
内存可见性是指一个goroutine对内存的修改是否对其他goroutine可见。
如果你的代码运行在多核处理器上，那么每个核心都有自己的缓存，它们可能不会及时同步内存的修改。
所以，即使o.done被设置为1了，其他的goroutine可能还是看到它为0，然后继续执行f函数。为了解决这个问题，
你需要使用原子操作（atomic operation）或者内存屏障（memory barrier）来保证内存的一致性。
你可以参考这篇文章来了解更多关于内存可见性和原子操作的内容。

Go语言里有atomic包，可以提供一些原子操作的函数，比如atomic.LoadInt32和atomic.StoreInt32等。
这些函数可以保证内存的一致性，但是它们不是内存屏障。内存屏障是一种特殊的指令，
它可以阻止屏障两侧的指令重排序，并强制把缓存中的数据同步到内存123。
Go语言目前没有提供专门的包或者函数来实现内存屏障，但是你可以使用cgo或者汇编来调用底层的指令45。
不过，这样做可能会降低代码的可移植性和可读性，所以你需要谨慎使用。
另外，Go语言的垃圾回收器也使用了一种写屏障技术，它可以在并行标记清理时保证对象的可达性6。
你可以参考这篇文章来了解更多关于写屏障技术的内容。

你说的例子是一个简化版的sync.Once，没有使用atomic包。这个例子中，o.done的值在解锁时也会写回内存，但是它可能会有一个问题。如果有两个goroutine同时调用o.Do(f)，并且o.done的初始值为0，那么可能会发生这样的情况：
goroutine A检查o.done的值为0，然后进入慢路径，加锁。
goroutine B检查o.done的值为0，然后等待锁。
goroutine A执行f()，然后把o.done的值设为1，解锁。
goroutine B拿到锁，然后再次检查o.done的值。
这时候，goroutine B看到的o.done的值可能还是0，因为它没有使用atomic.LoadUint32函数来保证内存一致性。所以，goroutine B也会执行f()，导致f()被执行了两次。这可能会引起一些错误或者副作用。所以，这个例子中的o.done需要使用atomic包来保证内存一致性和指令重排序。
*/
