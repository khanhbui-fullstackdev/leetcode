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

	res := hasCycle(head)
	if res {
		fmt.Printf("There is a cycle in the linked list")
	} else {
		fmt.Printf("There is no cycle in the linked list.")
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *models.ListNode) bool {
	return true
}
