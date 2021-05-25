package main
import "strconv"

type battery struct {
	columnList        []column
	numberColumn      int
	floorNumbersTotal int
}

//                                                        *****INITIALISING BATTERY*****
func newBattery(elevatorNumbersTotal, numberFloorsTotal, numberColumn, numberOfBasement int) *battery {
	b := battery{}
	b.numberColumn = numberColumn
	b.floorNumbersTotal = numberFloorsTotal
	elevatorsPerColumn := elevatorNumbersTotal / numberColumn
	floorsPerColumn := numberFloorsTotal / (numberColumn - 1)

	for i := 1; i <= numberColumn; i++ {
		col := column{
			name:         "Column " + string(i+64),
			elevatorList: []elevator{},
			floorList:    []int{},
		}
		for j := 0; j < elevatorsPerColumn; j++ {
			elev := elevator{
				name:         "elevator " + string(i+64) + strconv.Itoa(j+1),
				currentFloor: 1,
				direction:    "idle"}
			col.elevatorList = append(col.elevatorList, elev)
		}
		b.columnList = append(b.columnList, col)
	}

	for i := 0; i < numberColumn; i++ {
		for k := 1; k <= floorsPerColumn; k++ {
			if i == 0 {
				if k == 1 {
					b.columnList[i].floorList = append(b.columnList[i].floorList, k)
				}
				if k <= numberOfBasement {
					b.columnList[i].floorList = append(b.columnList[i].floorList, -k)
				}
			} else if i == 1 {
				b.columnList[i].floorList = append(b.columnList[i].floorList, k)
			} else {
				if k == 1 {
					b.columnList[i].floorList = append(b.columnList[i].floorList, 1)
				}

				b.columnList[i].floorList = append(b.columnList[i].floorList, k+((i-1)*floorsPerColumn))

			}
		}
	}
	return &b
}


func (b *battery) requestElevator(floorNumber int) {
	if floorNumber != 1 {
		requestedFloor := 1
		b.findColumn(floorNumber, requestedFloor)
	}

}

func (b *battery) assignElevator(requestedFloor int) {
	if requestedFloor != 1 {
		floorNumber := 1
		b.findColumn(floorNumber, requestedFloor)
	}
}

//                                                        *****FIND BEST COLUMN *****
func (b *battery) findColumn(floorNumber, requestedFloor int) *column {
	bestCol := column{}
	for i := 0; i < len(b.columnList); i++ {
		for _, k := range b.columnList[i].floorList {
			if floorNumber == 1 && requestedFloor == k {
				bestCol := b.columnList[i]
				bestCol.findElevator(floorNumber, requestedFloor)
				return &bestCol
			} else if floorNumber == k && requestedFloor == 1 {
				bestCol := b.columnList[i]
				bestCol.findElevator(floorNumber, requestedFloor)
				return &bestCol
			}
		}
	}
	bestCol.findElevator(floorNumber, requestedFloor)
	return &bestCol
}