package main

import "fmt"

func m6Task1() {
	fmt.Println("Задание 1. Написание простого цикла")
	var customDigit int
	fmt.Println("Введите произвольное целое число:")
	_, _ = fmt.Scan(&customDigit)

	if customDigit == 0 {
		fmt.Println("Вы ввели 0, цикл невозможен")
	} else {
		fmt.Println("Вывод цикла от 0 до", customDigit)
		if customDigit > 0 {
			for i := 0; i <= customDigit; i++ {
				fmt.Println(i)
			}
		} else {
			for i := 0; i >= customDigit; i-- {
				fmt.Println(i)
			}
		}
	}
}

func m6Task2() {

	fmt.Println("Задание 2. Сумма двух чисел по единице")
	var sumDigit1, sumDigit2 int
	fmt.Println("Введите первое целое число:")
	_, _ = fmt.Scan(&sumDigit1)
	fmt.Println("Введите второе целое число:")
	_, _ = fmt.Scan(&sumDigit2)

	fmt.Println("Вывод суммы", sumDigit1, "и", sumDigit2, "с шагом 1:")
	sum := sumDigit1 + sumDigit2
	if sum < sumDigit1 {
		for i := sumDigit1; i >= sum; i-- {
			fmt.Println(i)
		}
	} else {
		for i := sumDigit1; i <= sum; i++ {
			fmt.Println(i)
		}
	}
	fmt.Println("Cумма чисел", sumDigit1, "и", sumDigit2, "равна", sum)
}
func m6Task3() {
	fmt.Println("Задание 3. Расчёт суммы скидки")
	var price, discountPercent float32
	fmt.Println("Введите стоимость товара:")
	_, _ = fmt.Scan(&price)
	fmt.Println("Введите размер скидки в %:")
	_, _ = fmt.Scan(&discountPercent)

	if discountPercent > 30 {
		fmt.Println("Выбранный размер скидки превышает допустимый и был снижен до 30%.")
		discountPercent = 30
	}

	discountValue := price * (discountPercent / 100)
	if discountValue > 2000 {
		fmt.Println("Максимально возможный размер скидки - 2000 рублей")
		discountValue = 2000
		discountPercent = (discountValue / price) * 100
	}

	fmt.Println("Стоимость товара:", price)
	fmt.Print("Скидка (", int(discountPercent), "%): ", int(discountValue), " рублей\n")
	fmt.Println("Сумма к оплате:", price-discountValue)
}
func m6Task4() {
	fmt.Println("Задание 4. Предыдущее занятие на if")
	var variableMaxAt10, variableMaxAt100, variableMaxAt1000 int
	for i := 0; i <= 1000; i++ {
		if variableMaxAt10 < 10 {
			variableMaxAt10++
		}
		if variableMaxAt100 < 100 {
			variableMaxAt100++
		}
		variableMaxAt1000++
	}
}
func m6Task5() {
	fmt.Println("Задание 5 (по желанию). Задача на постепенное наполнение корзин разной ёмкости")
	var firstBasketCapacity, secondBasketCapacity, thirdBasketCapacity uint
	fmt.Println("Введите вместимость первой корзины:")
	_, _ = fmt.Scan(&firstBasketCapacity)

	fmt.Println("Введите вместимость второй корзины:")
	_, _ = fmt.Scan(&secondBasketCapacity)

	fmt.Println("Введите вместимость третьей корзины:")
	_, _ = fmt.Scan(&thirdBasketCapacity)

	var applesInFirst uint = 0
	var applesInSecond uint = 0
	var applesInThird uint = 0

	iterations := 0

	for applesInThird <= thirdBasketCapacity {
		iterations++
		fmt.Print("Шаг ", iterations, ":\nВ первой корзине ", applesInFirst, " из ", firstBasketCapacity, ". Во второй корзине ", applesInSecond, " из ", secondBasketCapacity, " В третьей корзине ", applesInThird, " из ", thirdBasketCapacity, "\n--------------------\n")
		if applesInFirst < firstBasketCapacity {
			applesInFirst++
			continue
		}
		if applesInSecond < secondBasketCapacity {
			applesInSecond++
			continue
		}
		applesInThird++
	}
	fmt.Println("Корзины заполнены за", iterations, "шагов цикла")
}

func m6Task6() {
	const topFloor int = 24
	const baseFloor int = 1
	const passengerInElevator int = -1

	var elevatorOnFloor int = 5 // Cтартуем с этого этажа

	passenger1 := 4
	passenger2 := 7
	passenger3 := 10

	currentCapacity := 2

	fmt.Println("Лифт начинает движение с", elevatorOnFloor, "этажа.")
	for {
		// логика поведения на 1-м этаже
		if elevatorOnFloor == baseFloor {
			if passenger1 == passengerInElevator {
				passenger1 = 1
				currentCapacity++
				fmt.Println("Пассажир №1 прибыл на 1-й этаж.")
			}
			if passenger2 == passengerInElevator {
				passenger2 = 1
				currentCapacity++
				fmt.Println("Пассажир №2 прибыл на 1-й этаж.")
			}
			if passenger3 == passengerInElevator {
				passenger3 = 1
				currentCapacity++
				fmt.Println("Пассажир №3 прибыл на 1-й этаж.")
			}
			if passenger1 == 1 && passenger2 == 1 && passenger3 == 1 {
				fmt.Println("Все пассажиры прибыли на 1-й этаж.")
				break
			}
			fmt.Println("Лифт на", baseFloor, "этаже. Двигается вверх без остановок.")
			elevatorOnFloor = topFloor
			fmt.Println("Лифт на", elevatorOnFloor, "этаже. Двигается вниз.")
		}
		// движение вниз
		for elevatorOnFloor > baseFloor {
			if currentCapacity > 0 && (passenger1 == elevatorOnFloor || passenger2 == elevatorOnFloor || passenger3 == elevatorOnFloor) {
				fmt.Println("Лифт двигался вниз и остановился на", elevatorOnFloor, "этаже.")
			}

			if passenger1 == elevatorOnFloor && currentCapacity > 0 {
				passenger1 = passengerInElevator
				currentCapacity--
				fmt.Println("Пассажир №1 входит в лифт.")
			}

			if passenger2 == elevatorOnFloor && currentCapacity > 0 {
				passenger2 = passengerInElevator
				currentCapacity--
				fmt.Println("Пассажир №2 входит в лифт.")
			}

			if passenger3 == elevatorOnFloor && currentCapacity > 0 {
				passenger3 = passengerInElevator
				currentCapacity--
				fmt.Println("Пассажир №3 входит в лифт.")
			}
			elevatorOnFloor--
		}
	}
	fmt.Println("===================================")
}

func m6RunAll() {
	fmt.Println("Это задания 3 практической работы в курсе Skillbox.")
	m6Task1()
	m6Task2()
	m6Task3()
	m6Task4()
	m6Task5()
	m6Task6()
}
