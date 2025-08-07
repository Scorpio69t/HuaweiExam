package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ========== 方法1: DFS递归回溯(值传递) ==========
func pathSum1(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var path []int
	dfs1(root, targetSum, path, &result)
	return result
}

func dfs1(node *TreeNode, remain int, path []int, result *[][]int) {
	if node == nil {
		return
	}

	// 添加当前节点到路径
	path = append(path, node.Val)
	remain -= node.Val

	// 如果是叶节点且路径和等于目标和
	if node.Left == nil && node.Right == nil && remain == 0 {
		// 复制路径，避免引用问题
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
		return
	}

	// 递归遍历左右子树
	dfs1(node.Left, remain, path, result)
	dfs1(node.Right, remain, path, result)
}

// ========== 方法2: DFS递归回溯(引用传递，手动回溯) ==========
func pathSum2(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var path []int
	dfs2(root, targetSum, &path, &result)
	return result
}

func dfs2(node *TreeNode, remain int, path *[]int, result *[][]int) {
	if node == nil {
		return
	}

	// 添加当前节点到路径
	*path = append(*path, node.Val)
	remain -= node.Val

	// 如果是叶节点且路径和等于目标和
	if node.Left == nil && node.Right == nil && remain == 0 {
		pathCopy := make([]int, len(*path))
		copy(pathCopy, *path)
		*result = append(*result, pathCopy)
	} else {
		// 继续搜索子树
		dfs2(node.Left, remain, path, result)
		dfs2(node.Right, remain, path, result)
	}

	// 回溯：移除当前节点
	*path = (*path)[:len(*path)-1]
}

// ========== 方法3: BFS层序遍历 ==========
func pathSum3(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	type pathInfo struct {
		node *TreeNode
		path []int
		sum  int
	}

	queue := []pathInfo{{root, []int{root.Val}, root.Val}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		node := current.node
		path := current.path
		sum := current.sum

		// 如果是叶节点
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				result = append(result, path)
			}
			continue
		}

		// 处理左子树
		if node.Left != nil {
			leftPath := make([]int, len(path))
			copy(leftPath, path)
			leftPath = append(leftPath, node.Left.Val)
			queue = append(queue, pathInfo{
				node: node.Left,
				path: leftPath,
				sum:  sum + node.Left.Val,
			})
		}

		// 处理右子树
		if node.Right != nil {
			rightPath := make([]int, len(path))
			copy(rightPath, path)
			rightPath = append(rightPath, node.Right.Val)
			queue = append(queue, pathInfo{
				node: node.Right,
				path: rightPath,
				sum:  sum + node.Right.Val,
			})
		}
	}

	return result
}

// ========== 方法4: DFS迭代栈实现 ==========
func pathSum4(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	type stackInfo struct {
		node *TreeNode
		path []int
		sum  int
	}

	stack := []stackInfo{{root, []int{root.Val}, root.Val}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := current.node
		path := current.path
		sum := current.sum

		// 如果是叶节点
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				result = append(result, path)
			}
			continue
		}

		// 右子树先入栈(后处理)
		if node.Right != nil {
			rightPath := make([]int, len(path))
			copy(rightPath, path)
			rightPath = append(rightPath, node.Right.Val)
			stack = append(stack, stackInfo{
				node: node.Right,
				path: rightPath,
				sum:  sum + node.Right.Val,
			})
		}

		// 左子树后入栈(先处理)
		if node.Left != nil {
			leftPath := make([]int, len(path))
			copy(leftPath, path)
			leftPath = append(leftPath, node.Left.Val)
			stack = append(stack, stackInfo{
				node: node.Left,
				path: leftPath,
				sum:  sum + node.Left.Val,
			})
		}
	}

	return result
}

// ========== 方法5: 优化DFS(预分配优化) ==========
func pathSum5(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	path := make([]int, 0, 1000) // 预分配容量
	dfs5(root, targetSum, path, &result)
	return result
}

func dfs5(node *TreeNode, remain int, path []int, result *[][]int) {
	if node == nil {
		return
	}

	// 添加当前节点
	path = append(path, node.Val)
	remain -= node.Val

	// 检查叶节点
	if node.Left == nil && node.Right == nil && remain == 0 {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
		return
	}

	// 继续搜索
	dfs5(node.Left, remain, path, result)
	dfs5(node.Right, remain, path, result)
}

// ========== 工具函数 ==========

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

// 打印路径列表
func printPaths(paths [][]int) {
	if len(paths) == 0 {
		fmt.Println("[]")
		return
	}

	fmt.Print("[")
	for i, path := range paths {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print("[")
		for j, val := range path {
			if j > 0 {
				fmt.Print(",")
			}
			fmt.Print(val)
		}
		fmt.Print("]")
	}
	fmt.Println("]")
}

// 比较两个路径列表是否相等
func equalPaths(paths1, paths2 [][]int) bool {
	if len(paths1) != len(paths2) {
		return false
	}

	// 转换为字符串集合进行比较（顺序无关）
	set1 := make(map[string]bool)
	set2 := make(map[string]bool)

	for _, path := range paths1 {
		set1[fmt.Sprintf("%v", path)] = true
	}

	for _, path := range paths2 {
		set2[fmt.Sprintf("%v", path)] = true
	}

	return reflect.DeepEqual(set1, set2)
}

// ========== 测试和性能评估 ==========
func main() {
	// 测试用例
	testCases := []struct {
		name      string
		treeVals  []interface{}
		targetSum int
		expected  [][]int
	}{
		{
			name:      "示例1: 复杂二叉树",
			treeVals:  []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, nil, 5, 1},
			targetSum: 22,
			expected:  [][]int{{5, 4, 11, 2}},
		},
		{
			name:      "示例2: 无解情况",
			treeVals:  []interface{}{1, 2, 3},
			targetSum: 5,
			expected:  [][]int{},
		},
		{
			name:      "示例3: 目标和为0",
			treeVals:  []interface{}{1, 2},
			targetSum: 0,
			expected:  [][]int{},
		},
		{
			name:      "测试4: 单节点树",
			treeVals:  []interface{}{5},
			targetSum: 5,
			expected:  [][]int{{5}},
		},
		{
			name:      "测试5: 单节点不匹配",
			treeVals:  []interface{}{5},
			targetSum: 3,
			expected:  [][]int{},
		},
		{
			name:      "测试6: 负数节点",
			treeVals:  []interface{}{1, -2, -3, 1, 3, -2, nil, -1},
			targetSum: 2,
			expected:  [][]int{{1, -2, 3}},
		},
		{
			name:      "测试7: 全负数",
			treeVals:  []interface{}{-1, -2, -3},
			targetSum: -3,
			expected:  [][]int{{-1, -2}},
		},
		{
			name:      "测试8: 链式结构",
			treeVals:  []interface{}{1, 2, nil, 3, nil, 4, nil},
			targetSum: 10,
			expected:  [][]int{{1, 2, 3, 4}},
		},
		{
			name:      "测试9: 多路径相同和",
			treeVals:  []interface{}{10, 5, 15, 3, 7, nil, 18},
			targetSum: 18,
			expected:  [][]int{{10, 5, 3}},
		},
		{
			name:      "测试10: 空树",
			treeVals:  []interface{}{},
			targetSum: 0,
			expected:  [][]int{},
		},
	}

	// 算法方法
	methods := []struct {
		name string
		fn   func(*TreeNode, int) [][]int
	}{
		{"DFS递归回溯(值传递)", pathSum1},
		{"DFS递归回溯(引用传递)", pathSum2},
		{"BFS层序遍历", pathSum3},
		{"DFS迭代栈", pathSum4},
		{"优化DFS", pathSum5},
	}

	fmt.Println("=== LeetCode 113. 路径总和 II - 测试结果 ===")
	fmt.Println()

	// 运行测试
	for _, tc := range testCases {
		fmt.Printf("测试用例: %s\n", tc.name)
		fmt.Printf("目标和: %d\n", tc.targetSum)

		// 构建测试树
		root := buildTree(tc.treeVals)
		fmt.Println("二叉树结构:")
		printTree(root)

		allPassed := true
		var results [][][]int
		var times []time.Duration

		for _, method := range methods {
			start := time.Now()
			result := method.fn(root, tc.targetSum)
			elapsed := time.Since(start)

			results = append(results, result)
			times = append(times, elapsed)

			status := "✅"
			if !equalPaths(result, tc.expected) {
				status = "❌"
				allPassed = false
			}

			fmt.Printf("  %s: %s (耗时: %v)\n", method.name, status, elapsed)
			fmt.Print("    结果: ")
			printPaths(result)
		}

		fmt.Print("期望结果: ")
		printPaths(tc.expected)

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
	fmt.Println("1. DFS递归回溯(值传递):")
	fmt.Println("   - 时间复杂度: O(N)")
	fmt.Println("   - 空间复杂度: O(H)")
	fmt.Println("   - 特点: 代码简洁，自动回溯")
	fmt.Println()
	fmt.Println("2. DFS递归回溯(引用传递):")
	fmt.Println("   - 时间复杂度: O(N)")
	fmt.Println("   - 空间复杂度: O(H)")
	fmt.Println("   - 特点: 空间效率高，需手动回溯")
	fmt.Println()
	fmt.Println("3. BFS层序遍历:")
	fmt.Println("   - 时间复杂度: O(N)")
	fmt.Println("   - 空间复杂度: O(W)")
	fmt.Println("   - 特点: 层次处理，适合宽树")
	fmt.Println()
	fmt.Println("4. DFS迭代栈:")
	fmt.Println("   - 时间复杂度: O(N)")
	fmt.Println("   - 空间复杂度: O(H)")
	fmt.Println("   - 特点: 显式栈，避免递归")
	fmt.Println()
	fmt.Println("5. 优化DFS:")
	fmt.Println("   - 时间复杂度: O(N)")
	fmt.Println("   - 空间复杂度: O(H)")
	fmt.Println("   - 特点: 预分配优化，减少内存分配")

	// 路径搜索演示
	fmt.Println("\n=== 路径搜索演示 ===")
	demoPathSearch()
}

// 性能测试
func performanceTest() {
	sizes := []int{100, 500, 1000, 2000}
	methods := []struct {
		name string
		fn   func(*TreeNode, int) [][]int
	}{
		{"DFS值传递", pathSum1},
		{"DFS引用传递", pathSum2},
		{"BFS层序", pathSum3},
		{"DFS迭代", pathSum4},
		{"优化DFS", pathSum5},
	}

	for _, size := range sizes {
		fmt.Printf("性能测试 - 树大小: %d个节点\n", size)

		// 构建测试树
		root := buildBalancedTree(size)
		targetSum := size + 100 // 不太可能匹配的目标和

		for _, method := range methods {
			start := time.Now()
			result := method.fn(root, targetSum)
			elapsed := time.Since(start)

			fmt.Printf("  %s: 找到%d条路径, 耗时=%v\n",
				method.name, len(result), elapsed)
		}
	}
}

// 构建平衡测试树
func buildBalancedTree(size int) *TreeNode {
	if size <= 0 {
		return nil
	}

	values := make([]interface{}, size)
	for i := 0; i < size; i++ {
		values[i] = i + 1
	}

	return buildTree(values)
}

// 路径搜索演示
func demoPathSearch() {
	fmt.Println("构建示例树:")
	treeVals := []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, nil, 5, 1}
	root := buildTree(treeVals)
	printTree(root)

	fmt.Println("寻找路径和为 22 的所有路径:")
	result := pathSum1(root, 22)

	fmt.Printf("找到 %d 条有效路径:\n", len(result))
	for i, path := range result {
		sum := 0
		for _, val := range path {
			sum += val
		}
		fmt.Printf("路径 %d: %v (和=%d)\n", i+1, path, sum)
	}

	fmt.Println("路径搜索完成!")
}
