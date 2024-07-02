package main

import "fmt"

func main() {
	runn(false)
	runn(true)
}

// Panic
func end() {
	fmt.Println("Aplikasi selesai")
	message := recover()
	fmt.Println("Pesan error: ", message)
}

func runn(error bool) {
	if error {
		defer end()
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}
