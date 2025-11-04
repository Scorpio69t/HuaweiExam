package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =========================== 方法一：递归上下界验证（最优解法） ===========================
func isValidBST1(root *TreeNode) bool {
	return isValidBSTHelper1(root, nil, nil)
}

func isValidBSTHelper1(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	// 检查当前节点值是否在范围内（严格小于/大于）
	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}

	// 递归验证左右子树
	// 左子树：上界更新为当前节点值
	// 右子树：下界更新为当前节点值
	return isValidBSTHelper1(node.Left, min, &node.Val) &&
		isValidBSTHelper1(node.Right, &node.Val, max)
}

// =========================== 方法二：中序遍历递归验证 ===========================
var prev2 *int

func isValidBST2(root *TreeNode) bool {
	prev2 = nil
	return inorder2(root)
}

func inorder2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 验证左子树
	if !inorder2(root.Left) {
		return false
	}

	// 检查当前节点（中序遍历，此时访问根节点）
	if prev2 != nil && root.Val <= *prev2 {
		return false
	}
	prev2 = &root.Val

	// 验证右子树
	return inorder2(root.Right)
}

// =========================== 方法三：中序遍历迭代验证 ===========================
func isValidBST3(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := []*TreeNode{}
	cur := root
	var prev *int

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

		// 检查递增性（BST中序遍历必须严格递增）
		if prev != nil && node.Val <= *prev {
			return false
		}
		prev = &node.Val

		// 转向右子树
		cur = node.Right
	}

	return true
}

// =========================== 方法四：Morris遍历验证 ===========================
func isValidBST4(root *TreeNode) bool {
	if root == nil {
		return true
	}

	cur := root
	var prev *int

	for cur != nil {
		if cur.Left == nil {
			// 访问当前节点
			if prev != nil && cur.Val <= *prev {
				return false
			}
			prev = &cur.Val
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
				if prev != nil && cur.Val <= *prev {
					return false
				}
				prev = &cur.Val
				cur = cur.Right
			}
		}
	}

	return true
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
	fmt.Println("=== LeetCode 98: 验证二叉搜索树 ===\n")

	testCases := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name:     "例1: [2,1,3] - 有效BST",
			root:     arrayToTreeLevelOrder([]interface{}{2, 1, 3}),
			expected: true,
		},
		{
			name:     "例2: [5,1,4,null,null,3,6] - 无效BST",
			root:     arrayToTreeLevelOrder([]interface{}{5, 1, 4, nil, nil, 3, 6}),
			expected: false,
		},
		{
			name:     "单节点 - 有效BST",
			root:     arrayToTreeLevelOrder([]interface{}{1}),
			expected: true,
		},
		{
			name:     "[1,1] - 重复值，无效BST",
			root:     arrayToTreeLevelOrder([]interface{}{1, 1}),
			expected: false,
		},
		{
			name:     "[2,2,3] - 左子节点重复，无效BST",
			root:     arrayToTreeLevelOrder([]interface{}{2, 2, 3}),
			expected: false,
		},
		{
			name:     "[5,6,7] - 左子节点大于根，无效BST",
			root:     arrayToTreeLevelOrder([]interface{}{5, 6, 7}),
			expected: false,
		},
		{
			name:     "[4,2,6,1,3,5,7] - 完全BST，有效",
			root:     arrayToTreeLevelOrder([]interface{}{4, 2, 6, 1, 3, 5, 7}),
			expected: true,
		},
		{
			name:     "[1,null,2,null,null,null,3] - 右链BST，有效",
			root:     arrayToTreeLevelOrder([]interface{}{1, nil, 2, nil, nil, nil, 3}),
			expected: true,
		},
		{
			name:     "[10,5,15,null,null,6,20] - 右子树中有小于根的值，无效",
			root:     arrayToTreeLevelOrder([]interface{}{10, 5, 15, nil, nil, 6, 20}),
			expected: false,
		},
		{
			name:     "[3,1,5,0,2,4,6] - 复杂有效BST",
			root:     arrayToTreeLevelOrder([]interface{}{3, 1, 5, 0, 2, 4, 6}),
			expected: true,
		},
	}

	methods := map[string]func(*TreeNode) bool{
		"递归上下界验证":  isValidBST1,
		"中序遍历递归":   isValidBST2,
		"中序遍历迭代":   isValidBST3,
		"Morris遍历": isValidBST4,
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
