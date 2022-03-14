package main // (1) Import package

import "fmt" // (2) This is a Package imported to go

/*
(3)
	The main Package or Func is wehere all the code logic will go Through.
	This is where other logic and functions are called
*/

func main() {

	// (4) Print using fmt
	fmt.Println("Welcome to Writing in Go")

	// (5.0) This function is call in below
	yo()

	/*
		(6)
		we will explain alot in here more later
		Whats here:
			Variables and Short Declaractions
			Control Flow and For Loops
				(1) Sequence
				(2) Loop and Iterate
				(3) Conditional

		below loop adds to 100
	*/
	for p := 0; p < 100; p++ {
		println(p)
	}

}

// (5.1) Create a function that will be called in the main fuction. *Note*...This func has to be called inside the main func to work
func yo() {

	// (5.2) Print can be used with out the fmt import
	println("Enbrace the learning process. Its not gonna be easy...But Definitely will set you apart from the world") // Print can be used with out the fmt import
}

// Go Playground link: https://go.dev/play/p/e3dq7P9Q9Qs
