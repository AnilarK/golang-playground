package main

import "fmt"

func main() {

	fmt.Println("Learning about Loops : ")

	days := []string{"a", "b", "c", "d", "e", "f"}

	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for _, day := range days {
		fmt.Println(day)
	}
}
