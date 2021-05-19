package main

import (
	"fmt"

	"dataStruct"
)

func main() {
	list := &dataStruct.LinkedList{}
	list.AddNode(0)
	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}
	list.PrintNodes()

	list.RemoveNode(list.Root.Next)

	list.PrintNodes()

	list.RemoveNode(list.root)
	list.PrintNodes()

	list.PrintReverse()

	a := []int{1, 2, 3, 4, 5}
	a = append(a[0:2], a[3:]...)
	fmt.Println(a)
}
