package main

import (
	"fmt"
)

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// =========================== 方法一：双链表算法（最优解法） ===========================

func partition(head *ListNode, x int) *ListNode {
	// 创建两个dummy节点
	dummy1 := &ListNode{} // 小于x的节点
	dummy2 := &ListNode{} // 大于等于x的节点

	tail1 := dummy1
	tail2 := dummy2

	curr := head

	for curr != nil {
		if curr.Val < x {
			tail1.Next = curr
			tail1 = tail1.Next
		} else {
			tail2.Next = curr
			tail2 = tail2.Next
		}
		curr = curr.Next
	}

	// 连接两个分区
	tail1.Next = dummy2.Next
	tail2.Next = nil

	return dummy1.Next
}

// =========================== 方法二：数组算法 ===========================

func partition2(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	// 收集所有节点值
	var values []int
	curr := head
	for curr != nil {
		values = append(values, curr.Val)
		curr = curr.Next
	}

	// 重新排列
	var less, greater []int
	for _, val := range values {
		if val < x {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}

	// 重建链表
	allValues := append(less, greater...)
	if len(allValues) == 0 {
		return nil
	}

	newHead := &ListNode{Val: allValues[0]}
	curr = newHead
	for i := 1; i < len(allValues); i++ {
		curr.Next = &ListNode{Val: allValues[i]}
		curr = curr.Next
	}

	return newHead
}

// =========================== 方法三：双链表算法（简化版） ===========================

func partition3(head *ListNode, x int) *ListNode {
	// 创建两个dummy节点
	less := &ListNode{}
	greater := &ListNode{}

	lessPtr, greaterPtr := less, greater

	for head != nil {
		if head.Val < x {
			lessPtr.Next = head
			lessPtr = lessPtr.Next
		} else {
			greaterPtr.Next = head
			greaterPtr = greaterPtr.Next
		}
		head = head.Next
	}

	greaterPtr.Next = nil
	lessPtr.Next = greater.Next

	return less.Next
}

// =========================== 方法四：优化版双链表 ===========================

func partition4(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	// 创建两个dummy节点
	dummy1 := &ListNode{}
	dummy2 := &ListNode{}

	tail1, tail2 := dummy1, dummy2

	for head != nil {
		if head.Val < x {
			tail1.Next = head
			tail1 = tail1.Next
		} else {
			tail2.Next = head
			tail2 = tail2.Next
		}
		head = head.Next
	}

	// 连接两个分区
	tail1.Next = dummy2.Next
	tail2.Next = nil

	return dummy1.Next
}

// =========================== 辅助函数 ===========================

// 创建链表
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	curr := head

	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}

	return head
}

// 链表转数组
func listToArray(head *ListNode) []int {
	var result []int
	curr := head

	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}

	return result
}

// 比较两个链表是否相等
func compareLists(l1, l2 *ListNode) bool {
	curr1, curr2 := l1, l2

	for curr1 != nil && curr2 != nil {
		if curr1.Val != curr2.Val {
			return false
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	return curr1 == nil && curr2 == nil
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 86: 分隔链表 ===\n")

	testCases := []struct {
		nums   []int
		x      int
		expect []int
	}{
		{
			[]int{1, 4, 3, 2, 5, 2},
			3,
			[]int{1, 2, 2, 4, 3, 5},
		},
		{
			[]int{2, 1},
			2,
			[]int{1, 2},
		},
		{
			[]int{1},
			0,
			[]int{1},
		},
		{
			[]int{},
			0,
			[]int{},
		},
		{
			[]int{1, 2, 3},
			5,
			[]int{1, 2, 3},
		},
		{
			[]int{4, 5, 6},
			3,
			[]int{4, 5, 6},
		},
		{
			[]int{3, 1, 2},
			3,
			[]int{1, 2, 3},
		},
		{
			[]int{1, 4, 3, 0, 2, 5, 2},
			3,
			[]int{1, 0, 2, 2, 4, 3, 5},
		},
	}

	fmt.Println("方法一：双链表算法（最优解法）")
	runTests(testCases, partition)

	fmt.Println("\n方法二：数组算法")
	runTests(testCases, partition2)

	fmt.Println("\n方法三：双链表算法（简化版）")
	runTests(testCases, partition3)

	fmt.Println("\n方法四：优化版双链表")
	runTests(testCases, partition4)
}

func runTests(testCases []struct {
	nums   []int
	x      int
	expect []int
}, fn func(*ListNode, int) *ListNode) {
	passCount := 0
	for i, tc := range testCases {
		input := createList(tc.nums)
		expected := createList(tc.expect)
		result := fn(input, tc.x)

		status := "✅"
		if !compareLists(result, expected) {
			status = "❌"
		} else {
			passCount++
		}
		fmt.Printf("  测试%d: %s\n", i+1, status)
		if status == "❌" {
			fmt.Printf("    输入: %v, x=%d\n", tc.nums, tc.x)
			fmt.Printf("    输出: %v\n", listToArray(result))
			fmt.Printf("    期望: %v\n", tc.expect)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}
