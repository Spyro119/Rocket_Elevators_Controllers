using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Rocket_Elevators_Commercial_Controller.Commercial_Controller.Classes
{
    internal class Column
    {
        public string Name;
        public int FloorsMin;
        public int FloorsMax;
        public List<Elevator> ElevatorList = new List<Elevator>();
        public int? SelectedIndex = null;
        public Elevator SelectedElevator = null;

        public Column(string name)
        {
            this.Name = name;
        }

        public void AddElevator(Elevator elevator)
        {
            this.ElevatorList.Add(elevator);
        }

        public Elevator FindElevator(int floorNumber, int requestedFloor)
        {
            for (var i = 0; i < this.ElevatorList.Count; i++)
            {
                Elevator currentElevator = this.ElevatorList[i];
                int floorsDiff = Math.Abs(currentElevator.CurrentFloor - floorNumber);
                currentElevator.CalculateScore(requestedFloor, floorNumber, floorsDiff);
            }

            this.SelectedElevator = this.ElevatorList.Aggregate((elevator1, elevator2) => elevator1.Score < elevator2.Score ? elevator1 : elevator2);
            Console.WriteLine("Best elevator is : " + this.SelectedElevator.Name);
            this.SelectedElevator.Move(requestedFloor, floorNumber);

            return this.SelectedElevator;
        }

    }
}
