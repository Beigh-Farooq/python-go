//● Using a COMPOSITE LITERAL: 
//● create an ARRAY which holds 5 VALUES of TYPE int 
//● assign VALUES to each index position. 
//● Range over the array and print the values out. 
//● Using format printing ○ print out the TYPE of the array
package main
func main(){
	x := [5]int{12, 13, 14, 15, 16}
	fmt.Println(x)
	for v, k := range x {
		fmt.Println(v, k)
	}
	fmt.Printf("%T\n", x)
} 