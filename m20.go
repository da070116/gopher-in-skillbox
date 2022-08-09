package main

import "fmt"

const (
	MATRIXSIZE3 = 3
	MATRIXSIZE4 = 4
	MATRIXSIZE5 = 5
)

func createSquareMatrix() (result [MATRIXSIZE3][MATRIXSIZE3]int) {
	for i := 0; i < MATRIXSIZE3; i++ {
		for j := 0; j < MATRIXSIZE3; j++ {
			fmt.Printf("Введите [%v]-й элемент [%v]-й строки матрицы:", j+1, i+1)
			_, _ = fmt.Scan(&result[i][j])
		}
	}
	return
}

func create3x5Matrix() (result [MATRIXSIZE3][MATRIXSIZE5]int) {
	for i := 0; i < MATRIXSIZE3; i++ {
		for j := 0; j < MATRIXSIZE5; j++ {
			//fmt.Printf("Введите [%v]-й элемент [%v]-й строки матрицы:", j+1, i+1)
			//_, _ = fmt.Scan(&result[i][j])
			result[i][j] = customRandInt(100)
		}
	}
	return
}

func create5x4Matrix() (result [MATRIXSIZE5][MATRIXSIZE4]int) {
	for i := 0; i < MATRIXSIZE5; i++ {
		for j := 0; j < MATRIXSIZE4; j++ {
			//fmt.Printf("Введите [%v]-й элемент [%v]-й строки матрицы:", j+1, i+1)
			//_, _ = fmt.Scan(&result[i][j])
			result[i][j] = customRandInt(10)
		}
	}
	return
}

func m20Task1() {
	// Вычислить определитель матрицы 3*3
	m := createSquareMatrix()

	det := m[0][0]*m[1][1]*m[2][2] + m[0][1]*m[1][2]*m[2][0] + m[0][2]*m[1][0]*m[2][1] - m[0][2]*m[1][1]*m[2][0] - m[0][0]*m[1][2]*m[2][1] - m[0][1]*m[1][0]*m[2][2]
	fmt.Println("Определитель матрицы равен", det)
}

func m20Task2() {
	// Вычислить произведение матриц
	m1 := create3x5Matrix()

	fmt.Println("Матрица 3 на 5: \n", m1)

	m2 := create5x4Matrix()
	fmt.Println("Матрица 5 на 4: \n", m2)

	var resultMatrix [MATRIXSIZE3][MATRIXSIZE4]int

	for i := 0; i < MATRIXSIZE3; i++ {
		for j := 0; j < MATRIXSIZE4; j++ {
			for k := 0; k < MATRIXSIZE5; k++ {
				resultMatrix[i][j] += m1[i][k] * m2[k][i]
			}
		}
	}
	// Вывести результат:
	for i := 0; i < len(resultMatrix); i++ {
		fmt.Println(resultMatrix[i])
	}
}

//Execute all tasks in this module
func m20RunAll() {
	wrapper(m20Task1)
	wrapper(m20Task2)
}
