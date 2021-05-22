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

	//1,2,3,4,5,6,7,8,9,10
	for len(s) > 0 {
		var last *TreeNode
		last, s = s[len(s)-1], s[0:len(s)-1] //last가 2,3,4
		// last, s = s[0], s[1:] //last가 3,2,10,9
		//last=4
		fmt.Printf("%d->", last.Val) //1

		//len(last.Child)=4,3,2
		for i := len(last.Childs) - 1; i >= 0; i-- {
			fmt.Print(last.Childs[i].Val)
			s = append(s, last.Childs[i]) //인덱스 2부터
		} //s=4,3,2
	}
}

func (t *Tree) BFS() {
	queue := []*TreeNode{}
	queue = append(queue, t.Root)

	for len(queue) > 0 {
		var first *TreeNode
		first, queue = queue[0], queue[1:]
		fmt.Printf("%d->", first.Val) //찾기
		for i := 0; i < len(first.Childs); i++ {
			queue = append(queue, first.Childs[i])
		}
	}
}

//BFS DFS
//너비우선 깊이우선
//깊이우선은 맨 아래부터 탐색->부모->부모의다른자식->올라감
//너비우선은 루트->층별로 탐색(위부터)
