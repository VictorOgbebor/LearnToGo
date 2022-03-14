// https://go.dev/ref/spec#Constants => variables that dont change
package main

import (
	"fmt"
)

func main() {
	constantsVariable()
}

const (
	typedConstant    int    = -55
	UnsignedConstant uint   = 11
	StringConstant   string = " Me"
)

func constantsVariable() {
	fmt.Println(typedConstant)
	fmt.Printf("%T\n", typedConstant)

	fmt.Println(UnsignedConstant)
	fmt.Printf("%T\n", UnsignedConstant)

	fmt.Println(StringConstant)
	fmt.Printf("%T\n", StringConstant)
}
