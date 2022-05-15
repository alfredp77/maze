package main

import (
	"math/rand"
	"time"
)

func (m maze) generatePaths() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	x := random.Intn(int(m.width) - 1)
	y := random.Intn(int(m.height) - 1)

	currentCell := m.getCell(uint32(x), uint32(y))
	pendingCells := []cell{currentCell}

	for len(pendingCells) > 0 {

	}
}
