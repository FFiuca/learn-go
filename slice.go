package main

import "fmt"

// kayak array[:] di python tapi slice hanya pointer bukan copy dari array asli
func main() {
	val := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	slice := val[0:3]
	slice2 := val[0:]
	slice3 := val[:4]
	var slice4 []int = val[:] // longhand slice declaration, equal var slice4 = make([]int, 8,8), var slice4 = []int{1,2,3,}
	fmt.Println(slice, slice2, slice3, slice4, cap(slice3), cap(slice))

	str := [...]string{"a", "b", "c", "d", "e", "f", "g"}

	strA := str[4:]
	strA[0] = "eNew"
	strB := append(strA, "h") // ini baru mengcopy karena sizenya/cap melebihi
	strB[0] = "eNew2"
	fmt.Println(str, strB)

	// slice pure syntax
	var slicePure = make(
		[]string, // tipe data
		2,        // panjang
		5,        // sizenya, jika nanti ini diappend, maka tidak akan buat array baru jika belum > 5
	)

	slicePure[0] = "a"
	slicePure[1] = "b"
	slicePure2 := append(slicePure, "c")

	fmt.Println(slicePure, slicePure2)

	slicePure2[0] = "aNew" // ikut mengubah nilai orinya
	fmt.Println(slicePure, slicePure2)

}
