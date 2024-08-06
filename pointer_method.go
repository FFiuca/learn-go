package main

import "fmt"

type Person struct {
	Name string
}

func (val *Person) TitleName() {
	val.Name = "Mr " + val.Name
}

func main() {
	a := Person{"Eko"}
	a.TitleName()

	fmt.Println(a)
}
