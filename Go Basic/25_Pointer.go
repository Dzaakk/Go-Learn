package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	 address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	 //pass by reference sehingga address2 menjadi pointer adress1
	 address2 := &address1
	 address2.City = "Bandung"
	//memakai bintang agar pada memory mana pun akan berubah sesuai data yang di sini
	 *address2 = Address{"Malang", "Jawa Timur", "Indonesia"}
	 
	 fmt.Println(address1)
	 fmt.Println(address2)
	//function new hanya mengembalikan pointer ke data kosongg
	var address3 *Address =new(Address)
	address3.City = "Jakarta"
	fmt.Println(address3)
}