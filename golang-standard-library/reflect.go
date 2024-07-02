package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func main() {
	sample := Sample{"John"}
	sampleType := reflect.TypeOf(sample)
	field := sampleType.Field(0)
	max := field.Tag.Get("max")

	fmt.Println(IsValid(sample))

	fmt.Println(max)
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}

	return true
}
