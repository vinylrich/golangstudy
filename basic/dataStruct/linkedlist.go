package dataStruct

import "fmt"

type Node struct {
	next *Node
	prev *Node
	val  int
}
type LinkedList struct {
	Root *Node
	Tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.Root == nil {
		l.Root = &Node{val: val}
		l.Tail = l.Root
		return
	}
	l.Tail.next = &Node{val: val}
	prev := l.Tail
	l.Tail = l.Tail.next
	l.Tail.prev = prev
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.Root {
		l.Root = l.Root.next
		l.Root.prev = nil
		if l.Root == nil {
			l.Tail = nil
		}
		return
	}
	prev := node.prev

	if node == l.Tail {
		prev.next = nil
		l.Tail.prev = nil
		l.Tail = prev
	} else {
		prev.next = prev.next.next
	}
}
func (l *LinkedList) PrintNodes() {
	node := l.Root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Println()
}

func (l *LinkedList) PrintReverse() {
	fmt.Println("접근")
	node := l.Tail
	for node.prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.prev
	}
	fmt.Println()
}
