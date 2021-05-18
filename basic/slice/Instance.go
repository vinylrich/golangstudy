package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

//Reference type
func (t *Student) SetName(newName string) {
	t.name = newName
}
func (t *Student) SetAge(age int) {
	t.age = age
}
func PrintStudent(a *Student) {
	fmt.Println(*a)
}
func main() {
	a := Student{"junwoo", 20, 3}
	a.SetName("bbb")
	a.SetAge(15)
	fmt.Println(a)
	PrintStudent(&a)
}

//a라는 구조체나 인터페이스와 관련된 함수
//구조체의 기능!!

//Instance
//go
