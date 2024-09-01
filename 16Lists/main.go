package main

import (
	"container/list"
	"fmt"
)

func main() {
	var newList list.List
	newList.PushBack(123)
	newList.PushBack(564)
	newList.PushBack(385)
	newList.PushBack(206)
	newList.PushBack(837)
	newList.PushBack(958)

	for element := newList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
