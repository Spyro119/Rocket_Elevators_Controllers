
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

