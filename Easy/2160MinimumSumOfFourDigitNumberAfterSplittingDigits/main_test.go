package MinimumSumOfFourDigitNumberAfterSplittingDigits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSum(t *testing.T) {
	assert.Equal(t,
		52,
		minimumSum(2932))
}
