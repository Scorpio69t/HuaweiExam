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

// navigationDevice 导航装置问题主函数
// 返回最少需要设置的导航装置数量（树的度量维数）
func navigationDevice(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 单节点
	if root.Left == nil && root.Right == nil {
		return 1
	}

	// 1) 构建无向邻接表与度
	adj := make(map[*TreeNode][]*TreeNode)
	deg := make(map[*TreeNode]int)
	var build func(node, parent *TreeNode)
	build = func(node, parent *TreeNode) {
		if node == nil {
			return
		}
		if _, ok := adj[node]; !ok {
			adj[node] = []*TreeNode{}
		}
		if parent != nil {
			adj[node] = append(adj[node], parent)
			adj[parent] = append(adj[parent], node)
		}
		build(node.Left, node)
		build(node.Right, node)
	}
	build(root, nil)
	for u, list := range adj {
		deg[u] = len(list)
	}

	// 2) 找出所有主结点(度>=3)与叶子(度==1)
	majors := []*TreeNode{}
	leaves := []*TreeNode{}
	for u, d := range deg {
		if d >= 3 {
			majors = append(majors, u)
		} else if d == 1 {
			leaves = append(leaves, u)
		}
	}
	// 若没有主结点，则是路径，答案为1
	if len(majors) == 0 {
		return 1
	}

	// 3) 统计每个主结点的终端叶子数：对每个叶子向内走直到遇到第一个主结点
	terminalCount := make(map[*TreeNode]int)
	for _, leaf := range leaves {
		cur := leaf
		prev := (*TreeNode)(nil)
		for {
			if deg[cur] >= 3 {
				terminalCount[cur]++
				break
			}
			// 继续向内走：唯一的未回退邻居
			next := (*TreeNode)(nil)
			for _, v := range adj[cur] {
				if v != prev {
					next = v
					break
				}
			}
			if next == nil {
				// 走到另一端也没遇到主结点（理论上不会，因为 len(majors)>0）
				break
			}
			prev, cur = cur, next
		}
	}

	// 4) 答案为 Σ(max(terminal(v)-1, 0))
	ans := 0
	for v, c := range terminalCount {
		_ = v
		if c > 1 {
			ans += c - 1
		}
	}
	return ans
}

// collectLeaves 收集所有叶子节点
func collectLeaves(root *TreeNode) []*TreeNode {
	var leaves []*TreeNode
	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 如果是叶子节点
		if node.Left == nil && node.Right == nil {
			leaves = append(leaves, node)
			return
		}

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return leaves
}

// findMinNavigationDevices 找到最少的导航装置数量
func findMinNavigationDevices(root *TreeNode, leaves []*TreeNode) int {
	n := len(leaves)

	// 尝试从1个导航装置开始
	for deviceCount := 1; deviceCount <= n; deviceCount++ {
		// 生成所有可能的导航装置组合
		combinations := generateCombinations(leaves, deviceCount)

		for _, combination := range combinations {
			if isValidNavigation(root, combination) {
				return deviceCount
			}
		}
	}

	return n // 最坏情况需要所有叶子节点
}

// generateCombinations 生成所有可能的导航装置组合
func generateCombinations(leaves []*TreeNode, k int) [][]*TreeNode {
	var result [][]*TreeNode
	var backtrack func(int, []*TreeNode)

	backtrack = func(start int, current []*TreeNode) {
		if len(current) == k {
			// 复制当前组合
			combination := make([]*TreeNode, k)
			copy(combination, current)
			result = append(result, combination)
			return
		}

		for i := start; i < len(leaves); i++ {
			current = append(current, leaves[i])
			backtrack(i+1, current)
			current = current[:len(current)-1]
		}
	}

	backtrack(0, []*TreeNode{})
	return result
}

// isValidNavigation 检查导航装置组合是否有效
func isValidNavigation(root *TreeNode, devices []*TreeNode) bool {
	// 计算每个节点到所有导航装置的距离
	nodeDistances := make(map[*TreeNode][]int)

	// 为每个节点计算到所有导航装置的距离
	var calculateDistances func(*TreeNode)
	calculateDistances = func(node *TreeNode) {
		if node == nil {
			return
		}

		distances := make([]int, len(devices))
		for i, device := range devices {
			distances[i] = calculateDistance(node, device)
		}

		nodeDistances[node] = distances

		calculateDistances(node.Left)
		calculateDistances(node.Right)
	}

	calculateDistances(root)

	// 检查是否有两个节点具有相同的距离向量
	seen := make(map[string]bool)
	for _, distances := range nodeDistances {
		// 将距离向量转换为字符串以便比较
		key := fmt.Sprintf("%v", distances)
		if seen[key] {
			return false // 有重复的距离向量
		}
		seen[key] = true
	}

	return true
}

// calculateDistance 计算两个节点之间的距离
func calculateDistance(node1, node2 *TreeNode) int {
	if node1 == nil || node2 == nil {
		return -1
	}

	// 找到最近公共祖先
	lca := findLCA(node1, node2)

	// 计算距离 = 到LCA的距离之和
	return getDepth(node1, lca) + getDepth(node2, lca)
}

// findLCA 找到两个节点的最近公共祖先
func findLCA(node1, node2 *TreeNode) *TreeNode {
	// 这里简化实现，假设node1和node2都在同一棵树中
	// 实际应用中需要更复杂的LCA算法

	// 简单实现：从根节点开始查找
	var find func(*TreeNode) *TreeNode
	find = func(root *TreeNode) *TreeNode {
		if root == nil || root == node1 || root == node2 {
			return root
		}

		left := find(root.Left)
		right := find(root.Right)

		if left != nil && right != nil {
			return root
		}

		if left != nil {
			return left
		}
		return right
	}

	// 需要找到包含这两个节点的树的根
	// 这里简化处理，假设从node1开始向上找
	return find(node1)
}

// getDepth 计算从node到target的深度
func getDepth(node, target *TreeNode) int {
	if node == nil || target == nil {
		return -1
	}

	if node == target {
		return 0
	}

	// 递归查找
	if node.Left != nil {
		leftDepth := getDepth(node.Left, target)
		if leftDepth != -1 {
			return leftDepth + 1
		}
	}

	if node.Right != nil {
		rightDepth := getDepth(node.Right, target)
		if rightDepth != -1 {
			return rightDepth + 1
		}
	}

	return -1
}

// 优化解法：基于图论的方法
func navigationDeviceOptimized(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 收集所有节点
	var nodes []*TreeNode
	var collectNodes func(*TreeNode)
	collectNodes = func(node *TreeNode) {
		if node == nil {
			return
		}
		nodes = append(nodes, node)
		collectNodes(node.Left)
		collectNodes(node.Right)
	}
	collectNodes(root)

	// 如果只有一个节点，需要1个导航装置
	if len(nodes) == 1 {
		return 1
	}

	// 计算所有节点之间的距离矩阵
	distances := make([][]int, len(nodes))
	for i := range distances {
		distances[i] = make([]int, len(nodes))
	}

	// 填充距离矩阵
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			dist := calculateDistanceOptimized(nodes[i], nodes[j])
			distances[i][j] = dist
			distances[j][i] = dist
		}
	}

	// 尝试不同的导航装置数量
	for deviceCount := 1; deviceCount <= len(nodes); deviceCount++ {
		if canDistinguishAllNodes(distances, deviceCount) {
			return deviceCount
		}
	}

	return len(nodes)
}

// calculateDistanceOptimized 优化的距离计算
func calculateDistanceOptimized(node1, node2 *TreeNode) int {
	if node1 == nil || node2 == nil {
		return -1
	}

	if node1 == node2 {
		return 0
	}

	// 使用LCA方法计算距离
	return calculateDistanceByLCA(node1, node2)
}

// calculateDistanceByLCA 使用LCA计算二叉树中两节点间的距离
func calculateDistanceByLCA(node1, node2 *TreeNode) int {
	if node1 == nil || node2 == nil {
		return -1
	}

	// 找到最近公共祖先
	lca := findLCASimple(node1, node2)
	if lca == nil {
		return -1
	}

	// 距离 = 到LCA的距离之和
	return getDepthFromLCA(lca, node1) + getDepthFromLCA(lca, node2)
}

// findLCASimple 找到两个节点的最近公共祖先
func findLCASimple(node1, node2 *TreeNode) *TreeNode {
	if node1 == nil || node2 == nil {
		return nil
	}

	// 从根节点开始查找
	var find func(*TreeNode) *TreeNode
	find = func(root *TreeNode) *TreeNode {
		if root == nil || root == node1 || root == node2 {
			return root
		}

		left := find(root.Left)
		right := find(root.Right)

		if left != nil && right != nil {
			return root
		}

		if left != nil {
			return left
		}
		return right
	}

	// 需要找到包含这两个节点的树的根
	// 这里简化处理，假设从node1开始向上找
	return find(node1)
}

// getDepthFromLCA 计算从LCA到target的深度
func getDepthFromLCA(lca, target *TreeNode) int {
	if lca == nil || target == nil {
		return -1
	}

	if lca == target {
		return 0
	}

	// 递归查找
	if lca.Left != nil {
		leftDepth := getDepthFromLCA(lca.Left, target)
		if leftDepth != -1 {
			return leftDepth + 1
		}
	}

	if lca.Right != nil {
		rightDepth := getDepthFromLCA(lca.Right, target)
		if rightDepth != -1 {
			return rightDepth + 1
		}
	}

	return -1
}

// canDistinguishAllNodes 检查是否可以用指定数量的导航装置区分所有节点
func canDistinguishAllNodes(distances [][]int, deviceCount int) bool {
	n := len(distances)

	// 生成所有可能的导航装置位置组合
	positions := make([]int, deviceCount)
	for i := range positions {
		positions[i] = i
	}

	// 检查所有组合
	for {
		// 检查当前组合是否能区分所有节点
		if canDistinguishWithPositions(distances, positions) {
			return true
		}

		// 生成下一个组合
		if !nextCombination(positions, n) {
			break
		}
	}

	return false
}

// canDistinguishWithPositions 检查指定位置的导航装置是否能区分所有节点
func canDistinguishWithPositions(distances [][]int, positions []int) bool {
	n := len(distances)

	// 为每个节点计算到所有导航装置的距离向量
	nodeSignatures := make(map[string]bool)

	for i := 0; i < n; i++ {
		signature := make([]int, len(positions))
		for j, pos := range positions {
			signature[j] = distances[i][pos]
		}

		// 将距离向量转换为字符串
		sigStr := fmt.Sprintf("%v", signature)
		if nodeSignatures[sigStr] {
			return false // 有重复的签名
		}
		nodeSignatures[sigStr] = true
	}

	return true
}

// nextCombination 生成下一个组合
func nextCombination(positions []int, n int) bool {
	k := len(positions)

	// 从右向左找到第一个可以增加的位置
	i := k - 1
	for i >= 0 && positions[i] == n-k+i {
		i--
	}

	if i < 0 {
		return false // 没有下一个组合
	}

	// 增加当前位置
	positions[i]++

	// 调整后面的位置
	for j := i + 1; j < k; j++ {
		positions[j] = positions[j-1] + 1
	}

	return true
}

// 辅助函数：从数组构建二叉树
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

// 辅助函数：打印二叉树（中序遍历）
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("nil ")
		return
	}

	fmt.Printf("%d ", root.Val)
	if root.Left != nil || root.Right != nil {
		fmt.Print("(")
		printTree(root.Left)
		printTree(root.Right)
		fmt.Print(")")
	}
}

func main() {
	fmt.Println("导航装置问题测试")
	fmt.Println("==================")

	// 测试用例1: [1,2,null,3,4] -> 2
	fmt.Println("\n测试用例1:")
	values1 := []interface{}{1, 2, nil, 3, 4}
	root1 := buildTree(values1)
	fmt.Printf("输入: ")
	printTree(root1)
	fmt.Println()

	result1 := navigationDevice(root1)
	fmt.Printf("输出: %d\n", result1)
	fmt.Printf("期望: 2\n")

	// 测试用例2: [1,2,3,4] -> 1
	fmt.Println("\n测试用例2:")
	values2 := []interface{}{1, 2, 3, 4}
	root2 := buildTree(values2)
	fmt.Printf("输入: ")
	printTree(root2)
	fmt.Println()

	result2 := navigationDevice(root2)
	fmt.Printf("输出: %d\n", result2)
	fmt.Printf("期望: 1\n")

	// 测试用例3: 单节点
	fmt.Println("\n测试用例3 - 单节点:")
	values3 := []interface{}{1}
	root3 := buildTree(values3)
	fmt.Printf("输入: ")
	printTree(root3)
	fmt.Println()

	result3 := navigationDevice(root3)
	fmt.Printf("输出: %d\n", result3)
	fmt.Printf("期望: 1\n")

	// 测试用例4: 线性结构
	fmt.Println("\n测试用例4 - 线性结构:")
	values4 := []interface{}{1, 2, nil, 3, nil, nil, nil, 4}
	root4 := buildTree(values4)
	fmt.Printf("输入: ")
	printTree(root4)
	fmt.Println()

	result4 := navigationDevice(root4)
	fmt.Printf("输出: %d\n", result4)
	fmt.Printf("期望: 2\n")

	// 测试优化版本
	fmt.Println("\n优化版本测试:")
	optResult1 := navigationDeviceOptimized(root1)
	optResult2 := navigationDeviceOptimized(root2)

	fmt.Printf("测试用例1优化版: %d\n", optResult1)
	fmt.Printf("测试用例2优化版: %d\n", optResult2)
}
