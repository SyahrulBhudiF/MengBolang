package main

func main() {
	// If
	var name1 = "John"
	var name2 = "Doe"

	if name1 == "John" {
		if name2 == "Doe" {
			println("Hello John Doe")
		} else {
			println("Hello John")
		}
	} else {
		println("Hello Doe")
	}

	// If Else
	if name1 == "John" && name2 == "Doe" {
		println("Hello John Doe")
	} else if name1 == "John" {
		println("Hello John")
	} else {
		println("Hello Doe")
	}

	// If Else If
	if name1 == "John" && name2 == "Doe" {
		println("Hello John Doe")
	} else if name1 == "John" {
		println("Hello John")
	} else if name2 == "Doe" {
		println("Hello Doe")
	} else {
		println("Hello")
	}

	// If Short Statement
	if length := len(name1); length > 5 {
		println("Terlalu panjang")
	} else {
		println("Nama sudah benar")
	}
}
