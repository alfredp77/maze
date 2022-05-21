package main

import "fmt"

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

func (m maze) getCellNeighbour(c *cell, dir direction) *cell {
	if dir == Up {
		if c.pos.y > 0 {
			return m.getCell(c.pos.x, c.pos.y-1)
		}
	}

	if dir == Down {
		if c.pos.y < m.height-1 {
			return m.getCell(c.pos.x, c.pos.y+1)
		}
	}

	if dir == Left {
		if c.pos.x > 0 {
			return m.getCell(c.pos.x-1, c.pos.y)
		}
	}

	if dir == Right {
		if c.pos.x < m.width-1 {
			return m.getCell(c.pos.x+1, c.pos.y)
		}
	}

	return nil
}

func (m maze) areNeighbours(cell1 *cell, cell2 *cell) direction {
	for _, dir := range ValidDirections {
		neighbour := m.getCellNeighbour(cell1, dir)
		if neighbour != nil {
			fmt.Println(dir.Name(), "cell2 found", *neighbour)
		} else {
			fmt.Println(dir.Name(), "nil!")
		}
		if neighbour != nil && neighbour == cell2 {
			return dir
		}
	}
	return None
}

func (m maze) areConnected(cell1 *cell, cell2 *cell) (bool, direction) {
	dir := m.areNeighbours(cell1, cell2)
	if dir != None {
		wall1 := cell1.getWall(dir)
		wall2 := cell2.getWall(dir.getOpposite())

		if wall1.isOpen && wall2.isOpen {
			return true, dir
		} else {
			return false, dir
		}
	}
	return false, None
}

func (m maze) connect(cell1 *cell, cell2 *cell) {
	isConnected, dir := m.areConnected(cell1, cell2)
	if !isConnected {
		cell1.setWall(createOpenWall(dir))
		cell2.setWall(createOpenWall(dir.getOpposite()))
	}
}
