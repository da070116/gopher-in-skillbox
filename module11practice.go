package main

import (
	"bufio"
	"fmt"
	"os"
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

	fmt.Println("Введите строку для подсчёта слов с большой буквы")
	// Ввод произвольной строки
	inputSource := bufio.NewReader(os.Stdin)
	rawInput, _ := inputSource.ReadString('\n')

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
	inputSource := bufio.NewReader(os.Stdin)
	rawInput, _ := inputSource.ReadString('\n')

	digitsString := ""

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
	if len(digitsString) > 0 {
		fmt.Print("В строке содержатся следующие числа в десятичном формате:\n", digitsString)
	} else {
		fmt.Println("В строке нет чисел в десятичном формате.")
	}

}

func main() {
	countWords()
	digitsOutput()
}
