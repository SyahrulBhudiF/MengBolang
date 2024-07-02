package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("Hello World!\napakah")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello ")
	_, _ = writer.WriteString("World!\n")
	_, _ = writer.WriteString("Apakah ")
	_, _ = writer.WriteString("Kabar?\n")
	writer.Flush()
}
