package main

import "fmt"

func NewMap(name string) any {
	if name == "" {
		return nil
	}

	return name
}

func main() {
	fmt.Println(NewMap("halo"), NewMap(""))
}
