package main

import "errors"

var (
	ValidationError = errors.New("Validation error")
	NotFoundError   = errors.New("Not found error")
)

func main() {
	err := GetById("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			println("Not found error")
		} else {
			println("Unknown error")
		}
	}
}

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "1" {
		return NotFoundError
	}

	return nil
}
