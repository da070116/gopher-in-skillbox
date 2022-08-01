package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Задание 1. Разложение ex в ряд Тейлора")
	/*
		E^x = Sum(x^n/n!)
	*/
	var x float64
	var myPrecious uint8
	fmt.Print("Введите точность (число символов после 0): ")
	_, _ = fmt.Scan(&myPrecious)

	fmt.Print("Введите значение переменной х: ")
	_, _ = fmt.Scan(&x)

	epsilon := 1 / math.Pow(10, x)

	funcResult := float64(0)
	funcPrevResult := float64(0)
	factorial := 1
	n := 0
	for {
		if n > 0 {
			factorial *= n
		}
		funcResult += math.Pow(x, float64(n)) / float64(factorial)
		if math.Abs(funcResult-funcPrevResult) < epsilon {
			fmt.Println(funcResult)
			break
		}
		n++
		funcPrevResult = funcResult
	}
}
