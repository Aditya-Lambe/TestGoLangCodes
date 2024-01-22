package main

import (
	"fmt"
)

func main() {
	p := person{
		Name: "Rahul",
		Age:  "26",
		address: address{
			City:  "Mumbai",
			State: "Maharashtra",
		},
	}
	fmt.Println(p.Name, p.City, p.address.State)
}

type address struct {
	City  string
	State string
}

type person struct {
	Name string
	Age  string
	address
}
