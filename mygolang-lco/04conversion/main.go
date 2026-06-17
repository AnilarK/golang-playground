package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our pizza app")
	fmt.Println("please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thenks for rating, ", input)

	numRating, err := strconv.ParseInt(strings.TrimSpace(input), 10, 32)

	if err != nil {
		fmt.Println(err)
	} else {
		numRating += 1
		fmt.Println("Added 1 to your rating: ", numRating)
	}

}
