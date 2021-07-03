package main

import "fmt"

//declare multiple variables
var (
	name  string //""
	age   int    //0
	alive bool   //false
)

func main() {
	name = "Rui Zhu"
	age = 16
	alive = true
	fmt.Print(alive) //out content on terminal
	fmt.Println()
	fmt.Printf("name:%s\n", name) //%s is used to replace with the content of variable name
	fmt.Println(age)              //get to the new line after printing the content of the variable
	//declare and initialize
	var s1 = "rz"
	fmt.Println(s1)
	var s2 = "20"
	fmt.Println(s2)
	//brief initialization
	s3 := "jajaja"
	fmt.Println(s3)
}
