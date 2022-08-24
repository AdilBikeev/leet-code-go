package RichestCustomerWealth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumWealth(t *testing.T) {
	exp := 6
	assert.Equal(t, exp, maximumWealth([][]int{
		[]int{1, 2, 3},
		[]int{3, 2, 1},
	}))
}
