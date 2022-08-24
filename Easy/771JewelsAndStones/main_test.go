package JewelsAndStones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumJewelsInStones(t *testing.T) {
	assert.Equal(t,
		3,
		numJewelsInStones("aA", "aAAbbbb"))
}
