package main

import (
	"fmt"
)

func main() {

	slice := []string{"Budi", "ipat", "isra"}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Tono"
	person["title"] = "Pekerja"
	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}