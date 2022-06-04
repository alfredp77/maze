package main

import "fmt"

func main() {
	theMaze := generateMazeBase(15, 5)
	theMaze.generatePaths()
	render(theMaze)

	fmt.Println("-- -- --")

	canvas := consoleCanvas{}
	canvas.render(theMaze)
}
