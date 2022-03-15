/*
Function => Allows us to break down our code logic into blocks.

When we create the function we will enter the params = Parameters we are defining what the function does
When we call on the function we will enter the (args = Argument) we are defining what the function does

Everything in Go is passed by VALUE!!! as long as the varible has a value, it can be used

func (reciever) Identifier(parms)(return(s) type){
	codeBlock
}

Similar to what we have been doing
*/

package main

import (
	"fmt"
)

func main() {
	simpleFunc()           // (1) simple we have done this before
	simpleParmsFunc("Vic") // (2) same as before, we will enter the argument now based on the Parameters required
	simpleMultiParmsFunc("Vic", "your the best thing to happen to Blockchain!")

	yo := simpleTypeReturnFunc("Vic", "your the best thing to happen to Blockchain!")
	fmt.Println(yo)

	v, o := multiTypeReturnFunc("Vic", "is he really the best thing to happen to Blockchain!")
	fmt.Println(v)
	fmt.Println("This is", o)

}

func simpleFunc() {
	fmt.Println("Congrats, You made a function...Again!")
}

func simpleParmsFunc(strings string) { // Define the Parameters and the type
	fmt.Println("Yooo", strings)

}

func simpleMultiParmsFunc(strings, greeting string) { // Define the Parameters and the type
	fmt.Println("Yooo", strings, ",", greeting)

}

func simpleTypeReturnFunc(strings, greeting string) string { // Define the Parameters and the type
	return fmt.Sprint("Hello from bic, ", strings, " ", greeting)

}

func multiTypeReturnFunc(strings, greeting string) (string, bool) { // Define the Parameters and the type
	vo := fmt.Sprint("Hello from bic, ", strings, " ", greeting)
	vog := true

	return vo, vog
}
