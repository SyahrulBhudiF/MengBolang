package main

import (
	"fmt"
	"math"
)

func main() {
	round := math.Round(1.5)
	fmt.Println(round)

	fmt.Println(math.Ceil(1.2))
	fmt.Println(math.Floor(1.8))
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Sqrt(16))
}
