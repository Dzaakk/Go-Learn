package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
func main() {
	var dodi Person
	dodi.Name = "dodi"

	SayHello(dodi)
	cat := Animal{
		Name: "kucing",
	}
	SayHello(cat)
}
