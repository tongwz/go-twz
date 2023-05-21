package main

const (
	R = iota // 具体方向
	L
	F
	B
)
const (
	ZR = iota // 具体操作
	ZL
	ZF
	ZB
)

func main() {
	var x, y int
	x, y = x+0, y+1 // ZF, F
}

func move() {

}
