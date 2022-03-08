package main

import (
	"fmt"
	"math"
)

func addSomething(x, y int) int {
	return x + y
}

func subSomething(x, y int) int {
	return x - y
}

func mulSomething(x, y int) int {
	return x * y
}

func divSomething(x, y int) int {
	return x / y
}

func sqRoot(x float64) string {
	if x < 0 {
		return sqRoot(-x) + "OG"
	}

	return fmt.Sprint(math.Sqrt(x))
}

// Named return values

func main() {
	println(addSomething(5, 8))
	println(subSomething(8, 5))
	println(mulSomething(8, 5))
	println(divSomething(40, 5))
	fmt.Println(sqRoot(144), sqRoot(64))

}
