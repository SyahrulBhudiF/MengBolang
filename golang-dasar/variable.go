package main

import "fmt"

func main() {
	name := "John Doe"
	fmt.Println(name)

	name = "John Wick"
	fmt.Println(name)

	// Variable multiple can be changed
	var (
		firstName = "John"
		lastName  = "Doe"
	)
	fmt.Println(firstName, lastName)

	// Variable constant, cant be changed
	const (
		age  = 22
		city = "Jakarta"
	)
	fmt.Println(age)
	fmt.Println(city)
}
