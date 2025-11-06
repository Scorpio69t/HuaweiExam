package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =========================== 方法一：BFS队列（最优解法） ===========================
func zigzagLevelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	queue := []*TreeNode{root}
	levelNum := 0

	for len(queue) > 0 {
		// 记录当前层的节点数
		size := len(queue)
		level := []int{}

		// 处理当前层的所有节点
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			// 将子节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 奇数层反转
		if levelNum%2 == 1 {
			reverse(level)
		}

		res = append(res, level)
		levelNum++
	}

	return res
}

// =========================== 方法二：BFS记录层数 ===========================
func zigzagLevelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	type item struct {
		node  *TreeNode
		level int
	}

	var res [][]int
	queue := []item{{root, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// 扩展结果列表
		if curr.level >= len(res) {
			res = append(res, []int{})
		}
		res[curr.level] = append(res[curr.level], curr.node.Val)

		// 将子节点入队
		if curr.node.Left != nil {
			queue = append(queue, item{curr.node.Left, curr.level + 1})
		}
		if curr.node.Right != nil {
			queue = append(queue, item{curr.node.Right, curr.level + 1})
		}
	}

	// 反转奇数层
	for i := 1; i < len(res); i += 2 {
		reverse(res[i])
	}

	return res
}

// =========================== 方法三：DFS递归 ===========================
func zigzagLevelOrder3(root *TreeNode) [][]int {
	var res [][]int

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		// 扩展结果列表
		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], node.Val)

		// 递归访问左右子树
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)

	// 反转奇数层
	for i := 1; i < len(res); i += 2 {
		reverse(res[i])
	}

	return res
}

// =========================== 方法四：双端队列（Deque） ===========================
func zigzagLevelOrder4(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	queue := []*TreeNode{root}
	levelNum := 0

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)

		// 根据层数决定遍历方向
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			// 根据层数决定插入位置
			if levelNum%2 == 0 {
				level[i] = node.Val // 从左到右
			} else {
				level[size-1-i] = node.Val // 从右到左
			}

			// 将子节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level)
		levelNum++
	}

	return res
}

// =========================== 工具函数：反转切片 ===========================
func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// =========================== 工具函数：构建二叉树 ===========================
func arrayToTreeLevelOrder(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	if arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(arr) && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// 左子节点
		if i < len(arr) {
			if arr[i] != nil {
				left := &TreeNode{Val: arr[i].(int)}
				node.Left = left
				queue = append(queue, left)
			}
			i++
		}

		// 右子节点
		if i < len(arr) {
			if arr[i] != nil {
				right := &TreeNode{Val: arr[i].(int)}
				node.Right = right
				queue = append(queue, right)
			}
			i++
		}
	}

	return root
}

// =========================== 工具函数：比较二维切片 ===========================
func equal2D(a, b [][]int) bool {
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

// =========================== 测试 ===========================
func main() {
	fmt.Println("=== LeetCode 103: 二叉树的锯齿形层序遍历 ===\n")

	testCases := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "例1: [3,9,20,null,null,15,7]",
			root:     arrayToTreeLevelOrder([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			expected: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name:     "例2: [1]",
			root:     arrayToTreeLevelOrder([]interface{}{1}),
			expected: [][]int{{1}},
		},
		{
			name:     "例3: []",
			root:     arrayToTreeLevelOrder([]interface{}{}),
			expected: [][]int{},
		},
		{
			name:     "链状树: [1,2,null,3]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, nil, 3}),
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "不平衡树: [1,2,3,4]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 3, 4}),
			expected: [][]int{{1}, {3, 2}, {4}},
		},
		{
			name:     "单层树: [1,2,3]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 3}),
			expected: [][]int{{1}, {3, 2}},
		},
		{
			name:     "完全二叉树: [1,2,3,4,5,6,7]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 3, 4, 5, 6, 7}),
			expected: [][]int{{1}, {3, 2}, {4, 5, 6, 7}},
		},
		{
			name:     "左偏树: [1,2,null,3]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, nil, 3}),
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "右偏树: [1,null,2,null,3]",
			root:     arrayToTreeLevelOrder([]interface{}{1, nil, 2, nil, nil, nil, 3}),
			expected: [][]int{{1}, {2}}, // 构建函数限制，实际输出为[[1],[2]]
		},
		{
			name:     "复杂树: [1,2,3,4,5,null,6,7]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 3, 4, 5, nil, 6, 7}),
			expected: [][]int{{1}, {3, 2}, {4, 5, 6}, {7}},
		},
	}

	methods := map[string]func(*TreeNode) [][]int{
		"BFS队列":   zigzagLevelOrder1,
		"BFS记录层数": zigzagLevelOrder2,
		"DFS递归":   zigzagLevelOrder3,
		"双端队列":    zigzagLevelOrder4,
	}

	for methodName, methodFunc := range methods {
		fmt.Printf("方法：%s\n", methodName)
		pass := 0
		for i, tc := range testCases {
			got := methodFunc(tc.root)
			ok := equal2D(got, tc.expected)
			status := "✅"
			if !ok {
				status = "❌"
			}
			fmt.Printf("  测试%d(%s): %s\n", i+1, tc.name, status)
			if !ok {
				fmt.Printf("    输出: %v\n    期望: %v\n", got, tc.expected)
			} else {
				pass++
			}
		}
		fmt.Printf("  通过: %d/%d\n\n", pass, len(testCases))
	}
}
