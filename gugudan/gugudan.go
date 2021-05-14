package main

import "fmt"

func fx(x int) int {
	return (x * x) + 2
}
func main() {
	var x int
	fmt.Scanln(&x)
	fmt.Println(fx(x))
}
