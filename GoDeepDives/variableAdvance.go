package main

import "fmt"

var (
	v, o = 1, 2
	Profit, notProfit bool
)

func main() {
	var JavaScript, Python, Typescript, Go = "good", "good", "bad", "great"

	// Short Declaration
	Parky := "name"
	ParkyAge := 5

	fmt.Println(v, o, JavaScript, Python, Typescript, Go)
	fmt.Println(v, Go, o)
	fmt.Println(Profit, notProfit)
	fmt.Println(Parky, ParkyAge)

}
