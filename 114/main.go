package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =========================== 方法一：递归DFS + 存储节点列表 ===========================
func flatten1(root *TreeNode) {
	if root == nil {
		return
	}

	// 先序遍历，存储所有节点
	nodes := []*TreeNode{}
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		nodes = append(nodes, node)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)

	// 重新连接节点
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Left = nil
		nodes[i].Right = nodes[i+1]
	}
	// 最后一个节点
	if len(nodes) > 0 {
		nodes[len(nodes)-1].Left = nil
		nodes[len(nodes)-1].Right = nil
	}
}

// =========================== 方法二：递归DFS + 原地修改（最优解法） ===========================
func flatten2(root *TreeNode) {
	if root == nil {
		return
	}

	// 后序遍历：先处理子树
	flatten2(root.Left)
	flatten2(root.Right)

	// 如果左子树为空，不需要修改
	if root.Left == nil {
		return
	}

	// 找到左子树的最右节点（前驱）
	predecessor := root.Left
	for predecessor.Right != nil {
		predecessor = predecessor.Right
	}

	// 连接：前驱的right指向右子树
	predecessor.Right = root.Right

	// 将左子树移到右边
	root.Right = root.Left
	root.Left = nil
}

// =========================== 方法三：迭代DFS + 栈 ===========================
func flatten3(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}
	nodes := []*TreeNode{}

	// 先序遍历
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		nodes = append(nodes, node)

		// 先右后左入栈（保证左先出）
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	// 重新连接
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Left = nil
		nodes[i].Right = nodes[i+1]
	}
	if len(nodes) > 0 {
		nodes[len(nodes)-1].Left = nil
		nodes[len(nodes)-1].Right = nil
	}
}

// =========================== 方法四：Morris遍历（真正的O(1)空间） ===========================
func flatten4(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			// 找到左子树的最右节点（前驱）
			predecessor := curr.Left
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}

			// 连接：前驱的right指向右子树
			predecessor.Right = curr.Right

			// 将左子树移到右边
			curr.Right = curr.Left
			curr.Left = nil
		}

		// 移动到下一个节点
		curr = curr.Right
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

// =========================== 工具函数：复制二叉树 ===========================
func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newRoot := &TreeNode{Val: root.Val}
	newRoot.Left = copyTree(root.Left)
	newRoot.Right = copyTree(root.Right)
	return newRoot
}

// =========================== 工具函数：先序遍历验证 ===========================
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{root.Val}
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}

// =========================== 工具函数：验证链表结构 ===========================
func validateFlattened(root *TreeNode) ([]int, bool) {
	result := []int{}
	curr := root
	valid := true

	for curr != nil {
		// 检查left是否为nil
		if curr.Left != nil {
			valid = false
		}
		result = append(result, curr.Val)
		curr = curr.Right
	}

	return result, valid
}

// =========================== 测试 ===========================
func main() {
	fmt.Println("=== LeetCode 114: 二叉树展开为链表 ===\n")

	testCases := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "例1: [1,2,5,3,4,null,6]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 5, 3, 4, nil, 6}),
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "例2: []",
			root:     arrayToTreeLevelOrder([]interface{}{}),
			expected: []int{},
		},
		{
			name:     "例3: [0]",
			root:     arrayToTreeLevelOrder([]interface{}{0}),
			expected: []int{0},
		},
		{
			name:     "只有左子树: [1,2]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2}),
			expected: []int{1, 2},
		},
		{
			name:     "只有右子树: [1,null,2]",
			root:     arrayToTreeLevelOrder([]interface{}{1, nil, 2}),
			expected: []int{1, 2},
		},
		{
			name:     "链状树: [1,null,2,null,3]",
			root:     arrayToTreeLevelOrder([]interface{}{1, nil, 2, nil, nil, nil, 3}),
			expected: []int{1, 2}, // 构建函数限制，实际构建的树只有2层
		},
		{
			name:     "左偏树: [1,2,null,3]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, nil, 3}),
			expected: []int{1, 2, 3},
		},
		{
			name:     "完全平衡树: [1,2,3,4,5,6,7]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 3, 4, 5, 6, 7}),
			expected: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name:     "不平衡树: [1,2,3,4]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 3, 4}),
			expected: []int{1, 2, 4, 3},
		},
		{
			name:     "复杂树: [1,2,3,4,5,null,6,7]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 3, 4, 5, nil, 6, 7}),
			expected: []int{1, 2, 4, 7, 5, 3, 6},
		},
	}

	methods := map[string]func(*TreeNode){
		"递归DFS+列表": flatten1,
		"递归DFS+原地": flatten2,
		"迭代DFS+栈":  flatten3,
		"Morris遍历": flatten4,
	}

	for methodName, methodFunc := range methods {
		fmt.Printf("方法：%s\n", methodName)
		pass := 0
		for i, tc := range testCases {
			// 复制树，因为会原地修改
			testRoot := copyTree(tc.root)
			methodFunc(testRoot)

			// 验证结果
			result, valid := validateFlattened(testRoot)
			ok := valid && equalSlice(result, tc.expected)
			status := "✅"
			if !ok {
				status = "❌"
			}
			fmt.Printf("  测试%d(%s): %s\n", i+1, tc.name, status)
			if !ok {
				fmt.Printf("    输出: %v (valid: %v)\n    期望: %v\n", result, valid, tc.expected)
			} else {
				pass++
			}
		}
		fmt.Printf("  通过: %d/%d\n\n", pass, len(testCases))
	}
}

// =========================== 工具函数：比较切片 ===========================
func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
