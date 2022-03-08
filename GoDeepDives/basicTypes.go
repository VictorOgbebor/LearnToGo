package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println("+++++++++++++++++++++++++++++++++++++++")
	// Zero Value
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	fmt.Println("+++++++++++++++++++++++++++++++++++++++")
	var x, y int = 3, 4
	var g float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(g)
	fmt.Println(x, y, z)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++")

	q := 0.867 + 0.5i // change me!
	f = 3.142
	fmt.Printf("q is of type %T\n", q, "f is of type %T\n", f)

	fmt.Println("+++++++++++++++++++++++++++++++++++++++")
}
