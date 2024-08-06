package main

import "fmt"

func main() {
	var name string = "hh"
	var test = "testt" // jika di init langsung tdk wajib declare type data
	name2 := "short"   // short syntax var, tapi harus init

	// multiple
	var (
		fName = "cc"
		lName = "haha"
	)

	const uName = "caco"
	const (
		fName2 = "cc2"
		lName2 = "haha2"
	)

	fmt.Println(name, test, name2, fName, lName, uName, fName2)

	name = "yuhu"
	test = "jj"
	name2 = "sh2"
	fmt.Println(name, test, name2, fName, lName)

	// convert
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	nilai16 := int16(nilai64) // bakal kembali ke mulai bawah

	var float32 float32 = float32(nilai64)

	fmt.Println(nilai32, nilai64, nilai16, float32)

	e := name[0]
	e2 := string(e)

	fmt.Println(e, e2)

	// type / alias

	type fieldString string

	var name3 fieldString = "haha"
	var name4 string = "haha2"

	name5 := fieldString(name4)

	fmt.Println(name3, name4, name5)
}
