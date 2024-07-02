package main

func random() interface{} {
	return 1
}

func main() {
	result := random()
	switch value := result.(type) {
	case string:
		println("Value ", value, " is string")
	case int:
		println("Value ", value, " is int")
	default:
		println("Unknown type")
	}
}
