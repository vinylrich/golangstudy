package main

import "fmt"

//상태+기능=Object
//어떤 한 상태가 있을 때 기능을 통해 외부와 소통
//각각의 잼의 기능 빵의 기능 샌드위치 기능의 관계 성립
type Bread struct {
	val string
}
type Jam struct {
	val string
}

//fmt.Println의 경우 이 String method가 있으면
//string 함수를 리턴한다
func (b *Bread) String() string {
	return b.val
}
func (j *Jam) GetVal() string {
	return "+ jam"
}
func (b *Bread) PutJam(jam *Jam) {
	b.val += jam.GetVal()
}

func main() {
	bread := &Bread{val: "bread"}
	jam := &Jam{}

	bread.PutJam(jam)

	fmt.Println(bread)
}

//객체

//상태+기능 외부 메서드-관계 이를 따로 정의한게
//인터페이스
// 내부메서드
