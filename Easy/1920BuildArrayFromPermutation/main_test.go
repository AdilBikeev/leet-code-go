package BuildArrayFromPermutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildArray(t *testing.T) {
	assert.Equal(t,
		[]int{0, 1, 2, 4, 5, 3},
		buildArray([]int{0, 2, 1, 5, 3, 4}),
		"they shoudl be equal")
}
