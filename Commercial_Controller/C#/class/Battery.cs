using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Rocket_Elevators_Commercial_Controller.Commercial_Controller.Classes
{
    internal class Battery
    {
        public int ColumnCount;
        public int FloorsCount;
        public List<Column> ColumnList = new List<Column>();
        public Column SelectedColumn = null;

        public Battery(int elevatorsCount, int columnCount, int floorsCount, int basementFloorCount)
        {
            this.ColumnCount = columnCount;
            this.FloorsCount = floorsCount;
            var numberOfElevators = elevatorsCount / this.ColumnCount;
            int numberOfFloors = this.FloorsCount / (this.ColumnCount - 1);

            for (var i = 1; i <= columnCount; i++)
            {
                Column column = new Column("Column " + i);
                this.ColumnList.Add(column);
                    for (var j = 1; j <= numberOfElevators; j++)
                {
                    column.AddElevator(new Elevator("Elevator " + i + "-" + j, 1, "Idle"));
                }
            }

            for (var i = 0; i < columnCount; i++)
            {
                if (i == 0)
                {
                    this.ColumnList[i].FloorsMin = -basementFloorCount;
                    this.ColumnList[i].FloorsMax = 1;
                }
                else
                {
                    this.ColumnList[i].FloorsMin = (i - 1) * numberOfFloors + 1;
                    this.ColumnList[i].FloorsMax = i * numberOfFloors;
                }
            }
        }

        public void RequestElevator(int floorNumber)
        {
            if (floorNumber != 1)
            {
                int requestedFloor = 1;
                this.FindColumn(floorNumber, requestedFloor);

            }
        }
        public void AssignElevator(int requestedFloor)
        {
            if (requestedFloor != 1)
            {
                int floorNumber = 1;
                this.FindColumn(floorNumber, requestedFloor);
            }
        }


        Column FindColumn(int requestedFloor, int destinationFloor)
        {
            if (destinationFloor < 0) this.SelectedColumn = this.ColumnList[0];
            else
            {
                for (var i = 0; i < this.ColumnList.Count; i++)
                {
                    if (requestedFloor == 1 && destinationFloor <= this.ColumnList[i].FloorsMax && destinationFloor >= this.ColumnList[i].FloorsMin) this.SelectedColumn = this.ColumnList[i];
                    else if (requestedFloor <= this.ColumnList[i].FloorsMax && requestedFloor >= this.ColumnList[i].FloorsMin && destinationFloor == 1) this.SelectedColumn = this.ColumnList[i];
                }
            }
            
            if (this.SelectedColumn == null) throw new Exception("Aucune colonne n'a été sélectionné.");
            this.SelectedColumn.FindElevator(requestedFloor, destinationFloor);
            return this.SelectedColumn;
        }
    }
}
