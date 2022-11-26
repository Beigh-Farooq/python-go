package main

import "fmt"

func main() {
	var rows = 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows; j++ {
			fmt.Print("*")
		}
		fmt.Println()
		rows = rows - 1
	}
}
