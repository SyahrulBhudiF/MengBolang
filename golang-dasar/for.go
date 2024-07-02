package main

func main() {
	// For
	for i := 0; i < 5; i++ {
		println("Angka", i)
	}

	// For with condition
	var i = 0
	for i < 5 {
		println("Angka", i)
		i++
	}

	// For without argument
	var j = 0
	for {
		println("Angka", j)
		j++
		if j == 5 {
			break
		}
	}
}
