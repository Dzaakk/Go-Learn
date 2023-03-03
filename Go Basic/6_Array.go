package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "Toto"
	names[1] = "bambang"
	names[2] = "tuti"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		98, 70, 100,
	}
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

}
