package classes

import (
	"fmt"
	"strconv"
)

type Elevator struct {
	Name         string
	CurrentFloor int
	Direction    string
	Score        int
}

// *****FUNCTION MOVE Elevator TO FLOORNUMBER*****
func (bestElev *Elevator) Move(floorNumber int) {
	fmt.Println(bestElev.Name, "is sent.", bestElev.Name, " is currently at floor : ", strconv.Itoa(bestElev.CurrentFloor))
	for bestElev.CurrentFloor < floorNumber {
		bestElev.Direction = "up"
		bestElev.CurrentFloor++
		fmt.Println(bestElev.Name, " is at floor : ", strconv.Itoa(bestElev.CurrentFloor))
	}
	for bestElev.CurrentFloor > floorNumber {
		bestElev.Direction = "down"
		bestElev.CurrentFloor--
		fmt.Println(bestElev.Name, " is at floor : ", strconv.Itoa(bestElev.CurrentFloor))
	}

	fmt.Println("============================================================")
	fmt.Println(bestElev.Name, "is arrived at destination. Doors open. You enter.")
	fmt.Println("============================================================")
	bestElev.Direction = "idle"

}

// Function to move Elevator to destination
func (bestElev *Elevator) MoveD(requestedFloor int) {
	fmt.Println("Doors closes. ", bestElev.Name, " starts moving.")
	for bestElev.CurrentFloor < requestedFloor {
		bestElev.Direction = "up"
		bestElev.CurrentFloor++
		if bestElev.CurrentFloor != 0 {
			fmt.Println(bestElev.Name, " is at floor : ", strconv.Itoa(bestElev.CurrentFloor))
		}

	}
	for bestElev.CurrentFloor > requestedFloor {
		bestElev.Direction = "down"
		bestElev.CurrentFloor--
		if bestElev.CurrentFloor != 0 {
			fmt.Println(bestElev.Name, " is at floor : ", strconv.Itoa(bestElev.CurrentFloor))
		}

	}

	fmt.Println("============================================================")
	fmt.Println(bestElev.Name, "is arrived at destination. Doors open. You exit.")
	fmt.Println("============================================================")
	bestElev.Direction = "idle"
}
