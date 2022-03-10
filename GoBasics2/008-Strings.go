/*
A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes.
The number of bytes is called the length of the string and is never negative. Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is string; it is a defined type.
https://go.dev/ref/spec#String_types
*/

package main

import (
	"fmt"
)

func main() {
	basicString()
	stringConversion()
}

func basicString() {
	aString := `This is a String`
	fmt.Println(aString)

	fmt.Printf("%T\n", aString)
}

func stringConversion() {
	aString := "This is a String"
	fmt.Println(aString)

	fmt.Printf("%T\n", aString)

	convertString := []byte(aString)
	fmt.Println(convertString)
	fmt.Printf("%T\n", convertString) // uint8 = byte

	// index positions || https://go.dev/blog/strings => In Depth
}
