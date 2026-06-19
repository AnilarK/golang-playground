package main

import "fmt"

func main() {
	fmt.Println("Learning about maps : ")

	languages := make(map[string]string)

	languages["1"] = "123"
	languages["2"] = "234,a"
	languages["3"] = "456"

	fmt.Println(languages)
	fmt.Println(languages["1"])

	delete(languages, "1")
	fmt.Println(languages)
	fmt.Println(languages["1"])

	for key, value := range languages {
		fmt.Println("the key is : ", key, " the value is : ", value)
	}

}
