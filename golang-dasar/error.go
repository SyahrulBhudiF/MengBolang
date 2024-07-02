package main

import "errors"

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	result, err := Pembagian(100, 0)
	if err == nil {
		println("Hasil", result)
	} else {
		println("Error", err.Error())
	}
}
