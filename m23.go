package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"
)

const fixedSize = uint8(55)

func generateArray() (array [fixedSize]int) {
	rand.Seed(time.Now().UnixNano())
	for i := range array {
		array[i] = customRandInt(int(fixedSize))
	}

	return
}

func splitterArray(raw [fixedSize]int) (odds, evens []int) {
	for _, v := range raw {
		if v%2 == 0 {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}
	return
}

func m23Task1() {
	array := generateArray()
	odds, evens := splitterArray(array)
	fmt.Printf("Массив\n%v\nсодержит нечётные элементы:\n%v\nи чётные элементы:\n%v\n", array, odds, evens)
}

func parseTest(sentences []string, chars []rune) (foundIndices [][]int) {
	if len(sentences) == 0&len(chars) {
		panic("Один из переданных массивов пуст")
	}
	for s, sentence := range sentences {
		fmt.Println(s, ":", reverseString(sentence))
		for l, letter := range sentence {
			fmt.Printf("%v %v\n", l, letter)
		}
	}
	return
}

func reverseString(raw string) (result string) {
	size := len(raw)
	buffer := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(raw[start:])
		start += n
		utf8.EncodeRune(buffer[size-start:], r)
	}
	result = string(buffer)
	return
}

func m23Task2() {
	var sentences []string
	var letters []rune
	for {
		fmt.Print("Введите строчку для массива предложений или нажмите Enter для завершения ввода: ")
		sentence := readStringFromConsole()
		if sentence == "" {
			break
		}
		sentences = append(sentences, sentence)
	}
	fmt.Println(sentences)

	for {
		var letter string
		fmt.Print("Введите символ для массива букв или '_' для завершения ввода: ")
		_, _ = fmt.Scan(&letter)
		fmt.Println(letter)
		if letter == "_" {
			break
		}
		letters = append(letters, rune(letter[0]))
	}
	fmt.Println(letters)

	parseTest(sentences, letters)
}

func m23RunAll() {
	wrapper(m23Task1)
	wrapper(m23Task2)
}
