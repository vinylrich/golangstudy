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
