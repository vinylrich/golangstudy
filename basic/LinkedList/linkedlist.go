package main

import "fmt"



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
}

//0,1
