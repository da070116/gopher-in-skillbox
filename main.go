package main

import (
	"bufio"
	"fmt"
	"gopher-in-skillbox/module_app/storage"
	"gopher-in-skillbox/module_app/student"
	"os"
	"strings"
)

// main - выполнение основной логики программы
func main() {
	const EOT = 4
	counter := 0
	studentsStorage := storage.NewStudentStorage()

	for {
		inputSource := bufio.NewReader(os.Stdin)
		fmt.Print("Введите данные о студенте в формате `имя` `возраст` `курс`: ")
		rawInput, _ := inputSource.ReadString('\n')
		rawInput = strings.Trim(rawInput, "\r\n")
		if rune(rawInput[0]) == EOT {
			fmt.Println("Пользователь запросил выход")
			break
		}

		err, st := student.NewStudent(rawInput)
		if err != nil {
			fmt.Println("Ошибка при вводе данных")
		} else {
			studentsStorage.Put(counter, &st)
			fmt.Println("Запись внесена")
		}
		counter++
	}
	fmt.Println("Хранилище студентов содержит записи:")
	for i := range studentsStorage {
		st, err := studentsStorage.Get(i)
		if err != nil {
			fmt.Print(fmt.Errorf("Ошибка при выводе: %v\n", err))
		} else {
			fmt.Println(st.StudentInfo())
		}
	}

}
