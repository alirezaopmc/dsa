package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	begin := head
	var finalHead *ListNode = nil
	for ok, end := kNextNode(begin, k); ok; {
		if finalHead == nil {
			finalHead = end
		}
		next := end.Next
		reverseGroup(begin, end, next)
		begin = next
	}
	return finalHead
}

func reverseGroup(begin, end, after *ListNode) {
	trav := begin
	next := begin.Next
	begin.Next = after
	for trav != end {
		temp := next.Next
		next.Next = trav
		trav = next
		next = temp
	}
}

func kNextNode(head *ListNode, k int) (bool, *ListNode) {
	if k == 0 && head != nil {
		return true, head
	}
	if head == nil {
		return false, nil
	}
	return kNextNode(head.Next, k-1)
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
	}
	return head
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: [1,2,3,4,5], k=2
	list1 := createList([]int{1, 2, 3, 4, 5})
	fmt.Println("Test 1: k=2")
	fmt.Print("Input:  ")
	printList(list1)
	result1 := reverseKGroup(list1, 2)
	fmt.Print("Output: ")
	printList(result1)
	fmt.Println()

	// Test case 2: [1,2,3,4,5], k=3
	list2 := createList([]int{1, 2, 3, 4, 5})
	fmt.Println("Test 2: k=3")
	fmt.Print("Input:  ")
	printList(list2)
	result2 := reverseKGroup(list2, 3)
	fmt.Print("Output: ")
	printList(result2)
	fmt.Println()

	// Test case 3: [1], k=1
	list3 := createList([]int{1})
	fmt.Println("Test 3: k=1")
	fmt.Print("Input:  ")
	printList(list3)
	result3 := reverseKGroup(list3, 1)
	fmt.Print("Output: ")
	printList(result3)
}
