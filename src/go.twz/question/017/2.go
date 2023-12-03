package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	mySlice := []int{
		23, 32, 78, 43, 76, 65, 345, 762, 999, 998,
		555, 915, 86, 934, 351, 524, 241, 746, 260, 804,
		355, 805, 4, 840, 707, 872, 679, 463, 295, 431,
		674, 290, 60, 516, 370, 395, 345, 633, 318, 976,
		879, 109, 179, 962, 895, 321, 345, 27, 747, 599,
		651, 567, 483, 58, 156, 283, 956, 956, 980, 545,
		454, 667, 414, 158, 602, 717, 422, 701, 767, 514,
		696, 580, 425, 653, 763, 395, 401, 763, 676, 212,
		768, 106, 767, 442, 995, 593, 51, 941, 141, 800,
		782, 223, 734, 193, 940, 106, 269, 318, 929, 49,
	}
	cou := len(mySlice)
	fmt.Printf("slice长度是：%d \n", cou)
	ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	result := make(chan int)
	defer close(result)
	FindValue(mySlice, ctx, 345, result)
	timer := time.NewTimer(time.Second * 5)

	for {
		select {
		case <-timer.C:
			fmt.Printf("五秒超时结束 \n")
			cancel()
			return
		case res := <-result:
			if res > 0 {
				fmt.Printf("我们拿到了结果：index：%d \n", res)
				cancel()
				// close(result)
				time.Sleep(time.Second * 1)

				return
			}
		}
	}
}

func FindValue(mySlice []int, ctx context.Context, findValue int, res chan int) {
	sLen := len(mySlice)
	gNum := (sLen + 9) / 10

	for i := 0; i < 10; i++ {
		go func(times int, ctx context.Context) {
			for j := gNum * times; j < gNum*(times+1); j++ {
				select {
				case <-ctx.Done():
					fmt.Printf("分支查找被终止！\n")
					return
				default:
					fmt.Printf("正在查询！index = %d \n", j)
				}
				time.Sleep(time.Millisecond * 600)
				if j >= sLen {
					return
				}
				if mySlice[j] == findValue && res != nil {

					res <- j
					return
				}
			}

		}(i, ctx)
	}
}
