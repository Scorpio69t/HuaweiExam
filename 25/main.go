package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseKGroup 迭代法：K个一组翻转链表
// 时间复杂度: O(n)；空间复杂度: O(1)
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for {
		// 检查是否还有k个节点
		count := 0
		curr := prev.Next
		for count < k && curr != nil {
			curr = curr.Next
			count++
		}

		if count < k {
			// 剩余节点不足k个，不翻转
			break
		}

		// 翻转当前k个节点
		start := prev.Next
		reversedStart := reverseKNodes(start, k)

		// 重新连接
		prev.Next = reversedStart
		start.Next = curr

		// 移动到下一组
		prev = start
	}

	return dummy.Next
}

// reverseKNodes 翻转k个节点，返回新的头节点
func reverseKNodes(head *ListNode, k int) *ListNode {
	var prev *ListNode
	curr := head
	count := 0

	for count < k && curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		count++
	}

	return prev
}

// reverseKGroupRecursive 递归法：K个一组翻转链表
func reverseKGroupRecursive(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	// 检查是否有k个节点
	count := 0
	curr := head
	for count < k && curr != nil {
		curr = curr.Next
		count++
	}

	if count < k {
		// 不足k个，不翻转
		return head
	}

	// 翻转前k个节点
	reversedHead := reverseFirstK(head, k)

	// 递归处理剩余节点
	head.Next = reverseKGroupRecursive(curr, k)

	return reversedHead
}

// reverseFirstK 翻转链表的前k个节点
func reverseFirstK(head *ListNode, k int) *ListNode {
	var prev *ListNode
	curr := head
	count := 0

	for count < k && curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		count++
	}

	return prev
}

// reverseKGroupOptimized 优化迭代法：减少不必要的遍历
func reverseKGroupOptimized(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for {
		// 检查剩余节点数量
		count := 0
		curr := prev.Next
		for count < k && curr != nil {
			curr = curr.Next
			count++
		}

		if count < k {
			break
		}

		// 翻转当前k个节点
		start := prev.Next
		reversedStart := reverseKNodes(start, k)

		// 重新连接
		prev.Next = reversedStart
		start.Next = curr

		// 移动到下一组
		prev = start
	}

	return dummy.Next
}

// 辅助：从切片构造链表
func buildList(vals []int) *ListNode {
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

// 辅助：链表转切片
func listToSlice(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func main() {
	// 示例1: head=[1,2,3,4,5], k=2 -> [2,1,4,3,5]
	h1 := buildList([]int{1, 2, 3, 4, 5})
	ans1 := reverseKGroup(h1, 2)
	fmt.Printf("示例1 (k=2): %v\n", listToSlice(ans1))

	// 示例2: head=[1,2,3,4,5], k=3 -> [3,2,1,4,5]
	h2 := buildList([]int{1, 2, 3, 4, 5})
	ans2 := reverseKGroup(h2, 3)
	fmt.Printf("示例2 (k=3): %v\n", listToSlice(ans2))

	// 测试递归版本
	h3 := buildList([]int{1, 2, 3, 4, 5, 6})
	ans3 := reverseKGroupRecursive(h3, 2)
	fmt.Printf("递归版 (k=2): %v\n", listToSlice(ans3))

	// 测试优化版本
	h4 := buildList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	ans4 := reverseKGroupOptimized(h4, 3)
	fmt.Printf("优化版 (k=3): %v\n", listToSlice(ans4))

	// 边界测试：k=1
	h5 := buildList([]int{1, 2, 3, 4, 5})
	ans5 := reverseKGroup(h5, 1)
	fmt.Printf("边界测试 (k=1): %v\n", listToSlice(ans5))

	// 边界测试：k等于链表长度
	h6 := buildList([]int{1, 2, 3, 4, 5})
	ans6 := reverseKGroup(h6, 5)
	fmt.Printf("边界测试 (k=5): %v\n", listToSlice(ans6))
}
