package main

import "fmt"

func main() {
	runApp(false)
	runApp(true)
}

// Panic
func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	if error {
		defer endApp()
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}
