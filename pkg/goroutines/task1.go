package goroutines

import (
	"errors"
	"fmt"
	"strconv"
)

// Task1Conveyor - выполнение основной логики программы
func Task1Conveyor() {
	for {
		chanString := make(chan string)
		go writer(chanString)
		userInput := <-chanString
		if userInput == "exit" {
			return
		} else {
			digit, err := strconv.Atoi(userInput)
			if err != nil {
				panic(errors.New("не могу преобразовать значение в число"))
			}
			firstIntegerChannel := power(digit)
			secondIntegerChannel := multiply(firstIntegerChannel)
			fmt.Println(<-secondIntegerChannel)
		}
	}
}

// writer - function to obtain user input
func writer(ch chan string) {
	var userInput string
	fmt.Print("Введите целое число или 'exit' для выхода из программы: ")
	_, _ = fmt.Scan(&userInput)
	ch <- userInput
	close(ch)
}

// power - return value powered by 2
func power(val int) chan int {
	intChan := make(chan int)
	go func() {
		fmt.Printf("Квадрат: %v\n", val*val)
		intChan <- val * val
		close(intChan)
	}()
	return intChan
}

// multiply - return value multiplied by 2
func multiply(intChan chan int) chan int {
	thirdChan := make(chan int)
	val := <-intChan
	go func() {
		fmt.Printf("Произведение на 2: %v\n", val*2)
		thirdChan <- val * 2
		close(thirdChan)
	}()
	return thirdChan
}
