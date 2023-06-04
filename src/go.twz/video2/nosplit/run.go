package main

func main() {
	MySplit()
}

//go:nosplit
func MySplit() int {
	return MySplit()
}
