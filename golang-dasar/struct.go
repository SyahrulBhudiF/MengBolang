package main

import "fmt"

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Indonesia"
	eko.Age = 30

	fmt.Println(eko)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     35,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 40}
	fmt.Println(budi)

	eko.sayHello("Budi")
	joko.sayHello("Eko")
	budi.sayHello("Joko")
}

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}
