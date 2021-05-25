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
                columnList.get(i).floorList.add(k);
            }
            if (k <= numberOfBasement) {
                columnList.get(i).floorList.add(-k);
            }
        } else if (i == 1){
            columnList.get(i).floorList.add(k);
        } else {
            if (k == 1) {
                columnList.get(i).floorList.add(1);
            }
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
      for (int k : this.columnList.get(i).floorList){
          if (floorNumber == 1 && requestedFloor == k){
              bestCol = this.columnList.get(i);
          } else if (floorNumber == k && requestedFloor == 1){
              bestCol = this.columnList.get(i);
          }
      }
  }
  
    bestCol.findElevator(floorNumber, requestedFloor, bestCol);
  return bestCol;
}
}

