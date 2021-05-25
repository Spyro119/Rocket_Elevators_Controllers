using System;
using System.Collections.Generic;
using Class.ElevatorObject;

namespace Class.ColumnObject {

    public class Column{
    public string name;
    public List<int> floorList;
    public List<Elevator> elevatorList;

    public Column(string name, int numberOfElevators,int numberOfFloors){
    this.name = name;
    this.elevatorList = new List<Elevator>();
    this.floorList = new List<int>();
        }

        public Elevator findElevator(int floorNumber, int requestedFloor, Column bestCol){
        for (var i = 0; i < bestCol.elevatorList.Count; i++){
            int floorsDiff = Math.Abs(bestCol.elevatorList[i].currentFloor - floorNumber);

            //                                  ***** FIND ELEVATOR FOR UPPER FLOORS *****
            if(requestedFloor >= 1 && floorNumber >= 1 ){
                //                              *****ELEVATOR AT SAME FLOOR AS USER AND DIRECTION EITHER UP OR IDLE*****
                if (bestCol.elevatorList[i].currentFloor == floorNumber) {
                    bestCol.elevatorList[i].score = floorsDiff;
                    //                              *****IF USER AT FLOOR 1 AND ELEVATOR IS GOING DOWN *****
                } else if (floorNumber == 1 && bestCol.elevatorList[i].direction == "down"){
                    bestCol.elevatorList[i].score = floorsDiff + 100;
                    //                              *****IF USER AT FLOOR 1 AND ELEVATOR IS IDLE *****
                } else if (floorNumber == 1 && bestCol.elevatorList[i].direction == "idle" ){
                    bestCol.elevatorList[i].score = floorsDiff + 200;
                    //                              *****IF USER AT FLOOR 1 AND ELEVATOR IS GOING UP *****
                } else if (floorNumber == 1 && bestCol.elevatorList[i].direction == "up") {
                    bestCol.elevatorList[i].score = floorsDiff + 300;
                    //                              *****IF USER IS GOING DOWN, ELEVATOR IS > USER FLOOR AND ELEVATOR IS GOING DOWN *****
                } else if (floorNumber > 1 && bestCol.elevatorList[i].currentFloor > floorNumber && bestCol.elevatorList[i].direction == "down") {
                    bestCol.elevatorList[i].score = floorsDiff + 400;
                    //                              *****IF USER IS GOING DOWN, ELEVATOR IS < USER FLOOR AND ELEVATOR IS IDLE *****
                } else if (floorNumber > 1 && bestCol.elevatorList[i].currentFloor > floorNumber && bestCol.elevatorList[i].direction == "idle" ){
                    bestCol.elevatorList[i].score = floorsDiff + 500;
                    //                              *****IF USER IS GOING DOWN, ELEVATOR IS > USER AND DIRECTION IS UP *****
                } else if (floorNumber > 1 && bestCol.elevatorList[i].currentFloor > floorNumber && bestCol.elevatorList[i].direction == "up") {
                    bestCol.elevatorList[i].score = floorsDiff + 600;
                    //                              *****IF USER IS GOING DOWN, ELEVATOR IS < USER *****
                } else if (floorNumber > 1 && bestCol.elevatorList[i].currentFloor < floorNumber){
                    bestCol.elevatorList[i].score = floorsDiff + 700;
                }

                //                                  *****FIND BEST ELEVATOR FOR BASEMENTS*****
            } else if (requestedFloor <= 1 && floorNumber <= 1) {
                //                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR IS AT 1ST FLOOR*****
                if( bestCol.elevatorList[i].direction == "idle" && bestCol.elevatorList[i].currentFloor == floorNumber ){
                    bestCol.elevatorList[i].score = floorsDiff;
                    //                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR GOES UP*****
                } else if (floorNumber == 1 && bestCol.elevatorList[i].direction == "up" ){
                    bestCol.elevatorList[i].score = floorsDiff + 100;
                    //                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR IS IDLE*****
                } else if (floorNumber == 1 && bestCol.elevatorList[i].direction == "idle" ){
                    bestCol.elevatorList[i].score = floorsDiff + 200;
                    //                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR GOES DOWN*****
                } else if (floorNumber == 1 && bestCol.elevatorList[i].direction == "down") {
                    bestCol.elevatorList[i].score = floorsDiff + 300;
                    //                              *****USER IS <1 AND ELEVATOR IS UNDER USER FLOOR AND ELEVATOR GOES UP*****
                } else if (floorNumber < 1 && bestCol.elevatorList[i].currentFloor < floorNumber && bestCol.elevatorList[i].direction == "up") {
                    bestCol.elevatorList[i].score = floorsDiff + 400;
                    //                              *****USER IS <1 AND ELEVATOR IS UNDER USER FLOOR AND ELEVATOR IS IDLE*****
                } else if (floorNumber < 1 && bestCol.elevatorList[i].currentFloor < floorNumber && bestCol.elevatorList[i].direction == "idle") {
                    bestCol.elevatorList[i].score = floorsDiff + 500;
                    //                              *****USER IS <1 AND ELEVATOR IS UNDER USER FLOOR AND ELEVATOR GOES DOWN*****
                } else if (floorNumber < 1 && bestCol.elevatorList[i].currentFloor < floorNumber && bestCol.elevatorList[i].direction == "down") {
                    bestCol.elevatorList[i].score = floorsDiff + 600;
                    //                              *****USER IS <1 AND ELEVATOR IS ABOVE USER FLOOR *****
                } else if (floorNumber < 1 && bestCol.elevatorList[i].currentFloor > floorNumber) {
                    bestCol.elevatorList[i].score = floorsDiff + 700;
                    //                              *****USER IS <1 AND ELEVATOR DIRECTION IS DOWN*****
                } else if (floorNumber < 1 && bestCol.elevatorList[i].direction == "down") {
                    bestCol.elevatorList[i].score = floorsDiff + 800;
                }
            }
        }
        //										*****SORT ELEVATOR LIST ARRAY FROM LOWEST SCORE TO HIGHEST*****
    for(var i = 0; i < bestCol.elevatorList.Count; i++)
        bestCol.elevatorList.Sort((p1, p2) =>
    {
    return p1.score - p2.score;
    });

        
        //							*****bestElev ALWAYS IS IN INDEX 0. THIS WAY, ONE ELEVATOR IS ALWAYS SENT EVEN IF 2 ELEVATORS HAS THE SAME LOWEST SCORE*****
        var bestElev = bestCol.elevatorList[0];
    bestElev.move(floorNumber, bestElev);
    bestElev.moveD(requestedFloor, bestElev);
    Console.WriteLine("Best elevator is : " + bestElev.name);

        return bestElev;

    }

    }
}
