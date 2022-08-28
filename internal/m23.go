package internal

import (
	"fmt"
	"gopher-in-skillbox/pkg"
	"math/rand"
	"time"
	"unicode/utf8"
)

const fixedSize = uint8(55)

func generateArray() (array [fixedSize]int) {
	rand.Seed(time.Now().UnixNano())
	for i := range array {
		array[i] = pkg.CustomRandInt(int(fixedSize))
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

func findChar(char rune, reversed string) (index int) {
	index = -1 // значение по умолчанию 0 приводит к некорректному отображению, когда искомый символ первый в предложении
	for i, r := range reversed {
		if r == char {
			index = len(reversed) - i - 1
			break
		}
	}
	return
}

func parseTest(sentences []string, chars []rune) (foundIndices [][]int) {
	if len(sentences) == 0&len(chars) {
		panic("Один из переданных массивов пуст")
	}
	for _, sentence := range sentences {
		var indicesInSentence []int
		for _, char := range chars {
			indicesInSentence = append(indicesInSentence, findChar(char, reverseString(sentence)))
		}
		foundIndices = append(foundIndices, indicesInSentence)
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
		if letter == "_" {
			break
		}
		letters = append(letters, rune(letter[0]))
	}
	fmt.Println(letters)

	fmt.Println(parseTest(sentences, letters))
}

func m23RunAll() {
	pkg.Wrapper(m23Task1)
	pkg.Wrapper(m23Task2)
}
