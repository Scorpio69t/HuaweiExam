package main

import (
	"fmt"
	"strings"
)

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =========================== 方法一：BFS + 反转（最直观） ===========================

// levelOrderBottom BFS层序遍历 + 反转结果
// 时间复杂度：O(n)，n为节点数
// 空间复杂度：O(n)，队列和结果数组
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	// 正常BFS遍历
	for len(queue) > 0 {
		size := len(queue)
		level := []int{}

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
	}

	// 反转结果
	reverse(result)
	return result
}

// reverse 反转二维数组
func reverse(arr [][]int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

// =========================== 方法二：BFS + 头插法（避免反转） ===========================

// levelOrderBottom2 BFS + 头部插入，避免最后反转
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func levelOrderBottom2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		level := []int{}

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 头部插入（避免最后反转）
		result = append([][]int{level}, result...)
	}

	return result
}

// =========================== 方法三：DFS 递归（代码简洁） ===========================

// levelOrderBottom3 DFS递归实现
// 时间复杂度：O(n)
// 空间复杂度：O(h)，h为树高度
func levelOrderBottom3(root *TreeNode) [][]int {
	result := [][]int{}
	dfs(root, 0, &result)

	// 反转结果
	reverse(result)
	return result
}

// dfs 深度优先遍历
func dfs(node *TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}

	// 扩展结果数组
	if len(*result) <= level {
		*result = append(*result, []int{})
	}

	// 将当前节点加入对应层
	(*result)[level] = append((*result)[level], node.Val)

	// 递归遍历左右子树
	dfs(node.Left, level+1, result)
	dfs(node.Right, level+1, result)
}

// =========================== 方法四：预分配 + 反向索引 ===========================

// levelOrderBottom4 预分配空间 + 反向索引
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func levelOrderBottom4(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	// 先计算树的深度
	depth := getDepth(root)

	// 预分配空间
	result := make([][]int, depth)

	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		size := len(queue)
		// 反向索引：从底部开始填充
		reverseLevel := depth - level - 1
		result[reverseLevel] = make([]int, 0, size)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			result[reverseLevel] = append(result[reverseLevel], node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		level++
	}

	return result
}

// getDepth 计算树的深度
func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getDepth(root.Left), getDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// =========================== 辅助函数 ===========================

// buildTree 从数组构建二叉树（层序遍历方式）
// -1 表示 nil 节点
func buildTree(arr []int) *TreeNode {
	if len(arr) == 0 || arr[0] == -1 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]

		// 左子节点
		if i < len(arr) && arr[i] != -1 {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		// 右子节点
		if i < len(arr) && arr[i] != -1 {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// visualizeTree 可视化打印树结构
func visualizeTree(root *TreeNode, prefix string, isLeft bool) {
	if root == nil {
		return
	}

	fmt.Print(prefix)
	if isLeft {
		fmt.Print("├── ")
	} else {
		fmt.Print("└── ")
	}
	fmt.Println(root.Val)

	if root.Left != nil || root.Right != nil {
		if root.Left != nil {
			newPrefix := prefix
			if isLeft {
				newPrefix += "│   "
			} else {
				newPrefix += "    "
			}
			visualizeTree(root.Left, newPrefix, true)
		}

		if root.Right != nil {
			newPrefix := prefix
			if isLeft {
				newPrefix += "│   "
			} else {
				newPrefix += "    "
			}
			visualizeTree(root.Right, newPrefix, false)
		}
	}
}

// printResult 打印二维数组结果
func printResult(result [][]int) {
	fmt.Print("[")
	for i, level := range result {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print("[")
		for j, val := range level {
			if j > 0 {
				fmt.Print(",")
			}
			fmt.Print(val)
		}
		fmt.Print("]")
	}
	fmt.Println("]")
}

// resultsEqual 比较两个二维数组是否相等
func resultsEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

// =========================== 扩展功能 ===========================

// averageOfLevelsBottom 返回每层平均值（自底向上）
func averageOfLevelsBottom(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	result := []float64{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		sum := 0

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, float64(sum)/float64(size))
	}

	// 反转
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// largestValuesBottom 返回每层最大值（自底向上）
func largestValuesBottom(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		maxVal := queue[0].Val

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Val > maxVal {
				maxVal = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, maxVal)
	}

	// 反转
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// printTreeBottom 自底向上打印树（带缩进）
func printTreeBottom(root *TreeNode) {
	levels := levelOrderBottom(root)

	for i, level := range levels {
		indent := strings.Repeat("  ", len(levels)-i-1)
		fmt.Printf("%s层%d: %v\n", indent, len(levels)-i, level)
	}
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 107: 二叉树的层序遍历 II ===\n")

	// 测试用例
	testCases := []struct {
		name   string
		arr    []int
		expect [][]int
	}{
		{
			name:   "示例1: 标准二叉树",
			arr:    []int{3, 9, 20, -1, -1, 15, 7},
			expect: [][]int{{15, 7}, {9, 20}, {3}},
		},
		{
			name:   "示例2: 单节点",
			arr:    []int{1},
			expect: [][]int{{1}},
		},
		{
			name:   "示例3: 空树",
			arr:    []int{},
			expect: [][]int{},
		},
		{
			name:   "左偏树",
			arr:    []int{1, 2, -1, 3},
			expect: [][]int{{3}, {2}, {1}},
		},
		{
			name:   "右偏树",
			arr:    []int{1, -1, 2, -1, -1, -1, 3},
			expect: [][]int{{3}, {2}, {1}},
		},
		{
			name:   "完全二叉树",
			arr:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: [][]int{{4, 5, 6, 7}, {2, 3}, {1}},
		},
		{
			name:   "不平衡树",
			arr:    []int{1, 2, 3, 4, -1, -1, 5, 6, 7},
			expect: [][]int{{6, 7}, {4, 5}, {2, 3}, {1}},
		},
	}

	methods := []struct {
		name string
		fn   func(*TreeNode) [][]int
	}{
		{"方法一：BFS + 反转", levelOrderBottom},
		{"方法二：BFS + 头插法", levelOrderBottom2},
		{"方法三：DFS 递归", levelOrderBottom3},
		{"方法四：预分配 + 反向索引", levelOrderBottom4},
	}

	// 对每种方法运行测试
	for _, method := range methods {
		fmt.Printf("\n%s\n", method.name)
		fmt.Println(strings.Repeat("=", 60))

		passCount := 0
		for i, tc := range testCases {
			root := buildTree(tc.arr)
			result := method.fn(root)

			status := "✅"
			if !resultsEqual(result, tc.expect) {
				status = "❌"
			} else {
				passCount++
			}

			fmt.Printf("  测试%d: %s\n", i+1, status)
			fmt.Printf("    名称: %s\n", tc.name)
			fmt.Printf("    输入: %v\n", tc.arr)
			fmt.Printf("    输出: ")
			printResult(result)

			if !resultsEqual(result, tc.expect) {
				fmt.Printf("    期望: ")
				printResult(tc.expect)
			}

			// 为第一个示例打印树结构
			if i == 0 {
				fmt.Println("    树结构:")
				if root != nil {
					visualizeTree(root, "      ", false)
				}
			}
		}

		fmt.Printf("\n  通过: %d/%d\n", passCount, len(testCases))
	}

	// 扩展功能测试
	fmt.Println("\n\n=== 扩展功能测试 ===\n")
	testExtensions()

	// 与102题对比
	fmt.Println("\n=== 与102题对比 ===\n")
	compareWith102()

	// 性能对比
	fmt.Println("\n=== 性能对比 ===\n")
	performanceTest()
}

// testExtensions 测试扩展功能
func testExtensions() {
	fmt.Println("1. 每层平均值（自底向上）")
	root1 := buildTree([]int{3, 9, 20, -1, -1, 15, 7})
	averages := averageOfLevelsBottom(root1)
	fmt.Printf("   树: [3,9,20,null,null,15,7]\n")
	fmt.Printf("   每层平均值: %v\n", averages)

	fmt.Println("\n2. 每层最大值（自底向上）")
	root2 := buildTree([]int{1, 3, 2, 5, 3, -1, 9})
	maxValues := largestValuesBottom(root2)
	fmt.Printf("   树: [1,3,2,5,3,null,9]\n")
	fmt.Printf("   每层最大值: %v\n", maxValues)

	fmt.Println("\n3. 自底向上打印树（带缩进）")
	root3 := buildTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("   树: [1,2,3,4,5,6,7]")
	printTreeBottom(root3)

	fmt.Println("\n4. 树结构可视化")
	root4 := buildTree([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println("   树: [3,9,20,null,null,15,7]")
	visualizeTree(root4, "   ", false)
}

// compareWith102 与102题对比
func compareWith102() {
	fmt.Println("使用同一棵树，对比正序和倒序层序遍历")

	root := buildTree([]int{1, 2, 3, 4, 5, 6, 7})

	fmt.Println("  树结构:")
	visualizeTree(root, "    ", false)

	// 正序（102题）
	fmt.Println("\n  102题（自顶向下）:")
	result102 := levelOrder102(root)
	fmt.Print("    ")
	printResult(result102)

	// 倒序（107题）
	fmt.Println("\n  107题（自底向上）:")
	result107 := levelOrderBottom(root)
	fmt.Print("    ")
	printResult(result107)

	fmt.Println("\n  关键区别:")
	fmt.Println("  - 102题: 从根节点开始，逐层向下")
	fmt.Println("  - 107题: 从叶子节点开始，逐层向上")
	fmt.Println("  - 实现: 107题只需在102题基础上反转结果")
}

// levelOrder102 正常层序遍历（102题）
func levelOrder102(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		level := []int{}

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}

// performanceTest 性能测试
func performanceTest() {
	// 构建深度为10的完全二叉树
	size := (1 << 10) - 1 // 2^10 - 1 = 1023个节点
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i + 1
	}

	root := buildTree(arr)

	fmt.Printf("测试数据：完全二叉树，节点数=%d，深度=10\n\n", size)

	fmt.Println("各方法性能测试:")
	result1 := levelOrderBottom(root)
	fmt.Printf("  方法一（BFS + 反转）: 层数=%d\n", len(result1))

	result2 := levelOrderBottom2(root)
	fmt.Printf("  方法二（BFS + 头插法）: 层数=%d\n", len(result2))

	result3 := levelOrderBottom3(root)
	fmt.Printf("  方法三（DFS 递归）: 层数=%d\n", len(result3))

	result4 := levelOrderBottom4(root)
	fmt.Printf("  方法四（预分配 + 反向索引）: 层数=%d\n", len(result4))

	fmt.Println("\n说明：")
	fmt.Println("  - 方法一（BFS + 反转）：O(n)时间，O(n)空间，最直观")
	fmt.Println("  - 方法二（BFS + 头插法）：O(n)时间，O(n)空间，避免反转")
	fmt.Println("  - 方法三（DFS 递归）：O(n)时间，O(h)空间，代码简洁")
	fmt.Println("  - 方法四（预分配）：O(n)时间，O(n)空间，避免反转和头插")
}
