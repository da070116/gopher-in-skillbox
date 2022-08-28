package internal

import (
	"fmt"
	"gopher-in-skillbox/pkg"
)

func m14Task1() {
	fmt.Println("Задание 1. Функция, возвращающая результат")
	fmt.Print("Введите максимальное значение для случайного числа: ")
	var maxRand int
	_, _ = fmt.Scan(&maxRand)
	rand := pkg.CustomRandInt(maxRand)
	fmt.Printf("Сгенерировано число %v, результат isEven(%v) = %v\n", rand, rand, isEven(rand))

}

func generateRandCoordinate() (int, int) {
	v1, v2 := pkg.CustomRandInt(100), pkg.CustomRandInt(100)
	if pkg.CustomRandInt(100) < 33 {
		v1 = -v1
	}
	if pkg.CustomRandInt(100) < 15 {
		v2 = -v2
	}
	return v1, v2
}

func modifyCoordinate(x, y *int) (int, int) {
	a, b := *x, *y

	a = 2*a + 10
	b = -3*b - 5

	return a, b
}

func m14Task2() {
	fmt.Println("Задание 2. Функция, возвращающая несколько значений")
	for i := 0; i < 3; i++ {
		x, y := generateRandCoordinate()
		a, b := modifyCoordinate(&x, &y)
		fmt.Printf("%v. Создана точка (%v, %v). Преобразование по формулам x=2×x+10, y=−3×y−5 даёт значения (%v, %v) \n", i+1, x, y, a, b)
	}
}

func namedSum(v int) (res int) {
	res = v + pkg.CustomRandInt(20)
	return
}

func namedMult(v int) (res int) {
	res = v * pkg.CustomRandInt(10)
	return
}

func m14Task3() {
	fmt.Println("Задание 3. Именованные возвращаемые значения")
	digit := pkg.CustomRandInt(100)
	fmt.Println("digit =", digit)
	digit = namedSum(digit)
	fmt.Println("digit =", digit)
	digit = namedMult(digit)
	fmt.Println("digit =", digit)

}

var var1 = pkg.CustomRandInt(10)
var var2 = pkg.CustomRandInt(30)

const var3 = -42

func sumNoFirst(v int) int {
	return v + var2 + var3
}

func sumNoSecond(v int) int {
	return v + var3 + var1
}

func sumNoThird(v int) int {
	return v + var2 + var1
}

func m14Task4() {
	fmt.Println("Задание 4. Область видимости переменных")
	digit := pkg.CustomRandInt(10)
	fmt.Println(digit, var1, var2, var3)
	fmt.Println(sumNoFirst(digit))
	fmt.Println(sumNoSecond(digit))
	fmt.Println(sumNoThird(digit))
}

func isEven(val int) bool {
	return val%2 == 0
}

func m14RunAll() {
	fmt.Println("Задания для модуля 14")
	pkg.Wrapper(m14Task1)
	pkg.Wrapper(m14Task2)
	pkg.Wrapper(m14Task3)
	pkg.Wrapper(m14Task4)
}
