package SubtractTheProductAndSumOfDigitsOfAnInteger

// Given an integer number n, return the difference between the product of its digits and the sum of its digits.

// Example 1:

// Input: n = 234
// Output: 15
// Explanation:
// Product of digits = 2 * 3 * 4 = 24
// Sum of digits = 2 + 3 + 4 = 9
// Result = 24 - 9 = 15

func parseDigits(num int) []int {
	digits := []int{}
	for num != 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	return digits
}

func Calc(nums []int, calc func(a, b int) int) int {
	res := 0
	for num := range nums {
		res = calc(res, num)
	}

	return res
}

func sum(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}

	return res
}

func multiple(nums []int) int {
	res := 1
	for _, num := range nums {
		res *= num
	}

	return res
}

func subtractProductAndSum(n int) int {
	digits := parseDigits(n)
	productRes := multiple(digits)
	sumRes := sum(digits)

	return productRes - sumRes
}
