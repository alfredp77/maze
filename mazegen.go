package main

// func (m maze) generatePaths() {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	random := rand.New(source)

// 	x := random.Intn(int(m.width) - 1)
// 	y := random.Intn(int(m.height) - 1)

// 	startCell := m.getCell(uint32(x), uint32(y))
// 	pendingCells := []cell{startCell}
// 	dirs := []direction{Up, Down, Left, Right}

// 	for len(pendingCells) > 0 {
// 		lastIndex := len(pendingCells) - 1
// 		currentCell := pendingCells[lastIndex]
// 		pendingCells = pendingCells[:lastIndex]

// 	}
// }

func (m maze) isConnected(cell1 cell, cell2 cell) direction {
	for _, dir := range Directions {
		neighbour := m.getCellNeighbour(cell1, dir)
		if neighbour != nil && *neighbour == cell2 {
			wall1 := cell1.getWall(dir)
			wall2 := cell2.getWall(dir.getOpposite())

			if wall1.isOpen && wall2.isOpen {
				return dir
			} else {
				return None
			}
		}
	}

	return None
}
