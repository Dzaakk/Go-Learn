package main

import (
	"fmt"
	"strconv"
)
//string to boolean
func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}
//string to number
	number, err := strconv.ParseInt("100000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	valueInt, err := strconv.Atoi("2000000") //langsung tanpa harus mendeskripsikan jumlah bytenya
	fmt.Println(valueInt)
	
//integer to string
value := strconv.FormatInt(10000000, 2)
fmt.Println(value) //print jadi biner

}