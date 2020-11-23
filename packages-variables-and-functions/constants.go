package main

import "fmt"

const four = 4

func main() {
	const World = "死んでいる"
	fmt.Println("お前はもうすでに", World)
	fmt.Println("お前はもうすでに", four, "んでいる")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
