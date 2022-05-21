package main

import (
	"testing"
)

var theMaze maze = generateMazeBase(4, 3)

func runNoNeighbourTest(t *testing.T, dir direction, cellPos position) {
	noNeighbour := theMaze.getCellNeighbour(theMaze.getCell(cellPos.x, cellPos.y), dir)
	if noNeighbour != nil {
		t.Errorf("Expected no %v neighbour for (%d,%d)", dir.Name(), cellPos.x, cellPos.y)
	}
}

func runGetCellNeighbourTest(t *testing.T, dir direction, cellPos position, neighbourPos position) {
	neighbour := theMaze.getCellNeighbour(theMaze.getCell(cellPos.x, cellPos.y), dir)
	expected := theMaze.getCell(neighbourPos.x, neighbourPos.y)
	if neighbour == nil {
		t.Errorf("Expected to have %v neighbour for (%d,%d) but found none", dir.Name(), cellPos.x, cellPos.y)
	} else if neighbour != expected {
		t.Errorf("Expected %v neighbour for (%d,%d) is %v, actual %v", dir.Name(), cellPos.x, cellPos.y, expected, *neighbour)
	}
}

func TestGetCellNeighbour(t *testing.T) {
	runNoNeighbourTest(t, Up, position{0, 0})
	runGetCellNeighbourTest(t, Up, position{1, 1}, position{1, 0})

	runNoNeighbourTest(t, Down, position{1, 2})
	runGetCellNeighbourTest(t, Down, position{0, 1}, position{0, 2})

	runNoNeighbourTest(t, Left, position{0, 1})
	runGetCellNeighbourTest(t, Left, position{2, 1}, position{1, 1})

	runNoNeighbourTest(t, Right, position{3, 0})
	runGetCellNeighbourTest(t, Right, position{0, 2}, position{1, 2})
}

func runAreNeighboursTest(t *testing.T, expectedDir direction, cell1Pos position, cell2Pos position) {
	cell1 := theMaze.getCell(cell1Pos.x, cell1Pos.y)
	cell2 := theMaze.getCell(cell2Pos.x, cell2Pos.y)

	dir := theMaze.areNeighbours(cell1, cell2)

	if dir != expectedDir {
		t.Errorf("Expected %v, actual %v", expectedDir.Name(), dir.Name())
	}
}

func TestAreNeighbours(t *testing.T) {
	runAreNeighboursTest(t, Up, position{2, 2}, position{2, 1})
	runAreNeighboursTest(t, Down, position{2, 1}, position{2, 2})
	runAreNeighboursTest(t, Left, position{2, 2}, position{1, 2})
	runAreNeighboursTest(t, Right, position{1, 2}, position{2, 2})
	runAreNeighboursTest(t, None, position{2, 0}, position{3, 2})
}

func TestAreConnected_NotConnectedNeighbours(t *testing.T) {
	cell1 := theMaze.getCell(2, 2)
	cell2 := theMaze.getCell(3, 2)

	isConnected, dir := theMaze.areConnected(cell1, cell2)
	if isConnected {
		t.Error("Expected not connected, actual is connected")
	}
	if dir != Right {
		t.Errorf("Expected %v, actual %v", Right.Name(), dir.Name())
	}
}

func TestAreConnected_ConnectedNeighbours(t *testing.T) {
	cell1 := theMaze.getCell(2, 2)
	cell2 := theMaze.getCell(3, 2)

	theMaze.connect(cell1, cell2)

	isConnected, dir := theMaze.areConnected(cell1, cell2)
	if !isConnected {
		t.Error("Expected connected, actual is not connected")
	}
	if dir != Right {
		t.Errorf("Expected %v, actual %v", Right.Name(), dir.Name())
	}
}
