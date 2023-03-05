package main

import "fmt"

func getGoodBye(name string) string{
	return "Good bye " + name
}
func main() {
	sayGoodBye := getGoodBye
	result := sayGoodBye("dodo")
	fmt.Println(result)
	fmt.Println(getGoodBye("toto"))	
}

