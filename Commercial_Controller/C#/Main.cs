using System;
using Rocket_Elevators_Commercial_Controller.Commercial_Controller.Classes;

namespace Rocket_Elevators_Commercial_Controller.Commercial_Controller
{
    internal class Program
    {
        static void Main()
        {
            var elevatorNumbersTotal = 20;
            var floorNumbersTotal = 60;
            var numberOfBasement = 6;
            var numberOfColumn = 4;
            var program = new Program();

            //                                *****CREATING BATTERY ******
            Battery battery = new Battery(elevatorNumbersTotal, numberOfColumn, floorNumbersTotal, numberOfBasement);

            Console.WriteLine("Please choose a scenario: 1, 2, 3 or 4");
            int.TryParse(Console.ReadLine(), out int scenario);
            while (scenario <= 0 || scenario > 4)
            {
                Console.WriteLine("Please choose a scenario: 1, 2, 3 or 4");
                int.TryParse(Console.ReadLine(), out scenario);
            }
            switch (scenario)
            {
                case 1:
                    program.scenario_one(battery);
                    break;
                case 2:
                    program.scenario_two(battery);
                    break;
                case 3:
                    program.scenario_three(battery);
                    break;
                case 4:
                    program.scenario_four(battery);
                    break;
                default:
                    break;
            }

            program.Retry();
        }

        void Retry()
        {
            Console.WriteLine("Want to retry? y/n");
            string retryInput = Console.ReadLine();
            while (retryInput != "y" && retryInput != "n")
            {
                Console.WriteLine("Please choose between y for Yes or n for no.");
                retryInput = Console.ReadLine();
            }
            if (retryInput == "y")
            {
                Main();
            }
            else if (retryInput == "n")
                Console.WriteLine("Thank you. Exiting program.");
        }


        public void scenario_one(Battery battery)
        {
            //Elevator B1 at 20th Floor, going "Down" to 5th
            battery.ColumnList[1].ElevatorList[0].CurrentFloor = 20;
            battery.ColumnList[1].ElevatorList[0].Direction = "Down";

            //Elevator B2 at 3rd, going "Up" to 15th
            battery.ColumnList[1].ElevatorList[1].CurrentFloor = 3;
            battery.ColumnList[1].ElevatorList[1].Direction = "Up";

            //Elevator B3 at 13th, going "Down" to 1st
            battery.ColumnList[1].ElevatorList[2].CurrentFloor = 13;
            battery.ColumnList[1].ElevatorList[2].Direction = "Down";

            //Elevator B4 at 15th, going "Down" to 2nd
            battery.ColumnList[1].ElevatorList[3].CurrentFloor = 15;
            battery.ColumnList[1].ElevatorList[3].Direction = "Down";

            //Elevator B5 at 6th, going "Down" to 1st
            battery.ColumnList[1].ElevatorList[4].CurrentFloor = 6;
            battery.ColumnList[1].ElevatorList[4].Direction = "Down";

            //1st, someone is at 1st floor and requests the 20th floor
            var requestedFloor = 20;
            battery.AssignElevator(requestedFloor);
        }

        public void scenario_two(Battery battery)
        {
            //Column C
            //Elevator C1 at 1ts Floor, going "Up" to 21st
            battery.ColumnList[2].ElevatorList[0].CurrentFloor = 1;
            battery.ColumnList[2].ElevatorList[0].Direction = "Up";

            //Elevator C2 at 23rd, going "Up" to 28th
            battery.ColumnList[2].ElevatorList[1].CurrentFloor = 23;
            battery.ColumnList[2].ElevatorList[1].Direction = "Up";

            //Elevator C3 at 33th, going "Down" to 1st
            battery.ColumnList[2].ElevatorList[2].CurrentFloor = 33;
            battery.ColumnList[2].ElevatorList[2].Direction = "Down";

            //Elevator C4 at 40th, going "Down" to 24th
            battery.ColumnList[2].ElevatorList[3].CurrentFloor = 40;
            battery.ColumnList[2].ElevatorList[3].Direction = "Down";

            //Elevator C5 at 39th, going "Down" to 1st
            battery.ColumnList[2].ElevatorList[4].CurrentFloor = 39;
            battery.ColumnList[2].ElevatorList[4].Direction = "Down";

            //1st, someone is at 1st floor and requests the 36th floor, elevator 3-1 is expected to be sent
            var requestedFloor = 36;
            battery.AssignElevator(requestedFloor);
        }

        public void scenario_three(Battery battery)
        {
            //Column D
            //Elevator D1 at 58th Floor, going "Down" to 1st
            battery.ColumnList[3].ElevatorList[0].CurrentFloor = 58;
            battery.ColumnList[3].ElevatorList[0].Direction = "Down";

            //Elevator D2 at 50th, going "Up" to 60th
            battery.ColumnList[3].ElevatorList[1].CurrentFloor = 50;
            battery.ColumnList[3].ElevatorList[1].Direction = "Up";

            //Elevator D3 at 46th, going "Up" to 58th
            battery.ColumnList[3].ElevatorList[2].CurrentFloor = 46;
            battery.ColumnList[3].ElevatorList[2].Direction = "Up";

            //Elevator D4 at 1st, going "Down" to 54th
            battery.ColumnList[3].ElevatorList[3].CurrentFloor = 1;
            battery.ColumnList[3].ElevatorList[3].Direction = "Up";

            //Elevator D5 at 60th, going "Down" to 1st
            battery.ColumnList[3].ElevatorList[4].CurrentFloor = 60;
            battery.ColumnList[3].ElevatorList[4].Direction = "Down";

            //1st, someone is at 54th floor and requests the 1st floor, elevator 4-1 is expected to be sent
            var floorNumber = 54;
            battery.RequestElevator(floorNumber);

        }

        public void scenario_four(Battery battery)
        {
            //Column A
            //Elevator A1 at B4th Floor, Being "Idle"
            battery.ColumnList[0].ElevatorList[0].CurrentFloor = -4;
            battery.ColumnList[0].ElevatorList[0].Direction = "Idle";

            //Elevator A2 "Idle" at 1st Floor
            battery.ColumnList[0].ElevatorList[1].CurrentFloor = 1;
            battery.ColumnList[0].ElevatorList[1].Direction = "Idle";

            //Elevator A3 at B3rd, going "Down" to B5th
            battery.ColumnList[0].ElevatorList[2].CurrentFloor = -3;
            battery.ColumnList[0].ElevatorList[2].Direction = "Down";

            //Elevator A4 at B6th, going "Up" to 1st
            battery.ColumnList[0].ElevatorList[3].CurrentFloor = -6;
            battery.ColumnList[0].ElevatorList[3].Direction = "Up";

            //Elevator A5 at B1st, going "Down" to b6th
            battery.ColumnList[0].ElevatorList[4].CurrentFloor = -1;
            battery.ColumnList[0].ElevatorList[4].Direction = "Down";

            //1st, someone is at B3rd floor and requests the 1st floor, elevator 1-4 is expected to be sent
            var numberFloor = -3;
            battery.RequestElevator(numberFloor);
        }
    }
}
