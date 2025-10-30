package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =========================== 方法一：递归 ===========================
func inorderTraversal1(root *TreeNode) []int {
	var res []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// =========================== 方法二：显式栈迭代 ===========================
func inorderTraversal2(root *TreeNode) []int {
	var res []int
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n]
		res = append(res, node.Val)
		cur = node.Right
	}
	return res
}

// =========================== 方法三：颜色标记统一迭代 ===========================
// 白色：0 表示展开；黑色：1 表示访问
func inorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	type item struct {
		node *TreeNode
		clr  int
	}
	const white, black = 0, 1
	var res []int
	stack := []item{{root, white}}
	for len(stack) > 0 {
		n := len(stack) - 1
		i := stack[n]
		stack = stack[:n]
		if i.node == nil {
			continue
		}
		if i.clr == white {
			// 中序：右 黑 自 黑 左 白
			stack = append(stack, item{i.node.Right, white})
			stack = append(stack, item{i.node, black})
			stack = append(stack, item{i.node.Left, white})
		} else {
			res = append(res, i.node.Val)
		}
	}
	return res
}

// =========================== 方法四：Morris 遍历 ===========================
func inorderTraversal4(root *TreeNode) []int {
	var res []int
	cur := root
	for cur != nil {
		if cur.Left == nil {
			res = append(res, cur.Val)
			cur = cur.Right
			continue
		}
		// 寻找前驱（左子树最右）
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

// =========================== 构建/工具 ===========================
func arrayToTreeLevelOrder(arr []any) *TreeNode {
	// 基于层序队列的构建：按顺序为每个出队节点填充左/右孩子
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
		n := queue[0]
		queue = queue[1:]
		// 左孩子
		if i < len(arr) {
			if arr[i] != nil {
				left := &TreeNode{Val: arr[i].(int)}
				n.Left = left
				queue = append(queue, left)
			}
			i++
		}
		// 右孩子
		if i < len(arr) {
			if arr[i] != nil {
				right := &TreeNode{Val: arr[i].(int)}
				n.Right = right
				queue = append(queue, right)
			}
			i++
		}
	}
	return root
}

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

func buildRightChain(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	cur := root
	for i := 1; i < len(vals); i++ {
		cur.Right = &TreeNode{Val: vals[i]}
		cur = cur.Right
	}
	return root
}

// =========================== 测试 ===========================
func main() {
	fmt.Println("=== LeetCode 94: 二叉树的中序遍历 ===\n")

	testCases := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "例1: [1,null,2,3]",
			root:     arrayToTreeLevelOrder([]any{1, nil, 2, 3}),
			expected: []int{1, 3, 2},
		},
		{
			name:     "空树",
			root:     arrayToTreeLevelOrder([]any{}),
			expected: []int{},
		},
		{
			name:     "单节点",
			root:     arrayToTreeLevelOrder([]any{1}),
			expected: []int{1},
		},
		{
			name:     "完全二叉树",
			root:     arrayToTreeLevelOrder([]any{4, 2, 6, 1, 3, 5, 7}),
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "全左链",
			root:     arrayToTreeLevelOrder([]any{3, 2, nil, 1}),
			expected: []int{1, 2, 3},
		},
		{
			name:     "全右链",
			root:     buildRightChain([]int{1, 2, 3}),
			expected: []int{1, 2, 3},
		},
	}

	methods := map[string]func(*TreeNode) []int{
		"递归":          inorderTraversal1,
		"栈迭代":         inorderTraversal2,
		"颜色标记":        inorderTraversal3,
		"Morris O(1)": inorderTraversal4,
	}

	for name, f := range methods {
		fmt.Printf("方法：%s\n", name)
		pass := 0
		for i, tc := range testCases {
			got := f(tc.root)
			ok := equalSlice(got, tc.expected)
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
