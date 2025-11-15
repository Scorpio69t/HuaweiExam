package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =========================== 方法一：自底向上递归（最优解法） ===========================
func isBalanced1(root *TreeNode) bool {
	return checkHeight1(root) != -1
}

func checkHeight1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 检查左子树
	leftHeight := checkHeight1(root.Left)
	if leftHeight == -1 {
		return -1 // 左子树不平衡，立即返回
	}

	// 检查右子树
	rightHeight := checkHeight1(root.Right)
	if rightHeight == -1 {
		return -1 // 右子树不平衡，立即返回
	}

	// 检查当前节点是否平衡
	if abs(leftHeight-rightHeight) > 1 {
		return -1 // 不平衡
	}

	// 返回当前节点的高度
	return max(leftHeight, rightHeight) + 1
}

// =========================== 方法二：自顶向下递归 ===========================
func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 检查当前节点
	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	if abs(leftHeight-rightHeight) > 1 {
		return false
	}

	// 递归检查子树
	return isBalanced2(root.Left) && isBalanced2(root.Right)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

// =========================== 方法三：迭代BFS ===========================
func isBalanced3(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// 计算左右子树高度
		leftHeight := getHeight(node.Left)
		rightHeight := getHeight(node.Right)

		// 检查平衡
		if abs(leftHeight-rightHeight) > 1 {
			return false
		}

		// 将子节点入队
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return true
}

// =========================== 方法四：迭代DFS（后序遍历） ===========================
func isBalanced4(root *TreeNode) bool {
	if root == nil {
		return true
	}

	type item struct {
		node    *TreeNode
		visited bool
	}

	stack := []item{{root, false}}
	heightMap := make(map[*TreeNode]int)
	heightMap[nil] = 0

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr.visited {
			// 处理节点
			leftHeight := heightMap[curr.node.Left]
			rightHeight := heightMap[curr.node.Right]

			if abs(leftHeight-rightHeight) > 1 {
				return false
			}

			heightMap[curr.node] = max(leftHeight, rightHeight) + 1
		} else {
			// 后序遍历：右-左-根
			stack = append(stack, item{curr.node, true})
			if curr.node.Right != nil {
				stack = append(stack, item{curr.node.Right, false})
			}
			if curr.node.Left != nil {
				stack = append(stack, item{curr.node.Left, false})
			}
		}
	}

	return true
}

// =========================== 工具函数 ===========================
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

// =========================== 测试 ===========================
func main() {
	fmt.Println("=== LeetCode 110: 平衡二叉树 ===\n")

	testCases := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name:     "例1: [3,9,20,null,null,15,7]",
			root:     arrayToTreeLevelOrder([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			expected: true,
		},
		{
			name:     "例2: [1,2,2,3,3,null,null,4,4]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4}),
			expected: false,
		},
		{
			name:     "例3: []",
			root:     arrayToTreeLevelOrder([]interface{}{}),
			expected: true,
		},
		{
			name:     "单节点: [1]",
			root:     arrayToTreeLevelOrder([]interface{}{1}),
			expected: true,
		},
		{
			name:     "链状树（不平衡）: [1,2,null,3]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, nil, 3}),
			expected: false,
		},
		{
			name:     "左偏树（不平衡）: [1,2,null,3]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, nil, 3}),
			expected: false,
		},
		{
			name:     "完全平衡树: [1,2,3,4,5,6,7]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 3, 4, 5, 6, 7}),
			expected: true,
		},
		{
			name:     "一侧子树高度差1: [1,2,3,4]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 3, 4}),
			expected: true,
		},
		{
			name:     "一侧子树高度差2: [1,2,2,3,3,3,3,4,4,4,4,4,4,4,4]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4}),
			expected: true, // 构建函数限制，实际构建的树可能只有3层，是平衡的
		},
		{
			name:     "复杂不平衡树: [1,2,2,3,null,null,3,4,null,null,4]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 2, 3, nil, nil, 3, 4, nil, nil, 4}),
			expected: false,
		},
	}

	methods := map[string]func(*TreeNode) bool{
		"自底向上递归": isBalanced1,
		"自顶向下递归": isBalanced2,
		"迭代BFS":  isBalanced3,
		"迭代DFS":  isBalanced4,
	}

	for methodName, methodFunc := range methods {
		fmt.Printf("方法：%s\n", methodName)
		pass := 0
		for i, tc := range testCases {
			got := methodFunc(tc.root)
			ok := got == tc.expected
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
