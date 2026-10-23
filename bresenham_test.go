package bresenham

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type point struct{ x, y int }

func pt(x, y int) point { return point{x: x, y: y} }

func TestBresenham(t *testing.T) {
	testCases := []struct {
		name           string
		p1, p2         point
		expectedPoints []point
	}{
		{"single point", pt(42, 43), pt(42, 43), []point{{42, 43}}},
		{"test line 1", pt(3, 3), pt(15, 10), []point{{3, 3}, {4, 4}, {5, 4}, {6, 5}, {7, 5}, {8, 6}, {9, 6}, {10, 7}, {11, 8}, {12, 8}, {13, 9}, {14, 9}, {15, 10}}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var accumulatedPoints []point

			Bresenham(testCase.p1.x, testCase.p1.y, testCase.p2.x, testCase.p2.y, func(x, y int) {
				accumulatedPoints = append(accumulatedPoints, pt(x, y))
			})

			assert.Equal(t, testCase.p1, accumulatedPoints[0], "first plotted point is not equal to p1")
			assert.Equal(t, testCase.p2, accumulatedPoints[len(accumulatedPoints)-1], "last plotted point is not equal to p2")
			assert.Equal(t, testCase.expectedPoints, accumulatedPoints, "all plotted points in the line are not exactly as expected")
		})
	}
}

// TestBresenhamOctant tests that the plot of a line between two points
// in any of the 8 octants, produce the expected result.
func TestBresenhamOctant(t *testing.T) {
	testCases := []struct {
		name           string
		p1, p2         point
		expectedPoints []point
	}{
		{"Octant 1", pt(0, 0), pt(10, 5), []point{{0, 0}, {1, 0}, {2, 1}, {3, 1}, {4, 2}, {5, 2}, {6, 3}, {7, 3}, {8, 4}, {9, 4}, {10, 5}}},
		{"Octant 2", pt(0, 0), pt(5, 10), []point{{0, 0}, {0, 1}, {1, 2}, {1, 3}, {2, 4}, {2, 5}, {3, 6}, {3, 7}, {4, 8}, {4, 9}, {5, 10}}},
		{"Octant 3", pt(0, 0), pt(-5, 10), []point{{0, 0}, {0, 1}, {-1, 2}, {-1, 3}, {-2, 4}, {-2, 5}, {-3, 6}, {-3, 7}, {-4, 8}, {-4, 9}, {-5, 10}}},
		{"Octant 4", pt(0, 0), pt(-10, 5), []point{{0, 0}, {-1, 0}, {-2, 1}, {-3, 1}, {-4, 2}, {-5, 2}, {-6, 3}, {-7, 3}, {-8, 4}, {-9, 4}, {-10, 5}}},
		{"Octant 5", pt(0, 0), pt(-10, -5), []point{{0, 0}, {-1, 0}, {-2, -1}, {-3, -1}, {-4, -2}, {-5, -2}, {-6, -3}, {-7, -3}, {-8, -4}, {-9, -4}, {-10, -5}}},
		{"Octant 6", pt(0, 0), pt(-5, -10), []point{{0, 0}, {0, -1}, {-1, -2}, {-1, -3}, {-2, -4}, {-2, -5}, {-3, -6}, {-3, -7}, {-4, -8}, {-4, -9}, {-5, -10}}},
		{"Octant 7", pt(0, 0), pt(5, -10), []point{{0, 0}, {0, -1}, {1, -2}, {1, -3}, {2, -4}, {2, -5}, {3, -6}, {3, -7}, {4, -8}, {4, -9}, {5, -10}}},
		{"Octant 8", pt(0, 0), pt(10, -5), []point{{0, 0}, {1, 0}, {2, -1}, {3, -1}, {4, -2}, {5, -2}, {6, -3}, {7, -3}, {8, -4}, {9, -4}, {10, -5}}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var accumulatedPoints []point

			Bresenham(testCase.p1.x, testCase.p1.y, testCase.p2.x, testCase.p2.y, func(x, y int) {
				accumulatedPoints = append(accumulatedPoints, pt(x, y))
			})

			assert.Equal(t, testCase.p1, accumulatedPoints[0], "first plotted point is not equal to p1")
			assert.Equal(t, testCase.p2, accumulatedPoints[len(accumulatedPoints)-1], "last plotted point is not equal to p2")
			assert.Equal(t, testCase.expectedPoints, accumulatedPoints, "all plotted points in the line are not exactly as expected", accumulatedPoints)
		})
	}
}

// TestBresenhamOctantBoundaries tests lines exactly on the boundary between
// two octants.
func TestBresenhamOctantBoundaries(t *testing.T) {
	testCases := []struct {
		name           string
		p1, p2         point
		expectedPoints []point
	}{
		{"Octant boundary 8&1 at   0°", pt(0, 0), pt(10, 0), []point{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0}, {9, 0}, {10, 0}}},
		{"Octant boundary 1&2 at  45°", pt(0, 0), pt(10, 10), []point{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}, {10, 10}}},
		{"Octant boundary 2&3 at  90°", pt(0, 0), pt(0, 10), []point{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}, {0, 8}, {0, 9}, {0, 10}}},
		{"Octant boundary 3&4 at 135°", pt(0, 0), pt(-10, 10), []point{{0, 0}, {-1, 1}, {-2, 2}, {-3, 3}, {-4, 4}, {-5, 5}, {-6, 6}, {-7, 7}, {-8, 8}, {-9, 9}, {-10, 10}}},
		{"Octant boundary 4&5 at 180°", pt(0, 0), pt(-10, 0), []point{{0, 0}, {-1, 0}, {-2, 0}, {-3, 0}, {-4, 0}, {-5, 0}, {-6, 0}, {-7, 0}, {-8, 0}, {-9, 0}, {-10, 0}}},
		{"Octant boundary 5&6 at 225°", pt(0, 0), pt(-10, -10), []point{{0, 0}, {-1, -1}, {-2, -2}, {-3, -3}, {-4, -4}, {-5, -5}, {-6, -6}, {-7, -7}, {-8, -8}, {-9, -9}, {-10, -10}}},
		{"Octant boundary 6&7 at 270°", pt(0, 0), pt(0, -10), []point{{0, 0}, {0, -1}, {0, -2}, {0, -3}, {0, -4}, {0, -5}, {0, -6}, {0, -7}, {0, -8}, {0, -9}, {0, -10}}},
		{"Octant boundary 7&8 at 315°", pt(0, 0), pt(10, -10), []point{{0, 0}, {1, -1}, {2, -2}, {3, -3}, {4, -4}, {5, -5}, {6, -6}, {7, -7}, {8, -8}, {9, -9}, {10, -10}}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var accumulatedPoints []point

			Bresenham(testCase.p1.x, testCase.p1.y, testCase.p2.x, testCase.p2.y, func(x, y int) {
				accumulatedPoints = append(accumulatedPoints, pt(x, y))
			})

			assert.Equal(t, testCase.p1, accumulatedPoints[0], "bresenham: first plotted point is not equal to p1")
			assert.Equal(t, testCase.p2, accumulatedPoints[len(accumulatedPoints)-1], "bresenham: last plotted point is not equal to p2")
			assert.Equal(t, testCase.expectedPoints, accumulatedPoints, "bresenham: all plotted points in the line are not exactly as expected")
		})
	}
}

func TestBresenhamIntegerBoundaries(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	testCases := []struct {
		name           string
		p1, p2         point
		expectedPoints []point
	}{
		{"maximum horizontal endpoint", pt(maxInt-1, 0), pt(maxInt, 0), []point{{maxInt - 1, 0}, {maxInt, 0}}},
		{"minimum horizontal endpoint", pt(minInt, 0), pt(minInt+1, 0), []point{{minInt, 0}, {minInt + 1, 0}}},
		{"maximum vertical endpoint", pt(0, maxInt-1), pt(0, maxInt), []point{{0, maxInt - 1}, {0, maxInt}}},
		{"minimum vertical endpoint", pt(0, minInt), pt(0, minInt+1), []point{{0, minInt}, {0, minInt + 1}}},
		{"maximum point", pt(maxInt, maxInt), pt(maxInt, maxInt), []point{{maxInt, maxInt}}},
		{"minimum point", pt(minInt, minInt), pt(minInt, minInt), []point{{minInt, minInt}}},
		{"near maximum diagonal", pt(maxInt-1, maxInt-1), pt(maxInt, maxInt), []point{{maxInt - 1, maxInt - 1}, {maxInt, maxInt}}},
		{"near minimum diagonal", pt(minInt, minInt), pt(minInt+1, minInt+1), []point{{minInt, minInt}, {minInt + 1, minInt + 1}}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var accumulatedPoints []point

			Bresenham(testCase.p1.x, testCase.p1.y, testCase.p2.x, testCase.p2.y, func(x, y int) {
				accumulatedPoints = append(accumulatedPoints, pt(x, y))
			})

			assert.Equal(t, testCase.expectedPoints, accumulatedPoints)
		})
	}
}

func TestUnsignedDistance(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	assert.Equal(t, ^uint(0), unsignedDistance(minInt, maxInt))
	assert.Equal(t, uint(0), unsignedDistance(minInt, minInt))
	assert.Equal(t, uint(1), unsignedDistance(maxInt-1, maxInt))
}

func TestBresenhamSmallGrid(t *testing.T) {
	const minCoordinate = -3
	const maxCoordinate = 3

	for p1x := minCoordinate; p1x <= maxCoordinate; p1x++ {
		for p1y := minCoordinate; p1y <= maxCoordinate; p1y++ {
			for p2x := minCoordinate; p2x <= maxCoordinate; p2x++ {
				for p2y := minCoordinate; p2y <= maxCoordinate; p2y++ {
					p1, p2 := pt(p1x, p1y), pt(p2x, p2y)
					points := collectPoints(p1, p2)

					if len(points) == 0 {
						t.Fatalf("line from %v to %v produced no points", p1, p2)
					}
					if points[0] != p1 {
						t.Fatalf("line from %v to %v started at %v", p1, p2, points[0])
					}
					if points[len(points)-1] != p2 {
						t.Fatalf("line from %v to %v ended at %v", p1, p2, points[len(points)-1])
					}

					expectedCount := absInt(p2x-p1x) + 1
					if dy := absInt(p2y - p1y); dy+1 > expectedCount {
						expectedCount = dy + 1
					}
					if len(points) != expectedCount {
						t.Fatalf("line from %v to %v produced %d points, want %d", p1, p2, len(points), expectedCount)
					}

					for i := 1; i < len(points); i++ {
						dx := absInt(points[i].x - points[i-1].x)
						dy := absInt(points[i].y - points[i-1].y)
						if dx > 1 || dy > 1 || (dx == 0 && dy == 0) {
							t.Fatalf("line from %v to %v has invalid step from %v to %v", p1, p2, points[i-1], points[i])
						}
					}
				}
			}
		}
	}
}

func TestBresenhamTranslatedAndReversed(t *testing.T) {
	testCases := []struct {
		name string
		p1   point
		p2   point
	}{
		{"translated octant one", pt(7, -4), pt(19, 3)},
		{"translated octant two", pt(7, -4), pt(12, 15)},
		{"translated octant three", pt(7, -4), pt(-3, 11)},
		{"translated octant four", pt(7, -4), pt(-8, 2)},
		{"translated octant five", pt(7, 8), pt(-5, 1)},
		{"translated octant six", pt(7, 8), pt(2, -5)},
		{"translated octant seven", pt(-7, 8), pt(4, -5)},
		{"translated octant eight", pt(-7, 8), pt(8, 2)},
		{"reversed line", pt(19, 3), pt(7, -4)},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			points := collectPoints(testCase.p1, testCase.p2)
			if len(points) == 0 {
				t.Fatal("line produced no points")
			}
			assert.Equal(t, testCase.p1, points[0])
			assert.Equal(t, testCase.p2, points[len(points)-1])
		})
	}
}

func collectPoints(p1, p2 point) []point {
	var points []point

	Bresenham(p1.x, p1.y, p2.x, p2.y, func(x, y int) {
		points = append(points, pt(x, y))
	})

	return points
}

func absInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
