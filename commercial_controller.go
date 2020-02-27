package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

//																	JUST RUN PROGRAMM FOR SCENARIOS
//																	INTERACTIVE INPUT ALREADY IN PLACE
//			    																Have fun ;)

// 											*****FUNCTION RETRY TO START ANOTHER SCENARIO ONCE 1 HAS FINISHED*****
func retry() {
	var input string
	fmt.Println("Want to try a new scenario? y/n")
	fmt.Scan(&input)
	for input != "y" && input != "n" {
		fmt.Println("Please type y for yes or n for no")
		fmt.Scan(&input)
	}
	if input == "y" {
		main()
	} else if input == "n" {
		fmt.Println("Program exiting. Thank you.")
	}

}

//                                     									 *****STRUCTURES*****
type battery struct {
	columnList        []column
	numberColumn      int
	floorNumbersTotal int
}
type column struct {
	name         string
	elevatorList []elevator
	floorList    []int
}

type elevator struct {
	name         string
	currentFloor int
	direction    string
	score        int
}

//                                                        *****INITIALISING BATTERY*****
func newBattery(elevatorNumbersTotal, numberFloorsTotal, numberColumn, numberOfBasement int) *battery {
	b := battery{}
	b.numberColumn = numberColumn
	b.floorNumbersTotal = numberFloorsTotal
	elevatorsPerColumn := elevatorNumbersTotal / numberColumn
	floorsPerColumn := numberFloorsTotal / (numberColumn - 1)

	//                                                    *****COLUMN GENERATOR*****
	for i := 1; i <= numberColumn; i++ {
		col := column{
			name:         "Column " + string(i+64),
			elevatorList: []elevator{},
			floorList:    []int{},
		}
		//                                                *****ELEVATOR GENERATOR*****
		for j := 0; j < elevatorsPerColumn; j++ {
			elev := elevator{
				name:         "elevator " + string(i+64) + strconv.Itoa(j+1),
				currentFloor: 1,
				direction:    "idle"}
			col.elevatorList = append(col.elevatorList, elev)
		}
		b.columnList = append(b.columnList, col)
	}

	for i := 0; i < numberColumn; i++ {
		for k := 1; k <= floorsPerColumn; k++ {
			if i == 0 {
				if k == 1 {
					//                               *****ADDING RC TO COLUMN DEDICATED TO BASEMENT *****
					b.columnList[i].floorList = append(b.columnList[i].floorList, k)
				}
				if k <= numberOfBasement {
					//                                     ***** BASEMENT FLOOR LIST TO COLUMN 0 *****
					b.columnList[i].floorList = append(b.columnList[i].floorList, -k)
				}
			} else if i == 1 {
				//                                            ***** FLOOR 1 TO 20 FOR 1ST COLUMN *****
				b.columnList[i].floorList = append(b.columnList[i].floorList, k)
			} else {
				//                            ***** ADD RC TO EVERY COLUMN THAT ISN'T IN INDEX 0 AND 1 *****
				if k == 1 {
					b.columnList[i].floorList = append(b.columnList[i].floorList, 1)
				}
				//       ***** ADDING FLOORS TO EVERY OTHER COLUMN, FloorsPerColumn BEING THE RANGE OF FLOORS PER COLUMN *****

				b.columnList[i].floorList = append(b.columnList[i].floorList, k+((i-1)*floorsPerColumn))

			}
		}
	}
	return &b
}

//												***** requestElevator METHOD FOR GOING TO RC *****
func (b *battery) requestElevator(floorNumber int) {
	if floorNumber != 1 {
		requestedFloor := 1
		b.findColumn(floorNumber, requestedFloor)
	}

}

//											***** assignElevator METHOD FROM RC TO DESTINATION*****
func (b *battery) assignElevator(requestedFloor int) {
	if requestedFloor != 1 {
		floorNumber := 1
		b.findColumn(floorNumber, requestedFloor)
	}
}

//                                                        *****FIND BEST COLUMN *****
func (b *battery) findColumn(floorNumber, requestedFloor int) *column {
	bestCol := column{}
	//                                                  ******ITERATING THROUGH COLUMN LIST *****
	for i := 0; i < len(b.columnList); i++ {
		//                                              *****ITERATING THROUGH FLOORLIST OF EACH COLUMN*****
		for _, k := range b.columnList[i].floorList {
			//              ******FOR K = FLOOR IN floorList, REQUEST FLOOR HAS TO BE RC (1) AND requestedFloor HAS TO BE IN COLUMN floorList ARRAY *****
			if floorNumber == 1 && requestedFloor == k {
				bestCol := b.columnList[i]
				bestCol.findElevator(floorNumber, requestedFloor)
				return &bestCol
				//             ***** FOR K = FLOOR in floorList, REQUEST FLOOR HAS TO BE IN COLUMN floorList ARRAY AND requestedFloor MUST BE RC *****
			} else if floorNumber == k && requestedFloor == 1 {
				bestCol := b.columnList[i]
				bestCol.findElevator(floorNumber, requestedFloor)
				return &bestCol
			}
		}
	}
	bestCol.findElevator(floorNumber, requestedFloor)
	return &bestCol
}

//                                          *****FIND ELEVATOR*****
func (bestCol *column) findElevator(floorNumber, requestedFloor int) *elevator {
	for i := 0; i < len(bestCol.elevatorList); i++ {
		var floorsDiff int = int(math.Abs(float64(bestCol.elevatorList[i].currentFloor - floorNumber)))

		//                                  ***** FIND ELEVATOR FOR UPPER FLOORS *****
		if requestedFloor >= 1 && floorNumber >= 1 {
			//                              *****ELEVATOR AT SAME FLOOR AS USER AND DIRECTION EITHER UP OR IDLE*****
			if bestCol.elevatorList[i].currentFloor == floorNumber {
				bestCol.elevatorList[i].score = floorsDiff
				//                              *****IF USER AT FLOOR 1 AND ELEVATOR IS GOING DOWN *****
			} else if floorNumber == 1 && bestCol.elevatorList[i].direction == "down" {
				bestCol.elevatorList[i].score = floorsDiff + 100
				//                              *****IF USER AT FLOOR 1 AND ELEVATOR IS IDLE *****
			} else if floorNumber == 1 && bestCol.elevatorList[i].direction == "idle" {
				bestCol.elevatorList[i].score = floorsDiff + 200
				//                              *****IF USER AT FLOOR 1 AND ELEVATOR IS GOING UP *****
			} else if floorNumber == 1 && bestCol.elevatorList[i].direction == "up" {
				bestCol.elevatorList[i].score = floorsDiff + 300
				//                              *****IF USER IS GOING DOWN, ELEVATOR IS > USER FLOOR AND ELEVATOR IS GOING DOWN *****
			} else if floorNumber > 1 && bestCol.elevatorList[i].currentFloor > floorNumber && bestCol.elevatorList[i].direction == "down" {
				bestCol.elevatorList[i].score = floorsDiff + 400
				//                              *****IF USER IS GOING DOWN, ELEVATOR IS < USER FLOOR AND ELEVATOR IS IDLE *****
			} else if floorNumber > 1 && bestCol.elevatorList[i].currentFloor > floorNumber && bestCol.elevatorList[i].direction == "idle" {
				bestCol.elevatorList[i].score = floorsDiff + 500
				//                              *****IF USER IS GOING DOWN, ELEVATOR IS > USER AND DIRECTION IS UP *****
			} else if floorNumber > 1 && bestCol.elevatorList[i].currentFloor > floorNumber && bestCol.elevatorList[i].direction == "up" {
				bestCol.elevatorList[i].score = floorsDiff + 600
				//                              *****IF USER IS GOING DOWN, ELEVATOR IS < USER *****
			} else if floorNumber > 1 && bestCol.elevatorList[i].currentFloor < floorNumber {
				bestCol.elevatorList[i].score = floorsDiff + 700
			}

			//                                  *****FIND BEST ELEVATOR FOR BASEMENTS*****
		} else if requestedFloor <= 1 && floorNumber <= 1 {
			//                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR IS AT 1ST FLOOR*****
			if bestCol.elevatorList[i].direction == "idle" && bestCol.elevatorList[i].currentFloor == floorNumber {
				bestCol.elevatorList[i].score = floorsDiff
				//                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR GOES UP*****
			} else if floorNumber == 1 && bestCol.elevatorList[i].direction == "up" {
				bestCol.elevatorList[i].score = floorsDiff + 100
				//                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR IS IDLE*****
			} else if floorNumber == 1 && bestCol.elevatorList[i].direction == "idle" {
				bestCol.elevatorList[i].score = floorsDiff + 200
				//                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR GOES DOWN*****
			} else if floorNumber == 1 && bestCol.elevatorList[i].direction == "down" {
				bestCol.elevatorList[i].score = floorsDiff + 300
				//                              *****USER IS <1 AND ELEVATOR IS UNDER USER FLOOR AND ELEVATOR GOES UP*****
			} else if floorNumber < 1 && bestCol.elevatorList[i].currentFloor < floorNumber && bestCol.elevatorList[i].direction == "up" {
				bestCol.elevatorList[i].score = floorsDiff + 400
				//                              *****USER IS <1 AND ELEVATOR IS UNDER USER FLOOR AND ELEVATOR IS IDLE*****
			} else if floorNumber < 1 && bestCol.elevatorList[i].currentFloor < floorNumber && bestCol.elevatorList[i].direction == "idle" {
				bestCol.elevatorList[i].score = floorsDiff + 500
				//                              *****USER IS <1 AND ELEVATOR IS UNDER USER FLOOR AND ELEVATOR GOES DOWN*****
			} else if floorNumber < 1 && bestCol.elevatorList[i].currentFloor < floorNumber && bestCol.elevatorList[i].direction == "down" {
				bestCol.elevatorList[i].score = floorsDiff + 600
				//                              *****USER IS <1 AND ELEVATOR IS ABOVE USER FLOOR *****
			} else if floorNumber < 1 && bestCol.elevatorList[i].currentFloor > floorNumber {
				bestCol.elevatorList[i].score = floorsDiff + 700
				//                              *****USER IS <1 AND ELEVATOR DIRECTION IS DOWN*****
			} else if floorNumber < 1 && bestCol.elevatorList[i].direction == "down" {
				bestCol.elevatorList[i].score = floorsDiff + 800
			}
		}
	}
	//										*****SORT ELEVATOR LIST ARRAY FROM LOWEST SCORE TO HIGHEST*****
	sort.Slice(bestCol.elevatorList[:], func(i, j int) bool {
		return bestCol.elevatorList[i].score < bestCol.elevatorList[j].score
	})
	//							*****bestElev ALWAYS IS IN INDEX 0. THIS WAY, ONE ELEVATOR IS ALWAYS SENT EVEN IF 2 ELEVATORS HAS THE SAME LOWEST SCORE*****
	bestElev := bestCol.elevatorList[0]
	bestElev.move(floorNumber)
	bestElev.moveD(requestedFloor)

	return &bestElev

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

func main() {
	//                                            *****VARIABLES*****
	var elevatorNumbersTotal = 20
	var numberFloorsTotal = 60
	var numberOfBasement = 6
	var numberColumn = 4

	//                                           *****FUNCTION NEW BATTERY*****
	b := newBattery(elevatorNumbersTotal, numberFloorsTotal, numberColumn, numberOfBasement)

	//                                           *****SCENARIOS INPUT****

	var scenario int
	fmt.Println("What scenario do you want to do? Choose 1, 2, 3 or 4.")
	fmt.Scan(&scenario)
	for scenario != 1 && scenario != 2 && scenario != 3 && scenario != 4 {
		fmt.Println("Please enter a valid number. 1, 2, 3 or 4.")
		fmt.Scan(&scenario)
	}
	if scenario == 1 {
		//Column B
		//Elevator B1 at 20th Floor, going "down" to 5th
		b.columnList[1].elevatorList[0].currentFloor = 20
		b.columnList[1].elevatorList[0].direction = "down"

		//Elevator B2 at 3rd, going "up" to 15th
		b.columnList[1].elevatorList[1].currentFloor = 3
		b.columnList[1].elevatorList[1].direction = "up"

		//Elevator B3 at 13th, going "down" to 1st
		b.columnList[1].elevatorList[2].currentFloor = 13
		b.columnList[1].elevatorList[2].direction = "down"

		//Elevator B4 at 15th, going "down" to 2nd
		b.columnList[1].elevatorList[3].currentFloor = 15
		b.columnList[1].elevatorList[3].direction = "down"

		//Elevator B5 at 6th, going "down" to 1st
		b.columnList[1].elevatorList[4].currentFloor = 6
		b.columnList[1].elevatorList[4].direction = "down"

		//1st, someone is at 1st floor and requests the 20th floor
		var requestedFloor = 20
		b.assignElevator(requestedFloor)

		//, elevator B5 is expected to be sent in Scenario 1

	} else if scenario == 2 {

		//Column C
		//Elevator C1 at 1ts Floor, going "up" to 21st
		b.columnList[2].elevatorList[0].currentFloor = 1
		b.columnList[2].elevatorList[0].direction = "up"

		//Elevator C2 at 23rd, going "up" to 28th
		b.columnList[2].elevatorList[1].currentFloor = 23
		b.columnList[2].elevatorList[1].direction = "up"

		//Elevator C3 at 33th, going "down" to 1st
		b.columnList[2].elevatorList[2].currentFloor = 33
		b.columnList[2].elevatorList[2].direction = "down"

		//Elevator C4 at 40th, going "down" to 24th
		b.columnList[2].elevatorList[3].currentFloor = 40
		b.columnList[2].elevatorList[3].direction = "down"

		//Elevator C5 at 39th, going "down" to 1st
		b.columnList[2].elevatorList[4].currentFloor = 39
		b.columnList[2].elevatorList[4].direction = "down"

		//1st, someone is at 1st floor and requests the 36th floor, elevator C1 is expected to be sent
		var requestedFloor = 36
		b.assignElevator(requestedFloor)

	} else if scenario == 3 {
		//Column D
		//Elevator D1 at 58th Floor, going "down" to 1st
		b.columnList[3].elevatorList[0].currentFloor = 58
		b.columnList[3].elevatorList[0].direction = "down"

		//Elevator D2 at 50th, going "up" to 60th
		b.columnList[3].elevatorList[1].currentFloor = 50
		b.columnList[3].elevatorList[1].direction = "up"

		//Elevator D3 at 46th, going "up" to 58th
		b.columnList[3].elevatorList[2].currentFloor = 46
		b.columnList[3].elevatorList[2].direction = "up"

		//Elevator D4 at 1st, going "down" to 54th
		b.columnList[3].elevatorList[3].currentFloor = 1
		b.columnList[3].elevatorList[3].direction = "up"

		//Elevator D5 at 60th, going "down" to 1st
		b.columnList[3].elevatorList[4].currentFloor = 60
		b.columnList[3].elevatorList[4].direction = "down"

		//1st, someone is at 54th floor and requests the 1st floor, elevator D1 is expected to be sent
		var floorNumber = 54
		b.requestElevator(floorNumber)

	} else if scenario == 4 {
		//Column A
		//Elevator A1 at B4th Floor, Being "idle"
		b.columnList[0].elevatorList[0].currentFloor = -4
		b.columnList[0].elevatorList[0].direction = "idle"

		//Elevator A2 "idle" at 1st Floor
		b.columnList[0].elevatorList[1].currentFloor = 1
		b.columnList[0].elevatorList[1].direction = "idle"

		//Elevator A3 at B3rd, going "down" to B5th
		b.columnList[0].elevatorList[2].currentFloor = -3
		b.columnList[0].elevatorList[2].direction = "down"

		//Elevator A4 at B6th, going "up" to 1st
		b.columnList[0].elevatorList[3].currentFloor = -6
		b.columnList[0].elevatorList[3].direction = "up"

		//Elevator A5 at B1st, going "down" to b6th
		b.columnList[0].elevatorList[4].currentFloor = -1
		b.columnList[0].elevatorList[4].direction = "down"

		//1st, someone is at B3rd floor and requests the 1st floor, elevator A4 is expected to be sent
		var numberFloor = -3
		b.requestElevator(numberFloor)

	}
	retry()
}
