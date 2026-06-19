package main

import "fmt"

func main() {

	fmt.Println("learning about functions :")
	greet()

	fmt.Println(adder(3, 5))

}

func greet() {
	fmt.Println("you have been greeted!!  \nhe he")
}

func adder(var1 int, var2 int) int {
	return var1 + var2
}
