package main

import "fmt"

const LoginToken string = "ghabbhjd" //public

func main() {
	var username string = "startupkashmir"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLogged bool = false
	fmt.Println(isLogged)
	fmt.Printf("variable is of type: %T \n", isLogged)

	welcome := "welcome to go"
	fmt.Println(welcome)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)
}
