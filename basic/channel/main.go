package main

import "fmt"

func pop(c chan int) {
	fmt.Println("pop") //3
	v := <-c           //4
	fmt.Println(v)     //5
}
func main() {
	c := make(chan int) //1

	go pop(c) //2
	c <- 10   //5

	fmt.Println("end of program") //5,6
}

/*goThread 1
Main함수
goThread생성
pop 실행
*/

/*process
Main thread 실행
pop thread 실행
pop thread에서 v에 c값이 들어갈때까지 기다림
c에 10이 들어감
v출력

이걸 통해서 컨베이어시스템
producer-consumer시스템을 만들 수 있음
자신이 1을 다 하면 2를 하기 위해 다른사람에게 넘김
2를 다 하면 3을 함 넘기는 과정에서 channel사용
*/
