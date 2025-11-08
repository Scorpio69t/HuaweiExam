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

// buildTree 从前序与中序遍历序列构造二叉树
// 时间复杂度：O(n)，n为节点数，每个节点访问一次
// 空间复杂度：O(n)，哈希表O(n) + 递归栈O(h)
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 构建哈希表：值 -> 索引，用于快速定位根节点
	indexMap := make(map[int]int)
	for i, val := range inorder {
		indexMap[val] = i
	}

	return helper(preorder, 0, len(preorder)-1,
		inorder, 0, len(inorder)-1, indexMap)
}

// helper 递归辅助函数
func helper(preorder []int, preStart, preEnd int,
	inorder []int, inStart, inEnd int,
	indexMap map[int]int) *TreeNode {
	// 递归终止条件
	if preStart > preEnd {
		return nil
	}

	// 前序遍历第一个是根节点
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中定位根节点（O(1)查找）
	rootIndex := indexMap[rootVal]

	// 左子树大小
	leftSize := rootIndex - inStart

	// 递归构造左右子树
	// 左子树：前序[preStart+1, preStart+leftSize]，中序[inStart, rootIndex-1]
	root.Left = helper(preorder, preStart+1, preStart+leftSize,
		inorder, inStart, rootIndex-1, indexMap)

	// 右子树：前序[preStart+leftSize+1, preEnd]，中序[rootIndex+1, inEnd]
	root.Right = helper(preorder, preStart+leftSize+1, preEnd,
		inorder, rootIndex+1, inEnd, indexMap)

	return root
}

// =========================== 方法二：递归+切片（简洁版） ===========================

// buildTree2 递归+切片，代码简洁但效率稍低
// 时间复杂度：O(n²)，线性查找O(n) × 递归n次
// 空间复杂度：O(n²)，切片复制导致额外空间
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// 根节点
	rootVal := preorder[0]
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
	root.Left = buildTree2(preorder[1:rootIndex+1],
		inorder[:rootIndex])
	root.Right = buildTree2(preorder[rootIndex+1:],
		inorder[rootIndex+1:])

	return root
}

// =========================== 方法三：迭代+栈（避免递归） ===========================

// buildTree3 迭代+栈实现
// 时间复杂度：O(n)
// 空间复杂度：O(n)，栈空间
func buildTree3(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{root}
	inorderIndex := 0

	for i := 1; i < len(preorder); i++ {
		node := &TreeNode{Val: preorder[i]}
		parent := stack[len(stack)-1]

		// 当前节点应该是左子节点
		if parent.Val != inorder[inorderIndex] {
			parent.Left = node
		} else {
			// 找到应该作为右子节点的位置
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				parent = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			parent.Right = node
		}

		stack = append(stack, node)
	}

	return root
}

// =========================== 方法四：全局变量优化 ===========================

var preIndex int
var indexMap map[int]int

// buildTree4 使用全局变量优化
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func buildTree4(preorder []int, inorder []int) *TreeNode {
	preIndex = 0
	indexMap = make(map[int]int)
	for i, val := range inorder {
		indexMap[val] = i
	}

	return build(preorder, 0, len(inorder)-1)
}

func build(preorder []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	rootVal := preorder[preIndex]
	preIndex++
	root := &TreeNode{Val: rootVal}

	rootIdx := indexMap[rootVal]

	// 注意：必须先构造左子树
	root.Left = build(preorder, left, rootIdx-1)
	root.Right = build(preorder, rootIdx+1, right)

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

// buildTreeInPost 从中序与后序遍历序列构造二叉树
func buildTreeInPost(inorder []int, postorder []int) *TreeNode {
	indexMap := make(map[int]int)
	for i, val := range inorder {
		indexMap[val] = i
	}

	return helperInPost(inorder, 0, len(inorder)-1,
		postorder, 0, len(postorder)-1, indexMap)
}

func helperInPost(inorder []int, inStart, inEnd int,
	postorder []int, postStart, postEnd int,
	indexMap map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}

	// 后序最后一个是根节点
	rootVal := postorder[postEnd]
	root := &TreeNode{Val: rootVal}

	rootIndex := indexMap[rootVal]
	leftSize := rootIndex - inStart

	// 注意：后序是左右根，所以左子树在前
	root.Left = helperInPost(inorder, inStart, rootIndex-1,
		postorder, postStart, postStart+leftSize-1, indexMap)
	root.Right = helperInPost(inorder, rootIndex+1, inEnd,
		postorder, postStart+leftSize, postEnd-1, indexMap)

	return root
}

// serialize 将树序列化为前序和中序
func serialize(root *TreeNode) ([]int, []int) {
	return preorderTraversal(root), inorderTraversal(root)
}

// deserialize 反序列化
func deserialize(preorder, inorder []int) *TreeNode {
	return buildTree(preorder, inorder)
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 105: 从前序与中序遍历序列构造二叉树 ===\n")

	// 测试用例
	testCases := []struct {
		name      string
		preorder  []int
		inorder   []int
		expectArr []interface{}
	}{
		{
			name:      "示例1: 标准二叉树",
			preorder:  []int{3, 9, 20, 15, 7},
			inorder:   []int{9, 3, 15, 20, 7},
			expectArr: []interface{}{3, 9, 20, nil, nil, 15, 7},
		},
		{
			name:      "示例2: 单节点",
			preorder:  []int{-1},
			inorder:   []int{-1},
			expectArr: []interface{}{-1},
		},
		{
			name:      "左偏树",
			preorder:  []int{1, 2, 3},
			inorder:   []int{3, 2, 1},
			expectArr: []interface{}{1, 2, nil, 3},
		},
		{
			name:      "右偏树",
			preorder:  []int{1, 2, 3},
			inorder:   []int{1, 2, 3},
			expectArr: []interface{}{1, nil, 2, nil, 3},
		},
		{
			name:      "完全二叉树",
			preorder:  []int{1, 2, 4, 5, 3, 6, 7},
			inorder:   []int{4, 2, 5, 1, 6, 3, 7},
			expectArr: []interface{}{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:      "不平衡树",
			preorder:  []int{1, 2, 4, 8, 9, 5, 3, 6, 7},
			inorder:   []int{8, 4, 9, 2, 5, 1, 6, 3, 7},
			expectArr: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:      "负数节点",
			preorder:  []int{-1, -2, -3},
			inorder:   []int{-2, -1, -3},
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
			root := method.fn(tc.preorder, tc.inorder)
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
			fmt.Printf("    前序: %v\n", tc.preorder)
			fmt.Printf("    中序: %v\n", tc.inorder)
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
			pre := preorderTraversal(root)
			in := inorderTraversal(root)
			fmt.Printf("    验证前序: %v (匹配: %v)\n", pre, slicesEqual(pre, tc.preorder))
			fmt.Printf("    验证中序: %v (匹配: %v)\n", in, slicesEqual(in, tc.inorder))
		}

		fmt.Printf("\n  通过: %d/%d\n", passCount, len(testCases))
	}

	// 扩展功能测试
	fmt.Println("\n\n=== 扩展功能测试 ===\n")
	testExtensions()

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
	fmt.Println("1. 从中序和后序构造二叉树")
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTreeInPost(inorder, postorder)
	fmt.Printf("   中序: %v\n", inorder)
	fmt.Printf("   后序: %v\n", postorder)
	fmt.Printf("   结果: ")
	printArray(treeToArray(root))
	fmt.Println("   树结构:")
	visualizeTree(root, "     ", false)

	fmt.Println("\n2. 序列化与反序列化")
	original := buildTree([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7})
	pre, in := serialize(original)
	fmt.Printf("   序列化前序: %v\n", pre)
	fmt.Printf("   序列化中序: %v\n", in)
	restored := deserialize(pre, in)
	fmt.Printf("   反序列化: ")
	printArray(treeToArray(restored))
	fmt.Printf("   树相等: %v\n", treesEqual(original, restored))

	fmt.Println("\n3. 验证构造正确性")
	testPre := []int{3, 9, 20, 15, 7}
	testIn := []int{9, 3, 15, 20, 7}
	testRoot := buildTree(testPre, testIn)
	verifyPre := preorderTraversal(testRoot)
	verifyIn := inorderTraversal(testRoot)
	fmt.Printf("   原始前序: %v\n", testPre)
	fmt.Printf("   验证前序: %v\n", verifyPre)
	fmt.Printf("   匹配: %v\n", slicesEqual(testPre, verifyPre))
	fmt.Printf("   原始中序: %v\n", testIn)
	fmt.Printf("   验证中序: %v\n", verifyIn)
	fmt.Printf("   匹配: %v\n", slicesEqual(testIn, verifyIn))
}

// performanceTest 性能测试
func performanceTest() {
	// 构建深度为10的完全二叉树
	size := (1 << 10) - 1 // 2^10 - 1 = 1023个节点

	// 生成完全二叉树的前序和中序
	preorder := make([]int, size)
	inorder := make([]int, size)

	// 前序：根-左-右
	preIdx := 0
	var genPreorder func(int, int)
	genPreorder = func(start, end int) {
		if start > end {
			return
		}
		mid := (start + end) / 2
		preorder[preIdx] = mid
		preIdx++
		genPreorder(start, mid-1)
		genPreorder(mid+1, end)
	}
	genPreorder(1, size)

	// 中序：左-根-右（就是升序）
	for i := 0; i < size; i++ {
		inorder[i] = i + 1
	}

	fmt.Printf("测试数据：完全二叉树，节点数=%d，深度=10\n\n", size)

	fmt.Println("各方法性能测试:")
	root1 := buildTree(preorder, inorder)
	fmt.Printf("  方法一（递归+哈希表）: 节点数=%d\n", countNodes(root1))

	root2 := buildTree2(preorder, inorder)
	fmt.Printf("  方法二（递归+切片）: 节点数=%d\n", countNodes(root2))

	root3 := buildTree3(preorder, inorder)
	fmt.Printf("  方法三（迭代+栈）: 节点数=%d\n", countNodes(root3))

	root4 := buildTree4(preorder, inorder)
	fmt.Printf("  方法四（全局变量）: 节点数=%d\n", countNodes(root4))

	fmt.Println("\n说明：")
	fmt.Println("  - 方法一（递归+哈希表）：O(n)时间，O(n)空间，最优解法")
	fmt.Println("  - 方法二（递归+切片）：O(n²)时间，O(n²)空间，简洁但低效")
	fmt.Println("  - 方法三（迭代+栈）：O(n)时间，O(n)空间，避免递归")
	fmt.Println("  - 方法四（全局变量）：O(n)时间，O(n)空间，简化参数")
}

// countNodes 计算节点总数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
