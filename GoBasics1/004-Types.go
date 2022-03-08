package main

/*
Review Docs:
	This build upon previous lessons. We will be a basic overview into
	what types are.

what are types?
	(Basic Programming)Primitive = Bool, String, Int
	Composite = data that we create(Structs, Slices, Methods)


how to use them?
	values can only hold to the same TYPEs
		Int = Int || String = String
		Int != String || String != Int

*Remember This Please*

*/

import "fmt"

// Try removing the intial values here to give the zero value
var name string = "NAME YOURSELF"
var numbers int = 25

type Type1 int

var int1 Type1

type Type2 string

var string1 Type2

func main() {

	types()
	fmt.Println("\t")
	createTypes()
}

func types() {
	println(name)
	fmt.Printf("%T\n", name)

	name := "BlockchainBic"
	fmt.Printf("%T\n", name)
	println(name)

	println(numbers)
	fmt.Printf("%T\n", numbers) // "%T\n" => (%T = shows the type  || \n = new line)

	numbers := 00
	println(numbers)
	fmt.Printf("%T\n", numbers)

}

func createTypes() {

	int1 := 22
	println(int1)
	fmt.Printf("%T\n", int1)

	string1 := "YOOOO"
	println(string1)
	fmt.Printf("%T\n", string1)

	// Conversions in type {Fill in here}
}

// Go Playground link:  https://go.dev/play/p/ZAyliKwGBz9
