package main

import "fmt"

// kayak interfacelah atau encap atau abstract tapi lebih ke variable

type Customer struct {
	name, address string
	status        bool
	dateJoin      string
}

func extractCustomer(cust Customer) {
	fmt.Println("extract", cust.name)
}

// struct method : menempel pada strucnya nanti pas di init
func (customer Customer) getStr() string {
	return customer.name + " " + customer.dateJoin
}

func main() {
	var cust Customer

	cust.name = "caco"
	cust.status = true

	cust2 := Customer{
		name:    "caco2",
		address: "aaa",
	}

	cust3 := Customer{"joko", "aaa", false, "2020"}

	fmt.Println(cust, cust2, cust3)
	extractCustomer(cust3)

	fmt.Println(cust3.getStr())
}
