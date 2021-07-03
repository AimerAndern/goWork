package main

import "fmt"

func foo() (int, string) {
	return 710, "Rui"
}

func main() {
	//anonymous variable
	x, _ := foo()
	_, y := foo()
	fmt.Println("x =", x)
	fmt.Println("y =", y)
}
