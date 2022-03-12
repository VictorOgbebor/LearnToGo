package main

import (
	"fmt"
)

func main() {
	fmt.Println("+++++++++++++++++++")
	conditionalsIF()
	fmt.Println("+++++++++++++++++++")
	conditionalsIF_ElSE()
	fmt.Println("+++++++++++++++++++")
	switchStatement()
	fmt.Println("+++++++++++++++++++")
	usingBoolOperators()
}

// if Else = true/false statement to execute statements
func conditionalsIF() {
	fmt.Println("IF Statement")

	if 2 == 2 {
		fmt.Println("1) Print True")
	}

	if 3 != 4 {
		fmt.Println("2) Print False")
	}

	if x := 12 + 6; x == 18 {
		fmt.Println("3) Print True")
	}

	if x := 12 + 6; x == 5 {
		fmt.Println("4) Print False") // should not print anything
	}
}
func conditionalsIF_ElSE() {
	fmt.Println("IF Else Statement")

	if 2 == 2 {
		fmt.Println("1) Print True")
	}

	if 3 != 4 {
		fmt.Println("2) Print False")
	}

	if x := 12 + 6; x == 18 {
		fmt.Println("3) Print True")
	}

	if x := 12 + 6; x == 5 {
		fmt.Println("4) Print True") // should not pting
	} else {
		fmt.Println("4) Print False") // printFalse
	}

	fmt.Println("Example 2")
	b := 25

	if b == 20 {
		fmt.Println("We was valued at 20")
	} else if b == 22 {
		fmt.Println("We was valued at 22")
	} else if b == 25 {
		fmt.Println("We are valued at 25")
	} else {
		fmt.Println("Not The 25 value => The orignal value") // Will loop until Condition is met
	}

}

/*
https://go.dev/ref/spec#Switch_statements

"Switch" statements provide multi-way execution. An expression or type is compared to the "cases" inside the "switch" to determine which branch to execute.
Expression Switch = compare values in expressions
Type Switch = Compare values in types
*/

func switchStatement() {
	fmt.Println("Switch Statement")
	// switch = allows you to switch between test or case
	// case = the reqaurement or action given in code

	switch {
	case false:
		fmt.Println("Should Not Print case 1")

	case (2 == 4):
		fmt.Println("Should Not Print case 2")

	case (3 == 3):
		fmt.Println("This should Print")

	case (4 == 4):
		fmt.Println("also true, does it print")

	}

	fmt.Println("### Using fallthrough ###")
	switch {
	case false:
		fmt.Println("Should Not Print case 1")
		fallthrough
	case (2 == 4):
		fmt.Println("Should Not Print case 2")
		fallthrough
	case (3 == 3):
		fmt.Println("This should Print")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print")

	}

	fmt.Println("### Using Default ###")
	switch {
	case false:
		fmt.Println("Should Not Print case 1")
		fallthrough
	case (2 == 4):
		fmt.Println("Should Not Print case 2")
		fallthrough
	case (3 == 2):
		fmt.Println("This should Print")
		fallthrough
	case (4 == 5):
		fmt.Println("also true, does it print")
	default:
		fmt.Println("LOL...All Failed so should be default")

	}

	fmt.Println("Assigning Variables")
	n := "Bic"

	switch n {
	case "parky":
		fmt.Println("parky")
	case "Bic":
		fmt.Println("Blockchain Bic")
	default:
		fmt.Println("LOL...All Failed so should be default")

	}

}

func usingBoolOperators() {
	fmt.Println("Bool Loops")

	fmt.Println(true && true) // True and true
	fmt.Println(true && false)  // True and false
	fmt.Println(true || true)  // True or true
	fmt.Println(true || false) // True or false
	fmt.Println(!true) // not true
	fmt.Println(!false) // not false

	fmt.Printf("true && true\t %v\n", true && true)
	fmt.Printf("true && false\t %v\n", true && false)
	fmt.Printf("true || true\t %v\n", true || true)
	fmt.Printf("true || false\t %v\n", true || false)
	fmt.Printf("!true\t\t %v\n", !true)
	fmt.Printf("!false\t\t %v\n", !false)
}
