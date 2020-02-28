package app;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.Collections;

public class Column{
    public String name;
    public ArrayList<Integer> floorList;
    public ArrayList<Elevator> elevatorList;
  
    public Column(String name, int numberOfElevators,int numberOfFloors){
    this.name = name;
    this.elevatorList = new ArrayList<Elevator>();
    this.floorList = new ArrayList<Integer>();
      }
      public Elevator findElevator(int floorNumber, int requestedFloor, Column bestCol){
          int size = bestCol.elevatorList.size();
        for (var i = 0; i < size; i++){
            int floorsDiff = Math.abs(bestCol.elevatorList.get(i).currentFloor - floorNumber);
    
            //                                  ***** FIND ELEVATOR FOR UPPER FLOORS *****
            if(requestedFloor >= 1 && floorNumber >= 1 ){
                //                              *****ELEVATOR AT SAME FLOOR AS USER AND DIRECTION EITHER UP OR IDLE*****
                if (bestCol.elevatorList.get(i).currentFloor == floorNumber) {
                    bestCol.elevatorList.get(i).score = floorsDiff;
                    //                              *****IF USER AT FLOOR 1 AND ELEVATOR IS GOING DOWN *****
                } else if (floorNumber == 1 && bestCol.elevatorList.get(i).direction == "down"){
                    bestCol.elevatorList.get(i).score = floorsDiff + 100;
                    //                              *****IF USER AT FLOOR 1 AND ELEVATOR IS IDLE *****
                } else if (floorNumber == 1 && bestCol.elevatorList.get(i).direction == "idle" ){
                    bestCol.elevatorList.get(i).score = floorsDiff + 200;
                    //                              *****IF USER AT FLOOR 1 AND ELEVATOR IS GOING UP *****
                } else if (floorNumber == 1 && bestCol.elevatorList.get(i).direction == "up") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 300;
                    //                              *****IF USER IS GOING DOWN, ELEVATOR IS > USER FLOOR AND ELEVATOR IS GOING DOWN *****
                } else if (floorNumber > 1 && bestCol.elevatorList.get(i).currentFloor > floorNumber && bestCol.elevatorList.get(i).direction == "down") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 400;
                    //                              *****IF USER IS GOING DOWN, ELEVATOR IS < USER FLOOR AND ELEVATOR IS IDLE *****
                } else if (floorNumber > 1 && bestCol.elevatorList.get(i).currentFloor > floorNumber && bestCol.elevatorList.get(i).direction == "idle" ){
                    bestCol.elevatorList.get(i).score = floorsDiff + 500;
                    //                              *****IF USER IS GOING DOWN, ELEVATOR IS > USER AND DIRECTION IS UP *****
                } else if (floorNumber > 1 && bestCol.elevatorList.get(i).currentFloor > floorNumber && bestCol.elevatorList.get(i).direction == "up") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 600;
                    //                              *****IF USER IS GOING DOWN, ELEVATOR IS < USER *****
                } else if (floorNumber > 1 && bestCol.elevatorList.get(i).currentFloor < floorNumber){
                    bestCol.elevatorList.get(i).score = floorsDiff + 700;
                }
    
                //                                  *****FIND BEST ELEVATOR FOR BASEMENTS*****
            } else {
                //                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR IS AT 1ST FLOOR*****
                if( bestCol.elevatorList.get(i).direction == "idle" && bestCol.elevatorList.get(i).currentFloor == floorNumber ){
                    bestCol.elevatorList.get(i).score = floorsDiff;
                    //                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR GOES UP*****
                } else if (floorNumber == 1 && bestCol.elevatorList.get(i).direction == "up" ){
                    bestCol.elevatorList.get(i).score = floorsDiff + 100;
                    //                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR IS IDLE*****
                } else if (floorNumber == 1 && bestCol.elevatorList.get(i).direction == "idle" ){
                    bestCol.elevatorList.get(i).score = floorsDiff + 200;
                    //                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR GOES DOWN*****
                } else if (floorNumber == 1 && bestCol.elevatorList.get(i).direction == "down") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 300;
                    //                              *****USER IS <1 AND ELEVATOR IS UNDER USER FLOOR AND ELEVATOR GOES UP*****
                } else if (floorNumber < 1 && bestCol.elevatorList.get(i).currentFloor < floorNumber && bestCol.elevatorList.get(i).direction == "up") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 400;
                    //                              *****USER IS <1 AND ELEVATOR IS UNDER USER FLOOR AND ELEVATOR IS IDLE*****
                } else if (floorNumber < 1 && bestCol.elevatorList.get(i).currentFloor < floorNumber && bestCol.elevatorList.get(i).direction == "idle") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 500;
                    //                              *****USER IS <1 AND ELEVATOR IS UNDER USER FLOOR AND ELEVATOR GOES DOWN*****
                } else if (floorNumber < 1 && bestCol.elevatorList.get(i).currentFloor < floorNumber && bestCol.elevatorList.get(i).direction == "down") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 600;
                    //                              *****USER IS <1 AND ELEVATOR IS ABOVE USER FLOOR *****
                } else if (floorNumber < 1 && bestCol.elevatorList.get(i).currentFloor > floorNumber) {
                    bestCol.elevatorList.get(i).score = floorsDiff + 700;
                    //                              *****USER IS <1 AND ELEVATOR DIRECTION IS DOWN*****
                } else if (floorNumber < 1 && bestCol.elevatorList.get(i).direction == "down") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 800;
                }
            } 
        }
        
        //										*****SORT ELEVATOR LIST ARRAY FROM LOWEST SCORE TO HIGHEST*****


            
            

            Collections.sort(bestCol.elevatorList, new Comparator<Elevator>() {
                @Override
                public int compare(Elevator o1, Elevator o2) {
                    return o1.score.compareTo(o2.score);
                }
            });
            
            var bestElev = bestCol.elevatorList.get(0);
            bestElev.move(floorNumber, bestElev);
            bestElev.moveD(requestedFloor, bestElev);

            
                return bestElev;
    
      }
    
    }   
 