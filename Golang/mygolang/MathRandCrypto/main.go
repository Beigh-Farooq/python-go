package main

import (
	"fmt"
	"math/big"

	//"math/big"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	//var mynumberone int = 2
	//var mynumbertwo float64 = 4
	//fmt.Println("The sum is: ", mynumberone+int(mynumbertwo))

	//random number
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(10))

	//random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
