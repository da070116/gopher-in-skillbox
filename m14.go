package main

import "fmt"

func m14Task1() {
	fmt.Println("Задание 1. Функция, возвращающая результат")
	fmt.Print("Введите максимальное значение для случайного числа: ")
	var maxRand int
	_, _ = fmt.Scan(&maxRand)
	rand := customRandInt(maxRand)
	fmt.Printf("Сгенерировано число %v, результат isEven(%v) = %v\n", rand, rand, isEven(rand))

}

func generateRandCoordinate() (int, int) {
	v1, v2 := customRandInt(100), customRandInt(100)
	if customRandInt(100) < 33 {
		v1 = -v1
	}
	if customRandInt(100) < 15 {
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
	res = v + customRandInt(20)
	return
}

func namedMult(v int) (res int) {
	res = v * customRandInt(10)
	return
}

func m14Task3() {
	fmt.Println("Задание 3. Именованные возвращаемые значения")
	digit := customRandInt(100)
	fmt.Println("digit =", digit)
	digit = namedSum(digit)
	fmt.Println("digit =", digit)
	digit = namedMult(digit)
	fmt.Println("digit =", digit)

}

var task4GlobalVar1 = customRandInt(10)
var task4GlobalVar2 = customRandInt(30)

const task4GlobalVar3 = -42

func sumGlobalVar1(v int) int {
	return v + task4GlobalVar1
}

func sumGlobalVar2(v int) int {
	return v + task4GlobalVar2
}

func sumGlobalVars(v int) int {
	return v + task4GlobalVar2 + task4GlobalVar3 + task4GlobalVar1
}

func m14Task4() {
	fmt.Println("Задание 4. Область видимости переменных")
	digit := customRandInt(10)
	fmt.Println(digit, task4GlobalVar1, task4GlobalVar2, task4GlobalVar3)
	digit = sumGlobalVar1(digit)
	fmt.Println(sumGlobalVars(sumGlobalVar2(digit)))

}

func isEven(val int) bool {
	return val%2 == 0
}

func m14RunAll() {
	fmt.Println("Задания для модуля 14")
	wrapper(m14Task1)
	wrapper(m14Task2)
	wrapper(m14Task3)
	wrapper(m14Task4)
}
