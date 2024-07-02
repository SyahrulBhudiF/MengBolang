package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack(1)
	data.PushBack(2)
	data.PushBack(3)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("")

	data.Remove(data.Front())

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("")

	data.PushFront(0)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("")

	data.Remove(data.Back())

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("")

	data.PushBack(3)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("")

	data.PushBack(4)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("")

	data.PushFront(-1)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	head := data.Front()

	fmt.Println(head.Value, head.Next().Value, head.Next().Next().Value, head.Next().Next().Next().Value, head.Next().Next().Next().Next().Value, head.Next().Next().Next().Next().Next().Value, head.Next().Next().Next().Next().Next().Next().Value, head.Next().Next().Next().Next().Next().Next().Next().Value, head.Next().Next().Next().Next().Next().Next().Next().Next().Value, head.Next().Next().Next().Next().Next().Next().Next().Next().Next().Value, head.Next().Next().Next().Next().Next().Next().Next().Next().Next().Next().Value)
}
