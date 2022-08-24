package DeepestLeavesSum

// Given the root of a binary tree, return the sum of values of its deepest leaves.

// Example 1:

// Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
// Output: 15
// Example 2:

// Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
// Output: 19

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

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

func (st *TreeNode) GetHeight() int {
	if st == nil {
		return 0
	}

	height := 0
	stack := Stack{
		values: []TreeNode{*st},
	}

	// Обходим каждый уровень дерева
	for ; stack.Size() != 0; height++ {
		newHeightStack := Stack{
			values: []TreeNode{},
		}
		// Формируем массив из вершин нового уровня
		for stack.Size() != 0 {
			v := stack.Pop()
			if v.Left != nil {
				newHeightStack.Push(*v.Left)
			}

			if v.Right != nil {
				newHeightStack.Push(*v.Right)
			}
		}

		// Обновляем текущий уровень
		stack = *NewStack(newHeightStack.values)
	}

	return height
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil { // is empty Tree
		return 0
	} else if root.Left == nil && root.Right == nil { // Tree consists of 1 node
		return root.Val
	}

	stack := Stack{
		values: []TreeNode{*root},
	}

	// Обходим каждый уровень дерева
	for i := 0; stack.Size() != 0; i++ {
		newHeightStack := Stack{
			values: []TreeNode{},
		}
		sum := 0
		// Формируем массив из вершин нового уровня
		for stack.Size() != 0 {
			v := stack.Pop()
			sum += v.Val

			if v.Left != nil {
				newHeightStack.Push(*v.Left)
			}

			if v.Right != nil {
				newHeightStack.Push(*v.Right)
			}
		}

		if newHeightStack.Size() == 0 {
			return sum
		}

		// Обновляем текущий уровень
		stack = *NewStack(newHeightStack.values)
	}

	// Потенциально ошибочный кейс поэтому возвращаем -1
	return -1
}
