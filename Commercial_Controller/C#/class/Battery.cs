using System;
using System.Collections.Generic;
using Class.ColumnObject;
using Class.ElevatorObject;

namespace Class.Battery {


public class Battery{
  public int numberOfColumn;
  public int floorNumbersTotal;
  public List<Column> columnList;

  public Battery(int elevatorNumbersTotal, int numberOfColumn, int floorNumbersTotal, int numberOfBasement){
  this.numberOfColumn = numberOfColumn;
  this.floorNumbersTotal = floorNumbersTotal;
  this.columnList = new List<Column>(); 
  var numberOfElevators = elevatorNumbersTotal / numberOfColumn;
  var numberOfFloors = floorNumbersTotal / (numberOfColumn - 1);

//                                  *****GENERATING COLUMNS AND ELEVATORS FOR EACH COLUMNS*****
    for (var i = 1; i <= numberOfColumn; i++){
        var col = new Column("Column" + i, numberOfElevators, numberOfFloors);
        columnList.Add(col);

        for (var j =1; j <= numberOfElevators; j++){
            var elev = new Elevator("Elevator" + i + "." + j, 1, "idle", 1);
            col.elevatorList.Add(elev);
            }
        }
    
    for (var i = 0; i < numberOfColumn; i++ ){
		for (var k = 1; k <= numberOfFloors; k++){
			if (i == 0){
				if (k == 1) {
					//                               *****ADDING RC TO COLUMN DEDICATED TO BASEMENT *****
					columnList[i].floorList.Add(k);
				}
				if (k <= numberOfBasement) {
					//                                     ***** BASEMENT FLOOR LIST TO COLUMN 0 *****
					columnList[i].floorList.Add(-k);
				}
			} else if (i == 1){
				//                                            ***** FLOOR 1 TO 20 FOR 1ST COLUMN *****
				columnList[i].floorList.Add(k);
			} else {
				//                            ***** ADD RC TO EVERY COLUMN THAT ISN'T IN INDEX 0 AND 1 *****
				if (k == 1) {
					columnList[i].floorList.Add(1);
				}
				//       ***** ADDING FLOORS TO EVERY OTHER COLUMN, FloorsPerColumn BEING THE RANGE OF FLOORS PER COLUMN *****

				columnList[i].floorList.Add( k+((i-1)*numberOfFloors));


        }
      }
    }
    }
 public void requestElevator(int floorNumber){
      if (floorNumber != 1){
        int requestedFloor = 1;
        findColumn(floorNumber, requestedFloor);
    
      } 
}
    public void assignElevator(int requestedFloor){
      if (requestedFloor != 1){
       int floorNumber = 1;
      findColumn(floorNumber, requestedFloor);
    } 

    

}
    Column findColumn (int floorNumber, int requestedFloor){
      Column bestCol = this.columnList[0];
      for (var i =0; i < this.columnList.Count; i++)
      {
		foreach (int k in this.columnList[i].floorList)
    {
			if (floorNumber == 1 && requestedFloor == k){
				bestCol = this.columnList[i];
			} else if (floorNumber == k && requestedFloor == 1){
				bestCol = this.columnList[i];
				
			}
		}
	}
    
	    bestCol.findElevator(floorNumber, requestedFloor, bestCol);
        
        return bestCol;
        }
    }
}
