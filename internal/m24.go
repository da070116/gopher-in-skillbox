package internal

import (
	"fmt"
	"gopher-in-skillbox/pkg"
	"math/rand"
	"time"
)

func m24Task1() {
	//Написать функцию для сортировки вставками массива.
	const arraySize = 10

	data := make([]int, arraySize)
	for i := 0; i < arraySize; i++ {
		data[i] = pkg.CustomRandInt(100)
	}
	fmt.Printf("Неотсортированный массив: %v\n", data)

	for i := 1; i < len(data); i++ {
		key := data[i]
		lst := i
		for j := i - 1; j > -1; j-- {
			if data[j] < key {
				break
			}
			data[j+1] = data[j]
			lst = j
		}
		data[lst] = key
	}

	fmt.Printf("Отсортированный массив: %v\n", data)
}

func m24Task2() {
	// Реализовать анонимную функцию, которая на вход получает массив типа integer, сортирует его пузырьком и переворачивает

	// подготовка и заполнение массива
	rand.Seed(time.Now().UnixNano())
	var arraySize int
	fmt.Printf("Введите размер массива: ")
	_, _ = fmt.Scan(&arraySize)
	arrSlice := make([]int, arraySize)
	for i := 0; i < arraySize; i++ {
		arrSlice[i] = pkg.CustomRandInt(100)
	}

	// анонимная функция
	bubbleSort := func(raw []int) []int {
		if len(raw) == 0 {
			panic("Недостаточно аргументов")
		}
		for i := len(raw); i > 0; i-- {
			for j := 1; j < i; j++ {
				if raw[j-1] < raw[j] {
					raw[j], raw[j-1] = raw[j-1], raw[j]
				}
			}
		}
		return raw
	}

	fmt.Printf("Неотсортированный массив: %v\n", arrSlice)
	fmt.Printf("Отсортированный по убыванию массив: %v\n", bubbleSort(arrSlice))
}

func m24RunAll() {
	pkg.Wrapper(m24Task1)
	pkg.Wrapper(m24Task2)
}
