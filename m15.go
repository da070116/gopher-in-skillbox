package main

import "fmt"

func inputArray() (array [10]int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Введите %v-й из 10 элемент массива: ", i+1)
		_, _ = fmt.Scan(&array[i])
	}
	return
}

func reverse(arrayPtr *[10]int) (reversed [10]int) {
	array := *arrayPtr
	reversed = array
	for idx, val := range array {
		reversed[len(reversed)-idx-1] = val
	}
	return
}

func m15Task1() {
	fmt.Println("Задание 1. Подсчёт чётных и нечётных чисел в массиве")
	mainArray := inputArray()
	var oddNum, evenNum int

	for _, val := range mainArray {
		if val%2 == 0 {
			evenNum++
		} else {
			oddNum++
		}
	}
	fmt.Printf("Для массива %v чётных %v,  нечётных %v \n", mainArray, evenNum, oddNum)
}

func m15Task2() {
	fmt.Println("Функция, реверсирующая массив")
	mainArray := inputArray()
	fmt.Println(mainArray, reverse(&mainArray))
}

func m15RunAll() {
	wrapper(m15Task1)
	wrapper(m15Task2)
}
