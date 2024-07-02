package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("a(x*)b")

	fmt.Println(regex.MatchString("ab"))
	fmt.Println(regex.MatchString("axb"))
	fmt.Println(regex.MatchString("axxb"))

}
