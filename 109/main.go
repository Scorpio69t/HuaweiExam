package main

import (
	"fmt"
	"math"
)

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ==================== 方法一：快慢指针 + 递归 ====================
// 时间复杂度：O(n log n)，每层递归O(n)找中点，深度O(log n)
// 空间复杂度：O(log n)，递归栈深度
func sortedListToBST(head *ListNode) *TreeNode {
	// 边界条件
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	// 使用快慢指针找中点
	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 断开链表：prev.Next = nil
	if prev != nil {
		prev.Next = nil
	}

	// 创建根节点（中点）
	root := &TreeNode{Val: slow.Val}

	// 递归构建左右子树
	root.Left = sortedListToBST(head)       // 左半链表
	root.Right = sortedListToBST(slow.Next) // 右半链表

	return root
}

// ==================== 方法二：转换为数组 ====================
// 时间复杂度：O(n)，遍历一次链表 + 构建树
// 空间复杂度：O(n)，数组存储
func sortedListToBST2(head *ListNode) *TreeNode {
	// 链表转数组
	nums := []int{}
	curr := head
	for curr != nil {
		nums = append(nums, curr.Val)
		curr = curr.Next
	}

	// 使用108题的方法构建BST
	return arrayToBST(nums, 0, len(nums)-1)
}

func arrayToBST(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}

	root.Left = arrayToBST(nums, left, mid-1)
	root.Right = arrayToBST(nums, mid+1, right)

	return root
}

// ==================== 方法三：中序遍历模拟（最优解）====================
// 时间复杂度：O(n)，每个节点访问一次
// 空间复杂度：O(log n)，递归栈
func sortedListToBST3(head *ListNode) *TreeNode {
	// 计算链表长度
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}

	// 使用全局指针，按中序遍历顺序消费链表节点
	return inorderBuild(&head, 0, length-1)
}

func inorderBuild(head **ListNode, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2

	// 先构建左子树（中序遍历：左-根-右）
	leftTree := inorderBuild(head, left, mid-1)

	// 创建根节点，消费当前链表节点
	root := &TreeNode{Val: (*head).Val}
	*head = (*head).Next // 指针后移

	// 再构建右子树
	root.Left = leftTree
	root.Right = inorderBuild(head, mid+1, right)

	return root
}

// ==================== 方法四：递归 + 计算长度优化 ====================
// 时间复杂度：O(n)
// 空间复杂度：O(log n)
func sortedListToBST4(head *ListNode) *TreeNode {
	// 计算链表长度
	length := getLength(head)
	return buildWithLength(head, length)
}

func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

func buildWithLength(head *ListNode, length int) *TreeNode {
	if length == 0 {
		return nil
	}
	if length == 1 {
		return &TreeNode{Val: head.Val}
	}

	// 找到中点位置
	mid := length / 2

	// 移动到中点
	curr := head
	for i := 0; i < mid; i++ {
		curr = curr.Next
	}

	// 创建根节点
	root := &TreeNode{Val: curr.Val}

	// 递归构建左右子树
	root.Left = buildWithLength(head, mid)
	root.Right = buildWithLength(curr.Next, length-mid-1)

	return root
}

// ==================== 辅助函数 ====================

// 创建链表
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	curr := head

	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
	}

	return head
}

// 打印链表
func printList(head *ListNode) {
	fmt.Print("[")
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Print("]")
}

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

// 层序遍历
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
	fmt.Print("输入链表: ")
	head := createList(nums)
	printList(head)
	fmt.Println()

	// 方法一：快慢指针
	head1 := createList(nums)
	root1 := sortedListToBST(head1)
	fmt.Println("\n方法一（快慢指针）:")
	fmt.Printf("层序遍历: %v\n", levelOrder(root1))
	fmt.Printf("中序遍历: %v\n", inorderTraversal(root1))
	fmt.Printf("是否平衡: %v\n", isBalanced(root1))
	fmt.Printf("是否BST: %v\n", isValidBST(root1))
	fmt.Printf("树高度: %d\n", getHeight(root1))
	fmt.Println("树结构:")
	printTree(root1, "", false)

	// 方法二：转数组
	head2 := createList(nums)
	root2 := sortedListToBST2(head2)
	fmt.Println("\n方法二（转数组）:")
	fmt.Printf("层序遍历: %v\n", levelOrder(root2))
	fmt.Printf("中序遍历: %v\n", inorderTraversal(root2))
	fmt.Printf("是否平衡: %v\n", isBalanced(root2))

	// 方法三：中序遍历（最优解）
	head3 := createList(nums)
	root3 := sortedListToBST3(head3)
	fmt.Println("\n方法三（中序遍历-最优解）:")
	fmt.Printf("层序遍历: %v\n", levelOrder(root3))
	fmt.Printf("中序遍历: %v\n", inorderTraversal(root3))
	fmt.Printf("是否平衡: %v\n", isBalanced(root3))

	// 方法四：计算长度优化
	head4 := createList(nums)
	root4 := sortedListToBST4(head4)
	fmt.Println("\n方法四（计算长度优化）:")
	fmt.Printf("层序遍历: %v\n", levelOrder(root4))
	fmt.Printf("中序遍历: %v\n", inorderTraversal(root4))
	fmt.Printf("是否平衡: %v\n", isBalanced(root4))
}

// ==================== 扩展功能 ====================

// 比较108题和109题的性能差异
func compareWithArray() {
	fmt.Println("\n========== 108题 vs 109题性能对比 ==========")

	sizes := []int{100, 1000, 5000}

	for _, size := range sizes {
		// 生成数据
		nums := make([]int, size)
		for i := 0; i < size; i++ {
			nums[i] = i
		}

		fmt.Printf("\n数据规模: %d\n", size)

		// 108题：数组方式（理论最优）
		root1 := arrayToBST(nums, 0, len(nums)-1)
		fmt.Printf("108题(数组): 树高度=%d, 理论高度=%d\n",
			getHeight(root1), int(math.Ceil(math.Log2(float64(size+1)))))

		// 109题方法一：快慢指针
		head := createList(nums)
		root2 := sortedListToBST(head)
		fmt.Printf("109题(快慢指针): 树高度=%d\n", getHeight(root2))

		// 109题方法三：中序遍历
		head = createList(nums)
		root3 := sortedListToBST3(head)
		fmt.Printf("109题(中序遍历): 树高度=%d\n", getHeight(root3))
	}
}

// 链表中点查找演示
func demonstrateFindMiddle() {
	fmt.Println("\n========== 快慢指针找中点演示 ==========")

	testCases := [][]int{
		{1, 2, 3, 4, 5},    // 奇数个
		{1, 2, 3, 4, 5, 6}, // 偶数个
		{1},                // 单个
		{1, 2},             // 两个
	}

	for _, nums := range testCases {
		head := createList(nums)
		fmt.Print("\n链表: ")
		printList(head)
		fmt.Println()

		// 找中点
		slow, fast := head, head
		var prev *ListNode

		for fast != nil && fast.Next != nil {
			prev = slow
			slow = slow.Next
			fast = fast.Next.Next
		}

		fmt.Printf("中点: %d\n", slow.Val)
		if prev != nil {
			fmt.Printf("中点前一个: %d\n", prev.Val)
		}
	}
}

// 验证所有方法生成的树是否等价
func verifyAllMethods() {
	fmt.Println("\n========== 验证所有方法的等价性 ==========")

	nums := []int{-10, -3, 0, 5, 9}
	head := createList(nums)

	methods := []struct {
		name string
		fn   func(*ListNode) *TreeNode
	}{
		{"快慢指针", sortedListToBST},
		{"转数组", sortedListToBST2},
		{"中序遍历", sortedListToBST3},
		{"计算长度", sortedListToBST4},
	}

	fmt.Print("输入: ")
	printList(head)
	fmt.Println()

	for _, method := range methods {
		h := createList(nums)
		root := method.fn(h)
		inorder := inorderTraversal(root)

		fmt.Printf("\n%s:\n", method.name)
		fmt.Printf("  中序遍历: %v\n", inorder)
		fmt.Printf("  是否平衡: %v\n", isBalanced(root))
		fmt.Printf("  是否BST: %v\n", isValidBST(root))
		fmt.Printf("  树高度: %d\n", getHeight(root))
	}
}

func main() {
	// 测试用例1：示例1
	testCase("测试用例1：基本情况", []int{-10, -3, 0, 5, 9})

	// 测试用例2：空链表
	testCase("测试用例2：空链表", []int{})

	// 测试用例3：单个元素
	testCase("测试用例3：单个元素", []int{1})

	// 测试用例4：两个元素
	testCase("测试用例4：两个元素", []int{1, 3})

	// 测试用例5：奇数个元素
	testCase("测试用例5：奇数个元素", []int{1, 2, 3, 4, 5, 6, 7})

	// 测试用例6：偶数个元素
	testCase("测试用例6：偶数个元素", []int{1, 2, 3, 4, 5, 6})

	// 测试用例7：连续数字
	testCase("测试用例7：连续数字", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	// 扩展功能测试
	compareWithArray()
	demonstrateFindMiddle()
	verifyAllMethods()
}
