package internal

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func m8Task1() {
	fmt.Println("Задание 1. Времена года")
	var month, season string
	fmt.Println("Введите месяц")
	_, _ = fmt.Scan(&month)
	switch strings.ToLower(month) {
	case "декабрь", "январь", "февраль":
		season = "зима"
	case "март", "апрель", "май":
		season = "весна"
	case "июнь", "июль", "август":
		season = "лето"
	case "сентябрь", "октябрь", "ноябрь":
		season = "зима"
	default:
		fmt.Println("Неизвестный месяц")
	}
	if season != "" {
		fmt.Printf("Введённый вами месяц %v приходится на сезон: %v\n", month, season)
	}
	fmt.Println("========================================================================")
}

func m8Task2() {
	fmt.Println("Задание 2. Дни недели")
	var userInput string
	fmt.Print("Введите будний день в сокращённой форме: ")
	_, _ = fmt.Scan(&userInput)
	switch strings.ToLower(userInput) {
	case "пн":
		fmt.Println("понедельник")
		fallthrough
	case "вт":
		fmt.Println("вторник")
		fallthrough
	case "ср":
		fmt.Println("среда")
		fallthrough
	case "чт":
		fmt.Println("четверг")
		fallthrough
	case "пт":
		fmt.Println("пятница")
		fallthrough
	case "сб", "вск":
		fmt.Println(userInput, "- выходной")
	default:
		fmt.Println("Неизвестная форма")
	}
	fmt.Println("========================================================================")

}

func m8Task3() {
	fmt.Println("Задание 3 (по желанию). Расчёт сдачи")
	const queueSize int = 20
	rand.Seed(time.Now().UnixNano())

	banknotes := [3]int{5, 10, 20}

	fiveDollarsCount := 0
	tenDollarsCount := 0
	twentyDollarsCount := 0
	gainedTotal := 0

	// процедурная генерация очереди
	var queue [queueSize]int
	for i := 0; i < 20; i++ {
		queue[i] = banknotes[rand.Intn(3)]
	}
	fmt.Println(queue)

	i := 0
	for i = 0; i < 20; i++ {
		switch queue[i] {
		case 5:
			fiveDollarsCount++
		case 10:
			tenDollarsCount++
			fiveDollarsCount--
		case 20:
			twentyDollarsCount++
			if tenDollarsCount > 0 {
				tenDollarsCount--
				fiveDollarsCount--
			} else {
				fiveDollarsCount -= 3
			}
		}
		if fiveDollarsCount < 0 || tenDollarsCount < 0 {
			fmt.Printf("Обслужено клиентов: %v. Не хватило сдачи для $%v \n", i+1, queue[i])
			break
		} else {
			fmt.Printf("Обслужено %v клиентов. $5 - %v, $10 - %v, $20 - %v \n", i+1, fiveDollarsCount, tenDollarsCount, twentyDollarsCount)
		}
	}
	if i < queueSize {
		fmt.Println(false)
	} else {
		gainedTotal = fiveDollarsCount*5 + tenDollarsCount*10 + twentyDollarsCount*20
		fmt.Println(true, "Прибыль:", gainedTotal)
	}
}

func m8RunAll() {
	m8Task1()
	m8Task2()
	m8Task3()
}
