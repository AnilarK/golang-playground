package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Learning about time : ")

	presentTime := time.Now()
	fmt.Println("Current time is : ", presentTime)
	fmt.Println(presentTime.Format("01-01-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
