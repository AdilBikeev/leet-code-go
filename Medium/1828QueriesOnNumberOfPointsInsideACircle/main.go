package QueriesOnNumberOfPointsInsideACircle

import "math"

// You are given an array points where points[i] = [xi, yi] is the coordinates of the ith point on a 2D plane. Multiple points can have the same coordinates.

// You are also given an array queries where queries[j] = [xj, yj, rj] describes a circle centered at (xj, yj) with a radius of rj.

// For each query queries[j], compute the number of points inside the jth circle. Points on the border of the circle are considered inside.

// Return an array answer, where answer[j] is the answer to the jth query.

// Example 1:

// Input: points = [[1,3],[3,3],[5,3],[2,2]], queries = [[2,3,1],[4,3,1],[1,1,2]]
// Output: [3,2,2]
// Explanation: The points and circles are shown above.
// queries[0] is the green circle, queries[1] is the red circle, and queries[2] is the blue circle.

type Circle struct {
	x0 int
	y0 int
	r  int
}

func NewCircle(x0, y0, r int) *Circle {
	return &Circle{
		x0: x0,
		y0: y0,
		r:  r,
	}
}

func (c *Circle) IsBelongsCircle(x, y int) bool {
	return math.Pow(float64(x)-float64(c.x0), 2)+
		math.Pow(float64(y)-float64(c.y0), 2) <= math.Pow(float64(c.r), 2)
}

func parseCircles(queries [][]int, n int) []Circle {
	circles := make([]Circle, n)
	for i := 0; i < n; i++ {
		x := queries[i][0]
		y := queries[i][1]
		r := queries[i][2]
		circles[i] = *NewCircle(x, y, r)
	}

	return circles
}

func countPoints(points [][]int, queries [][]int) []int {
	n := len(queries)
	countInnerPointsRes := make([]int, n)
	circles := parseCircles(queries, n)

	// Пробегаемся по каждому кругу
	for i := 0; i < n; i++ {
		circle := circles[i]
		//Проверяем сколько точек лежат на круге
		for _, point := range points {
			x := point[0]
			y := point[1]
			if circle.IsBelongsCircle(x, y) {
				countInnerPointsRes[i]++
			}
		}
	}

	return countInnerPointsRes
}
