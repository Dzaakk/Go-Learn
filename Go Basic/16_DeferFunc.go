package main

import "fmt"
func logging(){
	fmt.Println("Selesai memanggil logging")
}
func runApp(){
	defer logging()
	fmt.Println("Run App")
}
func main() {

	runApp()
}