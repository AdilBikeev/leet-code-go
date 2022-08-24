package MaxIncreaseToKeepCitySkyline

import "math"

// There is a city composed of n x n blocks, where each block contains a single building shaped like a vertical square prism. You are given a 0-indexed n x n integer matrix grid where grid[r][c] represents the height of the building located in the block at row r and column c.

// A city's skyline is the the outer contour formed by all the building when viewing the side of the city from a distance. The skyline from each cardinal direction north, east, south, and west may be different.

// We are allowed to increase the height of any number of buildings by any amount (the amount can be different per building). The height of a 0-height building can also be increased. However, increasing the height of a building should not affect the city's skyline from any cardinal direction.

// Return the maximum total sum that the height of the buildings can be increased by without changing the city's skyline from any cardinal direction.

// Example 1:

// Input: grid = [[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]
// Output: 35
// Explanation: The building heights are shown in the center of the above image.
// The skylines when viewed from each cardinal direction are drawn in red.
// The grid after increasing the height of buildings without affecting skylines is:
// gridNew = [ [8, 4, 8, 7],
//             [7, 4, 7, 7],
//             [9, 4, 8, 7],
//             [3, 3, 3, 3] ]

// Находит максиммум среди элементов строки матрциы
func FindMaxByColumn(matrix [][]int, j int) int {
	if len(matrix) == 0 {
		return -1
	}

	max := matrix[0][j]

	for _, row := range matrix {
		if row[j] > max {
			max = row[j]
		}
	}

	return max
}

// Находит максиммум среди элементов строки матрциы
func FindMaxByRow(matrix [][]int, i int) int {
	if len(matrix) == 0 {
		return -1
	}
	row := matrix[i]

	max := row[0]

	for _, cell := range row {
		if cell > max {
			max = cell
		}
	}

	return max
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	total := 0
	columnMaxHorizonBoundary := make([]int, n)

	for i := 0; i < n; i++ {
		rowMaxHorizonBoundary := FindMaxByRow(grid, i)
		for j := 0; j < n; j++ {
			if i == 0 {
				columnMaxHorizonBoundary[j] = FindMaxByColumn(grid, j)
			}

			maxHorizonBoundary := int(math.Min(float64(rowMaxHorizonBoundary), float64(columnMaxHorizonBoundary[j])))
			if maxHorizonBoundary > grid[i][j] {
				total += maxHorizonBoundary - grid[i][j]
			}
		}
	}

	return total
}
