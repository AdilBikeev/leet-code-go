package PartitioningIntoMinimumNumberOfDeciBinaryNumbers

import (
	"math"
)

// A decimal number is called deci-binary if each of its digits is either 0 or 1 without any leading zeros. For example, 101 and 1100 are deci-binary, while 112 and 3001 are not.

// Given a string n that represents a positive decimal integer, return the minimum number of positive deci-binary numbers needed so that they sum up to n.

// Example 1:

// Input: n = "32"
// Output: 3
// Explanation: 10 + 11 + 11 = 32

// Example 2:

// Input: n = "16"
// Output: 3
// Explanation: 11 + 1 + 1 + 1 + 1 + 1 = 32

func minPartitions(n string) int {
	nLen := len(n)
	digitMax := int(n[0] - '0')

	for i := 1; i < nLen; i++ {
		digitMax = int(math.Max(float64(digitMax), float64(n[i]-'0')))
	}

	return digitMax
}

// Try to express the string(number) in sum of 1's.
// for eg :- 32
// a) 3:-> 1 +1+1
// b) 2 :-> 1+1
// We observe it has three 1's in 3 and two 1's in 2.
// Now 32 can be written as:-
// 32:- 10 + 11 + 11
