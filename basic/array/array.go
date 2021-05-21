package main

import "fmt"

func main() {
	arr := make([]int, 0, 0)
	for {
		i := 0
		temp := 0

		fmt.Scanln(&temp)
		if temp == 0 {
			break
		}
		arr = append(arr, temp)
		i++
	}
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Println(arr)

}

//천천히 하는 것을 두려워 하지 말고
//가다 멈추는 것을 두려워하라

//if dfs.child==nil
