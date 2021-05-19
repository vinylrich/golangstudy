package dataStruct

import "fmt"

type Node struct {
	Next *Node
	Prev *Node
	Val  int
}
type LinkedList struct {
	Root *Node
	Tail *Node
}

func (l *LinkedList) AddNode(Val int) {
	if l.Root == nil {
		l.Root = &Node{Val: Val}
		l.Tail = l.Root
		return
	}
	l.Tail.Next = &Node{Val: Val}
	Prev := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.Prev = Prev
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.Root {
		l.Root = l.Root.Next
		l.Root.Prev = nil
		if l.Root == nil {
			l.Tail = nil
		}
		return
	}
	Prev := node.Prev

	if node == l.Tail {
		Prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = Prev
	} else {
		Prev.Next = Prev.Next.Next
	}
}
func (l *LinkedList) PrintNodes() {
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func (l *LinkedList) PrintReverse() {
	fmt.Println("접근")
	node := l.Tail
	for node.Prev != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Prev
	}
	fmt.Println()
}
