package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// swapPairs 迭代法：虚拟头 + 三指针操作
// 时间复杂度: O(n)；空间复杂度: O(1)
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		// 保存当前两个节点
		first := prev.Next
		second := prev.Next.Next

		// 交换操作
		first.Next = second.Next
		second.Next = first
		prev.Next = second

		// 移动到下一组
		prev = first
	}

	return dummy.Next
}

// swapPairsRecursive 递归法：先处理后续，再交换当前两个
func swapPairsRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 保存第二个节点
	second := head.Next
	// 递归处理后续节点
	head.Next = swapPairsRecursive(second.Next)
	// 交换当前两个节点
	second.Next = head

	return second
}

// swapPairsOptimized 优化迭代：减少变量使用
func swapPairsOptimized(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		// 直接交换
		prev.Next.Next, prev.Next.Next.Next, prev.Next =
			prev.Next, prev.Next, prev.Next.Next
		prev = prev.Next.Next
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
	// 示例1: head=[1,2,3,4] -> [2,1,4,3]
	h1 := buildList([]int{1, 2, 3, 4})
	ans1 := swapPairs(h1)
	fmt.Printf("示例1: %v\n", listToSlice(ans1))

	// 示例2: head=[] -> []
	h2 := buildList([]int{})
	ans2 := swapPairs(h2)
	fmt.Printf("示例2: %v\n", listToSlice(ans2))

	// 示例3: head=[1] -> [1]
	h3 := buildList([]int{1})
	ans3 := swapPairs(h3)
	fmt.Printf("示例3: %v\n", listToSlice(ans3))

	// 额外: 奇数个节点 head=[1,2,3] -> [2,1,3]
	h4 := buildList([]int{1, 2, 3})
	ans4 := swapPairs(h4)
	fmt.Printf("额外: %v\n", listToSlice(ans4))

	// 测试递归版本
	h5 := buildList([]int{1, 2, 3, 4, 5})
	ans5 := swapPairsRecursive(h5)
	fmt.Printf("递归版: %v\n", listToSlice(ans5))
}
