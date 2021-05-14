package main

import "fmt"

func main() {
	arr := [10]int{1, 5, 4, 3, 2, 1, 5, 6, 7, 8}
	temp := [10]int{}
	for i := 0; i < len(arr); i++ {
		idx := arr[i] //idx 5
		temp[idx]++   //5번 인덱스가 ++

	}
	idx := 0
	for i := 0; i < len(temp); i++ { //10번
		for j := 0; j < temp[i]; j++ { //숫자가 있는 만큼
			arr[idx] = i
			idx++
		}
	}
	fmt.Println(arr)
}
