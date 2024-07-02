package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var alamat1 *Address = new(Address)
	alamat1.City = "Subang"

	var alamat2 *Address = alamat1

	alamat2.City = "Bandung"

	fmt.Println(alamat1.City)

}
