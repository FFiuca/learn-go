package main

import "fmt"

type Dump struct {
	City, Province, Country string
}

func ChangeAddressNotPointer(dump Dump) {
	dump.City = "change"
}

func ChangeAddressPointer(dump *Dump) {
	dump.City = "change pointer"
}

func main() {
	a := Dump{"Jombang", "Jatim", "Indo"}
	b := a

	b.City = "Mojokerto" // tidak merubah a
	fmt.Println(a, b)

	var c *Dump = &a // & indicate pointer
	c.City = "Lamongan"

	fmt.Println(a, c)
	fmt.Println("----")

	// asterik(*) : mengubah root pointer jadi variable baru
	address1 := Dump{"sana", "sini", "sono"}
	address2 := &address1

	fmt.Println(address1, address2)

	*address2 = Dump{"sana2", "sini2", "sono2"} // address1 berubah ke address2, dan pointer root sekarang berubah ke address2
	fmt.Println(address1, address2)

	address2.City = "sini3"
	fmt.Println(address1, address2)

	fmt.Println("----")
	address3 := new(Dump) // membuat struct dengan data kosong dan struct ini aslinya hanya pointer ke Dump => var address3 *Dump = &Dump{}
	address3.City = "address3"
	fmt.Println(address3, address2)

	// Pointer function, untuk mengubah variable di luar function, karena sifat default param function adalah pass by value(copy origin value)
	fmt.Println("----")
	address4 := Dump{}
	ChangeAddressNotPointer(address4)
	fmt.Println(address4)
	ChangeAddressPointer(&address4) // berguna sebagai closure
	fmt.Println(address4)
}
