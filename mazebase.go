package main

type direction int

const (
	None direction = iota
	Up
	Down
	Left
	Right
)

var ValidDirections []direction = []direction{Up, Down, Left, Right}
var DirectionNames []string = []string{"None", "Up", "Down", "Left", "Right"}

func (d direction) Name() string {
	return DirectionNames[d]
}

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

func (c *cell) setWall(wall cellWall) {
	c.walls[wall.dir-1] = wall
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
			createClosedWall(Up),
			createClosedWall(Down),
			createClosedWall(Left),
			createClosedWall(Right),
		},
	}
}

func createClosedWall(dir direction) cellWall {
	return cellWall{dir: dir, isOpen: false}
}

func createOpenWall(dir direction) cellWall {
	return cellWall{dir: dir, isOpen: true}
}

func (m maze) getCell(x uint32, y uint32) *cell {
	idx := m.getCellIndex(x, y)
	return &m.cells[idx]
}

func (m maze) getCellIndex(x uint32, y uint32) uint32 {
	return x + y*m.width
}
