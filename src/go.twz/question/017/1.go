package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var sData []int
	for i := 11; i < 125; i++ {
		sData = append(sData, i)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	isFind := make(chan int)
	go FindNeedNum(ctx, sData, 100, isFind)
	for {
		select {
		case val := <-isFind:
			if val >= 0 {
				fmt.Printf("找到了数字index = %d 结束代码 \n", val)
				cancel()
				time.Sleep(time.Second * 3)
				fmt.Println("结束")
				return
			}
		}
	}

}

func FindNeedNum(ctx context.Context, sData []int, num int, isFind chan int) {
	sLen := len(sData)
	// 进一取整
	gNum := (sLen + 10 - 1) / 10

	// 分10个协程处理
	for i := 0; i < 10; i++ {
		go func(times int, ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("程序被Done终止  \n")
					return
				default:
					// 按照索引查询数据
					for j := gNum * times; j < gNum*(times+1); j++ {
						time.Sleep(time.Second * 1)
						if j >= sLen {
							return
						}
						fmt.Println(sData[j])
						if sData[j] == num {
							isFind <- j
							return
						}
					}
				}
			}

		}(i, ctx)
	}
}
