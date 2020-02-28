package app;





public class App {
    public static void main(String[] args) throws Exception {
        var elevatorNumbersTotal = 20;
        var numberOfColumn = 4;
        var floorNumbersTotal = 60;
        var numberOfBasement = 6;
    
       Battery battery = new Battery(elevatorNumbersTotal, numberOfColumn, floorNumbersTotal, numberOfBasement);
        
       //                                      *****SCENARIOS *****

       
 /*        //Elevator B1 at 20th Floor, going "down" to 5th
       battery.columnList.get(1).elevatorList.get(0).currentFloor = 20;
       battery.columnList.get(1).elevatorList.get(0).direction = "down";

       //Elevator B2 at 3rd, going "up" to 15th
       battery.columnList.get(1).elevatorList.get(1).currentFloor = 3;
       battery.columnList.get(1).elevatorList.get(1).direction = "up";

       //Elevator B3 at 13th, going "down" to 1st
       battery.columnList.get(1).elevatorList.get(2).currentFloor = 13;
       battery.columnList.get(1).elevatorList.get(2).direction = "down";

       //Elevator B4 at 15th, going "down" to 2nd
       battery.columnList.get(1).elevatorList.get(3).currentFloor = 15;
       battery.columnList.get(1).elevatorList.get(3).direction = "down";

       //Elevator B5 at 6th, going "down" to 1st
       battery.columnList.get(1).elevatorList.get(4).currentFloor = 6;
       battery.columnList.get(1).elevatorList.get(4).direction = "down";

       //1st, someone is at 1st floor and requests the 20th floor
       var requestedFloor = 20;
       battery.assignElevator(requestedFloor); */

//                                            *****SCENARIO 2*****
         //Column C
/*       //Elevator C1 at 1ts Floor, going "up" to 21st
       battery.columnList.get(2).elevatorList.get(0).currentFloor = 1;
       battery.columnList.get(2).elevatorList.get(0).direction = "up";

       //Elevator C2 at 23rd, going "up" to 28th
       battery.columnList.get(2).elevatorList.get(1).currentFloor = 23;
       battery.columnList.get(2).elevatorList.get(1).direction = "up";

       //Elevator C3 at 33th, going "down" to 1st
       battery.columnList.get(2).elevatorList.get(2).currentFloor = 33;
       battery.columnList.get(2).elevatorList.get(2).direction = "down";

       //Elevator C4 at 40th, going "down" to 24th
       battery.columnList.get(2).elevatorList.get(3).currentFloor = 40;
       battery.columnList.get(2).elevatorList.get(3).direction = "down";

       //Elevator C5 at 39th, going "down" to 1st
       battery.columnList.get(2).elevatorList.get(4).currentFloor = 39;
       battery.columnList.get(2).elevatorList.get(4).direction = "down";

       //1st, someone is at 1st floor and requests the 36th floor, elevator C1 is expected to be sent
       var requestedFloor = 36;
       battery.assignElevator(requestedFloor); */

  //                                                *****SCENARIO 3*****
 /*      //Column D
       //Elevator D1 at 58th Floor, going "down" to 1st
       battery.columnList.get(3).elevatorList.get(0).currentFloor = 58;
       battery.columnList.get(3).elevatorList.get(0).direction = "down";

       //Elevator D2 at 50th, going "up" to 60th
       battery.columnList.get(3).elevatorList.get(1).currentFloor = 50;
       battery.columnList.get(3).elevatorList.get(1).direction = "up";

       //Elevator D3 at 46th, going "up" to 58th
       battery.columnList.get(3).elevatorList.get(2).currentFloor = 46;
       battery.columnList.get(3).elevatorList.get(2).direction = "up";

       //Elevator D4 at 1st, going "down" to 54th
       battery.columnList.get(3).elevatorList.get(3).currentFloor = 1;
       battery.columnList.get(3).elevatorList.get(3).direction = "up";

       //Elevator D5 at 60th, going "down" to 1st
       battery.columnList.get(3).elevatorList.get(4).currentFloor = 60;
       battery.columnList.get(3).elevatorList.get(4).direction = "down";

       //1st, someone is at 54th floor and requests the 1st floor, elevator D1 is expected to be sent
       var floorNumber = 54;
       battery.requestElevator(floorNumber);
*/
   //                                     *****SCENARIO***** 4
       //Column A
       //Elevator A1 at B4th Floor, Being "idle"
       battery.columnList.get(0).elevatorList.get(0).currentFloor = -4;
       battery.columnList.get(0).elevatorList.get(0).direction = "idle";

       //Elevator A2 "idle" at 1st Floor
       battery.columnList.get(0).elevatorList.get(1).currentFloor = 1;
       battery.columnList.get(0).elevatorList.get(1).direction = "idle";

       //Elevator A3 at B3rd, going "down" to B5th
       battery.columnList.get(0).elevatorList.get(2).currentFloor = -3;
       battery.columnList.get(0).elevatorList.get(2).direction = "down";

       //Elevator A4 at B6th, going "up" to 1st
       battery.columnList.get(0).elevatorList.get(3).currentFloor = -6;
       battery.columnList.get(0).elevatorList.get(3).direction = "up";

       //Elevator A5 at B1st, going "down" to b6th
       battery.columnList.get(0).elevatorList.get(4).currentFloor = -1;
       battery.columnList.get(0).elevatorList.get(4).direction = "down";

       //1st, someone is at B3rd floor and requests the 1st floor, elevator A4 is expected to be sent
       var numberFloor = -3;
       battery.requestElevator(numberFloor); 

   }
       
    }
