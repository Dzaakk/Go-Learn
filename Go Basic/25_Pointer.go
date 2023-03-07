package main

import "fmt"

type Address struct {
	City, Province, Country string
}
//tambahkan * agar jadi pointer pada suatu parameter 
func ChangeCountry(adress *Address){
	adress.Country = "Indonesia"
}
type Man struct{
	Name string
}
//contoh Pointer di Method/ struct func
//gunakan operator * untuk menjadikan struct sebagai pointer
func(man *Man) Married(){
	man.Name = "Mr." +man.Name
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

	var alamat = Address{
		City: "Subang",
		Province: "Jawa Barat",
		Country: "",
		}	 
		//gunakan & agar jadi pointer
		ChangeCountry(&alamat)
		fmt.Println(alamat)

	toni := Man{"Toni"}
	toni.Married()
	fmt.Println(toni)
}