package main

import "fmt"

//golang ducktyping(유추)

type SpoonOfJam interface {
	String() string
}
type Jam interface {
	GetOneSpoon() SpoonOfJam
}
type Bread struct {
	val string
}
type StrawBerryJam struct {
}
type OrangeJame struct {
}
type SpoonOfStrawBerry struct {
}
type SpoonOfOrangeJam struct {
}
type AppleJam struct {
}
type SpoonOfAppleJam struct {
}

func (j *SpoonOfAppleJam) String() string {
	return "+applejam"
}

//SpoonOfJam interface에 있음
func (s *SpoonOfOrangeJam) String() string {
	return "+orangejam"
}
func (s *SpoonOfStrawBerry) String() string {
	return "+strawberryjam"
}

func (b *Bread) Putjam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}
func (j *SpoonOfAppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}
func (j *SpoonOfOrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}

}
func (j *StrawBerryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawBerry{}
}

func (b *Bread) String() string {
	return "Bread" + b.val
}
func main() {
	bread := &Bread{}
	apjam := &SpoonOfAppleJam{}
	bread.Putjam(apjam)

	fmt.Println(bread)
}
