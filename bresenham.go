package main

import (
	"fmt"
)

//  Octant            y
//                   ^
//                   |
//               \ 3 | 2 /
//                \  |  /
//         4       \ | /       1
//                  \|/
// ------------------*------------------> x
//                  /|\
//         5       / | \       8
//                /  |  \
//               / 6 | 7 \

// If line goes from point p1 to p2 then
// line goes from coordinate (p1x, p1y) to (p2x, p2y)
//
// Line is in specific octant if coordinates satisfy:
// Octant 1: p1x <= p2x, p1y <= p2y, (p2y-p1y) <= (p2x-p1x)
// Octant 2: p1x <= p2x, p1y <= p2y, (p2x-p1x) <= (p2y-p1y)
// Octant 3: p1x >= p2x, p1y <= p2y, (p1x-p2x) <= (p2y-p1y)
// Octant 4: p1x >= p2x, p1y <= p2y, (p2y-p1y) <= (p1x-p2x)
// Octant 5: p1x >= p2x, p1y >= p2y, (p1y-p2y) <= (p1x-p2x)
// Octant 6: p1x >= p2x, p1y >= p2y, (p1x-p2x) <= (p1y-p2y)
// Octant 7: p1x <= p2x, p1y >= p2y, (p2x-p1x) <= (p1y-p2y)
// Octant 8: p1x <= p2x, p1y >= p2y, (p2y-p1y) <= (p2x-p1x)
//
// Note: A line can, with definitions above, be considered belong to
//       two octant at the same time if the line is on the boundary
//       between the both octants.

// Bresenham is a multipurpose function that plots coordinates along a line
// from point/coordinate p1 to p2 and applies the function point on each point it plots.
//
// The point function can, but is not limited to, draw a pixel in an image.
func Bresenham(p1x, p1y, p2x, p2y int, point func(x, y int)) {
	if (p1x <= p2x) && (p1y <= p2y) && ((p2y - p1y) <= (p2x - p1x)) {
		bresenhamOctant1(p1x, p1y, p2x, p2y, point)
	} else if (p1x <= p2x) && (p1y <= p2y) && ((p2x - p1x) <= (p2y - p1y)) {
		bresenhamOctant2(p1x, p1y, p2x, p2y, point)
	} else if (p1x >= p2x) && (p1y <= p2y) && ((p1x - p2x) <= (p2y - p1y)) {
		bresenhamOctant3(p1x, p1y, p2x, p2y, point)
	} else if (p1x >= p2x) && (p1y <= p2y) && ((p2y - p1y) <= (p1x - p2x)) {
		bresenhamOctant4(p1x, p1y, p2x, p2y, point)
	} else if (p1x >= p2x) && (p1y >= p2y) && ((p1y - p2y) <= (p1x - p2x)) {
		bresenhamOctant5(p1x, p1y, p2x, p2y, point)
	} else if (p1x >= p2x) && (p1y >= p2y) && ((p1x - p2x) <= (p1y - p2y)) {
		bresenhamOctant6(p1x, p1y, p2x, p2y, point)
	} else if (p1x <= p2x) && (p1y >= p2y) && ((p2x - p1x) <= (p1y - p2y)) {
		bresenhamOctant7(p1x, p1y, p2x, p2y, point)
	} else if (p1x <= p2x) && (p1y >= p2y) && ((p2y - p1y) <= (p2x - p1x)) {
		bresenhamOctant8(p1x, p1y, p2x, p2y, point)
	} else {
		fmt.Printf("Unknown case: (%d, %d) --> (%d, %d)\n", p1x, p1y, p2x, p2y)
	}
}

func bresenhamOctant1(p1x, p1y, p2x, p2y int, point func(x, y int)) {
	bresenhamLine(p1x, p1y, p2x, p2y, point, func(x, y int) (int, int) { return x, y })
}

func bresenhamOctant2(p1x, p1y, p2x, p2y int, point func(x, y int)) {
	p1x, p1y, p2x, p2y = p1y, p1x, p2y, p2x
	bresenhamLine(p1x, p1y, p2x, p2y, point, func(x, y int) (int, int) { return y, x })
}

func bresenhamOctant3(p1x, p1y, p2x, p2y int, point func(x, y int)) {
	p1x, p1y, p2x, p2y = p1y, -p1x, p2y, -p2x
	bresenhamLine(p1x, p1y, p2x, p2y, point, func(x, y int) (int, int) { return -y, x })
}

func bresenhamOctant4(p1x, p1y, p2x, p2y int, point func(x, y int)) {
	p1x, p1y, p2x, p2y = -p1x, p1y, -p2x, p2y
	bresenhamLine(p1x, p1y, p2x, p2y, point, func(x, y int) (int, int) { return -x, y })
}

func bresenhamOctant5(p1x, p1y, p2x, p2y int, point func(x, y int)) {
	p1x, p1y, p2x, p2y = -p1x, -p1y, -p2x, -p2y
	bresenhamLine(p1x, p1y, p2x, p2y, point, func(x, y int) (int, int) { return -x, -y })
}

func bresenhamOctant6(p1x, p1y, p2x, p2y int, point func(x, y int)) {
	p1x, p1y, p2x, p2y = -p1y, -p1x, -p2y, -p2x
	bresenhamLine(p1x, p1y, p2x, p2y, point, func(x, y int) (int, int) { return -y, -x })
}

func bresenhamOctant7(p1x, p1y, p2x, p2y int, point func(x, y int)) {
	p1x, p1y, p2x, p2y = -p1y, p1x, -p2y, p2x
	bresenhamLine(p1x, p1y, p2x, p2y, point, func(x, y int) (int, int) { return y, -x })
}

func bresenhamOctant8(p1x, p1y, p2x, p2y int, point func(x, y int)) {
	p1x, p1y, p2x, p2y = p1x, -p1y, p2x, -p2y
	bresenhamLine(p1x, p1y, p2x, p2y, point, func(x, y int) (int, int) { return x, -y })
}

// bresenhamLine plot an integer coordinate path from (x1, y1) to (x2, y2).
// Conditions: x1<=x2, y1<=y2, 0 <= slope <=1
//
//	Line is in first octant and
//	increase from lower left (x1, y1) to upper right (x2, y2).
//
// Source of information:
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// https://www.geeksforgeeks.org/bresenhams-line-generation-algorithm/
func bresenhamLine(p1x, p1y, p2x, p2y int, point func(x, y int), transform func(x, y int) (int, int)) {
	mNew := 2 * (p2y - p1y)
	slopeErrorNew := mNew - (p2x - p1x)

	y := p1y
	for x := p1x; x <= p2x; x++ {
		point(transform(x, y))

		// Add the slope to increment the angle formed
		slopeErrorNew += mNew

		// Slope error reached limit, time to increment y and update slope error.
		if slopeErrorNew > 0 {
			y++
			slopeErrorNew -= 2 * (p2x - p1x)
		}
	}
}
