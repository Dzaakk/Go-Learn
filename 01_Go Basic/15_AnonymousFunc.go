package main

import "fmt"

type Blacklist func(string) bool

func registeruser (name string, blacklist Blacklist){
	if blacklist(name) {
		fmt.Println("Blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}
func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registeruser("admin", blacklist)
	registeruser("jojo", blacklist)


}

