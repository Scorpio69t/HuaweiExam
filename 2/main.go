package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一：模拟加法解法（推荐）
// 时间复杂度：O(max(m,n))，空间复杂度：O(max(m,n))
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0} // 虚拟头节点
	current := dummy
	carry := 0

	// 同时遍历两个链表
	for l1 != nil || l2 != nil || carry > 0 {
		val1, val2 := 0, 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		// 计算当前位的和和进位
		sum := val1 + val2 + carry
		carry = sum / 10

		// 创建新节点
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}

// 方法二：递归解法
// 时间复杂度：O(max(m,n))，空间复杂度：O(max(m,n))
func addTwoNumbersRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersWithCarry(l1, l2, 0)
}

func addTwoNumbersWithCarry(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	// 递归终止条件
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	// 计算当前位的值
	val1, val2 := 0, 0
	if l1 != nil {
		val1 = l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		val2 = l2.Val
		l2 = l2.Next
	}

	// 计算当前位的和和进位
	sum := val1 + val2 + carry
	newCarry := sum / 10
	currentVal := sum % 10

	// 创建当前节点并递归处理下一位
	current := &ListNode{Val: currentVal}
	current.Next = addTwoNumbersWithCarry(l1, l2, newCarry)

	return current
}

// 方法三：转换为数字后相加（仅适用于小规模数据）
// 时间复杂度：O(m+n)，空间复杂度：O(max(m,n))
func addTwoNumbersConvert(l1 *ListNode, l2 *ListNode) *ListNode {
	// 将链表转换为数字
	num1 := listToNumber(l1)
	num2 := listToNumber(l2)

	// 数字相加
	sum := num1 + num2

	// 将结果转换回链表
	return numberToList(sum)
}

// 辅助函数：将链表转换为数字
func listToNumber(head *ListNode) int {
	if head == nil {
		return 0
	}

	var result strings.Builder
	current := head

	// 从链表头部开始构建数字字符串（逆序）
	for current != nil {
		result.WriteString(strconv.Itoa(current.Val))
		current = current.Next
	}

	// 反转字符串得到正确的数字
	numStr := result.String()
	num, _ := strconv.Atoi(reverseString(numStr))
	return num
}

// 辅助函数：将数字转换为链表
func numberToList(num int) *ListNode {
	if num == 0 {
		return &ListNode{Val: 0}
	}

	// 转换为字符串
	numStr := strconv.Itoa(num)
	
	// 创建虚拟头节点
	dummy := &ListNode{Val: 0}
	current := dummy

	// 从右到左（低位到高位）创建节点
	for i := len(numStr) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(numStr[i]))
		current.Next = &ListNode{Val: digit}
		current = current.Next
	}

	return dummy.Next
}

// 辅助函数：反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 方法四：优化的模拟加法解法
func addTwoNumbersOptimized(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	current := dummy
	carry := 0

	// 优化：先处理两个链表都有的部分
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		carry = sum / 10
		
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		
		l1 = l1.Next
		l2 = l2.Next
	}

	// 处理剩余部分
	remaining := l1
	if l2 != nil {
		remaining = l2
	}

	for remaining != nil {
		sum := remaining.Val + carry
		carry = sum / 10
		
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		
		remaining = remaining.Next
	}

	// 处理最后的进位
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}

// 辅助函数：从数组构建链表
func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	dummy := &ListNode{Val: 0}
	current := dummy

	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}

	return dummy.Next
}

// 辅助函数：将链表转换为数组
func listToSlice(head *ListNode) []int {
	var result []int
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

// 辅助函数：打印链表
func printList(head *ListNode, name string) {
	fmt.Printf("%s: %v\n", name, listToSlice(head))
}

// 辅助函数：验证结果是否正确
func validateResult(l1 *ListNode, l2 *ListNode, result *ListNode) bool {
	// 将链表转换为数字进行验证
	num1 := listToNumber(l1)
	num2 := listToNumber(l2)
	expectedSum := num1 + num2
	actualSum := listToNumber(result)

	return expectedSum == actualSum
}

// 辅助函数：比较两个链表是否相等
func compareLists(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func main() {
	fmt.Println("=== 2. 两数相加 ===")

	// 测试用例1
	l1 := buildList([]int{2, 4, 3})
	l2 := buildList([]int{5, 6, 4})
	
	fmt.Printf("测试用例1: l1=[2,4,3], l2=[5,6,4]\n")
	printList(l1, "l1")
	printList(l2, "l2")
	
	result1 := addTwoNumbers(l1, l2)
	fmt.Printf("模拟加法解法结果: %v\n", listToSlice(result1))
	
	result1Recursive := addTwoNumbersRecursive(l1, l2)
	fmt.Printf("递归解法结果: %v\n", listToSlice(result1Recursive))
	
	result1Optimized := addTwoNumbersOptimized(l1, l2)
	fmt.Printf("优化解法结果: %v\n", listToSlice(result1Optimized))
	
	// 验证结果
	if validateResult(l1, l2, result1) {
		fmt.Println("✅ 结果验证通过！")
	} else {
		fmt.Println("❌ 结果验证失败！")
	}
	fmt.Println()

	// 测试用例2
	l3 := buildList([]int{0})
	l4 := buildList([]int{0})
	
	fmt.Printf("测试用例2: l1=[0], l2=[0]\n")
	printList(l3, "l1")
	printList(l4, "l2")
	
	result2 := addTwoNumbers(l3, l4)
	fmt.Printf("模拟加法解法结果: %v\n", listToSlice(result2))
	fmt.Println()

	// 测试用例3
	l5 := buildList([]int{9, 9, 9, 9, 9, 9, 9})
	l6 := buildList([]int{9, 9, 9, 9})
	
	fmt.Printf("测试用例3: l1=[9,9,9,9,9,9,9], l2=[9,9,9,9]\n")
	printList(l5, "l1")
	printList(l6, "l2")
	
	result3 := addTwoNumbers(l5, l6)
	fmt.Printf("模拟加法解法结果: %v\n", listToSlice(result3))
	
	// 验证结果
	if validateResult(l5, l6, result3) {
		fmt.Println("✅ 结果验证通过！")
	} else {
		fmt.Println("❌ 结果验证失败！")
	}
	fmt.Println()

	// 边界测试用例
	testCases := []struct {
		l1   []int
		l2   []int
		desc string
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}, "边界测试1"},
		{[]int{1}, []int{9, 9, 9}, "边界测试2"},
		{[]int{9, 9, 9}, []int{1}, "边界测试3"},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, "边界测试4"},
	}

	for _, tc := range testCases {
		ll1 := buildList(tc.l1)
		ll2 := buildList(tc.l2)
		
		fmt.Printf("%s: l1=%v, l2=%v\n", tc.desc, tc.l1, tc.l2)
		result := addTwoNumbers(ll1, ll2)
		fmt.Printf("结果: %v\n", listToSlice(result))
		
		// 验证结果
		if validateResult(ll1, ll2, result) {
			fmt.Println("✅ 验证通过")
		} else {
			fmt.Println("❌ 验证失败")
		}
		fmt.Println()
	}

	// 算法正确性验证
	fmt.Println("=== 算法正确性验证 ===")
	verifyL1 := buildList([]int{2, 4, 3})
	verifyL2 := buildList([]int{5, 6, 4})
	
	fmt.Printf("验证链表: l1=%v, l2=%v\n", listToSlice(verifyL1), listToSlice(verifyL2))
	
	verifyResult1 := addTwoNumbers(verifyL1, verifyL2)
	verifyResult2 := addTwoNumbersRecursive(verifyL1, verifyL2)
	verifyResult3 := addTwoNumbersOptimized(verifyL1, verifyL2)
	
	fmt.Printf("模拟加法解法: %v\n", listToSlice(verifyResult1))
	fmt.Printf("递归解法: %v\n", listToSlice(verifyResult2))
	fmt.Printf("优化解法: %v\n", listToSlice(verifyResult3))

	// 验证所有解法结果一致
	if compareLists(verifyResult1, verifyResult2) && compareLists(verifyResult2, verifyResult3) {
		fmt.Println("✅ 所有解法结果一致！")
		
		// 验证结果正确性
		if validateResult(verifyL1, verifyL2, verifyResult1) {
			fmt.Println("✅ 结果验证通过！")
		} else {
			fmt.Println("❌ 结果验证失败！")
		}
	} else {
		fmt.Println("❌ 解法结果不一致，需要检查！")
	}

	// 性能测试
	fmt.Println("\n=== 性能测试 ===")
	
	// 创建大链表进行测试
	largeL1 := buildList(make([]int, 1000))
	largeL2 := buildList(make([]int, 1000))
	
	// 填充一些测试数据
	current := largeL1
	for i := 0; i < 1000; i++ {
		current.Val = i % 10
		current = current.Next
	}
	
	current = largeL2
	for i := 0; i < 1000; i++ {
		current.Val = (i + 5) % 10
		current = current.Next
	}

	fmt.Printf("大链表测试: 长度=%d\n", 1000)
	result := addTwoNumbers(largeL1, largeL2)
	fmt.Printf("模拟加法解法结果长度: %d\n", len(listToSlice(result)))
	
	// 验证大链表结果
	if validateResult(largeL1, largeL2, result) {
		fmt.Println("✅ 大链表结果验证通过！")
	} else {
		fmt.Println("❌ 大链表结果验证失败！")
	}
}
