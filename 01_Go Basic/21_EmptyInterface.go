package main

import "fmt"
/* fungsi dibawah menggunakan parameter interface untuk 
menampung data apapun sebagai parameter
*/

// func Ups(apapun interface{}) interface{}{
	
// }

//return apapun karena interface bisa menerima apapun data
func ups(i int) interface{}{
	if i == 1 {
		return 1
	} else if i ==2{
		return true
	} else {
		return "ups"
	}
}
func main() {
	var data interface{} = ups(3)
	fmt.Println(data)	
}