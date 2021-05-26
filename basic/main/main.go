package main

import (
	"golangstudy/basic/dataStruct"
)

//기본(알고리즘)-> 개발(웹)-로직,처리 등 실무코딩->상용화
// func main() {
// 	tree := dataStruct.Tree{}
// 	val := 1

// 	tree.AddNode(val)
// 	val++

// 	for i := 0; i < 3; i++ {
// 		tree.Root.AddNode(val)
// 		val++
// 	}
// 	for i := 0; i < len(tree.Root.Childs); i++ {
// 		for j := 0; j < 2; j++ {
// 			tree.Root.Childs[i].AddNode(val)
// 			val++
// 		}
// 	}
// 	tree.DFS2()
// }

// func main() {
// 	heap1 := &dataStruct.Heap{}

// 	nums1 := []int{-1, 3, -1, 5, 4}
// 	for i := 0; i < len(nums1); i++ {
// 		heap1.Push(nums1[i])
// 		if heap1.Count() > 2 {
// 			heap1.Pop()
// 		}
// 	}

// 	fmt.Println(heap1.Pop())

// 	heap2 := &dataStruct.Heap{}
// 	nums2 := []int{2, 8, -2, -3, 4}
// 	for i := 0; i < len(nums2); i++ {
// 		heap2.Push(nums2[i])
// 		if heap2.Count() > 1 {
// 			heap2.Pop()
// 		}
// 	}

// 	fmt.Println(heap2.Pop())

// }

// func main() {
// 	m := dataStruct.CreateMap()
// 	m.Add("aaa", "010123123")
// 	m.Add("bbb", "12321414")
// 	m.Add("ccc", "213531231")
// 	m.Add("sdasd", "567876532")

// 	fmt.Println("aaa=", m.Get("aaa"))
// 	fmt.Println("bbb=", m.Get("bbb"))
// 	fmt.Println("ccc=", m.Get("ccc"))
// 	fmt.Println("sdasd=", m.Get("sdasd"))
// }

func main() {
	root := dataStruct.LinkedList{}

	root.AddNode(5)

	root.PrintNodes()
	root.RemoveNode(root.Root)
	root.AddNode(6)
	root.AddNode(8)
	root.AddNode(9)
	root.AddNode(10)

	root.PrintNodes()
}
