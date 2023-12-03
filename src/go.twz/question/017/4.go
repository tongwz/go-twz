package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	mySlice := []int{
		23, 32, 78, 43, 76, 65, 345, 762, 999, 998, 555, 915,
		86, 934, 351, 524, 241, 746, 260, 355, 805, 4, 840, 707, 872, 679,
		463, 295, 431, 674, 290, 60, 516, 370, 395, 699, 633, 318, 976, 804,
		879, 109, 179, 962, 895, 321, 403, 27, 747, 599, 651, 567, 483, 58,
		156, 283, 956, 956, 980, 545, 454, 667, 414, 158, 602, 717, 422,
		701, 767, 514, 696, 580, 425, 653, 763, 395, 401, 763, 676, 212,
		768, 106, 767, 442, 995, 593, 51, 941, 141, 800, 782, 223, 734,
		193, 940, 106, 269, 318, 929, 49, 400, 457, 334, 285, 466, 758,
		384, 634, 485, 509, 878, 14, 919,
	}
	cou := len(mySlice)

	fmt.Printf("我们的slice长度是 %d \n", cou)

	ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	result := make(chan int)
	go FindColumnValue(ctx, mySlice, 345, result)

	timer := time.NewTimer(time.Second * 5)
	for {
		select {
		case <-timer.C:
			fmt.Printf("五秒超时 结束~ \n")
			cancel()
			time.Sleep(time.Second * 2)
			return
		case res := <-result:
			fmt.Printf("我们拿到了结果：index=%d\n", res)
			cancel()
			time.Sleep(time.Second * 2)
			return
		}
	}
}

func FindColumnValue(ctx context.Context, mySlice []int, findValue int, res chan int) {
	sLen := len(mySlice)
	gNum := (sLen + 9) / 10

	for i := 0; i < 10; i++ {
		go func(times int, ctx context.Context) {
			for j := gNum * times; j < gNum*(times+1); j++ {
				select {
				case <-ctx.Done():
					fmt.Printf("分支查询被终止！\n")
					return
				default:
					time.Sleep(time.Millisecond * 500)
					if j >= sLen {
						return
					}
					if mySlice[j] == findValue {
						res <- j
						return
					}
				}
			}
		}(i, ctx)
	}
}
