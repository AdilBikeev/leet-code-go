package SubrectangleQueries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateSubrectangle(t *testing.T) {
	obj := Constructor([][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	})
	obj.UpdateSubrectangle(0, 0, 3, 2, 5)
	param_2 := obj.GetValue(0, 2)

	assert.Equal(t,
		5,
		param_2)
}
