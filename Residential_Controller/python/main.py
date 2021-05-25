from includes.column import Column

#                   ***** SCENARIO #1 *****
def scenario():
    while True:
        try:
            scenario = int(input("Which scenario you'd like to start with? (Type 1, 2 or 3) :"))
            while(scenario < 1 or scenario > 3):
                scenario = int(input("Please enter a valid number."))

            if(scenario == 1):
                #Elevator 1 idle at floor 2, elevator 2 idle at floor 6, someone request elevator at floor 3 and wants to go to Floor 7
                elevatorNumbers = 2 
                numberFloors = 10
                column1 = Column(elevatorNumbers, numberFloors)
                #RequestElevator sent at floor 3
                requestFloor = 3
                requestDirection = "up"
                #Wants to go to floor 7
                destination = 7
                #elevator 1
                column1.elevatorList[0].currentFloor = 2
                column1.elevatorList[0].direction = "idle"
                #elevator 2
                column1.elevatorList[1].currentFloor = 6
                column1.elevatorList[1].direction = "idle"
                #Move elevator to floor 3, then floor 7
                column1.requestElevator(requestFloor, requestDirection)
                column1.requestFloor(column1.elevatorList[0] , destination)

                retry()

            elif(scenario == 2):
                #Elevator 1 idle at floor 10, elevator 2 idle at floor 3, someone on 1st floor and request 6th floor.
                elevatorNumbers = 2 
                numberFloors = 10
                column1 = Column(elevatorNumbers, numberFloors)
                #RequestElevator sent at floor 1
                requestFloor = 1
                requestDirection = "up"
                #Wants to go to floor 6
                destination = 6
                #elevator 1 idle at floor 10
                column1.elevatorList[0].currentFloor = 10
                column1.elevatorList[0].direction = "idle"
                #elevator 2 idle at floor 3
                column1.elevatorList[1].currentFloor = 3
                column1.elevatorList[1].direction = "idle"
                #Move elevator to floor 1, then floor 6
                column1.requestElevator(requestFloor, requestDirection)
                column1.requestFloor(column1.elevatorList[0] , destination)

                print("")
                print("*****2 minutes later, another request is sent*****")
                print("")

                #New request from floor 3
                requestFloor = 3
                requestDirection = "up"
                #Wants to go to floor 5
                destination = 5
                #move elevator to floor 3, then 5
                column1.requestElevator(requestFloor, requestDirection)
                column1.requestFloor(column1.elevatorList[0], destination)

                print("")
                print("*****A third request is sent.*****")
                print("")

                #New request from floor 3
                requestFloor = 9
                requestDirection = "down"
                #Wants to go to floor 2
                destination = 2
                #move elevator to floor 9, then 2
                column1.requestElevator(requestFloor, requestDirection)
                column1.requestFloor(column1.elevatorList[0] , destination)

                retry()
            
            elif(scenario == 3):
                #Elevator 1 idle at floor 10, elevator 2 moving up from floor 3 to 6, someone on 3rd floor and request 2nd floor.
                elevatorNumbers = 2 
                numberFloors = 10
                column1 = Column(elevatorNumbers, numberFloors)
                #RequestElevator sent at floor 3
                requestFloor = 3
                requestDirection = "down"
                #Wants to go to floor 6
                destination = 2
                #elevator 1 idle at floor 10
                column1.elevatorList[0].currentFloor = 10
                column1.elevatorList[0].direction = "idle"
                #elevator 2 moving up at floor 3 to 6
                column1.elevatorList[1].currentFloor = 3
                column1.elevatorList[1].direction = "up"
                #Move elevator to floor 1, then floor 6
                column1.requestElevator(requestFloor, requestDirection)
                column1.requestFloor(column1.elevatorList[0], destination)

                print("")
                print("*****5 minutes later, another request is sent.*****")
                print("")

                #New request from floor 10
                requestFloor = 10
                requestDirection = "down"
                #Wants to go to floor 3
                destination = 3
                #elevator 2 is now at floor 6
                column1.elevatorList[1].currentFloor = 6
                column1.elevatorList[1].direction = "idle"
                #move elevator to floor 3, then 5
                column1.requestElevator(requestFloor, requestDirection)
                column1.requestFloor(column1.elevatorList[0], destination)
                retry()
            break
        except ValueError:
            print("Please enter a valid number.")

def retry():
    retry = input("Want to retry another scenario? y/n : ")
    while(retry != "y" and retry != "n"):
        input("please choose y or n.").lower
        print(retry)
    if(retry == "y"):
        scenario()
    else:
        print("Thank you for trying out my program. Exiting.")

scenario()
