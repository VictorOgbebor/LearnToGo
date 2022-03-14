package main

/*
Maps is like an unordered lisy
*/
import (
	"fmt"
)

func main() {
	maps()

	mapsAddRange()

	mapDelete()
}

func maps() {
	mapExample := map[string]int{
		"Bic":   25,
		"Parky": 5,
	}

	fmt.Println(mapExample) // Shows

	fmt.Println(mapExample["Parky"]) // Will Return the value of the mapping or object

	fmt.Println(mapExample["parky"]) // Naming convention matters = will returns zero value

	v, ok := mapExample["parky"]

	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := mapExample["parky"]; ok {
		fmt.Println("This prints If", v)
	}

	if v, ok := mapExample["Parky"]; ok {
		fmt.Println("This is the Value = ", v)
	}

}

func mapsAddRange() {

	mapExample := map[string]int{
		"Bic":   25,
		"Parky": 5,
	}

	fmt.Println(mapExample) // Shows

	fmt.Println(mapExample["Parky"]) // Will Return the value of the mapping or object

	fmt.Println(mapExample["parky"]) // Naming convention matters = will returns zero value

	v, ok := mapExample["parky"]

	fmt.Println(v)
	fmt.Println(ok)

	mapExample["OG"] = 50

	if v, ok := mapExample["parky"]; ok {
		fmt.Println("This prints If", v)
	}

	if v, ok := mapExample["Parky"]; ok {
		fmt.Println("This is the Value = ", v)
	}

	for k, v := range mapExample {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 7, 8, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}

}

func mapDelete() {

	mapExample := map[string]int{
		"Bic":   25,
		"Parky": 5,
	}

	delete(mapExample, "Bic")
	fmt.Println(mapExample)

	if v, ok := mapExample["Parky"]; ok {
		fmt.Println("value:", v)
		delete(mapExample, "Parky")
	}

}
