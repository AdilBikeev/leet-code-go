package DeepestLeavesSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeepestLeavesSum(t *testing.T) {
	mock := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}

	assert.Equal(t,
		15,
		deepestLeavesSum(&mock))
}
