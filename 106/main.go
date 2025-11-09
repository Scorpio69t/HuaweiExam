package main

import (
	"fmt"
	"strings"
)

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =========================== 方法一：递归+哈希表（最优解法） ===========================

// buildTree 从中序与后序遍历序列构造二叉树
// 时间复杂度：O(n)，n为节点数，每个节点访问一次
// 空间复杂度：O(n)，哈希表O(n) + 递归栈O(h)
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 构建哈希表：值 -> 索引，用于快速定位根节点
	indexMap := make(map[int]int)
	for i, val := range inorder {
		indexMap[val] = i
	}

	return helper(inorder, 0, len(inorder)-1,
		postorder, 0, len(postorder)-1, indexMap)
}

// helper 递归辅助函数
func helper(inorder []int, inStart, inEnd int,
	postorder []int, postStart, postEnd int,
	indexMap map[int]int) *TreeNode {
	// 递归终止条件
	if postStart > postEnd {
		return nil
	}

	// 后序遍历最后一个是根节点
	rootVal := postorder[postEnd]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中定位根节点（O(1)查找）
	rootIndex := indexMap[rootVal]

	// 左子树大小
	leftSize := rootIndex - inStart

	// 递归构造左右子树
	// 左子树：中序[inStart, rootIndex-1]，后序[postStart, postStart+leftSize-1]
	root.Left = helper(inorder, inStart, rootIndex-1,
		postorder, postStart, postStart+leftSize-1, indexMap)

	// 右子树：中序[rootIndex+1, inEnd]，后序[postStart+leftSize, postEnd-1]
	root.Right = helper(inorder, rootIndex+1, inEnd,
		postorder, postStart+leftSize, postEnd-1, indexMap)

	return root
}

// =========================== 方法二：递归+切片（简洁版） ===========================

// buildTree2 递归+切片，代码简洁但效率稍低
// 时间复杂度：O(n²)，线性查找O(n) × 递归n次
// 空间复杂度：O(n²)，切片复制导致额外空间
func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	// 根节点（后序最后一个）
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}

	// 在中序中找到根节点位置（线性查找）
	rootIndex := 0
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	// 递归构造左右子树（使用切片，会复制数组）
	// 左子树大小为rootIndex
	root.Left = buildTree2(inorder[:rootIndex],
		postorder[:rootIndex])
	root.Right = buildTree2(inorder[rootIndex+1:],
		postorder[rootIndex:len(postorder)-1])

	return root
}

// =========================== 方法三：迭代+栈（避免递归） ===========================

// buildTree3 迭代+栈实现，从后向前遍历
// 时间复杂度：O(n)
// 空间复杂度：O(n)，栈空间
func buildTree3(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	stack := []*TreeNode{root}
	inorderIndex := len(inorder) - 1

	// 从后向前遍历后序数组
	for i := len(postorder) - 2; i >= 0; i-- {
		node := &TreeNode{Val: postorder[i]}
		parent := stack[len(stack)-1]

		// 当前节点应该是右子节点
		if parent.Val != inorder[inorderIndex] {
			parent.Right = node
		} else {
			// 找到应该作为左子节点的位置
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				parent = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex--
			}
			parent.Left = node
		}

		stack = append(stack, node)
	}

	return root
}

// =========================== 方法四：全局变量优化 ===========================

var postIndex int
var indexMap map[int]int

// buildTree4 使用全局变量优化
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func buildTree4(inorder []int, postorder []int) *TreeNode {
	postIndex = len(postorder) - 1 // 从后向前
	indexMap = make(map[int]int)
	for i, val := range inorder {
		indexMap[val] = i
	}

	return build(postorder, 0, len(inorder)-1)
}

func build(postorder []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	rootVal := postorder[postIndex]
	postIndex-- // 从后向前移动
	root := &TreeNode{Val: rootVal}

	rootIdx := indexMap[rootVal]

	// 注意：后序是左右根，所以要先构造右子树
	root.Right = build(postorder, rootIdx+1, right)
	root.Left = build(postorder, left, rootIdx-1)

	return root
}

// =========================== 辅助函数 ===========================

// treeToArray 将树转换为数组（层序遍历）
func treeToArray(root *TreeNode) []interface{} {
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

	// 移除末尾的nil
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

// printArray 打印数组
func printArray(arr []interface{}) {
	fmt.Print("[")
	for i, v := range arr {
		if i > 0 {
			fmt.Print(",")
		}
		if v == nil {
			fmt.Print("null")
		} else {
			fmt.Print(v)
		}
	}
	fmt.Println("]")
}

// visualizeTree 可视化打印树结构
func visualizeTree(root *TreeNode, prefix string, isLeft bool) {
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
			newPrefix := prefix
			if isLeft {
				newPrefix += "│   "
			} else {
				newPrefix += "    "
			}
			visualizeTree(root.Left, newPrefix, true)
		}

		if root.Right != nil {
			newPrefix := prefix
			if isLeft {
				newPrefix += "│   "
			} else {
				newPrefix += "    "
			}
			visualizeTree(root.Right, newPrefix, false)
		}
	}
}

// preorderTraversal 前序遍历
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{root.Val}
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}

// inorderTraversal 中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := inorderTraversal(root.Left)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}

// postorderTraversal 后序遍历
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := postorderTraversal(root.Left)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}

// treesEqual 判断两棵树是否相等
func treesEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return treesEqual(t1.Left, t2.Left) && treesEqual(t1.Right, t2.Right)
}

// =========================== 扩展功能 ===========================

// getAllTraversals 从树生成所有三种遍历序列
func getAllTraversals(root *TreeNode) (pre, in, post []int) {
	return preorderTraversal(root), inorderTraversal(root), postorderTraversal(root)
}

// validateTraversal 验证中序和后序是否匹配
func validateTraversal(inorder, postorder []int) bool {
	if len(inorder) != len(postorder) {
		return false
	}

	// 检查元素是否相同
	inSet := make(map[int]bool)
	for _, val := range inorder {
		inSet[val] = true
	}

	for _, val := range postorder {
		if !inSet[val] {
			return false
		}
	}

	return true
}

// serialize 将树序列化为中序和后序
func serialize(root *TreeNode) ([]int, []int) {
	return inorderTraversal(root), postorderTraversal(root)
}

// deserialize 反序列化
func deserialize(inorder, postorder []int) *TreeNode {
	return buildTree(inorder, postorder)
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 106: 从中序与后序遍历序列构造二叉树 ===\n")

	// 测试用例
	testCases := []struct {
		name      string
		inorder   []int
		postorder []int
		expectArr []interface{}
	}{
		{
			name:      "示例1: 标准二叉树",
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			expectArr: []interface{}{3, 9, 20, nil, nil, 15, 7},
		},
		{
			name:      "示例2: 单节点",
			inorder:   []int{-1},
			postorder: []int{-1},
			expectArr: []interface{}{-1},
		},
		{
			name:      "左偏树",
			inorder:   []int{3, 2, 1},
			postorder: []int{3, 2, 1},
			expectArr: []interface{}{1, 2, nil, 3},
		},
		{
			name:      "右偏树",
			inorder:   []int{1, 2, 3},
			postorder: []int{1, 2, 3},
			expectArr: []interface{}{1, nil, 2, nil, 3},
		},
		{
			name:      "完全二叉树",
			inorder:   []int{4, 2, 5, 1, 6, 3, 7},
			postorder: []int{4, 5, 2, 6, 7, 3, 1},
			expectArr: []interface{}{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:      "不平衡树",
			inorder:   []int{8, 4, 9, 2, 5, 1, 6, 3, 7},
			postorder: []int{8, 9, 4, 5, 2, 6, 7, 3, 1},
			expectArr: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:      "负数节点",
			inorder:   []int{-2, -1, -3},
			postorder: []int{-2, -3, -1},
			expectArr: []interface{}{-1, -2, -3},
		},
	}

	methods := []struct {
		name string
		fn   func([]int, []int) *TreeNode
	}{
		{"方法一：递归+哈希表", buildTree},
		{"方法二：递归+切片", buildTree2},
		{"方法三：迭代+栈", buildTree3},
		{"方法四：全局变量", buildTree4},
	}

	// 对每种方法运行测试
	for _, method := range methods {
		fmt.Printf("\n%s\n", method.name)
		fmt.Println(strings.Repeat("=", 60))

		passCount := 0
		for i, tc := range testCases {
			root := method.fn(tc.inorder, tc.postorder)
			result := treeToArray(root)

			// 验证结果
			status := "✅"
			if !arraysEqual(result, tc.expectArr) {
				status = "❌"
			} else {
				passCount++
			}

			fmt.Printf("  测试%d: %s\n", i+1, status)
			fmt.Printf("    名称: %s\n", tc.name)
			fmt.Printf("    中序: %v\n", tc.inorder)
			fmt.Printf("    后序: %v\n", tc.postorder)
			fmt.Printf("    输出: ")
			printArray(result)

			if !arraysEqual(result, tc.expectArr) {
				fmt.Printf("    期望: ")
				printArray(tc.expectArr)
			}

			// 为第一个示例打印树结构
			if i == 0 {
				fmt.Println("    树结构:")
				if root != nil {
					visualizeTree(root, "      ", false)
				}
			}

			// 验证遍历结果
			in := inorderTraversal(root)
			post := postorderTraversal(root)
			fmt.Printf("    验证中序: %v (匹配: %v)\n", in, slicesEqual(in, tc.inorder))
			fmt.Printf("    验证后序: %v (匹配: %v)\n", post, slicesEqual(post, tc.postorder))
		}

		fmt.Printf("\n  通过: %d/%d\n", passCount, len(testCases))
	}

	// 扩展功能测试
	fmt.Println("\n\n=== 扩展功能测试 ===\n")
	testExtensions()

	// 与105题对比
	fmt.Println("\n=== 与105题对比 ===\n")
	compareWith105()

	// 性能对比
	fmt.Println("\n=== 性能对比 ===\n")
	performanceTest()
}

// arraysEqual 比较两个interface数组是否相等
func arraysEqual(a, b []interface{}) bool {
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

// slicesEqual 比较两个int切片是否相等
func slicesEqual(a, b []int) bool {
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

// testExtensions 测试扩展功能
func testExtensions() {
	fmt.Println("1. 获取所有遍历序列")
	root := buildTree([]int{4, 2, 5, 1, 6, 3, 7}, []int{4, 5, 2, 6, 7, 3, 1})
	pre, in, post := getAllTraversals(root)
	fmt.Printf("   前序: %v\n", pre)
	fmt.Printf("   中序: %v\n", in)
	fmt.Printf("   后序: %v\n", post)
	fmt.Println("   树结构:")
	visualizeTree(root, "     ", false)

	fmt.Println("\n2. 序列化与反序列化")
	original := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	inSeq, postSeq := serialize(original)
	fmt.Printf("   序列化中序: %v\n", inSeq)
	fmt.Printf("   序列化后序: %v\n", postSeq)
	restored := deserialize(inSeq, postSeq)
	fmt.Printf("   反序列化: ")
	printArray(treeToArray(restored))
	fmt.Printf("   树相等: %v\n", treesEqual(original, restored))

	fmt.Println("\n3. 验证遍历合法性")
	fmt.Printf("   合法: %v\n", validateTraversal([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
	fmt.Printf("   不合法: %v\n", validateTraversal([]int{1, 2, 3}, []int{1, 2}))
}

// compareWith105 与105题对比
func compareWith105() {
	fmt.Println("使用同一棵树，对比前序+中序和中序+后序构造的结果")

	// 构造一棵树
	inorder := []int{4, 2, 5, 1, 6, 3, 7}
	postorder := []int{4, 5, 2, 6, 7, 3, 1}
	tree := buildTree(inorder, postorder)

	// 获取前序
	preorder := preorderTraversal(tree)

	fmt.Printf("  中序遍历: %v\n", inorder)
	fmt.Printf("  前序遍历: %v\n", preorder)
	fmt.Printf("  后序遍历: %v\n", postorder)

	fmt.Println("\n  关键区别:")
	fmt.Println("  - 105题: 前序第一个是根节点，从前往后遍历")
	fmt.Println("  - 106题: 后序最后一个是根节点，从后往前遍历")
	fmt.Println("  - 全局变量法: 105题先左后右，106题必须先右后左")

	fmt.Println("\n  树结构:")
	visualizeTree(tree, "    ", false)
}

// performanceTest 性能测试
func performanceTest() {
	// 构建深度为10的完全二叉树
	size := (1 << 10) - 1 // 2^10 - 1 = 1023个节点

	// 生成完全二叉树的中序和后序
	inorder := make([]int, size)
	postorder := make([]int, size)

	// 中序：左-根-右（升序）
	for i := 0; i < size; i++ {
		inorder[i] = i + 1
	}

	// 后序：左-右-根
	postIdx := 0
	var genPostorder func(int, int)
	genPostorder = func(start, end int) {
		if start > end {
			return
		}
		mid := (start + end) / 2
		genPostorder(start, mid-1)
		genPostorder(mid+1, end)
		postorder[postIdx] = mid
		postIdx++
	}
	genPostorder(1, size)

	fmt.Printf("测试数据：完全二叉树，节点数=%d，深度=10\n\n", size)

	fmt.Println("各方法性能测试:")
	root1 := buildTree(inorder, postorder)
	fmt.Printf("  方法一（递归+哈希表）: 节点数=%d\n", countNodes(root1))

	root2 := buildTree2(inorder, postorder)
	fmt.Printf("  方法二（递归+切片）: 节点数=%d\n", countNodes(root2))

	root3 := buildTree3(inorder, postorder)
	fmt.Printf("  方法三（迭代+栈）: 节点数=%d\n", countNodes(root3))

	root4 := buildTree4(inorder, postorder)
	fmt.Printf("  方法四（全局变量）: 节点数=%d\n", countNodes(root4))

	fmt.Println("\n说明：")
	fmt.Println("  - 方法一（递归+哈希表）：O(n)时间，O(n)空间，最优解法")
	fmt.Println("  - 方法二（递归+切片）：O(n²)时间，O(n²)空间，简洁但低效")
	fmt.Println("  - 方法三（迭代+栈）：O(n)时间，O(n)空间，避免递归")
	fmt.Println("  - 方法四（全局变量）：O(n)时间，O(n)空间，注意构造顺序")
}

// countNodes 计算节点总数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
