package app;
import java.util.ArrayList;
//import  java.lang.reflect.Array;

public class Battery{
    
    public int numberOfColumn;
    public int floorNumbersTotal;
    public ArrayList<Column> columnList;
      
    public Battery(int elevatorNumbersTotal, int numberOfColumn, int floorNumbersTotal, int numberOfBasement){
    this.numberOfColumn = numberOfColumn;
    this.floorNumbersTotal = floorNumbersTotal;
    this.columnList = new ArrayList<Column>(); 
    var numberOfElevators = elevatorNumbersTotal / numberOfColumn;
    var numberOfFloors = floorNumbersTotal / (numberOfColumn - 1);

    
//                                  *****GENERATING COLUMNS*****
for (var i = 1; i <= numberOfColumn; i++){
    var col = new Column("Column" + i, numberOfElevators, numberOfFloors);
    columnList.add(col);

//                                  *****GENERATING ELEVATORS FOR EACH COLUMN*****
    for (var j =1; j <= numberOfElevators; j++){
        var elev = new Elevator("Elevator" + i + "." + j, 1, "idle", 1);
        col.elevatorList.add(elev);
        }
    }

for (var i = 0; i < numberOfColumn; i++ ){
    for (var k = 1; k <= numberOfFloors; k++){
        if (i == 0){
            if (k == 1) {
                //                               *****ADDING RC TO COLUMN DEDICATED TO BASEMENT *****
                columnList.get(i).floorList.add(k);
            }
            if (k <= numberOfBasement) {
                //                                     ***** BASEMENT FLOOR LIST TO COLUMN 0 *****
                columnList.get(i).floorList.add(-k);
            }
        } else if (i == 1){
            //                                            ***** FLOOR 1 TO 20 FOR 1ST COLUMN *****
            columnList.get(i).floorList.add(k);
        } else {
            //                            ***** ADD RC TO EVERY COLUMN THAT ISN'T IN INDEX 0 AND 1 *****
            if (k == 1) {
                columnList.get(i).floorList.add(1);
            }
            //       ***** ADDING FLOORS TO EVERY OTHER COLUMN, FloorsPerColumn BEING THE RANGE OF FLOORS PER COLUMN *****

            columnList.get(i).floorList.add( k+((i-1)*numberOfFloors));


            }
        }
    }
    }
    public void requestElevator(int floorNumber){
        if (floorNumber != 1){
          int requestedFloor = 1;
          this.findColumn(floorNumber, requestedFloor);
      
        } 
  }
      public void assignElevator(int requestedFloor){
        if (requestedFloor != 1){
         int floorNumber = 1;
        this.findColumn(floorNumber, requestedFloor);
      } 
  }
  Column findColumn (int floorNumber, int requestedFloor){
    Column bestCol = this.columnList.get(0);
     int size = this.columnList.size();
    for (var i =0; i < size; i++)
    {
      //                                              *****ITERATING THROUGH FLOORLIST OF EACH COLUMN*****
      for (int k : this.columnList.get(i).floorList){
          //              ******FOR K = FLOOR IN floorList, REQUEST FLOOR HAS TO BE RC (1) AND requestedFloor HAS TO BE IN COLUMN floorList ARRAY *****
          if (floorNumber == 1 && requestedFloor == k){
              bestCol = this.columnList.get(i);
          //	bestCol.findElevator(floorNumber, requestedFloor)
              //             ***** FOR K = FLOOR in floorList, REQUEST FLOOR HAS TO BE IN COLUMN floorList ARRAY AND requestedFloor MUST BE RC *****
          } else if (floorNumber == k && requestedFloor == 1){
              bestCol = this.columnList.get(i);
              
          }
      }
  }
  
    bestCol.findElevator(floorNumber, requestedFloor, bestCol);
  return bestCol;
}
}

