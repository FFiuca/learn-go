package main

import (
	"fmt"
)

func sayHello(a string) string {
	return "Hello " + a
}

func getFullName() (string, string) {
	return "caco", "cici"
}

func declaredReturnVal() (a, b string, c int) {
	a = "a"
	b = "b"
	c = 1
	return a, b, c
}

func variadicParam(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func funcParam(name string, getName func(string) string) string {
	return "Hai " + getName(name)
}

func getName(name string) string {
	return name
}

type Name func(string) string

func funcParam2(name string, getName Name) string {
	return "Hai " + getName(name)
}

// defer function : func yang akan dieksekusi di akhir function dan tetap berjalan walau func pemanggilnya error
func log() {
	fmt.Println("Logging run")

	// to get panic message
	message := recover()
	fmt.Println("message", message)
}

func exec() {
	defer log()

	fmt.Println("run exec")
}

// panic : menghentikan eksekuis program
func pnc() {
	defer log()
	fmt.Println("Run Panic")
	if true {
		panic("Error")
	}
	fmt.Println("End Panic")
}

func main() {
	fmt.Println(sayHello("a"))
	fmt.Println(getFullName())

	fName, lName := getFullName()

	fName2, _ := getFullName() // _ to ignore indexed value
	a, b, c := declaredReturnVal()

	numbers := []int{10, 2, 3}
	fmt.Println(fName, lName, fName2, a, b, c, variadicParam(1, 2, 3), variadicParam(numbers...))

	fullName := sayHello
	fmt.Println(fullName("a"), funcParam("caco", getName), funcParam2("caco", getName))

	// anon func
	fmt.Println(funcParam2("anon", func(name string) string {
		return name
	}))

	exec()

	pnc()
}
