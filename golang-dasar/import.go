package main

import (
	"fmt"
	"golang-dasar/database"
	"golang-dasar/helper"
	_ "golang-dasar/internal" // Blank Identifier
)

func main() {
	result := helper.SayHello("Eko")
	fmt.Println(result)

	fmt.Println(database.GetDatabase())
}
