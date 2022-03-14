package main

/*
Struct = composite data type. (composite data types, aka, aggregate data types, aka, complex data types)

Embedded structs = take one struct and embed it in another struct. When you do this the inner type gets promoted to the outer type.
*/
import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type superStar struct {
	person
	MVP bool
}

func main() {
	structs()

	EmbeddedStructs()
}

func structs() {

	person1 := person{
		first: "Blockchain",
		last:  "Bic",
		age:   3,
	}

	person2 := person{
		first: "Parky",
		last:  "OG",
		age:   5,
	}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.first, person1.last, person1.age)

	fmt.Println(person2.first, person2.last, person2.age)

}

func EmbeddedStructs() {

	superStar := superStar{
		person: person{
			first: "OG",
			last:  "Parky",
			age:   5,
		},
		MVP: true,
	}

	fmt.Println(superStar)
	fmt.Println(superStar.first, superStar.last, superStar.age, superStar.MVP)
}
