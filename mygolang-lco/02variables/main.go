package main

import "fmt"

const LoginToekn string = "ABXYZ" // public

func main() {

	fmt.Println("Learning about variables : ")

	var username string = "anilark"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float32 = .123398429387248578944
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	//	Alias and default value

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	//	implicit type
	var website = "www.google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)

	//	 no var style (only inside a method)
	numberOfUser := 30000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type : %T \n", numberOfUser)

	fmt.Println(LoginToekn)
	fmt.Printf("Variable is of type : %T \n", LoginToekn)
}
