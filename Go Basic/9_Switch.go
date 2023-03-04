package main

import "fmt"

func main() {
	name := "nam"
	switch length := len(name); length < 3 {
	case true:
		fmt.Println("Nama terlalu pendek")
	case false:
		fmt.Println("Nama sudah benar")
	}

	lengthName := len(name)
	switch {
	case lengthName > 20:
		fmt.Println("nama terlalu panjang")
	case lengthName > 5:
		fmt.Println("panjang karakter nama sudah terpenuhi")
	default:
		fmt.Println("nama terlalu pendek")
	}

}
