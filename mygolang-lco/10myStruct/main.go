package main

import "fmt"

func main() {
	fmt.Println("Learning about struct : ")

	user := User{"Abhay", 201, "abhay@gmail.com", true, 23}

	fmt.Println(user)
	fmt.Println(user.Email)
}

type User struct {
	Name   string
	RollNo int
	Email  string
	Status bool
	Age    int
}
