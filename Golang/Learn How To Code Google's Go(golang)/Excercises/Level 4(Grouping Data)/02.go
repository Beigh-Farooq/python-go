//● Using a COMPOSITE LITERAL: 
//● create a SLICE of TYPE int 
//● assign 10 VALUES 
//● Range over the slice and print the values out. 
//● Using format printing ○ print out the TYPE of the slice 
package main

import "fmt"

func main() {
	x := []int{12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	fmt.Println(x)
	for v, k := range x {
		fmt.Println(v, k)
	}
	fmt.Printf("%T\n", x)
}