package main

import "fmt"

func RemoveBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1] //10
}
func RemoveFront(a []int) ([]int, int) {
	return a[1:], a[0] //11
}
func main() {

	a := make([]int, 0, 12)
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)

	for i := 0; i < 5; i++ {
		var back int
		a, back = RemoveBack(a)
		fmt.Printf("%d, ", back)
	}
	fmt.Println(a)
	for i := 0; i < 5; i++ {
		var front int
		a, front = RemoveFront(a)
		fmt.Printf("%d, ", front)
	}
	fmt.Println(a)
}
