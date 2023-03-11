package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "dzk",
		"address": "Bogor",
		"salah":   "salah",
	}

	//input key baru
	person["title"] = "programmer"
	delete(person, "salah")
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

}
