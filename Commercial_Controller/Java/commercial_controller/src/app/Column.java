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
    
            if(requestedFloor >= 1 && floorNumber >= 1 ){
                if (bestCol.elevatorList.get(i).currentFloor == floorNumber) {
                    bestCol.elevatorList.get(i).score = floorsDiff;
                } else if (floorNumber == 1 && bestCol.elevatorList.get(i).direction == "down"){
                    bestCol.elevatorList.get(i).score = floorsDiff + 100;
                } else if (floorNumber == 1 && bestCol.elevatorList.get(i).direction == "idle" ){
                    bestCol.elevatorList.get(i).score = floorsDiff + 200;
                } else if (floorNumber == 1 && bestCol.elevatorList.get(i).direction == "up") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 300;
                } else if (floorNumber > 1 && bestCol.elevatorList.get(i).currentFloor > floorNumber && bestCol.elevatorList.get(i).direction == "down") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 400;
                } else if (floorNumber > 1 && bestCol.elevatorList.get(i).currentFloor > floorNumber && bestCol.elevatorList.get(i).direction == "idle" ){
                    bestCol.elevatorList.get(i).score = floorsDiff + 500;
                } else if (floorNumber > 1 && bestCol.elevatorList.get(i).currentFloor > floorNumber && bestCol.elevatorList.get(i).direction == "up") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 600;
                } else if (floorNumber > 1 && bestCol.elevatorList.get(i).currentFloor < floorNumber){
                    bestCol.elevatorList.get(i).score = floorsDiff + 700;
                }
    
            } else {
                if( bestCol.elevatorList.get(i).direction == "idle" && bestCol.elevatorList.get(i).currentFloor == floorNumber ){
                    bestCol.elevatorList.get(i).score = floorsDiff;
                } else if (floorNumber == 1 && bestCol.elevatorList.get(i).direction == "up" ){
                    bestCol.elevatorList.get(i).score = floorsDiff + 100;
                } else if (floorNumber == 1 && bestCol.elevatorList.get(i).direction == "idle" ){
                    bestCol.elevatorList.get(i).score = floorsDiff + 200;
                } else if (floorNumber == 1 && bestCol.elevatorList.get(i).direction == "down") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 300;
                } else if (floorNumber < 1 && bestCol.elevatorList.get(i).currentFloor < floorNumber && bestCol.elevatorList.get(i).direction == "up") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 400;
                } else if (floorNumber < 1 && bestCol.elevatorList.get(i).currentFloor < floorNumber && bestCol.elevatorList.get(i).direction == "idle") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 500;
                } else if (floorNumber < 1 && bestCol.elevatorList.get(i).currentFloor < floorNumber && bestCol.elevatorList.get(i).direction == "down") {
                    bestCol.elevatorList.get(i).score = floorsDiff + 600;
                } else if (floorNumber < 1 && bestCol.elevatorList.get(i).currentFloor > floorNumber) {
                    bestCol.elevatorList.get(i).score = floorsDiff + 700;
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
 