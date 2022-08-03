package main

import (
	"fmt"
)

func CheckTicket(ticket int) bool {

	d1 := ticket % 1000000 / 100000
	d2 := ticket % 100000 / 10000
	d3 := ticket % 10000 / 1000
	d4 := ticket % 1000 / 100
	d5 := ticket % 100 / 10
	d6 := ticket % 10
	return (d1 + d2 + d3) == (d4 + d5 + d6)
}

func module7Task1() {

	fmt.Print("Задание 1. Зеркальные билеты")
	var mirrorTickets uint = 0
	for hundredsThousands := 1; hundredsThousands < 10; hundredsThousands++ {
		for dozensThousands := 0; dozensThousands < 10; dozensThousands++ {
			for thousands := 0; thousands < 10; thousands++ {
				for hundreds := 0; hundreds < 10; hundreds++ {
					for dozens := 0; dozens < 10; dozens++ {
						for units := 0; units < 10; units++ {
							if hundredsThousands == units && dozensThousands == dozens && thousands == hundreds {
								mirrorTickets++
								// Возможно, игла тоже тут... © Кащей
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("Число зеркальных билетов:", mirrorTickets)
	fmt.Println("===========================================================")
}
func module7Task2() {
	fmt.Println("Задание 2. Шахматная доска")
	const whiteCell string = " "
	const blackCell string = "*"
	var boardWidth, boardHeight, vertical, horizontal, frame uint

	fmt.Println("Ширина доски: ")
	_, _ = fmt.Scan(&boardWidth)
	fmt.Println("Высота доски: ")
	_, _ = fmt.Scan(&boardHeight)

	for frame = 0; frame <= boardWidth; frame++ {
		fmt.Print("=")
	}
	fmt.Println("")
	for horizontal = 1; horizontal <= boardHeight; horizontal++ {
		fmt.Print("|")
		for vertical = 1; vertical <= boardWidth; vertical++ {
			if horizontal%2 == vertical%2 {
				fmt.Print(whiteCell)
			} else {
				fmt.Print(blackCell)
			}
		}
		fmt.Print("|")
		fmt.Println("")
	}
	for frame = 0; frame <= boardWidth; frame++ {
		fmt.Print("=")
	}
	fmt.Println("")
	fmt.Println("===========================================================")
}
func module7Task3() {

	fmt.Println("Задание 3. Вывод ёлочки")
	var canvasHeight int
	const space string = "."
	const star string = "*"
	fmt.Println("Высота изображения в строках:")
	_, _ = fmt.Scan(&canvasHeight)

	spacesHalfCount := canvasHeight - 1
	starsCount := 1

	for row := 1; row <= canvasHeight; row++ {
		for item := 0; item < spacesHalfCount; item++ {
			fmt.Print(space)
		}
		for item := 0; item < starsCount; item++ {
			fmt.Print(star)
		}
		for item := 0; item < spacesHalfCount; item++ {
			fmt.Print(space)
		}
		fmt.Println("")
		spacesHalfCount -= 1
		starsCount += 2
	}
	fmt.Println("===========================================================")
}
func module7Task4() {
	fmt.Println("Задание 4 (по желанию). Счастливые билеты")

	var maxTicketsToBuy, countTicketsToFindLucky, previousCountTicketsToFindLucky int

	currentTicket := 100000
	for currentTicket <= 999999 {
		if CheckTicket(currentTicket) {

			if countTicketsToFindLucky != previousCountTicketsToFindLucky {
				previousCountTicketsToFindLucky = countTicketsToFindLucky
				/*
					        // интересный блок с детальной информацией по интервалам

					fmt.Printf("В интервале от %v до %v нужно купить минимум %v билетов, чтобы выпал счастливый\n", currentTicket, currentTicket - countTicketsToFindLucky, countTicketsToFindLucky)
				*/
			}
			if previousCountTicketsToFindLucky > maxTicketsToBuy {
				maxTicketsToBuy = previousCountTicketsToFindLucky
			}

			countTicketsToFindLucky = 0
		}
		countTicketsToFindLucky++
		currentTicket++
	}

	fmt.Println("Наибольшее число билетов, чтобы выпал счастливый -", maxTicketsToBuy)
}

func runAllTasksModule7() {
	module7Task1()
	module7Task2()
	module7Task3()
	module7Task4()
}
