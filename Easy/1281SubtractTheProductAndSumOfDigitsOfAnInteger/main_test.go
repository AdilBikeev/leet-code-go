package SubtractTheProductAndSumOfDigitsOfAnInteger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtractProductAndSum(t *testing.T) {
	assert.Equal(t,
		15,
		subtractProductAndSum(234))
}
