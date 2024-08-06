package main

import (
	"fmt"
)

func main() {
	data := map[string]string{
		"name":    "caco",
		"address": "haha",
	}

	data["yuhu"] = "yoho"
	delete(data, "address")
	fmt.Println(data)

	switch aa := len("ha"); {
	case aa == 2:
		fmt.Println("true")
	}

	for counter := 0; counter < 10; counter++ {
		fmt.Println(counter)
	}

	data2 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(data2); i++ {
		fmt.Println(data2[i])
	}

	// for map
	for idx, val := range data {
		fmt.Println(idx, val)
	}
}
