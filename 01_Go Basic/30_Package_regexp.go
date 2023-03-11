package main

import (
	"fmt"
	"regexp"
)

//validasi email dll
func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])o")
	fmt.Println(regex.MatchString("aho"))
	fmt.Println(regex.MatchString("ato"))
	fmt.Println(regex.MatchString("aJo"))

	search := regex.FindAllString("ako ajo ayo aio", -1)
	fmt.Println(search)
	// -1 untuk mendapatkan semuanya
}