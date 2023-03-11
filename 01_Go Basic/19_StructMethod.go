package main

import "fmt"

type Customer struct {
	Name, Address string
	age           int
}

func (a Customer) sayHi(name string) {
	fmt.Println("Hi", name, "My name is", a.Name)
}
func main() {
	var mamat Customer
	mamat.Name = "mamat"
	mamat.Address = "Indonesia"
	mamat.age = 22
	mamat.sayHi("toto")
}