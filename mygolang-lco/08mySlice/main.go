package main

import "fmt"

func main() {

	fmt.Println("Learning about slices : ")

	var fruitList = []string{"apple", "samosa", "chutney"}

	//fruitList[0] = "samosa"
	//fruitList[1] = "samosa"
	//fruitList[2] = "samosa"

	fmt.Println("printing the fruit list : ", fruitList)

	fruitList = append(fruitList, "bengun")
	fmt.Println("printing the fruit list : ", fruitList)

	fruitList = append(fruitList[1:], "mongo")
	fmt.Println("printing the fruit list : ", fruitList)
}
