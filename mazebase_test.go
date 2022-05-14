package main

import "testing"

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

	idx := 0
	for x := 0; x < 2; x++ {
		for y := 0; y < 3; y++ {
			currentCell := cells[idx]
			if currentCell.pos.x != x || currentCell.pos.y != y {
				t.Errorf("Expected cell %d position (x:%d,y:%d), actual (x:%d,y:%d)",
					idx, x, y, currentCell.pos.x, currentCell.pos.y)
			}
			idx += 1
		}
	}
}
