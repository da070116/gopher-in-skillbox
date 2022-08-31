package main

import (
	"errors"
	"fmt"
	"strconv"
)

// main - выполнение основной логики программы
func main() {
	for {
		chanString := make(chan string)
		go writer(chanString)
		userInput := <-chanString
		if userInput == "exit" {
			return
		} else {
			digit, err := strconv.Atoi(userInput)
			if err != nil {
				panic(errors.New("can't convert value"))
			}
			firstIntegerChannel := power(digit)
			secondIntegerChannel := multiply(firstIntegerChannel)
			fmt.Println(<-secondIntegerChannel)
		}
	}
}

func writer(ch chan string) {
	var userInput string
	fmt.Print("Enter a digit: ")
	_, _ = fmt.Scan(&userInput)
	ch <- userInput
	close(ch)
}

func power(val int) chan int {
	intChan := make(chan int)
	go func() {
		fmt.Printf("value %v powered by 2 is %v\n", val, val*val)
		intChan <- val * val
		close(intChan)
	}()
	return intChan
}

func multiply(intChan chan int) chan int {
	thirdChan := make(chan int)
	val := <-intChan
	go func() {
		fmt.Printf("value %v multiplied by 2 is %v\n", val, val*2)
		thirdChan <- val * 2
	}()
	return thirdChan
}
