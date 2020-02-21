class Column: 
    def __init__ (self, elevatorNumbers, numberFloors):
        self.elevatorList = []
        self.floorsList = []

#           ***** CREATING FLOORLIST *****
        for i in range(numberFloors):
            self.floorsList.append(i+1)

#           ***** CREATING ELEVATORLIST WITH OBJECT ELEVATOR *****
        for i in range(elevatorNumbers):
            elevator = Elevator(1, "", "elevator " + str(i+1))
            self.elevatorList.append(elevator)

    
    def requestElevator(self, requestedFloor, requestDirection):
        self.requestedFloor = requestedFloor
        self.requestDirection = requestDirection
        self.findElevator(requestedFloor, requestDirection)

    def requestFloor(self, destination):
        self.destination = destination
        self.elevatorList[0].moveD(destination)
        

#              ***** FIND BEST ELEVATOR FOR SCENARIO DEPENDING ON VARIABLES *****
    def findElevator(self, requestedFloor, requestDirection):
        for elevator in self.elevatorList:
            floorsDiff = abs(elevator.currentFloor - requestedFloor)
            if(elevator.currentFloor == requestedFloor and elevator.direction =="idle"):
                elevator.elevatorScore = 0 + floorsDiff
            elif(elevator.currentFloor < requestedFloor and requestDirection == elevator.direction and elevator.direction == "up"):
                elevator.elevatorScore = 100 + floorsDiff
            elif(elevator.currentFloor > requestedFloor and requestDirection == elevator.direction and elevator.direction == "down"):
                elevator.elevatorScore = 200 + floorsDiff
            elif(elevator.currentFloor == requestedFloor and elevator.direction == "idle"):
                elevator.elevatorScore = 300 + floorsDiff
            elif(elevator.currentFloor != requestedFloor and elevator.direction != requestDirection and elevator.direction != "idle"):
                elevator.elevatorScore = 400 + floorsDiff
            elif(elevator.currentFloor == requestedFloor and elevator.direction != requestDirection):
                elevator.elevatorScore = 500 + floorsDiff
            elif(elevator.currentFloor != requestedFloor and elevator.direction == requestDirection):
                elevator.elevatorScore = 600 + floorsDiff
            elif(elevator.direction == "idle"):
                elevator.elevatorScore = 700 + floorsDiff
        self.elevatorList.sort(key=lambda x: x.elevatorScore)
        bestElevator = self.elevatorList[0]
        bestElevator.move(requestedFloor)


class Elevator:
    def __init__ (self, currentFloor, direction, name):
        self.requestFloor = []
        self.direction = direction
        self.currentFloor = currentFloor
        self.name = name
        self.elevatorScore = 1000

    def move(self, requestFloor):
        print(self.name, "is sent.", self.name, " is currently at floor : ", self.currentFloor)
        while(self.currentFloor < requestFloor):
            self.direction = "up"
            self.currentFloor += 1
            print(self.name , "at floor : " , self.currentFloor)
        while(self.currentFloor > requestFloor):
            self.direction = "down"
            self.currentFloor -= 1
            print(self.name , "at floor: " , self.currentFloor)
        print("============================================================")
        print(self.name, "is arrived at destination. Doors open. You enter.")
        print("============================================================")
        self.direction = "idle"
        

    def moveD(self, destination):
        while(self.currentFloor < destination):
            self.direction = "up"
            self.currentFloor += 1
            print(self.name , "at floor : " , self.currentFloor)
        while(self.currentFloor > destination):
            self.direction = "down"
            self.currentFloor -= 1
            print(self.name , "at floor: " , self.currentFloor)
        print("===========================================================")
        print(self.name, " is arrived. Doors open. You exit.")
        print("===========================================================")
        self.direction = "idle"



    
#                                   ***** Changeable Variables *****
#If you change elevatorNumbers, make sure to uncomment "column1.elevatorList[i].blabla" below

#requestDirection = "up"
#requestFloor = 5
#destination = 10


#column1 = Column(elevatorNumbers, numberFloors)

#                   ***** SCENARIOS *****

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
                column1.requestFloor(destination)

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
                column1.requestFloor(destination)

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
                column1.requestFloor(destination)

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
                column1.requestFloor(destination)

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
                column1.requestFloor(destination)

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
                column1.requestFloor(destination)
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
