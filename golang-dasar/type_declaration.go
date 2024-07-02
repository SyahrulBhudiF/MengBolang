package main

import "fmt"

func main() {
	// Type Declaration
	type NoKTP string
	type Married bool

	var noKTP NoKTP = "1234567890"
	var marriedStatus Married = true

	fmt.Println(noKTP)
	fmt.Println(marriedStatus)
}
