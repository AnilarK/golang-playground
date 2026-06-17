package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Learning user Input : ")

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the reading of pizza : ")

	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of the Rating is %T \n", input)

}
