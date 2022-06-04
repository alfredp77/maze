package main

import (
	"fmt"
	"strings"
)

type consoleCanvas struct{}

var wallChar byte = '#'
var emptyChar byte = ' '
var cellSize uint32 = 3 //2x2 cell

func (c consoleCanvas) render(m maze) {
	rows := initRows(c, m)
	for x := uint32(0); x < m.width; x++ {
		for y := uint32(0); y < m.height; y++ {
			cellTop := y*2 + 1
			cellLeft := x*2 + 1
			drawCell(cellLeft, cellTop, emptyChar, rows)
			cell := m.getCell(x, y)
			if cell.getWall(Up).isOpen {
				drawCell(cellLeft, cellTop-1, emptyChar, rows)
			}
			if cell.getWall(Down).isOpen {
				drawCell(cellLeft, cellTop+1, emptyChar, rows)
			}
			if cell.getWall(Left).isOpen {
				drawCell(cellLeft-1, cellTop, emptyChar, rows)
			}
			if cell.getWall(Right).isOpen {
				drawCell(cellLeft+1, cellTop, emptyChar, rows)
			}
		}
	}

	builder := strings.Builder{}
	for _, row := range rows {
		builder.Write(row)
		builder.WriteByte('\n')
	}
	fmt.Println(builder.String())
}

func drawCell(x uint32, y uint32, value byte, rows [][]byte) {
	cellRow := y * cellSize
	cellCol := x * cellSize
	for cr := uint32(0); cr < cellSize; cr++ {
		for cc := uint32(0); cc < cellSize; cc++ {
			rows[cellRow+cr][cellCol+cc] = value
		}
	}
}

func initRows(c consoleCanvas, m maze) [][]byte {
	totalRows := (m.height*2 + 1) * cellSize
	totalCols := (m.width*2 + 1) * cellSize
	rows := make([][]byte, totalRows)
	for i := uint32(0); i < totalRows; i++ {
		rows[i] = make([]byte, totalCols)
		for j := uint32(0); j < totalCols; j++ {
			rows[i][j] = wallChar
		}
	}
	return rows
}

func render(m maze) {
	// each cell (*) is surrounded by "walls" (@)
	// @@@
	// @*@
	// @@@

	wall := "\xff"
	empty := " "
	rows := []string{}

	for y := uint32(0); y < m.height; y++ {
		topBuilder := strings.Builder{}
		topBuilder.WriteString(wall)
		for x := uint32(0); x < m.width; x++ {
			cell := m.getCell(x, y)
			if cell.getWall(Up).isOpen {
				topBuilder.WriteString(empty)
			} else {
				topBuilder.WriteString(wall)
			}
			topBuilder.WriteString(wall)
		}
		rows = append(rows, topBuilder.String())

		midBuilder := strings.Builder{}
		midBuilder.WriteString(wall)
		for x := uint32(0); x < m.width; x++ {
			cell := m.getCell(x, y)
			midBuilder.WriteString(empty)
			if cell.getWall(Right).isOpen {
				midBuilder.WriteString(empty)
			} else {
				midBuilder.WriteString(wall)
			}
		}

		rows = append(rows, midBuilder.String())
	}

	rows = append(rows, strings.Repeat(wall, 2*int(m.width+1)))

	for _, row := range rows {
		fmt.Println(row)
	}
}
