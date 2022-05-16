package main

type direction int

const (
	None direction = iota
	Up
	Down
	Left
	Right
)

var Directions []direction = []direction{Up, Down, Left, Right}

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

func (d direction) getOpposite() direction {
	if d == Up {
		return Down
	}
	if d == Down {
		return Up
	}
	if d == Left {
		return Right
	}

	return Left
}

func (c cell) getWall(dir direction) cellWall {
	return c.walls[dir-1]
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

func (m maze) getCellNeighbour(c cell, dir direction) *cell {
	if dir == Up {
		if c.pos.y > 0 {
			found := m.getCell(c.pos.x, c.pos.y-1)
			return &found
		}
	}

	if dir == Down {
		if c.pos.y < m.height-1 {
			found := m.getCell(c.pos.x, c.pos.y+1)
			return &found
		}
	}

	if dir == Left {
		if c.pos.x > 0 {
			found := m.getCell(c.pos.x-1, c.pos.y)
			return &found
		}
	}

	if dir == Right {
		if c.pos.x < m.width-1 {
			found := m.getCell(c.pos.x+1, c.pos.y)
			return &found
		}
	}

	return nil
}
