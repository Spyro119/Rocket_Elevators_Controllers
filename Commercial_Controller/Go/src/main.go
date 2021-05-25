package main

import "fmt"

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
