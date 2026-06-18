package main

import "fmt"

func main() {

	fmt.Println("Welcome to class of pointers : ")

	//var ptr *int
	//fmt.Println("Value of the pointer is : ", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of the pointer is : ", ptr)
	fmt.Println("Value of the * pointer is : ", *ptr)
	fmt.Println("Value of the & pointer is : ", &ptr)

	*ptr = *ptr * 2
	fmt.Println("Value of the *pointer is : ", *ptr)
	fmt.Println("Value of the number is : ", myNumber)

}
