package main

func getConcatenation(nums []int) []int {
	n := len(nums)
	newLenght := n * 2
	ans := make([]int, newLenght)

	for i := 0; i < newLenght; i++ {
		ans[i] = nums[i%n]
	}

	return ans
}

func main() {

}
