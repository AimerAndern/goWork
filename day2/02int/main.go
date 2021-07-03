package main

import "fmt"

func main() {
	//decimal
	i1 := 101
	fmt.Printf("%d\n", i1)
	//decimal to octal
	fmt.Printf("%o\n", i1)
	//decimal to hex
	fmt.Printf("%x\n", i1)
	//decimal to binary
	fmt.Printf("%b\n", i1)

	//octal
	i2 := 077
	fmt.Printf("%d\n", i2)
	//hex
	i3 := 0xee
	fmt.Printf("%d\n", i3)

	i4 := int16(9)
	fmt.Printf("%T\n", i4)
}
