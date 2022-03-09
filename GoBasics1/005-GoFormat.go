/*
	How fo we Print and Format our values
	https://gobyexample.com/string-formatting
	https://go.dev/ref/spec#Rune_literals
	https://pkg.go.dev/fmt#hdr-Printing

	Printf => Print Format
*/

package main

import "fmt"

var y = 21

func main() {

	fmt.Println(y)
	printNumbers()
	
	fmt.Println("\t")
	printStrings()

}

func printNumbers() {

	// Number Systems =>
	fmt.Printf("%T\n", y)              // Display Type
	fmt.Printf("%b\n", y)              // Display Binary
	fmt.Printf("%x\n", y)              // Display Hexadecimal
	fmt.Printf("%#x\n", y)             // Display 0x
	fmt.Printf("%#x\t%b\t%x", y, y, y) // Display 0x, binary, and precentage

}

func printStrings() {

	fmt.Println("YOOOOO")
	lol := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y) /// String Printing

	fmt.Println(lol)

	// FPrint is another one we will use for web server and files stuff || "variadic parameter" = means that the func can take as many values of that type as you want to pass in
}

// Go Playground link: https://go.dev/play/p/7bileHGdH-K
