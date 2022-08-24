package MergeNodesInBetweenZeros

// You are given the head of a linked list, which contains a series of integers separated by 0's. The beginning and end of the linked list will have Node.val == 0.

// For every two consecutive 0's, merge all the nodes lying in between them into a single node whose value is the sum of all the merged nodes. The modified list should not contain any 0's.

// Return the head of the modified linked list.

// Example 1:

// Input: head = [0,3,1,0,4,5,2,0]
// Output: [4,11]
// Explanation:
// The above figure represents the given linked list. The modified list contains
// - The sum of the nodes marked in green: 3 + 1 = 4.
// - The sum of the nodes marked in red: 4 + 5 + 2 = 11.

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	headSumList := &ListNode{}

	for lastSum := headSumList; head != nil && head.Next != nil; {
		if head.Val == 0 {
			sum := 0
			head = head.Next
			for head != nil && head.Val != 0 {
				sum += head.Val
				head = head.Next
			}

			if lastSum.Val == 0 {
				lastSum.Val = sum
			} else {
				lastSum.Next = &ListNode{
					Val: sum,
				}
				lastSum = lastSum.Next
			}

			if head == nil {
				return headSumList
			}
		}
	}

	return headSumList
}
