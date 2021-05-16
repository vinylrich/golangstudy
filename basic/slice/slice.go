package main

import "fmt"

//동적배열 "slice"
//동적 배열은 실제 배열을 포인트 하고 있다가
//길이가 늘어나면 더 큰 길이의 배열을 다시 만들고
//거기에 복사하고 포인터를 더 큰 길이의 배열에 참조
func Printlencap(s []int) {
	fmt.Println(len(s))
	fmt.Println(cap(s))

}
func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 0, 8)
	Printlencap(s1)
	Printlencap(s2)

	s1 = append(s1, 5, 6, 7, 8)
	fmt.Println(s1)

	a := []int{1, 2}
	b := append(a, 3)
	c := append(b, 4)
	fmt.Printf("%p %p %p\n", a, b, c)
}

//공간이 늘어나면 두배씩 메모리 길이가 늘어남
