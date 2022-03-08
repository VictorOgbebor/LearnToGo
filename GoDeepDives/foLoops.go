package main

import "fmt"

func main() {

	// For loop is Go's "while" loops
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 4; j++ {
			fmt.Printf("The Outer loop: %d\t The Inner loop: %d\n", i, j)
		}
	}
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum)

	// If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	println(sum)

}
