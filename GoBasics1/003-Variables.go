package main
/*
	cannot have unused variables in your code

	Variables Assigns keys to values....
	[Name => Person || CallerId => Phone Number]
	[Name = "String" || CallerId => Int or Some Number]

		Global variables an be used anywhere in code.
		Local variables an be used inside functions.
				Short Declarations used function body. ( := )
*/

import "fmt"


var p = 5 // This variable can everywhere

var zero int // just Assigned the fact its a int type

func main()  {
	x := 50 // This Value can only be used in the function
	fmt.Println(x, ": Is the inital value x")

	x = 100
	fmt.Println(x, ": Notice the Variable has changed for x")

	y := x + 25 // Operators => This is like doing math stuff
	fmt.Println(y, ": Is the inital value y")

	y = 20
	fmt.Println(y, ": Notice the Variable has changed for y")

	fmt.Println(p, ": Notice the global Variable p")

	proofItWorks()
}

func proofItWorks()  {
	fmt.Println(p, ": see it again")

	fmt.Println(zero, ": and this is a zero value...if you dont Assign it something, its zero by default")
}

// Bsst tip is to limit the scope of the variable... If the variable has to be used alot make it global, 
// else keep variable to a function scope

// Go Playground link:  https://go.dev/play/p/9odbSj_0Xf7