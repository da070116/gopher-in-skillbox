package main

import (
	"fmt"
	"math"
)

func taylorSeriesForExponent() {
	fmt.Println("Задание 1. Разложение ex в ряд Тейлора")
	/*
		E^x = Sum(x^n/n!)
	*/
	var x float64
	var myPrecious uint8
	fmt.Print("Введите точность (число символов после 0): ")
	_, _ = fmt.Scan(&myPrecious)

	fmt.Print("Введите значение переменной х: ")
	_, _ = fmt.Scan(&x)

	epsilon := 1 / math.Pow(10, float64(myPrecious))

	funcResult := float64(0)
	funcPrevResult := float64(0)
	factorial := 1
	n := 0
	for {
		if n > 0 {
			factorial *= n
		}
		funcResult += math.Pow(x, float64(n)) / float64(factorial)
		if math.Abs(funcResult-funcPrevResult) < epsilon {
			fmt.Println(funcResult)
			break
		}
		n++
		funcPrevResult = funcResult

	}

	fmt.Printf("Чтобы добиться требуемой точности %v понадобилось %v итераций\n", epsilon, n+1)
	fmt.Println("=========================================")
}

func percentRoundingTroubles() {
	var clientBaseSum float64
	var yearsToKeep uint8
	var monthInterestRate uint8

	fmt.Print("Сумма вклада: ")
	_, _ = fmt.Scan(&clientBaseSum)
	fmt.Print("Сколько (целых) лет хранить вклад: ")
	_, _ = fmt.Scan(&yearsToKeep)
	fmt.Print("Ежемесячная процентная ставка: ")
	_, _ = fmt.Scan(&monthInterestRate)

	augmentation := float64(monthInterestRate) / float64(100)

	resultSum := clientBaseSum

	bankInterest := float64(0)

	for i := 0; i < int(yearsToKeep)*12; i++ {

		beforeTrunc := resultSum + resultSum*augmentation
		truncatedSum := math.Trunc(beforeTrunc*100) / 100
		bankInterest += beforeTrunc - truncatedSum
		resultSum = truncatedSum
	}
	fmt.Println("Итоговая сумма на счете:", resultSum, "Выгода банка за счёт округлений: ", bankInterest)
}

func main() {
	taylorSeriesForExponent()
	percentRoundingTroubles()
}
