package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}
type Plane struct {
	val string
}

func MakeTire(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "c_Tire, "
			outCarChan <- car

		case plane := <-planeChan:
			plane.val += "p_Tire, "
			outPlaneChan <- plane
		}

	}
}

func MakeEngine(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "c_Engine, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "p_Engine, "
			outPlaneChan <- plane
		}
	}
}
func StartCarWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car: " + strconv.Itoa(i)}
		i++
	}
}
func StartPlaneWork(chan1 chan Plane) {
	i := 0
	for {
		time.Sleep(time.Second)
		chan1 <- Plane{val: "Plane: " + strconv.Itoa(i)}
		i++
	}
}
func main() {
	Carchan1 := make(chan Car)
	Carchan2 := make(chan Car)
	Carchan3 := make(chan Car)

	Planechan1 := make(chan Plane)
	Planechan2 := make(chan Plane)
	Planechan3 := make(chan Plane)

	go StartCarWork(Carchan1)
	go StartPlaneWork(Planechan1)
	go MakeTire(Carchan1, Planechan1, Carchan2, Planechan2)
	go MakeEngine(Carchan2, Planechan2, Carchan3, Planechan3)

	for {
		select {
		case CarResult := <-Carchan3:
			fmt.Println(CarResult.val)
		case PlaneResult := <-Planechan3:
			fmt.Println(PlaneResult.val)
		}

	}

}
