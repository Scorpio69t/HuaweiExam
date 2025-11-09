package main

import (
	"fmt"
	"math"
	"strings"
)

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ==================== 方法一：递归（中间偏左）====================
// 时间复杂度：O(n)，每个元素访问一次
// 空间复杂度：O(log n)，递归栈深度
func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, left, right int) *TreeNode {
	// 递归终止条件
	if left > right {
		return nil
	}

	// 选择中间位置（偏左）作为根节点
	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}

	// 递归构建左右子树
	root.Left = buildBST(nums, left, mid-1)
	root.Right = buildBST(nums, mid+1, right)

	return root
}

// ==================== 方法二：递归（中间偏右）====================
// 选择中间偏右的元素作为根节点，会生成不同的平衡BST
func sortedArrayToBST2(nums []int) *TreeNode {
	return buildBST2(nums, 0, len(nums)-1)
}

func buildBST2(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// 选择中间位置（偏右）作为根节点
	mid := left + (right-left+1)/2
	root := &TreeNode{Val: nums[mid]}

	root.Left = buildBST2(nums, left, mid-1)
	root.Right = buildBST2(nums, mid+1, right)

	return root
}

// ==================== 方法三：迭代法（栈模拟）====================
// 使用栈模拟递归过程
type StackNode struct {
	left   int
	right  int
	parent *TreeNode
	isLeft bool
}

func sortedArrayToBST3(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var root *TreeNode
	stack := []StackNode{{0, len(nums) - 1, nil, true}}

	for len(stack) > 0 {
		// 弹出栈顶元素
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr.left > curr.right {
			continue
		}

		// 创建当前节点
		mid := curr.left + (curr.right-curr.left)/2
		node := &TreeNode{Val: nums[mid]}

		// 处理父节点连接
		if curr.parent == nil {
			root = node
		} else if curr.isLeft {
			curr.parent.Left = node
		} else {
			curr.parent.Right = node
		}

		// 先压入右子树，再压入左子树（栈是后进先出）
		stack = append(stack, StackNode{mid + 1, curr.right, node, false})
		stack = append(stack, StackNode{curr.left, mid - 1, node, true})
	}

	return root
}

// ==================== 方法四：中序遍历模拟 ====================
// 按中序遍历的顺序构建BST
func sortedArrayToBST4(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx := 0
	return inorderBuild(nums, &idx, 0, len(nums)-1)
}

func inorderBuild(nums []int, idx *int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2

	// 先构建左子树
	leftNode := inorderBuild(nums, idx, left, mid-1)

	// 创建根节点
	root := &TreeNode{Val: nums[*idx]}
	*idx++
	root.Left = leftNode

	// 再构建右子树
	root.Right = inorderBuild(nums, idx, mid+1, right)

	return root
}

// ==================== 辅助函数 ====================

// 中序遍历验证BST
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}

// 层序遍历（用于显示树结构）
func levelOrder(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	result := []interface{}{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	// 移除末尾的 nil
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

// 检查是否为平衡二叉树
func isBalanced(root *TreeNode) bool {
	return checkHeight(root) != -1
}

func checkHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := checkHeight(root.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := checkHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}

// 检查是否为BST
func isValidBST(root *TreeNode) bool {
	return validateBST(root, math.MinInt64, math.MaxInt64)
}

func validateBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return validateBST(root.Left, min, root.Val) && validateBST(root.Right, root.Val, max)
}

// 获取树的高度
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

// 树形打印
func printTree(root *TreeNode, prefix string, isLeft bool) {
	if root == nil {
		return
	}

	fmt.Print(prefix)
	if isLeft {
		fmt.Print("├── ")
	} else {
		fmt.Print("└── ")
	}
	fmt.Println(root.Val)

	if root.Left != nil || root.Right != nil {
		if root.Left != nil {
			printTree(root.Left, prefix+getTreePrefix(isLeft, true), true)
		} else {
			fmt.Println(prefix + getTreePrefix(isLeft, true) + "├── nil")
		}

		if root.Right != nil {
			printTree(root.Right, prefix+getTreePrefix(isLeft, false), false)
		} else {
			fmt.Println(prefix + getTreePrefix(isLeft, false) + "└── nil")
		}
	}
}

func getTreePrefix(isLeft, hasNext bool) string {
	if isLeft {
		return "│   "
	}
	return "    "
}

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

// ==================== 测试函数 ====================

func testCase(name string, nums []int) {
	fmt.Printf("\n========== %s ==========\n", name)
	fmt.Printf("输入: %v\n", nums)

	// 方法一：中间偏左
	root1 := sortedArrayToBST(nums)
	fmt.Println("\n方法一（中间偏左）:")
	fmt.Printf("层序遍历: %v\n", levelOrder(root1))
	fmt.Printf("中序遍历: %v\n", inorderTraversal(root1))
	fmt.Printf("是否平衡: %v\n", isBalanced(root1))
	fmt.Printf("是否BST: %v\n", isValidBST(root1))
	fmt.Printf("树高度: %d\n", getHeight(root1))
	fmt.Println("树结构:")
	printTree(root1, "", false)

	// 方法二：中间偏右
	root2 := sortedArrayToBST2(nums)
	fmt.Println("\n方法二（中间偏右）:")
	fmt.Printf("层序遍历: %v\n", levelOrder(root2))
	fmt.Printf("中序遍历: %v\n", inorderTraversal(root2))
	fmt.Printf("是否平衡: %v\n", isBalanced(root2))

	// 方法三：迭代法
	root3 := sortedArrayToBST3(nums)
	fmt.Println("\n方法三（迭代法）:")
	fmt.Printf("层序遍历: %v\n", levelOrder(root3))
	fmt.Printf("中序遍历: %v\n", inorderTraversal(root3))
	fmt.Printf("是否平衡: %v\n", isBalanced(root3))

	// 方法四：中序遍历
	root4 := sortedArrayToBST4(nums)
	fmt.Println("\n方法四（中序遍历）:")
	fmt.Printf("层序遍历: %v\n", levelOrder(root4))
	fmt.Printf("中序遍历: %v\n", inorderTraversal(root4))
	fmt.Printf("是否平衡: %v\n", isBalanced(root4))
}

// ==================== 扩展功能 ====================

// 从BST转回有序数组
func bstToSortedArray(root *TreeNode) []int {
	return inorderTraversal(root)
}

// 获取BST中第k小的元素
func kthSmallest(root *TreeNode, k int) int {
	arr := inorderTraversal(root)
	if k > 0 && k <= len(arr) {
		return arr[k-1]
	}
	return -1
}

// 将BST转换为更平衡的BST（重新构建）
func balanceBST(root *TreeNode) *TreeNode {
	arr := inorderTraversal(root)
	return sortedArrayToBST(arr)
}

// 比较两种方法生成的树结构差异
func compareTreeStructures(nums []int) {
	fmt.Printf("\n========== 比较不同方法生成的树结构 ==========\n")
	fmt.Printf("输入数组: %v\n", nums)

	root1 := sortedArrayToBST(nums)
	root2 := sortedArrayToBST2(nums)

	fmt.Println("\n中间偏左策略:")
	printTree(root1, "", false)

	fmt.Println("\n中间偏右策略:")
	printTree(root2, "", false)

	fmt.Printf("\n两棵树的中序遍历相同: %v\n",
		strings.Trim(strings.Join(strings.Fields(fmt.Sprint(inorderTraversal(root1))), ","), "[]") ==
			strings.Trim(strings.Join(strings.Fields(fmt.Sprint(inorderTraversal(root2))), ","), "[]"))
}

// 性能测试
func performanceTest() {
	fmt.Println("\n========== 性能测试 ==========")

	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		nums := make([]int, size)
		for i := 0; i < size; i++ {
			nums[i] = i
		}

		fmt.Printf("\n数组大小: %d\n", size)

		// 测试方法一
		root1 := sortedArrayToBST(nums)
		fmt.Printf("方法一 - 树高度: %d, 期望高度: %d\n",
			getHeight(root1), int(math.Ceil(math.Log2(float64(size+1)))))

		// 测试方法二
		root2 := sortedArrayToBST2(nums)
		fmt.Printf("方法二 - 树高度: %d, 期望高度: %d\n",
			getHeight(root2), int(math.Ceil(math.Log2(float64(size+1)))))
	}
}

func main() {
	// 测试用例1：示例1
	testCase("测试用例1：基本情况", []int{-10, -3, 0, 5, 9})

	// 测试用例2：示例2
	testCase("测试用例2：两个元素", []int{1, 3})

	// 测试用例3：单个元素
	testCase("测试用例3：单个元素", []int{1})

	// 测试用例4：奇数个元素
	testCase("测试用例4：奇数个元素", []int{1, 2, 3, 4, 5, 6, 7})

	// 测试用例5：偶数个元素
	testCase("测试用例5：偶数个元素", []int{1, 2, 3, 4, 5, 6})

	// 测试用例6：连续数字
	testCase("测试用例6：连续数字", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	// 测试用例7：负数
	testCase("测试用例7：负数", []int{-10, -5, -3, 0, 5, 10})

	// 扩展功能测试
	fmt.Println("\n========== 扩展功能测试 ==========")
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	root := sortedArrayToBST(nums)

	fmt.Printf("原数组: %v\n", nums)
	fmt.Printf("BST转回数组: %v\n", bstToSortedArray(root))
	fmt.Printf("第3小的元素: %d\n", kthSmallest(root, 3))

	// 比较不同构建策略
	compareTreeStructures([]int{1, 2, 3, 4, 5, 6, 7})

	// 性能测试
	performanceTest()
}
