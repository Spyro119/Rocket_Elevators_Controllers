package classes

import "strconv"

type Battery struct {
	ColumnList        []Column
	NumberColumn      int
	FloorNumbersTotal int
}

//                                                        *****INITIALISING Battery*****
func NewBattery(elevatorNumbersTotal, numberFloorsTotal, NumberColumn, numberOfBasement int) *Battery {
	b := Battery{}
	b.NumberColumn = NumberColumn
	b.FloorNumbersTotal = numberFloorsTotal
	elevatorsPerColumn := elevatorNumbersTotal / NumberColumn
	floorsPerColumn := numberFloorsTotal / (NumberColumn - 1)

	for i := 1; i <= NumberColumn; i++ {
		col := Column{
			Name:         "Column " + string(i+64),
			ElevatorList: []Elevator{},
			FloorList:    []int{},
		}
		for j := 0; j < elevatorsPerColumn; j++ {
			elev := Elevator{
				Name:         "elevator " + string(i+64) + strconv.Itoa(j+1),
				CurrentFloor: 1,
				Direction:    "idle"}
			col.ElevatorList = append(col.ElevatorList, elev)
		}
		b.ColumnList = append(b.ColumnList, col)
	}

	for i := 0; i < NumberColumn; i++ {
		for k := 1; k <= floorsPerColumn; k++ {
			if i == 0 {
				if k == 1 {
					b.ColumnList[i].FloorList = append(b.ColumnList[i].FloorList, k)
				}
				if k <= numberOfBasement {
					b.ColumnList[i].FloorList = append(b.ColumnList[i].FloorList, -k)
				}
			} else if i == 1 {
				b.ColumnList[i].FloorList = append(b.ColumnList[i].FloorList, k)
			} else {
				if k == 1 {
					b.ColumnList[i].FloorList = append(b.ColumnList[i].FloorList, 1)
				}

				b.ColumnList[i].FloorList = append(b.ColumnList[i].FloorList, k+((i-1)*floorsPerColumn))

			}
		}
	}
	return &b
}

func (b *Battery) RequestElevator(floorNumber int) {
	if floorNumber != 1 {
		requestedFloor := 1
		b.findColumn(floorNumber, requestedFloor)
	}

}

func (b *Battery) AssignElevator(requestedFloor int) {
	if requestedFloor != 1 {
		floorNumber := 1
		b.findColumn(floorNumber, requestedFloor)
	}
}

//                                                        *****FIND BEST Column *****
func (b *Battery) findColumn(floorNumber, requestedFloor int) *Column {
	bestCol := Column{}
	for i := 0; i < len(b.ColumnList); i++ {
		for _, k := range b.ColumnList[i].FloorList {
			if floorNumber == 1 && requestedFloor == k {
				bestCol := b.ColumnList[i]
				bestCol.findElevator(floorNumber, requestedFloor)
				return &bestCol
			} else if floorNumber == k && requestedFloor == 1 {
				bestCol := b.ColumnList[i]
				bestCol.findElevator(floorNumber, requestedFloor)
				return &bestCol
			}
		}
	}
	bestCol.findElevator(floorNumber, requestedFloor)
	return &bestCol
}
