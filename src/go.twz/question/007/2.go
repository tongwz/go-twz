package main

func main() {

}

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student:
		_ = msg.Name
	}
}
