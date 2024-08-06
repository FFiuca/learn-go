package main

import "fmt"

func main() {
	a := 10
	b := 10
	c := 5
	d := a + b*c
	a /= 3
	var f float32 = float32(a) / 5

	fmt.Println(d, a, f)

	f++
	fmt.Println(f, f == 1.6, f != 1.6, "a" > "b", "a" < "b")
}
