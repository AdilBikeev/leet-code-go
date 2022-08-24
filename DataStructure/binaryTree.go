package DataStructure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (this *TreeNode) GetHeight() int {
	if this == nil {
		return 0
	}

	height := 0
	stack := Stack[TreeNode]{
		values: []TreeNode{*this},
	}

	// Обходим каждый уровень дерева
	for ; stack.Size() != 0; height++ {
		newHeightStack := Stack[TreeNode]{
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
