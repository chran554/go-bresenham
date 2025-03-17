package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type point struct{ x, y int }

func pt(x, y int) point { return point{x: x, y: y} }

// TestBresenhamOctant tests that the plot of a line between two points
// in any of the 8 octants, produce the expected result.
func TestBresenhamOctant(t *testing.T) {
	testCases := []struct {
		name           string
		p1, p2         point
		expectedPoints []point
	}{
		{"Octant 1", pt(0, 0), pt(10, 5), []point{{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 2}, {5, 3}, {6, 3}, {7, 4}, {8, 4}, {9, 5}, {10, 5}}},
		{"Octant 2", pt(0, 0), pt(5, 10), []point{{0, 0}, {1, 1}, {1, 2}, {2, 3}, {2, 4}, {3, 5}, {3, 6}, {4, 7}, {4, 8}, {5, 9}, {5, 10}}},
		{"Octant 3", pt(0, 0), pt(-5, 10), []point{{0, 0}, {-1, 1}, {-1, 2}, {-2, 3}, {-2, 4}, {-3, 5}, {-3, 6}, {-4, 7}, {-4, 8}, {-5, 9}, {-5, 10}}},
		{"Octant 4", pt(0, 0), pt(-10, 5), []point{{0, 0}, {-1, 1}, {-2, 1}, {-3, 2}, {-4, 2}, {-5, 3}, {-6, 3}, {-7, 4}, {-8, 4}, {-9, 5}, {-10, 5}}},
		{"Octant 5", pt(0, 0), pt(-10, -5), []point{{0, 0}, {-1, -1}, {-2, -1}, {-3, -2}, {-4, -2}, {-5, -3}, {-6, -3}, {-7, -4}, {-8, -4}, {-9, -5}, {-10, -5}}},
		{"Octant 6", pt(0, 0), pt(-5, -10), []point{{0, 0}, {-1, -1}, {-1, -2}, {-2, -3}, {-2, -4}, {-3, -5}, {-3, -6}, {-4, -7}, {-4, -8}, {-5, -9}, {-5, -10}}},
		{"Octant 7", pt(0, 0), pt(5, -10), []point{{0, 0}, {1, -1}, {1, -2}, {2, -3}, {2, -4}, {3, -5}, {3, -6}, {4, -7}, {4, -8}, {5, -9}, {5, -10}}},
		{"Octant 8", pt(0, 0), pt(10, -5), []point{{0, 0}, {1, -1}, {2, -1}, {3, -2}, {4, -2}, {5, -3}, {6, -3}, {7, -4}, {8, -4}, {9, -5}, {10, -5}}},
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

// TestBresenhamOctantBoundaries tests that either one of the two possible bresenham octant functions
// can be used and produce the same result for a line exactly on the boundary line between two octants.
func TestBresenhamOctantBoundaries(t *testing.T) {
	testCases := []struct {
		name           string
		bresenhamFn1   func(int, int, int, int, func(x, y int))
		bresenhamFn2   func(int, int, int, int, func(x, y int))
		p1, p2         point
		expectedPoints []point
	}{
		{"Octant boundary 8&1 at   0°", bresenhamOctant8, bresenhamOctant1, pt(0, 0), pt(10, 0), []point{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0}, {9, 0}, {10, 0}}},
		{"Octant boundary 1&2 at  45°", bresenhamOctant1, bresenhamOctant2, pt(0, 0), pt(10, 10), []point{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}, {10, 10}}},
		{"Octant boundary 2&3 at  90°", bresenhamOctant2, bresenhamOctant3, pt(0, 0), pt(0, 10), []point{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}, {0, 8}, {0, 9}, {0, 10}}},
		{"Octant boundary 3&4 at 135°", bresenhamOctant3, bresenhamOctant4, pt(0, 0), pt(-10, 10), []point{{0, 0}, {-1, 1}, {-2, 2}, {-3, 3}, {-4, 4}, {-5, 5}, {-6, 6}, {-7, 7}, {-8, 8}, {-9, 9}, {-10, 10}}},
		{"Octant boundary 4&5 at 180°", bresenhamOctant4, bresenhamOctant5, pt(0, 0), pt(-10, 0), []point{{0, 0}, {-1, 0}, {-2, 0}, {-3, 0}, {-4, 0}, {-5, 0}, {-6, 0}, {-7, 0}, {-8, 0}, {-9, 0}, {-10, 0}}},
		{"Octant boundary 5&6 at 225°", bresenhamOctant5, bresenhamOctant6, pt(0, 0), pt(-10, -10), []point{{0, 0}, {-1, -1}, {-2, -2}, {-3, -3}, {-4, -4}, {-5, -5}, {-6, -6}, {-7, -7}, {-8, -8}, {-9, -9}, {-10, -10}}},
		{"Octant boundary 6&7 at 270°", bresenhamOctant6, bresenhamOctant7, pt(0, 0), pt(0, -10), []point{{0, 0}, {0, -1}, {0, -2}, {0, -3}, {0, -4}, {0, -5}, {0, -6}, {0, -7}, {0, -8}, {0, -9}, {0, -10}}},
		{"Octant boundary 7&8 at 315°", bresenhamOctant7, bresenhamOctant8, pt(0, 0), pt(10, -10), []point{{0, 0}, {1, -1}, {2, -2}, {3, -3}, {4, -4}, {5, -5}, {6, -6}, {7, -7}, {8, -8}, {9, -9}, {10, -10}}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var accumulatedPoints []point
			var accumulatedPointsFn1 []point
			var accumulatedPointsFn2 []point

			// Line plotted points using Bresenham choice of octant function
			Bresenham(testCase.p1.x, testCase.p1.y, testCase.p2.x, testCase.p2.y, func(x, y int) {
				accumulatedPoints = append(accumulatedPoints, pt(x, y))
			})

			// Line plotted points using function 1 of 2 possible octant functions
			testCase.bresenhamFn1(testCase.p1.x, testCase.p1.y, testCase.p2.x, testCase.p2.y, func(x, y int) {
				accumulatedPointsFn1 = append(accumulatedPointsFn1, pt(x, y))
			})

			// Line plotted points using function 2 of 2 possible octant functions
			testCase.bresenhamFn2(testCase.p1.x, testCase.p1.y, testCase.p2.x, testCase.p2.y, func(x, y int) {
				accumulatedPointsFn2 = append(accumulatedPointsFn2, pt(x, y))
			})

			assert.Equal(t, testCase.p1, accumulatedPoints[0], "bresenham: first plotted point is not equal to p1")
			assert.Equal(t, testCase.p2, accumulatedPoints[len(accumulatedPoints)-1], "bresenham: last plotted point is not equal to p2")
			assert.Equal(t, testCase.expectedPoints, accumulatedPoints, "bresenham: all plotted points in the line are not exactly as expected")

			assert.Equal(t, testCase.p1, accumulatedPointsFn1[0], "bresenhamFn1: first plotted point is not equal to p1")
			assert.Equal(t, testCase.p2, accumulatedPointsFn1[len(accumulatedPointsFn1)-1], "bresenhamFn1: last plotted point is not equal to p2")
			assert.Equal(t, testCase.expectedPoints, accumulatedPointsFn1, "bresenhamFn1: all plotted points in the line are not exactly as expected")

			assert.Equal(t, testCase.p1, accumulatedPointsFn2[0], "bresenhamFn2: first plotted point is not equal to p1")
			assert.Equal(t, testCase.p2, accumulatedPointsFn2[len(accumulatedPointsFn2)-1], "bresenhamFn2: last plotted point is not equal to p2")
			assert.Equal(t, testCase.expectedPoints, accumulatedPointsFn2, "bresenhamFn2: all plotted points in the line are not exactly as expected")
		})
	}
}
