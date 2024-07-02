package main

import "fmt"

func main() {
	var names [4]string

	names[0] = "John"
	names[1] = "Wick"
	names[2] = "Ethan"
	names[3] = "Hunt"

	for i := 0; i < len(names); i++ {
		println(names[i])
	}

	// Array Initialization
	var fruits = [...]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println(fruits)
	fmt.Println(len(fruits))

}
