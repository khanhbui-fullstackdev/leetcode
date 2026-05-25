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
	middleNode(head1).PrintAllListNode()

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

func RunReverseList() {
	testcase1_head := &models.ListNode{
		Val: 1, Next: &models.ListNode{
			Val: 2,
			Next: &models.ListNode{
				Val: 3,
				Next: &models.ListNode{
					Val:  4,
					Next: &models.ListNode{Val: 5}}}},
	}
	testcase1_reveredNode := reverseList(testcase1_head)
	testcase1_reveredNode.PrintAllListNode()
}

func reverseList(head *models.ListNode) *models.ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var prevNode *models.ListNode = nil
	currentNode := head
	nextNode := currentNode

	for currentNode != nil {
		nextNode = currentNode.Next
		currentNode.Next = prevNode

		prevNode = currentNode
		currentNode = nextNode
	}

	return prevNode
}

func RunMergeTwoLists() {
	testcase1_list1 := &models.ListNode{Val: 1, Next: &models.ListNode{Val: 2, Next: &models.ListNode{Val: 4}}}
	testcase1_list2 := &models.ListNode{Val: 1, Next: &models.ListNode{Val: 3, Next: &models.ListNode{Val: 4}}}
	testcase1_mergedList := mergeTwoLists(testcase1_list1, testcase1_list2)
	testcase1_mergedList.PrintAllListNode()

	testcase3_list2 := &models.ListNode{Val: 0}
	testcase3_mergedList := mergeTwoLists(nil, testcase3_list2)
	testcase3_mergedList.PrintAllListNode()
}

func mergeTwoLists(list1 *models.ListNode, list2 *models.ListNode) *models.ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	// Khởi tạo node output
	mergedListNode := &models.ListNode{Val: -1, Next: nil}
	workerNode := mergedListNode

	for list1 != nil && list2 != nil {
		list1Val := list1.Val
		list2Val := list2.Val

		if list1Val <= list2Val {
			// Add all of list 1 into workerNode
			node := &models.ListNode{Val: list1Val, Next: nil}
			workerNode.Next = node
			workerNode = workerNode.Next
			list1 = list1.Next
		} else {
			// Add all of list 2 into workerNode
			node := &models.ListNode{Val: list2Val, Next: nil}
			workerNode.Next = node
			workerNode = workerNode.Next
			list2 = list2.Next
		}
	}

	for list1 != nil {
		// Add all of list 1 into workerNode
		node := &models.ListNode{Val: list1.Val, Next: nil}
		workerNode.Next = node
		workerNode = workerNode.Next
		list1 = list1.Next
	}

	for list2 != nil {
		// Add all of list 1 into workerNode
		node := &models.ListNode{Val: list2.Val, Next: nil}
		workerNode.Next = node
		workerNode = workerNode.Next
		list2 = list2.Next
	}
	return mergedListNode.Next
}

func RunIsPalindrome() {
	testcase1_head := &models.ListNode{
		Val: 1,
		Next: &models.ListNode{
			Val: 2,
			Next: &models.ListNode{
				Val:  2,
				Next: &models.ListNode{Val: 1}}}}

	fmt.Printf("Is palindrome:%v", isPalindrome(testcase1_head))
	fmt.Println()

	testcase2_head := &models.ListNode{
		Val: 1,
		Next: &models.ListNode{
			Val:  2,
			Next: nil}}

	fmt.Printf("Is palindrome:%v", isPalindrome(testcase2_head))
	fmt.Println()

	testcase3_node := &models.ListNode{
		Val: 2,
		Next: &models.ListNode{
			Val: 3,
			Next: &models.ListNode{
				Val:  4,
				Next: &models.ListNode{Val: 5}}}}

	testcase3_head := &models.ListNode{
		Val:  1,
		Next: testcase3_node,
	}

	fmt.Printf("Is palindrome:%v", isPalindrome(testcase3_head))
	fmt.Println()
}

func isPalindrome(head *models.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	middleNode := findMiddleNode(head)
	fmt.Printf("Print middle node:%v", middleNode)
	middleNode.PrintAllListNode()

	reversedNode := reverseListNodes(middleNode)
	fmt.Printf("Print reverse node:%v", reversedNode)
	reversedNode.PrintAllListNode()

	currentNode := head
	for reversedNode != nil {
		if currentNode.Val == reversedNode.Val {
			currentNode = currentNode.Next
			reversedNode = reversedNode.Next
		} else {
			return false
		}
	}

	return true
}

func findMiddleNode(node *models.ListNode) *models.ListNode {
	currentNode := node
	fastNode := node
	for currentNode != nil {
		if fastNode == nil || fastNode.Next == nil {
			return currentNode
		}
		fastNode = currentNode.Next.Next
		currentNode = currentNode.Next
	}
	return nil
}

func reverseListNodes(node *models.ListNode) *models.ListNode {
	var previousNode *models.ListNode = nil
	currentNode := node
	var nextNode *models.ListNode = nil

	for currentNode != nil {
		nextNode = currentNode.Next
		currentNode.Next = previousNode

		previousNode = currentNode
		currentNode = nextNode
	}
	return previousNode
}

func BuildDoubleLinkedList() {
	linkedList := &models.LinkedList{}

	// Test AddFirst()
	// linkedList.AddFirst(10)
	// linkedList.AddFirst(20)
	// linkedList.AddFirst(30)
	// linkedList.PrintAllNodes()

	// Test AddLast()
	linkedList.AddLast(30)
	linkedList.AddLast(75)
	linkedList.AddLast(860)
	linkedList.AddLast(730)
	linkedList.AddLast(20)
	linkedList.PrintAllNodes()

	// Test DesignateNodeAsHead()
	headNode, err := linkedList.DesignateNodeAsHead("node#1")
	if err != nil {
		fmt.Printf("\n Error:%s", err.Error())
	} else {
		fmt.Printf("\n Head node:%v", headNode)
	}
	fmt.Println()
	linkedList.PrintAllNodes()

	// linkedList.RemoveNodeByVal(860)
	// linkedList.PrintAllNodes()
	// linkedList.PrintAllNodeMaps()

	// foundNodes, err := linkedList.FindNodesByValue(30)
	// if err != nil {
	// 	fmt.Printf("\n Error:%s", err.Error())
	// } else {
	// 	for _, node := range foundNodes {
	// 		fmt.Printf("\nNodeId:%s, val:%d", node.Id, node.Val)
	// 	}
	// }
	// fmt.Println()

	// foundNode, err := linkedList.FindNodeById("node#7")
	// if err != nil {
	// 	fmt.Printf("\n Error:%s", err.Error())
	// } else {
	// 	fmt.Printf("\n Found node:%v", foundNode)
	// }
	// fmt.Println()

	// newNode, err := linkedList.AddNewNodeAfterExistingNode("node#6", 235)
	// if err != nil {
	// 	fmt.Printf("\n Error:%s", err.Error())
	// } else {
	// 	fmt.Printf("\n New node:%v", newNode)
	// }
	// fmt.Println()
	// linkedList.PrintAllNodes()

}

func RunLRUCache() {
	// Test case 1
	// lruCache := models.Constructor(2)
	// lruCache.Put(1, 1)
	// lruCache.Put(2, 2)
	// lruCache.PrintAllListNode()

	// fmt.Println("Get(1):", lruCache.Get(1))
	// lruCache.PrintAllListNode()

	// lruCache.Put(3, 3)
	// lruCache.PrintAllListNode()
	// lruCache.PrintAllCacheMap()

	// fmt.Println("Get(2):", lruCache.Get(2))

	// lruCache.Put(4, 4)
	// lruCache.PrintAllListNode()
	// lruCache.PrintAllCacheMap()

	// fmt.Println("Get(1):", lruCache.Get(1))
	// fmt.Println("Get(3):", lruCache.Get(3))
	// fmt.Println("Get(4):", lruCache.Get(4))

	// Test case 2
	// lruCache := models.Constructor(1)
	// lruCache.Put(2, 1)
	// lruCache.PrintAllListNode()

	// fmt.Println("Get(2):", lruCache.Get(2))
	// lruCache.Put(3, 2)
	// lruCache.PrintAllListNode()

	// fmt.Println("Get(2):", lruCache.Get(2))
	// fmt.Println("Get(3):", lruCache.Get(3))

	// Test case 14
	lruCache := models.Constructor(3)
	lruCache.Put(1, 1)

	// headNode1 := lruCache.Head
	// fmt.Printf("\nHead1:%v | Next:%v | Prev:%v", headNode1, headNode1.Next, headNode1.Previous)

	lruCache.Put(2, 2)
	// headNode2 := lruCache.Head
	// fmt.Printf("\nHead2:%v | Next:%v | Prev:%v", headNode2, headNode2.Next, headNode2.Previous)

	lruCache.Put(3, 3)
	// headNode3 := lruCache.Head
	// fmt.Printf("\nHead3:%v | Next:%v | Prev:%v", headNode3, headNode3.Next, headNode3.Previous)
	// fmt.Printf("\nNode2:%v | Next:%v | Prev:%v", headNode3.Next, headNode3.Next.Next, headNode3.Next.Previous)
	// fmt.Printf("\nNode1:%v | Next:%v | Prev:%v", headNode3.Next.Next, headNode3.Next.Next.Next, headNode3.Next.Next.Previous)

	lruCache.Put(4, 4)
	// headNode4 := lruCache.Head
	// fmt.Printf("\nHead4:%v | Next:%v | Prev:%v", headNode4, headNode4.Next, headNode4.Previous)
	// fmt.Printf("\nNode3:%v | Next:%v | Prev:%v", headNode4.Next, headNode4.Next.Next, headNode4.Next.Previous)
	// fmt.Printf("\nNode2:%v | Next:%v | Prev:%v", headNode4.Next.Next, headNode4.Next.Next.Next, headNode4.Next.Next.Previous)

	fmt.Println("Get(4):", lruCache.Get(4))
	headNode4 := lruCache.Head
	fmt.Printf("\nHead4:%v | Next:%v | Prev:%v", headNode4, headNode4.Next, headNode4.Previous)
	fmt.Printf("\nNode3:%v | Next:%v | Prev:%v", headNode4.Next, headNode4.Next.Next, headNode4.Next.Previous)
	fmt.Printf("\nNode2:%v | Next:%v | Prev:%v", headNode4.Next.Next, headNode4.Next.Next.Next, headNode4.Next.Next.Previous)

	// lruCache.PrintAllListNode()

	fmt.Println("Get(3):", lruCache.Get(3))
	headNode3 := lruCache.Head
	fmt.Printf("\nHead3:%v | Next:%v | Prev:%v", headNode3, headNode3.Next, headNode3.Previous)
	fmt.Printf("\nNode4:%v | Next:%v | Prev:%v", headNode3.Next, headNode3.Next.Next, headNode3.Next.Previous)
	fmt.Printf("\nNode2:%v | Next:%v | Prev:%v", headNode3.Next.Next, headNode3.Next.Next.Next, headNode3.Next.Next.Previous)

	fmt.Println("Get(2):", lruCache.Get(2))
	headNode2 := lruCache.Head
	fmt.Printf("\nHead2:%v | Next:%v | Prev:%v", headNode2, headNode2.Next, headNode2.Previous)
	fmt.Printf("\nNode3:%v | Next:%v | Prev:%v", headNode2.Next, headNode2.Next.Next, headNode2.Next.Previous)
	fmt.Printf("\nNode4:%v | Next:%v | Prev:%v", headNode2.Next.Next, headNode2.Next.Next.Next, headNode2.Next.Next.Previous)

	// lruCache.PrintAllListNode()

	// fmt.Println("Get(1):", lruCache.Get(1))
	// lruCache.PrintAllListNode()

	lruCache.Put(5, 5)
	headNode5 := lruCache.Head
	fmt.Printf("\nHead2:%v | Next:%v | Prev:%v", headNode5, headNode5.Next, headNode5.Previous)
	fmt.Printf("\nNode3:%v | Next:%v | Prev:%v", headNode5.Next, headNode5.Next.Next, headNode5.Next.Previous)
	fmt.Printf("\nNode4:%v | Next:%v | Prev:%v", headNode5.Next.Next, headNode5.Next.Next.Next, headNode5.Next.Next.Previous)
}
