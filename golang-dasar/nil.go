package main

func main() {
	// Nil
	var person map[string]string
	if person == nil {
		println("Data kosong")
	}

	// Nil with function
	var person2 = NewMap("")
	if person2 == nil {
		println("Data kosong")
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
