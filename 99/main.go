package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =========================== 方法一：中序遍历递归 + 错误节点记录（推荐解法） ===========================
var prev1, first1, second1 *TreeNode

func recoverTree1(root *TreeNode) {
	prev1, first1, second1 = nil, nil, nil
	inorder1(root)
	// 交换两个错误节点的值
	if first1 != nil && second1 != nil {
		first1.Val, second1.Val = second1.Val, first1.Val
	}
}

func inorder1(root *TreeNode) {
	if root == nil {
		return
	}

	inorder1(root.Left)

	// 检查逆序
	if prev1 != nil && prev1.Val > root.Val {
		if first1 == nil {
			// 第一个逆序对
			first1 = prev1
		}
		// 第二个逆序对（可能不存在，如果相邻交换）
		second1 = root
	}
	prev1 = root

	inorder1(root.Right)
}

// =========================== 方法二：中序遍历迭代 ===========================
func recoverTree2(root *TreeNode) {
	var prev, first, second *TreeNode
	stack := []*TreeNode{}
	cur := root

	for cur != nil || len(stack) > 0 {
		// 一路向左
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		// 访问节点
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n]

		// 检查逆序
		if prev != nil && prev.Val > node.Val {
			if first == nil {
				first = prev
			}
			second = node
		}
		prev = node

		// 转向右子树
		cur = node.Right
	}

	// 交换
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

// =========================== 方法三：中序遍历优化（闭包） ===========================
func recoverTree3(root *TreeNode) {
	var prev, first, second *TreeNode

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)

		if prev != nil && prev.Val > node.Val {
			if first == nil {
				first = prev
			}
			second = node
		}
		prev = node

		dfs(node.Right)
	}

	dfs(root)
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

// =========================== 方法四：Morris遍历（O(1)空间，最优解法） ===========================
func recoverTree4(root *TreeNode) {
	var prev, first, second *TreeNode
	cur := root

	for cur != nil {
		if cur.Left == nil {
			// 访问当前节点
			if prev != nil && prev.Val > cur.Val {
				if first == nil {
					first = prev
				}
				second = cur
			}
			prev = cur
			cur = cur.Right
		} else {
			// 寻找前驱节点（左子树的最右节点）
			pre := cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}

			if pre.Right == nil {
				// 建立线索
				pre.Right = cur
				cur = cur.Left
			} else {
				// 拆除线索并访问当前节点
				pre.Right = nil
				if prev != nil && prev.Val > cur.Val {
					if first == nil {
						first = prev
					}
					second = cur
				}
				prev = cur
				cur = cur.Right
			}
		}
	}

	// 交换
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
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

// =========================== 工具函数：中序遍历验证（用于测试） ===========================
func inorderTraversal(root *TreeNode) []int {
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

// =========================== 工具函数：验证BST ===========================
func isValidBST(root *TreeNode) bool {
	var prev *int
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Left) {
			return false
		}
		if prev != nil && node.Val <= *prev {
			return false
		}
		prev = &node.Val
		return dfs(node.Right)
	}
	return dfs(root)
}

// =========================== 工具函数：复制树（用于测试） ===========================
func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newRoot := &TreeNode{Val: root.Val}
	newRoot.Left = copyTree(root.Left)
	newRoot.Right = copyTree(root.Right)
	return newRoot
}

// =========================== 测试 ===========================
func main() {
	fmt.Println("=== LeetCode 99: 恢复二叉搜索树 ===\n")

	testCases := []struct {
		name         string
		root         *TreeNode
		expected     []int // 恢复后的中序遍历结果
		expectedSwap []int // 被交换的两个值（用于验证）
	}{
		{
			name:         "例1: [1,3,null,null,2] - 相邻交换",
			root:         arrayToTreeLevelOrder([]interface{}{1, 3, nil, nil, 2}),
			expected:     []int{1, 2, 3},
			expectedSwap: []int{1, 3},
		},
		{
			name:         "例2: [3,1,4,null,null,2] - 非相邻交换",
			root:         arrayToTreeLevelOrder([]interface{}{3, 1, 4, nil, nil, 2}),
			expected:     []int{1, 2, 3, 4},
			expectedSwap: []int{3, 2},
		},
		{
			name:         "两个节点: [2,1]",
			root:         arrayToTreeLevelOrder([]interface{}{2, 1}),
			expected:     []int{1, 2},
			expectedSwap: []int{1, 2},
		},
		{
			name:         "根节点交换: [2,1,3]",
			root:         arrayToTreeLevelOrder([]interface{}{2, 1, 3}),
			expected:     []int{1, 2, 3},
			expectedSwap: []int{1, 2},
		},
		{
			name:         "复杂交换: [5,3,9,1,4,7,10,null,null,2]",
			root:         arrayToTreeLevelOrder([]interface{}{5, 3, 9, 1, 4, 7, 10, nil, nil, 2}),
			expected:     []int{1, 2, 3, 4, 5, 7, 9, 10},
			expectedSwap: []int{1, 5}, // 交换1和5
		},
		{
			name:         "链状交换: [3,2,null,1]",
			root:         arrayToTreeLevelOrder([]interface{}{3, 2, nil, 1}),
			expected:     []int{1, 2, 3},
			expectedSwap: []int{1, 3},
		},
		{
			name:         "完全BST交换: [4,2,6,1,3,5,7] -> [4,2,6,1,5,3,7]",
			root:         arrayToTreeLevelOrder([]interface{}{4, 2, 6, 1, 5, 3, 7}), // 交换3和5
			expected:     []int{1, 2, 3, 4, 5, 6, 7},
			expectedSwap: []int{3, 5},
		},
	}

	methods := map[string]func(*TreeNode){
		"中序遍历递归": recoverTree1,
		"中序遍历迭代": recoverTree2,
		"中序遍历优化": recoverTree3,
		"Morris遍历":   recoverTree4,
	}

	for methodName, methodFunc := range methods {
		fmt.Printf("方法：%s\n", methodName)
		pass := 0
		for i, tc := range testCases {
			// 复制树以避免修改原始测试用例
			testRoot := copyTree(tc.root)
			methodFunc(testRoot)

			// 验证恢复后的BST
			got := inorderTraversal(testRoot)
			isValid := isValidBST(testRoot)
			expectedMatch := true
			if len(got) != len(tc.expected) {
				expectedMatch = false
			} else {
				for j := range got {
					if got[j] != tc.expected[j] {
						expectedMatch = false
						break
					}
				}
			}

			ok := expectedMatch && isValid
			status := "✅"
			if !ok {
				status = "❌"
			}
			fmt.Printf("  测试%d(%s): %s\n", i+1, tc.name, status)
			if !ok {
				fmt.Printf("    输出中序遍历: %v\n    期望中序遍历: %v\n    是否有效BST: %v\n", got, tc.expected, isValid)
			} else {
				pass++
			}
		}
		fmt.Printf("  通过: %d/%d\n\n", pass, len(testCases))
	}
}