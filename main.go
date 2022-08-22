package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

// newStudent - создание новой сущности типа Student из необработанной строки
func newStudent(rawData string) (st Student) {
	return
}

// main - выполнение основной логики программы
func main() {
	for {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		inputSource := bufio.NewReader(os.Stdin)
		rawInput, _ := inputSource.ReadString('\n')
		rawInput = strings.Trim(rawInput, "\n")
		fmt.Printf("->%v<-\n", rawInput)

		fmt.Println("Go down")
		break
	}
}
