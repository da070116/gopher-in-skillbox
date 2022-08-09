package main

import (
	"fmt"
	"math/rand"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	m20RunAll()
}
