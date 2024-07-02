package main

import "fmt"

func main() {
	name := "John Doe"

	switch name {
	case "John Doe":
		fmt.Println("Nama saya John Doe")
	case "John Wick":
		fmt.Println("Nama saya John Wick")
	default:
		fmt.Println("Nama saya tidak diketahui")
	}
}
