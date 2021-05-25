package app;



public class Elevator{
    public String name;
    public int currentFloor;
    public String direction;
    public Integer score;
  
public Elevator(String name, int currentFloor, String direction, Integer score){
    this.name = name;
    this.currentFloor = currentFloor;
    this.direction = direction;
    this.score = score;
  }

    // 														*****FUNCTION MOVE ELEVATOR TO FLOORNUMBER*****
public void move(int floorNumber, Elevator bestElev) {
    System.out.print("\n" + bestElev.name + " is sent. " + bestElev.name + " is currently at floor : " + bestElev.currentFloor);
        while (bestElev.currentFloor < floorNumber ){
            bestElev.direction = "up";
            bestElev.currentFloor++;
        System.out.print("\n" + bestElev.name + " is at floor : " + bestElev.currentFloor);
        }
        while (bestElev.currentFloor > floorNumber){
            bestElev.direction = "down";
            bestElev.currentFloor--;
        System.out.print("\n" + bestElev.name + " is at floor : " + bestElev.currentFloor);
        }
    
    System.out.print("\n============================================================");
    System.out.print("\n" +bestElev.name + "is arrived at destination. Doors open. You enter.");
    System.out.print("\n============================================================");
        bestElev.direction = "idle";
    
    }
    
    //														*****FUNCTION MOVE ELEVATOR TO DESTINATION*****
    //													(2 functions only because print is slightly different)
    public void moveD(int requestedFloor, Elevator bestElev) {
        System.out.print(" \nDoors closes. " + bestElev.name + " starts moving.");
        while (bestElev.currentFloor < requestedFloor ){
            bestElev.direction = "up";
            bestElev.currentFloor++;
            if (bestElev.currentFloor != 0 ){
                System.out.print("\n" + bestElev.name + " is at floor : " + bestElev.currentFloor);
            }
    
        }
        while (bestElev.currentFloor > requestedFloor) {
            bestElev.direction = "down";
            bestElev.currentFloor--;
            if (bestElev.currentFloor != 0 ){
            System.out.print("\n" +bestElev.name + " is at floor : "+ bestElev.currentFloor);
            }
    
        }
    
    System.out.print("\n============================================================");
    System.out.print("\n" + bestElev.name + " is arrived at destination. Doors open. You exit.");
    System.out.print("\n============================================================");
        bestElev.direction = "idle";
    }
    }

