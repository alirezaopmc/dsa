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
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
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
