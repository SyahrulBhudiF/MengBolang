package main

import "strconv"

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		println(err)
	} else {
		println(result)
	}

	resultInt, err := strconv.ParseInt("42", 10, 64)
	if err != nil {
		println(err)
	} else {
		println(resultInt)
	}

	resultUint, err := strconv.ParseUint("42", 10, 64)
	if err != nil {
		println(err)
	} else {
		println(resultUint)
	}

	resultFloat, err := strconv.ParseFloat("42.42", 64)
	if err != nil {
		println(err)
	} else {
		println(resultFloat)
	}

	resultInt64, err := strconv.ParseInt("42", 10, 32)
	if err != nil {
		println(err)
	} else {
		println(resultInt64)
	}

	resultUint64, err := strconv.ParseUint("42", 10, 32)
	if err != nil {
		println(err)
	} else {
		println(resultUint64)
	}

	resultAtoi, err := strconv.Atoi("aku")
	if err != nil {
		println(err)
	} else {
		println(resultAtoi)
	}

	resultItoa := strconv.Itoa(42)
	println(resultItoa)

	binary := strconv.FormatInt(42, 2)
	println(binary)
}
