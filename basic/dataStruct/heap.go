package dataStruct

import "fmt"

//이진트리 형태

//slice로 구현 0번인덱스:루트
//부모idx=(현재-1)/2
//3번인덱스의 부모: 1번인덱스
//4번인덱스의 부모: 2번인덱스
type Heap struct {
	Root  int
	slice []int
}

//push value in the slice
func (h *Heap) Push(val int) {
	h.slice = append(h.slice, val)
	idx := len(h.slice) - 1 //4번
	for idx >= 0 {
		parentIdx := (idx - 1) / 2 //3번째가 1번째보다 크면
		if parentIdx < 0 {
			break
		}
		if h.slice[idx] < h.slice[parentIdx] {
			h.slice[idx], h.slice[parentIdx] = h.slice[parentIdx], h.slice[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *Heap) Print() {
	fmt.Println(h.slice)
}
func (h *Heap) Count() int {
	return len(h.slice)
}
func (h *Heap) Pop() int {
	if len(h.slice) == 0 {
		return 0
	}

	top := h.slice[0]
	last := h.slice[len(h.slice)-1]
	h.slice = h.slice[:len(h.slice)-1]

	if len(h.slice) == 0 {
		return top
	}

	h.slice[0] = last
	idx := 0
	for idx < len(h.slice) { //
		swapIdx := -1
		leftIdx := idx*2 + 1
		if leftIdx >= len(h.slice) {
			break
		}
		if h.slice[leftIdx] < h.slice[idx] {
			swapIdx = leftIdx
		}
		rightIdx := idx*2 + 2
		if rightIdx < len(h.slice) {
			if h.slice[rightIdx] < h.slice[idx] {
				if swapIdx < 0 || h.slice[swapIdx] > h.slice[rightIdx] {
					swapIdx = rightIdx
				}
			}
		}
		if swapIdx < 0 {
			break
		}
		h.slice[idx], h.slice[swapIdx] = h.slice[swapIdx], h.slice[idx]
		idx = swapIdx
	}
	return top
}
