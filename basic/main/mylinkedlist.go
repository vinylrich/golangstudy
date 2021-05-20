package main

import "fmt"

type Node struct {
	Next *Node
	val  int
}

//Add,Remove

func main() {
	var root *Node
	root = &Node{}

	root.val = 0
	tail := root //루트==tail
	node := root //tail을 계속 바꿔줘야함
	for i := 1; i <= 10; i++ {
		node.AddNode(tail, i) //tail을 다음 노드로 바꾸기
	}
	node = root
	for node.Next != nil {
		node.PrintNode()
	}
}

//root
func (n *Node) AddNode(tail *Node, val int) {
	for tail.Next != nil {
		tail = tail.Next
	}
	node := &Node{val: val}
	tail.Next = node
}

func (n *Node) PrintNode() {
	fmt.Println(n.val)
	n = n.Next
}
