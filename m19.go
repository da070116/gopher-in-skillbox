package main

import "fmt"

func sortLen5(raw [5]int) (sorted [5]int) {
	for i := 0; i < len(raw)-1; i++ {

		isSwitched := false
		for j := 0; j < len(raw)-1-i; j++ {
			if raw[j] > raw[j+1] {
				raw[j], raw[j+1] = raw[j+1], raw[j]
				isSwitched = true
			}
		}
		if isSwitched == false {
			break
		}
	}
	sorted = raw
	return
}

func sortLen4(raw [4]int) (sorted [4]int) {
	for i := 0; i < len(raw)-1; i++ {

		isSwitched := false
		for j := 0; j < len(raw)-1-i; j++ {
			if raw[j] > raw[j+1] {
				raw[j], raw[j+1] = raw[j+1], raw[j]
				isSwitched = true
			}
		}
		if isSwitched == false {
			break
		}
	}
	sorted = raw
	return
}

func sumArr(a1 [4]int, a2 [5]int) (res [9]int) {
	for i := 0; i <= 8; i++ {
		if i > 3 {
			res[i] = a2[i-3-1]
		} else {
			res[i] = a1[i]
		}
	}
	return
}

func m19RunAll() {
	var arrayOf5 [5]int
	var arrayOf4 [4]int
	var arrayOf9 [9]int

	for i := 0; i <= len(arrayOf5)-1; i++ {
		fmt.Println("Введите элемент массива:")
		_, _ = fmt.Scan(&arrayOf5[i])
	}
	fmt.Println(arrayOf5)
	for i := 0; i <= len(arrayOf4)-1; i++ {
		fmt.Println("Введите элемент массива:")
		_, _ = fmt.Scan(&arrayOf4[i])
	}
	fmt.Println(arrayOf4)
	sorted5 := sortLen5(arrayOf5)
	sorted4 := sortLen4(arrayOf4)
	arrayOf9 = sumArr(sorted4, sorted5)
	fmt.Println(arrayOf9)
}
