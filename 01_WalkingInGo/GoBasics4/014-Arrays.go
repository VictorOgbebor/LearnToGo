package main

/*
Array = a numbered sequence of elements of a single type, called the element type. The number of elements is called the length of the array and is never negative.
*/

import (
	"fmt"
)

func main() {
	arrays()
}

func arrays() {
	var x [5]int // The number of items in the array
	fmt.Println(x)

	x[3] = 42 // redececlare the 4th spot( 4 = 3 in computers)

	fmt.Println(x)
	fmt.Println(len(x)) // the length of the array
}

// We use Slice more than arrays
