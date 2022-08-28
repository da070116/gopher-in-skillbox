package internal

import (
	"fmt"
	"gopher-in-skillbox/pkg"
	"math/rand"
	"time"
)

func multiTypeSolve(x int16, y uint8, z float32) (result float32) {
	x_ := float32(x)
	y_ := float32(y)
	result = 2*x_ + y_*y_ - float32(3)/z

	return
}

func m21Task1() {
	var x int16 = -1<<15 + 1
	var y uint8 = 254
	var z float32 = 0.01
	fmt.Println(multiTypeSolve(x, y, z))
}

func mathWithRandom(F func(int, int) int) {

	a := rand.Intn(100)
	b := rand.Intn(a)
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	defer fmt.Println("Func result =", F(a, b))
}

func m21Task2() {
	rand.Seed(time.Now().UnixNano())
	mathWithRandom(func(i1 int, i2 int) int { return i1 + i2 })
	mathWithRandom(func(i1 int, i2 int) int { return i1 - i2 })
	mathWithRandom(func(i1 int, i2 int) int { return i1 / i2 })
}

func m21RunAll() {
	pkg.Wrapper(m21Task1)
	pkg.Wrapper(m21Task2)
}
