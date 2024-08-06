package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	result := random()
	var assertionString string
	var assertionInt int

	switch result.(type) {
	case string:
		assertionString = result.(string)
	case int:
		assertionInt = result.(int)
	}

	fmt.Println(assertionString, assertionInt)

}
