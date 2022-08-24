package CountNodesEqualToAverageOfSubtree

// Given the root of a binary tree, return the number of nodes where the value of the node is equal to the average of the values in its subtree.

// Note:

// The average of n elements is the sum of the n elements divided by n and rounded down to the nearest integer.
// A subtree of root is a tree consisting of root and all of its descendants.

// Example 1:

// Input: root = [4,8,5,0,1,null,6]
// Output: 5
// Explanation:
// For the node with value 4: The average of its subtree is (4 + 8 + 5 + 0 + 1 + 6) / 6 = 24 / 6 = 4.
// For the node with value 5: The average of its subtree is (5 + 6) / 2 = 11 / 2 = 5.
// For the node with value 0: The average of its subtree is 0 / 1 = 0.
// For the node with value 1: The average of its subtree is 1 / 1 = 1.
// For the node with value 6: The average of its subtree is 6 / 1 = 6.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	values []TreeNode
}

func NewStack(values []TreeNode) *Stack {
	return &Stack{
		values: values,
	}
}

func (st *Stack) Size() int {
	return len(st.values)
}

func (st *Stack) Push(val TreeNode) {
	st.values = append(st.values, val)
}

func (st *Stack) Top() TreeNode {
	return st.values[len(st.values)-1]
}

func (st *Stack) Pop() TreeNode {
	n := len(st.values)
	top := st.values[n-1]
	st.values = st.values[:n-1]
	return top
}

func averageOfSubtree(root *TreeNode) int {
	if root == nil { // is empty Tree
		return 0
	} else if root.Left == nil && root.Right == nil { // Tree consists of 1 node
		return 1
	}

	count := 0
	stack := Stack{
		values: []TreeNode{*root},
	}

	// Обходим каждый уровень дерева
	for i := 0; stack.Size() != 0; i++ {
		newHeightStack := Stack{
			values: []TreeNode{},
		}

		// Формируем массив из вершин нового уровня
		for stack.Size() != 0 {
			v := stack.Pop()
			if avgSum(&v) == v.Val {
				count++
			}

			if v.Left != nil {
				newHeightStack.Push(*v.Left)
			}

			if v.Right != nil {
				newHeightStack.Push(*v.Right)
			}
		}

		if newHeightStack.Size() == 0 {
			return count
		}

		// Обновляем текущий уровень
		stack = *NewStack(newHeightStack.values)
	}

	return count
}

func avgSum(root *TreeNode) int {
	if root == nil { // is empty Tree
		return 0
	} else if root.Left == nil && root.Right == nil { // Tree consists of 1 node
		return root.Val
	}

	n := 0
	sum := 0
	stack := Stack{
		values: []TreeNode{*root},
	}

	// Обходим каждый уровень дерева
	for i := 0; stack.Size() != 0; i++ {
		newHeightStack := Stack{
			values: []TreeNode{},
		}

		// Формируем массив из вершин нового уровня
		for stack.Size() != 0 {
			v := stack.Pop()
			sum += v.Val
			n++

			if v.Left != nil {
				newHeightStack.Push(*v.Left)
			}

			if v.Right != nil {
				newHeightStack.Push(*v.Right)
			}
		}

		if newHeightStack.Size() == 0 {
			return sum / n
		}

		// Обновляем текущий уровень
		stack = *NewStack(newHeightStack.values)
	}

	return sum / n
}
