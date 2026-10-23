package bresenham

// Bresenham is a multipurpose function that plots coordinates along a line
// from point/coordinate p1 to p2 and applies the function point on each point it plots.
//
// The point function can, but is not limited to, draw a pixel in an image.
func Bresenham(p1x, p1y, p2x, p2y int, point func(x, y int)) {
	dx := unsignedDistance(p1x, p2x)
	dy := unsignedDistance(p1y, p2y)
	sx := direction(p1x, p2x)
	sy := direction(p1y, p2y)

	if dx == 0 && dy == 0 {
		point(p1x, p1y)
		return
	}

	if dx >= dy {
		bresenhamShallow(p1x, p1y, p2x, sx, sy, dx, dy, point)
		return
	}

	bresenhamSteep(p1x, p1y, p2y, sx, sy, dx, dy, point)
}

func unsignedDistance(a, b int) uint {
	// Avoid signed subtraction: the full distance between two int values fits
	// in uint, and unsigned subtraction gives that distance modulo 2^N.
	if a < b {
		return uint(b) - uint(a)
	}
	return uint(a) - uint(b)
}

func direction(from, to int) int {
	if from < to {
		return 1
	}
	if from > to {
		return -1
	}
	return 0
}

func bresenhamShallow(p1x, p1y, p2x, sx, sy int, dx, dy uint, point func(x, y int)) {
	accumulatedError := (dx - 1) / 2
	threshold := dx - dy
	x, y := p1x, p1y

	for {
		point(x, y)
		if x == p2x {
			return
		}

		if accumulatedError >= threshold {
			accumulatedError -= threshold
			y += sy
		} else {
			accumulatedError += dy
		}
		x += sx
	}
}

func bresenhamSteep(p1x, p1y, p2y, sx, sy int, dx, dy uint, point func(x, y int)) {
	accumulatedError := (dy - 1) / 2
	threshold := dy - dx
	x, y := p1x, p1y

	for {
		point(x, y)
		if y == p2y {
			return
		}

		if accumulatedError >= threshold {
			accumulatedError -= threshold
			x += sx
		} else {
			accumulatedError += dx
		}
		y += sy
	}
}
