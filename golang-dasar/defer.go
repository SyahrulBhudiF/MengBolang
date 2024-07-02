package main

import "fmt"

func main() {
	run()
}

func logging() {
	fmt.Println("Logging")
}

func run() {
	defer logging()
	fmt.Println("Run application")
}
