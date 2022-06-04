package main

import (
	"fmt"
	"math/rand"
	"time"
)

func (m maze) generatePaths() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	x := random.Intn(int(m.width) - 1)
	y := random.Intn(int(m.height) - 1)
	randoms := make([]byte, m.getCellCount()*4)
	random.Read(randoms)
	currentRandomIdx := 0

	startCell := m.getCell(uint32(x), uint32(y))
	pendingCells := []*cell{startCell}
	visitedFlags := m.createCellFlags()

	for len(pendingCells) > 0 {
		lastIndex := len(pendingCells) - 1
		currentCell := pendingCells[lastIndex]
		visitedFlags[m.getCellIndex(currentCell.pos.x, currentCell.pos.y)] = true

		unvisitedNeighbours := m.getUnvisitedNeighbours(currentCell, visitedFlags)
		idx := len(unvisitedNeighbours) - 1
		if idx < 0 {
			pendingCells = pendingCells[:lastIndex]
		} else {
			if idx > 0 {
				idx = (int)(randoms[currentRandomIdx]) % len(unvisitedNeighbours)
				currentRandomIdx++
			}
			m.connect(currentCell, unvisitedNeighbours[idx])
			pendingCells = append(pendingCells, unvisitedNeighbours[idx])
		}
	}
}

func (m maze) createCellFlags() []bool {
	result := []bool{}
	cellCount := m.width * m.height
	for i := uint32(0); i < cellCount; i++ {
		result = append(result, false)
	}
	return result
}

func (m maze) getUnvisitedNeighbours(c *cell, visitedFlags []bool) []*cell {
	result := []*cell{}
	for _, dir := range ValidDirections {
		neighbour := m.getCellNeighbour(c, dir)
		if neighbour == nil {
			continue
		}

		idx := m.getCellIndex(neighbour.pos.x, neighbour.pos.y)
		if !visitedFlags[idx] {
			result = append(result, neighbour)
		}
	}
	return result
}

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
	if cell1 == nil || cell2 == nil {
		return false, None
	}
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

func (m maze) getConnectedNeighbours(cell1 *cell) []*cell {
	result := []*cell{}
	for _, dir := range ValidDirections {
		neighbour := m.getCellNeighbour(cell1, dir)
		areConnected, _ := m.areConnected(cell1, neighbour)
		if areConnected {
			result = append(result, neighbour)
		}
	}
	return result
}

func (m maze) connect(cell1 *cell, cell2 *cell) {
	isConnected, dir := m.areConnected(cell1, cell2)
	if !isConnected {
		cell1.setWall(createOpenWall(dir))
		cell2.setWall(createOpenWall(dir.getOpposite()))
	}
}
