package main

import "fmt"

//Variadic func
func SumAll (numbers ...int) int{

	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
func main() {
	
	//Variadic Func
	total :=  SumAll(10, 20, 10, 20, 10, 200)
	fmt.Println(total)

	//slice parameter
	slice := []int{10,10,10,10}
	total = SumAll(slice...)
	fmt.Println(total)
}

