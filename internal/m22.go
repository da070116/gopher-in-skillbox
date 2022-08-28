package internal

import (
	"fmt"
	"gopher-in-skillbox/pkg"
	"math/rand"
	"sort"
	"time"
)

const (
	ARRAY1SIZE = 10
	ARRAY2SIZE = 12
)

func findInArray(needle int, haystack [ARRAY1SIZE]int) (resultIndex int8) {
	/*
		Find an int value in unsorted int array of size 10. Returns index and flag whether value was found.
	*/
	resultIndex = -1
	for idx, val := range haystack {
		if val == needle {
			resultIndex = int8(idx)
			break
		}
	}
	return
}

func fillArray(maxDigit int) (resultArray [ARRAY1SIZE]int) {
	// Generate random values from 0 to maxDigit for integer array of size 10.
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(resultArray); i++ {
		resultArray[i] = pkg.CustomRandInt(maxDigit)
	}
	return
}

func fillSortedArray() (resultArray [ARRAY2SIZE]int) {
	// Generate random values from 0 to 10 for integer array of size 12.
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(resultArray); i++ {
		resultArray[i] = pkg.CustomRandInt(10)
	}
	sort.Ints(resultArray[:])
	return
}

func findInSorted(haystack [ARRAY2SIZE]int, needle int) (index int8) {
	index = int8(-1)
	min := int8(0)
	max := int8(ARRAY2SIZE - 1)

	for min <= max {
		median := (max + min) / 2
		switch {
		case haystack[median] > needle:
			max = median - 1
		case haystack[median] < needle:
			min = median + 1
		default:
			index = median
			max = median - 1
		}
	}
	return
}

func m22Task1() {
	// Программа должна найти это число в массиве и вывести, сколько чисел находится в массиве после введённого.
	var findMe int
	arr := fillArray(100)
	fmt.Println(arr)
	fmt.Print("Введите число, которое нужно найти: ")
	_, _ = fmt.Scan(&findMe)
	position := findInArray(findMe, arr)
	if position > 0 {
		fmt.Printf("В массиве %v найдено %v чисел после заданного %v\n", arr, int8(len(arr))-position-1, findMe)
	} else {
		fmt.Println("Число не найдено")
	}

}

func m22Task2() {
	// Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться)
	arr := fillSortedArray()
	var digit int
	fmt.Println(arr)
	fmt.Println("Введите число, которое нужно найти в массиве: ")
	_, _ = fmt.Scan(&digit)

	searchResult := findInSorted(arr, digit)
	if searchResult == -1 {
		fmt.Printf("В массиве %v элемента %v не найдено", arr, digit)
	} else {
		fmt.Printf("В массиве %v первое вхождение элемента %v на позиции %v\n", arr, digit, searchResult+1)
	}
}

func m22RunAll() {
	pkg.Wrapper(m22Task1)
	pkg.Wrapper(m22Task2)
}
