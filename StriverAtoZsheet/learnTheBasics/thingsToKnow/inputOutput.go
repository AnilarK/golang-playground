package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("lets solve")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Taking input of a int : ")
	xStr, _ := reader.ReadString('\n')
	xInt, _ := strconv.ParseInt(strings.TrimSpace(xStr), 10, 32)
	fmt.Println("the integer input is : ", xInt)

	fmt.Println("Taking input of a float : ")
	xStr, _ = reader.ReadString('\n')
	xFloat, _ := strconv.ParseFloat(strings.TrimSpace(xStr), 32)
	fmt.Println("the float input is : ", xFloat)

	fmt.Println("Taking input of a bool : ")
	xStr, _ = reader.ReadString('\n')
	xBool, _ := strconv.ParseBool(strings.TrimSpace(xStr))
	fmt.Println("the float input is : ", xBool)

	fmt.Println("Taking input of a string : ")
	xStr, _ = reader.ReadString('\n')
	fmt.Println("the string input is : ", xStr)

}
