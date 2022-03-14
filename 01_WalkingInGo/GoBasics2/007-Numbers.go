/*
https://go.dev/ref/spec#Numeric_types
A numeric type represents sets of integer or floating-point values.
	The predeclared architecture-independent numeric types are:


	Integers = Whole Number || No Decimals
	Real Numbers = Floating points || Deciemals

	8-bit
	16-bit
	32-bit
	64-bit

	uint = Unsigned Integers || No Negatives
	int = Signed Integers || Negatives

	float32
	float64

	complex64
	complex128

	bytes = uint8 (0 - 255)
	runes = int32 (-2147483648 to 2147483647) (can do all the Characters)
*/

package main

import (
	"fmt"
)

var coinbasePrice float64
var chainlinkPrice uint

var badPrice rune = -64
var goodPrice byte = 124

func main() {
	differanceBetweenInts()
	NumberSystem()
}

func differanceBetweenInts() {

	coinAPrice := 75
	coinBPrice := 89.64850

	fmt.Println(coinAPrice)
	fmt.Printf("%T\n", coinAPrice)

	fmt.Println(coinBPrice)
	fmt.Printf("%T\n", coinBPrice)

	fmt.Println(coinbasePrice)
	fmt.Printf("%T\n", coinbasePrice)

	fmt.Println(chainlinkPrice)
	fmt.Printf("%T\n", chainlinkPrice)

	fmt.Println(badPrice)
	fmt.Printf("%T\n", badPrice)

	fmt.Println(goodPrice)
	fmt.Printf("%T\n", goodPrice)

}

func NumberSystem() {

	// Create Simple Algebra
	// Create A Subnetting
}
