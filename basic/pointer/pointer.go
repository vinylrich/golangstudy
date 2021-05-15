package main

import "fmt"

//golang pointer 목적
//본래 pointer존재 명시
//연산,casting 막기

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintSungjeok() {
	fmt.Println(s.grade, s.class)
}

//모든 function의 class은 copy
func (s *Student) InputSungjeok(class string, grade string) {
	s.class = class
	s.grade = grade
}

func main() {
	var s Student
	s = Student{name: "junwoo", age: 18, class: "수학", grade: "A++"}
	s.PrintSungjeok()
	s.InputSungjeok("과학", "B+")
	s.PrintSungjeok()
}

func Increase(x *int) {
	*x = *x + 1
}

//포인터를 인자로 받으면 메모리 주소만 복사
//값을 함수 인자로 받으면 전체 속성이 복사
