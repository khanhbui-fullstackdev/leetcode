package main

import (
	"fmt"
	"strconv"

	"leetcode/engineeringpro/graph/models"
)

/*
====================================
100. Same Tree
====================================

🔗 Link:
https://leetcode.com/problems/same-tree/description/

🧠 Idea:

- DFS

- Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

⏱ Complexity:
- Time: O(n)
- Space: O(h)

====================================
*/
func RunIsSameTree() {
	// Test case 1
	pLeftNode1 := &models.TreeNode{Val: 2}
	pRightNode1 := &models.TreeNode{Val: 3}
	pRoot1 := &models.TreeNode{Val: 1}
	pRoot1.Left = pLeftNode1
	pRoot1.Right = pRightNode1

	qLeftNode1 := &models.TreeNode{Val: 2}
	qRightNode1 := &models.TreeNode{Val: 3}
	qRoot1 := &models.TreeNode{Val: 1}
	qRoot1.Left = qLeftNode1
	qRoot1.Right = qRightNode1

	fmt.Printf("\nIs same tree:%v", isSameTree(pRoot1, qRoot1))

	// Test case 2
	pLeftNode2 := &models.TreeNode{Val: 2}
	pRoot2 := &models.TreeNode{Val: 1}
	pRoot2.Left = pLeftNode2

	qRightNode2 := &models.TreeNode{Val: 2}
	qRoot2 := &models.TreeNode{Val: 1}
	qRoot2.Right = qRightNode2

	fmt.Printf("\nIs same tree:%v", isSameTree(pRoot2, qRoot2))

	// Test case 3
	pLeftNode3 := &models.TreeNode{Val: 2}
	pRightNode3 := &models.TreeNode{Val: 1}
	pRoot3 := &models.TreeNode{Val: 1}
	pRoot3.Left = pLeftNode3
	pRoot3.Right = pRightNode3

	qLeftNode3 := &models.TreeNode{Val: 1}
	qRightNode3 := &models.TreeNode{Val: 2}
	qRoot3 := &models.TreeNode{Val: 1}
	qRoot3.Left = qLeftNode3
	qRoot3.Right = qRightNode3

	fmt.Printf("\nIs same tree:%v", isSameTree(pRoot3, qRoot3))
	fmt.Println()
}

func isSameTree(p *models.TreeNode, q *models.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

/*
====================================
101. Symmetric Tree
====================================

🔗 Link:
https://leetcode.com/problems/symmetric-tree/description/

🧠 Idea:
- DFS (recursion)

- BFS (queue)

- Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
⏱ Complexity:
- Time: O(n)
- Space: O(h)

====================================
*/
func RunIsSymmetric() {
	leftNode1_LeftNode1 := &models.TreeNode{Val: 3}
	leftNode1_RightNode1 := &models.TreeNode{Val: 4}
	leftNode1 := &models.TreeNode{Val: 2, Left: leftNode1_LeftNode1, Right: leftNode1_RightNode1}

	rightNode1_LeftNode1 := &models.TreeNode{Val: 4}
	rightNode1_RightNode1 := &models.TreeNode{Val: 3}
	rightNode1 := &models.TreeNode{Val: 2, Left: rightNode1_LeftNode1, Right: rightNode1_RightNode1}

	root1 := &models.TreeNode{Val: 1, Left: leftNode1, Right: rightNode1}
	fmt.Printf("Is symmetric:%v", isSymmetric(root1))

	leftNode2_LeftNode2 := &models.TreeNode{Val: 3}
	leftNode2 := &models.TreeNode{Val: 2, Left: nil, Right: leftNode2_LeftNode2}

	rightNode2_RightNode2 := &models.TreeNode{Val: 3}
	rightNode2 := &models.TreeNode{Val: 2, Left: nil, Right: rightNode2_RightNode2}

	root2 := &models.TreeNode{Val: 1, Left: leftNode2, Right: rightNode2}
	fmt.Printf("\nIs symmetric:%v", isSymmetric(root2))

	fmt.Println()
}

func isSymmetric(root *models.TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(lNode *models.TreeNode, rNode *models.TreeNode) bool {
	if lNode == nil && rNode == nil {
		return true
	}

	if lNode.Val != rNode.Val {
		return false
	}

	if (lNode == nil && rNode != nil) || (lNode != nil && rNode == nil) {
		return false
	}

	if lNode.Left == nil && lNode.Right == nil && rNode.Left == nil && rNode.Right == nil {
		if lNode.Val == rNode.Val {
			return true
		}
	}

	return checkSymmetric(lNode.Left, rNode.Right) && checkSymmetric(lNode.Right, rNode.Left)
}

func RunMaxDepth() {
	node3_child5 := &models.Node{Val: 5}
	node3_child6 := &models.Node{Val: 6}

	node3 := &models.Node{Val: 3, Children: []*models.Node{node3_child5, node3_child6}}
	node2 := &models.Node{Val: 2}
	node4 := &models.Node{Val: 4}

	root1 := &models.Node{Val: 1, Children: []*models.Node{node3, node2, node4}}

	fmt.Printf("Max depth:%d", maxDepth(root1))
	fmt.Println()

	// Tầng sâu nhất (Tầng 5)
	node14 := &models.Node{Val: 14}

	// Tầng 4
	node11 := &models.Node{Val: 11, Children: []*models.Node{node14}}
	node12 := &models.Node{Val: 12}
	node13 := &models.Node{Val: 13}

	// Tầng 3
	node6 := &models.Node{Val: 6}
	node7 := &models.Node{Val: 7, Children: []*models.Node{node11}}
	node8 := &models.Node{Val: 8, Children: []*models.Node{node12}}
	node9 := &models.Node{Val: 9, Children: []*models.Node{node13}}
	node10 := &models.Node{Val: 10}

	// Tầng 2
	node22 := &models.Node{Val: 2}
	node33 := &models.Node{Val: 3, Children: []*models.Node{node6, node7}}
	node44 := &models.Node{Val: 4, Children: []*models.Node{node8}}
	node55 := &models.Node{Val: 5, Children: []*models.Node{node9, node10}}

	// Tầng 1 (Gốc)
	root2 := &models.Node{
		Val:      1,
		Children: []*models.Node{node22, node33, node44, node55},
	}

	fmt.Printf("Max depth:%d", maxDepth(root2))
	fmt.Println()
}

func maxDepth(root *models.Node) int {
	if root == nil {
		return 0
	}

	if root.Children == nil {
		return 1
	}

	return calculateHeight(root)
}

func calculateHeight(node *models.Node) int {
	if node.Children == nil {
		return 1
	}

	maxHeight := 0

	for _, child := range node.Children {
		nodeHeight := calculateHeight(child)
		if nodeHeight > maxHeight {
			maxHeight = nodeHeight
		}
	}

	return maxHeight + 1
}

func RunHasPathSumI() {
	leftNode11_Node7 := &models.TreeNode{Val: 7}
	rightNode11_Node2 := &models.TreeNode{Val: 2}
	leftNode1_Node11 := &models.TreeNode{Val: 11, Left: leftNode11_Node7, Right: rightNode11_Node2}
	leftNode1 := &models.TreeNode{Val: 4, Left: leftNode1_Node11}

	leftNode8_Node13 := &models.TreeNode{Val: 13}
	rightNode8_Node4 := &models.TreeNode{Val: 4}
	rightNode1 := &models.TreeNode{Val: 8, Left: leftNode8_Node13, Right: rightNode8_Node4}

	root1 := &models.TreeNode{Val: 5, Left: leftNode1, Right: rightNode1}
	fmt.Printf("Has path sum:%v", hasPathSum(root1, 22))
	fmt.Println()

	leftNode2 := &models.TreeNode{Val: 2}
	rightNode2 := &models.TreeNode{Val: 3}
	root2 := &models.TreeNode{Val: 1, Left: leftNode2, Right: rightNode2}
	fmt.Printf("Has path sum:%v", hasPathSum(root2, 5))
	fmt.Println()
}

func hasPathSum(root *models.TreeNode, targetSum int) bool {
	// if the tree is empty, there are no root to leaf paths
	if root == nil {
		return false
	}
	// check if the node is a leaf.
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	return checkPath(root, targetSum)
}

func checkPath(node *models.TreeNode, targetSum int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil && node.Val == targetSum {
		return true
	}
	targetSum -= node.Val
	return checkPath(node.Left, targetSum) || checkPath(node.Right, targetSum)
}

func RunInvertTree() {
	left2 := &models.TreeNode{Val: 1}
	right2 := &models.TreeNode{Val: 3}
	root2 := &models.TreeNode{Val: 2, Left: left2, Right: right2}

	fmt.Printf("Invert tree:%v", invertTree(root2))
}

func invertTree(root *models.TreeNode) *models.TreeNode {
	// check edge cases
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	leftNode := root.Left
	rightNode := root.Right

	// swap node
	root.Left = invertTree(leftNode)
	root.Right = invertTree(rightNode)

	return root
}

func RunIsBalanced() {
	childLeftNode15_1 := &models.TreeNode{Val: 15}
	childRightNode7_2 := &models.TreeNode{Val: 7}
	rightNode20_1 := &models.TreeNode{Val: 20, Left: childLeftNode15_1, Right: childRightNode7_2}

	leftNode9_1 := &models.TreeNode{Val: 9}
	rootNode3_1 := &models.TreeNode{Val: 3, Left: leftNode9_1, Right: rightNode20_1}

	fmt.Printf("is blance:%v", isBalanced(rootNode3_1))
	fmt.Println()

	childLeftNode4_2 := &models.TreeNode{Val: 4}
	childRightNode4_2 := &models.TreeNode{Val: 4}
	childLeftNode3_2 := &models.TreeNode{Val: 3, Left: childLeftNode4_2, Right: childRightNode4_2}

	childRightNode3_2 := &models.TreeNode{Val: 3}

	leftNode2_2 := &models.TreeNode{Val: 2, Left: childLeftNode3_2, Right: childRightNode3_2}
	rightNode2_2 := &models.TreeNode{Val: 2}
	rootNode1_2 := &models.TreeNode{Val: 1, Left: leftNode2_2, Right: rightNode2_2}

	fmt.Printf("is blance:%v", isBalanced(rootNode1_2))
	fmt.Println()

	childLeftNode4_1 := &models.TreeNode{Val: 4}

	childLeftNode3_1 := &models.TreeNode{
		Val:  3,
		Left: childLeftNode4_1,
	}

	leftNode2_1 := &models.TreeNode{
		Val:  2,
		Left: childLeftNode3_1,
	}

	childRightNode4_1 := &models.TreeNode{Val: 4}

	childRightNode3_1 := &models.TreeNode{
		Val:   3,
		Right: childRightNode4_1,
	}

	rightNode2_1 := &models.TreeNode{
		Val:   2,
		Right: childRightNode3_1,
	}

	rootNode1_1 := &models.TreeNode{
		Val:   1,
		Left:  leftNode2_1,
		Right: rightNode2_1,
	}

	fmt.Printf("is balance:%v", isBalanced(rootNode1_1))
	fmt.Println()
}

func isBalanced(root *models.TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	return calculateNodeHeight(root) != -1
}

func calculateNodeHeight(node *models.TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}

	nodeLeftHeight := calculateNodeHeight(node.Left)
	if nodeLeftHeight == -1 {
		return -1
	}
	nodeRightHeight := calculateNodeHeight(node.Right)
	if nodeRightHeight == -1 {
		return -1
	}
	res := absInt(nodeLeftHeight - nodeRightHeight)
	if res > 1 {
		return -1
	}
	if nodeLeftHeight >= nodeRightHeight {
		return nodeLeftHeight + 1
	}
	return nodeRightHeight + 1
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func RunBinaryTreePaths() {
	childLeft1 := &models.TreeNode{Val: 5}
	leftNode := &models.TreeNode{Val: 2, Left: childLeft1}
	rightNode := &models.TreeNode{Val: 3}

	root := &models.TreeNode{
		Val:   1,
		Left:  leftNode,
		Right: rightNode,
	}
	paths := binaryTreePaths(root)
	fmt.Printf("Binary tree path:%v", paths)
	fmt.Println()

	root2 := &models.TreeNode{
		Val: 1,
	}
	paths2 := binaryTreePaths(root2)
	fmt.Printf("Binary tree path:%v", paths2)
	fmt.Println()
}

func binaryTreePaths(root *models.TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	return findTreePaths(root)
}

func findTreePaths(node *models.TreeNode) []string {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return []string{strconv.Itoa(node.Val)}
	}
	paths := []string{}

	leftPaths := findTreePaths(node.Left)
	for _, leftPath := range leftPaths {
		path := strconv.Itoa(node.Val) + "->" + leftPath
		paths = append(paths, path)
	}

	rightPaths := findTreePaths(node.Right)
	for _, rightPath := range rightPaths {
		path := strconv.Itoa(node.Val) + "->" + rightPath
		paths = append(paths, path)
	}

	return paths
}

func RunPathSum() {

	leftNode_node07 := &models.TreeNode{Val: 7}
	rightNode_node02 := &models.TreeNode{Val: 2}
	leftNode_node11 := &models.TreeNode{Val: 11, Left: leftNode_node07, Right: rightNode_node02}

	leftNode := &models.TreeNode{Val: 4, Left: leftNode_node11}

	leftNode_node13 := &models.TreeNode{Val: 13}
	leftNode_node05 := &models.TreeNode{Val: 5}
	rightNode_node01 := &models.TreeNode{Val: 1}
	rightNode_node04 := &models.TreeNode{Val: 4, Left: leftNode_node05, Right: rightNode_node01}
	rightNode := &models.TreeNode{Val: 8, Left: leftNode_node13, Right: rightNode_node04}

	root := &models.TreeNode{Val: 5, Left: leftNode, Right: rightNode}
	paths := pathSum(root, 22)
	fmt.Printf("Paths:%v", paths)
	fmt.Println()
}

func pathSum(root *models.TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return [][]int{{targetSum}}
	}

	var paths *[][]int = new([][]int)
	var arr *[]int = new([]int)

	findPathSum(root, targetSum, paths, arr)
	return *paths
}

func findPathSum(node *models.TreeNode, targetSum int, paths *[][]int, arr *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil && node.Val == targetSum {
		*arr = append(*arr, node.Val)
		tempArr := make([]int, len(*arr))
		copy(tempArr, *arr)
		*paths = append(*paths, tempArr)
	} else {
		*arr = append(*arr, node.Val)
		targetSum = targetSum - node.Val

		findPathSum(node.Left, targetSum, paths, arr)
		findPathSum(node.Right, targetSum, paths, arr)
	}
	// remove the last element
	if len(*arr) > 0 {
		*arr = (*arr)[:len(*arr)-1]
	}
}
