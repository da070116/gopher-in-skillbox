package main

import (
	"fmt"
	"math/rand"
	"time"
)

func module4Task1() {
	fmt.Println("Задача 1: баллы ЕГЭ.")
	var firstExamScore, secondExamScore, thirdExamScore, totalScore int
	const enterThresholdScore int = 275

	fmt.Println("Введите результат первого экзамена:")
	_, _ = fmt.Scan(&firstExamScore)
	fmt.Println("Введите результат второго экзамена:")
	_, _ = fmt.Scan(&secondExamScore)
	fmt.Println("Введите результат третьего экзамена:")
	_, _ = fmt.Scan(&thirdExamScore)

	totalScore = firstExamScore + secondExamScore + thirdExamScore

	fmt.Println("Сумма проходных баллов:", enterThresholdScore)
	fmt.Println("Количество набранных баллов:", totalScore)

	if totalScore >= enterThresholdScore {
		fmt.Println("Вы поступили.")
	} else {
		fmt.Println("Вы не поступили")
	}
}
func module4Task2() {
	fmt.Println("Задача 2: Три числа")
	var firstDigit, secondDigit, thirdDigit int
	const compareToDigit int = 5
	fmt.Println("Введите первое число:")
	_, _ = fmt.Scan(&firstDigit)
	fmt.Println("Введите второе число:")
	_, _ = fmt.Scan(&secondDigit)
	fmt.Println("Введите третье число:")
	_, _ = fmt.Scan(&thirdDigit)
	if firstDigit > compareToDigit || secondDigit > compareToDigit || thirdDigit > compareToDigit {
		fmt.Printf("Среди введённых чисел есть число больше %v.\n", compareToDigit)
	} else {
		fmt.Printf("Среди введённых чисел нет числа больше %v.\n", compareToDigit)
	}
}
func module4Task3() {
	fmt.Println("Задача 3: Банкомат")
	const minValue int = 100
	const maxValue int = 100000
	const stateError string = "Ошибка. Операция не выполнена:"
	const stateOK string = "Операция успешно выполнена."

	var requiredValue int
	fmt.Println("Введите сумму снятия со счёта:")
	_, _ = fmt.Scan(&requiredValue)
	if requiredValue < minValue {
		fmt.Println(stateError, "сумма снятия меньше минимального номинала в", minValue, "рублей")
	} else if requiredValue%minValue > 0 {
		fmt.Println(stateError, "часть суммы", requiredValue%minValue, "рублей меньше минимального номинала в", minValue, "рублей")
	} else if requiredValue > maxValue {
		fmt.Println(stateError, "сумма снятия больше максимально допустимой суммы в", maxValue, "рублей")
	} else {
		fmt.Print(stateOK, "\nВы сняли ", requiredValue, " рублей.\n")
	}
}
func module4Task4() {
	fmt.Println("Задача 4: Три числа: ещё попытка")
	var _firstDigit, _secondDigit, _thirdDigit, counter int
	fmt.Println("Введите первое число:")
	_, _ = fmt.Scan(&_firstDigit)
	fmt.Println("Введите второе число:")
	_, _ = fmt.Scan(&_secondDigit)
	fmt.Println("Введите третье число:")
	_, _ = fmt.Scan(&_thirdDigit)
	if _firstDigit >= 5 {
		counter += 1
	}
	if _secondDigit >= 5 {
		counter += 1
	}
	if _thirdDigit >= 5 {
		counter += 1
	}
	fmt.Println("Среди введённых чисел", counter, "больше или равны 5.")
}
func module4Task5() {
	var dayOfWeek, visitorsGroupSize, billSum, sumToPay int
	fmt.Println("Введите день недели:")
	_, _ = fmt.Scan(&dayOfWeek)
	fmt.Println("Введите число гостей:")
	_, _ = fmt.Scan(&visitorsGroupSize)
	fmt.Println("Введите сумму чека:")
	_, _ = fmt.Scan(&billSum)

	if (dayOfWeek > 7 || dayOfWeek < 1) || visitorsGroupSize < 1 || billSum <= 0 {
		fmt.Println("Ошибка при введении исходных данных")
	} else {
		sumToPay = billSum
		if dayOfWeek == 1 {
			fmt.Println("Скидка по понедельникам:", billSum/10)
			sumToPay -= billSum / 10

		} else if dayOfWeek == 5 && billSum > 10000 {
			fmt.Println("Скидка по пятницам:", billSum*5/100)
			sumToPay -= billSum * 5 / 100
		}
		if visitorsGroupSize > 5 {
			fmt.Println("Надбавка на обслуживание:", billSum/10)
			sumToPay += billSum / 10
		}
		fmt.Println("Сумма к оплате:", sumToPay)
	}
}

func module4Task6() {
	fmt.Println("Задача 5: Студенты")
	var studentsAmount, groupsAmount, studentNumber int
	fmt.Println("Введите количество студентов:")
	_, _ = fmt.Scan(&studentsAmount)
	fmt.Println("Введите количество групп:")
	_, _ = fmt.Scan(&groupsAmount)
	fmt.Println("Введите порядковый номер студента:")
	_, _ = fmt.Scan(&studentNumber)

	if studentsAmount > 0 && groupsAmount > 1 && (studentNumber > 0 && studentNumber <= studentsAmount) {
		rand.Seed(time.Now().Unix())

		pickedGroup := rand.Intn(groupsAmount-1) + 1

		fmt.Println("Студент с номером", studentNumber, "распределён в группу", pickedGroup)

	} else {
		fmt.Println("Ошибка при введении исходных данных")
	}
}

func runAllTasksModule4() {
	fmt.Println("Это задания 1 практической работы в курсе Skillbox.")
	wrapper(module4Task1)
	wrapper(module4Task2)
	wrapper(module4Task3)
	wrapper(module4Task4)
	wrapper(module4Task5)
	wrapper(module4Task6)
}
