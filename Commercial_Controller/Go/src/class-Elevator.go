package main

import (
	"fmt"
	"strconv"
)

type elevator struct {
	name         string
	currentFloor int
	direction    string
	score        int
}

// 														*****FUNCTION MOVE ELEVATOR TO FLOORNUMBER*****
func (bestElev *elevator) move(floorNumber int) {
	fmt.Println(bestElev.name, "is sent.", bestElev.name, " is currently at floor : ", strconv.Itoa(bestElev.currentFloor))
	for bestElev.currentFloor < floorNumber {
		bestElev.direction = "up"
		bestElev.currentFloor++
		fmt.Println(bestElev.name, " is at floor : ", strconv.Itoa(bestElev.currentFloor))
	}
	for bestElev.currentFloor > floorNumber {
		bestElev.direction = "down"
		bestElev.currentFloor--
		fmt.Println(bestElev.name, " is at floor : ", strconv.Itoa(bestElev.currentFloor))
	}

	fmt.Println("============================================================")
	fmt.Println(bestElev.name, "is arrived at destination. Doors open. You enter.")
	fmt.Println("============================================================")
	bestElev.direction = "idle"

}

//														*****FUNCTION MOVE ELEVATOR TO DESTINATION*****
//													(2 functions only because print is slightly different)
func (bestElev *elevator) moveD(requestedFloor int) {
	fmt.Println("Doors closes. ", bestElev.name, " starts moving.")
	for bestElev.currentFloor < requestedFloor {
		bestElev.direction = "up"
		bestElev.currentFloor++
		if bestElev.currentFloor != 0 {
			fmt.Println(bestElev.name, " is at floor : ", strconv.Itoa(bestElev.currentFloor))
		}

	}
	for bestElev.currentFloor > requestedFloor {
		bestElev.direction = "down"
		bestElev.currentFloor--
		if bestElev.currentFloor != 0 {
			fmt.Println(bestElev.name, " is at floor : ", strconv.Itoa(bestElev.currentFloor))
		}

	}

	fmt.Println("============================================================")
	fmt.Println(bestElev.name, "is arrived at destination. Doors open. You exit.")
	fmt.Println("============================================================")
	bestElev.direction = "idle"
}