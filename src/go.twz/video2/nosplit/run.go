package main

func main() {
	split()
}

//go:nosplit
func split() int {
	return split()
}
