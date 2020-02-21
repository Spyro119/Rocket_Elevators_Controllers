/*                            ***** RUN JS IN https://repl.it/@SamuelJubinvill/residentialController   ******
                                                  Also, sorry for the inputs ;) 
*/

class Column {
  constructor(numberFloors, numberElevators){
      this.elevatorList = [];
      this.floorsList = [];

//                               ***** SETTING FLOORLIST *****
      for(let i = 1; i <= numberFloors ; i++){
          this.floorsList.push(i);
      }
//        ***** SETTING ELEVATORLIST AND ELEVATORS DIRECTION AND ELEVATORS POSITION *****
      for(var i = 1 ; i <= numberElevators; i++){
          var currentFloor = parseInt(prompt(" Where is Elevator " + [i] + "? floor :", ""));
            while(isNaN(currentFloor) || currentFloor < 1 || currentFloor > numberFloors){
              currentFloor = parseInt(prompt(" please enter a valid number for elevator's floor"));
            }
            var direction = String(prompt("Where is it heading", "")).toLowerCase();
            while (direction != "up" && direction != "down" && direction != "null" && direction != "idle"){
              direction = String(prompt("Please choose between up, down or null/idle"));
            }
            if(direction == "idle"){
              direction = "null";
            }
            let id = "Elevator " + parseInt(i) ;
            var elevator = new Elevator(currentFloor, direction, id);
            this.elevatorList.push(elevator);
            }
      }
  
//                    ***** REQUEST ELEVATOR FROM USER FLOOR *****

requestElevator(requestFloor, requestDirection){
  var requestFloor = parseInt(prompt("What floor are you requesting the elevator?"));
  while(isNaN(requestFloor) || requestFloor < 1 || requestFloor > numberFloors){
  requestFloor = parseInt(prompt("Please enter a valid number for request floor"));
  }
var requestDirection = String(prompt("Going up or down?")).toLowerCase();
  while (requestDirection != "up" && requestDirection != "down" ){
    requestDirection = String(prompt("Please choose between up or down values")).toLowerCase();
  }
column1.findElevator(requestFloor, requestDirection);
}

//                            ***** FIND ELEVATOR *****

findElevator(requestFloor, requestDirection){
this.elevatorList.forEach(function(elevator){
  if(elevator.currentFloor == requestFloor && requestDirection == elevator.direction){
    elevator.elevatorScore = 0;
  }
  else if(elevator.currentFloor < requestFloor && requestDirection == elevator.direction && elevator.direction == "up"){
    elevator.elevatorScore = 10;
  }
  else if(elevator.currentFloor > requestFloor && requestDirection == "down" && elevator.direction == requestDirection){
    elevator.elevatorScore = 20;
  }
  else if(elevator.currentFloor == requestFloor && elevator.direction == "null"){
    elevator.elevatorScore = 30;
  }
  else if(elevator.direction == "null"){
    elevator.elevatorScore = 40;
  }
  else if(elevator.currentFloor == requestFloor && elevator.direction != requestDirection){
    elevator.elevatorScore = 50;
  }
  else if(elevator.currentFloor != requestFloor && elevator.direction == requestDirection){
    elevator.elevatorScore = 60;
  }
  else if(elevator.currentFloor != requestFloor && elevator.direction != requestDirection && elevator.direction != "null"){
    elevator.elevatorScore = 70;
  }
  console.log(elevator.name + ' score is :' + elevator.elevatorScore);
})
  column1.elevatorList.sort(compare);
  var bestElevator = column1.elevatorList[0];
  bestElevator.move(requestFloor);

  function compare(a, b) {
    if(a.elevatorScore == b.elevatorScore){
     console.log("Recalculating score depending on elevators' position...");
     column1.elevatorList.forEach(function(elevator){
       var floorsDiff = Math.abs(elevator.currentFloor - requestFloor);
       elevator.elevatorScore = elevator.elevatorScore + floorsDiff;
       console.log(elevator.name + ' score is now :' + elevator.elevatorScore);
     });
   }

   
   return a.elevatorScore - b.elevatorScore;
   
}  
}

//                *****FUNCTION TO FIND DESTINATION*****

destination(requestFloor){
  requestFloor = parseInt(prompt("What floor are you going?"));
  while (isNaN(requestFloor) || requestFloor < 0 || requestFloor > numberFloors){
    requestFloor = parseInt(prompt("please enter a valid number."));
  }
  var elevator = this.elevatorList[0];
  elevator.moveD(requestFloor);
}

}

class Elevator{
  constructor(currentFloor, direction, id, elevatorScore){
      this.requestFloor = [];
      this.direction = direction;
      this.currentFloor = currentFloor;
      this.name = id;
      this.elevatorScore = elevatorScore;
      
      }

//                         ***** MOVE ELEVATOR TO USER *****
  move(requestFloor){
    console.log(this.name + " Is sent. " + this.name + " Is currently at floor : " + this.currentFloor);
    while(this.currentFloor < requestFloor){
      this.direction = "up";
      this.currentFloor++; 
      console.log(this.name + " is at floor: " + this.currentFloor);
    }
    while(this.currentFloor > requestFloor){
      this.direction = "down";
      this.currentFloor--; 
      console.log(this.name + " at floor: " + this.currentFloor);
    }
    console.log("==============================================");
    console.log(this.name +" is arrived. Doors open. You enter.");
    console.log("Doors closed.")
    this.direction = "null";
    column1.destination(requestFloor);
  }

//                  ***** MOVE ELEVATOR TO DESTINATION *****

  moveD(requestFloor){
    while(this.currentFloor < requestFloor){
      this.direction = "up";
      this.currentFloor++; 
      console.log(this.name + " at floor: " + this.currentFloor);
    }
    while(this.currentFloor > requestFloor){
      this.direction = "down";
      this.currentFloor--; 
      console.log(this.name + " at floor: " + this.currentFloor);
    }
    console.log("=============================================");
    console.log(this.name +" is arrived. Doors open. You exit.");
    this.direction = "null";
    retry();
}
}

//               ***** SETTING SITUATION PARAMETERS *****

var numberElevators = parseInt(prompt("How many elevators?"));
while (isNaN(numberElevators) || numberElevators < 1) {
numberElevators = prompt("please enter a valid number of elevator", "");}
numberFloors = prompt("How many floors?");
while(isNaN(numberFloors) || numberFloors < 1){
numberFloors = parseInt(prompt("please enter a valid number of Floors", ""));
}

//              ***** RETRY PROGRAM WITH SAME PARAMETERS *****
function retry(){
var accept = String(prompt("Do you want to continue running this programm with same values? y/n")).toLowerCase();
if(accept == "y"){
  this.elevatorScore=('');
  column1.requestFloor();
}
if(accept == "n"){
  console.log("exiting program. Thank you.");
} else {
  accept = String(prompt("Please type 'y' for yes or 'n' for no."));
}
}

let column1 = new Column(numberFloors, numberElevators);

column1.requestElevator();