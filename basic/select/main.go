package main

import (
	"fmt"
	"time"
)

func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++
	}
}
func main() {
	c := make(chan int)
	timer := time.After(11 * time.Second)
	tickTimer := time.Tick(2 * time.Second)
	go push(c)
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-timer:
			fmt.Println("timeout")
			return
		case <-tickTimer:
			fmt.Println("Tick")
		}

	}
}
