package MergeNodesInBetweenZeros

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeNodes(t *testing.T) {
	assert.EqualValues(t,
		&ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 11,
			},
		},
		mergeNodes(&ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 2,
									Next: &ListNode{
										Val: 0,
									},
								},
							},
						},
					},
				},
			},
		}))
}

//[0,3,1,0,4,5,2,0]
