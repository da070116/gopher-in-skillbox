package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(raw []int, asc bool) (sorted []int) {
	for i := len(raw); i > 0; i-- {
		for j := 1; j < i; j++ {
			var clause bool
			if asc {
				clause = raw[j-1] < raw[j]
			} else {
				clause = raw[j-1] > raw[j]
			}
			if clause {
				raw[j], raw[j-1] = raw[j-1], raw[j]
			}
		}
	}
	sorted = raw
	return
}

func m24Task1() {
	rand.Seed(time.Now().UnixNano())
	var arraySize int
	fmt.Printf("Введите размер массива: ")
	_, _ = fmt.Scan(&arraySize)
	arrSlice := make([]int, arraySize)
	for i := 0; i < arraySize; i++ {
		arrSlice[i] = customRandInt(100)
	}
	fmt.Printf("Неотсортированный массив: %v\n", arrSlice)
	fmt.Printf("Отсортированный по убыванию массив: %v\n", bubbleSort(arrSlice, true))
	fmt.Printf("Отсортированный по возрастанию массив: %v\n", bubbleSort(arrSlice, false))
}

func m24RunAll() {
	wrapper(m24Task1)
}
