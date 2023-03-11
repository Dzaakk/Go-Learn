package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}
func sayGoodbyeFIlter(name string, filter Filter) {
	NameFilter := filter(name)
	fmt.Println("Goodbye", NameFilter)
}
func NameFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
func main() {
	sayHelloWithFilter("lili", NameFilter)
	sayHelloWithFilter("Anjing", NameFilter)
	sayGoodbyeFIlter("Tro", NameFilter)
	sayGoodbyeFIlter("Anjing", NameFilter)
}
