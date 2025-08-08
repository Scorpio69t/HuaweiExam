package main

import (
	"fmt"
	"strings"
	"time"
)

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ========== 方法1: 树形动态规划（最优解） ==========
func rob1(root *TreeNode) int {
	result := robHelper1(root)
	return max(result[0], result[1])
}

// 返回 [不偷当前节点的最大值, 偷当前节点的最大值]
func robHelper1(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}

	// 递归计算左右子树的状态
	left := robHelper1(node.Left)
	right := robHelper1(node.Right)

	// 不偷当前节点：左右子节点可偷可不偷，取最大值
	notRob := max(left[0], left[1]) + max(right[0], right[1])

	// 偷当前节点：左右子节点都不能偷
	rob := node.Val + left[0] + right[0]

	return []int{notRob, rob}
}

// ========== 方法2: 记忆化递归 ==========
func rob2(root *TreeNode) int {
	memo := make(map[*TreeNode]int)
	return robHelper2(root, memo)
}

func robHelper2(node *TreeNode, memo map[*TreeNode]int) int {
	if node == nil {
		return 0
	}

	// 检查缓存
	if val, exists := memo[node]; exists {
		return val
	}

	// 偷当前节点：不能偷左右子节点，但可以偷孙子节点
	robCurrent := node.Val
	if node.Left != nil {
		robCurrent += robHelper2(node.Left.Left, memo) + robHelper2(node.Left.Right, memo)
	}
	if node.Right != nil {
		robCurrent += robHelper2(node.Right.Left, memo) + robHelper2(node.Right.Right, memo)
	}

	// 不偷当前节点：可以偷左右子节点
	notRobCurrent := robHelper2(node.Left, memo) + robHelper2(node.Right, memo)

	// 取最大值并缓存
	result := max(robCurrent, notRobCurrent)
	memo[node] = result
	return result
}

// ========== 方法3: 暴力递归（会超时，仅用于对比） ==========
func rob3(root *TreeNode) int {
	return robHelper3(root)
}

func robHelper3(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// 偷当前节点：不能偷左右子节点，但可以偷孙子节点
	robCurrent := node.Val
	if node.Left != nil {
		robCurrent += robHelper3(node.Left.Left) + robHelper3(node.Left.Right)
	}
	if node.Right != nil {
		robCurrent += robHelper3(node.Right.Left) + robHelper3(node.Right.Right)
	}

	// 不偷当前节点：可以偷左右子节点
	notRobCurrent := robHelper3(node.Left) + robHelper3(node.Right)

	return max(robCurrent, notRobCurrent)
}

// ========== 方法4: 改进的树形DP（一次遍历） ==========
func rob4(root *TreeNode) int {
	robbed, notRobbed := robHelper4(root)
	return max(robbed, notRobbed)
}

func robHelper4(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	// 递归获取左右子树的状态
	leftRobbed, leftNotRobbed := robHelper4(node.Left)
	rightRobbed, rightNotRobbed := robHelper4(node.Right)

	// 偷当前节点：左右子节点都不能偷
	robbed := node.Val + leftNotRobbed + rightNotRobbed

	// 不偷当前节点：左右子节点可偷可不偷
	notRobbed := max(leftRobbed, leftNotRobbed) + max(rightRobbed, rightNotRobbed)

	return robbed, notRobbed
}

// ========== 方法5: 状态压缩优化 ==========
func rob5(root *TreeNode) int {
	states := robHelper5(root)
	return max(states[0], states[1])
}

func robHelper5(node *TreeNode) [2]int {
	if node == nil {
		return [2]int{0, 0} // [不偷, 偷]
	}

	leftStates := robHelper5(node.Left)
	rightStates := robHelper5(node.Right)

	// 不偷当前节点
	notRob := max(leftStates[0], leftStates[1]) + max(rightStates[0], rightStates[1])

	// 偷当前节点
	rob := node.Val + leftStates[0] + rightStates[0]

	return [2]int{notRob, rob}
}

// ========== 工具函数 ==========

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 根据数组构建二叉树
func buildTree(values []interface{}) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	index := 1

	for len(queue) > 0 && index < len(values) {
		node := queue[0]
		queue = queue[1:]

		// 左子节点
		if index < len(values) && values[index] != nil {
			node.Left = &TreeNode{Val: values[index].(int)}
			queue = append(queue, node.Left)
		}
		index++

		// 右子节点
		if index < len(values) && values[index] != nil {
			node.Right = &TreeNode{Val: values[index].(int)}
			queue = append(queue, node.Right)
		}
		index++
	}

	return root
}

// 层序打印树结构
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("Empty tree")
		return
	}

	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		size := len(queue)
		fmt.Printf("Level %d: ", level)

		allNull := true
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node != nil {
				fmt.Printf("%d ", node.Val)
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
				allNull = false
			} else {
				fmt.Print("null ")
				queue = append(queue, nil, nil)
			}
		}
		fmt.Println()
		level++

		// 如果下一层全是null，停止打印
		if allNull {
			break
		}
	}
}

// 计算树的节点数量
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// 生成测试树
func generateTestTree(size int) *TreeNode {
	if size <= 0 {
		return nil
	}

	values := make([]interface{}, size)
	for i := 0; i < size; i++ {
		values[i] = i + 1 // 节点值为1到size
	}

	return buildTree(values)
}

// ========== 测试和性能评估 ==========
func main() {
	// 测试用例
	testCases := []struct {
		name     string
		treeVals []interface{}
		expected int
	}{
		{
			name:     "示例1: 不规则树",
			treeVals: []interface{}{3, 2, 3, nil, 3, nil, 1},
			expected: 7,
		},
		{
			name:     "示例2: 平衡树",
			treeVals: []interface{}{3, 4, 5, 1, 3, nil, 1},
			expected: 9,
		},
		{
			name:     "测试3: 单节点",
			treeVals: []interface{}{5},
			expected: 5,
		},
		{
			name:     "测试4: 两层树",
			treeVals: []interface{}{2, 1, 3},
			expected: 4, // 正确答案：偷左右子节点1+3=4，比偷根节点2更优
		},
		{
			name:     "测试5: 链式树",
			treeVals: []interface{}{2, 1, nil, 4, nil},
			expected: 6,
		},
		{
			name:     "测试6: 完全二叉树",
			treeVals: []interface{}{4, 1, nil, 2, nil, 3},
			expected: 7,
		},
		{
			name:     "测试7: 大值根节点",
			treeVals: []interface{}{10, 5, 5, 1, 1, 1, 1},
			expected: 14,
		},
		{
			name:     "测试8: 空树",
			treeVals: []interface{}{},
			expected: 0,
		},
	}

	// 算法方法
	methods := []struct {
		name string
		fn   func(*TreeNode) int
	}{
		{"树形DP(最优)", rob1},
		{"记忆化递归", rob2},
		// {"暴力递归", rob3}, // 会超时，注释掉
		{"改进树形DP", rob4},
		{"状态压缩DP", rob5},
	}

	fmt.Println("=== LeetCode 337. 打家劫舍 III - 测试结果 ===")
	fmt.Println()

	// 运行测试
	for _, tc := range testCases {
		fmt.Printf("测试用例: %s\n", tc.name)
		fmt.Printf("树结构: %v\n", tc.treeVals)

		// 构建测试树
		root := buildTree(tc.treeVals)
		fmt.Println("二叉树结构:")
		printTree(root)

		allPassed := true
		var results []int
		var times []time.Duration

		for _, method := range methods {
			start := time.Now()
			result := method.fn(root)
			elapsed := time.Since(start)

			results = append(results, result)
			times = append(times, elapsed)

			status := "✅"
			if result != tc.expected {
				status = "❌"
				allPassed = false
			}

			fmt.Printf("  %s: %s (结果: %d, 耗时: %v)\n", method.name, status, result, elapsed)
		}

		fmt.Printf("期望结果: %d\n", tc.expected)

		if allPassed {
			fmt.Println("✅ 所有方法均通过")
		} else {
			fmt.Println("❌ 存在失败的方法")
		}
		fmt.Println(strings.Repeat("-", 60))
	}

	// 性能对比测试
	fmt.Println("\n=== 性能对比测试 ===")
	performanceTest()

	// 算法特性总结
	fmt.Println("\n=== 算法特性总结 ===")
	fmt.Println("1. 树形DP(最优):")
	fmt.Println("   - 时间复杂度: O(N)")
	fmt.Println("   - 空间复杂度: O(H)")
	fmt.Println("   - 特点: 最优解法，一次遍历解决")
	fmt.Println()
	fmt.Println("2. 记忆化递归:")
	fmt.Println("   - 时间复杂度: O(N)")
	fmt.Println("   - 空间复杂度: O(N)")
	fmt.Println("   - 特点: 避免重复计算，容易理解")
	fmt.Println()
	fmt.Println("3. 改进树形DP:")
	fmt.Println("   - 时间复杂度: O(N)")
	fmt.Println("   - 空间复杂度: O(H)")
	fmt.Println("   - 特点: 返回元组，思路清晰")
	fmt.Println()
	fmt.Println("4. 状态压缩DP:")
	fmt.Println("   - 时间复杂度: O(N)")
	fmt.Println("   - 空间复杂度: O(H)")
	fmt.Println("   - 特点: 数组存储状态，内存优化")

	// 打家劫舍演示
	fmt.Println("\n=== 打家劫舍演示 ===")
	demoRobHouse()
}

// 性能测试
func performanceTest() {
	sizes := []int{100, 500, 1000, 5000}
	methods := []struct {
		name string
		fn   func(*TreeNode) int
	}{
		{"树形DP", rob1},
		{"记忆化递归", rob2},
		{"改进DP", rob4},
		{"状态压缩", rob5},
	}

	for _, size := range sizes {
		fmt.Printf("性能测试 - 树大小: %d个节点\n", size)

		// 生成测试树
		root := generateTestTree(size)

		for _, method := range methods {
			start := time.Now()
			result := method.fn(root)
			elapsed := time.Since(start)

			fmt.Printf("  %s: 最大金额=%d, 耗时=%v\n",
				method.name, result, elapsed)
		}
	}
}

// 打家劫舍演示
func demoRobHouse() {
	fmt.Println("构建示例房屋树:")
	treeVals := []interface{}{3, 4, 5, 1, 3, nil, 1}
	root := buildTree(treeVals)

	fmt.Printf("房屋布局: %v\n", treeVals)
	fmt.Println("房屋结构:")
	printTree(root)

	fmt.Println("分析最优偷窃策略:")
	result := rob1(root)

	fmt.Printf("最大能偷窃金额: %d\n", result)
	fmt.Println("策略分析:")
	fmt.Println("  - 如果偷根节点(3)，不能偷子节点(4,5)")
	fmt.Println("  - 如果不偷根节点，可以偷子节点")
	fmt.Println("  - 最优策略：偷左子节点4和右子节点5，总计9")
	fmt.Println("打家劫舍完成!")
}
