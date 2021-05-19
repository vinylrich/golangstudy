package dataStruct

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}

//BFS DFS
//너비우선 깊이우선
//깊이우선은 맨 아래부터 탐색->부모->부모의다른자식->올라감
//너비우선은 루트->층별로 탐색(위부터)
