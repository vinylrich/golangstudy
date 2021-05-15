package main

import "fmt"

type Student struct {
	name  string
	class int

	sungjeok Sungjeok
}
type Sungjeok struct {
	name  string
	grade int
}

//위와 아래는 같은 함수(기능도 같음)
func (s Student) ViewSungJuk() {
	fmt.Println(s.sungjeok)
}
func ViewSungJuk(s Student) {
	fmt.Println(s.sungjeok)
}

//struct에서 함수에 포인터를 어떻게,왜 사용할까?
//값을 InputSungjuk에서 바꿨기 때문에 그 값을 main 함수까지 오게 만들기 위해서 사용하는 것
//golang에서의 함수 호출은 무조건 복사로 이루어짐
//여기서는 정상적으로 들어갔는데 main함수에서는 반영이 안된거임
func (s *Student) InputSungjuk(name string, grade int) {
	s.sungjeok.name = name
	s.sungjeok.grade = grade
}
func InputSungjuk(s Student, name string, grade int) {
	s.sungjeok.name = name
	s.sungjeok.grade = grade
}
func main() {
	s := Student{name: "철수", class: 1}
	s.sungjeok.name = "과학"
	s.sungjeok.grade = 80
	s.ViewSungJuk() //view 성적이라는 기능을 가지고 있음
	ViewSungJuk(s)
	s.InputSungjuk("수학", 70)
	//여기 입력값은 위 InputSungjuk이랑 다름
	s.ViewSungJuk()
}

/*
struct 배운내용
	- 구조 뿐만 아니라 기능도 가지고 있음(함수)
	- struct안에 struct
	- struct=class
	- OOP의 기본
	- java에서는 클래스 안에 method가 있지만 밖에 method를 만들 수 있음
*/
