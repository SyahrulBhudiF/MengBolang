package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "a string")
	port := flag.Int("port", 8080, "an int")
	debug := flag.Bool("debug", false, "a bool")

	flag.Parse()

	fmt.Println(*host, *port, *debug)
}
