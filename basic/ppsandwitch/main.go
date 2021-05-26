package main

import "fmt"

type Bread struct {
	val string
}

type StrawBerryJam struct {
	opened bool
}
type SpoonOfStrawBerry struct {
}

type Sandwitch struct {
	val string
}

func GetBreads(cnt int) []*Bread {
	breads := make([]*Bread, cnt)
	for i := 0; i < cnt; i++ {
		breads[i] = &Bread{val: "bread"}
	}
	return breads
}
func OpenStrawberryJam(jam *StrawBerryJam) {
	jam.opened = true
}

func GetOnespoon(_ *StrawBerryJam) *SpoonOfStrawBerry {
	return &SpoonOfStrawBerry{}
}
func PutJamOnBread(bread *Bread, jam *SpoonOfStrawBerry) {
	bread.val += "+StrawBerryJam"
}
func MakeSandwitch(breads []*Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	for i := 0; i < len(breads); i++ {
		sandwitch.val = breads[0].val + "+bread "
	}
	return sandwitch
}

func main() {
	// 	1. 빵 두개를 꺼낸다
	breads := GetBreads(2)

	jam := &StrawBerryJam{}
	// 2. 딸기잼 뚜껑을 연다
	OpenStrawberryJam(jam)
	// 3. 딸기잼을 한수저 떠서 빵 위에 올린다
	spoon := GetOnespoon(jam)
	// 4. 딸기잼을 잘 바른다
	PutJamOnBread(breads[0], spoon)
	// 5. 빵을 덮는다
	sandwitch := MakeSandwitch(breads)

	fmt.Println(sandwitch.val)
}

/*위와 같은 방법은
기능이 바뀌었을 때 또 새로운 코딩을 해야함
유지보수가 힘들다
수정을 거듭6할수록 스파게티 코드가 됨
산탄총 수정:
단순한 변경사항인데도 많은 양을 수정해야함
->OOP
어떤 코드를 짜더라도 잘 짜는거
Object->상태+기능
코딩-상태(메모리)를 어떻게 조정하는가
*/
