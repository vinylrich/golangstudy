package main

import "fmt"

func main() {
	var m map[string]string
	m = make(map[string]string)
	m["abd"] = "bbb"
	fmt.Println(m["abd"])

	m3 := make(map[int]bool)
	m3[4] = true

	fmt.Println(m3[6])
	//value 뒤 ok를 통해서
	//값이 있는지 없는지 확인 가능
	v, ok := m["abd"]
	fmt.Println(v, ok)

	//내장 delete함수로 지울 수 있음
	//delete(m, "abd")

	//range 키워드를 통해서
	//map, 배열 등을 순회할 수 있음

	h := []string{"hi", "hello"}
	for idx, value := range h {
		fmt.Println(idx, value)
	}
}
