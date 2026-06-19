package main

import (
	"fmt"
	"sort"
)

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

	highScores := make([]int, 4)
	highScores[0] = 43
	highScores[1] = 98
	highScores[2] = 76
	highScores[3] = 58

	highScores = append(highScores, 34, 49)

	fmt.Println("the scores are : ", highScores)

	sort.Ints(highScores)
	fmt.Println("the sorted scores are : ", highScores)
}
