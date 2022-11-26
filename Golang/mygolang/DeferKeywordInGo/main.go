// defer will move the execution of the statement to the very end inside a function
// Multiple defer statements are allowed in the same program and will be executed in
// reverse order in which they are deffered.
// In other words they will follow the LIFO order.
package main

import "fmt"

func main() {
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello")
	myDefer()
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
