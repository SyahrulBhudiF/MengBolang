package main

import "fmt"

func main() {
	sayHello("syahrul")
	fmt.Println(aritmatika(10, 5))
	data1, data2 := multiple(10, 5)
	fmt.Println(data1, data2)

	dt1, _ := multiple(10, 5)
	fmt.Println(dt1)

	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)

	total := sumALl(10, 10, 10, 10, 10)
	fmt.Println(total)

	numbers := []int{10, 10, 10, 10, 10}
	total = sumALl(numbers...)
	fmt.Println(total)

	original := changeValue
	fmt.Println(original(5))

	filter := spamFilter
	fmt.Println(sayHelloWithFilter("anjing", filter))

	anonymous()

	fmt.Println(factorial(5))
}

// Function with parameter
func sayHello(name string) {
	fmt.Println("Hello, " + name)
}

// Function with return value
func aritmatika(a int16, b int16) int16 {
	return (a + b)
}

// Function with multiple return value
func multiple(a, b int) (int, int) {
	return a + b, a * b
}

// Function with named return value
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Syahrul"
	middleName = "Ridho"
	lastName = "Hilman"
	return
}

// Function with variadic parameter
func sumALl(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Function as a value
func changeValue(original int) int {
	return original
}

// Function as a Parameter
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) string {
	return "Hello, " + filter(name)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

// Function Anonymous
func anonymous() {
	goodbye := func(name string) {
		fmt.Println("Goodbye, " + name)
	}

	goodbye("Syahrul")
}

// Function Recursive
func factorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorial(value-1)
	}
}
