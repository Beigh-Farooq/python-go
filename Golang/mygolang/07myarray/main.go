package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit List is:", fruitList)
	fmt.Println("Length of Fruit List is:", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy List is:", vegList)
	fmt.Println("Length of Vegy List is:", len(vegList))
}
