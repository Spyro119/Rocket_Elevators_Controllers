from .elevator import Elevator

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

    def requestFloor(self, elevator, destination):
        self.destination = destination
        self.elevator = self.elevatorList[0]
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

