using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Rocket_Elevators_Commercial_Controller.Commercial_Controller.Classes
{
    public class Elevator
    {
        public string Name;
        public int CurrentFloor;
        public string Direction;
        public int Score = -1;

        public Elevator(string name, int currentFloor, string direction)
        {
            this.Name = name;
            this.CurrentFloor = currentFloor;
            this.Direction = direction;
        }

        public void CalculateScore(int requestedFloor, int floorNumber, int floorsDiff)
        {
            if (requestedFloor >= 1 && floorNumber >= 1)
            {
                //                              *****ELEVATOR AT SAME FLOOR AS USER AND Direction EITHER Up OR Idle*****
                if (this.CurrentFloor == floorNumber)
                {
                    this.Score = floorsDiff;
                    //                              *****IF USER AT FLOOR 1 AND ELEVATOR IS GOING Down *****
                }
                else if (floorNumber == 1 && this.Direction == "Down")
                {
                    this.Score = floorsDiff + 100;
                    //                              *****IF USER AT FLOOR 1 AND ELEVATOR IS Idle *****
                }
                else if (floorNumber == 1 && this.Direction == "Idle")
                {
                    this.Score = floorsDiff + 200;
                    //                              *****IF USER AT FLOOR 1 AND ELEVATOR IS GOING Up *****
                }
                else if (floorNumber == 1 && this.Direction == "Up")
                {
                    this.Score = floorsDiff + 300;
                    //                              *****IF USER IS GOING Down, ELEVATOR IS > USER FLOOR AND ELEVATOR IS GOING Down *****
                }
                else if (floorNumber > 1 && this.CurrentFloor > floorNumber && this.Direction == "Down")
                {
                    this.Score = floorsDiff + 400;
                    //                              *****IF USER IS GOING Down, ELEVATOR IS < USER FLOOR AND ELEVATOR IS Idle *****
                }
                else if (floorNumber > 1 && this.CurrentFloor > floorNumber && this.Direction == "Idle")
                {
                    this.Score = floorsDiff + 500;
                    //                              *****IF USER IS GOING Down, ELEVATOR IS > USER AND Direction IS Up *****
                }
                else if (floorNumber > 1 && this.CurrentFloor > floorNumber && this.Direction == "Up")
                {
                    this.Score = floorsDiff + 600;
                    //                              *****IF USER IS GOING Down, ELEVATOR IS < USER *****
                }
                else if (floorNumber > 1 && this.CurrentFloor < floorNumber)
                {
                    this.Score = floorsDiff + 700;
                }

                //                                  *****FIND BEST ELEVATOR FOR BASEMENTS*****
            } else if (requestedFloor <= 1 && floorNumber <= 1)
            {
                //                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR IS AT 1ST FLOOR*****
                if (this.Direction == "Idle" && this.CurrentFloor == floorNumber)
                {
                    this.Score = floorsDiff;
                    //                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR GOES Up*****
                }
                else if (floorNumber == 1 && this.Direction == "Up")
                {
                    this.Score = floorsDiff + 100;
                    //                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR IS Idle*****
                }
                else if (floorNumber == 1 && this.Direction == "Idle")
                {
                    this.Score = floorsDiff + 200;
                    //                              *****IF USER IS AT 1ST FLOOR AND ELEVATOR GOES Down*****
                }
                else if (floorNumber == 1 && this.Direction == "Down")
                {
                    this.Score = floorsDiff + 300;
                    //                              *****USER IS <1 AND ELEVATOR IS UNDER USER FLOOR AND ELEVATOR GOES Up*****
                }
                else if (floorNumber < 1 && this.CurrentFloor < floorNumber && this.Direction == "Up")
                {
                    this.Score = floorsDiff + 400;
                    //                              *****USER IS <1 AND ELEVATOR IS UNDER USER FLOOR AND ELEVATOR IS Idle*****
                }
                else if (floorNumber < 1 && this.CurrentFloor < floorNumber && this.Direction == "Idle")
                {
                    this.Score = floorsDiff + 500;
                    //                              *****USER IS <1 AND ELEVATOR IS UNDER USER FLOOR AND ELEVATOR GOES Down*****
                }
                else if (floorNumber < 1 && this.CurrentFloor < floorNumber && this.Direction == "Down")
                {
                    this.Score = floorsDiff + 600;
                    //                              *****USER IS <1 AND ELEVATOR IS ABOVE USER FLOOR *****
                }
                else if (floorNumber < 1 && this.CurrentFloor > floorNumber)
                {
                    this.Score = floorsDiff + 700;
                    //                              *****USER IS <1 AND ELEVATOR Direction IS Down*****
                }
                else if (floorNumber < 1 && this.Direction == "Down")
                {
                    this.Score = floorsDiff + 800;
                }
            } else
            {
                this.Score = floorsDiff + 1000;
            }
        }

        /// <summary>
        /// Function to move the elevator.
        /// </summary>
        /// <param name="requestedFloor">Floor where the elvator is requested.</param>
        /// <param name="destinationFloor">Destination floor.</param>
        public void Move(int requestedFloor, int destinationFloor)
        {
            Console.WriteLine($"{this.Name} has been sent. {this.Name} is currently at floor : {this.CurrentFloor}");
            if (this.CurrentFloor < requestedFloor) this.MoveUp(requestedFloor);
            else if (this.CurrentFloor > requestedFloor) this.MoveDown(requestedFloor);
            Console.WriteLine("============================================================\n" +
                               this.Name + " is arrived at destination. Doors open. You enter.\n" +
                               "============================================================");
            this.Direction = "Idle";

            // sleep there
            Console.WriteLine("Doors closes. ", this.Name, " starts moving.");
            if (this.CurrentFloor < destinationFloor) this.MoveUp(destinationFloor);
            else this.MoveDown(destinationFloor);

            Console.WriteLine("============================================================\n" +
                               this.Name + " is arrived at destination. Doors open. You exit.\n" +
                               "============================================================");
            this.Direction = "Idle";
        }

        /// <summary>
        /// Function to move the elevator upwards.
        /// </summary>
        /// <param name="destinationFloor">Destination floor.</param>
        private void MoveUp(int destinationFloor)
        {
            this.Direction = "Up";
            while (this.CurrentFloor < destinationFloor)
            {
                this.CurrentFloor++;
                Console.WriteLine($"{this.Name} is at floor : {this.CurrentFloor}.");
            }
        }

        /// <summary>
        /// Function to move the elevator downwards
        /// </summary>
        /// <param name="destinationFloor">Destination floor</param>
        private void MoveDown(int destinationFloor)
        {
            this.Direction = "Down";
            while (this.CurrentFloor > destinationFloor)
            {
                this.CurrentFloor--;
                Console.WriteLine($"{this.Name} is at floor : {this.CurrentFloor}.");
            }
        }
    }
}
