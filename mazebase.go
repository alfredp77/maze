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
	x uint32
	y uint32
}

type cell struct {
	pos   position
	walls [4]cellWall
}

type cellWall struct {
	dir    direction
	isOpen bool
}

func generateCells(width uint32, height uint32) []cell {
	cells := []cell{}
	for x := uint32(0); x < width; x++ {
		for y := uint32(0); y < height; y++ {
			newCell := createCell(x, y)
			cells = append(cells, newCell)
		}
	}
	return cells
}

func createCell(x uint32, y uint32) cell {
	return cell{
		pos: position{
			x, y,
		},
		walls: [4]cellWall{
			createWall(Up),
			createWall(Down),
			createWall(Left),
			createWall(Right),
		},
	}
}

func createWall(dir direction) cellWall {
	return cellWall{dir: dir, isOpen: false}
}
