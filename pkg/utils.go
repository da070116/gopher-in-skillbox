package pkg

import (
	"fmt"
	"math/rand"
)

// CustomRandInt  - создать случайную переменную типа int от 0 до maxValue
func CustomRandInt(maxValue int) int {
	return rand.Intn(maxValue)
}

// Wrapper  - обернуть вывод функции в "="
func Wrapper(f func()) {
	fmt.Println("==========================================")
	f()
	fmt.Println("==========================================")
}
