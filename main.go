package main

import (
	"fmt"
	"math/rand"
)

func customRandInt(maxValue int) int {
	return rand.Intn(maxValue)
}

func wrapper(f func()) {
	fmt.Println("==========================================")
	f()
	fmt.Println("==========================================")
}

func main() {

	m23RunAll()
}
