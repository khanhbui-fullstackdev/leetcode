package main

import (
	"fmt"
	"leetcode/engineeringpro/linkedlist/models"
)

func RunMiddleNode() {
	node1 := &models.ListNode{
		Val: 2,
		Next: &models.ListNode{
			Val: 3,
			Next: &models.ListNode{
				Val:  4,
				Next: &models.ListNode{Val: 5}}}}

	head1 := &models.ListNode{
		Val:  1,
		Next: node1,
	}

	fmt.Println("Middle node:", middleNode(head1))

	testcase3_node1 := &models.ListNode{Val: 1, Next: &models.ListNode{Val: 2}}
	fmt.Println("Middle node:", middleNode(testcase3_node1))

	testcase4_node1 := &models.ListNode{Val: 1, Next: &models.ListNode{Val: 2, Next: &models.ListNode{Val: 3}}}
	fmt.Println("Middle node:", middleNode(testcase4_node1))
}

func middleNode(head *models.ListNode) *models.ListNode {
	if head.Next == nil {
		return head
	}

	currentNode := head
	fastNode := head

	for currentNode != nil {
		if fastNode == nil || fastNode.Next == nil {
			return currentNode
		}
		currentNode = currentNode.Next
		fastNode = fastNode.Next.Next
	}
	return nil
}

func RunDeleteDuplicates() {
	testcase_1 := &models.ListNode{
		Val:  1,
		Next: &models.ListNode{Val: 1, Next: &models.ListNode{Val: 2}},
	}
	testcase_1 = deleteDuplicates(testcase_1)
	testcase_1.PrintAllListNode()

	testcase_2 := &models.ListNode{
		Val: 1,
		Next: &models.ListNode{
			Val: 1,
			Next: &models.ListNode{
				Val: 2,
				Next: &models.ListNode{
					Val: 3,
					Next: &models.ListNode{
						Val: 3}}}},
	}
	testcase_2 = deleteDuplicates(testcase_2)
	testcase_2.PrintAllListNode()
}

func deleteDuplicates(head *models.ListNode) *models.ListNode {
	currentNode := head

	for currentNode != nil {
		nextNode := currentNode.Next
		if nextNode != nil && currentNode.Val == nextNode.Val {
			currentNode.Next = nextNode.Next
			nextNode = nil
		} else {
			currentNode = currentNode.Next
		}
	}

	return head
}

func RunRemoveElement() {
	testcase_head := &models.ListNode{
		Val: 1,
		Next: &models.ListNode{
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
								Val: 6}}},
				},
			},
		},
	}

	remainingNodes1 := removeElements(testcase_head, 6)
	remainingNodes1.PrintAllListNode()

	testcase_head1 := &models.ListNode{
		Val: 7,
		Next: &models.ListNode{
			Val: 7,
			Next: &models.ListNode{
				Val:  7,
				Next: &models.ListNode{Val: 7}}},
	}

	remainingNodes2 := removeElements(testcase_head1, 7)
	remainingNodes2.PrintAllListNode()
}

func removeElements(head *models.ListNode, val int) *models.ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil && head.Val == val {
		return head
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

func RunGetIntersectionNode() {
	testcase1_sharedNode := &models.ListNode{Val: 8, Next: &models.ListNode{Val: 4, Next: &models.ListNode{Val: 5}}}
	testcase1_nodeA := &models.ListNode{Val: 4, Next: &models.ListNode{Val: 1, Next: testcase1_sharedNode}}
	testcase1_nodeB := &models.ListNode{Val: 5, Next: &models.ListNode{Val: 6, Next: &models.ListNode{Val: 1, Next: testcase1_sharedNode}}}

	testcase1_intersectionNode := getIntersectionNodeUsingMath(testcase1_nodeA, testcase1_nodeB)
	testcase1_intersectionNode.PrintAllListNode()
	fmt.Printf("\nPrint memory address of intersection node:%p", testcase1_intersectionNode)
	fmt.Println()

	// testcase2_sharedNode := &models.ListNode{Val: 2, Next: &models.ListNode{Val: 4}}
	// testcase2_nodeA := &models.ListNode{Val: 1, Next: &models.ListNode{Val: 9, Next: &models.ListNode{Val: 1, Next: testcase2_sharedNode}}}
	// testcase2_nodeB := &models.ListNode{Val: 3, Next: testcase2_sharedNode}

	// testcase2_intersectionNode := getIntersectionNodeUsingMap(testcase2_nodeA, testcase2_nodeB)
	// testcase2_intersectionNode.PrintAllListNode()
	// fmt.Printf("\nPrint memory address of intersection node:%p", testcase2_intersectionNode)
	// fmt.Println()

	// testcase40_sharedNode := &models.ListNode{Val: 1}
	// testcase40_nodeA := &models.ListNode{Val: 1, Next: testcase40_sharedNode}
	// testcase40_nodeB := &models.ListNode{Val: 1, Next: testcase40_sharedNode}
	// testcase40_intersectionNode := getIntersectionNodeUsingMap(testcase40_nodeA, testcase40_nodeB)
	// testcase40_intersectionNode.PrintAllListNode()
	// fmt.Printf("\nPrint memory address of intersection node:%p", testcase40_intersectionNode)
	// fmt.Println()
}

func getIntersectionNodeUsingMap(headA, headB *models.ListNode) *models.ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	memoryAddressMap := make(map[string]bool)
	currentNodeA := headA

	for currentNodeA != nil {
		memoryAddress := fmt.Sprintf("%p", currentNodeA)
		existed := memoryAddressMap[memoryAddress]
		if !existed {
			memoryAddressMap[memoryAddress] = true
		}
		currentNodeA = currentNodeA.Next
	}
	currentNodeB := headB
	for currentNodeB != nil {
		memoryAddress := fmt.Sprintf("%p", currentNodeB)
		existed := memoryAddressMap[memoryAddress]
		if !existed {
			memoryAddressMap[memoryAddress] = true
		} else {
			// release memory
			clear(memoryAddressMap)
			return currentNodeB
		}
		currentNodeB = currentNodeB.Next
	}
	clear(memoryAddressMap)
	return nil
}

func getIntersectionNodeUsingMath(headA, headB *models.ListNode) *models.ListNode {
	currentNodeA := headA
	currentNodeB := headB

	for currentNodeA != currentNodeB {
		if currentNodeA != nil {
			currentNodeA = currentNodeA.Next
		} else {
			currentNodeA = headB
		}

		if currentNodeB != nil {
			currentNodeB = currentNodeB.Next
		} else {
			currentNodeB = headA
		}
	}

	headA.PrintAllListNode()
	fmt.Println()

	headB.PrintAllListNode()
	fmt.Println()

	return currentNodeA
}
