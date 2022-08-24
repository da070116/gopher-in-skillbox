package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

// strToInt - преобраозовать строку в число
func strToInt(value string) (result int) {
	return
}

// newStudent - создание новой сущности типа Student из необработанной строки
func newStudent(rawData string) (st Student) {
	items := strings.Split(rawData, " ")
	st.name = items[0]
	_age, err := strconv.Atoi(items[1])
	if err != nil {
		fmt.Println("Ошибка:", err)
		_age = 0
	}
	st.age = _age
	_grade, err := strconv.Atoi(items[2])
	if err != nil {
		fmt.Println("Ошибка:", err)
		_grade = 0
	}
	st.age = _grade
	return
}

func (s Student) studentInfo() (info string) {
	info = fmt.Sprintf("Student %v (age: %d) of grade %d\n", s.name, s.age, s.grade)
	return
}

// main - выполнение основной логики программы
func main() {
	for {

		inputSource := bufio.NewReader(os.Stdin)
		rawInput, err := inputSource.ReadString('\n')
		if err == io.EOF {
			break
		}
		rawInput = strings.Trim(rawInput, "\n")
		fmt.Println("Go down")
	}
}
