package main

import (
	"fmt"
	"math/rand"
)

func customRandInt(maxValue int) int {
	// customRandInt - создать случайную переменную типа int от 0 до maxValue
	return rand.Intn(maxValue)
}

func wrapper(f func()) {
	// wrapper - обернуть вывод функции в "="
	fmt.Println("==========================================")
	f()
	fmt.Println("==========================================")
}
