package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{2, 3, 1, 4}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "John"))
	fmt.Println(slices.Index(names, "Ringo"))
}
