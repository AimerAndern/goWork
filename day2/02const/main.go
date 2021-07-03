package main

import "fmt"

// constant

//iota+=1 if there's a new line declaration
const (
	d1, d2 = iota + 1, iota + 2 //iota=0
	d3, d4 = 100, 200           //iota=1
	d5, d6 = iota + 1, iota + 2 //iota=2
)

const (
	_  = iota             //anonymous variable =0
	KB = 1 << (10 * iota) //binary:1^(10*(iota=1))
	MB = 1 << (10 * iota) //binary:1^(10*(iota=2))
	GB = 1 << (10 * iota) //binary:1^(10*(iota=3))
	TB = 1 << (10 * iota) //binary:1^(10*(iota=4))
)

func main() {

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
	fmt.Println(d5)
	fmt.Println(d6)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)

}
