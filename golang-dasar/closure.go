package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		fmt.Println("Increment ", counter, " executed")
		counter++
	}

	for i := 10; i >= 0; i-- {
		increment()
	}
	fmt.Println(counter)
}
