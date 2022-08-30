package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Conveyor() {

	for {
		fc := input()
		a := <-fc
		if a == "exit" {
			fmt.Println("Завершение программы по команде пользователя")
			close(fc)
			break
		}
		digit, err := strconv.Atoi(a)
		if err != nil {
			fmt.Print(fmt.Errorf("Ошибка при обработке данных: %v\n", err))
			break
		}
		intChannel := make(chan int)
		intChannel <- digit

		sc := square(intChannel)
		tc := multiply(sc)
		fmt.Println(<-tc)
	}

}

func input() chan string {
	c1 := make(chan string, 3)
	go func() {
		fmt.Print("input a value: ")
		consoleInputReader := bufio.NewReader(os.Stdin)
		result, _ := consoleInputReader.ReadString('\n')
		result = strings.Trim(result, " \n")
		c1 <- result
	}()
	return c1
}

func square(ic chan int) chan int {
	c2 := make(chan int)
	go func() {
		value := <-ic
		c2 <- value * value
	}()
	return c2
}

func multiply(ic chan int) chan int {
	c3 := make(chan int)
	go func() {
		value := <-ic
		fmt.Println(value)
		c3 <- value * 2
	}()
	return c3
}
