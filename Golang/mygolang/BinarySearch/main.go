package main

import "fmt"

func BinarySearch(items []int, value int) {
	min := 0
	max := len(items)
	mid := 0
	for min < max {
		mid = (min + max) / 2
		if value == items[mid] {
			break
		} else if value > items[mid] {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	if min >= len(items) {
		fmt.Println("Item not found")
	} else {
		fmt.Println("Item Found at=", mid+1)
	}
}
func main() {
	items := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	BinarySearch(items, 50)
}
