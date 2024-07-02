package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello, World!", "World"))
	fmt.Println(strings.Split("apalah", ""))
	fmt.Println(strings.Join([]string{"Hello", "World!"}, ", "))
	fmt.Println(strings.Replace("Hello, World!", "World", "Golang", 1))
	fmt.Println(strings.ToLower("Hello, World!"))
	fmt.Println(strings.ToUpper("Hello, World!"))
	fmt.Println(strings.TrimSpace(" Hello, World! "))
	fmt.Println(strings.Trim(" Hello, World! ", " "))
	fmt.Println(strings.TrimLeft(" Hello, World! ", " "))
	fmt.Println(strings.TrimRight(" Hello, World! ", " "))
	fmt.Println(strings.HasPrefix("Hello, World!", "Hello"))
	fmt.Println(strings.HasSuffix("Hello, World!", "World!"))
	fmt.Println(strings.Index("Hello, World!", "World"))
	fmt.Println(strings.LastIndex("Hello, World!", "o"))
	fmt.Println(strings.Count("Hello, World!", "o"))
	fmt.Println(strings.Repeat("Hello, World! ", 3))
	fmt.Println(strings.Compare("Hello, World!", "Hello, World!"))
	fmt.Println(strings.EqualFold("Hello, World!", "hello, world!"))
	fmt.Println(strings.Fields("Hello, World!"))
}
