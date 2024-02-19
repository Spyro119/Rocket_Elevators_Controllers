package classes

import (
	"math"
	"sort"
)

type Column struct {
	Name         string
	ElevatorList []Elevator
	FloorList    []int
}

// *****FIND Elevator*****
func (bestCol *Column) findElevator(floorNumber, requestedFloor int) *Elevator {
	for i := 0; i < len(bestCol.ElevatorList); i++ {
		var floorsDiff int = int(math.Abs(float64(bestCol.ElevatorList[i].CurrentFloor - floorNumber)))

		if requestedFloor >= 1 && floorNumber >= 1 {
			if bestCol.ElevatorList[i].CurrentFloor == floorNumber {
				bestCol.ElevatorList[i].Score = floorsDiff
			} else if floorNumber == 1 && bestCol.ElevatorList[i].Direction == "down" {
				bestCol.ElevatorList[i].Score = floorsDiff + 100
			} else if floorNumber == 1 && bestCol.ElevatorList[i].Direction == "idle" {
				bestCol.ElevatorList[i].Score = floorsDiff + 200
			} else if floorNumber == 1 && bestCol.ElevatorList[i].Direction == "up" {
				bestCol.ElevatorList[i].Score = floorsDiff + 300
			} else if floorNumber > 1 && bestCol.ElevatorList[i].CurrentFloor > floorNumber && bestCol.ElevatorList[i].Direction == "down" {
				bestCol.ElevatorList[i].Score = floorsDiff + 400
			} else if floorNumber > 1 && bestCol.ElevatorList[i].CurrentFloor > floorNumber && bestCol.ElevatorList[i].Direction == "idle" {
				bestCol.ElevatorList[i].Score = floorsDiff + 500
			} else if floorNumber > 1 && bestCol.ElevatorList[i].CurrentFloor > floorNumber && bestCol.ElevatorList[i].Direction == "up" {
				bestCol.ElevatorList[i].Score = floorsDiff + 600
			} else if floorNumber > 1 && bestCol.ElevatorList[i].CurrentFloor < floorNumber {
				bestCol.ElevatorList[i].Score = floorsDiff + 700
			}

			//                                  *****FIND BEST Elevator FOR BASEMENTS*****
		} else if requestedFloor <= 1 && floorNumber <= 1 {
			if bestCol.ElevatorList[i].Direction == "idle" && bestCol.ElevatorList[i].CurrentFloor == floorNumber {
				bestCol.ElevatorList[i].Score = floorsDiff
			} else if floorNumber == 1 && bestCol.ElevatorList[i].Direction == "up" {
				bestCol.ElevatorList[i].Score = floorsDiff + 100
			} else if floorNumber == 1 && bestCol.ElevatorList[i].Direction == "idle" {
				bestCol.ElevatorList[i].Score = floorsDiff + 200
			} else if floorNumber == 1 && bestCol.ElevatorList[i].Direction == "down" {
				bestCol.ElevatorList[i].Score = floorsDiff + 300
			} else if floorNumber < 1 && bestCol.ElevatorList[i].CurrentFloor < floorNumber && bestCol.ElevatorList[i].Direction == "up" {
				bestCol.ElevatorList[i].Score = floorsDiff + 400
			} else if floorNumber < 1 && bestCol.ElevatorList[i].CurrentFloor < floorNumber && bestCol.ElevatorList[i].Direction == "idle" {
				bestCol.ElevatorList[i].Score = floorsDiff + 500
			} else if floorNumber < 1 && bestCol.ElevatorList[i].CurrentFloor < floorNumber && bestCol.ElevatorList[i].Direction == "down" {
				bestCol.ElevatorList[i].Score = floorsDiff + 600
			} else if floorNumber < 1 && bestCol.ElevatorList[i].CurrentFloor > floorNumber {
				bestCol.ElevatorList[i].Score = floorsDiff + 700
			} else if floorNumber < 1 && bestCol.ElevatorList[i].Direction == "down" {
				bestCol.ElevatorList[i].Score = floorsDiff + 800
			}
		}
	}
	//										*****SORT Elevator LIST ARRAY FROM LOWEST Score TO HIGHEST*****
	sort.Slice(bestCol.ElevatorList[:], func(i, j int) bool {
		return bestCol.ElevatorList[i].Score < bestCol.ElevatorList[j].Score
	})
	bestElev := bestCol.ElevatorList[0]
	bestElev.Move(floorNumber)
	bestElev.MoveD(requestedFloor)

	return &bestElev

}
