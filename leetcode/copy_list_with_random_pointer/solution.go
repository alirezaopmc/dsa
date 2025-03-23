package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	trav := head
	next := head.Next
	for trav != nil {
		next = trav.Next
		trav.Next = &Node{trav.Val, next, nil} // Cloning
		trav = next
	}

	trav = head
	for trav != nil {
		clone := trav.Next
		next = clone.Next
		if trav.Random != nil {
			clone.Random = trav.Random.Next
		}
		trav = next
	}

	origTrav := head
	copyTrav := head.Next
	copyHead := head.Next
	for origTrav != nil {
		origNext := copyTrav.Next
		origTrav.Next = origNext
		origTrav = copyTrav.Next
		if origNext != nil {
			copyTrav.Next = origNext.Next
			copyTrav = origTrav.Next
		}
	}

	return copyHead
}

// Helper function to create a test linked list
func createTestList(values []int, randomIndices []int) *Node {
	if len(values) == 0 {
		return nil
	}

	// Create nodes
	nodes := make([]*Node, len(values))
	for i := range values {
		nodes[i] = &Node{Val: values[i]}
	}

	// Connect next pointers
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	// Connect random pointers
	for i, randIdx := range randomIndices {
		if randIdx >= 0 {
			nodes[i].Random = nodes[randIdx]
		}
	}

	return nodes[0]
}

// Helper function to verify if two lists are identical
func verifyLists(original, copied *Node) bool {
	if original == nil && copied == nil {
		return true
	}
	if original == nil || copied == nil {
		return false
	}

	origMap := make(map[*Node]int)
	copyMap := make(map[*Node]int)

	// Map original nodes to indices
	idx := 0
	for curr := original; curr != nil; curr = curr.Next {
		origMap[curr] = idx
		idx++
	}

	// Map copied nodes to indices
	idx = 0
	for curr := copied; curr != nil; curr = curr.Next {
		copyMap[curr] = idx
		idx++
	}

	// Verify structure
	orig := original
	copy := copied
	for orig != nil && copy != nil {
		if orig.Val != copy.Val {
			return false
		}

		// Check random pointers
		origRandomIdx := -1
		if orig.Random != nil {
			origRandomIdx = origMap[orig.Random]
		}

		copyRandomIdx := -1
		if copy.Random != nil {
			copyRandomIdx = copyMap[copy.Random]
		}

		if origRandomIdx != copyRandomIdx {
			return false
		}

		orig = orig.Next
		copy = copy.Next
	}

	return orig == nil && copy == nil
}

func main() {
	// Test case 1: [[7,null],[13,0],[11,4],[10,2],[1,0]]
	test1 := createTestList([]int{7, 13, 11, 10, 1}, []int{-1, 0, 4, 2, 0})
	result1 := copyRandomList(test1)
	fmt.Printf("Test 1 passed: %v\n", verifyLists(test1, result1))

	// Test case 2: [[1,1],[2,1]]
	test2 := createTestList([]int{1, 2}, []int{1, 1})
	result2 := copyRandomList(test2)
	fmt.Printf("Test 2 passed: %v\n", verifyLists(test2, result2))

	// Test case 3: [[3,null],[3,0],[3,null]]
	test3 := createTestList([]int{3, 3, 3}, []int{-1, 0, -1})
	result3 := copyRandomList(test3)
	fmt.Printf("Test 3 passed: %v\n", verifyLists(test3, result3))
}
