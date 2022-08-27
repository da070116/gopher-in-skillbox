package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

// strToInt - преобразовать строку в число
func strToInt(value string) (result int) {
	value = strings.Trim(value, " ")
	result, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Ошибка:", err)
		result = -1
	}
	return
}

// newStudent - создание новой сущности типа Student из необработанной строки
func newStudent(rawData string) (err any, st Student) {
	items := strings.Split(rawData, " ")

	st.name = items[0]
	st.age = strToInt(items[1])
	st.grade = strToInt(items[2])
	if st.age == -1 || st.grade == -1 {
		err = -1
	}
	return
}

// studentInfo - вывод данных о структуре Student
func (s Student) studentInfo() (info string) {
	info = fmt.Sprintf("Student %v (age: %d) of grade %d\n", s.name, s.age, s.grade)
	return
}

// main - выполнение основной логики программы
func main() {
	const EOT = 4
	counter := 0
	studentsStorage := make(map[int]Student)

	for {
		inputSource := bufio.NewReader(os.Stdin)
		fmt.Print("Введите данные о студенте в формате `имя` `возраст` `курс`: ")
		rawInput, _ := inputSource.ReadString('\n')
		rawInput = strings.Trim(rawInput, "\r\n")
		if rune(rawInput[0]) == EOT {
			fmt.Println("Пользователь запросил выход")
			break
		}

		err, student := newStudent(rawInput)
		if err != nil {
			fmt.Println("Ошибка при вводе данных")
		} else {
			studentsStorage[counter] = student
			fmt.Println("Запись внесена")
		}
		counter++
	}
	fmt.Println("Хранилище студентов содержит записи:")
	for i := range studentsStorage {
		fmt.Print(studentsStorage[i].studentInfo())
	}

}
