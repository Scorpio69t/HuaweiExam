package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =========================== 方法一：递归比较（最优解法） ===========================
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	// 两棵树都为空
	if p == nil && q == nil {
		return true
	}

	// 一棵树为空，另一棵不为空
	if p == nil || q == nil {
		return false
	}

	// 节点值不同
	if p.Val != q.Val {
		return false
	}

	// 递归比较左右子树
	return isSameTree1(p.Left, q.Left) && isSameTree1(p.Right, q.Right)
}

// =========================== 方法二：BFS迭代比较 ===========================
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	pQueue := []*TreeNode{p}
	qQueue := []*TreeNode{q}

	for len(pQueue) > 0 {
		// 同时出队
		pNode := pQueue[0]
		qNode := qQueue[0]
		pQueue = pQueue[1:]
		qQueue = qQueue[1:]

		// 检查nil
		if pNode == nil && qNode == nil {
			continue
		}
		if pNode == nil || qNode == nil {
			return false
		}
		if pNode.Val != qNode.Val {
			return false
		}

		// 同时入队左右子节点
		pQueue = append(pQueue, pNode.Left, pNode.Right)
		qQueue = append(qQueue, qNode.Left, qNode.Right)
	}

	return true
}

// =========================== 方法三：DFS迭代比较 ===========================
func isSameTree3(p *TreeNode, q *TreeNode) bool {
	pStack := []*TreeNode{p}
	qStack := []*TreeNode{q}

	for len(pStack) > 0 {
		// 同时出栈
		n1 := len(pStack) - 1
		pNode := pStack[n1]
		qNode := qStack[n1]
		pStack = pStack[:n1]
		qStack = qStack[:n1]

		// 检查nil
		if pNode == nil && qNode == nil {
			continue
		}
		if pNode == nil || qNode == nil {
			return false
		}
		if pNode.Val != qNode.Val {
			return false
		}

		// 同时入栈左右子节点（注意顺序：先右后左，因为栈是后进先出）
		pStack = append(pStack, pNode.Right, pNode.Left)
		qStack = append(qStack, qNode.Right, qNode.Left)
	}

	return true
}

// =========================== 方法四：序列化比较 ===========================
func isSameTree4(p *TreeNode, q *TreeNode) bool {
	return serialize4(p) == serialize4(q)
}

func serialize4(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return fmt.Sprintf("%d,%s,%s",
		root.Val,
		serialize4(root.Left),
		serialize4(root.Right))
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
		if i < len(arr) && arr[i] != nil {
			left := &TreeNode{Val: arr[i].(int)}
			node.Left = left
			queue = append(queue, left)
		}
		i++

		// 右子节点
		if i < len(arr) && arr[i] != nil {
			right := &TreeNode{Val: arr[i].(int)}
			node.Right = right
			queue = append(queue, right)
		}
		i++
	}

	return root
}

// =========================== 测试 ===========================
func main() {
	fmt.Println("=== LeetCode 100: 相同的树 ===\n")

	testCases := []struct {
		name     string
		p        *TreeNode
		q        *TreeNode
		expected bool
	}{
		{
			name:     "例1: [1,2,3] vs [1,2,3] - 完全相同",
			p:        arrayToTreeLevelOrder([]interface{}{1, 2, 3}),
			q:        arrayToTreeLevelOrder([]interface{}{1, 2, 3}),
			expected: true,
		},
		{
			name:     "例2: [1,2] vs [1,null,2] - 结构不同",
			p:        arrayToTreeLevelOrder([]interface{}{1, 2}),
			q:        arrayToTreeLevelOrder([]interface{}{1, nil, 2}),
			expected: false,
		},
		{
			name:     "例3: [1,2,1] vs [1,1,2] - 值不同",
			p:        arrayToTreeLevelOrder([]interface{}{1, 2, 1}),
			q:        arrayToTreeLevelOrder([]interface{}{1, 1, 2}),
			expected: false,
		},
		{
			name:     "都为空 - 相同",
			p:        arrayToTreeLevelOrder([]interface{}{}),
			q:        arrayToTreeLevelOrder([]interface{}{}),
			expected: true,
		},
		{
			name:     "单节点相同: [1] vs [1]",
			p:        arrayToTreeLevelOrder([]interface{}{1}),
			q:        arrayToTreeLevelOrder([]interface{}{1}),
			expected: true,
		},
		{
			name:     "单节点不同: [1] vs [2]",
			p:        arrayToTreeLevelOrder([]interface{}{1}),
			q:        arrayToTreeLevelOrder([]interface{}{2}),
			expected: false,
		},
		{
			name:     "一棵为空: [1] vs []",
			p:        arrayToTreeLevelOrder([]interface{}{1}),
			q:        arrayToTreeLevelOrder([]interface{}{}),
			expected: false,
		},
		{
			name:     "完全二叉树相同: [4,2,6,1,3,5,7] vs [4,2,6,1,3,5,7]",
			p:        arrayToTreeLevelOrder([]interface{}{4, 2, 6, 1, 3, 5, 7}),
			q:        arrayToTreeLevelOrder([]interface{}{4, 2, 6, 1, 3, 5, 7}),
			expected: true,
		},
		{
			name:     "链状树相同: [1,null,2,null,3] vs [1,null,2,null,3]",
			p:        arrayToTreeLevelOrder([]interface{}{1, nil, 2, nil, nil, nil, 3}),
			q:        arrayToTreeLevelOrder([]interface{}{1, nil, 2, nil, nil, nil, 3}),
			expected: true,
		},
		{
			name:     "结构细微不同: [1,2,3] vs [1,2,3,null]",
			p:        arrayToTreeLevelOrder([]interface{}{1, 2, 3}),
			q:        arrayToTreeLevelOrder([]interface{}{1, 2, 3, nil}),
			expected: true, // 层序遍历中，末尾的nil不影响树结构
		},
	}

	methods := map[string]func(*TreeNode, *TreeNode) bool{
		"递归比较":   isSameTree1,
		"BFS迭代":  isSameTree2,
		"DFS迭代":  isSameTree3,
		"序列化比较": isSameTree4,
	}

	for methodName, methodFunc := range methods {
		fmt.Printf("方法：%s\n", methodName)
		pass := 0
		for i, tc := range testCases {
			got := methodFunc(tc.p, tc.q)
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