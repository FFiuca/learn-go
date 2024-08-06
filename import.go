package main

import (
	"fmt"
	"learn/database"
	"learn/helper"
	_ "learn/internal" // _ this is to we want import and run init in other package due needed but not implement any function next
)

func main() {
	a := helper.SayHello("caco")
	fmt.Println(a)

	fmt.Println(helper.Application)
	// b := Dog{Name: "cupit"} // will call function in interface  // still error
	// fmt.Println(Run(b))

	fmt.Println(database.GetConnection())
}
