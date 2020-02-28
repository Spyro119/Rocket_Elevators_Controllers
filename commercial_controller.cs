using System;
using System.Collections.Generic;


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

//                                  *****GENERATING COLUMNS*****
    for (var i = 1; i <= numberOfColumn; i++){
        var col = new Column("Column" + i, numberOfElevators, numberOfFloors);
        columnList.Add(col);

//                                  *****GENERATING ELEVATORS FOR EACH COLUMN*****
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
		//                                              *****ITERATING THROUGH FLOORLIST OF EACH COLUMN*****
		foreach (int k in this.columnList[i].floorList)
    {
			//              ******FOR K = FLOOR IN floorList, REQUEST FLOOR HAS TO BE RC (1) AND requestedFloor HAS TO BE IN COLUMN floorList ARRAY *****
			if (floorNumber == 1 && requestedFloor == k){
				bestCol = this.columnList[i];
			//	bestCol.findElevator(floorNumber, requestedFloor)
				//             ***** FOR K = FLOOR in floorList, REQUEST FLOOR HAS TO BE IN COLUMN floorList ARRAY AND requestedFloor MUST BE RC *****
			} else if (floorNumber == k && requestedFloor == 1){
				bestCol = this.columnList[i];
				
			}
		}
	}
    
	bestCol.findElevator(floorNumber, requestedFloor, bestCol);

	return bestCol;
}
}

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
public class Elevator{
  public string name;
  public int currentFloor;
  public string direction;
  public int score;

  public Elevator(string name, int currentFloor, string direction, int score){
    this.name = name;
    this.currentFloor = currentFloor;
    this.direction = direction;
    this.score = score;
  }

  // 														*****FUNCTION MOVE ELEVATOR TO FLOORNUMBER*****
public void move(int floorNumber, Elevator bestElev) {
Console.WriteLine(bestElev.name + "is sent. " + bestElev.name + " is currently at floor : " + bestElev.currentFloor);
	while (bestElev.currentFloor < floorNumber ){
		bestElev.direction = "up";
		bestElev.currentFloor++;
	Console.WriteLine(bestElev.name + " is at floor : " + bestElev.currentFloor);
	}
	while (bestElev.currentFloor > floorNumber){
		bestElev.direction = "down";
		bestElev.currentFloor--;
	Console.WriteLine(bestElev.name + " is at floor : " + bestElev.currentFloor);
	}

Console.WriteLine("============================================================");
Console.WriteLine(bestElev.name + "is arrived at destination. Doors open. You enter.");
Console.WriteLine("============================================================");
	bestElev.direction = "idle";

}

//														*****FUNCTION MOVE ELEVATOR TO DESTINATION*****
//													(2 functions only because print is slightly different)
public void moveD(int requestedFloor, Elevator bestElev) {
	Console.WriteLine("Doors closes. ", bestElev.name, " starts moving.");
	while (bestElev.currentFloor < requestedFloor ){
		bestElev.direction = "up";
		bestElev.currentFloor++;
		if (bestElev.currentFloor != 0 ){
			Console.WriteLine(bestElev.name + " is at floor : " + bestElev.currentFloor);
		}

	}
	while (bestElev.currentFloor > requestedFloor) {
		bestElev.direction = "down";
		bestElev.currentFloor--;
		if (bestElev.currentFloor != 0 ){
		Console.WriteLine(bestElev.name + " is at floor : "+ bestElev.currentFloor);
		}

	}

Console.WriteLine("============================================================");
Console.WriteLine(bestElev.name + " is arrived at destination. Doors open. You exit.");
Console.WriteLine("============================================================");
	bestElev.direction = "idle";
}
}



class Program
{
    static void Main()
    {



      var elevatorNumbersTotal = 20;
	    var floorNumbersTotal = 60;
	    var numberOfBasement = 6;
	    var numberOfColumn = 4;
      //                                *****CREATING BATTERY ******
      Battery battery = new Battery(elevatorNumbersTotal,numberOfColumn,floorNumbersTotal,numberOfBasement);

//                                      *****SCENARIOS INPUTS*****

        Console.WriteLine("Please choose a scenario: 1, 2, 3 or 4");
        string scenario = Console.ReadLine();
        while (scenario != "1" && scenario != "2" && scenario != "3" && scenario != "4"){
        Console.WriteLine("Please choose a scenario: 1, 2, 3 or 4");
        scenario = Console.ReadLine();
        }
        if (scenario == "1"){
          //Elevator B1 at 20th Floor, going "down" to 5th
		battery.columnList[1].elevatorList[0].currentFloor = 20;
		battery.columnList[1].elevatorList[0].direction = "down";

		//Elevator B2 at 3rd, going "up" to 15th
		battery.columnList[1].elevatorList[1].currentFloor = 3;
		battery.columnList[1].elevatorList[1].direction = "up";

		//Elevator B3 at 13th, going "down" to 1st
		battery.columnList[1].elevatorList[2].currentFloor = 13;
		battery.columnList[1].elevatorList[2].direction = "down";

		//Elevator B4 at 15th, going "down" to 2nd
		battery.columnList[1].elevatorList[3].currentFloor = 15;
		battery.columnList[1].elevatorList[3].direction = "down";

		//Elevator B5 at 6th, going "down" to 1st
		battery.columnList[1].elevatorList[4].currentFloor = 6;
		battery.columnList[1].elevatorList[4].direction = "down";

		//1st, someone is at 1st floor and requests the 20th floor
		var requestedFloor = 20;
		battery.assignElevator(requestedFloor);

 } else if(scenario == "2"){
          //Column C
		//Elevator C1 at 1ts Floor, going "up" to 21st
		battery.columnList[2].elevatorList[0].currentFloor = 1;
		battery.columnList[2].elevatorList[0].direction = "up";

		//Elevator C2 at 23rd, going "up" to 28th
		battery.columnList[2].elevatorList[1].currentFloor = 23;
		battery.columnList[2].elevatorList[1].direction = "up";

		//Elevator C3 at 33th, going "down" to 1st
		battery.columnList[2].elevatorList[2].currentFloor = 33;
		battery.columnList[2].elevatorList[2].direction = "down";

		//Elevator C4 at 40th, going "down" to 24th
		battery.columnList[2].elevatorList[3].currentFloor = 40;
		battery.columnList[2].elevatorList[3].direction = "down";

		//Elevator C5 at 39th, going "down" to 1st
		battery.columnList[2].elevatorList[4].currentFloor = 39;
		battery.columnList[2].elevatorList[4].direction = "down";

		//1st, someone is at 1st floor and requests the 36th floor, elevator C1 is expected to be sent
		var requestedFloor = 36;
		battery.assignElevator(requestedFloor);

   }else if (scenario == "3") {
		//Column D
		//Elevator D1 at 58th Floor, going "down" to 1st
		battery.columnList[3].elevatorList[0].currentFloor = 58;
		battery.columnList[3].elevatorList[0].direction = "down";

		//Elevator D2 at 50th, going "up" to 60th
		battery.columnList[3].elevatorList[1].currentFloor = 50;
		battery.columnList[3].elevatorList[1].direction = "up";

		//Elevator D3 at 46th, going "up" to 58th
		battery.columnList[3].elevatorList[2].currentFloor = 46;
		battery.columnList[3].elevatorList[2].direction = "up";

		//Elevator D4 at 1st, going "down" to 54th
		battery.columnList[3].elevatorList[3].currentFloor = 1;
		battery.columnList[3].elevatorList[3].direction = "up";

		//Elevator D5 at 60th, going "down" to 1st
		battery.columnList[3].elevatorList[4].currentFloor = 60;
		battery.columnList[3].elevatorList[4].direction = "down";

		//1st, someone is at 54th floor and requests the 1st floor, elevator D1 is expected to be sent
		var floorNumber = 54;
		battery.requestElevator(floorNumber);

	} else if (scenario == "4") {
		//Column A
		//Elevator A1 at B4th Floor, Being "idle"
		battery.columnList[0].elevatorList[0].currentFloor = -4;
		battery.columnList[0].elevatorList[0].direction = "idle";

		//Elevator A2 "idle" at 1st Floor
		battery.columnList[0].elevatorList[1].currentFloor = 1;
		battery.columnList[0].elevatorList[1].direction = "idle";

		//Elevator A3 at B3rd, going "down" to B5th
		battery.columnList[0].elevatorList[2].currentFloor = -3;
		battery.columnList[0].elevatorList[2].direction = "down";

		//Elevator A4 at B6th, going "up" to 1st
		battery.columnList[0].elevatorList[3].currentFloor = -6;
		battery.columnList[0].elevatorList[3].direction = "up";

		//Elevator A5 at B1st, going "down" to b6th
		battery.columnList[0].elevatorList[4].currentFloor = -1;
		battery.columnList[0].elevatorList[4].direction = "down";

		//1st, someone is at B3rd floor and requests the 1st floor, elevator A4 is expected to be sent
		var numberFloor = -3;
		battery.requestElevator(numberFloor);

	}

     void retry(){
      Console.WriteLine("Want to retry? y/n");
      string retryInput = Console.ReadLine();
        while (retryInput != "y" && retryInput != "n"){
          Console.WriteLine("Please choose between y for Yes or n for no.");
          retryInput = Console.ReadLine();
        }
        if (retryInput == "y"){
          Main();
        } else if(retryInput == "n")
        Console.WriteLine("Thank you. Exiting program.");
    }
	retry(); 
        
              
    }
}
