package internal

import "fmt"

func countSum(from, to int) {
	sumEvens := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			sumEvens += i
		}
	}
	fmt.Printf("Сумма чётных чисел в интервале от %v до %v составит %v\n", from, to, sumEvens)
}

func swap(from, to *int) {
	fmt.Println("Вы ввели:", *from, *to)
	// swap if from > to
	if *from > *to {
		*to, *from = *from, *to
		fmt.Println("Первое число больше второго. Меняем местами.")
		fmt.Println("Новые значения:", *from, *to)
	}

}

func m13RunAll() {
	var digit1, digit2 int
	fmt.Print("Введите через пробел два целых числа: ")
	_, _ = fmt.Scan(&digit1, &digit2)
	swap(&digit1, &digit2)
	countSum(digit1, digit2)
}
