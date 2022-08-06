package main

import "fmt"

func wrapper(f func()) {
	fmt.Println("==========================================")
	f()
	fmt.Println("==========================================")
}

func main() {
	m13RunAll()
}
