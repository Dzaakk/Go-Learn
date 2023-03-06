package main

import "fmt"

func random() interface{} {
	return "ups"
}
func RandomWithSwitch() interface{} {
	return 10
}

func main() {
	var result interface{} = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	var resultRandom interface{} = RandomWithSwitch()
	switch RandomReturn := resultRandom.(type) {
	case string:
		fmt.Println("Return is", RandomReturn, "with type string")
	case int:
		fmt.Println("Return is", RandomReturn, "with type int")
	default:
		fmt.Println("Unkonw Type")
	}
}