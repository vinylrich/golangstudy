package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
	val string
}

type StructB struct {
	val int
}

func (s *StructA) String() string {
	return "Val:" + s.val
}

type Printable interface {
	String() string
}

//이런식으로 인터페이스를 매개변수로 넣어서
//매개변수를 참조하여 인터페이스 안에 있는
//메서드를 사용 할 수 있다

func Println(p Printable) {
	fmt.Println(p.String())
}

func (s *StructB) String() string {
	return "StructB" + strconv.Itoa(s.val)
}

func main() {
	a := &StructA{val: "AAA"}
	Println(a)

	b := &StructB{val: 10}
	Println(b)
}
