package main

import (
	"math"
	"sort"
)

type column struct {
	name         string
	elevatorList []elevator
	floorList    []int
}

//                                          *****FIND ELEVATOR*****
func (bestCol *column) findElevator(floorNumber, requestedFloor int) *elevator {
	for i := 0; i < len(bestCol.elevatorList); i++ {
		var floorsDiff int = int(math.Abs(float64(bestCol.elevatorList[i].currentFloor - floorNumber)))

		if requestedFloor >= 1 && floorNumber >= 1 {
			if bestCol.elevatorList[i].currentFloor == floorNumber {
				bestCol.elevatorList[i].score = floorsDiff
			} else if floorNumber == 1 && bestCol.elevatorList[i].direction == "down" {
				bestCol.elevatorList[i].score = floorsDiff + 100
			} else if floorNumber == 1 && bestCol.elevatorList[i].direction == "idle" {
				bestCol.elevatorList[i].score = floorsDiff + 200
			} else if floorNumber == 1 && bestCol.elevatorList[i].direction == "up" {
				bestCol.elevatorList[i].score = floorsDiff + 300
			} else if floorNumber > 1 && bestCol.elevatorList[i].currentFloor > floorNumber && bestCol.elevatorList[i].direction == "down" {
				bestCol.elevatorList[i].score = floorsDiff + 400
			} else if floorNumber > 1 && bestCol.elevatorList[i].currentFloor > floorNumber && bestCol.elevatorList[i].direction == "idle" {
				bestCol.elevatorList[i].score = floorsDiff + 500
			} else if floorNumber > 1 && bestCol.elevatorList[i].currentFloor > floorNumber && bestCol.elevatorList[i].direction == "up" {
				bestCol.elevatorList[i].score = floorsDiff + 600
			} else if floorNumber > 1 && bestCol.elevatorList[i].currentFloor < floorNumber {
				bestCol.elevatorList[i].score = floorsDiff + 700
			}

			//                                  *****FIND BEST ELEVATOR FOR BASEMENTS*****
		} else if requestedFloor <= 1 && floorNumber <= 1 {
			if bestCol.elevatorList[i].direction == "idle" && bestCol.elevatorList[i].currentFloor == floorNumber {
				bestCol.elevatorList[i].score = floorsDiff
			} else if floorNumber == 1 && bestCol.elevatorList[i].direction == "up" {
				bestCol.elevatorList[i].score = floorsDiff + 100
			} else if floorNumber == 1 && bestCol.elevatorList[i].direction == "idle" {
				bestCol.elevatorList[i].score = floorsDiff + 200
			} else if floorNumber == 1 && bestCol.elevatorList[i].direction == "down" {
				bestCol.elevatorList[i].score = floorsDiff + 300
			} else if floorNumber < 1 && bestCol.elevatorList[i].currentFloor < floorNumber && bestCol.elevatorList[i].direction == "up" {
				bestCol.elevatorList[i].score = floorsDiff + 400
			} else if floorNumber < 1 && bestCol.elevatorList[i].currentFloor < floorNumber && bestCol.elevatorList[i].direction == "idle" {
				bestCol.elevatorList[i].score = floorsDiff + 500
			} else if floorNumber < 1 && bestCol.elevatorList[i].currentFloor < floorNumber && bestCol.elevatorList[i].direction == "down" {
				bestCol.elevatorList[i].score = floorsDiff + 600
			} else if floorNumber < 1 && bestCol.elevatorList[i].currentFloor > floorNumber {
				bestCol.elevatorList[i].score = floorsDiff + 700
			} else if floorNumber < 1 && bestCol.elevatorList[i].direction == "down" {
				bestCol.elevatorList[i].score = floorsDiff + 800
			}
		}
	}
	//										*****SORT ELEVATOR LIST ARRAY FROM LOWEST SCORE TO HIGHEST*****
	sort.Slice(bestCol.elevatorList[:], func(i, j int) bool {
		return bestCol.elevatorList[i].score < bestCol.elevatorList[j].score
	})
	bestElev := bestCol.elevatorList[0]
	bestElev.move(floorNumber)
	bestElev.moveD(requestedFloor)

	return &bestElev

}