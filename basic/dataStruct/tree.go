package dataStruct

import "fmt"

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *TreeNode) AddNode(val int) {
	t.Childs = append(t.Childs, &TreeNode{Val: val})
}
func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}
func (t *Tree) DFS1() {
	DFS1(t.Root)
}
func DFS1(Node *TreeNode) {
	fmt.Printf("%d->", Node.Val)

	for i := 0; i < len(Node.Childs); i++ {
		DFS1(Node.Childs[i])
	}
}

func (t *Tree) DFS2() {
	s := []*TreeNode{}
	s = append(s, t.Root)

	for len(s) > 0 {
		var last *TreeNode
		last, s = s[len(s)-1], s[0:len(s)-1]

		fmt.Printf("%d->", last.Val)
		for i := len(last.Childs) - 1; i >= 0; i-- {
			s = append(s, last.Childs[i])
		}
	}
}

//BFS DFS
//너비우선 깊이우선
//깊이우선은 맨 아래부터 탐색->부모->부모의다른자식->올라감
//너비우선은 루트->층별로 탐색(위부터)
