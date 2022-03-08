package main

func strings(x, y string) (string, string) {
	return x, y
}

func swapStrings(x, y string) (string, string) {
	return y, x
}

func main() {

	x, y := strings("Hello", "world")
	println(x, y)

	y, x = swapStrings("Hello", "world")
	println(y, x)

}
