package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd 双指针法：快指针先走n步，然后快慢同步，快到尾时慢在待删前一个
// 时间复杂度: O(sz)；空间复杂度: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	// fast 先走 n 步
	for i := 0; i < n; i++ {
		if fast != nil {
			fast = fast.Next
		}
	}
	// fast 和 slow 同步走，直到 fast 到尾
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// slow.Next 就是待删除结点
	if slow != nil && slow.Next != nil {
		slow.Next = slow.Next.Next
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
	// 示例1: head=[1,2,3,4,5], n=2 -> [1,2,3,5]
	h1 := buildList([]int{1, 2, 3, 4, 5})
	ans1 := removeNthFromEnd(h1, 2)
	fmt.Printf("示例1: %v\n", listToSlice(ans1))

	// 示例2: head=[1], n=1 -> []
	h2 := buildList([]int{1})
	ans2 := removeNthFromEnd(h2, 1)
	fmt.Printf("示例2: %v\n", listToSlice(ans2))

	// 示例3: head=[1,2], n=1 -> [1]
	h3 := buildList([]int{1, 2})
	ans3 := removeNthFromEnd(h3, 1)
	fmt.Printf("示例3: %v\n", listToSlice(ans3))

	// 额外: 删除头结点场景 head=[1,2], n=2 -> [2]
	h4 := buildList([]int{1, 2})
	ans4 := removeNthFromEnd(h4, 2)
	fmt.Printf("额外: %v\n", listToSlice(ans4))
}
