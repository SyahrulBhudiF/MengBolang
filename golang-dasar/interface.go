package main

import "fmt"

type HashName interface {
	GetName() string
}

func SayHello(value HashName) {
	fmt.Println("Hello", value.GetName())
}

func main() {
	var eko Person
	eko.Name = "Eko"

	SayHello(eko)

	var cat Animal
	cat.Name = "Kucing"

	SayHello(cat)
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (person Person) GetName() string {
	return person.Name
}
