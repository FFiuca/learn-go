package main

import "fmt"

// di golang tidak ada class hmm makanya bingung awal. diganti struct sama interface. berarti gk ada inheritence
// dan disini polimorfismenya agak beda, jadi langsung level paling dibawah OOP biasnya gk kayak Python, misal Animal langsung ke Dog, jadi gak ada typenya
type Animal struct {
	Name, Type, LatinName, Sex string
	IsAmphibie, IsHasAge       bool
}

// ini mah gak kepake yang animal langsung pake dog.
// intinya gini. polimorphisme hanya digunakan di method karena variable harus dideclare tiap struct yang mana meskipun bisa sama tapi ya memang harus di diclare, aneh" aja :)

type Behaviour interface {
	Speak() string
	Go() string
	Breath() string // capitalize indicate is public accessor, lower-first indicete private
}

type Dog struct {
	Name, Type, LatinName, Sex string
	IsAmphibie, IsHasAge       bool
}

func (a Dog) Speak() string {
	return "I am " + a.Name + " and I am menggonggong"
}

func (a Dog) Go() string {
	return "I am walking not swiming"
}

func (a Dog) Breath() string {
	return "I am breathing with paru-paru"
}

type Fish struct {
	Name, Type, LatinName, Sex string
	IsAmphibie, IsHasAge       bool
}

// disini pengecekan kontrak terjadi
func Run(a Behaviour) {
	fmt.Println(a.Speak())
}

// inetrface kosng untuk any
func Kosong1() interface{} {
	return 1
}

func Kosong2() any {
	return "bukan OOP"
}

func main() {
	a := Dog{Name: "Cupit"}
	Run(a)

	// f := Fish{Name: "arwana"}
	// Run(f)

	b, c := Kosong1(), Kosong2()
	fmt.Println(b, c)
}
