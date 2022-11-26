//create a func with the identifier foo that returns an int  
//create a func with the identifier bar that returns an int and a string  
//call both funcs 
// print out their results

package main 
func main(){
	n:=foo()
	x,s:=bar()
	fmt.Println(n,x,s)
}
func foo() int{
	return 42
}
func bar() (int,string){
	return 1984,"Big Brother is watching"
}