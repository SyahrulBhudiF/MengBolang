package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	fmt.Println(parse.Local())

	duration1 := time.Second * 10
	fmt.Println(duration1)
}
