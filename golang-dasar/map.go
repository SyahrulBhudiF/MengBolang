package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "John Doe",
		"address": "St. Petersburg",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Println(len(person))
	person["name"] = "John Wick"
	fmt.Println(person)
	delete(person, "name")
	fmt.Println(person)
}
