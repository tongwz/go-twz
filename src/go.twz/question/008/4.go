package main

type People struct{}

func (p *People) ShowA() {
	println("showA")
	// fmt.Printf("当前对象：%#v \n", p)
	p.ShowB()
}
func (p *People) ShowB() {
	println("showB")
}

type Teacher struct {
	P1 People
	People2
}

func (t *Teacher) ShowB() {
	println("teacher showB")
}

func main() {
	t := Teacher{}

	println("获取teacher中People地址：%p", &t.People2)

	t.ShowA()
}

type People2 struct{}

func (p *People2) ShowA() {
	println("show People2 A")
	// fmt.Printf("当前对象：%#v \n", p)
	p.ShowB()
}
func (p *People2) ShowB() {
	println("show People2 B")
}
