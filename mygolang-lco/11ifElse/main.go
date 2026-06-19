package main

import "fmt"

func main() {

	fmt.Println("learning about if else : ")

	loginCount := 23

	if loginCount < 10 {
		fmt.Println("regular user")
	} else if loginCount < 30 {
		fmt.Println("active user")
	} else {
		fmt.Println("way too many logins!! WATCH OUT")
	}

}
