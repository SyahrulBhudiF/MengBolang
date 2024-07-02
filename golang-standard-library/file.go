package main

import (
	"bufio"
	"io"
	"os"
)

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func main() {
	//err := createNewFile("message.log", "Hello, World!")
	//if err != nil {
	//	println("Error creating file")
	//	return
	//}
	//println("File created")
	//result, _ := readFile("message.log")
	//fmt.Println(result)
	addToFile("message.log", "\nHello, World!")
}
