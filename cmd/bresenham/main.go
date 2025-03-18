package main

import (
	"fmt"
	"github.com/chran554/go-bresenham"
)

func main() {
	fmt.Println("Plot a line of integer coordinates from (3,3) to (10,15):")

	printBresenhamPlotCoordinate := func(x, y int) { fmt.Printf("(%d,%d) ", x, y) }
	bresenham.Bresenham(3, 3, 15, 10, printBresenhamPlotCoordinate)
	fmt.Println()
}
