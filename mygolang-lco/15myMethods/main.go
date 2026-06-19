package main

import "fmt"

func main() {
	fmt.Println("Learning about methods :")

	user := User{"abhay", "abhay@gmail.com", true}
	fmt.Println(user)
}

type User struct {
	Name   string
	Email  string
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("is user verified : ", u.Status)
}
