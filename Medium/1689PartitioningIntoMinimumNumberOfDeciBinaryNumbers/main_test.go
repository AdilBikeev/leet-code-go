package PartitioningIntoMinimumNumberOfDeciBinaryNumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeater(t *testing.T) {
	assert.Equal(t,
		"1",
		repeater("1", 1))

	assert.Equal(t,
		"11",
		repeater("1", 2))

	assert.Equal(t,
		"111",
		repeater("1", 3))
}

func TestFreqInNum(t *testing.T) {
	assert.Equal(t,
		2,
		freqInNum(32, 11))

	assert.Equal(t,
		3,
		freqInNum(33, 11))
}

func TestValidGetMaxBinaryBelowNum(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{32, 11},
		{33, 11},
		{10, 10},
		{11, 11},
		{12, 11},
	}

	for _, test := range tests {
		binaryMax, err := getMaxBinaryBelowNum(test.input)
		if err != nil {
			t.Error("Handling error: ", err)
		}
		assert.Equal(test.expected, binaryMax)
	}
}

func TestMinPartitions(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    string
		expected int
	}{
		{"32", 3},
		{"82734", 8},
	}

	for _, test := range tests {
		res := minPartitions(test.input)
		assert.Equal(test.expected, res)
	}
}
