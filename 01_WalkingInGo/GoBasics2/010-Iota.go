// https://go.dev/ref/spec#Iota => Used to Increment constant declaration, starting at zero
package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)

func main() {
	iotaStuff()
}

func iotaStuff() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
