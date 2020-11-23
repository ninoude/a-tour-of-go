package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("入れ替わってる？？！！", "これってもしかして、俺たち・私たち、")
	fmt.Println(a, b)
}
