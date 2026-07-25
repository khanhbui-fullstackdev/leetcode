package main

import (
	"fmt"
	"leetcode/neetcode/linkedlist/models"
)

func RunMergeTwoLists() {
	list1 := &models.ListNode{Val: 1}
	list1_node2 := &models.ListNode{Val: 2}
	list1_node4 := &models.ListNode{Val: 4}
	list1.Next = list1_node2
	list1_node2.Next = list1_node4

	list2 := &models.ListNode{Val: 1}
	list2_node3 := &models.ListNode{Val: 3}
	list2_node4 := &models.ListNode{Val: 4}
	list2.Next = list2_node3
	list2_node3.Next = list2_node4

	mergedLinkedList := mergeTwoLists(list1, list2)
	mergedLinkedList.PrintAllListNodes()
}

/*
*

	list1 = [1,2,4]
	list2 = [1,3,4]

	mergedLinkedList = [-101] // dummy node = -101

	1.tranverse list1 && list2
	2.compare value between list1.Val & list2.Val
	    if list1.Val <= list2.Val
	    {
	        mergedLinkedList.Next = list1
	        list1 = list1.Next
	    }else{
	        mergedLinkedList.Next = list2
	        list2 = list2.Next
	    }

	3.add list into output mergedLinkedList

	list1=node1
	list2=node1
	-> compare
	        mergedLinkedList.Next = node1
	        <=> node-101.Next = node1 //  [-101 -> 1]

	        list1=list1.Next = node1.Next = node2
	        mergedLinkedList=mergedLinkedList.Next // node1

	list1=node2
	list2=node1
	-> compare
	        mergedLinkedList.Next = node1 //  [-101->1->1]
	        <=> node1.Next = node1

	        list2=list2.Next = node3
	        mergedLinkedList=mergedLinkedList.Next // node1

	list1=node2
	list2=node3
	-> compare
	        mergedLinkedList.Next = node2 //  [-101->1->1->2]
	        <=> node1.next = node2

	        list1 = list1.Next = node4
	        mergedLinkedList=mergedLinkedList.Next // node2

	list1=node4
	list2=node3
	->compare
	        mergedLinkedList.Next = node3 //  [-101->1->1->2->3]
	        <=> node2.next = node3
	        list2 = list2.Next = node4
	        mergedLinkedList=mergedLinkedList.Next // node3

	list1=node4
	list2=node4
	->compare
	        mergedLinkedList.Next = node4 //  [-101->1->1->2->3->4]
	        <=> node3.next = node4
	        list1=list1.Next = nil
	        mergedLinkedList=mergedLinkedList.Next // node4

	for list1!=nil{
	    mergedLinkedList.Next = list1
	    list1=list1.next
	}

	for list2!=nil {
	    mergedLinkedList.Next = list2 // node4 //  [-101->1->1->2->3->4->4]
	    list2=list2.next //
	}

	return mergedLinkedList.next
*/
/**
  mergedLists:={val:-101} // 0x000
  worker:=mergedLists// 0x000
  vì mergedLists & worker đều có chung 1 vùng nhớ là 0x00

*/
func mergeTwoLists(list1 *models.ListNode, list2 *models.ListNode) *models.ListNode {
	dummyNode := &models.ListNode{Val: -1, Next: nil}
	mergedList, workerNode := dummyNode, dummyNode

	// traverse list1 & list2
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			// add list1 into workernode
			workerNode.Next = list1
			list1 = list1.Next
		} else {
			// add list2 into workernode
			workerNode.Next = list2
			list2 = list2.Next
		}
		workerNode = workerNode.Next
	}

	if list1 != nil {
		workerNode.Next = list1
	}

	if list2 != nil {
		workerNode.Next = list2
	}

	workerNode.PrintAllListNodes()

	return mergedList.Next
}

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

func RunDeleteDuplicates() {
	head := &models.ListNode{Val: 1}
	node1 := &models.ListNode{Val: 1}
	node2 := &models.ListNode{Val: 2}

	head.Next = node1
	node1.Next = node2

	nodes := deleteDuplicates(head)
	nodes.PrintAllListNodes()

	head = &models.ListNode{Val: 1}
	node1 = &models.ListNode{Val: 1}
	node2 = &models.ListNode{Val: 2}
	node3 := &models.ListNode{Val: 3}
	node3_1 := &models.ListNode{Val: 3}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node3_1

	nodes = deleteDuplicates(head)
	nodes.PrintAllListNodes()
}

func deleteDuplicates(head *models.ListNode) *models.ListNode {
	currentNode := head
	for currentNode != nil {
		nextNode := currentNode.Next
		if nextNode != nil && nextNode.Val == currentNode.Val {
			currentNode.Next = nextNode.Next
			nextNode = nil
		} else {
			currentNode = currentNode.Next
		}
	}
	return head
}

func RunRemoveElements() {
	head := &models.ListNode{Val: 1, Next: &models.ListNode{
		Val: 2,
		Next: &models.ListNode{
			Val: 6,
			Next: &models.ListNode{
				Val: 3,
				Next: &models.ListNode{
					Val: 4,
					Next: &models.ListNode{
						Val: 5,
						Next: &models.ListNode{
							Val:  6,
							Next: nil}}}}},
	}}
	nodes := removeElements(head, 6)
	nodes.PrintAllListNodes()

	nodes = removeElements(nil, 6)
	nodes.PrintAllListNodes()

	head = &models.ListNode{Val: 7, Next: &models.ListNode{
		Val: 7,
		Next: &models.ListNode{
			Val: 7,
			Next: &models.ListNode{
				Val:  7,
				Next: nil},
		}}}

	nodes = removeElements(head, 7)
	nodes.PrintAllListNodes()
}

func removeElements(head *models.ListNode, val int) *models.ListNode {
	if head == nil {
		return nil
	}
	dummyNode := &models.ListNode{Val: -1, Next: head}
	currentNode := dummyNode
	for currentNode != nil {
		nextNode := currentNode.Next
		if nextNode != nil && nextNode.Val == val {
			currentNode.Next = nextNode.Next
			nextNode = nil
		} else {
			currentNode = currentNode.Next
		}
	}
	return dummyNode.Next
}
