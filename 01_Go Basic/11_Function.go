package main

import "fmt"

func Namefunc(a string) string {
	return "Hello" + a
}
func getFullName() (string, string) {
	return "Budi", "Bob"
}
func getCompleteName()(FrName, MdName, LaName string){
	FrName = "tato"
	MdName = "tuta"
	LaName = "tuti"

	return
}
func main() {
	fname := Namefunc("lin")
	fmt.Println(fname)

	firstName, _ := getFullName()
	fmt.Println(firstName)
	//fmt.Println(LastName)

	b,c,d:= getCompleteName()
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
