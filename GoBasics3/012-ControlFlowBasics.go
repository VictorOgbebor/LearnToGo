/*
|| https://go.dev/doc/effective_go#control-structures
Control Flow allows = How the computer will read the code. Like a recipe. Like a play in a playbook. The steps to Success based on Conditions
	Loops
	Conditions

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
	//forStatement() // This will run forever
	breakContinue()
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

	goodPrice := 6 //
	buy := true

	// simple

	for goodPrice < 10 {

		fmt.Println("assets = ", goodPrice)
		fmt.Println(buy, "Good Buy")
		goodPrice++
	}

}

func breakContinue() {
	fmt.Println("Break and Continue")

	goodPrice := 1 //
	buy := true

	for {

		if goodPrice < 9 {
			break
		}
		fmt.Println("assets = ", goodPrice)
		fmt.Println(buy, "Good Buy")
		goodPrice++

	}

}

func conditionalsIF() {
	fmt.Println("IF Statement")
}
func conditionalsIF_ElSE() {
	fmt.Println("IF Else Statement")
}

func switchStatement() {
	fmt.Println("Switch Statement")
}

func usingBoolOperators() {
	fmt.Println("Bool Loops")
}

/*
go fmt
go run
*/
