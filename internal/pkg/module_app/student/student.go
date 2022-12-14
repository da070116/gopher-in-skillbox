package student

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	name  string
	age   int
	grade int
}

// NewStudent  - создание новой сущности типа Person из необработанной строки
func NewStudent(rawData string) (err error, st Person) {
	items := strings.Split(rawData, " ")

	st.name = items[0]
	st.age = strToInt(items[1])
	st.grade = strToInt(items[2])
	if st.age == -1 || st.grade == -1 {
		err = errors.New("bad data")
	}
	return
}

// DisplayStudentInfo - вывод данных о структуре Person
func DisplayStudentInfo(s *Person) (info string) {
	info = fmt.Sprintf("Person %v (age: %d) of grade %d\n", s.name, s.age, s.grade)
	return
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
