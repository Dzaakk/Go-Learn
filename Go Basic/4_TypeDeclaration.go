package main

import "fmt"

func main() {

	type NoKTP string
	type Married bool

	var NoKTPUser NoKTP = "123113123"
	var MarriedStatus Married = true
	fmt.Println(NoKTPUser)
	fmt.Println(MarriedStatus)
}
