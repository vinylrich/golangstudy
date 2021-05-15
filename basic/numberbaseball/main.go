package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	strikes int
	balls   int
}

//make 3 random integer
func MakeNumbers() [3]int {
	var rst [3]int
	for i := 0; i < len(rst); i++ {
		for {
			n := rand.Intn(10)
			duplicated := false
			for j := 0; j < i; j++ {
				if rst[j] == n {
					duplicated = true
					break
				}
			}
			if !duplicated {
				rst[i] = n
				break
			}
		}
	}
	return rst
}

//usr Input 3 integer
func InputNumbers() [3]int {
	// 키보드로부터 0~9 사이의 겹치지 않는 숫자 3개를 입력받아 반환한다.
	var rst [3]int

	for {
		fmt.Println("겹치지 않는 0~9 사이의 숫자 3개를 입력하세요.")
		var no int
		_, err := fmt.Scanf("%d\n", &no)
		if err != nil {
			fmt.Println("잘못 입력하셨습니다.")
			continue
		}

		success := true
		idx := 0
		for no > 0 {
			n := no % 10
			no = no / 10

			duplicated := false
			for j := 0; j < idx; j++ {
				if rst[j] == n {
					// 겹친다. 다시 뽑는다.
					duplicated = true
					break
				}
			}
			if duplicated {
				fmt.Println("숫자가 겹치지 않아야 합니다.")
				success = false
				break
			}

			if idx >= 3 {
				fmt.Println("3개보다 많은 숫자를 입력하셨습니다.")
				success = false
				break
			}

			rst[idx] = n
			idx++
		}
		if success && idx < 3 {
			fmt.Println("3개의 숫자를 입력하세요.")
			success = false
		}

		if !success {
			continue
		}
		break
	}
	rst[0], rst[2] = rst[2], rst[0]
	//fmt.Println(rst)
	return rst
}

//com vs usr
func CompareNumbers(com, usr [3]int) Result {
	var r Result
	for i := 0; i < len(com); i++ {
		for j := 0; j < 3; j++ {
			if com[i] == usr[j] {
				if i == j {
					r.strikes++
				} else {
					r.balls++
				}
				break
			}
		}
	}

	return r
}
func PrintResult(result Result) {
	fmt.Printf("%dS%dB\n", result.strikes, result.balls)
}

//
func IsGameEnd(result Result) bool {
	return result.strikes == 3
}
func main() {
	rand.Seed(time.Now().UnixNano())
	// 무작위 숫자 3개 만들기
	com_numbers := MakeNumbers()
	// fmt.Println(com_numbers)
	cnt := 0
	for {
		cnt++
		//입력받기
		usr_numbers := InputNumbers()
		// fmt.Println(usr_numbers)
		result := CompareNumbers(com_numbers, usr_numbers)
		PrintResult(result)
		if IsGameEnd(result) {
			break
		}
	}
	fmt.Println(cnt, "번 만에 게임 끝!")
}
