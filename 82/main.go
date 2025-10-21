package main

import (
	"fmt"
)

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// =========================== 方法一：双指针算法（最优解法） ===========================

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 创建dummy节点简化边界处理
	dummy := &ListNode{Next: head}
	prev := dummy
	curr := head

	for curr != nil {
		// 检查curr是否重复
		if curr.Next != nil && curr.Val == curr.Next.Val {
			// 记录重复值
			val := curr.Val
			// 跳过所有重复节点
			for curr != nil && curr.Val == val {
				curr = curr.Next
			}
			// 删除重复节点
			prev.Next = curr
		} else {
			// 保留curr节点
			prev = curr
			curr = curr.Next
		}
	}

	return dummy.Next
}

// =========================== 方法二：递归算法 ===========================

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		// 跳过所有重复节点
		val := head.Val
		for head != nil && head.Val == val {
			head = head.Next
		}
		return deleteDuplicates2(head)
	} else {
		// 保留当前节点
		head.Next = deleteDuplicates2(head.Next)
		return head
	}
}

// =========================== 方法三：哈希表 ===========================

func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 统计每个值的出现次数
	count := make(map[int]int)
	curr := head
	for curr != nil {
		count[curr.Val]++
		curr = curr.Next
	}

	// 创建新链表，只保留出现一次的值
	dummy := &ListNode{}
	prev := dummy
	curr = head

	for curr != nil {
		if count[curr.Val] == 1 {
			prev.Next = curr
			prev = prev.Next
		}
		curr = curr.Next
	}
	prev.Next = nil

	return dummy.Next
}

// =========================== 方法四：优化版双指针 ===========================

func deleteDuplicates4(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil {
		curr := prev.Next

		// 检查是否有重复
		if curr.Next != nil && curr.Val == curr.Next.Val {
			val := curr.Val
			// 跳过所有重复节点
			for curr != nil && curr.Val == val {
				curr = curr.Next
			}
			prev.Next = curr
		} else {
			prev = prev.Next
		}
	}

	return dummy.Next
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
	fmt.Println("=== LeetCode 82: 删除排序链表中的重复元素 II ===\n")

	testCases := []struct {
		nums   []int
		expect []int
	}{
		{
			[]int{1, 2, 3, 3, 4, 4, 5},
			[]int{1, 2, 5},
		},
		{
			[]int{1, 1, 1, 2, 3},
			[]int{2, 3},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1, 1, 1, 1, 1},
			[]int{},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 1, 2, 2, 3, 3},
			[]int{},
		},
		{
			[]int{1, 2, 2, 3, 3, 4},
			[]int{1, 4},
		},
	}

	fmt.Println("方法一：双指针算法（最优解法）")
	runTests(testCases, deleteDuplicates)

	fmt.Println("\n方法二：递归算法")
	runTests(testCases, deleteDuplicates2)

	fmt.Println("\n方法三：哈希表")
	runTests(testCases, deleteDuplicates3)

	fmt.Println("\n方法四：优化版双指针")
	runTests(testCases, deleteDuplicates4)
}

func runTests(testCases []struct {
	nums   []int
	expect []int
}, fn func(*ListNode) *ListNode) {
	passCount := 0
	for i, tc := range testCases {
		input := createList(tc.nums)
		expected := createList(tc.expect)
		result := fn(input)

		status := "✅"
		if !compareLists(result, expected) {
			status = "❌"
		} else {
			passCount++
		}
		fmt.Printf("  测试%d: %s\n", i+1, status)
		if status == "❌" {
			fmt.Printf("    输入: %v\n", tc.nums)
			fmt.Printf("    输出: %v\n", listToArray(result))
			fmt.Printf("    期望: %v\n", tc.expect)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}
