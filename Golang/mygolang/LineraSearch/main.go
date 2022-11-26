package main

import "fmt"

func LinearSearch(items []int, num int) bool {
	for _, item := range items {
		if item == num {
			return true
		}
	}
	return false
}
func main() {
	items := []int{10, 20, 30, 40}
	if LinearSearch(items, 200) {
		fmt.Println("Element found")
	} else {
		fmt.Println("Element Not found")
	}
}
