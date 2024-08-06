package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "cac"
	name[1] = "2"
	name[2] = "3"

	var val = [3]int{
		1, 2, 3,
	}

	var val2 [3]int   // default 0
	var str [3]string // default ''

	var dynamic = [...]int{
		1, 2, 3, 4, // akan menghitung ini
	}

	fmt.Println(name, val, val2, str, len(val2), dynamic, len(dynamic))

}
