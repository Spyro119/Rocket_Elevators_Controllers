using System;

namespace Class.ElevatorObject{

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
}