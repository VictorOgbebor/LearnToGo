package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	var myint int8
	for i := 0; i < 129; i++ {
		myint += 1
	}
	fmt.Println(myint) // prints out -127

	var tokenID uint8
	for i := 0; i < 120; i++ {
		tokenID += 1
	}
	fmt.Println(tokenID) // prints out -127
}
