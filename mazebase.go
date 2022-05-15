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

type maze struct {
	cells  []cell
	width  uint32
	height uint32
}

func generateMazeBase(width uint32, height uint32) maze {
	cells := []cell{}
	for y := uint32(0); y < height; y++ {
		for x := uint32(0); x < width; x++ {
			newCell := createCell(x, y)
			cells = append(cells, newCell)
		}
	}
	return maze{
		cells, width, height,
	}
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

func (m maze) getCell(x uint32, y uint32) cell {
	idx := m.getCellIndex(x, y)
	return m.cells[idx]
}

func (m maze) getCellIndex(x uint32, y uint32) uint32 {
	return x + y*m.width
}
