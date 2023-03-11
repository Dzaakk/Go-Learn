package main

import (
	"errors"
	"fmt"
)

func Pembagian(angka int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := angka / pembagi
		return result, nil
	}
}
func main() {
	hasil, err := Pembagian(100, 0)

	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}

