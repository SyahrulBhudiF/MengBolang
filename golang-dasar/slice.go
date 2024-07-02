package main

import "fmt"

func main() {
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daysSlice1 := days[5:]
	fmt.Println(daysSlice1)

	daysSlice1[0] = "Sabtu Lagi"
	fmt.Println(days)

	daysSlice2 := append(daysSlice1, "Libur")
	fmt.Println(daysSlice2)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "John"
	newSlice[1] = "Doe"
	newSlice2 := append(newSlice, "Wick")

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(newSlice2)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
	fmt.Println(fromSlice)
}
