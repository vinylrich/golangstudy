package main

import "fmt"

type calculator interface {
	add(int, int)
}

func add(x, y int) int {
	return x + y
}
func subtract(x, y int) int {
	return x - y
}
func divine(x, y int) int {
	return x / y
}
func mul(x, y int) int {
	return x * y
}
func main() {
	var c int
	var x, y int
	fmt.Scanln(&c)
	fmt.Scanln(&x, &y)
	switch c {
	case 1:
		fmt.Println(add(x, y))
	case 2:
		fmt.Println(subtract(x, y))
	case 3:
		fmt.Println(divine(x, y))
	case 4:
		fmt.Println(mul(x, y))
	default:
		return
	}

}
