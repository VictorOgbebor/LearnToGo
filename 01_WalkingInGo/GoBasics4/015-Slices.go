package main

/*

What are Composite Data Types = A data type that is created using other basxict types and composited types(Structs, Arrays, and List)

slices = Allow you to group together values of the same type ... Like a arrays do
*/
import (
	"fmt"
)

func main() {
	slices()

	sliceRangeLoop()

	slicingSlice()

	sliceAppend()

	sliceDelete()

	sliceMake()

	multiSlice()
}

// 	https://go.dev/ref/spec#Composite_literals
func slices() {

	// Composite Literal = construct values for structs, arrays, slices, and maps and create a new value each time they are evaluated. They consist of the type of the literal followed by a brace-bound list of elements

	// Groups values of same type

	c := []int{4, 51, 67, 8, 25}
	fmt.Println(c)

	// Index the slice = Prints out the  value per position
	fmt.Println(c[0])
	fmt.Println(c[1])
	fmt.Println(c[2])
	fmt.Println(c[3])
	fmt.Println(c[4])

	names := []string{"Bic", "Parky", "OG"}
	fmt.Println(names)

	// Index the slice = Prints out the  value per position
	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

}

func sliceRangeLoop() {
	c := []int{4, 51, 67, 8, 25}
	fmt.Println(c)

	// Index the slice = Prints out the  value per position
	fmt.Println(c[0])
	fmt.Println(c[1])
	fmt.Println(c[2])
	fmt.Println(c[3])
	fmt.Println(c[4])

	for v, o := range c { // Use the for range to get the index and value
		fmt.Println(v, o)
	}

	names := []string{"Bic", "Parky", "OG"}
	fmt.Println(names)

	// Index the slice = Prints out the  value per position
	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	for i, v := range names {
		fmt.Println(i, v)
	}

}

func slicingSlice() {
	fmt.Println("Slicing A Slice")
	c := []int{4, 51, 67, 8, 25}
	fmt.Println(c)

	// Index the slice = Prints out the  value per position
	fmt.Println(c[0])
	fmt.Println(c[1])
	fmt.Println(c[2])
	fmt.Println(c[3])
	fmt.Println(c[4])

	for v, o := range c { // Use the for range to get the index and value
		fmt.Println(v, o)
	}

	names := []string{"Bic", "Parky", "OG"}
	fmt.Println(names)

	// Index the slice = Prints out the  value per position
	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	for i, v := range names {
		fmt.Println(i, v)
	}

	fmt.Println(c[2])

	fmt.Println(c[0:3]) // Selects the index in the code

	for i := 0; i <= 4; i++ {
		fmt.Println(i, c[i]) // will loop the index
	}
}

func sliceAppend() {
	fmt.Println("Appending Slice => Adding to slice")

	c := []int{4, 51, 67, 8, 25}
	fmt.Println(c)

	c = append(c, 11, 12, 13, 140)
	fmt.Println(c)

	y := []int{41, 31, 64, 86, 28}
	c = append(c, y...)
	fmt.Println(y)
	fmt.Println(c)

}

func sliceDelete() {
	fmt.Println("Deleting Slice => Deleting from slice") // Deep Dive into Slices

	c := []int{4, 51, 67, 8, 25}
	fmt.Println(c)

	c = append(c, 11, 12, 13, 140)
	fmt.Println(c)

	y := []int{41, 31, 64, 86, 28}
	c = append(c, y...)
	fmt.Println(c)

	c = append(c[:2], c[4:]...)
	fmt.Println(c)

}

func sliceMake() {

	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3424)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3425)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

func multiSlice() {
	OG := []string{"Victor", "OG", "The Programer I Know", "King"}
	fmt.Println(OG)

	PO := []string{"Parky", "OG", "The GreatestDog", "Elite"}
	fmt.Println(PO)

	DUO := [][]string{OG, PO}
	fmt.Println(DUO)
}
