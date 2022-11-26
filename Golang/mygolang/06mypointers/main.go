package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// 	var ptr *int
	// 	fmt.Println("value of pointers is ", ptr)

	myNumber := 23
	var ptr = &myNumber //& means refrence

	fmt.Println("Value of actual pointer is ", ptr)  //refrence of memory location
	fmt.Println("Value of actual pointer is ", *ptr) //Value

	*ptr = *ptr + 2
	fmt.Println("New Value is: ", myNumber)

}
