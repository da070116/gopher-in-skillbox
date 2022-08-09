package main

import (
	"fmt"
	"math/rand"
	"time"
)

func customRandInt(maxValue int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxValue)
}

func wrapper(f func()) {
	fmt.Println("==========================================")
	f()
	fmt.Println("==========================================")
}

func main() {
	m20RunAll()
}
