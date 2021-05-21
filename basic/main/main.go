package main

import (
	"fmt"
	"golangstudy/basic/dataStruct"
)

func main() {
	tree := dataStruct.NewBinaryTree(5)
	tree.Root.AddNode(3)
	tree.Root.AddNode(2)
	tree.Root.AddNode(4)
	tree.Root.AddNode(8)
	tree.Root.AddNode(7)
	tree.Root.AddNode(6)
	tree.Root.AddNode(10)
	tree.Root.AddNode(9)

	tree.PrintTree()
	val := 6
	if found, cnt := tree.Search(val); found {
		fmt.Println("found ", val, "count:", cnt)
	} else {
		fmt.Println("Not Found", val, "count:", cnt)
	}
}
