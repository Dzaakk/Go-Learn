package main

import "fmt"
func endApp(){
	fmt.Println("End App")
	message:= recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	
}
func runApp(error bool){
	defer endApp()
	if error {
		panic("ERROR")
	}
	fmt.Println("Aplikasi berjalan lancar")
}
func main() {

	runApp(true)
}