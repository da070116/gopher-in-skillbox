package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func getMessage() string {
	consoleInputReader := bufio.NewReader(os.Stdin)
	userInput, _ := consoleInputReader.ReadString('\n')
	userInput = strings.Trim(userInput, " \n")
	return userInput
}

func getFormattedTime() string {
	currentTime := time.Now()
	timeFormat := "2006-01-02 15:04:05"
	return fmt.Sprint(currentTime.Format(timeFormat))
}

func formatString(fileLine int) string {
	if fileLine == 1 {
		return "%v: %v %v"
	}
	return "\n%v: %v %v"
}

func writeMessageToFile(file *os.File, stringCount int, datetime string, userInput string) {
	format := formatString(stringCount)
	_, err := file.WriteString(fmt.Sprintf(format, stringCount, datetime, userInput))
	if err != nil {
		panic(err)
	}
}

func module12Task1() {
	fmt.Println("Задание 1. Работа с файлами")

	// работа с файлом
	file, err := os.Create("records.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var stringCount = 0

	for {
		if stringCount == 0 {
			fmt.Println("Введите произвольное сообщение или `exit` для завершения.")
		}
		// получаем сообщение от пользователя
		userInput := getMessage()
		if userInput == "exit" {
			break
		}
		stringCount++

		// получаем время в нужном формате
		datetime := getFormattedTime()

		// запись в файл без лишних переносов строки
		writeMessageToFile(file, stringCount, datetime, userInput)
		fmt.Println("Введите ещё одну запись или `exit` для завершения.")
	}
	fmt.Println("Работа приложения завершена")

}

func module12Task2() {
	fmt.Println("Задание 2. Интерфейс io.Reader")

	// проверить наличие:
	f, err := os.Open("records.txt")
	if err != nil {
		fmt.Println("Файл не найден")
		return
	}
	defer f.Close()

	// проверить размер:
	fileInfo, err := f.Stat()
	if err != nil || fileInfo.Size() == 0 {
		fmt.Println("Файл пуст")
		return
	}

	// считать в буфер по размеру файла:
	buf := make([]byte, fileInfo.Size())
	if _, err := io.ReadFull(f, buf); err != nil {
		panic(err)
	}
	fmt.Println("Содержимое файла", fileInfo.Name())
	fmt.Println(string(buf))
}

func module12Task3() {
	fmt.Println("Задание 3. Уровни доступа")
	// работа с файлом
	file, err := os.Create("readonly.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
	defer file.Close()

	_, err = file.WriteString("Единственная строка в файле")
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}

	fileInfo, err := file.Stat()
	fmt.Println(fileInfo.Mode())

	err = file.Chmod(0444)
	if err != nil {
		fmt.Println("Ошибка при смене прав доступа файла:", err)
	}

	fileInfo, err = file.Stat()
	fmt.Println(fileInfo.Mode())
}

func module12Task4() {
	var b bytes.Buffer
	filename := "records.txt"
	var stringCount = 0

	for {
		if stringCount == 0 {
			fmt.Println("Введите произвольное сообщение или `exit` для завершения.")
		}
		// получаем сообщение от пользователя
		userInput := getMessage()
		if userInput == "exit" {
			break
		}
		stringCount++

		// получаем время в нужном формате
		datetime := getFormattedTime()

		// запись в буфер
		format := formatString(stringCount)
		_, err := b.WriteString(fmt.Sprintf(format, stringCount, datetime, userInput))
		if err != nil {
			panic(err)
		}
		fmt.Println("Введите ещё одну запись или `exit` для завершения.")
	}
	if err := ioutil.WriteFile(filename, b.Bytes(), 0666); err != nil {
		panic(err)
	}

	// читаем файл
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Файл:")
	fmt.Println(string(resultBytes))
}

func generate(openBracketNum, closeBracketNum uint8, currentRecord string, result *[]string) {
	if openBracketNum == 0 {
		*result = append(*result, currentRecord+strings.Repeat(")", int(closeBracketNum)))
		return
	}
	if closeBracketNum > openBracketNum {
		generate(openBracketNum, closeBracketNum-1, currentRecord+")", result)
	}
	generate(openBracketNum-1, closeBracketNum, currentRecord+"(", result)
}

func module12Task5() {
	/* Скопировано с решения в модуле 13	*/
	result := make([]string, 0)

	var bracketPairsNumber uint8
	fmt.Print("Введите число пар скобок: ")
	_, _ = fmt.Scan(&bracketPairsNumber)

	generate(bracketPairsNumber, bracketPairsNumber, "", &result)
	fmt.Printf("Возможные варианты положения скобок: %v\n", result)
}

func runAllTasksModule12() {
	wrapper(module12Task1)
	wrapper(module12Task2)
	wrapper(module12Task3)
	wrapper(module12Task4)
	wrapper(module12Task5)
}
