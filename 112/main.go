package main

import (
	"fmt"
)

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// hasPathSum 方法一：递归解法（推荐）
// 时间复杂度：O(n)，空间复杂度：O(h)
func hasPathSum(root *TreeNode, targetSum int) bool {
	// 边界条件：空节点
	if root == nil {
		return false
	}

	// 叶子节点：检查节点值是否等于剩余目标和
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	// 非叶子节点：递归检查左右子树
	// 更新目标和为 targetSum - root.Val
	remainingSum := targetSum - root.Val
	return hasPathSum(root.Left, remainingSum) || hasPathSum(root.Right, remainingSum)
}

// hasPathSumIterative 方法二：迭代解法（使用栈）
// 时间复杂度：O(n)，空间复杂度：O(n)
func hasPathSumIterative(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// 使用两个栈分别存储节点和对应的剩余目标和
	nodeStack := []*TreeNode{root}
	sumStack := []int{targetSum}

	for len(nodeStack) > 0 {
		// 弹出栈顶元素
		node := nodeStack[len(nodeStack)-1]
		remainingSum := sumStack[len(sumStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		sumStack = sumStack[:len(sumStack)-1]

		// 检查是否为叶子节点
		if node.Left == nil && node.Right == nil {
			if node.Val == remainingSum {
				return true
			}
			continue
		}

		// 将子节点和对应的剩余和压入栈
		newRemainingSum := remainingSum - node.Val
		if node.Left != nil {
			nodeStack = append(nodeStack, node.Left)
			sumStack = append(sumStack, newRemainingSum)
		}
		if node.Right != nil {
			nodeStack = append(nodeStack, node.Right)
			sumStack = append(sumStack, newRemainingSum)
		}
	}

	return false
}

// hasPathSumDFS 方法三：DFS辅助函数解法
// 时间复杂度：O(n)，空间复杂度：O(h)
func hasPathSumDFS(root *TreeNode, targetSum int) bool {
	var dfs func(*TreeNode, int, int) bool

	dfs = func(node *TreeNode, currentSum, target int) bool {
		if node == nil {
			return false
		}

		currentSum += node.Val

		// 叶子节点检查
		if node.Left == nil && node.Right == nil {
			return currentSum == target
		}

		// 递归检查左右子树
		return dfs(node.Left, currentSum, target) || dfs(node.Right, currentSum, target)
	}

	return dfs(root, 0, targetSum)
}

// createTree 构建测试用的二叉树
func createTree() *TreeNode {
	// 构建示例1的树：[5,4,8,11,null,13,4,7,2,null,null,null,1]
	//       5
	//      / \
	//     4   8
	//    /   / \
	//   11  13  4
	//  / \      \
	// 7   2      1

	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}

	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}

	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 1}

	return root
}

// createSimpleTree 构建示例2的简单树
func createSimpleTree() *TreeNode {
	// 构建示例2的树：[1,2,3]
	//   1
	//  / \
	// 2   3

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	return root
}

// printTree 打印二叉树（中序遍历）
func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	printTree(root.Left)
	fmt.Printf("%d ", root.Val)
	printTree(root.Right)
}

// runTests 运行所有测试用例
func runTests() {
	fmt.Println("=== 112. 路径总和 测试用例 ===")

	// 测试用例1：示例1
	fmt.Println("测试用例1：示例1")
	tree1 := createTree()
	fmt.Print("树的中序遍历：")
	printTree(tree1)
	fmt.Println()

	targetSum1 := 22
	result1_1 := hasPathSum(tree1, targetSum1)
	result1_2 := hasPathSumIterative(tree1, targetSum1)
	result1_3 := hasPathSumDFS(tree1, targetSum1)

	fmt.Printf("目标和：%d\n", targetSum1)
	fmt.Printf("递归解法结果：%t\n", result1_1)
	fmt.Printf("迭代解法结果：%t\n", result1_2)
	fmt.Printf("DFS解法结果：%t\n", result1_3)
	fmt.Printf("预期结果：true（路径：5→4→11→2 = 22）\n\n")

	// 测试用例2：示例2
	fmt.Println("测试用例2：示例2")
	tree2 := createSimpleTree()
	fmt.Print("树的中序遍历：")
	printTree(tree2)
	fmt.Println()

	targetSum2 := 5
	result2_1 := hasPathSum(tree2, targetSum2)
	result2_2 := hasPathSumIterative(tree2, targetSum2)
	result2_3 := hasPathSumDFS(tree2, targetSum2)

	fmt.Printf("目标和：%d\n", targetSum2)
	fmt.Printf("递归解法结果：%t\n", result2_1)
	fmt.Printf("迭代解法结果：%t\n", result2_2)
	fmt.Printf("DFS解法结果：%t\n", result2_3)
	fmt.Printf("预期结果：false（路径1→2=3, 1→3=4，都不等于5）\n\n")

	// 测试用例3：示例3（空树）
	fmt.Println("测试用例3：空树")
	var tree3 *TreeNode = nil
	targetSum3 := 0
	result3_1 := hasPathSum(tree3, targetSum3)
	result3_2 := hasPathSumIterative(tree3, targetSum3)
	result3_3 := hasPathSumDFS(tree3, targetSum3)

	fmt.Printf("目标和：%d\n", targetSum3)
	fmt.Printf("递归解法结果：%t\n", result3_1)
	fmt.Printf("迭代解法结果：%t\n", result3_2)
	fmt.Printf("DFS解法结果：%t\n", result3_3)
	fmt.Printf("预期结果：false\n\n")

	// 测试用例4：单节点树
	fmt.Println("测试用例4：单节点树")
	tree4 := &TreeNode{Val: 5}
	targetSum4 := 5
	result4_1 := hasPathSum(tree4, targetSum4)
	result4_2 := hasPathSumIterative(tree4, targetSum4)
	result4_3 := hasPathSumDFS(tree4, targetSum4)

	fmt.Printf("目标和：%d\n", targetSum4)
	fmt.Printf("递归解法结果：%t\n", result4_1)
	fmt.Printf("迭代解法结果：%t\n", result4_2)
	fmt.Printf("DFS解法结果：%t\n", result4_3)
	fmt.Printf("预期结果：true\n\n")

	// 测试用例5：负数节点
	fmt.Println("测试用例5：负数节点")
	tree5 := &TreeNode{Val: -3}
	tree5.Left = &TreeNode{Val: -2}
	targetSum5 := -5
	result5_1 := hasPathSum(tree5, targetSum5)
	result5_2 := hasPathSumIterative(tree5, targetSum5)
	result5_3 := hasPathSumDFS(tree5, targetSum5)

	fmt.Printf("目标和：%d\n", targetSum5)
	fmt.Printf("递归解法结果：%t\n", result5_1)
	fmt.Printf("迭代解法结果：%t\n", result5_2)
	fmt.Printf("DFS解法结果：%t\n", result5_3)
	fmt.Printf("预期结果：true（路径：-3→-2 = -5）\n\n")
}

func main() {
	runTests()
}
