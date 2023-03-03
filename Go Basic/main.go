package main

import "fmt"

func main() {

	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	var slice1 = months[2:4]
	// slice1[0] = "bukan april"
	// months[4] = "bukan mei"
	// fmt.Println(slice1)
	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	var slice2 = append(slice1, "bulan1")
	fmt.Println(slice2)
	slice2[1] = "bulanbaru"
	// fmt.Println(slice2)
	// fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "hay"
	newSlice[1] = "mike"
	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)

}
