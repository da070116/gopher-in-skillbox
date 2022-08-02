package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func isCapitalized(word string) bool {
	return unicode.IsUpper(rune(word[0]))
}

func countWords() {
	fmt.Println("Задание 1. Определение количества слов, начинающихся с большой буквы")
	capitalizedWords := 0
	var rawInput string
	fmt.Println("Введите строку для подсчёта слов с большой буквы")
	_, _ = fmt.Scan(&rawInput)
	for len(rawInput) > 0 {
		spaceIndex := strings.Index(rawInput, " ")
		if spaceIndex == -1 {
			break
		}
		if isCapitalized(rawInput[:spaceIndex]) {
			capitalizedWords++
		}
		rawInput = rawInput[spaceIndex+1:]
	}

	if isCapitalized(rawInput) {
		capitalizedWords++
	}

	fmt.Printf("Строка содержит %v слов с большой буквы\n", capitalizedWords)
	fmt.Println("=============================================================")
}

func digitsOutput() {
	var rawInput = "a10 10 20b 20 30c30 30 dd"
	digitsString := ""
	fmt.Println(rawInput)

	for len(rawInput) > 0 {
		spaceIndex := strings.Index(rawInput, " ")
		if spaceIndex == -1 {
			break
		}
		word := rawInput[:spaceIndex]
		_, err := strconv.ParseInt(word, 10, 0)
		if err == nil {
			digitsString += word + " "
		}
		rawInput = rawInput[spaceIndex+1:]
	}
	fmt.Print("В строке содержатся числа в десятичном формате:\n", digitsString)
}

func main() {
	countWords()
	digitsOutput()
}
