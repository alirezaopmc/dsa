package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	} else if Abs(head.Val) > 100000 {
		return true
	} else {
		head.Val += Sign(head.Val) * 100000
		return hasCycle(head.Next)
	}
}

func Abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func Sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

func main() {
	// Create test cases
	tests := []struct {
		name     string
		list     *ListNode
		expected bool
	}{
		{
			name: "Cycle at pos 1",
			list: func() *ListNode {
				nodes := make([]*ListNode, 4)
				for i := range nodes {
					nodes[i] = &ListNode{Val: i + 1}
				}
				for i := 0; i < len(nodes)-1; i++ {
					nodes[i].Next = nodes[i+1]
				}
				nodes[len(nodes)-1].Next = nodes[1] // Create cycle
				return nodes[0]
			}(),
			expected: true,
		},
		{
			name: "No cycle",
			list: func() *ListNode {
				nodes := make([]*ListNode, 3)
				for i := range nodes {
					nodes[i] = &ListNode{Val: i + 1}
				}
				for i := 0; i < len(nodes)-1; i++ {
					nodes[i].Next = nodes[i+1]
				}
				return nodes[0]
			}(),
			expected: false,
		},
	}

	for _, t := range tests {
		actual := hasCycle(t.list)
		if actual != t.expected {
			fmt.Printf("Error in test %s: expected %v, got %v\n", t.name, t.expected, actual)
		} else {
			fmt.Printf("Test %s passed\n", t.name)
		}
	}
}
