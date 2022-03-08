package main

import "fmt"

func main() {

	fmt.Println("Relational Operators")
	if 5 < 6 {
		fmt.Println("Correct Statement #1")
	}
	// Assigning to || => =
	// equal to || => ==
	if 6 == 6 {
		fmt.Println("Correct Statement #2")
	}

	if 9 == 6 {
		fmt.Println("Correct Statement #3")
	}

	if 9 != 6 {
		fmt.Println("Correct Statement #4")
	}

	fmt.Println("should only work for 3")
	fmt.Println("Logicial Operators")

	// If this OR that works it is fine
	if 5 < 6 || 9 != 6 {
		fmt.Println("Correct Statement #1")
	}

	// this AND this has to works
	if 6 == 6 && 9 != 6 {
		fmt.Println("Correct Statement #2")
	}

	if 6 == 6 && 9 == 6 {
		fmt.Println("Correct Statement #2")
	}

	fmt.Println("Logicial Operators")

}
