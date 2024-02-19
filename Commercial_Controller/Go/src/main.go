package main

import (
	"fmt"

	"github.com/Spyro119/Rocket_Elevators_Controllers/classes"
)

//																	JUST RUN PROGRAMM FOR SCENARIOS
//																	INTERACTIVE INPUT ALREADY IN PLACE

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

func main() {
	//                                            *****VARIABLES*****
	var elevatorNumbersTotal = 20
	var numberFloorsTotal = 60
	var numberOfBasement = 6
	var numberColumn = 4

	b := classes.NewBattery(elevatorNumbersTotal, numberFloorsTotal, numberColumn, numberOfBasement)

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
		b.ColumnList[1].ElevatorList[0].CurrentFloor = 20
		b.ColumnList[1].ElevatorList[0].Direction = "down"

		//Elevator B2 at 3rd, going "up" to 15th
		b.ColumnList[1].ElevatorList[1].CurrentFloor = 3
		b.ColumnList[1].ElevatorList[1].Direction = "up"

		//Elevator B3 at 13th, going "down" to 1st
		b.ColumnList[1].ElevatorList[2].CurrentFloor = 13
		b.ColumnList[1].ElevatorList[2].Direction = "down"

		//Elevator B4 at 15th, going "down" to 2nd
		b.ColumnList[1].ElevatorList[3].CurrentFloor = 15
		b.ColumnList[1].ElevatorList[3].Direction = "down"

		//Elevator B5 at 6th, going "down" to 1st
		b.ColumnList[1].ElevatorList[4].CurrentFloor = 6
		b.ColumnList[1].ElevatorList[4].Direction = "down"

		//1st, someone is at 1st floor and requests the 20th floor
		var requestedFloor = 20
		b.AssignElevator(requestedFloor)

		//, elevator B5 is expected to be sent in Scenario 1

	} else if scenario == 2 {

		//Column C
		//Elevator C1 at 1ts Floor, going "up" to 21st
		b.ColumnList[2].ElevatorList[0].CurrentFloor = 1
		b.ColumnList[2].ElevatorList[0].Direction = "up"

		//Elevator C2 at 23rd, going "up" to 28th
		b.ColumnList[2].ElevatorList[1].CurrentFloor = 23
		b.ColumnList[2].ElevatorList[1].Direction = "up"

		//Elevator C3 at 33th, going "down" to 1st
		b.ColumnList[2].ElevatorList[2].CurrentFloor = 33
		b.ColumnList[2].ElevatorList[2].Direction = "down"

		//Elevator C4 at 40th, going "down" to 24th
		b.ColumnList[2].ElevatorList[3].CurrentFloor = 40
		b.ColumnList[2].ElevatorList[3].Direction = "down"

		//Elevator C5 at 39th, going "down" to 1st
		b.ColumnList[2].ElevatorList[4].CurrentFloor = 39
		b.ColumnList[2].ElevatorList[4].Direction = "down"

		//1st, someone is at 1st floor and requests the 36th floor, elevator C1 is expected to be sent
		var requestedFloor = 36
		b.AssignElevator(requestedFloor)

	} else if scenario == 3 {
		//Column D
		//Elevator D1 at 58th Floor, going "down" to 1st
		b.ColumnList[3].ElevatorList[0].CurrentFloor = 58
		b.ColumnList[3].ElevatorList[0].Direction = "down"

		//Elevator D2 at 50th, going "up" to 60th
		b.ColumnList[3].ElevatorList[1].CurrentFloor = 50
		b.ColumnList[3].ElevatorList[1].Direction = "up"

		//Elevator D3 at 46th, going "up" to 58th
		b.ColumnList[3].ElevatorList[2].CurrentFloor = 46
		b.ColumnList[3].ElevatorList[2].Direction = "up"

		//Elevator D4 at 1st, going "down" to 54th
		b.ColumnList[3].ElevatorList[3].CurrentFloor = 1
		b.ColumnList[3].ElevatorList[3].Direction = "up"

		//Elevator D5 at 60th, going "down" to 1st
		b.ColumnList[3].ElevatorList[4].CurrentFloor = 60
		b.ColumnList[3].ElevatorList[4].Direction = "down"

		//1st, someone is at 54th floor and requests the 1st floor, elevator D1 is expected to be sent
		var floorNumber = 54
		b.RequestElevator(floorNumber)

	} else if scenario == 4 {
		//Column A
		//Elevator A1 at B4th Floor, Being "idle"
		b.ColumnList[0].ElevatorList[0].CurrentFloor = -4
		b.ColumnList[0].ElevatorList[0].Direction = "idle"

		//Elevator A2 "idle" at 1st Floor
		b.ColumnList[0].ElevatorList[1].CurrentFloor = 1
		b.ColumnList[0].ElevatorList[1].Direction = "idle"

		//Elevator A3 at B3rd, going "down" to B5th
		b.ColumnList[0].ElevatorList[2].CurrentFloor = -3
		b.ColumnList[0].ElevatorList[2].Direction = "down"

		//Elevator A4 at B6th, going "up" to 1st
		b.ColumnList[0].ElevatorList[3].CurrentFloor = -6
		b.ColumnList[0].ElevatorList[3].Direction = "up"

		//Elevator A5 at B1st, going "down" to b6th
		b.ColumnList[0].ElevatorList[4].CurrentFloor = -1
		b.ColumnList[0].ElevatorList[4].Direction = "down"

		//1st, someone is at B3rd floor and requests the 1st floor, elevator A4 is expected to be sent
		var numberFloor = -3
		b.RequestElevator(numberFloor)

	}
	retry()
}
