package main

type direction int

const (
	None direction = iota
	Up
	Down
	Left
	Right
)

type position struct {
	x int
	y int
}

type cell struct {
	pos   position
	walls []direction
}

func generateCells(width int, height int) []cell {
	cells := []cell{}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			newCell := cell{
				pos: position{
					x: x,
					y: y,
				},
				walls: []direction{},
			}
			cells = append(cells, newCell)
		}
	}
	return cells
}
