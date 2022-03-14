/*
|| https://go.dev/doc/effective_go#control-structures
https://go.dev/ref/spec#Statements

Control Flow allows = How the computer will read the code. Like a recipe. Like a play in a playbook. The steps to Success based on Conditions
	Loops
	Conditions

	for i := 0; i > 10: i++ {

	}

	for i := "init or Initialize Variable using the short Declaration"; Condition statement = Some math or Boolean; post = the PostAction result {
		print a statement or call a function
	}
*/

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Standard Go ControlFlow")

	//simpleLoop()
	//nestedLoop()
	// forStatement() // This will run forever
	breakContinue()
	// printingStatement()
	//conditionalsIF()
	//conditionalsIF_ElSE()
	//switchStatement()
	//usingBoolOperators()

}

func simpleLoop() {
	fmt.Println("SimpleLoops")

	for i := 25; i <= 100; i++ {
		fmt.Println(i)
	}
}

func nestedLoop() {
	fmt.Println("Nesting Loops")

	// Inner Loop Will run when the outer loop runs
	for loop1 := 0; loop1 <= 10; loop1++ {
		fmt.Printf("outsideLoop: %d\n", loop1)
		for loop2 := 0; loop2 < 5; loop2++ {
			fmt.Printf("\t\t insideLoop: %d\t", loop2)

		}
	}

}

// https://go.dev/ref/spec#For_statements
// repeated execution of a block of code => like a while statement
func forStatement() {
	fmt.Println("For Statements in loops")

	// simple for loop or While Statment
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")

	// For Statement W
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")

}

//  for loop
func breakContinue() {
	fmt.Println("Break and Continue")

	x := 0
	for {
		x++
		if x > 100 {
			break // Break when hit over 100
		}

		if x%2 != 0 {
			continue // will continue if conditon is met
		}

		fmt.Println(x)
		// math Remander in Go
		x := 83 / 40
		y := 83 % 40
		fmt.Println(x, y)

	}
	fmt.Println("done.")

}

func printingStatement() {
	x := 33
	for {
		if x > 122 {
			break
		}
		fmt.Printf("%v corresponds to %+q in ASCII\n", x, x)
		x++
	}
}

/*
go fmt
go run
*/
