package main

import (
	"fmt"
)

type pup struct {
	name    string
	color   string
	age     int
	goodDog bool
}

type yaba struct {
	pup
	beautifulDog bool
}

type parky struct {
	pup
	beautifulDog bool
}

type dog interface {
	bark() // In this interface, if it calls the bark function at all...it will classify itself as a dog. *Inheritance*
}

func main() {
	defer defer1() // When We Defer it will not excute until others after it does
	defer defer2() // This allows us to be concise...so we dont have so many functions running at once

	myDogs()

}

// Defer =
func defer1() {
	fmt.Println("LOL 1")
}

func defer2() {
	fmt.Println("LOL 2")
}

func myDogs() {

	p := parky{
		pup: pup{
			name:    "Parka",
			color:   "GreyishBrown",
			age:     5,
			goodDog: true,
		},
		beautifulDog: true,
	}

	yaya := yaba{
		pup: pup{
			name:    "Ayaba",
			color:   "Black",
			age:     2,
			goodDog: true,
		},
		beautifulDog: true,
	}

	chaz := pup{
		"chaz",
		"Whatever",
		0,
		true,
	}

	fmt.Println(p, yaya, chaz)

	yaya.bark()

	newDog(chaz)
	newDog(p)
}

func (bark pup) bark() {
	fmt.Println("Who is Doing all that barking?...It might be....", bark.name)
}

func newDog(d dog) {
	fmt.Println("I AM a Dog!!!", d)
}
