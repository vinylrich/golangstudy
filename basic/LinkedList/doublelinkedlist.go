package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	val  int
}
type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{val: val}
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		l.root.prev = nil
		if l.root == nil {
			l.tail = nil
		}
		return
	}
	prev := node.prev

	if node == l.tail {
		prev.next = nil
		l.tail.prev = nil
		l.tail = prev
	} else {
		prev.next = prev.next.next
	}
}
func (l *LinkedList) PrintNodes() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Println()
}

func (l *LinkedList) PrintReverse() {
	fmt.Println("접근")
	node := l.tail
	for node.prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.prev
	}
	fmt.Println()
}

//1번 방법(계속 순차적으로 접근)
// func main() {
// 	var root *Node

// 	root = &Node{val: 0}
// 	for i := 1; i < 10; i++ {
// 		AddNode(root, i)
// 	}
// 	node := root//0
// 	for node.next != nil {
// 		fmt.Printf("%d->", node.val)
// 		node = node.next
// 	}
// 	fmt.Printf("%d\n", node.val)
// }

// func AddNode(root *Node, val int) {
// 	var tail *Node
// 	tail = root
// 	for tail.next != nil {
// 		tail = tail.next
// 	}
// 	node := &Node{val: val}
// 	tail.next = node
// }

//2번 방법 tail을 따로 저장해놔서 하나만 접근
func main() {
	list := &LinkedList{}
	list.AddNode(0)
	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}
	list.PrintNodes()

	list.RemoveNode(list.root.next)

	list.PrintNodes()

	list.RemoveNode(list.root)
	list.PrintNodes()

	list.PrintReverse()

	a := []int{1, 2, 3, 4, 5}
	a = append(a[0:2], a[3:]...)
	fmt.Println(a)
}

//0,1
