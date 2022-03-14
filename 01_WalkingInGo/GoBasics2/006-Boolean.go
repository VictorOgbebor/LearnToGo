/*
BOOLEAN TYPE REPRESENTS A VALUE THAR WILL THE TRUE OR FALSE
https://go.dev/ref/spec#Boolean_types

this is also a Introduction to calling functions
*/

package main

import (
	"fmt"
)

var boolTypeTrue bool
var boolTypeFalse bool // zero value = false

func main() {
	bool1()
	bool2()

}

func bool1() {
	fmt.Println(boolTypeFalse, boolTypeTrue)
	boolTypeTrue = true
	fmt.Println(boolTypeFalse, boolTypeTrue)
}

func bool2() {
	o := 25
	g := 96

	fmt.Println(o == g)
	fmt.Println(o <= g)
	fmt.Println(o >= g)
	fmt.Println(o != g)
}
