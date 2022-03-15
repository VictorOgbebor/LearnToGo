/*
You can create a func which takes an unlimited number of arguments.
When you do this, this is known as a “variadic parameter.”
When use the lexical element operator “...T” to signify a variadic parameter (there “T” represents some type).
*/
package main

import (
	"fmt"
)

func main() {
	variadicParms(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	sumOfParms(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
}
// when we use (...) => allows us to use Unlimited amount of parms(enter unlimited args)
func variadicParms(z ...int) {
	fmt.Println(z) // This will print int a slice
	fmt.Printf("%T", z)


}

func sumOfParms(z ...int) int {
	fmt.Println(z) // This will print int a slice
	fmt.Printf("%T", z)

	sum := 0
	for t, g := range z {
		sum += g
		fmt.Println("Index position", t, "add", g, "to get the total", sum)

	}
	fmt.Println("The total is", sum)

	return sum
	
}

func sliceFunctions()  {
	
}
