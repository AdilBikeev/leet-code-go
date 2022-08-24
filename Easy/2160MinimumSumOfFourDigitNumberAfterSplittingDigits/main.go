package MinimumSumOfFourDigitNumberAfterSplittingDigits

import (
	"sort"
	"strconv"
	"strings"
)

// You are given a positive integer num consisting of exactly four digits. Split num into two new integers new1 and new2 by using the digits found in num. Leading zeros are allowed in new1 and new2, and all the digits found in num must be used.

// For example, given num = 2932, you have the following digits: two 2's, one 9 and one 3. Some of the possible pairs [new1, new2] are [22, 93], [23, 92], [223, 9] and [2, 329].
// Return the minimum possible sum of new1 and new2.

// Example 1:

// Input: num = 2932
// Output: 52
// Explanation: Some possible pairs [new1, new2] are [29, 23], [223, 9], etc.
// The minimum sum can be obtained by the pair [29, 23]: 29 + 23 = 52.

func sliceAtoi(numsStr []string) []int {
	n := len(numsStr)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i], _ = strconv.Atoi(numsStr[i])
	}

	return res
}

func minimumSum(num int) int {
	sortedDigitsStr := strings.Split(strconv.Itoa(num), "")
	sort.Strings(sortedDigitsStr) // сортируем по возврастанию

	sortedDigitsInt := sliceAtoi(sortedDigitsStr)

	pair1 := sortedDigitsInt[0]*10 + sortedDigitsInt[2]
	pair2 := sortedDigitsInt[1]*10 + sortedDigitsInt[3]

	return pair1 + pair2
}
