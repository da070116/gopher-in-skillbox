package main

import (
	"fmt"
	"math"
)

func module5Task1() {
	fmt.Println("Задание 1. Определение координатной плоскости точки")
	var x, y int
	var result string
	fmt.Println("Введите координату X:")
	_, _ = fmt.Scan(&x)
	fmt.Println("Введите координату Y:")
	_, _ = fmt.Scan(&y)
	if (x == y) && (y == 0) {
		fmt.Printf("Искомая точка (%v, %v) не принадлежит ни одной координатной четверти. Это начало координат\n", x, y)
	} else {
		if x > 0 {
			if y > 0 {
				result = "I"
			} else {
				result = "IV"
			}
		} else {
			if y > 0 {
				result = "II"
			} else {
				result = "III"
			}
		}
		fmt.Printf("Искомая точка (%v, %v) принадлежит %v координатной четверти\n", x, y, result)
	}
}
func module5Task2() {
	fmt.Println("Задание 2. Проверить, что одно из чисел — положительное")

	var d1, d2, d3 int
	fmt.Println("Введите число #1:")
	_, _ = fmt.Scan(&d1)
	fmt.Println("Введите число #2:")
	_, _ = fmt.Scan(&d2)
	fmt.Println("Введите число #3:")
	_, _ = fmt.Scan(&d3)

	if d1 > 0 || d2 > 0 || d3 > 0 {
		fmt.Println("Одно из чисел - положительное")
	} else {
		fmt.Println("Ни одно из чисел не является положительным")
	}
}
func module5Task3() {

	fmt.Println("Задание 3. Проверить, что есть совпадающие числа")
	var digit1, digit2, digit3 int
	fmt.Println("Введите число #1:")
	_, _ = fmt.Scan(&digit1)
	fmt.Println("Введите число #2:")
	_, _ = fmt.Scan(&digit2)
	fmt.Println("Введите число #3:")
	_, _ = fmt.Scan(&digit3)

	if digit1 == digit2 || digit1 == digit3 || digit2 == digit3 {
		fmt.Println("Есть совпадающие числа")
	} else {
		fmt.Println("Совпадающих чисел нет")
	}
}
func module5Task4() {
	fmt.Println("Задание 4. Сумма без сдачи")

	var sum, nom1, nom2, nom3 int
	fmt.Println("Введите номинал первой монеты:")
	_, _ = fmt.Scan(&nom1)
	fmt.Println("Введите номинал второй монеты:")
	_, _ = fmt.Scan(&nom2)
	fmt.Println("Введите номинал третьей монеты:")
	_, _ = fmt.Scan(&nom3)
	fmt.Println("Введите сумму к оплате:")
	_, _ = fmt.Scan(&sum)
	canBeFlag := " не "
	sumIsEqualToOneCoin := (sum == nom1) || (sum == nom2) || (sum == nom3)
	sumIsEqualToTwoCoins := (sum == nom1+nom2) || (sum == nom1+nom3) || (sum == nom3+nom2)
	sumIsEqualToAllCoins := sum == nom1+nom2+nom3
	if sumIsEqualToAllCoins || sumIsEqualToTwoCoins || sumIsEqualToOneCoin {
		canBeFlag = " "
	}
	fmt.Printf("Сумма %v %vможет быть оплачена без сдачи монетами следующих номиналов: %v, %v, %v\n", sum, canBeFlag, nom1, nom2, nom3)

}
func module5Task5() {
	fmt.Println("Задание 5. Решение квадратного уравнения a*x^2 + b*x + c = 0")

	var a, b, c, discriminant float64
	fmt.Println("Введите коэффициент a:")
	_, _ = fmt.Scan(&a)
	fmt.Println("Введите коэффициент b:")
	_, _ = fmt.Scan(&b)
	fmt.Println("Введите коэффициент c:")
	_, _ = fmt.Scan(&c)

	discriminant = math.Pow(b, 2.0) - 4*a*c

	if discriminant < 0.0 {
		fmt.Println("Действительных корней нет")
	} else if discriminant == 0 {
		x := -b / (2 * a)
		fmt.Println("Корень уравнения равен", x)
	} else {
		x1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b + math.Sqrt(discriminant)) / (2 * a)
		fmt.Println("Корни уравнения равны", x1, x2)
	}
}

func module5Task6() {
	fmt.Println("Задание 6. Счастливый билет")
	var ticket int
	fmt.Println("Введите номер билета:")
	_, _ = fmt.Scan(&ticket)

	t1 := ticket / 1000
	t2 := ticket / 100 % 10
	t3 := ticket / 10 % 10
	t4 := ticket % 10
	if t1 == t4 && t2 == t3 {
		fmt.Println("зеркальный билет")
	} else if t1+t2 == t3+t4 {
		fmt.Println("счастливый билет")
	} else {
		fmt.Println("обычный билет")
	}
}

func module5Task7() {
	fmt.Println("Задание 7 (по желанию). Игра «Угадай число»")
	var possibleAnswer, triesToGuess int
	var userAnswer string
	var isGuessed bool
	min := 1
	max := 10

	for {
		triesToGuess += 1
		if triesToGuess > 4 {
			fmt.Println("Число не угадано. Превышено число попыток")
			break
		}
		possibleAnswer = (max + min) / 2
		fmt.Printf("Попытка %v. Это %v? Введите `=`, если ответ верен, `>` если Ваше число больше, и `<` если меньше\n", triesToGuess, possibleAnswer)
		_, _ = fmt.Scan(&userAnswer)

		if userAnswer == "=" {
			isGuessed = true
		} else if userAnswer == ">" {
			min = possibleAnswer + 1
		} else {
			max = possibleAnswer - 1
		}
		if isGuessed {
			fmt.Println("Число угадано. Попыток:", triesToGuess)
			break
		}
	}
}

func runAllTasksModule5() {
	fmt.Println("Это задания 2 практической работы в курсе Skillbox.")
	wrapper(module5Task1)
	wrapper(module5Task2)
	wrapper(module5Task3)
	wrapper(module5Task4)
	wrapper(module5Task5)
	wrapper(module5Task6)
	wrapper(module5Task7)
}
