package MaxIncreaseToKeepCitySkyline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxByColumn(t *testing.T) {
	grid := [][]int{
		{8, 4, 8, 7},
		{7, 4, 7, 7},
		{9, 4, 8, 7},
		{3, 3, 3, 3},
	}

	assert.Equal(t,
		9,
		FindMaxByColumn(grid, 0))

	assert.Equal(t,
		4,
		FindMaxByColumn(grid, 1))

	assert.Equal(t,
		8,
		FindMaxByColumn(grid, 2))

	assert.Equal(t,
		7,
		FindMaxByColumn(grid, 3))
}

func TestFindMaxByRow(t *testing.T) {
	grid := [][]int{
		{8, 4, 8, 7},
		{7, 4, 7, 7},
		{9, 4, 8, 7},
		{3, 3, 3, 3},
	}

	assert.Equal(t,
		8,
		FindMaxByRow(grid, 0))

	assert.Equal(t,
		7,
		FindMaxByRow(grid, 1))

	assert.Equal(t,
		9,
		FindMaxByRow(grid, 2))

	assert.Equal(t,
		3,
		FindMaxByRow(grid, 3))
}

func TestMaxIncreaseKeepingSkyline(t *testing.T) {
	grid := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}

	assert.Equal(t,
		35,
		maxIncreaseKeepingSkyline(grid))

}
