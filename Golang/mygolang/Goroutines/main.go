package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}
func main() {
	//calling Goroutine
	go display("Welcome")

	//calling normal function
	display("GreeksforGeeks")

}
