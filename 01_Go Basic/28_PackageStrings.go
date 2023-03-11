package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("budi subudi", "mamat"))
	fmt.Println(strings.Split("Budi Budiman", "yoyo"))

	fmt.Println(strings.Trim("Yaya sugiman", " "))
	fmt.Println(strings.ToTitle("menuju hari besar"))
	fmt.Println(strings.ReplaceAll("a a a a a a a", "a", "b"))
}