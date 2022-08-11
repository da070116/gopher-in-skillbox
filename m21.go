package main

import (
	"fmt"
	"math/rand"
	"time"
)

func m22Task1() {
	var findMe int
	arr := fillArray(100)
	fmt.Println(arr)
	fmt.Print("Введите число, которое нужно найти: ")
	_, _ = fmt.Scan(&findMe)
	position, flag := findInArray(findMe, arr)
	if flag {
		fmt.Printf("В массиве %v найдено %v чисел после заданного %v\n", arr, uint8(len(arr))-position-1, findMe)
	} else {
		fmt.Println("Число не найдено")
	}

}

func findInArray(needle int, haystack [10]int) (resultIndex uint8, isFound bool) {
	for idx, val := range haystack {
		if val == needle {
			resultIndex = uint8(idx)
			isFound = true
			break
		}
	}
	return
}

func fillArray(maxDigit int) (resultArray [10]int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(resultArray); i++ {
		resultArray[i] = customRandInt(maxDigit)
	}
	return
}

func m22Task2() {

}

func m22RunAll() {
	wrapper(m22Task1)
	wrapper(m22Task2)
}
