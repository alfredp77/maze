package main

import (
	"testing"
)

func TestGenerateCells_CorrectNumberOfCells(t *testing.T) {
	cells := generateCells(5, 6)

	expected := 30
	actual := len(cells)
	if actual != expected {
		t.Errorf("Expected %d cells, actual %d", expected, actual)
	}
}

func TestGenerateCells_CorrectPositions(t *testing.T) {
	cells := generateCells(2, 3)

	idx := uint32(0)
	for x := uint32(0); x < 2; x++ {
		for y := uint32(0); y < 3; y++ {
			currentCell := cells[idx]
			if currentCell.pos.x != x || currentCell.pos.y != y {
				t.Errorf("Expected cell %d position (x:%d,y:%d), actual (x:%d,y:%d)",
					idx, x, y, currentCell.pos.x, currentCell.pos.y)
			}
			idx += 1
		}
	}
}

func TestCreateCell(t *testing.T) {
	x := uint32(2)
	y := uint32(3)
	newCell := createCell(x, y)

	if newCell.pos.x != x || newCell.pos.y != y {
		t.Errorf("Expected cell position (x:%d,y:%d), actual (x:%d,y:%d)",
			x, y, newCell.pos.x, newCell.pos.y)
	}

	verifyWall(t, cellWall{Up, false}, newCell.walls[0])
	verifyWall(t, cellWall{Down, false}, newCell.walls[1])
	verifyWall(t, cellWall{Left, false}, newCell.walls[2])
	verifyWall(t, cellWall{Right, false}, newCell.walls[3])
}

func TestCreateWall(t *testing.T) {
	wall := createWall(Down)

	verifyWall(t, cellWall{Down, false}, wall)
}

func verifyWall(t *testing.T, expected cellWall, actual cellWall) {
	if expected.dir != actual.dir || expected.isOpen != actual.isOpen {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}
