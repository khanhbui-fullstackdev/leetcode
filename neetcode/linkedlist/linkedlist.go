package main

import (
	"fmt"
	"leetcode/neetcode/linkedlist/models"
)

func RunHasCycle() {
	// 1. Tạo độc lập các Node
	head := &models.ListNode{Val: 3}
	node2 := &models.ListNode{Val: 2}
	node3 := &models.ListNode{Val: 0}
	node4 := &models.ListNode{Val: -4}

	// 2. Nối chuỗi theo thứ tự
	head.Next = node2
	node2.Next = node3
	node3.Next = node4

	// 3. Tạo "quay xe" (cycle): Cho node cuối trỏ về node2
	node4.Next = node2

	isCycle := hasCycleUsingHashTable(head)
	if isCycle {
		fmt.Printf("There is a cycle in the linked list")
	} else {
		fmt.Printf("There is no cycle in the linked list.")
	}
	fmt.Println()

	// 1 create node
	head = &models.ListNode{Val: 1}
	node2 = &models.ListNode{Val: 2}

	// 2 create a link between nodes
	head.Next = node2
	node2.Next = head
	isCycle = hasCycleUsingHashTable(head)
	if isCycle {
		fmt.Printf("There is a cycle in the linked list")
	} else {
		fmt.Printf("There is no cycle in the linked list.")
	}
	fmt.Println()

	// 1 create node
	head = &models.ListNode{Val: 1}
	isCycle = hasCycleUsingHashTable(head)
	if isCycle {
		fmt.Printf("There is a cycle in the linked list")
	} else {
		fmt.Printf("There is no cycle in the linked list.")
	}
	fmt.Println()
}

/*
*
Time complexity O(n)
Space complexity O(n)
*/
func hasCycleUsingHashTable(head *models.ListNode) bool {
	nodeMap := make(map[*models.ListNode]bool, 10)

	currentNode := head
	for currentNode != nil {
		_, existed := nodeMap[currentNode]
		if existed {
			return true
		} else {
			nodeMap[currentNode] = true
		}
		currentNode = currentNode.Next
	}

	return false
}

func hasCycleUsingTwoPointers(head *models.ListNode) bool {
	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}
