package main

import (
	"fmt"
	"strconv"
)

// main - выполнение основной логики программы
func main() {
	for {
		chanStr := make(chan string)
		go writer(chanStr)
		val := <-chanStr
		if val == "exit" {
			break
		} else {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				fmt.Print(fmt.Errorf("Error: %v\n", err))
				break
			}
			chanInt := make(chan int)
			chanInt <- intVal
			sc := power(chanInt)
			tc := multiply(sc)
			fmt.Println(<-tc)
		}
	}
	fmt.Println("Terminated by user")
}

func writer(ch chan string) {
	var userInput string
	fmt.Print("Enter a digit: ")
	_, _ = fmt.Scan(&userInput)
	ch <- userInput
	close(ch)
}

func power(firstChan chan int) chan int {
	secondChan := make(chan int)
	val := <-firstChan
	go func() {
		secondChan <- val * val
	}()
	return secondChan
}

func multiply(secondChan chan int) chan int {
	thirdChan := make(chan int)
	val := <-secondChan
	go func() {
		thirdChan <- val * 2
	}()
	return thirdChan
}
