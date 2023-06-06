package main

import (
	"fmt"
	"time"
)

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}

func (p *Project) exec(msgchan chan interface{}) {
	defer p.deferError()
	for msg := range msgchan {
		m := msg.(int)
		fmt.Println("msg: ", m)
	}
}

func (p *Project) run(msgchan chan interface{}) {
	i := 0
	for {
		go p.exec(msgchan)
		fmt.Println("我们循环了", i)
		i++
		time.Sleep(time.Second * 2)
	}

	fmt.Println("aaa")
}

func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			a <- "1"
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 1000000)
}

func main() {
	p := new(Project)

	var i = 1<<63 - 1

	fmt.Printf("值是：%d \n", i)
	p.Main()

}

func p() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
	}()
	fmt.Println("a")
	panic("异常信息")
	fmt.Println("b") // 这里开始下面代码不会再执行
}
