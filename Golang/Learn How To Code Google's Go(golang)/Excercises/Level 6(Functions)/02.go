
//create a func with the identifier foo that takes in a variadic parameter of type int 
//pass in a value of type []int into your func (unfurl the []int) 
//returns the sum of all values of type int passed in 

//● create a func with the identifier bar that takes in a parameter of type []int  
//returns the sum of all values of type int passed in package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(ii...)
	fmt.Println(n)

	ii = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := bar(ii)
	fmt.Println(b)
}
func foo(xi ...int) int { //unfurl the slice of int
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
