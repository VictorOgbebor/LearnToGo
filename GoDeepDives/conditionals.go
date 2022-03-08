package main

import "fmt"

func main() {

	var a = 25
	var b = 29

	// if - Else statments
	if 9 > 10 {
		fmt.Println("Statement 1 is good")
	} else {
		fmt.Println("Statement 1 is bad: Fix it")
	}

	// if - Else statments BOTH Conditions must work
	if a < b && a == 25 {
		fmt.Println("Statement 2 is good")
	} else {
		fmt.Println("Statement 2 is bad: Fix it")
	}

	// if - Else statments BOTH Conditions must work
	if a > b || a == 26 {
		fmt.Println("Statement 3 is good")
	} else {
		fmt.Println("Statement 3 is bad: Fix it")
	}

	// if - Else If statments => Gives the sript options to execute...ells is the last statement

	if a < b {
		fmt.Println("Statement 4 is good, try other option")
	} else if a > b {
		fmt.Println("Statement 4 is good")
	} else {
		fmt.Println("Statement 4 is bad: Fix it")
	}
	loops()
}

func loops() {
	println("If Statements in Go")
	var customerHeight int = 140
	customerAge := 18

	if customerHeight >= 150 || customerAge >= 18 {
		println("can access ride")
	} else if customerHeight >= 120 {
		println("customer can access children's rides")
	} else {
		println("customer is too small")
	}

}
