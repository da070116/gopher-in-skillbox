package main

import (
	"fmt"
	"math"
)

func module9Task1() {
	/*
		В цикле с использованием встроенных констант
		(на предельные значения целых чисел, в пакете math) программа будет подсчитывать,
		сколько приходится переполнений чисел типа uint8, uint16 в диапазоне от 0 до uint32.
	*/

	fmt.Println("Задание 1. Переполнение")

	var uint8Counter uint8
	var uint16Counter uint16
	var uint32Counter uint32

	uint8OverflowCounts := 0
	uint16OverflowCounts := 0

	for uint32Counter = 0; uint32Counter < math.MaxUint32; uint32Counter++ {
		if uint8Counter == math.MaxUint8 {
			uint8Counter = 0
			uint8OverflowCounts++
		}
		if uint16Counter == math.MaxUint16 {
			uint16Counter = 0
			uint16OverflowCounts++
		}
		uint8Counter++
		uint16Counter++
	}
	fmt.Printf("На интервале от 0 до %v есть %v переполнений типа uint8 и %v переполнений типа uint16\n", math.MaxUint32, uint8OverflowCounts, uint16OverflowCounts)
}
func module9Task2() {
	var digit1, digit2 int16
	var resultType string
	fmt.Print("Введите первое число:")
	_, _ = fmt.Scan(&digit1)
	fmt.Print("Введите второе число:")
	_, _ = fmt.Scan(&digit2)

	multiply := int32(digit1) * int32(digit2)

	switch {
	case multiply <= math.MaxUint8 && multiply >= 0:
		resultType = "uint8"
	case multiply <= math.MaxUint16 && multiply >= 0:
		resultType = "uint16"
	case multiply <= math.MaxInt8 && multiply >= math.MinInt8:
		resultType = "int8"
	case multiply <= math.MaxInt16 && multiply >= math.MinInt16:
		resultType = "int16"
	case multiply >= 0:
		resultType = "uint32"
	default:
		resultType = "int32"
	}
	fmt.Printf("Результат произведения %v и %v может быть записан как %v и равен %v\n", digit1, digit2, resultType, multiply)
}

func runAllTasksModule9() {
	wrapper(module9Task1)
	wrapper(module9Task2)
}
