package main

import (
	"fmt"
)

func main() {
	anonymousFunctions()
	return1 := returningFunctionString()
	fmt.Println(return1)

	someNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	o := add(someNumbers...)
	fmt.Println(o)

	fmt.Println("**************************************************")

	a := ClosureFuncs()

	b := ClosureFuncs()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println("**************************************************")

	fmt.Println(b())
	fmt.Println(b())

	fmt.Println("**************************************************")

	r := RecursiveFuncs(2)
	fmt.Println(r)

}

func anonymousFunctions() {

	// Anonymous Functions or Local Functions
	func() {
		fmt.Println("Anonymous with no Params")
	}()

	func(x int) {
		fmt.Println("Anonymous with Params", x)
	}(25)

	f := func() {
		fmt.Println("Function as a Variable")
	}
	f()

	k := func(x int) {
		fmt.Println("Function as a Variable ad assign type int", x)
	}
	k(50)
}

func returningFunctionString() string {
	s := "Hello"
	return s
}

// Callback Funcs = Pass a func as a argument
func CallbackFuncs() {

}

func add(xInt ...int) int {
	total := 0

	for _, jz := range xInt {
		total += jz
	}
	return total
}

// Closure => limits the scope of each
func ClosureFuncs() func() int {

	var x int

	return func() int {
		x++
		return x
	}
}

// Recursive => Func call on itself ||=> We can Use a Loop instead
func RecursiveFuncs(r int) int {

	if r == 0 {
		return 1
	}
	return r * RecursiveFuncs(r-1)

}
