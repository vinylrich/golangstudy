package dataStruct

import "fmt"

//이진 탐색 트리(BST:Binary Search Tree)
type BinTreeNode struct {
	Val int

	Left  *BinTreeNode //left<=parent
	Right *BinTreeNode //right>=parent
}

type BinTree struct {
	Root *BinTreeNode
}

func NewBinaryTree(v int) *BinTree {
	tree := &BinTree{}
	tree.Root = &BinTreeNode{Val: v}
	return tree
}
func (n *BinTreeNode) AddNode(v int) *BinTreeNode {
	if n.Val > v {
		if n.Left == nil {
			n.Left = &BinTreeNode{Val: v}
			return n.Left
		} else {
			return n.Left.AddNode(v)
		}
	} else {
		if n.Right == nil {
			n.Right = &BinTreeNode{Val: v}
			return n.Right
		} else {
			return n.Right.AddNode(v)
		}
	}
}

type depthNode struct {
	depth int
	node  *BinTreeNode
}

func (t *BinTree) PrintTree() {
	q := []depthNode{}
	q = append(q, depthNode{depth: 0, node: t.Root})
	currentDepth := 0

	for len(q) > 0 {
		var first depthNode
		first, q = q[0], q[1:]

		if first.depth != currentDepth {
			fmt.Println()
			currentDepth = first.depth
		}
		fmt.Print(first.node.Val, " ")

		if first.node.Left != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.Left})
		}
		if first.node.Right != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.Right})
		}
	}
}

func (t *BinTree) Search(v int) (bool, int) {
	return t.Root.Search(v, 1)
}

func (n *BinTreeNode) Search(v int, cnt int) (bool, int) {
	if n.Val == v {
		return true, cnt
	} else if n.Val > v {
		if n.Left != nil {
			return n.Left.Search(v, cnt+1)
		}
		return false, cnt
	} else {
		if n.Right != nil {
			return n.Right.Search(v, cnt+1)
		}
		return false, cnt
	}

}
