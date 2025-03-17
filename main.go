package main

import (
	"fmt"
	"image"
)

func main() {
	printBresenhamPlotCoordinate := func(x, y int) { fmt.Printf("(%d,%d) ", x, y) }

	p1 := image.Pt(0, 0)
	p2 := image.Pt(10, 5)
	fmt.Printf("Print plotted coordinates from %+v to %+v using bresenham algorithm.\n", p1, p2)

	Bresenham(p1.X, p1.Y, p2.X, p2.Y, printBresenhamPlotCoordinate)
	fmt.Println()
}
