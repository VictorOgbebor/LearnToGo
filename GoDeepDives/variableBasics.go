package main

import "fmt"

func main() {
	var var1 = 3
	var str1 = "stringInHere"
	var c, python, java bool
	var i int
	// all numeric types default to 0

	fmt.Println(var1, str1)
	fmt.Println("   ")

	changeVar()
	fmt.Println("   ")
	types()
	fmt.Println("   ")
	floatingNumbers()
	fmt.Println("   ")
	boolean()

	fmt.Println(i, c, python, java)
}

func changeVar() {
	var x = 10
	const y = 24
	fmt.Println(x)
	fmt.Println("   ")
	x = x + y
	fmt.Println(x)
	fmt.Println("   ")
	x = 2
	fmt.Println(x)
	fmt.Println(y)
}

// const is like const in JavaScript
// var is like let in JavaScript

func types() {
	var men uint64
	var women int64
	var people int

	men = 5
	women = 7

	people = int(men) + int(women)

	fmt.Println(people)

	// // unsigned int with 8 bits
	// // Can store: 0 to 255
	// var myint uint8

	// // signed int with 8 bits
	// // Can store: -127 to 127
	// var myint int8

	// // unsigned int with 16 bits
	// var myint uint16

	// // signed int with 16 bits
	// var myint int16

	// // unsigned int with 32 bits
	// var myint uint32

	// // signed int with 32 bits
	// var myint int32

	// // unsigned int with 64 bits
	// var myint uint64

	// // signed int with 64 bits
	// var myint int64

}

func floatingNumbers() {
	// These come in 2 distinct sizes, either float32 or float64 and allow you to work with exceptionally large numbers that don’t fit inside a standard int64 data type.
	// var f1 float32
	// var f2 float64
	var maxFloat32 float32
	maxFloat32 = 16777216
	fmt.Println(maxFloat32 == maxFloat32+10) // you would typically expect this to return false
	// it returns true
	fmt.Println(maxFloat32 + 10)      // 16777216
	fmt.Println(maxFloat32 + 2000000) // 16777216

	// converting from int to float
	var myint1 int = -29
	myfloat := float64(myint1)

	// converting from float to int
	var myfloat2 float64 = 1
	myint2 := int(myfloat2)

	fmt.Println(myfloat)
	fmt.Println(myint2)

}

func boolean() {
	var isProfitable bool = true
	var isNotProfitable bool = false
	// And
	if isProfitable && isNotProfitable {
		fmt.Println("Both Conditions need to be True")
	}
	// Or
	if isProfitable || isNotProfitable {
		fmt.Println("Only one condition needs to be True")
	}
}
