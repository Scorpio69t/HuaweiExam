package main

import "fmt"

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：前缀和 + 哈希表解法（推荐）
// 时间复杂度：O(n)，空间复杂度：O(h)
func pathSum(root *TreeNode, targetSum int) int {
	prefixMap := make(map[int]int)
	prefixMap[0] = 1 // 初始化，空路径的前缀和为0
	return dfs(root, 0, targetSum, prefixMap)
}

func dfs(node *TreeNode, currentSum, targetSum int, prefixMap map[int]int) int {
	if node == nil {
		return 0
	}

	currentSum += node.Val
	count := prefixMap[currentSum-targetSum] // 查找目标前缀和的个数

	prefixMap[currentSum]++ // 将当前前缀和加入哈希表

	// 递归处理左右子树
	count += dfs(node.Left, currentSum, targetSum, prefixMap)
	count += dfs(node.Right, currentSum, targetSum, prefixMap)

	prefixMap[currentSum]-- // 回溯，移除当前前缀和

	return count
}

// 方法二：双重DFS解法
// 时间复杂度：O(n²)，空间复杂度：O(h)
func pathSumDoubleDFS(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// 以当前节点为起点的路径和
	count := dfsFromNode(root, targetSum, 0)

	// 递归处理左右子树
	count += pathSumDoubleDFS(root.Left, targetSum)
	count += pathSumDoubleDFS(root.Right, targetSum)

	return count
}

func dfsFromNode(node *TreeNode, targetSum, currentSum int) int {
	if node == nil {
		return 0
	}

	currentSum += node.Val
	count := 0

	if currentSum == targetSum {
		count++
	}

	// 继续向下搜索
	count += dfsFromNode(node.Left, targetSum, currentSum)
	count += dfsFromNode(node.Right, targetSum, currentSum)

	return count
}

// 方法三：递归回溯解法
// 时间复杂度：O(n²)，空间复杂度：O(h)
func pathSumBacktrack(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// 从根节点开始的所有路径
	count := backtrack(root, targetSum, 0)

	// 递归处理左右子树
	count += pathSumBacktrack(root.Left, targetSum)
	count += pathSumBacktrack(root.Right, targetSum)

	return count
}

func backtrack(node *TreeNode, targetSum, currentSum int) int {
	if node == nil {
		return 0
	}

	currentSum += node.Val
	count := 0

	if currentSum == targetSum {
		count++
	}

	// 继续向下搜索
	count += backtrack(node.Left, targetSum, currentSum)
	count += backtrack(node.Right, targetSum, currentSum)

	return count
}

// 方法四：优化的前缀和解法（更清晰的实现）
func pathSumOptimized(root *TreeNode, targetSum int) int {
	prefixMap := make(map[int]int)
	prefixMap[0] = 1
	return dfsOptimized(root, 0, targetSum, prefixMap)
}

func dfsOptimized(node *TreeNode, currentSum, targetSum int, prefixMap map[int]int) int {
	if node == nil {
		return 0
	}

	currentSum += node.Val
	target := currentSum - targetSum

	// 查找目标前缀和的个数
	count := 0
	if freq, exists := prefixMap[target]; exists {
		count = freq
	}

	// 更新当前前缀和的频率
	prefixMap[currentSum]++

	// 递归处理左右子树
	count += dfsOptimized(node.Left, currentSum, targetSum, prefixMap)
	count += dfsOptimized(node.Right, currentSum, targetSum, prefixMap)

	// 回溯，移除当前前缀和
	prefixMap[currentSum]--

	return count
}

// 辅助函数：根据层序遍历数组构建二叉树
func buildTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// 左子节点
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// 右子节点
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// 辅助函数：打印二叉树（层序遍历）
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("空树")
		return
	}

	queue := []*TreeNode{root}
	var result []interface{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}

	// 移除末尾的nil
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	fmt.Printf("二叉树: %v\n", result)
}

func main() {
	fmt.Println("=== 437. 路径总和 III ===")

	// 测试用例1
	root1 := buildTree([]interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1})
	targetSum1 := 8
	fmt.Printf("测试用例1: targetSum=%d\n", targetSum1)
	printTree(root1)
	fmt.Printf("前缀和+哈希表解法结果: %d\n", pathSum(root1, targetSum1))
	fmt.Printf("双重DFS解法结果: %d\n", pathSumDoubleDFS(root1, targetSum1))
	fmt.Printf("递归回溯解法结果: %d\n", pathSumBacktrack(root1, targetSum1))
	fmt.Printf("优化前缀和解法结果: %d\n", pathSumOptimized(root1, targetSum1))
	fmt.Println()

	// 测试用例2
	root2 := buildTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1})
	targetSum2 := 22
	fmt.Printf("测试用例2: targetSum=%d\n", targetSum2)
	printTree(root2)
	fmt.Printf("前缀和+哈希表解法结果: %d\n", pathSum(root2, targetSum2))
	fmt.Printf("双重DFS解法结果: %d\n", pathSumDoubleDFS(root2, targetSum2))
	fmt.Printf("递归回溯解法结果: %d\n", pathSumBacktrack(root2, targetSum2))
	fmt.Printf("优化前缀和解法结果: %d\n", pathSumOptimized(root2, targetSum2))
	fmt.Println()

	// 测试用例3（包含负数）
	root3 := buildTree([]interface{}{1, -2, -3, 1, 3, -2, nil, -1})
	targetSum3 := -1
	fmt.Printf("测试用例3: targetSum=%d\n", targetSum3)
	printTree(root3)
	fmt.Printf("前缀和+哈希表解法结果: %d\n", pathSum(root3, targetSum3))
	fmt.Printf("双重DFS解法结果: %d\n", pathSumDoubleDFS(root3, targetSum3))
	fmt.Printf("递归回溯解法结果: %d\n", pathSumBacktrack(root3, targetSum3))
	fmt.Printf("优化前缀和解法结果: %d\n", pathSumOptimized(root3, targetSum3))
	fmt.Println()

	// 边界测试用例
	testCases := []struct {
		values    []interface{}
		targetSum int
		desc      string
	}{
		{[]interface{}{1}, 1, "单个节点"},
		{[]interface{}{1, 2, 3}, 3, "简单三节点"},
		{[]interface{}{1, 2, 3}, 6, "整个树的和"},
		{[]interface{}{0, 1, 1}, 1, "包含零值"},
		{[]interface{}{}, 0, "空树"},
	}

	for _, tc := range testCases {
		var root *TreeNode
		if len(tc.values) > 0 {
			root = buildTree(tc.values)
		}
		result := pathSum(root, tc.targetSum)
		fmt.Printf("%s: targetSum=%d, 结果=%d\n", tc.desc, tc.targetSum, result)
	}

	// 算法正确性验证
	fmt.Println("\n=== 算法正确性验证 ===")
	simpleRoot := buildTree([]interface{}{1, 2, 3})
	simpleTarget := 3
	fmt.Printf("验证树: targetSum=%d\n", simpleTarget)
	printTree(simpleRoot)
	fmt.Printf("前缀和解法: %d\n", pathSum(simpleRoot, simpleTarget))
	fmt.Printf("双重DFS解法: %d\n", pathSumDoubleDFS(simpleRoot, simpleTarget))
	fmt.Printf("递归回溯解法: %d\n", pathSumBacktrack(simpleRoot, simpleTarget))
	fmt.Printf("优化前缀和解法: %d\n", pathSumOptimized(simpleRoot, simpleTarget))

	// 验证所有解法结果一致
	if pathSum(simpleRoot, simpleTarget) == pathSumDoubleDFS(simpleRoot, simpleTarget) &&
		pathSum(simpleRoot, simpleTarget) == pathSumBacktrack(simpleRoot, simpleTarget) &&
		pathSum(simpleRoot, simpleTarget) == pathSumOptimized(simpleRoot, simpleTarget) {
		fmt.Println("✅ 所有解法结果一致，算法正确！")
	} else {
		fmt.Println("❌ 解法结果不一致，需要检查！")
	}
}
