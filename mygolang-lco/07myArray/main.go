package main

import "fmt"

func main() {
	fmt.Println("Learning about arrays : ")

	var fruitList [5]string

	fruitList[0] = "sasmosa"
	fruitList[2] = "apple"

	fmt.Println("Printing the array : ", fruitList)
	fmt.Println("length the array : ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("veggie list is : ", vegList)
	fmt.Println("length of veggie list is : ", len(vegList))

}
