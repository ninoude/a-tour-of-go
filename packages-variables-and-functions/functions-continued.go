package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("1 + 2 =", add(1, 2))
}
